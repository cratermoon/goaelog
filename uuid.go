package main

import (
	"fmt"
	"io"
	"crypto/rand"

	"github.com/cratermoon/goaelog/app"
	"github.com/goadesign/goa"
)

// UUIDController implements the uuid resource.
type UUIDController struct {
	*goa.Controller
}

// NewUUIDController creates a uuid controller.
func NewUUIDController(service *goa.Service) *UUIDController {
	return &UUIDController{Controller: service.NewController("UUIDController")}
}

// Generate runs the generate action.
func (c *UUIDController) Generate(ctx *app.GenerateUUIDContext) error {
	// UUIDController_Generate: start_implement

        uuid := make([]byte, 16)
        n, err := io.ReadFull(rand.Reader, uuid)
        if n != len(uuid) || err != nil {
                return err
        }
        // variant bits; see section 4.1.1
        uuid[8] = uuid[8]&^0xc0 | 0x80
        // version 4 (pseudo-random); see section 4.1.3
        uuid[6] = uuid[6]&^0xf0 | 0x40
	s := fmt.Sprintf("%x-%x-%x-%x-%x", uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:])
        return ctx.OK([]byte(s))

	// UUIDController_Generate: end_implement
	return nil
}
