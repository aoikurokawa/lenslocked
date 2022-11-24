package main

import (
	"fmt"
	"net/http"

	"github.com/Aoi1011/lenslocked/controllers"
	"github.com/Aoi1011/lenslocked/templates"
	"github.com/Aoi1011/lenslocked/views"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Get("/", controllers.StaticHandler(views.Must(
		views.ParseFS(templates.FS, "home.gohtml", "tailwind.gohtml"))))
	r.Get("/contact", controllers.StaticHandler(views.Must(
		views.ParseFS(templates.FS, "contact.gohtml", "tailwind.gohtml"))))
	r.Get("/faq", controllers.FAQ(views.Must(
		views.ParseFS(templates.FS, "faq.gohtml", "tailwind.gohtml"))))
	var userC controllers.Users
	userC.Templates.New = views.Must(
		views.ParseFS(templates.FS, "signup.gohtml", "tailwind.gohtml"))
	r.Get("/signup", userC.New)
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page Not Found", http.StatusNotFound)
	})

	fmt.Println("Starting the server on :3000...")
	http.ListenAndServe(":3000", r)
}
