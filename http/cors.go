package http

import (
	"net/http"
)

// Cross-Origin Resource Sharing: permit any :-)
// useful for debug.
// this should not be used in product ;D
func PermitAnyCORS(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
}
