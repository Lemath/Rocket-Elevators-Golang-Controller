package main

//Button on a floor or basement to go back to lobby
type CallButton struct {
	ID int
	floor int
	direction string
	status string
}

func NewCallButton(_ID int, _floor int, _direction string, _status string) *CallButton {
	callButton := CallButton {
		ID: _ID,
		floor: _floor,
		direction: _direction,
		status: _status,
	}
	return &callButton
}
