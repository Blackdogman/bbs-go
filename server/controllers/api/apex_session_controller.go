package api

import (
	"github.com/kataras/iris/v12"
	"github.com/mlogclub/simple"
	"github.com/sirupsen/logrus"
)

type ApexSessionController struct {
	Ctx iris.Context
}

func (c *ApexSessionController) GetBy(articleId int64) *simple.JsonResult {
	return simple.JsonErrorCode(404, "还在施工中")
}

func (c *ApexSessionController) GetSessions() *simple.JsonResult {
	logrus.Info("进入了方法")
	return simple.JsonData("{a:1,b:2}")
}
