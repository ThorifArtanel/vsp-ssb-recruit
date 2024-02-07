package apis

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SignUpForm struct {
	Name    string `form:"name"`
	Email   string `form:"email"`
	Subject string `form:"subject"`
	Message string `form:"message"`
}

func FormHandler(c *gin.Context) {
	log.Print("INFO: Parsing Payload")
	var req SignUpForm
	if err := c.Bind(&req); err != nil {
		log.Printf("ERROR: %+v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"kode":   http.StatusBadRequest,
			"status": "Bad Request",
			"pesan":  "The request could not be understood by the server due to incorrect syntax."})
		return
	}

	fmt.Println(req)
}
