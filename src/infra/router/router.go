package router

import (
	"net/http"

	"github.com/tora0091/get-account-data/infra/handler"
)

type Router struct {
	AccountHandler *handler.AccountHandler
}

func NewRouter(aH *handler.AccountHandler) *Router {
	return &Router{
		AccountHandler: aH,
	}
}

func (router *Router) SetRouting(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/list":
		router.AccountHandler.GetList(w, r)
	case "/item":
		router.AccountHandler.GetItem(w, r)
	default:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("bad request"))
	}
}
