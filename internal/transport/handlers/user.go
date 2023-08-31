package handlers

import (
	"avitotest/internal/models"
	bodyparcer "avitotest/pkg/bodyParcer"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

func (h *handler) AddUserIntoSegment(w http.ResponseWriter, r *http.Request) {
	userSegments := &models.UserSegments{}
	bodyparcer.ParseBody(r, userSegments)
	err := h.service.User.AddUserIntoSegment(userSegments)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("error in add user into segment"))
		return
	}
	f := fmt.Sprintf("User with id-%d, was added into this segments:"+strings.Join(userSegments.SegmentsForAdd, " ")+" and was removed from this segments:"+strings.Join(userSegments.SegmentsForDelete, " "), userSegments.UserId)
	res, _ := json.Marshal(f)
	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}

func (h *handler) GetUserSegments(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["id"]
	userSegments, err := h.service.User.GetUserSegments(userId)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("error in Get User Segment"))
		return
	}
	type NewStruct struct {
		ActiveSegments string `json:"segment_name"`
	}
	var newStruct []NewStruct
	for _, us := range userSegments {
		newStruct = append(newStruct, NewStruct{
			ActiveSegments: us.SegmentName,
		})
	}
	res, _ := json.Marshal(newStruct)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
