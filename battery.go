package main

import "math"

type Battery struct {
	ID int 
	status string
	columnsList []*Column
	floorRequestsButtonsList []*FloorRequestButton
	columnID byte
	floorRequestButtonID int
}
func NewBattery(_id, _amountOfColumns, _amountOfFloors, _amountOfBasements, _amountOfElevatorPerColumn int) *Battery {
	battery := Battery{
		ID: _id,
		status: "online",
		columnsList: []*Column{},
		floorRequestsButtonsList: []*FloorRequestButton{},
		columnID: 'A',
		floorRequestButtonID: 1,
	}
	if _amountOfBasements > 0 {
		battery.createBasementFloorRequestButtons(_amountOfBasements)
		battery.createBasementColumn(_amountOfBasements, _amountOfElevatorPerColumn)
		_amountOfColumns--
	}
	battery.createFloorRequestButtons(_amountOfFloors)
	battery.createColumns(_amountOfColumns, _amountOfFloors, _amountOfElevatorPerColumn)
	return &battery
}

func (b *Battery) createBasementColumn(_amountOfBasements int, _amountOfElevatorPerColumn int) {
	servedFloors := []int{}
	floor := -1
	for floor >= _amountOfBasements {
		servedFloors = append(servedFloors, floor)
		floor--
	}
	column := NewColumn(string(b.columnID), _amountOfElevatorPerColumn, servedFloors, true, "online", _amountOfBasements)
	b.columnsList = append(b.columnsList, column)
	b.columnID++
}

func (b *Battery) createColumns(_amountOfColumns int, _amountOfFloors int, _amountOfElevator int) {
	amountOfFloorsPerColumn := int(math.Ceil(float64(_amountOfFloors) / float64(_amountOfColumns)))
	floor := 1
	for c := 1; c <= _amountOfColumns; c++ {
		servedFloors := []int{}
		for i := 0; i < amountOfFloorsPerColumn; i++ {
			if floor <= _amountOfFloors {
				servedFloors = append(servedFloors, floor)
				floor++
			}
		}
		column := NewColumn(string(b.columnID), _amountOfElevator, servedFloors, false, "online", _amountOfFloors)
		b.columnsList = append(b.columnsList, column)
		b.columnID++
	} 
}

func (b *Battery) createFloorRequestButtons(_amountOfFloors int) {
	for buttonFloor := 1; buttonFloor <= _amountOfFloors; buttonFloor++ {
		FloorRequestButton := NewFloorRequestButton(b.floorRequestButtonID, buttonFloor, "up", "OFF") 
		b.floorRequestsButtonsList = append(b.floorRequestsButtonsList, FloorRequestButton)
		b.floorRequestButtonID++
	}
}

func (b *Battery) createBasementFloorRequestButtons(_amountOfBasements int) {
	for buttonFloor := -1; buttonFloor >= -_amountOfBasements; buttonFloor-- {
		floorRequestButton := NewFloorRequestButton(b.floorRequestButtonID, buttonFloor, "down", "OFF")
		b.floorRequestsButtonsList = append(b.floorRequestsButtonsList, floorRequestButton)
		b.floorRequestButtonID++
	}
}


func (b *Battery) findBestColumn(_requestedFloor int) *Column {
	var bestColumn *Column
	for _, column := range b.columnsList {
		if Contains(column.servedFloorsList, _requestedFloor) {
			bestColumn = column
		}
	}
	return bestColumn
}

//Simulate when a user press a button at the lobby
func (b *Battery) assignElevator(_requestedFloor int, _direction string) (*Column, *Elevator) {
	column := b.findBestColumn(_requestedFloor)
	elevator := column.findElevator(1, _direction)
	elevator.addNewRequest(1)
	elevator.move()
	elevator.addNewRequest(_requestedFloor)
	elevator.move()
	return column, elevator
}
