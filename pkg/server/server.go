package server

import (
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/PrunedNeuron/Fluoride/internal/routes"
	"github.com/PrunedNeuron/Fluoride/pkg/logger"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Server type
type Server struct {
	logger     *zap.SugaredLogger
	router     chi.Router
	httpServer *http.Server
}

// New will set up a new Server
func New() (*Server, error) {
	router := chi.NewRouter()
	router.Use(Helmet())
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Recoverer)
	router.Use(Compression())
	router.Use(render.SetContentType(render.ContentTypeJSON))

	router.Use(logger.DefaultHTTPLogger(
		viper.GetBool("server.log_requests_body"),
		viper.GetStringSlice("server.log_disabled_http"),
	))

	// Do the logging
	if viper.GetBool("server.log_requests") {
		router.Use(func(next http.Handler) http.Handler {
			return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				start := time.Now()

				wrapResponseWriter := middleware.NewWrapResponseWriter(w, r.ProtoMajor)
				next.ServeHTTP(wrapResponseWriter, r)

				latency := time.Since(start)

				fields := []zapcore.Field{
					zap.Int("status", wrapResponseWriter.Status()),
					zap.Duration("took", latency),
					zap.String("remote", r.RemoteAddr),
					zap.String("request", r.RequestURI),
					zap.String("method", r.Method),
					zap.String("package", "server"),
				}

				zap.L().Info("API Request", fields...)
			})
		})
	}

	server := &Server{
		logger: zap.S().With("package", "server"),
		router: router,
	}

	// Routes
	routes.Route(server.router) // Don't need to pass address since chi.Router is already a pointer, since it is of type *chi.Mux

	return server, nil
}

// Serve will start the server and make it listen for requests
func (server *Server) Serve() error {

	// Create the http server and assign host, port from config
	server.httpServer = &http.Server{
		Addr:    net.JoinHostPort(viper.GetString("server.host"), viper.GetString("server.port")),
		Handler: server.router,
	}

	// Listen for requests
	listener, err := net.Listen(viper.GetString("server.network"), server.httpServer.Addr)
	if err != nil {
		return fmt.Errorf("Could not listen on %s: %v", server.httpServer.Addr, err)
	}

	go func() {
		if err = server.httpServer.Serve(listener); err != nil {
			server.logger.Fatalw("Listen error", "error", err, "address", server.httpServer.Addr)
		}
	}()

	server.logger.Infow("Server listening", "address", "http://"+server.httpServer.Addr)

	// Profiler
	if viper.GetBool("server.profiler_enabled") && viper.GetString("server.profiler_path") != "" {
		zap.S().Debugw("Profiler enabled on server", "path", viper.GetString("server.profiler_path"))
		server.router.Mount(viper.GetString("server.profiler_path"), middleware.Profiler())
	}

	return nil
}
