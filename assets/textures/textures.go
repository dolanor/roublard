package textures

import (
	"bytes"
	_ "embed"
	"image"
	"image/draw"
	_ "image/png"

	"github.com/g3n/engine/texture"
)

//go:embed wood_inlaid_stone_wall_nor_gl_1k.png
var woodInlaidNormal []byte

//go:embed wood_inlaid_stone_wall_diff_1k.jpg
var woodInlaidDiffuse []byte

//go:embed wood_inlaid_stone_wall_rough_1k.png
var woodInlaidRough []byte

//go:embed wood_inlaid_stone_wall_disp_1k.png
var woodInlaidDisplace []byte

func bytesToTex(b []byte) *texture.Texture2D {
	r := bytes.NewBuffer(b)
	img, _, err := image.Decode(r)
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

func WoodInlaidNormal() *texture.Texture2D {
	return bytesToTex(woodInlaidNormal)
}

func WoodInlaidDiffuse() *texture.Texture2D {
	return bytesToTex(woodInlaidDiffuse)
}

func WoodInlaidRough() *texture.Texture2D {
	return bytesToTex(woodInlaidRough)
}

func WoodInlaidDisplace() *texture.Texture2D {
	return bytesToTex(woodInlaidDisplace)
}
