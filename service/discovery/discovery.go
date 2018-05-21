package discovery

import (
	"net/http"

	"github.com/gorilla/mux"
)

// send requests in broadcast
func StartDiscoveryService(router *mux.Router, f func(http.ResponseWriter, *http.Request)) {
	router.Methods("GET").PathPrefix("v1").HandlerFunc(f)
}

func IsRoutePublic(path string) bool {
	return false
}

// return "" if route does not exist
func GetServiceURL(req *http.Request) string {
	return ""
	// return "serviceip:port/" + req.URL.Path
}
