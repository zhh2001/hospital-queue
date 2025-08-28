package models

import "time"

type Patient struct {
	ID         uint      `json:"id"`         // ID
	Number     uint      `json:"number"`     // 患者取号
	Name       string    `json:"name"`       // 患者姓名
	Phone      string    `json:"phone"`      // 手机号码
	Department uint      `json:"department"` // 诊室号
	Status     uint      `json:"status"`     // 0:等待，1:已叫号，2:过号
	CreateAt   time.Time `json:"create_at"`  // 创建时间
	UpdateAt   time.Time `json:"update_at"`  // 更新时间
}
