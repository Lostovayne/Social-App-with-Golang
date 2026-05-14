package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/Elevate-Techworks/social/internal/store"
	"github.com/fatih/color"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type application struct {
	config config
	store  store.Storage
}

type config struct {
	addr string
	db   dbConfig
	env  string
}

type dbConfig struct {
	addr         string
	maxOpenConns int
	maxIdleConns int
	maxIdleTime  string
}

// 1. Creamos nuestro middleware personalizado
func coloredLoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// WrapResponseWriter nos permite leer el status code y los bytes escritos
		ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)
		// Pasamos al siguiente middleware o a la ruta final
		next.ServeHTTP(ww, r)
		// Calculamos la latencia
		latency := time.Since(start)
		color.NoColor = false

		// Definimos los colores
		cyan := color.New(color.FgCyan).SprintFunc()
		green := color.New(color.FgGreen).SprintFunc()
		yellow := color.New(color.FgYellow).SprintFunc()
		red := color.New(color.FgRed).SprintFunc()

		// Coloreamos el status code dependiendo del resultado
		status := ww.Status()
		var coloredStatus string
		switch {
		case status >= 200 && status < 300:
			coloredStatus = green(status)
		case status >= 400 && status < 500:
			coloredStatus = yellow(status)
		case status >= 500:
			coloredStatus = red(status)
		default:
			coloredStatus = cyan(status)
		}

		// Colores adicionales para un look más limpio
		methodColor := color.New(color.FgCyan, color.Bold).SprintFunc()
		dim := color.New(color.FgHiBlack).SprintFunc()
		magenta := color.New(color.FgMagenta).SprintFunc()
		white := color.New(color.FgHiWhite, color.Bold).SprintFunc()

		timeStr := time.Now().Format("15:04:05")
		// Rellenamos el método a 6 espacios ("%-6s")
		// Así "GET   " y "DELETE" ocupan el mismo ancho.
		paddedMethod := fmt.Sprintf("%-6s", r.Method)
		paddedLatency := fmt.Sprintf("%-10s", latency.String())

		ip, _, err := net.SplitHostPort(r.RemoteAddr)
		if err != nil {
			ip = r.RemoteAddr
		}

		if ip == "::1" || ip == "127.0.0.1" {
			ip = "localhost"
		}

		fmt.Printf("%s %s %s %s %s %s %s %s %s %s\n",
			dim(timeStr),
			dim("│"),
			coloredStatus,
			dim("│"),
			methodColor(paddedMethod),
			dim("│"),
			magenta(paddedLatency),
			dim("│"),
			white(r.URL.Path),
			dim(fmt.Sprintf("(%s)", ip)),
		)
	})
}

func (app *application) mount() *chi.Mux {
	r := chi.NewRouter()

	// Middlewares
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)

	// Middleware con colores para los logs
	r.Use(coloredLoggerMiddleware)

	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	r.Route("/v1", func(r chi.Router) {
		r.Get("/health", app.healthCheckHandler)

		r.Route("/posts", func(r chi.Router) {
			r.Post("/", app.createPostHandler)
			r.Route("/{postID}", func(r chi.Router) {
				r.Use(app.postsContextMiddleware)
				r.Get("/", app.getPostHandler)
				r.Delete("/", app.deletePostHandler)
				r.Patch("/", app.updatePostHandler)
			})
		})

		r.Route("/users", func(r chi.Router) {
			r.Route("/{userID}", func(r chi.Router) {
				r.Use(app.userContextMiddleware)

				r.Get("/", app.getUserHandler)
				r.Put("follow", app.followUserHandler)
				r.Put("unfollow", app.unfollowUserHandler)
			})

			r.Group(func(r chi.Router) {
				r.Get("/feed", app.getUserFeedHandler)
			})
		})

	})

	return r
}

func (app *application) run(mux *chi.Mux) error {

	srv := &http.Server{
		Addr:         app.config.addr,
		Handler:      mux,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Minute,
	}

	log.Printf("Server has started %s", app.config.addr)
	return srv.ListenAndServe()
}
