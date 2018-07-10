package models

import (
	"time"

	"github.com/globalsign/mgo/bson"

	"github.com/coderminer/blog/db"
)

type BlogModel struct {
	Id       string     `bson:"_id"`
	Title    string     `bson:"title"`
	Author   string     `bson:"author"`
	Summary  string     `bson:"summary"`
	Original string     `bson:"original"` //原始的markdown格式文本
	Content  string     `bson:"content"`  //渲染之后的html文本
	Date     time.Time  `bson:"date"`
	Tags     []TagModel `bson:"tags"`
	Img      string     `bson:"img"`
}

type TagModel struct {
	Id   string `bson:"_id"`
	Name string `bson:"name"`
}

const (
	database   = "Blog"
	collection = "BlogModel"
)

func NewBlog() *BlogModel {
	return &BlogModel{}
}

func (b *BlogModel) PostBlog(blog *BlogModel) error {
	return db.Insert(database, collection, blog)
}

func (b *BlogModel) GetAllBlogs(page int) ([]BlogModel, error) {
	var blogs []BlogModel
	err := db.FindAllSort(database, collection, "-date", nil, nil, &blogs)
	return blogs, err
}

func (b *BlogModel) GetBlogById(id string) (BlogModel, error) {
	var blog BlogModel
	err := db.FindOne(database, collection, bson.M{"_id": id}, nil, &blog)
	return blog, err
}
