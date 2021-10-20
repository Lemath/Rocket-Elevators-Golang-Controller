package main

type Door struct {
	ID          int
	status      string
	sensorState string
}

func NewDoor() *Door {
	door := Door{
		ID: 1,
		status: "closed",
		sensorState: "OFF",
	}
	return &door
}

func (d *Door) isObstructed() bool {
	return d.sensorState == "ON"
}