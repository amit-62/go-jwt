package helper

import (
	"errors"
	"github.com/gin-gonic/gin"
)

func CheckUserType(c *gin.Context, role string) (err error) {
	userType := c.GetString("user_type")
	err = nil

	if userType != role {
		err = errors.New("unauthorised to access this resources")
		return err
	} 

	return err

}

func MatchUserTypeToUid(c *gin.Context, userId string) (err error){
	userType := c.GetString("user_type")
	uid := c.GetString("user_id")
	err = nil

	if userType == "USER" && userId != uid {
		err = errors.New("unauthorised to access this resources")
		return err
	} 

	err = CheckUserType(c, userType)
	return err
}