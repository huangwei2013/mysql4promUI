package boot

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"mysql4promUI/app/controller"
)

func InitRouter() {
	urlPath := g.Config().GetString("url-path")
	domain := g.Config().GetString("domain")
	s := g.Server()
	s.BindHandler(urlPath+"/", controller.Index)
	s.BindHandler(urlPath+"/index", controller.Index)
	s.BindHandler(urlPath+"/main.html", controller.Main)

	s.Domain(domain)

	s.Group(urlPath+"/", func(g *ghttp.RouterGroup) {
		tRuleLabelsController := new(controller.TRuleLabelsController)
		g.ALL("/rule/labels.html", tRuleLabelsController)
		g.ALL("/rule/labels", tRuleLabelsController)
		g.POST("/rule/labels/page", tRuleLabelsController.Page)
		g.GET("/rule/labels/get/{id}", tRuleLabelsController.Get)
		g.POST("/rule/labels/save", tRuleLabelsController.Save)
		g.POST("/rule/labels/update", tRuleLabelsController.Update)
		g.POST("/rule/labels/delete", tRuleLabelsController.Delete)

		tRuleAnnotationsController := new(controller.TRuleAnnotationsController)
		g.ALL("/rule/annotations.html", tRuleAnnotationsController)
		g.ALL("/rule/annotations", tRuleAnnotationsController)
		g.POST("/rule/annotations/page", tRuleAnnotationsController.Page)
		g.GET("/rule/annotations/get/{id}", tRuleAnnotationsController.Get)
		g.POST("/rule/annotations/save", tRuleAnnotationsController.Save)
		g.POST("/rule/annotations/update", tRuleAnnotationsController.Update)
		g.POST("/rule/annotations/delete", tRuleAnnotationsController.Delete)

		tConfigsController := new(controller.TConfigsController)
		g.ALL("/configs", tConfigsController)
		g.POST("/configs/page", tConfigsController.Page)
		g.GET("/configs/get/{id}", tConfigsController.Get)
		g.POST("/configs/save", tConfigsController.Save)
		g.POST("/configs/update", tConfigsController.Update)
		g.POST("/configs/delete", tConfigsController.Delete)

		tRulesController := new(controller.TRulesController)
		g.ALL("/rules.html", tRulesController)
		g.ALL("/rules", tRulesController)
		g.POST("/rules/page", tRulesController.Page)
		g.GET("/rules/get/{id}", tRulesController.Get)
		g.POST("/rules/save", tRulesController.Save)
		g.POST("/rules/update", tRulesController.Update)
		g.POST("/rules/delete", tRulesController.Delete)

	})
}
