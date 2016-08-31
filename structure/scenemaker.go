package structure

import (
	"log"
	"image/color"
//	"fmt"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/examples/common"
)

var err error

const (
	nx0 = 500
	nx1 = 560
	ny0 = 410
	ny1 = 450
	px0 = 40
	px1 = 100
	py0 = 410
	py1 = 450
)

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
	text string
}

type MyScene struct {
	count int
	parts *sceneParts
}

func NewScene(sName, imgName , stext string) *MyScene {
	return &MyScene{
		parts: &sceneParts{sName, Load(imgName), 0, stext},
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
		return  nil
	}

	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		mx, my := ebiten.CursorPosition()
		if inRange(mx, nx0, nx1, true) && inRange(my, ny0, ny1, true) {
			state.SceneManager.GoTo(SceneList[s.parts.name].next)
		}
		if inRange(mx, px0, px1, true) && inRange(my, py0, py1, true) {
			state.SceneManager.GoTo(SceneList[s.parts.name].prev)
		}
	}
	return nil
}


func inRange(f, min, max int, incl bool) (an bool) {
	an = false
	if incl {
		if f == min || f == max	{
			an = true
		}
	}
	if f > min && f < max {
		an = true
	}
	return  an
}

func (s *MyScene) Draw(r *ebiten.Image) error {
	op := &ebiten.DrawImageOptions{}
	// op.ColorM.Scale(0.5, 0.5, 0.5, 1)
	r.DrawImage(s.parts.image, op)

	x := (ScreenWidth - common.ArcadeFont.TextWidth(s.parts.text)*2) / 2
	y := ScreenHeight - 48
	if err := common.ArcadeFont.DrawTextWithShadow(r, s.parts.text, x, y, 2, color.NRGBA{0xa8, 0xbe, 0xf9, 0xff}); err != nil {
		return err
	}
	return nil

	//	mx, my := ebiten.CursorPosition()
	//
	//	msg := fmt.Sprintf(`X:  %d
	//Y: %d`, mx, my)
	//	if err := ebitenutil.DebugPrint(r, msg); err != nil {
	//		return err
	//	}
}