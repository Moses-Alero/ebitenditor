package event

import (
	"log"
	"reflect"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type EventMapper map[string]func(arg interface{}) bool

var mapper EventMapper = EventMapper{
	"onClick": OnClick,
}

type CursorState int

const (
	Idle CursorState = iota
	MouseBtnPressed
	CusrsorOverlapping
	Clicked
)
//Events as functions


var cursorState = Idle

func OnClick(entity interface{}) bool{
	//logic for handling click event
	//a click happens when 
	//1. the mouse key is pressed down
	//2. the cursor is overlapping the object.
	//3. the mouse key has been released
	mouseBtnPressed := inpututil.IsMouseButtonJustPressed(ebiten.MouseButton0)
	switch cursorState{
		case Idle:
			if mouseBtnPressed {
				cursorState = MouseBtnPressed
			}
		case MouseBtnPressed:
			if isCursorOverlapping(entity){
				cursorState = CusrsorOverlapping	
			}else {
				cursorState = Idle
				return false
			}
		case CusrsorOverlapping:
			if  inpututil.IsMouseButtonJustReleased(ebiten.MouseButton0) {
					cursorState = Clicked
			}
		case Clicked:
			cursorState = Idle
			return true
	}

	return false
}


func isCursorOverlapping(entity interface{})  bool{
	cursorPosX, cursorPosY := ebiten.CursorPosition() 

	entityVal := reflect.ValueOf(entity).Elem()
	
	pos := entityVal.FieldByName("Position")

	if !pos.IsValid() {
		log.Fatal("Position not found in type", entityVal)
	}
	X := pos.FieldByName("X").Interface().(float64)
	Y := pos.FieldByName("Y").Interface().(float64)


	bounds := entityVal.FieldByName("Bounds")
	if !bounds.IsValid() {
		log.Fatal("Bounds not found in type", entityVal)
	}

	W := bounds.FieldByName("W").Interface().(float64)
	H := bounds.FieldByName("H").Interface().(float64)


	entityRight := X + W
	entityBottom := Y + H

	return X < float64(cursorPosX) && entityRight > float64(cursorPosX) &&
	entityBottom > float64(cursorPosY) && Y < float64(cursorPosY)
}


