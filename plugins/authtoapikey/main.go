package main

import (
	"errors"
	"fmt"

	"github.com/luraproject/lura/v2/logging"
	"github.com/luraproject/lura/v2/proxy"
)

func main() {}

func init() {
	fmt.Println(string(ModifierRegisterer), "loaded!!!")
}

// ModifierRegisterer is the symbol the plugin loader will be looking for. It must
// implement the plugin.Registerer interface
// https://github.com/luraproject/lura/blob/master/proxy/plugin/modifier.go#L71
var ModifierRegisterer = registerer("auth-to-apikey")

type registerer string

var logger logging.Logger = logging.NoOp

func (r registerer) RegisterLogger(i interface{}) {
	l, ok := i.(logging.Logger)
	if !ok {
		return
	}
	logger = l
	logger.Debug(ModifierRegisterer, "server plugin loaded!!!")
}

// RegisterModifiers is the function the plugin loader will call to register the
// modifier(s) contained in the plugin using the function passed as argument.
// f will register the factoryFunc under the name and mark it as a request
// and/or response modifier.
func (r registerer) RegisterModifiers(f func(
	name string,
	factoryFunc func(map[string]interface{}) func(interface{}) (interface{}, error),
	appliesToRequest bool,
	appliesToResponse bool,
)) {
	f(string(r)+"-request", r.requestModifier, true, false)
	fmt.Println(string(r), "registered!!!")
}

func (r registerer) requestModifier(m map[string]interface{}) func(interface{}) (interface{}, error) {
	return func(i interface{}) (interface{}, error) {
		req, ok := i.(proxy.RequestWrapper)
		if !ok {
			return nil, errors.New("unknown request type")
		}

		ha := req.Headers()["Authorization"]
		if len(ha) == 0 {
			logger.Info("request headers empty, skipping request modifier")
			return req, nil
		}

		req.Headers()["X-Api-Key"] = ha
		return req, nil
	}
}
