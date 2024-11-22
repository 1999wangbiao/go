package menu

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gvb_server/global"
	"gvb_server/models/common/res"
	"gvb_server/models/system"
)

func (MenuApi) MenuRemoveView(c *gin.Context) {
	var cr system.IDList
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	var menuList []system.MenuModel
	rowsAffected := global.DB.Debug().Find(&menuList, cr.IDList).RowsAffected
	if rowsAffected == 0 {
		res.FailWithMsg("菜单不存在", c)
		return
	}
	fmt.Println(cr.IDList)
	//这两个表应该一起删除 作为事务进行处理
	err = global.DB.Transaction(func(tx *gorm.DB) error {
		//删除第三张连接表
		err = tx.Model(&menuList).Association("Banners").Clear()
		if err != nil {
			global.Log.Error(err)
			return err
		}
		//删除菜单表
		err = tx.Delete(&menuList).Error
		if err != nil {
			global.Log.Error(err)
			return err
		}
		return nil
	})
	if err != nil {
		global.Log.Error(err)
		res.FailWithMsg("删除失败", c)
		return
	}
	res.FailWithMsg(fmt.Sprintf("共删除了 %d 个菜单", rowsAffected), c)

}
