package main

import (
	"errors"
	"fmt"
	"io"

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
var ModifierRegisterer = registerer("example-response-modifier")

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
func (r registerer) RegisterModifiers(
	f func(
		name string,
		factoryFunc func(map[string]interface{}) func(interface{}) (interface{}, error),
		appliesToRequest bool,
		appliesToResponse bool,
	),
) {
	f(string(r), r.requestModifier, false, true)
	fmt.Println(string(r), "registered!!!")
}

func (r registerer) requestModifier(m map[string]interface{}) func(interface{}) (interface{}, error) {
	return func(input interface{}) (interface{}, error) {
		resp, ok := input.(proxy.ResponseWrapper)
		if !ok {
			return nil, errors.New("unknown response")
		}

		logger.Info("Request modifier called.")
		logger.Info("Modifying ", resp.Data())

		modifiedResp := &responseWrapper{
			data: map[string]interface{}{
				"s3content": map[string]interface{}{
					"test": "test",
				},
			},
			isComplete: resp.IsComplete(),
			headers:    resp.Headers(),
			statusCode: resp.StatusCode(),
			io:         resp.Io(),
		}

		return modifiedResp, nil
	}
}

type responseWrapper struct {
	data       map[string]interface{}
	io         io.Reader
	isComplete bool
	headers    map[string][]string
	statusCode int
}

func (r responseWrapper) Data() map[string]interface{} {
	return r.data
}

func (r responseWrapper) Io() io.Reader {
	return r.io
}

func (r responseWrapper) IsComplete() bool {
	return r.isComplete
}

func (r responseWrapper) Headers() map[string][]string {
	return r.headers
}

func (r responseWrapper) StatusCode() int {
	return r.statusCode
}
