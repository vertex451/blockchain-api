package http_handlers

import (
	"net/http"
)

type HttpProvider interface {
	LivenessCheck(w http.ResponseWriter, r *http.Request)
}
