package scheduler

import (
	"github.com/robfig/cron"
	"github.com/sirupsen/logrus"
)

func Start() {
	// c := cron.New()

	// Generate RSS
	// addCronFunc(c, "@every 30m", func() {
	// 	services.ArticleService.GenerateRss()
	// 	services.TopicService.GenerateRss()
	// 	services.ProjectService.GenerateRss()
	// })

	// Generate sitemap
	// addCronFunc(c, "0 0 4 ? * *", func() {
	// 	sitemap.Generate()
	// })

	// c.Start()
}

func addCronFunc(c *cron.Cron, sepc string, cmd func()) {
	err := c.AddFunc(sepc, cmd)
	if err != nil {
		logrus.Error(err)
	}
}
