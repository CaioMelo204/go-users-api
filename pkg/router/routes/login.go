package routes

import (
	"devbook-api/pkg/controllers"
	"net/http"
)

var routeLogin = Route{
	Uri:            "/login",
	Method:         http.MethodPost,
	Function:       controllers.Login,
	Authentication: false,
}
