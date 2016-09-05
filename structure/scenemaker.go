package structure

import (
	"log"
	"image/color"
//	"fmt"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/examples/common"
	"github.com/kr/text"
	"strings"
)

var err error

const (
	// Arrow - w: 69, h:50
	awidth = 69
	aheight = 50
	nx0 = ScreenWidth - 80
	nx1 = nx0 + awidth
	ny0 = ScreenHeight - 55
	ny1 = ny0 + aheight
	px0 = 5
	px1 = px0 + awidth
	py0 = ny0
	py1 = ny0 + aheight
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

	// TODO: Figure out how to make cursor change over the arrows
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
	r.DrawImage(s.parts.image, op)

	var txt []string

	w := len(s.parts.text)

	if w > 90 {
		// TODO: Create an actual error object and return that.
	}

	txt = strings.Split(text.Wrap(s.parts.text, 30), "\n")

	for l, t := range txt {
		x := (ScreenWidth - common.ArcadeFont.TextWidth(t) * 2) / 2
		y := ScreenHeight - (16 * (3-l))
		if err := common.ArcadeFont.DrawTextWithShadow(r, t, x, y, 2, color.NRGBA{0xa8, 0xbe, 0xf9, 0xff}); err != nil {
			return err
		}
	}

	rop := &ebiten.DrawImageOptions{}
	rop.GeoM.Translate( nx0, ny0 )
	r.DrawImage(rArrow, rop)

	lop := &ebiten.DrawImageOptions{}
	lop.GeoM.Translate( px0, py0)
	r.DrawImage(lArrow, lop)

	return nil
}