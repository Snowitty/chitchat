package routes

import (
	"net/http"

	"github.com/snowitty/chitchat/handlers"
)

type WebRoute struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type WebRoutes []WebRoute

var webRoutes = WebRoutes{
	{
		"home",
		"GET",
		"/",
		handlers.Index,
	},

	{
		"signup",
		"GET",
		"/signup",
		handlers.Signup,
	},

	{
		"signupAccount",
		"POST",
		"/signup_account",
		handlers.SignupAccount,
	},

	{
		"login",
		"GET",
		"/login",
		handlers.Login,
	},

	{
		"auth",
		"POST",
		"/authenticate",
		handlers.Authenticate,
	},

	{
		"logout",
		"GET",
		"/logout",
		handlers.Logout,
	},
}
