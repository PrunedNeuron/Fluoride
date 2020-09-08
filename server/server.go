package server

import (
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Server type
type Server struct {
	logger *zap.SugaredLogger
	router chi.Router
	server *http.Server
}

// New will setup the API listener
func New() (*Server, error) {

	router := chi.NewRouter()
	router.Use(Helmet())
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Recoverer)
	router.Use(Compression())
	router.Use(render.SetContentType(render.ContentTypeJSON))

	// Log Requests - Use appropriate format depending on the encoding
	if viper.GetBool("server.log_requests") {
		switch viper.GetString("logger.encoding") {
		case "stackdriver":
			router.Use(StackDriverHTTPLogger(viper.GetBool("server.log_requests_body"), viper.GetStringSlice("server.log_disabled_http")))
		default:
			router.Use(DefaultHTTPLogger(viper.GetBool("server.log_requests_body"), viper.GetStringSlice("server.log_disabled_http")))
		}
	}
	// Log Requests
	if viper.GetBool("server.log_requests") {
		router.Use(func(next http.Handler) http.Handler {
			return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				start := time.Now()
				var requestID string
				if reqID := r.Context().Value(middleware.RequestIDKey); reqID != nil {
					requestID = reqID.(string)
				}
				ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)
				next.ServeHTTP(ww, r)

				latency := time.Since(start)

				fields := []zapcore.Field{
					zap.Int("status", ww.Status()),
					zap.Duration("took", latency),
					zap.String("remote", r.RemoteAddr),
					zap.String("request", r.RequestURI),
					zap.String("method", r.Method),
					zap.String("package", "server.request"),
				}
				if requestID != "" {
					fields = append(fields, zap.String("request-id", requestID))
				}
				zap.L().Info("API Request", fields...)
			})
		})
	}

	s := &Server{
		logger: zap.S().With("package", "server"),
		router: router,
	}

	// Register routes
	s.router.Get("/version", GetVersion())

	return s, nil

}

// ListenAndServe will listen for requests
func (s *Server) ListenAndServe() error {

	s.server = &http.Server{
		Addr:    net.JoinHostPort(viper.GetString("server.host"), viper.GetString("server.port")),
		Handler: s.router,
	}

	// Listen
	listener, err := net.Listen("tcp", s.server.Addr)
	if err != nil {
		return fmt.Errorf("Could not listen on %s: %v", s.server.Addr, err)
	}

	go func() {
		if err = s.server.Serve(listener); err != nil {
			s.logger.Fatalw("API Listen error", "error", err, "address", s.server.Addr)
		}
	}()
	s.logger.Infow("API Listening", "address", s.server.Addr, "tls", viper.GetBool("server.tls"))

	// Enable profiler
	if viper.GetBool("server.profiler_enabled") && viper.GetString("server.profiler_path") != "" {
		zap.S().Debugw("Profiler enabled on API", "path", viper.GetString("server.profiler_path"))
		s.router.Mount(viper.GetString("server.profiler_path"), middleware.Profiler())
	}

	return nil
}

// Router returns the router
func (s *Server) Router() chi.Router {

	return s.router
}

// RenderOrErrInternal will render whatever you pass it (assuming it has Renderer) or prints an internal error
func RenderOrErrInternal(w http.ResponseWriter, r *http.Request, d render.Renderer) {
	if err := render.Render(w, r, d); err != nil {
		render.Render(w, r, ErrorInternal(err))
		return
	}
}
