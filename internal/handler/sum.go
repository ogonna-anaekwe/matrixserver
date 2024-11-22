package handler

import (
	"fmt"
	"net/http"

	u "github.com/ogonna-anaekwe/matrixserver/internal/utils"
)

// Handler for requests to /sum.
func (h *Handler) HandlerSum(w http.ResponseWriter, r *http.Request) {
	res, err := u.Reduce(h.Rows, u.Sum)
	h.handleError(err, w)

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, res)
}
