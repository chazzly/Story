package structure

import (
	"log"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

var imageBackground *ebiten.Image

func init() {
	var err error
	imageBackground, _, err = ebitenutil.NewImageFromFile("images/IMG_0462.png", ebiten.FilterLinear)
	if err != nil {
		log.Fatal(err)
	}
}

type openerImageParts struct {
	image *ebiten.Image
	count int
}

type OpenerScene struct {
	count int
	parts *openerImageParts
}

func NewOpenerScene() *OpenerScene {
	return &OpenerScene{
		parts: &openerImageParts{imageBackground, 0},
	}
}

func (s *OpenerScene) Update(state *GameState) error {
	s.count++
	if state.Input.StateForKey(ebiten.KeySpace) == 1 {
		state.SceneManager.GoTo(NewEmblemScene())
		return nil
	}
	return nil
}

func (s *OpenerScene) Draw(r *ebiten.Image) error {
	op := &ebiten.DrawImageOptions{}
	// op.ColorM.Scale(0.5, 0.5, 0.5, 1)
	op.GeoM.Scale(0.3, 0.3)
	r.DrawImage(imageBackground, op)
	return nil
}