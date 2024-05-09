package main

import (
	"eduFair/internal/component"
	"eduFair/internal/config"
	"eduFair/internal/module/feedback"
	"eduFair/internal/module/pengunjung"
	"embed"
	"io/fs"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

//go:embed client/dist/*
var app1 embed.FS

func main() {
	conf := config.Get()

	dbConnnection := component.GetDatabaseConnection(conf)

	pengunjungRepository := pengunjung.NewRepository(dbConnnection)
	feedbackRepository := feedback.NewRepository(dbConnnection)

	pengunjungService := pengunjung.NewService(pengunjungRepository)
	feedbackService := feedback.NewService(feedbackRepository)

	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://10.4.84.142",
		AllowMethods:     "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders:     "Origin,Content-Type,Accept,Authorization,Token",
		ExposeHeaders:    "Content-Length",
		AllowCredentials: true,
	}))
	app.Use(requestid.New())
	app.Use(logger.New(logger.Config{
		Format: "[${locals:requestid}] ${ip} - ${method} ${status} ${path}\n",
	}))

	pengunjung.NewApi(app, pengunjungService)
	feedback.NewApi(app, feedbackService)

	dist, err := fs.Sub(app1, "client/dist")
	if err != nil {
		log.Fatalf("failed to get sub filesystem: %v", err)
	}

	app.Use("/", filesystem.New(filesystem.Config{
		Root:   http.FS(dist),
		Browse: true,
	}))
	app.Use("/presensi", filesystem.New(filesystem.Config{
		Root:   http.FS(dist),
		Browse: true,
	}))
	app.Use("/catalog", filesystem.New(filesystem.Config{
		Root:   http.FS(dist),
		Browse: true,
	}))
	app.Use("/Twinbbon", filesystem.New(filesystem.Config{
		Root:   http.FS(dist),
		Browse: true,
	}))
	app.Use("/feedback", filesystem.New(filesystem.Config{
		Root:   http.FS(dist),
		Browse: true,
	}))
	_ = app.Listen(conf.Srv.Host + ":" + conf.Srv.Port)
}
