package main

import (
	"github.com/iris-contrib/template/django"
	"github.com/kataras/iris"
)

func main() {

	// iris.TemplateString
	// Executes and parses the template but instead of rendering to the client, it returns the contents.
	// Useful when you want to send a template via e-mail or anything you can imagine.

	// Note that: iris.TemplateString can be called outside of the context also

	iris.UseTemplate(django.New()).Directory("./templates", ".html")

	iris.Get("/", func(ctx *iris.Context) {
		// THIS WORKS WITH ALL TEMPLATE ENGINES, but I am not doing the same example for all engines again :) (the same you can do with templates using the iris.ResponseString)

		rawHtmlContents := iris.TemplateString("mypage.html", map[string]interface{}{"username": "iris", "is_admin": true}, iris.RenderOptions{"charset": "UTF-8"}) // defaults to UTF-8 already
		ctx.Log(rawHtmlContents)
		ctx.Write("The Raw HTML is:\n%s", rawHtmlContents)
	})

	iris.Listen(":8080")
}
