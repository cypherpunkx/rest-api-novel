package delivery

import (
	"enigmacamp.com/rest-api-novel/config"
	"enigmacamp.com/rest-api-novel/delivery/controller"
	"enigmacamp.com/rest-api-novel/repository"
	"enigmacamp.com/rest-api-novel/service"
	"github.com/gin-gonic/gin"
)

func SetupRouter(router *gin.Engine) error {

	repository := repository.NewNovelRepository(config.DB)

	service := service.NewNovelService(repository)

	controller := controller.NewNovelController(service)

	v1 := router.Group("/api/v1")
	{
		novels := v1.Group("/novels")
		{
			novels.GET("/", controller.GetListNovel)
			novels.GET("/:id", controller.GetDetailNovel)
			novels.POST("/", controller.CreateNovel)
			novels.PUT("/:id", controller.UpdateNovel)
			novels.DELETE("/:id", controller.DeleteNovel)
		}
	}

	return router.Run()
}
