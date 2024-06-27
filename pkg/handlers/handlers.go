package handlers

import (
	"github.com/mziyadabek/bookings/pkg/render"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplatest(w, "Home.html")
}
func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplatest(w, "About.html")
}
