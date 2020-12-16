package session

import (
	"errors"
	"github.com/garyburd/redigo/redis"
	uuid "github.com/satori/go.uuid"
	"sync"
	"time"
)

// RedisSessionMgr 设计
// 定义RedisSessionMgr 对象， （字段：存放所有session的map, 读写锁）
// 构造函数
// Init()
// CreateSession()
// GetSession

type RedisSessionMgr struct {
	// redis 地址
	addr string
	// 密码
	passwd string
	// 连接池
	pool *redis.Pool
	// 锁
	rwlock sync.RWMutex
	// 大map
	sessionMap map[string]Session
}

// 构造
func NewRedisSessionMgr() *RedisSessionMgr {
	return &RedisSessionMgr{
		sessionMap: make(map[string]Session, 32),
	}
}

func (r *RedisSessionMgr) Init(addr string, options ...string) (err error) {
	if len(options) > 0 {
		r.passwd = options[0]
	}

	// 创建连接池
	r.pool = RedisPool(addr, r.passwd)
	r.addr = addr

	return
}

func RedisPool(addr, passwd string) *redis.Pool {
	return &redis.Pool{
		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial("tcp", addr)
			if err != nil {
				return nil, err
			}

			// 若有密码， 判断
			if _, err := conn.Do("AUTH", passwd); err != nil {
				conn.Close()
				return nil, err
			}
			return conn, err
		},
		MaxIdle:     64,
		MaxActive:   1000,
		IdleTimeout: time.Second * 4,
		// 连接测试，开发时写， 上线时注释掉
		TestOnBorrow: func(conn redis.Conn, t time.Time) error {
			_, err := conn.Do("PING")
			return err
		},
	}
}

// 创建session
func (r *RedisSessionMgr) CreateSession() (session Session, err error) {
	r.rwlock.Lock()
	defer r.rwlock.Unlock()

	id := uuid.NewV4()

	// 转为string
	sessionId := id.String()

	// 创建session
	session = NewRedisSession(sessionId, r.pool)

	// 加入到大map
	r.sessionMap[sessionId] = session
	return
}

func (r *RedisSessionMgr) Get(sessionId string) (session Session, err error) {
	r.rwlock.Lock()
	defer r.rwlock.Unlock()

	session, ok := r.sessionMap[sessionId]
	if !ok {
		err = errors.New("session not exists")
	}
	return
}
