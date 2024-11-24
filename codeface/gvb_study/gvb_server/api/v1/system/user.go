package system

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models/common/request"
	"gvb_server/models/common/response"
	"gvb_server/models/ctype"
	"gvb_server/models/system"
	"gvb_server/service/common"
	"gvb_server/utils"
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
// @Router    /user/email_Login [post]
func (u *UserApi) EmailLoginView(c *gin.Context) {
	var cr request.EmailLogin
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		response.FailWithError(err, &cr, c)
		return
	}
	var userModel system.UserModel
	//err = global.DB.Where("user_name = ?", cr.UserName).Select("user_name").First(&userModel).Error
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
	//jwt.Claims()
	j := &utils.JWT{SigningKey: []byte(global.Config.Jwt.SigningKey)} // 唯一签名
	claims := j.CreateClaims(request.BaseClaims{
		ID:          userModel.ID,
		Username:    userModel.UserName,
		NickName:    userModel.NickName,
		AuthorityId: int(userModel.Role),
	})
	token, err := j.CreateToken(claims)
	if err != nil {
		global.Log.Fatal("设置token失败")
		response.FailWithMsg("设置token失败", c)
		return
	}
	fmt.Println("xixi")
	response.OkWithDetailed(response.LoginResponse{
		User:      system.UserModel{UserName: userModel.UserName},
		Token:     token,
		ExpiresAt: claims.RegisteredClaims.ExpiresAt.Unix() * 1000,
	}, "登录成功", c)

}

// UserListView 用户列表
// @Tags      用户管理
// @Summary	  用户列表
// @Description  用户列表
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   json
// @Success   200   {object}  response.LoginResponse
// @Router    /user/user_list [get]
func (u *UserApi) UserListView(c *gin.Context) {
	//如何判断是否是管理员(token)
	token := c.Request.Header.Get("token")
	if token == "" {
		response.FailWithMsg("未登录,没有携带token", c)
		return
	}
	j := utils.NewJWT()
	// parseToken 解析token包含的信息
	claims, err := j.ParseToken(token)
	if err != nil {
		response.FailWithDetailed(err, "token验证失败", c)
		return
	}

	var cr system.Page
	err = c.ShouldBindQuery(&cr)
	if err != nil {
		response.FailWithCode(response.ArgumentError, c)
		return
	}
	var users []system.UserModel
	list, count, _ := common.ComList(system.UserModel{}, common.Option{Page: cr})
	for _, user := range list {
		if ctype.Role(claims.BaseClaims.AuthorityId) != ctype.PermissionAdmin {
			//非  管理员
			fmt.Println("非管理员")
			user.UserName = ""
		}
		user.Email = utils.DesensitizationEmail(user.Email)
		user.Tel = utils.DesensitizationTel(user.Tel)
		users = append(users, user)
	}
	response.OKWithList(users, count, c)
	return
}
