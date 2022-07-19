package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/luraproject/lura/v2/logging"
)

// ClientRegisterer is the symbol the plugin loader will try to load. It must implement the RegisterClient interface
var ClientRegisterer = registerer("s3")

type registerer string

var logger logging.Logger = logging.NoOp

func (r registerer) RegisterLogger(v interface{}) {
	l, ok := v.(logging.Logger)
	if !ok {
		return
	}
	logger = l
	logger.Debug(ClientRegisterer, "client plugin loaded!!!")
}

func (r registerer) RegisterClients(f func(
	name string,
	handler func(context.Context, map[string]interface{}) (http.Handler, error),
)) {
	f(string(r), r.registerClients)
}

func (r registerer) registerClients(ctx context.Context, extra map[string]interface{}) (http.Handler, error) {
	// check the passed configuration and initialize the plugin
	cfg, err := NewConfig(extra)
	if err != nil {
		return nil, errors.New("unable to load config")
	}
	if cfg.Name != string(r) {
		return nil, fmt.Errorf("unknown register %s", cfg.Name)
	}

	if cfg.Bucket == "" {
		return nil, errors.New("unable to get bucket from config")
	}
	logger.Debug(fmt.Sprintf("using s3 bucket %s", cfg.Bucket))

	s, err := NewS3Client(logger)
	if err != nil {
		return nil, errors.New("unable to initialize s3 client")
	}

	// return the actual handler wrapping or your custom logic so it can be used as a replacement for the default http client
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		ctx2 := req.Context()
		k := strings.TrimPrefix(req.URL.Path, "/")
		logger.Info(fmt.Sprintf("Fetching file from s3 bucket %s with key %s", cfg.Bucket, k))

		b, err := s.GetBytes(ctx2, cfg.Bucket, k)
		if err != nil {
			logger.Info(fmt.Sprintf("error getting content from s3, returning 404 %#v", err))
			w.WriteHeader(404)
			fmt.Fprintf(w, "Not found")
			return
		}

		_, err = w.Write(b)
		if err != nil {
			logger.Info("error writing content to body, returning 500")
			w.WriteHeader(500)
		}
	}), nil
}

func init() {
	fmt.Println(ClientRegisterer, "client plugin loaded!!!")
}

func main() {}
