//go:generate goagen bootstrap -d github.com/tikasan/goa-datastore/design

package server

import (
	"net/http"

	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"github.com/tikasan/goa-datastore/app"
	"github.com/tikasan/goa-datastore/controller"
)

func init() {
	// Create service
	service := goa.New("appengine")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	// Mount "Account" controller
	c := controller.NewAccountController(service)
	app.MountAccountController(service, c)
	// Mount "swagger" controller
	c2 := controller.NewSwaggerController(service)
	app.MountSwaggerController(service, c2)

	// Start service
	http.HandleFunc("/", service.Mux.ServeHTTP)
}
