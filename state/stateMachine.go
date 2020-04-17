package state

import ()

type StateMachine struct {
	currState int
}

func NewStateMachine() StateMachine {
	obj := StateMachine{}
	obj.currState = mainMenu

	return obj
}

func (this *StateMachine) IsMainMenu() bool {
	return this.currState == mainMenu
}

func (this *StateMachine) IsGamePlay() bool {
	return this.currState == gamePlay
}

func (this *StateMachine) IsGameOVer() bool {
	return this.currState == gameOver
}

func (this *StateMachine) UpdateStateGameplay() {
	this.currState = gamePlay
}

func (this *StateMachine) UpdateStateGameOver() {
	this.currState = gameOver
}
