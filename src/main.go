package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	screenWidth int32 = 800
	screenHeight int32 = 800
	FPS int32 = 60
)


var (
	running bool = true
	dt float32 = rl.GetFrameTime()

	showDebugPanel bool = true
	
	bkgColor rl.Color = rl.NewColor(22, 22, 35, 255)

	cameraTarget rl.Vector2
	cameraOffset rl.Vector2
	camera rl.Camera2D
	cameraSpeed float32 = 250
	cameraMoving bool = false
	cameraUp, cameraDown, cameraRight, cameraLeft = false, false, false, false
)

func input() {
	if rl.IsKeyDown(rl.KeyW) || rl.IsKeyDown(rl.KeyUp) {
		cameraMoving = true
		cameraUp = true
	}
	if rl.IsKeyDown(rl.KeyS) || rl.IsKeyDown(rl.KeyDown) {
		cameraMoving = true
		cameraDown = true
	}
	if rl.IsKeyDown(rl.KeyD) || rl.IsKeyDown(rl.KeyRight) {
		cameraMoving = true
		cameraRight = true
	}
	if rl.IsKeyDown(rl.KeyA) || rl.IsKeyDown(rl.KeyLeft) {
		cameraMoving = true
		cameraLeft = true
	}
	if rl.IsKeyDown(rl.KeyQ) {

	}
}

func update() {
	running = !rl.WindowShouldClose()
	dt = rl.GetFrameTime()

	if cameraMoving {
		if cameraUp {
			cameraOffset.Y -= cameraSpeed * dt
		}
		if cameraDown {
			cameraOffset.Y += cameraSpeed * dt
		}
		if cameraRight {
			cameraOffset.X += cameraSpeed * dt
		}
		if cameraLeft {
			cameraOffset.X -= cameraSpeed * dt
		}
	}
	cameraMoving = false
	cameraUp, cameraDown, cameraRight, cameraLeft = false, false, false, false
	camera.Target = cameraOffset
}

func drawDebugPanel() {
	if !showDebugPanel {
		return
	}
	rl.DrawText(fmt.Sprintf("FPS: %.3f", float32(1.0 / dt)),
		int32(cameraOffset.X)-screenWidth/2+20, 
		int32(cameraOffset.Y)-screenHeight/2+20, 
		40, rl.LightGray)
}

func drawScene() {
	rl.ClearBackground(bkgColor)

	drawDebugPanel() // bool: showDebugPanel
}

func render() {
	rl.BeginDrawing()
	rl.BeginMode2D(camera)

	drawScene()

	rl.EndMode2D()
	rl.EndDrawing()
}

func initialize() {
	rl.InitWindow(screenWidth, screenHeight, "Wave Function Collapse")
	rl.SetTargetFPS(FPS)
	rl.SetExitKey(0)

}

func exit() {
	defer rl.CloseWindow()

}

func main() {
	initialize()


	for running {
		input()
		update()
		render()
	}
	exit()
}
