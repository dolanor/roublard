package assets

import (
	"bytes"
	_ "embed"
	"image"
	"image/draw"
	"image/png"

	"github.com/g3n/engine/loader/gltf"
	"github.com/g3n/engine/material"
	"github.com/g3n/engine/texture"
)

func wallMat() material.IMaterial {
	model, err := gltf.ParseBin("assets/wood_inlaid_stone_wall_1k.glb")
	if err != nil {
		panic(err)
	}

	mat, err := model.LoadMaterial(0)
	if err != nil {
		panic(err)
	}
	return mat
}

//go:embed floor.png
var floor []byte

func Floor() *texture.Texture2D {
	r := bytes.NewBuffer(floor)
	img, err := png.Decode(r)
	if err != nil {
		panic(err)
	}

	// Converts image to RGBA format
	rgba := image.NewRGBA(img.Bounds())
	if rgba.Stride != rgba.Rect.Size().X*4 {
		panic("unsupported stride")
	}
	draw.Draw(rgba, rgba.Bounds(), img, image.Point{0, 0}, draw.Src)

	tex := texture.NewTexture2DFromRGBA(rgba)

	return tex
}

//go:embed wall.png
var wall []byte

func Wall() *texture.Texture2D {
	r := bytes.NewBuffer(wall)
	img, err := png.Decode(r)
	if err != nil {
		panic(err)
	}

	// Converts image to RGBA format
	rgba := image.NewRGBA(img.Bounds())
	if rgba.Stride != rgba.Rect.Size().X*4 {
		panic("unsupported stride")
	}
	draw.Draw(rgba, rgba.Bounds(), img, image.Point{0, 0}, draw.Src)

	tex := texture.NewTexture2DFromRGBA(rgba)

	return tex
}
