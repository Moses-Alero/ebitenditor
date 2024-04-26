package entity

import (
	"ebitenditor/ui/component"
)

type Button struct {
	component.Transform
	component.Icon
	component.BackGround
	component.Clickable
	component.ActionHandler
	component.Text
}
