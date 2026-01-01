package model

import "time"

type ArticleM struct {
	Id        int64     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Title     string    `gorm:"column:title;type:varchar(255);not null;comment:文章标题" json:"title"`                                                                         // 文章标题
	Content   string    `gorm:"column:content;type:text;not null;comment:文章内容" json:"content"`                                                                             // 文章内容
	AuthorId  int64     `gorm:"column:author_id;type:int(11);not null;comment:作者ID;index:idx_aid_ctime" json:"authorId"`                                                                       // 作者ID
	CreatedAt time.Time `gorm:"column:created_at;type:datetime(3);not null;default:CURRENT_TIMESTAMP(3);comment:文章创建时间;index:idx_aid_ctime" json:"createdAt"`                                  // 文章创建时间
	UpdatedAt time.Time `gorm:"column:updated_at;type:datetime(3);not null;default:CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3);comment:文章最后修改时间" json:"updatedAt"` // 文章最后修改时间
}

func (ArticleM) TableName() string {
	return "article"
}
