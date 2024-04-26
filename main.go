package main

import (
	"fmt"
	"net/http"

	"lenslocked/controllers"
	"lenslocked/templates"
	"lenslocked/views"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", controllers.StaticHandler(views.Must(views.ParseFS(templates.FS, "home.html", "tailwind.html"))))
	r.Get("/contact", controllers.StaticHandler(views.Must(views.ParseFS(templates.FS, "contact.html", "tailwind.html"))))
	r.Get("/faq", controllers.FAQ(views.Must(views.ParseFS(templates.FS, "faq.html", "tailwind.html"))))
	r.Get("/signup", controllers.StaticHandler(views.Must(views.ParseFS(templates.FS, "signup.html", "tailwind.html"))))

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)
	})
	fmt.Println("Statring the server on :3000...")
	http.ListenAndServe(":3000", r)
}
