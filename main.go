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
	r.Use(Cors())
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "welcome to pandemic's time",
		})
	})
	NewRotateMatrix().Routes(r)
	// TODO: si queremos implementar nuevas rutas debemos crear nuevas
	// estructuras de tipo Module e implementar la interfaz ModuleInterfaz despues invocarlas.
	return r
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Header.Get("Content-Type") != "application/json" {
			c.AbortWithStatus(http.StatusUnsupportedMediaType)
			return
		}
		c.Next()
	}
}
