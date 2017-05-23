//go:generate goagen bootstrap -d github.com/cratermoon/goaelog/design

package main

import (
	"github.com/cratermoon/goaelog/app"
	"github.com/cratermoon/goaelog/elog"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
)

func main() {
	// Create service
	logger := elog.NewNulLogger()
	service := goa.New("demo")
	service.WithLogger(elog.New(logger))

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	// Mount "datetime" controller
	c := NewDatetimeController(service)
	app.MountDatetimeController(service, c)
	// Mount "uuid" controller
	c2 := NewUUIDController(service)
	app.MountUUIDController(service, c2)

	// Start service
	if err := service.ListenAndServe(":8080"); err != nil {
		service.LogError("startup", "err", err)
	}

}
