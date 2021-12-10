package framework

import "net/http"

type Core struct {
	router map[string]ControllerHandler
}

func NewCore() *Core {
	return &Core{
		router: map[string]ControllerHandler{},
	}
}

func (c *Core) Get(url string, handler ControllerHandler) {
	c.router[url] = handler
}

func (c *Core) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}
