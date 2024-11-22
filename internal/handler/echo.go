package handler

import (
	"fmt"
	"net/http"

	u "github.com/ogonna-anaekwe/matrixserver/internal/utils"
)

// Handler for requests to /echo.
func (h *Handler) HandlerEcho(w http.ResponseWriter, r *http.Request) {
	res, err := u.Reduce(h.Rows, u.Echo)
	h.handleError(err, w)

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, res)
}
