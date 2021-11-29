package model

import "blog_service.com/m/pkg/app"

// Article 创建文章 Model
type Article struct {
	*Model
	Title		string 		`json:"title"`
	Desc		string 		`json:"desc"`
	Content		string 		`json:"content"`
	CoverImageUrl string 	`json:"cover_image_url"`
	State 		string 		`json:"state"`
}

type ArticleSwagger struct {
	List []*Article
	Pager *app.Pager
}

// TableName 给文章类型的model重命名
func (a Article)TableName() string {
	return "blog_article"
}
