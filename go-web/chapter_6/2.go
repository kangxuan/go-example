package main

import (
	"encoding/base64"
	"fmt"
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"sync"
	"time"
)

// Manager 定义全局的session管理器
type Manager struct {
	cookieName  string     // 私有cookieName
	lock        sync.Mutex // 锁
	provider    Provider   // 提供器
	maxLifeTime int64      // 过期时间
}

// sessionId 生成全局唯一的sessionID
func (manager *Manager) sessionId() string {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		return ""
	}
	return base64.URLEncoding.EncodeToString(b)
}

// SessionStart 检测是否已经有某个 Session 与当前来访用户发生了关联，如果没有则创建之
func (manager *Manager) SessionStart(w http.ResponseWriter, r *http.Request) (session Session) {
	manager.lock.Lock()         // 加锁
	defer manager.lock.Unlock() // 解锁

	cookie, err := r.Cookie(manager.cookieName) // 获取session
	if err != nil || cookie.Value == "" {
		sid := manager.sessionId() // 生成sessionId
		session, _ = manager.provider.SessionInit(sid)
		// 种下cookie
		cookie := http.Cookie{Name: manager.cookieName, Value: url.QueryEscape(sid), Path: "/", HttpOnly: true, MaxAge: int(manager.maxLifeTime)}
		http.SetCookie(w, &cookie)
	} else {
		sid, _ := url.QueryUnescape(cookie.Value)
		session, _ = manager.provider.SessionRead(sid)
	}
	return
}

// SessionDestroy 销毁session
func (manager *Manager) SessionDestroy(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie(manager.cookieName)
	if err != nil || cookie.Value == "" { // 如果没有则不需要销毁
		return
	} else {
		manager.lock.Lock()
		defer manager.lock.Unlock()

		_ = manager.provider.SessionDestroy(cookie.Value) // 销毁数据

		// 浏览器的cookie进行失效
		expiration := time.Now()
		cookie := http.Cookie{Name: manager.cookieName, Path: "/", HttpOnly: true, Expires: expiration, MaxAge: -1}
		http.SetCookie(w, &cookie)
	}
}

// GC 过期session失效
func (manager *Manager) GC() {
	manager.lock.Lock()
	defer manager.lock.Unlock()

	manager.provider.SessionGC(manager.maxLifeTime)
	time.AfterFunc(time.Duration(manager.maxLifeTime), func() {
		manager.GC()
	})
}

// Provider 提供器接口
type Provider interface {
	SessionInit(sid string) (Session, error) // 初始化session，成功并返回新的session变量
	SessionRead(sid string) (Session, error) // 返回sid的session变量
	SessionDestroy(sid string) error         // 销毁sid对应的session变量
	SessionGC(maxLifeTime int64)             // 根据maxLifeTime来删除过期的数据
}

// Session 会话接口
type Session interface {
	Set(key, value interface{}) error // 设置session的值
	Get(key interface{}) interface{}  // 获取session的值
	Delete(key interface{}) error     // 删除session的值
	SessionID() string                // 返回当前的sessionId
}

var provides = make(map[string]Provider)

// Register 注册
func Register(name string, provider Provider) {
	if provider == nil {
		panic("session: Register provider is nil")
	}
	if _, dup := provides[name]; dup {
		panic("session: Register called twice for provider " + name)
	}
	provides[name] = provider
}

// NewManager 定义一个全局的session管理器
func NewManager(provideName, cookieName string, maxLifeTime int64) (*Manager, error) {
	provider, ok := provides[provideName]
	if !ok {
		return nil, fmt.Errorf("session: unknown provide %q（forgotten import?）", provideName)
	}
	return &Manager{provider: provider, cookieName: cookieName, maxLifeTime: maxLifeTime}, nil
}

// ------------------------------------------------------------------------------------------------------//

var globalSessions *Manager

// init 初始化
func init() {
	globalSessions, _ = NewManager("memory", "gosessionid", 3600)
	log.Println("init")
	//go globalSessions.GC()
}

func login(w http.ResponseWriter, r *http.Request) {
	log.Println(globalSessions)
	sess := globalSessions.SessionStart(w, r)
	err := r.ParseForm()
	if err != nil {
		return
	}

	if r.Method == "GET" {
		t, _ := template.ParseFiles("2_login.html")
		w.Header().Set("Content-Type", "text/html")
		_ = t.Execute(w, sess.Get("username"))
	} else {
		err := sess.Set("username", r.Form["username"])
		if err != nil {
			return
		}
		http.Redirect(w, r, "/", 302)
	}
}

func count(w http.ResponseWriter, r *http.Request) {
	sess := globalSessions.SessionStart(w, r)
	createTime := sess.Get("createTime")
	if createTime == nil {
		err := sess.Set("createTime", time.Now().Unix())
		if err != nil {
			return
		}
	} else if createTime.(int64)+360 < time.Now().Unix() {
		globalSessions.SessionDestroy(w, r)
		sess = globalSessions.SessionStart(w, r)
	}

	ct := sess.Get("countNum")
	if ct == nil {
		err := sess.Set("countNum", 1)
		if err != nil {
			return
		}
	} else {
		err := sess.Set("countNum", ct.(int)+1)
		if err != nil {
			return
		}
	}

	t, _ := template.ParseFiles("2_count.html")
	w.Header().Set("Content-Type", "text/html")
	err := t.Execute(w, sess.Get("countNum"))
	if err != nil {
		return
	}
}

func main() {
	// 注册路由
	http.HandleFunc("/login", login)
	http.HandleFunc("/count", count)

	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
