package model

import "github.com/jinzhu/gorm"

type Tag struct {
	*Model
	Name  string `json:"name"`
	State uint8  `json:"state"`
}

// TableName 重写TableName指定其对应返回的表名
func (t Tag) TableName() string {
	return "blog_tag"
}

// Count 查找标签数量
func (t Tag) Count(db *gorm.DB) (int, error) {
	var count int
	if t.Name != "" {
		// 使用name过滤
		db = db.Where("name = ?", t.Name)
	}
	// 使用state过滤
	db = db.Where("state = ?", t.State)
	// 统计可使用的标签
	if err := db.Model(&t).Where("is_del = ?", 0).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

// List 返回标签列表
func (t Tag) List(db *gorm.DB, pageOffset, pageSize int) ([]*Tag, error) {
	var tags []*Tag
	var err error
	if pageOffset >= 0 && pageSize > 0 {
		// 偏移并限制检索的记录数
		db = db.Offset(pageOffset).Limit(pageSize)
	}
	if t.Name != "" {
		// 使用name过滤
		db = db.Where("name = ?", t.Name)
	}
	// 使用state过滤
	db = db.Where("state = ?", t.State)
	// 查找可使用标签的所有记录
	if err = db.Where("is_del = ?", 0).Find(&tags).Error; err != nil {
		return nil, err
	}
	return tags, nil
}

// Create 创建标签记录
func (t Tag) Create(db *gorm.DB) error {
	return db.Create(&t).Error
}

// Update 更新标签记录
func (t Tag) Update(db *gorm.DB, values interface{}) error {
	// 使用字典形式可以进行更新零值
	if err := db.Model(t).Where("id = ? AND is_del = ?", t.ID, 0).Updates(values).Error; err != nil {
		return err
	}
	return nil
	// 在 GORM 中使用 struct 类型传入进行更新时，GORM 是不会对值为零值的字段进行变更
	// return db.Model(&Tag{}).Where("id = ? AND is_del = ?", t.ID, 0).Update(t).Error
}

// Delete 删除标签
func (t Tag) Delete(db *gorm.DB) error {
	return db.Where("id = ? AND is_del = ?", t.Model.ID, 0).Delete(&t).Error
}
