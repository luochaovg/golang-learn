package session

import (
	"errors"
	"sync"
)

// 对象
// 定义MemorySession 对象， （sessionId , 存kv的map, 读写锁）
// 构造函数
// Set()
// Get()
// Del()
// Save()

type MemorySession struct {
	sessionId string
	// 存kv
	data   map[string]interface{}
	rwlock sync.RWMutex
	// session过期时间
	expired int
}

// 构造函数
func NewMemorySession(id string) *MemorySession {
	return &MemorySession{
		sessionId: id,
		data:      make(map[string]interface{}, 16),
	}
}

func (m *MemorySession) ID() string {
	return m.sessionId
}

func (m *MemorySession) Set(key string, value interface{}) (err error) {
	// 加锁
	m.rwlock.Lock()
	defer m.rwlock.Unlock()

	// 设置值
	m.data[key] = value
	return
}

func (m *MemorySession) Get(key string) (value interface{}, err error) {
	// 加锁
	m.rwlock.Lock()
	defer m.rwlock.Unlock()

	value, ok := m.data[key]
	if !ok {
		err = errors.New("key not exists in session")
	}

	return
}

func (m *MemorySession) Del(key string) (err error) {
	// 加锁
	m.rwlock.Lock()
	defer m.rwlock.Unlock()

	delete(m.data, key)

	return
}

func (m *MemorySession) Save() (err error) {
	return
}

func (m *MemorySession) SetExpired(expired int) {
	m.expired = expired
}
