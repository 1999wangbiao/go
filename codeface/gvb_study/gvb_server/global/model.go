package global

import (
	"time"
)

type GVA_MODEL struct {
	ID        uint      `gorm:"primarykey"` // 主键ID
	CreatedAt time.Time `json:"created_at"` // 创建时间
	UpdatedAt time.Time `json:"updated_at"` // 更新时间
}
