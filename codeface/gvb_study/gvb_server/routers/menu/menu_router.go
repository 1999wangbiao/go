package menu

import (
	"github.com/gin-gonic/gin"
	"gvb_server/api"
)

type MenuRouter struct {
}

func (s *MenuRouter) InitMenuRouter(Router *gin.Engine) {
	menuRouter := Router.Group("menu")
	menuApi := api.ApiGroupApp.MenuApi
	menuRouter.POST("create", menuApi.MenuCreateView)
	menuRouter.GET("list", menuApi.MenuListView)
	menuRouter.GET("name_list", menuApi.MenuNameList)
	menuRouter.PUT("update", menuApi.MenuUpdateView)
}
