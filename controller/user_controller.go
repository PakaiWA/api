package controller

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type UserController interface {
	RegisterRoutes(router *httprouter.Router)
	Login(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	CreateUser(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}
