package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/chasenut/wave-function-collapse/src/config"
)

const (
	SCREEN_WIDTH int32 = config.SCREEN_WIDTH
	SCREEN_HEIGHT int32 = config.SCREEN_HEIGHT
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
			cameraTarget.Y -= cameraSpeed * dt
		}
		if cameraDown {
			cameraTarget.Y += cameraSpeed * dt
		}
		if cameraRight {
			cameraTarget.X += cameraSpeed * dt
		}
		if cameraLeft {
			cameraTarget.X -= cameraSpeed * dt
		}
	}
	cameraMoving = false
	cameraUp, cameraDown, cameraRight, cameraLeft = false, false, false, false
	camera.Target = cameraTarget
}

func drawDebugPanel() {
	if !showDebugPanel {
		return
	}
	rl.DrawText(fmt.Sprintf("FPS: %.3f", float32(1.0 / dt)),
		int32(cameraTarget.X)-SCREEN_WIDTH/2+20, 
		int32(cameraTarget.Y)-SCREEN_HEIGHT/2+20, 
		40, rl.LightGray)
}

func drawScene() {
	rl.ClearBackground(bkgColor)
	rl.DrawRectangle(0, 0, 100, 100, rl.White) // testing camera

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
	rl.InitWindow(SCREEN_WIDTH, SCREEN_HEIGHT, "Wave Function Collapse")
	rl.SetTargetFPS(FPS)
	rl.SetExitKey(0)

	cameraTarget.X = float32(SCREEN_WIDTH/2)
	cameraTarget.Y = float32(SCREEN_HEIGHT/2)
	cameraOffset.X = 0
	cameraOffset.Y = 0
	camera = rl.NewCamera2D(cameraTarget, cameraOffset, 0, 1)
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
