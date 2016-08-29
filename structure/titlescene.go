// Copyright 2014 Hajime Hoshi
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package structure

import (
	"image/color"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/examples/common"
)

var titleBackground *ebiten.Image

func init() {
	var err error
	titleBackground, _, err = ebitenutil.NewImageFromFile("images/Title.png", ebiten.FilterNearest)
	if err != nil {
		panic(err)
	}
}

type titleImageParts struct {
	name string
	image *ebiten.Image
	count int
}

func (t *titleImageParts) Len() int {
	w, h := t.image.Size()
	return (ScreenWidth/w + 1) * (ScreenHeight/h + 2)
}

func (t *titleImageParts) Dst(i int) (x0, y0, x1, y1 int) {
	w, h := t.image.Size()
	i, j := i%(ScreenWidth/w+1), i/(ScreenWidth/w+1)-1
	dx := (-t.count / 4) % w
	dy := (t.count / 4) % h
	dstX := i*w + dx
	dstY := j*h + dy
	return dstX, dstY, dstX + w, dstY + h
}

func (t *titleImageParts) Src(i int) (x0, y0, x1, y1 int) {
	w, h := t.image.Size()
	return 0, 0, w, h
}

type TitleScene struct {
	count int
	parts *titleImageParts
}

func NewTitleScene() *TitleScene {
	return &TitleScene{
		parts: &titleImageParts{"Title", titleBackground, 0},
	}
}

func (s *TitleScene) Update(state *GameState) error {
	s.count++
	if state.Input.StateForKey(ebiten.KeySpace) == 1 {
		state.SceneManager.GoTo(SceneList[s.parts.name].next)
		return nil
	}
	return nil
}

func (s *TitleScene) Draw(r *ebiten.Image) error {
	if err := s.drawTitleBackground(r, s.count); err != nil {
		return err
	}
	if err := drawLogo(r, "Thunder Snowflake"); err != nil {
		return err
	}

	message := "PRESS SPACE TO START"
	x := (ScreenWidth - common.ArcadeFont.TextWidth(message)) / 2
	y := ScreenHeight - 48
	if err := common.ArcadeFont.DrawTextWithShadow(r, message, x, y, 1, color.NRGBA{0x80, 0, 0, 0xff}); err != nil {
		return err
	}
	return nil
}

func (s *TitleScene) drawTitleBackground(r *ebiten.Image, c int) error {
	s.parts.count = c
	return r.DrawImage(titleBackground, &ebiten.DrawImageOptions{
		ImageParts: s.parts,
	})
}

func drawLogo(r *ebiten.Image, str string) error {
	scale := 4
	textWidth := common.ArcadeFont.TextWidth(str) * scale
	x := (ScreenWidth - textWidth) / 2
	y := 32
	return common.ArcadeFont.DrawTextWithShadow(r, str, x, y, scale, color.NRGBA{0x00, 0x00, 0x80, 0xff})
}
