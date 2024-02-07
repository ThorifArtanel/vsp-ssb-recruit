package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	api "gitlab.com/vsp-ssb/landing-page-recruitment/api"
)

func main() {
	r := gin.Default()

	r.StaticFS("/assets/images", http.Dir("./web/static/images"))
	r.StaticFS("/assets/svg", http.Dir("./web/static/svg"))

	r.StaticFS("/assets/css", http.Dir("./web/static/css"))
	r.StaticFS("/assets/js", http.Dir("./web/static/js"))
	r.StaticFS("/assets/plugins", http.Dir("./web/static/plugins"))

	r.Use(static.Serve("/", static.LocalFile("./web/static/html", true)))

	r.POST("/api/contact-form", api.FormHandler)

	port := "8080"
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}
	fmt.Printf("setting port as %s", port)
	r.Run(":" + port)
}
