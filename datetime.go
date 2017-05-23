package main

import (
	"time"

	"github.com/cratermoon/goaelog/app"
	"github.com/goadesign/goa"
)

// DatetimeController implements the datetime resource.
type DatetimeController struct {
	*goa.Controller
}

// NewDatetimeController creates a datetime controller.
func NewDatetimeController(service *goa.Service) *DatetimeController {
	return &DatetimeController{Controller: service.NewController("DatetimeController")}
}

// Time runs the time action.
func (c *DatetimeController) Time(ctx *app.TimeDatetimeContext) error {
	// DatetimeController_Time: start_implement

	layout := time.Kitchen
	if ctx.Long {
		layout = time.RFC1123Z
	}
	loc, err := time.LoadLocation(ctx.Tz)
	if err != nil {
		loc, _ = time.LoadLocation("UTC")
	}
	return ctx.OK([]byte(time.Now().In(loc).Format(layout)))

	// DatetimeController_Time: end_implement
	return nil
}
