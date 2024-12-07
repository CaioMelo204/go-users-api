package routes

import (
	"devbook-api/pkg/controllers"
	"net/http"
)

var userRoutes = []Route{
	{
		Uri:            "/user",
		Method:         http.MethodGet,
		Function:       controllers.GetUserList,
		Authentication: false,
	},
	{
		Uri:            "/user",
		Method:         http.MethodPost,
		Function:       controllers.CreateUser,
		Authentication: false,
	},
	{
		Uri:            "/user/{userId}",
		Method:         http.MethodGet,
		Function:       controllers.GetUser,
		Authentication: true,
	},
	{
		Uri:            "/user/{userId}",
		Method:         http.MethodPatch,
		Function:       controllers.UpdateUser,
		Authentication: true,
	},
	{
		Uri:            "/user/{userId}",
		Method:         http.MethodDelete,
		Function:       controllers.DeleteUser,
		Authentication: true,
	},
}
