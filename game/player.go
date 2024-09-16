package game

import (
	"github.com/Edu4rdoNeves/space-war/assets"
	"github.com/hajimehoshi/ebiten/v2"
)

type Player struct {
	image             *ebiten.Image
	position          Vector
	game              *Game
	laserLoadingTimer *Timer
}

func NewPlayer(game *Game) *Player {
	image := assets.PlayerSprite

	bounds := image.Bounds()

	halfWidth := float64(bounds.Dx()) / 2

	position := Vector{
		X: (screenWidth / 2) - halfWidth,
		Y: 500,
	}

	return &Player{
		game:              game,
		image:             image,
		position:          position,
		laserLoadingTimer: NewTimer(12),
	}
}

func (p *Player) Update() {
	speed := 6.0
	bounds := p.image.Bounds()
	halfWidth := float64(bounds.Dx()) / 2

	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		p.position.X -= speed
		if p.position.X < 0 {
			p.position.X = 0
		}
	}

	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		p.position.X += speed
		if p.position.X+float64(bounds.Dx()) > screenWidth {
			p.position.X = screenWidth - float64(bounds.Dx())
		}
	}

	p.laserLoadingTimer.Update()
	if ebiten.IsKeyPressed(ebiten.KeySpace) && p.laserLoadingTimer.IsReady() {
		p.laserLoadingTimer.Reset()

		spawnLaserPosition := Vector{
			p.position.X + halfWidth,
			p.position.Y,
		}

		laser := NewLaser(spawnLaserPosition)
		p.game.AddLasers(laser)
	}

}

func (p *Player) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}

	op.GeoM.Translate(p.position.X, p.position.Y)

	screen.DrawImage(p.image, op)
}

func (p *Player) Collider() Rect {
	bounds := p.image.Bounds()

	return NewRect(p.position.X,
		p.position.Y,
		float64(bounds.Dx()),
		float64(bounds.Dy()))

}
