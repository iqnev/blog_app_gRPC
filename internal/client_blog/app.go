package client_blog

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	gohandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/hashicorp/go-hclog"
	"github.com/iqnev/blog_app_gRPC/internal/client_blog/rest"
	protos "github.com/iqnev/blog_app_gRPC/proto/blog"
)

type App struct {
	Router *mux.Router
	Log    hclog.Logger
	Blog   protos.BlogServiceClient
	client protos.BlogService_ListBlogsClient
}

type RequestHandlerFunction func(w http.ResponseWriter, r *http.Request, b protos.BlogServiceClient)

func (a *App) Initialize(l hclog.Logger, b protos.BlogServiceClient) {
	a.Log = l
	a.Router = mux.NewRouter()
	a.Blog = b
	a.setRouters()

}

func (a *App) Run() {
	cor := gohandlers.CORS(gohandlers.AllowedOrigins([]string{"*"}))
	serv := &http.Server{
		Addr:         ":8787",
		Handler:      cor(a.Router),
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		a.Log.Info("Starting server on port %s", 2323)
		err := serv.ListenAndServe()

		if err != nil {
			a.Log.Error("Error starting server", "error", err)
			os.Exit(1)
		}
	}()

	sigChan := make(chan os.Signal)

	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan

	log.Println("Got signal:", sig)

	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)

	serv.Shutdown(ctx)
}

func (a *App) setRouters() {
	a.Get("/blogs", MiddlewareValidateBlog(a.handleRequest(rest.GetAllBlogs)))
}

func MiddlewareValidateBlog(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA")
		next.ServeHTTP(w, req)
	})
}

// Get wraps the router for GET method
func (a *App) Get(path string, f func(rw http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods(http.MethodGet)
}

// Post wraps the router for POST method
func (a *App) Post(path string, f func(rw http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods(http.MethodPost)
}

// Put wraps the router for PUT method
func (a *App) Put(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods(http.MethodPut)
}

// Delete wraps the router for DELETE method
func (a *App) Delete(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods(http.MethodDelete)
}

func (a *App) handleRequest(handler RequestHandlerFunction) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handler(w, r, a.Blog)
	}
}
