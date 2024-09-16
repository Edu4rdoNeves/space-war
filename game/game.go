package game

import (
	"fmt"

	"github.com/Edu4rdoNeves/space-war/assets"
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	player           *Player
	lasers           []*Laser
	meteors          []*Meteor
	meteorSpawnTimer *Timer
	stars            []*Stars
	starsSpawnTimer  *Timer
	score            int
	inMenu           bool
	position         Vector
	btnImg           *ebiten.Image
}

func NewGame() *Game {
	image := assets.StartSprite
	bounds := image.Bounds()
	halfWidth := float64(bounds.Dx()) / 2
	halfHeight := float64(bounds.Dy()) / 2

	position := Vector{
		X: (screenWidth / 2) - halfWidth,
		Y: (screenHeight / 2) - halfHeight,
	}

	g := &Game{
		meteorSpawnTimer: NewTimer(24),
		starsSpawnTimer:  NewTimer(1),
		inMenu:           true,
		position:         position,
		btnImg:           image,
	}
	player := NewPlayer(g)
	g.player = player

	return g
}

func (g *Game) Update() error {
	if g.inMenu {
		if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
			x, y := ebiten.CursorPosition()
			bounds := g.btnImg.Bounds()

			if float64(x) >= g.position.X && float64(x) <= g.position.X+float64(bounds.Dx()) &&
				float64(y) >= g.position.Y && float64(y) <= g.position.Y+float64(bounds.Dy()) {
				g.inMenu = false
				return nil
			}
		}
		return nil
	}
	g.player.Update()

	for _, l := range g.lasers {
		l.Update()
	}

	g.meteorSpawnTimer.Update()
	if g.meteorSpawnTimer.IsReady() {
		g.meteorSpawnTimer.Reset()

		m := NewMeteor()
		g.meteors = append(g.meteors, m)
	}

	for _, m := range g.meteors {
		m.Update()
	}

	for _, m := range g.meteors {
		if m.Collider().Intersects(g.player.Collider()) {
			fmt.Println("You lost")
			g.Reset()
			g.inMenu = true
		}
	}

	for i, m := range g.meteors {
		for j, l := range g.lasers {
			if m.Collider().Intersects(l.Collider()) {
				g.meteors = append(g.meteors[:i], g.meteors[i+1:]...)
				g.lasers = append(g.lasers[:j], g.lasers[j+1:]...)

				g.score += 1

			}
		}
	}

	g.starsSpawnTimer.Update()
	if g.starsSpawnTimer.IsReady() {
		g.starsSpawnTimer.Reset()

		s := NewStarts()
		g.stars = append(g.stars, s)
	}

	for _, s := range g.stars {
		s.Update()
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	if g.inMenu {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(g.position.X, g.position.Y)

		screen.DrawImage(g.btnImg, op)

	} else {

		g.player.Draw(screen)

		for _, l := range g.lasers {
			l.Draw(screen)
		}

		for _, s := range g.stars {
			s.Draw(screen)
		}

		for _, m := range g.meteors {
			m.Draw(screen)
		}
	}

	ScoreTitle(screen, g)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight

}

func (g *Game) AddLasers(laser *Laser) {
	g.lasers = append(g.lasers, laser)
}

func (g *Game) Reset() {
	g.player = NewPlayer(g)
	g.meteors = nil
	g.lasers = nil
	g.meteorSpawnTimer.Reset()
	g.stars = nil
	g.starsSpawnTimer.Reset()
	g.score = 0
}
