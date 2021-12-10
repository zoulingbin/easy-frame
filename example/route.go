package main

import (
	"github.com/zoulingbin/easy-frame/framework"
)

func RegisterRoute(core *framework.Core) {
	core.Get("foo", FooControllerHandler)
}
