package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// NewRotateMatrix: Constructor del modulo.
func NewRotateMatrix() Module {
	return Module{
		URL: "/rotate-matrix",
	}
}

// MatrixPayload: Mapper para el payload del request.
type MatrixPayload struct {
	Matrix [][]int `binding:"required" json:"matrix"`
}

// Routes: Define las rutas roportadas para la url declarada en la estructura.
func (c Module) Routes(r *gin.Engine) {
	route := r.Group(c.URL)
	{
		route.POST("/", PostHandler)
	}
}

// PostHandler: Manejador para el metodo POST  de la ruta del modulo
func PostHandler(c *gin.Context) {
	payload := new(MatrixPayload)

	err := c.BindJSON(payload)
	if err != nil {
		panic(err)
	}
	if len(payload.Matrix) != len(payload.Matrix[0]) {
		c.JSON(http.StatusBadRequest, gin.H{
			"messsage": "la matriz ingresada debe ser de tipo NxN",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"matrix": RotarMatriz(payload.Matrix),
	})
}

// RotarMatriz: Rota una matriz NxN
func RotarMatriz(matrizOriginal [][]int) [][]int {
	x := len(matrizOriginal)
	matrizRotada := make([][]int, x)
	for v := range matrizRotada {
		matrizRotada[v] = make([]int, x)
	}
	for indexRow, row := range matrizOriginal {
		for indexCol, item := range row {
			matrizRotada[x-indexCol-1][indexRow] = item
		}
	}
	return matrizRotada
}
