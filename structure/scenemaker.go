package structure

import (
	"log"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

var err error

func Load(imgName string) (sceneImage *ebiten.Image){
	sceneImage, _, err = ebitenutil.NewImageFromFile(imgName, ebiten.FilterLinear)
	if err != nil {
		log.Fatal(err)
	}
	return
}

type sceneParts struct {
	name string
	image *ebiten.Image
	count int
}

type MyScene struct {
	count int
	parts *sceneParts
}

func NewScene(sName, imgName string) *MyScene {
	return &MyScene{
		parts: &sceneParts{sName, Load(imgName), 0},
	}
}

func (s *MyScene) Update(state *GameState) error {
	s.count++

	if state.Input.StateForKey(ebiten.KeySpace) == 1 || state.Input.StateForKey(ebiten.KeyRight) == 1 {
		state.SceneManager.GoTo(SceneList[s.parts.name].next)
		return nil
	}

	if state.Input.StateForKey(ebiten.KeyLeft) == 1 {
		state.SceneManager.GoTo(SceneList[s.parts.name].prev)
		return nil
	}

	if state.Input.StateForKey(ebiten.KeyEscape) == 1 {
		state.SceneManager.GoTo(NewTitleScene())
	}
	return nil
}

func (s *MyScene) Draw(r *ebiten.Image) error {
	op := &ebiten.DrawImageOptions{}
	// op.ColorM.Scale(0.5, 0.5, 0.5, 1)
	op.GeoM.Scale(0.3, 0.3)
	r.DrawImage(s.parts.image, op)
	return nil
}