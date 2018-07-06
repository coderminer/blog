package controllers

import (
	"fmt"
	"time"

	"github.com/astaxie/beego"
	"github.com/coderminer/blog/models"
	"github.com/globalsign/mgo/bson"
)

type EditorController struct {
	beego.Controller
}

type BlogData struct {
	Title   string `json:"title"`
	Origin  string `json:"origin"`
	Content string `json:"content"`
}

func (this *EditorController) Get() {
	this.TplName = "editor.html"
}

func (this *EditorController) Post() {
	title := this.GetString("title")
	origin := this.GetString("origin")
	content := this.GetString("content")
	fmt.Println("post data", title, origin, content)

	blog := &models.BlogModel{
		Id:       bson.NewObjectId().Hex(),
		Title:    title,
		Original: origin,
		Content:  content,
		Date:     time.Now(),
	}

	blog.PostBlog(blog)

	this.ServeJSON()
}
