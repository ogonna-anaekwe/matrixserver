package handler

import (
	"net/http"

	config "github.com/ogonna-anaekwe/matrixserver/config"

	"github.com/sirupsen/logrus"
)

// Valid API paths
const (
	EchoPath     string = "/echo"
	SumPath      string = "/sum"
	InvertPath   string = "/invert"
	FlattenPath  string = "/flatten"
	MultiplyPath string = "/multiply"
)

// Request Handler.
type Handler struct {
	Cfg  config.Config  // service configs e.g. port, data file location
	Log  *logrus.Logger // logger
	Rows [][]string     // matrix data from CSV file
}

// Logs errors and writes status code when errors are encountered in request.
func (h *Handler) handleError(err error, w http.ResponseWriter) {
	if err != nil {
		h.Log.SetReportCaller(true)
		h.Log.Errorf("%v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}
