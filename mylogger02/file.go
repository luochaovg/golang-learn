package mylogger02

import (
	"fmt"
	"os"
	"path"
	"time"
)

//往文件里面写日志
// 切割：
// 按文件大小 每次记录之前判断文件大小
// 按日期：
// 1.在日志结构体中设置一个字段记录上一次切割时间
// 2.在写日志之前检查一下当前时间的小时数和之前保存的是否一致，不一致就切割

// FileLogger 文件日志结构体
type FileLogger struct {
	Level       LogLevel
	filePath    string // 日志文件保存路径
	fileName    string // 日志文件保存的文件名
	fileObj     *os.File
	errFileObj  *os.File
	maxFileSize int64
	logChan     chan *logMsg
}

type logMsg struct {
	level     LogLevel
	msg       string
	funcName  string
	fileName  string
	timestamp string
	line      int
}

// 日志通道缓冲区大小
var chanSize int = 5000

// NewFileLogger 构造函数
func NewFileLogger(levelStr, fp, fn string, maxSize int64) *FileLogger {
	LogLevel, err := parseLogLevel(levelStr)
	if err != nil {
		panic(err)
	}

	fl := &FileLogger{
		Level:       LogLevel,
		filePath:    fp,
		fileName:    fn,
		maxFileSize: maxSize,
		logChan:     make(chan *logMsg, chanSize),
	}

	err = fl.initFile() // 按照文件路径和文件名将文件打开
	if err != nil {
		panic(err)
	}
	return fl
}

// 根据指定的日志文件路径和文件名打开文件
func (f *FileLogger) initFile() (err error) {
	fullFileName := path.Join(f.filePath, f.fileName)
	fileObj, err := os.OpenFile(fullFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open log file failed , err:%v", err)
		return err
	}

	errFileObj, err := os.OpenFile(fullFileName+".error", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open err log file failed , err:%v", err)
		return err
	}

	// 日志文件都已经打开
	f.fileObj = fileObj
	f.errFileObj = errFileObj

	// 开启一个后台的goroutine写日志
	go f.writeLogBackground()

	return nil
}

// 判断文件是否需要切割
func (f *FileLogger) checkSize(file *os.File) bool {
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Printf("get file err %v", err)
		return false
	}

	// 如果当前文件大于等于日志文件最大值
	return fileInfo.Size() > f.maxFileSize

}

// 判断是否需要记录该日志
func (f *FileLogger) enable(loglevel LogLevel) bool {
	return loglevel >= f.Level
}

func (f *FileLogger) writeLogBackground() {
	for {
		if f.checkSize(f.fileObj) {
			newFile, err := f.splitFile(f.fileObj)
			if err != nil {
				return
			}
			f.fileObj = newFile
		}

		select {
		case logTmp := <-f.logChan:
			// 把日志并出来
			logInfo := fmt.Sprintf("[%s][%s][%s:%s:%d] %s \n", logTmp.timestamp, getLogString(logTmp.level), logTmp.fileName, logTmp.funcName, logTmp.line, logTmp.msg)
			fmt.Fprintf(f.fileObj, logInfo)

			// 如果要记录的日志大于等于ERROR我还要到errorlog 日志文件中在记录以下
			if logTmp.level >= ERROR {
				if f.checkSize(f.errFileObj) {
					newFile, err := f.splitFile(f.errFileObj)
					if err != nil {
						return
					}
					f.errFileObj = newFile
				}

				fmt.Fprintf(f.errFileObj, logInfo)
			}
		default:
			// 取不到日志休息500ms
			time.Sleep(time.Microsecond * 500)
		}

	}
}

// 记录日志方法
func (f *FileLogger) log(lv LogLevel, format string, a ...interface{}) {
	if f.enable(lv) {
		msg := fmt.Sprintf(format, a...)
		now := time.Now()
		funcName, fileName, lineNo := getInfo(3)

		// 先把日志发送到通道中
		// 1.造一个logMsg 对象
		logTmp := &logMsg{
			level:     lv,
			msg:       msg,
			funcName:  funcName,
			fileName:  fileName,
			timestamp: now.Format("2006-01-02 15:04:05"),
			line:      lineNo,
		}

		// 防止堵塞
		select {
		case f.logChan <- logTmp:
		default: // 如果堵塞，把日志丢掉保证不出现堵塞
		}
	}
}

// 切割文件
func (f *FileLogger) splitFile(file *os.File) (*os.File, error) {
	// 需要切割
	nowStr := time.Now().Format("20060102150405")
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Printf("get file info failed , err:%v\n", err)
		return nil, err
	}

	logName := path.Join(f.filePath, fileInfo.Name())      // 拿到当前的日志文件完整路径
	newLogName := fmt.Sprintf("%s.bak%s", logName, nowStr) // 拼接一个日志文件备份的名字

	// 1.关闭当前文件
	file.Close()
	// 2.rename 备份一下  xx.log -> xx.log.bak2019092934
	os.Rename(logName, newLogName)

	// 3.打开一个新的文件
	fileObj, err := os.OpenFile(logName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open new log file failed , err:%v\n", err)
		return nil, err
	}

	// 4.直接返回打开的文件对象
	return fileObj, nil
}

// Debug ...
func (f *FileLogger) Debug(format string, a ...interface{}) {
	f.log(DEBUG, format, a...)
}

// Trace ...
func (f *FileLogger) Trace(format string, a ...interface{}) {
	f.log(TRACE, format, a...)
}

// Info ...
func (f *FileLogger) Info(format string, a ...interface{}) {
	f.log(INFO, format, a...)
}

// Warning ...
func (f *FileLogger) Warning(format string, a ...interface{}) {
	f.log(WARNING, format, a...)
}

// Error ...
func (f *FileLogger) Error(format string, a ...interface{}) {
	f.log(ERROR, format, a...)
}

// Fatal ...
func (f *FileLogger) Fatal(format string, a ...interface{}) {
	f.log(FATAL, format, a...)
}

// Close ...
func (f *FileLogger) Close() {
	f.fileObj.Close()
	f.errFileObj.Close()
}
