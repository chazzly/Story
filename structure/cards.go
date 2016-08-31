package structure

import (
	"fmt"
	"os"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

var deckImage *ebiten.Image

func init() {
	var err error
	deckImage, _, err = ebitenutil.NewImageFromFile("images/Deck-72x100x16.gif", ebiten.FilterNearest)
	if err != nil {
		panic(err)
	}
}

var c int = 1

type deckImageParts struct {
	image *ebiten.Image
	count int
}

type CardScene struct {
	count int
	parts *deckImageParts
}

func NewDeckScene() *CardScene {
	return &CardScene{
		parts: &deckImageParts{deckImage, 0},
	}
}

func (s *CardScene) Update(state *GameState) error {
	if c > 51.0 {
		os.Exit(0)
	}

	if state.Input.StateForKey(ebiten.KeyEscape) == 1 {
		state.SceneManager.GoTo(NewEmblemScene())
	}

	if state.Input.StateForKey(ebiten.KeySpace) == 1 {
		c++
		err = s.Draw(s.parts.image)
		return nil
	}

	return nil
}

func (s *CardScene) Draw(r *ebiten.Image) error {
	w, h := deckImage.Size()
	op := &ebiten.DrawImageOptions{
		ImageParts: cardImageParts(c),
	}
	op.GeoM.Translate(ScreenWidth/2 - (float64(w) * z / 2) - (0 * 500), ScreenHeight/2 - (float64(h) * z / 2) - (0 * 500))
	r.DrawImage(deckImage, op)


	// FIX - WTF!!!
	msg := fmt.Sprintf("c:  %d\n", c)
	msg += fmt.Sprintln(op.ImageParts.Src(c))
	msg += fmt.Sprintln(op.ImageParts.Dst(c))

	if err := ebitenutil.DebugPrint(r, msg); err != nil {
		return err
	}
	return nil
}

type cardImageParts int

const cardwidth int = 936/13
const cardheight int = 100


func (b cardImageParts) Len() int {
	return 13 * 5
}

func (b cardImageParts) Dst(i int) (x0, y0, x1, y1 int) {
	x0, y0 = (ScreenWidth-cardwidth)/8, (ScreenHeight-cardheight)/8
	x1, y1 = x0 + cardwidth, y0 + cardheight

	return
}

func (b cardImageParts) Src(i int) (x0, y0, x1, y1 int) {
	x0 = (cardwidth * (i-1)) - 450
	y0 = 500 - (cardheight * (i/13))
	x1, y1 = x0 + cardwidth, y0 - cardheight
	return
}