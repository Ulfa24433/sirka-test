package endpoint

import (
	"fmt"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	modelV1 "github.com/Ulfa24433/sirka-test/backend/db/v1/model"
	utilV1 "github.com/Ulfa24433/sirka-test/backend/util/v1"
)

type userRequest struct {
	UserId string `json:"userid"`
}

func DisplayUser(c *gin.Context) {
	payload := &userRequest{}
	err := c.BindJSON(&payload)
	if err != nil {
		log.Error(err)
		utilV1.CallServerError(c, "invalid request", err)
		return
	}

	if payload.UserId == "" {
		err = fmt.Errorf("invalid request")
		log.Error(err)
		utilV1.CallServerError(c, "invalid request", err)
		return
	}

	//init connection to postgres
	db, err := utilV1.GetPostgreClient()
	if err != nil {
		log.Error(err)
		utilV1.CallServerError(c, "an error occure", err)
		return
	}

	usr := []*modelV1.Users{}
	err = db.
		Model(&usr).
		Where("userid = ?", payload.UserId).
		Find(&usr).
		Error
	if err != nil {
		log.Error(err)
		utilV1.CallServerError(c, "An error occure", err)
		return
	}
	utilV1.CallSuccess0K(c, "success", usr)

}

func DisplayAllUser(c *gin.Context) {

	//init connection to postgres
	db, err := utilV1.GetPostgreClient()
	if err != nil {
		log.Error(err)
		utilV1.CallServerError(c, "an error occure", err)
		return
	}

	usr := []*modelV1.Users{}
	err = db.
		Model(&usr).
		Find(&usr).
		Error
	if err != nil {
		log.Error(err)
		utilV1.CallServerError(c, "An error occure", err)
		return
	}
	utilV1.CallSuccess0K(c, "success", usr)
}
