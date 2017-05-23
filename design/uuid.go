package design

import (
        . "github.com/goadesign/goa/design"
        . "github.com/goadesign/goa/design/apidsl"
)

var _ = Resource("uuid", func() {
        Action("generate", func() {
                Routing(GET("uuid"))
                Description("generate returns a new randomly generated uuid")
                Response(OK, "text/plain")
        })

})
