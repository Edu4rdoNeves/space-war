package game

import (
	"math/rand"

	"github.com/Edu4rdoNeves/space-war/assets"
	"github.com/hajimehoshi/ebiten/v2"
)

type Stars struct {
	image    *ebiten.Image
	speed    float64
	position Vector
}

func NewStarts() *Stars {
	image := assets.StarsSprites[rand.Intn(len(assets.StarsSprites))]
	speed := float64(13)

	position := Vector{
		X: rand.Float64() * screenWidth,
		Y: -100,
	}

	return &Stars{
		image:    image,
		speed:    speed,
		position: position,
	}
}

func (s *Stars) Update() {
	s.position.Y += s.speed
}

func (s *Stars) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}

	op.GeoM.Translate(s.position.X, s.position.Y)

	screen.DrawImage(s.image, op)
}
