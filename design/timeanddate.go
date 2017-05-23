package design

import (
        . "github.com/goadesign/goa/design"
        . "github.com/goadesign/goa/design/apidsl"
)

var _ = API("demo", func() {
        Title("Demo goa APIs")
        Description("simple utils")
        Host("localhost:8080")
        Scheme("http")
})

var _ = Resource("datetime", func() {
        Action("time", func() {
                Routing(GET("time/:tz/:long"))
                Description("time returns the time for the timezone (optionally in long format)")
                Params(func() {
                        Param("tz", String, "Time Zone")
                        Param("long", Boolean, "long format")
                })
                Response(OK, "text/plain")
        })

})
