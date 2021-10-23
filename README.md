# Rocket-Elevators-Golang-Controller
### Description

This project simulate the selection and movement logic for elevators in a commercial building and is written in Go.

When someone want to call an elevator from the lobby, he will first select his destination. The controller will then decide which column of the battery to use and select the best elevator available to fill the request and carry the user to his selected floor. When someone call an elevator from a floor, the corresponding column will select his best available elevator to pick up the user and bring him back to the lobby.

### Installation

With golang installed on your computer, all you need to do is initialize the module:

`go mod init Rocket-Elevators-Commercial-Controller`

The code to run the scenarios is included, and can be executed with:

`go run . <SCENARIO-NUMBER>`

### Running the tests

To launch the tests:

`go test`

With a fully completed project, you should get an output like:

![golang_test1](https://user-images.githubusercontent.com/56204810/138209403-3d5e289d-1296-4d8a-b607-817bab6995d1.jpg)

You can also get more details about each test by adding the `-v` flag: 

`go test -v` 

which should give something like: 

![golang_test2](https://user-images.githubusercontent.com/56204810/138209419-b43cb11a-b297-4c55-a704-57088e127b2f.jpg)



