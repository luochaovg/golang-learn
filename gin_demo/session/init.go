package session

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

// 中间件让用户选择使用那个版本

var (
	sessionMgr SessionMgr
)

type SessionMgrType string

const (
	// SessionID在cookie里面的名字
	SessionCookieName = "session_id"
	// Session对象在Context里面的名字
	SessionContextName                = "session"
	Memory             SessionMgrType = "memory"
	Redis              SessionMgrType = "redis"
)

// Options Cookie对应的相关选项
type Options struct {
	Path   string
	Domain string
	// Cookie中的SessionID存活时间
	// MaxAge=0 means no 'Max-Age' attribute specified.
	// MaxAge<0 means delete cookie now, equivalently 'Max-Age: 0'.
	// MaxAge>0 means Max-Age attribute present and given in seconds.
	MaxAge   int
	Secure   bool
	HttpOnly bool
}

func Init(provider string, addr string, options ...string) (sm SessionMgr, err error) {
	switch provider {
	case "memory":
		sessionMgr = NewMemorySessionMgr()
	case "redis":
		sessionMgr = NewRedisSessionMgr()
	default:
		fmt.Errorf("不支持")

	}

	err = sessionMgr.Init(addr, options...)

	return
}

func SessionMiddleware(sm SessionMgr, options Options) gin.HandlerFunc {
	return func(c *gin.Context) {
		var session Session
		// 尝试从cookie获取session ID
		sessionID, err := c.Cookie(SessionCookieName)
		if err != nil {
			log.Printf("get session_id from cookie failed, err:%v\n", err)
			session, err = sm.CreateSession()
			sessionID = session.ID()
		} else {
			log.Printf("SessionId: %v\n", sessionID)
			session, err = sm.Get(sessionID)
			if err != nil {
				log.Printf("Get session by %s failed, err: %v\n", sessionID, err)
				session, err = sm.CreateSession()
				sessionID = session.ID()
			}
		}

		session.SetExpired(options.MaxAge)
		c.Set(SessionContextName, session)
		c.SetCookie(SessionCookieName, sessionID, options.MaxAge, options.Path, options.Domain, options.Secure, options.HttpOnly)
		//defer sm.Clear(sessionID)
		c.Next()
	}
}
