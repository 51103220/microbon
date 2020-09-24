package server

import (
	"context"
	"errors"
	"fmt"
	"github.com/51103220/microbon/core"
	"github.com/51103220/microbon/registry"
	"github.com/51103220/microbon/text"
	"log"
	"net/http"
	"net/http/pprof"
	"reflect"
	"time"
)

var defaultShutdownTimeout = 10

type GatewayConfig struct {
	path            string
	processRequest  func(data string) string
	processResponse func(response interface{}) string
}

type ServerConfig struct {
	ProfilingEnabled bool
}

type Server struct {
	port   int
	ctx    context.Context
	logger *log.Logger
	srv    *http.Server
	config *ServerConfig
	router *http.ServeMux
}

func NewServer(port int, ctx context.Context, logger *log.Logger, config *ServerConfig) (*Server, error) {
	if port == 0 {
		return nil, errors.New("port must not be empty")
	}

	address := fmt.Sprintf(":%d", port)
	router := http.NewServeMux()
	//TODO: add more configuration
	srv := &http.Server{
		Addr:     address,
		Handler:  nil,
		ErrorLog: logger,
	}

	if config == nil {
		config = &ServerConfig{}
	}

	if config.ProfilingEnabled {
		registerPProfileHandlers(router)
	}
	return &Server{port, ctx, logger, srv, config, router}, nil
}

func (server *Server) RegisterServices() error {
	for _, service := range registry.GetServices() {
		//register

		server.router.HandleFunc(service.Path, func(w http.ResponseWriter, r *http.Request) {

			gw := core.GBonResponseWriter{ResponseWriter: w}

			if service.Filter.Auth != nil {
				//apply authentication
				service.Filter.Auth.ServeHTTP(&gw, r)
				if gw.Done {
					return
				}
			}

			if len(service.Filter.RequestFilters) > 0 {
				//apply request filters
				for _, filter := range service.Filter.RequestFilters {
					filter.Handler(&gw, r)

					if gw.Done {
						return
					}
				}
			}

			request := text.DeserializeRequest(&gw, r, service.PayloadType)

			if request.IsNil() {
				gw.WriteHeader(400)
				gw.Write([]byte("Invalid payload"))

			}

			if gw.Done {
				//TODO: check if empty response or empty statusCode
				return
			}

			//TODO: check ten thousands of condition
			ctx := context.Background()
			//TODO: create timeout context or cancel context based on config
			requestContext, cancelRequest := context.WithCancel(ctx)

			defer func() {
				cancelRequest()
			}()

			method := request.MethodByName("Process")
			values := method.Call([]reflect.Value{reflect.ValueOf(requestContext)})

			errVal := values[0]
			responseVal := values[1]

			if !errVal.IsNil() {
				err := errVal.Elem().Interface().(error)
				gw.WriteHeader(500)
				gw.Write([]byte(err.Error()))
				//TODO: write error at function level
				return
			}

			var response interface{}
			if !responseVal.IsNil() {
				response = responseVal.Elem().Interface().(interface{})
			}

			if len(service.Filter.RequestFilters) > 0 {
				//apply response filters
				for _, filter := range service.Filter.ResponseFilters {
					filter.Handler(&gw, r)
				}
			}

			text.SerializeResponse(&gw, r, response)
		})
	}

	return nil
}

//TODO: do something to help application
func (server *Server) RegisterGateway(path string, handler http.HandlerFunc) error {
	server.router.HandleFunc(path, handler)
	return nil
}

func (server *Server) Init() {
	server.srv.Handler = server.router
}

func (server *Server) Start() error {

	return server.srv.ListenAndServe()

}

func (server *Server) Shutdown(timeout int) error {

	if timeout == 0 {
		timeout = defaultShutdownTimeout
	}

	ctx, cancel := context.WithTimeout(server.ctx, time.Duration(timeout)*time.Second) //and context is the mechanism to handle goRoutines timeout

	defer cancel()

	server.srv.SetKeepAlivesEnabled(false)

	if err := server.srv.Shutdown(ctx); err != nil {
		return errors.New(fmt.Sprintf("Could not gracefully shutdown the server: %v\n", err))
	}

	return nil
}

func registerPProfileHandlers(router *http.ServeMux) {
	router.HandleFunc("/debug/pprof", pprof.Index)
	router.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	router.HandleFunc("/debug/pprof/profile", pprof.Profile)
	router.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	router.HandleFunc("/debug/pprof/trace", pprof.Trace)
}
