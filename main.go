package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Module: estructura para los modulos a implementar
type Module struct {
	URL string
}

// ModuleInterface: interfaz para el modulo de rutas a implementar.
type ModuleInterface interface {
	Routes(gin.Engine)
}

func main() {
	r := InitServe()
	r.Run(":8080")
}

// InitServe: iniciamos el servidor.
func InitServe() *gin.Engine {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "welcome to pandemic's time",
		})
	})
	// TODO: si queremos implementar nuevas rutas debemos crear nuevas
	// estructuras de tipo Module e implementar la interfaz ModuleInterfaz despues invocarlas.
	return r
}
