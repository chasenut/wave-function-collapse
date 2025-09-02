package wfc

import (
	//"github.com/chasenut/wave-function-collapse/src/config"
	//rl "github.com/gen2brain/raylib-go/raylib"
)

type Cell struct {
	X int16
	Y int16
	Options map[int8][]int8
	State int8
	Entropy float32
}

type Grid struct {
	X int32
	Y int32
	Grid [][]Cell

}

const (
)

var (
	
)
