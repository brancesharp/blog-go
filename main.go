package main

import (
	"fmt"
	"net/http"

	"blog-go/controllers"
	"blog-go/templates"
	"blog-go/views"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", controllers.StaticHandler(views.Must(views.ParseFS(templates.FS, "home.html", "tailwind.html"))))
	r.Get("/contact", controllers.StaticHandler(views.Must(views.ParseFS(templates.FS, "contact.html", "tailwind.html"))))
	r.Get("/blog-posts", controllers.StaticHandler(views.Must(views.ParseFS(templates.FS, "blog-posts.html", "tailwind.html"))))

	fmt.Println("Statring the server on :3000...")
	http.ListenAndServe(":3000", r)
}
