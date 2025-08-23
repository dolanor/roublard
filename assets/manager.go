package assets

import (
	"sync"

	"github.com/g3n/engine/material"
	"github.com/g3n/engine/math32"
	"github.com/g3n/engine/texture"
)

type MaterialID string

type MaterialManager struct {
	mu    sync.RWMutex
	cache map[MaterialID]material.IMaterial
}

func NewMaterialManager() *MaterialManager {
	mm := MaterialManager{
		cache: map[MaterialID]material.IMaterial{},
	}

	floor := matFromTex(Floor())
	mm.Add(MaterialID("floor"), floor)

	wall := wallMat()
	mm.Add(MaterialID("wall"), wall)

	return &mm
}

func (mm *MaterialManager) Add(id MaterialID, mat material.IMaterial) {
	mm.mu.Lock()
	defer mm.mu.Unlock()

	mm.cache[id] = mat
}

func (mm *MaterialManager) Get(id MaterialID) material.IMaterial {
	mm.mu.RLock()
	defer mm.mu.RUnlock()

	mat, ok := mm.cache[id]
	if !ok {
		panic("no material found for: " + id)
	}
	return mat
}

func matFromTex(tex *texture.Texture2D) material.IMaterial {
	color := math32.NewColor("White")

	mat := material.NewStandard(color)
	mat.AddTexture(tex)
	return mat
}
