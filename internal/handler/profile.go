package handler

import (
	"encoding/json"
	"go-backend-template/internal/model"
	"go-backend-template/internal/service"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type ProfileHandler struct {
	service *service.ProfileService
}

func NewProfileHandler(service *service.ProfileService) *ProfileHandler {
	return &ProfileHandler{service: service}
}

func (h *ProfileHandler) CreateProfile(w http.ResponseWriter, r *http.Request) {
	var req model.ProfileCreate
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	profile, err := h.service.CreateProfile(r.Context(), &req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(profile)
}

func (h *ProfileHandler) GetProfile(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := uuid.Parse(vars["id"])
	if err != nil {
		http.Error(w, "Invalid profile ID", http.StatusBadRequest)
		return
	}

	profile, err := h.service.GetProfile(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(profile)
}

func (h *ProfileHandler) GetAllProfiles(w http.ResponseWriter, r *http.Request) {
	profiles, err := h.service.GetAllProfiles(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(profiles)
}

func (h *ProfileHandler) UpdateProfile(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := uuid.Parse(vars["id"])
	if err != nil {
		http.Error(w, "Invalid profile ID", http.StatusBadRequest)
		return
	}

	var req model.ProfileUpdate
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	profile, err := h.service.UpdateProfile(r.Context(), id, &req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(profile)
}

func (h *ProfileHandler) DeleteProfile(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := uuid.Parse(vars["id"])
	if err != nil {
		http.Error(w, "Invalid profile ID", http.StatusBadRequest)
		return
	}

	if err := h.service.DeleteProfile(r.Context(), id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *ProfileHandler) GetProfileByUserID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID, err := uuid.Parse(vars["id"])
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	profile, err := h.service.GetProfileByUserID(r.Context(), userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(profile)
}

func (h *ProfileHandler) SearchProfile(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	if username := query.Get("username"); username != "" {
		// Реализация поиска по username
		// Вам нужно добавить соответствующий метод в сервис и репозиторий
		http.Error(w, "Not implemented", http.StatusNotImplemented)
		return
	}

	if email := query.Get("email"); email != "" {
		// Реализация поиска по email
		// Вам нужно добавить соответствующий метод в сервис и репозиторий
		http.Error(w, "Not implemented", http.StatusNotImplemented)
		return
	}

	http.Error(w, "Please provide either username or email parameter", http.StatusBadRequest)
}
