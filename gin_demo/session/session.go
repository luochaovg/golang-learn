package session

type Session interface {
	// 获取Session对象的ID
	ID() string
	Set(key string, value interface{}) error
	Get(key string) (interface{}, error)
	Del(key string) error
	Save() error
	// 设置Redis数据过期时间,内存版本无效
	SetExpired(int)
}
