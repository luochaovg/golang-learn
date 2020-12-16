package session

import (
	"errors"
	uuid "github.com/satori/go.uuid"
	"sync"
)

// MemorySessionMgr 设计
// 定义MemorySessionMgr 对象， （字段：存放所有session的map, 读写锁）
// 构造函数
// Init()
// CreateSession()
// GetSession

type MemorySessionMgr struct {
	sessionMap map[string]Session
	rwlock     sync.RWMutex
}

func NewMemorySessionMgr() *MemorySessionMgr {
	return &MemorySessionMgr{
		sessionMap: make(map[string]Session, 1024),
	}
}

func (s *MemorySessionMgr) Init(addr string, options ...string) (err error) {
	return
}

// 创建session
func (s *MemorySessionMgr) CreateSession() (session Session, err error) {
	s.rwlock.Lock()
	defer s.rwlock.Unlock()
	// go get github.com/satori/go.uuid
	// 用uuid 作为sessionId
	id := uuid.NewV4()

	// 转为string
	sessionId := id.String()

	// 创建session
	session = NewMemorySession(sessionId)

	// 加入到大map
	s.sessionMap[sessionId] = session
	return
}

func (s *MemorySessionMgr) Get(sessionId string) (session Session, err error) {
	s.rwlock.Lock()
	defer s.rwlock.Unlock()

	session, ok := s.sessionMap[sessionId]
	if !ok {
		err = errors.New("session not exists")
	}
	return
}
