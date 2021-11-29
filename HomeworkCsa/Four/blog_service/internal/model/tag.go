package model

import (
	"blog_service.com/m/pkg/app"
	"github.com/jinzhu/gorm"
)

// Tag 创建标签Model
type Tag struct {
	*Model
	Name string `json:"name"`
	State uint8 `json:"state"`
}

type TagSwagger struct {
	List []*Tag
	Pager *app.Pager
}

// TableName 为每个tag声明对应的类型
func (t Tag) TableName() string {
	return "blog_tag"
}

/*
		·Model：指定运行 DB 操作的模型实例，默认解析该结构体的名字为表名，格式为大写驼峰转小 写下划线驼峰。若情况特殊，也可以编写该结构体的
		·TableName 方法用于指定其对应返回的表 名。
		·Where：设置筛选条件，接受 map，struct 或 string 作为条件。
		·Offset：偏移量，用于指定开始返回记录之前要跳过的记录数。
		·Limit：限制检索的记录数。
		·Find：查找符合筛选条件的记录。
		·Updates：更新所选字段。
		·Delete：删除数据。
		·Count：统计行为，用于统计模型的记录数。
 */

func (t Tag) Count(db *gorm.DB) (int, error) {
	var count int
	if t.Name != "" {
		db = db.Where("name = ?",t.Name)
	}
	db.Where("state = ?",t.State)
	if err := db.Model(&t).Where("is_del = ?", 0).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

func (t Tag)List(db *gorm.DB, pageOffset, pageSize int) ([]*Tag, error)  {
	var tags []*Tag
	var err error
	if pageOffset >= 0 && pageSize > 0 {
		db = db.Offset(pageOffset).Limit(pageSize)
	}
	if t.Name != "" {
		db = db.Where("name = ?",t.Name)
	}
	db = db.Where("state = ?",t.State)
	if err = db.Where("is_del = ?", 0).Find(&tags).Error; err != nil {
		return nil, err
	}

	return tags, nil
}

func (t Tag)Create(db *gorm.DB) error {
	return db.Create(&t).Error
}

func (t Tag)Update(db *gorm.DB, value interface{}) error {
	return db.Model(&t).Where("id = ? AND is_del = ?",t.ID,0).Update(value).Error
}

func (t Tag)Delete(db *gorm.DB) error {
	return db.Where("id = ? AND is_del = ?",t.Model.ID,0).Delete(&t).Error
}