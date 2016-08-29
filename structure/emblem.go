package structure

import (
	"fmt"
	"os"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

var emblemImage *ebiten.Image

func init() {
	var err error
	emblemImage, _, err = ebitenutil.NewImageFromFile("images/Emblem.png", ebiten.FilterNearest)
	if err != nil {
		panic(err)
	}
}

var z float64 = 0.02
var d float64 = 0.001

type emblemImageParts struct {
	image *ebiten.Image
	count int
}

type EmblemScene struct {
	count int
	parts *emblemImageParts
}

func NewEmblemScene() *EmblemScene {
	return &EmblemScene{
		parts: &emblemImageParts{emblemImage, 0},
	}
}

func (s *EmblemScene) Update(state *GameState) error {
	// // To make the emblem pulse
	//if z >= 0.1 || z <= 0.01 {
	//	d = d * float64(-1)
	//}
	z = z + d
	d = d + 0.001

	if z > 2.0 {
		os.Exit(0)
	}

	return nil
}

func (s *EmblemScene) Draw(r *ebiten.Image) error {
	w, h := emblemImage.Size()
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(z, z)
	op.GeoM.Translate(ScreenWidth/2 - (float64(w) * z / 2) - (0 * 500), ScreenHeight/2 - (float64(h) * z / 2) - (0 * 500))
	r.DrawImage(emblemImage, op)

	msg := fmt.Sprintf(`Z:  %0.2f
d: %0.3f
off-X:   %0.2f
off-Y:    %0.2f
W:   %d
Y:    %d`, z, d, ScreenWidth/2 - (float64(w) * z / 2) - (0 * 500), ScreenHeight/2 - (float64(h) * z / 2) - (0 * 500), w, h)
	if err := ebitenutil.DebugPrint(r, msg); err != nil {
		return err
	}

	return nil
}