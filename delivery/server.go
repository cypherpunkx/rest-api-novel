package delivery

import (
	"github.com/gin-gonic/gin"
)

type application struct {
	engine *gin.Engine
}

func (app *application) Run() {
	if err := SetupRouter(app.engine); err != nil {
		panic("Aplication error")
	}
}

func Server() *application {
	router := gin.Default()

	return &application{
		engine: router,
	}

}
