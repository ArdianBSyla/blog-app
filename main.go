package main

import (
	"flag"
	"log"

	"github.com/personal/blog-app/app"
	"github.com/personal/blog-app/app/controller"
	"github.com/personal/blog-app/app/helper"
	"github.com/personal/blog-app/app/service"
	"github.com/personal/blog-app/config"

	"go.uber.org/dig"
)

func main() {
	port := flag.String("port", "3000", "Port to run server on")
	flag.Parse()

	server := app.NewServer(container())
	log.Println(server.Serve(*port))
}

func container() *dig.Container {
	container := dig.New()
	container.Provide(config.NewConfig)
	container.Provide(app.NewChiRouter)
	container.Provide(helper.NewGormDB)

	// controllers
	container.Provide(controller.NewBlogPostController, dig.Group("controller"))
	container.Provide(controller.NewPostCommentController, dig.Group("controller"))

	// services
	container.Provide(service.NewBlogPostService)
	container.Provide(service.NewPostCommentService)

	return container
}
