package system

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models/common/request"
	"gvb_server/models/common/response"
	"gvb_server/models/system"
	"gvb_server/utils/pwd"
)

type UserApi struct {
}

// EmailLoginView email登陆
// @Tags      用户管理
// @Summary	  email登陆
// @Description  email登陆
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   json
// @Param body request.EmailLogin
// @Success   200   {object}  response.Response{}
// @Router    /user/emailLogin [post]
func (u *UserApi) EmailLoginView(c *gin.Context) {
	var cr request.EmailLogin
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		response.FailWithError(err, &cr, c)
	}

	var userModel system.UserModel
	err = global.DB.Take(&userModel, "user_name = ?", cr.UserName).Error
	if err != nil {
		//没找到该用户
		global.Log.Warn("该用户不存在")
		response.FailWithMsg("用户名或密码错误", c)
		return
	}
	//校验密码
	isCheck := pwd.ComparePassword(userModel.Password, cr.Password)
	if !isCheck {
		global.Log.Warn("用户名密码错误")
		response.FailWithMsg("用户名或密码错误", c)
		return
	}

}
