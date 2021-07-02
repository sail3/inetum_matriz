package main

func main() {

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
