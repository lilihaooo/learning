package user

import (
	"crypto/md5"
	"encoding/hex"
	"learning2.0/model/model_view"
	"learning2.0/util"
)

func checkParam(res *model_view.UserReq) error {
	if res.Username == "" {
		return util.BuildErrorInfo("用户名不能为空")
	}
	if res.Password == "" {
		return util.BuildErrorInfo("密码不能为空")
	}
	return nil
}

// 加密密码
func EncryptPassword(data []byte) (result string) {
	secret := "127.0.0.1:5050/learning"
	h := md5.New()
	h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum(data))
}
