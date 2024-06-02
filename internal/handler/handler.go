package handler

import (
    "encoding/json"
    "net/http"
    "send-email/internal/service"
)

type Handler struct {
    srv *service.Service
}

func NewHandler(srv *service.Service) *Handler {
    return &Handler{srv: srv}
}

func (h *Handler) CreateCampaign(w http.ResponseWriter, r *http.Request) {
    var input struct {
        Name    string   `json:"name"`
        Content string   `json:"content"`
        Emails  []string `json:"emails"`
    }
    if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }

    id, err := h.srv.CreateCampaign(input.Name, input.Content, input.Emails)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(map[string]string{"id": id})
}
