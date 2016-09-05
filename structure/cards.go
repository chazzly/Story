package structure

import (
	"fmt"
	// "log"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/engoengine/math"
)

var suits = [5]string{"Hearts", "Diamonds", "Clubs", "Spades", "Backs"}
var faces = [13]string{"Ace", "2", "3", "4", "5", "6", "7", "8", "9", "10", "Jack", "Queen", "King"}

type card struct {
	suit   string
	value  int
	imgLoc [2]int
	img     *ebiten.Image
}

var deckImage, bgImage *ebiten.Image
var deck []card
var c int = 1

const cardwidth int = 936 / 13
const cardheight int = 100


var piles [5][2]int = [5][2]int{
	[2]int{44, 16},
	[2]int{164, 16},
	[2]int{284, 16},
	[2]int{404, 16},
	[2]int{524, 16},
}


func init() {
	var err error
	deckImage, _, err = ebitenutil.NewImageFromFile("images/Deck-72x100x16.gif", ebiten.FilterNearest)
	if err != nil {
		panic(err)
	}

	bgImage, _, err = ebitenutil.NewImageFromFile("images/background.png", ebiten.FilterNearest)
	if err != nil {
		panic(err)
	}

	var x, y int

	for i := 1; i <= 13; i++ {
		for n, s := range suits {
			x = 72 * (i - 1)
			y = n * 100
			loc := [2]int{x,y}
			deck = append(deck, card{suit: s, value: i, imgLoc: loc })
		}
	}
}


type bgImageParts struct {
	image *ebiten.Image
	count int
}

type CardScene struct {
	count int
	parts *bgImageParts
}

func NewDeckScene() *CardScene {
	return &CardScene{
		parts: &bgImageParts{bgImage, 0},
	}
}

func (s *CardScene) Update(state *GameState) error {
	if state.Input.StateForKey(ebiten.KeyEscape) == 1 {
		state.SceneManager.GoTo(NewEmblemScene())
	}

	// TODO: why does it not show up until the second "spcace", and with the first card, but in the secons spot.
	if state.Input.StateForKey(ebiten.KeySpace) == 1 {
		y, _ := math.Modf(float32((c - 1) / 5))
		s.NextCard(c, c-int(y*5)-1)
		c++
	}
	return nil
}

func (s *CardScene) Draw(r *ebiten.Image) error {
	// w, h := deckImage.Size()
	bop := &ebiten.DrawImageOptions{}
	r.DrawImage(bgImage, bop)


	// FIX - WTF!!!
	// msg := fmt.Sprintf("c: %d\n", c)

	//if err := ebitenutil.DebugPrint(r, msg); err != nil {
	//	return err
	//}
	return nil
}

func (s *CardScene) NextCard(c, p int) error {
	op := &ebiten.DrawImageOptions{
		ImageParts: cardImageParts(c),
	}
	// op.GeoM.Translate(ScreenWidth/2 - (float64(w) * 0 / 2) - (0 * 500), ScreenHeight/2 - (float64(h) * 0 / 2) - (0 * 500r))
	// op.GeoM.Translate(0,0)
	op.GeoM.Translate(float64(piles[p][0]), float64(piles[p][1]))
	s.parts.image.DrawImage(deckImage, op)

	if  p == 0 {
		msg := fmt.Sprintf(`c:  %d
		x: %d
		y: %d
		p: %d`, c, piles[p][0], piles[p][1], p)
		if err := ebitenutil.DebugPrint(s.parts.image, msg); err != nil {
			return err
		}
	}
	return nil
}


type cardImageParts int



func (b cardImageParts) Len() int {
	return c
}

func (b cardImageParts) Dst(i int) (x0, y0, x1, y1 int) {
	// x0, y0 = (-cardwidth * i)/2, (-cardheight * i)/2
	// x1, y1 = x0 + cardwidth, y0 + cardheight
	// return

	x0 = 0
	y0 = 0
	x1, y1 = x0 + cardwidth, y0 + cardheight
	return
}

func (b cardImageParts) Src(i int) (x0, y0, x1, y1 int) {
	// log.Printf("i: %d", i)
	// x0, y0 = deck[i].imgLoc[0], deck[i].imgLoc[1]

	// x0 = (cardwidth * (i-1)) - 450
	// y0 = 500 - (cardheight * (i/13))
	// x = 72 * (i%13 - 1)
	// y = math.Remainder() * 100

	r, _ := math.Modf(float32((i - 1) / 13))
	c := float32(i) - (13.0 * r)

	x0 = int(72 * (c-1))
	y0 = int(100 * r)
	x1, y1 = x0 + cardwidth, y0 + cardheight
	return
}