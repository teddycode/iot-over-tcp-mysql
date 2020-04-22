package models

import "github.com/jinzhu/gorm"

// 数据表结构
type Item struct {
	ID        int     `gorm:"primary_key" json:"id"`
	CreatedOn int     `json:"created_on"`
	Did       int     `json:"did"`   // 节点ID
	Light     float32 `json:"light"` // 光强
	Mq2       float32 `json:"mq2"`   // Mq2传感器值
	Mq135     float32 `json:"mq135"` // Mq135 传感器
	Temp      float32 `json:"temp"`  //温度
	Wet       float32 `json:"wet"`   //湿度
}

// 增加
func NewItem(item *Item) (bool, error) {
	err := db.Create(item).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

// 查询100 条数据
func FindItemsByDId(did int) (*[]Item, error) {
	items := new([]Item)
	err := db.Limit(100).Order("created_on desc").Where("did = ?", did).Find(items).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return items, err
	}
	return items, err
}
