package handlers

import (
	"net/http"

	db "github.com/ciftci-mehmet/go_cms/db/sqlc"
)

type Handlers struct {
	Query *db.Queries
}

func (h *Handlers) HelloWorld(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}
