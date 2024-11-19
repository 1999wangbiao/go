package images

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gvb_server/models/common/res"
)

type ImagesApi struct{}

// 上传多个图片
func (ImagesApi) ImageLoad(c *gin.Context) {

	form, err := c.MultipartForm()
	if err != nil {
		res.FailWithMsg(err.Error(), c)
		return
	}
	fileHeaders := form.File["images"]
	for _, fileHeader := range fileHeaders {
		fmt.Println("Header", fileHeader.Header)
		fmt.Println("Filename", fileHeader.Filename)
		fmt.Println("size", fileHeader.Size)
	}
	res.OKWith(c)
}
