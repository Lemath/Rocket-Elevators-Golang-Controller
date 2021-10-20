package main

import (
	"math";
	"strconv"
)

type Column struct {
	ID string
	servedFloorsList []int
	isBasement bool
	status string
	elevatorsList []*Elevator
	callButtonsList []*CallButton
}

func NewColumn(_id string, _amountOfElevators int, _servedFloors []int, _isBasement bool, _status string, _amountOfFloors int) *Column {
	 var column Column = Column{
		ID: _id,
		servedFloorsList: _servedFloors,
		isBasement: _isBasement,
		status: _status,
		elevatorsList: []*Elevator{},
		callButtonsList: []*CallButton{},
	}
	column.createElevators(_amountOfFloors, _amountOfElevators)
	column.createCallButtons(_amountOfFloors, _isBasement)
	return &column
}

func (c *Column) createElevators(_amountOfFloors int, _amountOfElevators int) {
	for elevatorID := 1; elevatorID <=_amountOfElevators; elevatorID++ {
		elevator := NewElevator(c.ID + strconv.Itoa(elevatorID), "idle", 1)
		c.elevatorsList = append(c.elevatorsList, elevator)
	}
}

func (c *Column) createCallButtons(_amountOfFloors int, _isBasement bool) {
	var callButtonID int = 1
	if _isBasement {
		for buttonFloor := -1; buttonFloor >= -_amountOfFloors; buttonFloor-- {
			callButton := NewCallButton(callButtonID, buttonFloor, "up", "OFF")
			c.callButtonsList = append(c.callButtonsList, callButton)
			callButtonID++
		}
	} else {
		for buttonFloor := 1; buttonFloor <= _amountOfFloors; buttonFloor++ {
			callButton := NewCallButton(callButtonID, buttonFloor, "down", "OFF")
			c.callButtonsList = append(c.callButtonsList, callButton)
			callButtonID++
		}
	}
}

//Simulate when a user press a button on a floor to go back to the first floor
func (c *Column) requestElevator(_requestedFloor int, _direction string) *Elevator {
	elevator := c.findElevator(_requestedFloor, _direction)
	elevator.addNewRequest(_requestedFloor)
	elevator.move()
	elevator.addNewRequest(1)
	elevator.move()
	return elevator
}

func (c *Column) findElevator(_requestedFloor int, _requestedDirection string) *Elevator {
	var bestElevator *Elevator = nil
	var bestScore int = 6
	var referenceGap int = 100
	if  _requestedFloor == 1 {
		for _, elevator := range c.elevatorsList {
			var score int
			if elevator.currentFloor == 1 && elevator.status == "stopped" {
				score = 1
			} else if elevator.currentFloor == 1 && elevator.status == "idle" {
				score = 2 
			} else if 1 > elevator.currentFloor && elevator.direction == "up" {
				score = 3
			} else if 1 < elevator.currentFloor && elevator.direction == "down" {
				score = 3 
			} else if elevator.status == "idle" {
				score = 4 
			} else {
				score = 5 
			}
			bestElevator, bestScore, referenceGap = c.compareElevator(score, elevator, bestScore, referenceGap, bestElevator, _requestedFloor)
		}
	} else {
		for _, elevator := range c.elevatorsList {
			var score int
			if elevator.currentFloor == _requestedFloor && elevator.status == "stopped" && elevator.direction == _requestedDirection {
				score = 1 
			} else if _requestedFloor > elevator.currentFloor && elevator.direction == "up" && _requestedDirection == "up" {
				score = 2 
			} else if _requestedFloor < elevator.currentFloor && elevator.direction == "down" && _requestedDirection == "down" {
				score = 2 
			} else if elevator.status == "idle" {
				score = 4 
			} else {
				score = 5
			}
			bestElevator, bestScore, referenceGap = c.compareElevator(score, elevator, bestScore, referenceGap, bestElevator, _requestedFloor)
		}
	}
	return bestElevator
}

func (c *Column) compareElevator(score int, currentElevator *Elevator, bestScore int, referenceGap int, bestElevator *Elevator, requestedFloor int) (*Elevator, int, int) {
	if score < bestScore {
		bestScore = score
		bestElevator = currentElevator
		referenceGap = int(math.Abs(float64(currentElevator.currentFloor - requestedFloor)))
	} else if bestScore == score {
		var gap int = int(math.Abs(float64(currentElevator.currentFloor - requestedFloor)))
		if referenceGap > gap {
			bestElevator = currentElevator
			referenceGap = gap
		}
	}
	return bestElevator, bestScore, referenceGap
}