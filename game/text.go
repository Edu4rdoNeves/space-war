package game

import (
	"fmt"
	"image/color"

	"github.com/Edu4rdoNeves/space-war/assets"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
)

func ScoreTitle(screen *ebiten.Image, g *Game) {
	textString := fmt.Sprintf("Score: %d", g.score)
	textWidth := text.BoundString(assets.FontUi, textString).Dx()
	screenWidth := screen.Bounds().Dx()
	x := (screenWidth - textWidth) / 2

	text.Draw(screen, textString, assets.FontUi, x, 80, color.White)
}
