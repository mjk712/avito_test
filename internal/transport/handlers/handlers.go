package handlers

import (
	"avitotest/internal/service"

	"github.com/gorilla/mux"
)

type handler struct {
	service *service.Services
}

type Handler interface {
	DinamicSegmentRoutes(router *mux.Router)
}

func NewHandler(service *service.Services) Handler {
	return &handler{service}
}

func (h *handler) DinamicSegmentRoutes(router *mux.Router) {

	router.HandleFunc("/dinamicsegment/add_segment", h.AddSegment).Methods("POST")
	router.HandleFunc("/dinamicsegment/delete_segment", h.DeleteSegment).Methods("DELETE")

	router.HandleFunc("/dinamicsegment/add_user_into_segment", h.AddUserIntoSegment).Methods("POST")
	router.HandleFunc("/dinamicsegment/get_user_segments/{id}", h.GetUserSegments).Methods("GET")
}
