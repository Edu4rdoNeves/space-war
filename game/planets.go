package game

import (
	"math"
	"math/rand"

	"github.com/Edu4rdoNeves/space-war/assets"
	"github.com/hajimehoshi/ebiten/v2"
)

type Planets struct {
	image    *ebiten.Image
	speed    float64
	position Vector
	scale    float64
	opacity  float64
}

func NewPlanets(existingPlanets []*Planets) *Planets {
	image := assets.PlanetsSprites[rand.Intn(len(assets.PlanetsSprites))]

	distanceFactor := rand.Float64()*0.5 + 0.5

	scale := distanceFactor
	speed := 2.0 * distanceFactor
	opacity := distanceFactor
	var position Vector
	isOverlapping := true

	for isOverlapping {
		position = Vector{
			X: rand.Float64() * screenWidth,
			Y: -500,
		}

		isOverlapping = false
		for _, planet := range existingPlanets {
			if distance(position, planet.position) < 100 {
				isOverlapping = true
				break
			}
		}
	}

	return &Planets{
		image:    image,
		speed:    speed,
		position: position,
		scale:    scale,
		opacity:  opacity,
	}
}

func (p *Planets) Update() {
	p.position.Y += p.speed
}

func (p *Planets) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}

	// Aplicando a escala (tamanho baseado na distância)
	op.GeoM.Scale(p.scale, p.scale)

	// Aplicando a posição
	op.GeoM.Translate(p.position.X, p.position.Y)

	// Aplicando a opacidade/transparência (alpha)
	op.ColorM.Scale(1, 1, 1, p.opacity)

	screen.DrawImage(p.image, op)
}

func distance(a, b Vector) float64 {
	return math.Sqrt(math.Pow(a.X-b.X, 2) + math.Pow(a.Y-b.Y, 2))
}
