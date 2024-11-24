package system

import (
	"github.com/gin-gonic/gin"
	"gvb_server/api"
)

type MenuRouter struct {
}

func (s *MenuRouter) InitMenuRouter(Router *gin.Engine) {
	menuRouter := Router.Group("menu")
	menuApi := api.ApiGroupApp.SystemGroup.MenuApi
	menuRouter.POST("createMenu", menuApi.MenuCreateView)
	menuRouter.GET("menuList", menuApi.MenuListView)
	menuRouter.GET("menuNameList", menuApi.MenuNameList)
	menuRouter.PUT("updateMenu", menuApi.MenuUpdateView)
	menuRouter.DELETE("removeMenu", menuApi.MenuRemoveView)
	menuRouter.GET("menuDetail/:id", menuApi.MenuDetailView)
}
