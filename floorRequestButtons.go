package main

//FloorRequestButton is a button on the pannel at the lobby to request any floor
type FloorRequestButton struct {
	ID int
	floor int
	direction string
	status string
}

func NewFloorRequestButton(_ID int, _floor int, _direction string, _status string) *FloorRequestButton {
	floorRequestButton := FloorRequestButton{
		ID: _ID,
		floor: _floor,
		direction: _direction,
		status: _status,
	}
	return &floorRequestButton
}
