package request

type Advert struct {
	Title  string `json:"title" binding:"required" msg:"请输入标题"`       // 广告标题 唯一
	Href   string `json:"href" binding:"required,url" msg:"跳转链接非法"`   // 广告链接
	Images string `json:"images" binding:"required,url" msg:"图片地址非法"` // 图片
	IsShow bool   `json:"is_show" binding:"required" msg:"请选择是否展示"`   // 是否显示
}
