package image

import (
	"image"
	"image/color"
	"image/jpeg"
	"image/png"
	"os"
	"strings"
)

func ReadImage(filepath string) ([][]color.RGBA, error) {
	var img image.Image

	// Abrir o arquivo de imagem
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	if strings.HasSuffix(filepath, "png") {
		img, err = png.Decode(file)
	} else {
		img, err = jpeg.Decode(file)
	}

	if err != nil {
		return nil, err
	}

	// Criar a matriz de cores
	width := img.Bounds().Size().X
	height := img.Bounds().Size().Y
	colors := make([][]color.RGBA, height)
	for y := 0; y < height; y++ {
		row := make([]color.RGBA, width)
		for x := 0; x < width; x++ {
			// Obter a cor do pixel na posição (x, y) e adicioná-la à matriz
			r, g, b, a := img.At(x, y).RGBA()
			row[x] = color.RGBA{uint8(r >> 8), uint8(g >> 8), uint8(b >> 8), uint8(a >> 8)}
		}
		colors[y] = row
	}

	return colors, nil
}
