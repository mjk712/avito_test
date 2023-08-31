package handlers

import (
	"avitotest/internal/models"
	bodyparcer "avitotest/pkg/bodyParcer"
	"encoding/json"
	"fmt"
	"net/http"
)

func (h *handler) AddSegment(w http.ResponseWriter, r *http.Request) {
	segment := &models.Segment{}
	bodyparcer.ParseBody(r, segment)
	err := h.service.Segment.CreateSegment(segment.Name)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("error in add segment"))
		return
	}
	f := fmt.Sprintf("add new segment:%s", segment.Name)
	res, _ := json.Marshal(f)
	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}

func (h *handler) DeleteSegment(w http.ResponseWriter, r *http.Request) {
	segment := &models.Segment{}
	bodyparcer.ParseBody(r, segment)
	err := h.service.Segment.DeleteSegment(segment.Name)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("error in delete Segment"))
		return
	}
	f := fmt.Sprintf("Delete %s segment", segment.Name)
	res, _ := json.Marshal(f)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
