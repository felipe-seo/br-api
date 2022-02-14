package controllers

import "net/http"

type ControllerMethods interface {
	GetAccounts(http.ResponseWriter, *http.Request)
	GetBalance(http.ResponseWriter, *http.Request)
	CreatAccount(http.ResponseWriter, *http.Request)
}

type Controller struct {
}

//func newController(d database)
