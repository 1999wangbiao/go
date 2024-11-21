package system

type Page struct {
	Page  int    `form:"page"` //第几页
	Key   string `form:"key"`
	Limit int    `form:"limit"` //每一页最多几个
	Sort  string `form:"sort"`
}

type IDList struct {
	IDList []uint `json:"id_list"`
}
