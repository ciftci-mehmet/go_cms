package handlers

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func (h *Handlers) GetUser(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		w.WriteHeader(404)
		w.Write([]byte("invalid param"))
		return
	}
	user, err := h.Query.GetUserById(context.Background(), id)
	if err != nil {
		w.WriteHeader(404)
		w.Write([]byte("Not found"))
		return
	}

	userinfo := fmt.Sprintf("ID:%v\nUsername:%s\nHashedPassword:%s\nCreatedAt:%s", user.ID, user.Username, user.HashedPassword, user.CreatedAt)
	w.Write([]byte(userinfo))
}
