package main

import (
    "log"

    "github.com/hajimehoshi/ebiten/v2"
    "github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
    screenWidth  = 256
    screenHeight = 176
)

type Game struct {
}

func (g *Game) Update() error {

    return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

    ebitenutil.DebugPrint(screen, "Home")
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
    return screenWidth, screenHeight
}

func main() {

    g := &Game{}

    ebiten.SetWindowSize(screenWidth * 4, screenHeight * 4)
    ebiten.SetWindowTitle("Zelda and Link - Adventure")

    if err := ebiten.RunGame(g); err != nil {
        log.Fatal(err)
    }
}
