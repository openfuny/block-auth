package adminhandler

import (
	"fmt"
	"net/http"
)

type ManagerHandler struct {
}

func (h *ManagerHandler) ServeHTTP(w http.ResponseWriter,
	r *http.Request) {
	fmt.Fprintf(w, "Hi, This is an example of http service in golang!\n")
}
