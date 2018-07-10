package controllers

import (
	"github.com/coderminer/blog/models"

	"github.com/astaxie/beego"
)

type BlogController struct {
	beego.Controller
}

func (this *BlogController) Get() {
	blogs, err := models.NewBlog().GetAllBlogs(0)
	if err == nil {
		this.Data["Blogs"] = blogs
	}

	this.TplName = "index.html"
}

func (this *BlogController) Detail() {
	id := this.Ctx.Input.Param(":id")
	data, err := models.NewBlog().GetBlogById(id)
	if err == nil {
		this.Data["Content"] = data
	} else {
		this.Abort("404")
	}

	this.TplName = "detail.html"
}
