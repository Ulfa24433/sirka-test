package main

import (
	"fmt"
	"os"

	endpointV1 "github.com/Ulfa24433/sirka-test/backend/endpoint"
	"github.com/Ulfa24433/sirka-test/backend/util/v1/envvar"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func main() {
	//initialize gin

	r := gin.Default()
	v1 := r.Group("/v1")

	//endpoint to display user
	user := v1.Group("/user")
	user.POST("/displayUser", endpointV1.DisplayUser)

	//endpoint to display all user

	user.GET("/displayAllUser", endpointV1.DisplayAllUser)

	port := os.Getenv(envvar.Port)

	err := r.Run(fmt.Sprintf(":%s", port))
	if err != nil {
		log.Error(err)

	}

}
