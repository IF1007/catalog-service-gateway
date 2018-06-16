package dns

import (
	"strings"

	"../constants"
)

type SystemRoute struct {
	PathPrefix string
	IsPublic   bool
	ServiceURL string
}

var routes = []SystemRoute{
	SystemRoute{
		PathPrefix: constants.PathCrud,
		IsPublic:   false,
		ServiceURL: constants.CrudHost,
	},
}

func IsRoutePublic(path string) bool {
	contextPath := strings.Split(path, constants.PathAPI)[1]
	for _, route := range routes {
		if strings.Index(contextPath, route.PathPrefix) == 0 {
			return route.IsPublic
		}
	}

	return false
}

func GetServiceURL(path string) string {
	contextPath := strings.Split(path, constants.PathAPI)[1]
	for _, route := range routes {
		if strings.Index(contextPath, route.PathPrefix) == 0 {
			return route.ServiceURL + contextPath
		}
	}

	return ""
}
