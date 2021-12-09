package framework

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"sync"
	"time"
)

type GetParams interface {
	GetInt (key string, def int) int
	GetString (key, def string) string
	GetArray (key string, def []string) []string
}

type Param struct {
	p map[string][]string
}

type Context struct {
	request *http.Request
	responseWriter http.ResponseWriter
	ctx context.Context
	//是否超时标记位
	hasTimeout bool
	//写锁
	writerMux *sync.RWMutex
	Keys map[string]interface{}

}

func NewContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{
		request:        nil,
		responseWriter: nil,
		ctx:            nil,
		hasTimeout:     false,
		writerMux:      nil,
	}
}

func (c *Context) BaseContext() context.Context {
	return c.request.Context()
}

func (c *Context) Done() <-chan struct{} {
	return c.BaseContext().Done()
}

func (c *Context) WriterMux() *sync.RWMutex {
	return c.writerMux
}

func (c *Context) HasTimeout() bool {
	return c.hasTimeout
}

func (c *Context) GetRequest() *http.Request {
	return c.request
}

func (c *Context) GetResponse() http.ResponseWriter {
	return c.responseWriter
}

func (c *Context) SetHasTimeout() {
	c.hasTimeout = true
}

// #实现context标准库接口

func (c *Context)Deadline()(deadline time.Time, ok bool) {
	return c.ctx.Deadline()
}

func (c *Context) Err() error {
	return c.ctx.Err()
}

func (c *Context) Value(key interface{}) interface{} {
	return c.ctx.Value(key)
}


func (c *Context) Set(k string, v interface{}) {
	c.writerMux.Lock()
	defer c.writerMux.Unlock()
	if c.Keys == nil {
		c.Keys = make(map[string]interface{})
	}
	c.Keys[k] = v
}

func (c *Context) Get(k string) (v interface{}, ok bool) {
	c.writerMux.RLock()
	v, ok = c.Keys[k]
	c.writerMux.RUnlock()
	return
}

//接收uri参数的方法

func (c *Context) QueryAll() *Param {
	if c.request == nil {
		return &Param{p: map[string][]string{}}
	}
	return &Param{p: c.request.URL.Query()}
}


func (c *Context) FormAll() *Param {
	if c.request == nil {
		return &Param{p: map[string][]string{}}
	}
	return &Param{p: c.request.PostForm}
}


func (p *Param) GetInt( key string, def int) int {
	if v, ok := p.p[key]; ok {
		l := len(v)
		if l > 0 {
			i, err := strconv.Atoi(v[l-1])
			if err != nil {
				return def
			}
			return i
		}
	}
	return def
}

func (p *Param) GetString (key, def string) string {
	if v, ok := p.p[key]; ok {
		l := len(v)
		if l > 0 {
			return v[l-1]
		}
	}
	return def
}

func (p *Param) GetArray( key string, def []string) []string {
	if v, ok := p.p[key]; ok {
		return v
	}
	return def
}

//获取json请求数据

func (c *Context) BindJson(obj interface{}) error {
	if c.request != nil {
		body, err := ioutil.ReadAll(c.request.Body)
		if err != nil {
			return err
		}
		c.request.Body = ioutil.NopCloser(bytes.NewBuffer(body))
		err = json.Unmarshal(body, obj)
		if err != nil {
			return err
		}
	}else {
		return errors.New("ctx.request empty")
	}
	return nil
}

