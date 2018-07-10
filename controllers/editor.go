package controllers

import (
	"fmt"
	"strings"
	"time"

	"github.com/astaxie/beego"
	"github.com/coderminer/blog/models"
	"github.com/coderminer/blog/utils"
	"github.com/globalsign/mgo/bson"
	stripmd "github.com/writeas/go-strip-markdown"
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
		Summary:  getDes(origin),
		Content:  content,
		Date:     time.Now(),
		Author:   "CoderMiner",
		Img:      "http://7xplrz.com1.z0.glb.clouddn.com/jianshu/android/133H_1.jpg",
	}

	resp := utils.Response{
		Status: 0,
	}
	defer resp.WriteJson(this.Ctx.ResponseWriter)

	err := blog.PostBlog(blog)
	if err != nil {
		resp = utils.Response{
			Status: -1,
			Msg:    "post blog error",
		}
	}

}

func getDes(origin string) string {
	original := stripmd.Strip(origin)
	des := strings.Replace(original, "\n", "", -1)
	desrune := []rune(des)
	return string(desrune[:150])
}
