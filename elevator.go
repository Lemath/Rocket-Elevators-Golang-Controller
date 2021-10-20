package main

import "sort"

type Elevator struct {
	ID string
	status string
	currentFloor int
	door *Door
	floorRequestsList []int
	completedRequestsList []int
	screenDisplay int
	direction string
	overweight bool
	overweightAlarm bool
}

func NewElevator(_elevatorID string, _status string, _currentFloor int) *Elevator {
	elevator := Elevator{
		ID: _elevatorID,
		status: _status,
		currentFloor: _currentFloor,
		door: NewDoor(),
		floorRequestsList: []int{},
		completedRequestsList: []int{},
		direction: "any",
		overweight: false,
		overweightAlarm: false,
		screenDisplay: _currentFloor,
	}
	return &elevator

}

func (e *Elevator) move() {
	for len(e.floorRequestsList) > 0 {
		var destination int = e.floorRequestsList[0]
		e.status = "moving"
		if e.currentFloor < destination {
			e.direction = "up"
			e.sortFloorList()
			for e.currentFloor < destination {
				e.currentFloor++
				e.screenDisplay = e.currentFloor
			}
		} else if e.currentFloor > destination {
			e.direction = "down"
			e.sortFloorList()
			for e.currentFloor > destination {
				e.currentFloor--
				e.screenDisplay = e.currentFloor
			}
		}
		e.status = "stopped"
		e.operateDoors()
		if !Contains(e.completedRequestsList, destination) {
			e.completedRequestsList = append(e.completedRequestsList, destination)
		}
		e.floorRequestsList = e.floorRequestsList[1:]
	}
	e.status = "idle"
}

func (e *Elevator) sortFloorList() {
	if len(e.floorRequestsList) > 1 {
		if e.direction == "up" {
			sort.Sort(sort.IntSlice(e.floorRequestsList))
		} else if e.direction == "down" {
			sort.Sort(sort.Reverse(sort.IntSlice(e.floorRequestsList)))
		}
	}
}

func (e *Elevator) operateDoors() {
	e.door.status = "opening"
	// wait 5 second
	if !e.isOverweight() {
		e.door.status = "closing"
		if !e.door.isObstructed() {
			e.door.status = "closed"
		} else {
			e.operateDoors()
		}
	} else {
		for e.isOverweight() {
			e.overweightAlarm = true
		}
		e.overweightAlarm = false
		e.operateDoors()
	}
}

func (e *Elevator) addNewRequest(_requestedFloor int) {
	if !Contains(e.floorRequestsList,_requestedFloor) {
		e.floorRequestsList = append([]int{_requestedFloor}, e.floorRequestsList...)
	}
	if e.currentFloor < _requestedFloor {
		e.direction = "up"
	}
	if e.currentFloor > _requestedFloor {
		e.direction = "down"
	}
}

func (e *Elevator) isOverweight() bool {
	return e.overweight
}