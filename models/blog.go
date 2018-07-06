package models

import (
	"time"

	"github.com/coderminer/blog/db"
)

type BlogModel struct {
	Id       string     `bson:"_id"`
	Title    string     `bson:"title"`
	Summary  string     `bson:"summary"`
	Original string     `bson:"original"` //原始的markdown格式文本
	Content  string     `bson:"content"`  //渲染之后的html文本
	Date     time.Time  `bson:"date"`
	Tags     []TagModel `bson:"tags"`
}

type TagModel struct {
	Id   string `bson:"_id"`
	Name string `bson:"name"`
}

const (
	database   = "Blog"
	collection = "BlogModel"
)

func (b *BlogModel) PostBlog(blog *BlogModel) error {
	return db.Insert(database, collection, blog)
}
