package taillog

import (
	"fmt"
	"github.com/nxadm/tail"
	"time"
)

var (
	tails *tail.Tail
)

func Init(fileName string) (err error) {
	cnf := tail.Config{
		ReOpen:    true,
		Follow:    true,
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2},
		MustExist: false,
		Poll:      true,
	}

	tails, err = tail.TailFile(fileName, cnf)

	if err != nil {
		fmt.Println("tail file failer, err", err)
		return
	}
	return
}

func ReadLog() {
	var (
		line *tail.Line
		ok   bool
	)

	for {
		line, ok = <-tails.Lines

		if !ok {
			fmt.Printf("taile file close re open , filename:%s \n", tails.Filename)
			time.Sleep(time.Second)
			continue
		}

		fmt.Println("Line is :", line.Text)
	}
}

func ReadChan() <-chan *tail.Line {

	return tails.Lines
}
