package system

import (
	"github.com/gin-gonic/gin"
	"gvb_server/api"
	"gvb_server/middleware"
)

type UserRouter struct {
}

func (u *UserRouter) InitUserRouter(Router *gin.Engine) {
	userRouter := Router.Group("user")
	userApi := api.ApiGroupApp.SystemGroup.UserApi
	userRouter.POST("email_login", userApi.EmailLoginView)
	userRouter.GET("users", middleware.JWTAuth(), userApi.UserListView)
	//menuRouter.GET("menuNameList", menuApi.MenuNameList)
	//menuRouter.PUT("updateMenu", menuApi.MenuUpdateView)
	//menuRouter.DELETE("removeMenu", menuApi.MenuRemoveView)
	//menuRouter.GET("menuDetail/:id", menuApi.MenuDetailView)
}
