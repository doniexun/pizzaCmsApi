package controller

import (
	// "encoding/json"
	"github.com/ivpusic/neo"
	"pizzaCmsApi/model"
	"pizzaCmsApi/tools"
)



/**
 * @api {get} /article/:id 获取文章信息
 * @apiName 获取文章信息by path
 * @apiGroup article
 * @apiVersion 1.0.0
 * @apiDescription 用于后台管理员获取文章信息
 * @apiSampleRequest /article/:id
 * @apiParam {int} id文章id
 * @apiSuccess {bool} state 状态
 * @apiSuccess {String} msg 消息
 */
func ArticleGet(ctx *neo.Ctx) (int, error) {
	id := tools.ParseInt(ctx.Req.Params.Get("id"), 0)
	return 200, ctx.Res.Json(model.ArticleGet(id))
}

/**
* @api {PUT} /article/:id 更新article信息
* @apiName 更新article信息
* @apiGroup article
* @apiVersion 1.0.0
* @apiDescription 后台管理员更新文章信息
* @apiSampleRequest /article/:id
* @apiParam {string} title 文章名
* @apiParam {int} id 文章的id
* @apiSuccess {bool} state 状态
* @apiSuccess {String} msg 消息
* @apiPermission admin
 */
func ArticleUpdate(ctx *neo.Ctx) (int, error) {
	var article model.Article
	// err := ctx.Req.JsonBody(&article)
	// if err != nil {
	//   	return 200, ctx.Res.Json(model.ApiJson{State: false, Msg: err.Error() })
	// }

	// b, err := json.Marshal(article)
	// if err != nil {
	// 	log.Printf("json err:", err)
	// } else {
	// 	log.Printf(string(b))
	// }

	id := tools.ParseInt(ctx.Req.Params.Get("id"), 0)
	article.ID = id
	article.Title = ctx.Req.FormValue("title")
	return 200, ctx.Res.Json(model.ArticleUpdate(article))
}

/**
* @api {post} /article/ 创建article信息1
* @apiName 创建article信息1
* @apiGroup article
* @apiVersion 1.0.0
* @apiDescription 创建文章信息
* @apiSampleRequest /article/
* @apiParam {string} title title
* @apiParam {string} brief brief
* @apiParam {string} content content
* @apiSuccess {bool} state 状态
* @apiSuccess {String} msg 消息
* @apiPermission admin
 */
func ArticleCreate(ctx *neo.Ctx) (int, error) {
	var article model.Article

	article.Title = ctx.Req.FormValue("title")
	article.Brief = ctx.Req.FormValue("brief")
	article.Content = ctx.Req.FormValue("content")
	return 200, ctx.Res.Json(model.ArticleCreate(article))
}

/**
* @api {post} /article/page page article
* @apiName page article
* @apiGroup article
* @apiVersion 1.0.0
* @apiDescription page article
* @apiSampleRequest /article/page
* @apiParam {string} kw 关键字
* @apiParam {int} cp cp
* @apiParam {int} mp mp
* @apiParam {nodeid} nodeid 节点id
* @apiSuccess {bool} state 状态
* @apiSuccess {String} msg 消息
* @apiPermission admin
 */
func ArticlePage(ctx *neo.Ctx) (int, error) {
	cp := tools.ParseInt(ctx.Req.FormValue("cp"), 1)
	mp := tools.ParseInt(ctx.Req.FormValue("mp"), 20)
	nodeid := tools.ParseInt(ctx.Req.FormValue("nodeid"), 0)
	kw := ctx.Req.FormValue("kw")

	return 200, ctx.Res.Json(model.ArticlePage(kw, nodeid,cp, mp))
}
