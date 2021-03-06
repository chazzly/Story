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
	"log"
	"sync"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

const ScreenWidth = 640
const ScreenHeight = 480

var rArrow, lArrow *ebiten.Image

func init(){
	rArrow, _, err = ebitenutil.NewImageFromFile("images/rArrow.png", ebiten.FilterLinear)
	if err != nil {
		log.Fatal(err)
	}
	lArrow, _, err = ebitenutil.NewImageFromFile("images/lArrow.png", ebiten.FilterLinear)
	if err != nil {
		log.Fatal(err)
	}

	return
}

type GameState struct {
	SceneManager *SceneManager
	Input        *Input
}

type Game struct {
	once         sync.Once
	sceneManager *SceneManager
	input        Input
}

func NewGame() *Game {
	return &Game{
		sceneManager: NewSceneManager(NewTitleScene()),
	}
}

func (game *Game) Update(r *ebiten.Image) error {
	game.input.Update()
	if err := game.sceneManager.Update(&GameState{
		SceneManager: game.sceneManager,
		Input:        &game.input,
	}); err != nil {
		return err
	}
	if !ebiten.IsRunningSlowly() {
		if err := game.sceneManager.Draw(r); err != nil {
			return err
		}
	}
	return nil
}
