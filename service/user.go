package service

import (
	"github.com/gin-gonic/gin"
	"learning2.0/database"
	"learning2.0/model"
	"learning2.0/util"
	"log"
)

func GetUserById(c *gin.Context, id int64) (*model.User, error) {
	userDb := database.Query.User
	user, err := userDb.WithContext(c).Where(userDb.ID.Eq(id)).First()
	if err != nil {
		log.Printf("userDb query failed, ids:%v, err:%v\n", id, err)
		return nil, util.BuildErrorInfo("userDb query failed, id:%v, err:%v", id, err)
	}
	return user, nil
}

func GetUserByUsernameAndPassword(c *gin.Context, username string, password string) (*model.User, error) {
	userDb := database.Query.User
	user, err := userDb.WithContext(c).Where(userDb.Username.Eq(username)).Where(userDb.Password.Eq(password)).First()
	if err != nil {
		log.Printf("userDb query failed, err:%v\n", err)
		return nil, util.BuildErrorInfo("userDb query failed, err:%v", err)
	}
	return user, nil
}

func AddUser(c *gin.Context, user *model.User) error {
	userDb := database.Query.User
	err := userDb.WithContext(c).Create(user)
	if err != nil {
		log.Printf("userDb create failed, err:%v\n", err)
		return util.BuildErrorInfo("userDb create failed, err:%v", err)
	}
	return nil
}
