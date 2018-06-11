package discovery

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type SystemRoute struct {
	PathPrefix string
	Methods    []string
	IsPublic   bool
	ServiceURL string
}

var routes = []SystemRoute{
	SystemRoute{
		PathPrefix: "crud",
		Methods:    []string{"GET", "POST", "PUT", "DELETE"},
		IsPublic:   false,
		ServiceURL: "http://csd-crud/",
	},
}

// StartDiscoveryService send requests in broadcast
func StartDiscoveryService(router *mux.Router, f func(http.ResponseWriter, *http.Request)) {

	for _, route := range routes {
		router.Methods(route.Methods...).PathPrefix(route.PathPrefix).HandlerFunc(f)
	}
}

// GetServiceURL return "" if route does not exist
func GetServiceURL(req *http.Request) string {
	// TODO: Change method
	fmt.Println(req.URL.Path)
	return routes[0].ServiceURL + req.URL.Path
	// return "serviceip:port/" + req.URL.Path
}
