package main

import (
	"github.com/gen2brain/raylib-go/raylib"
)

var (
	sw    int32 = 400
	sh    int32 = 800
	fps   int32 = 60
	speed int32 = 5
)

func main() {
	rl.InitWindow(sw, sh, "Flappy Gopher")
	defer rl.CloseWindow()

	rl.SetTargetFPS(fps)

	bg := rl.LoadImage("assets/background.png")
	rl.ImageResize(bg, bg.Width, sh)

	pipeUp := rl.LoadImage("assets/pipe-up.png")
	pipeDown := rl.LoadImage("assets/pipe-down.png")

	birdUp := rl.LoadImage("assets/bird-up.png")
	birdDown := rl.LoadImage("assets/bird-down.png")
	birdMid := rl.LoadImage("assets/bird-mid.png")

	// We will load the bird-mid image as the default texture
	texture := rl.LoadTextureFromImage(birdMid)

	xCoords := sw/2 - texture.Width/2
	yCoords := sh/2 - texture.Height/2

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		// Draw the background
		rl.DrawTexture(rl.LoadTextureFromImage(bg), 0, 0, rl.RayWhite)

		// Draw the bird
		rl.DrawTexture(texture, xCoords, yCoords, rl.RayWhite)

		if rl.IsKeyDown(rl.KeySpace) {
			texture = rl.LoadTextureFromImage(birdUp)
			//yCoords -= speed
		} else {
			texture = rl.LoadTextureFromImage(birdDown)
			//yCoords += speed
		}

		// Draw the pipes
		rl.DrawTexture(rl.LoadTextureFromImage(pipeDown), 100, 0, rl.RayWhite)
		rl.DrawTexture(rl.LoadTextureFromImage(pipeUp), 120, sh-pipeUp.Height, rl.RayWhite)

		rl.EndDrawing()
	}
}
