package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/luraproject/lura/v2/logging"
)

// HandlerRegisterer is the symbol the plugin loader will try to load. It must implement the Registerer interface
var HandlerRegisterer = registerer("header-logger")
var logger logging.Logger = logging.NoOp

type registerer string

func (r registerer) RegisterLogger(v interface{}) {
	l, ok := v.(logging.Logger)
	if !ok {
		return
	}
	logger = l
	logger.Debug(HandlerRegisterer, "server plugin loaded!!!")
}

func (r registerer) RegisterHandlers(f func(
	name string,
	handler func(context.Context, map[string]interface{}, http.Handler) (http.Handler, error),
)) {
	f(string(r), r.registerHandlers)
}

func (r registerer) registerHandlers(ctx context.Context, cfg map[string]interface{}, next http.Handler) (http.Handler, error) {
	// check the passed configuration and initialize the plugin
	pluginNames, ok := cfg["name"].([]interface{})
	if !ok {
		return nil, errors.New("wrong config")
	}

	if !r.isEnabled(pluginNames) {
		return nil, fmt.Errorf("unknown register %s", r)
	}

	if logger == nil {
		return next, nil
	}

	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		logger.Info(fmt.Sprintf("%#v", req.Header))
		next.ServeHTTP(w, req)
	}), nil
}

func (r registerer) isEnabled(pluginNames []interface{}) bool {
	for _, name := range pluginNames {
		if name == string(r) {
			return true
		}
	}
	return false
}

func init() {
	fmt.Printf("%s handler plugin loaded!!!\n", HandlerRegisterer)
}

func main() {

}
