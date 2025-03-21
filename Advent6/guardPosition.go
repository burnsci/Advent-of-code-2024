package main

import(
	"slices"
	"errors"
)

type guardPosition struct{
	i int
	j int
	direction string
	patrolComplete bool
}

//func to return the correct facings for direction
func (pos guardPosition) getCorrectFacings() []string{
	return []string{"^", ">", "v", "<"}
}

//func to ensure only accepted directions are allowed
func (pos guardPosition) validateDirMarker(dirMarker string){
	correctFacings := pos.getCorrectFacings()
	if !slices.Contains(correctFacings, dirMarker){
		panic(errors.New("Invalid direction"))
	}
}
//constructor
func newGuardPosition(i int, j int, dirMarker string) *guardPosition{
	guardPos := guardPosition{i: i,j: j}
	guardPos.validateDirMarker(dirMarker)
	guardPos.direction = dirMarker
	guardPos.patrolComplete = false
	return &guardPos
}

//based on value of direction choose specific move method
func (pos *guardPosition) move(mz Maze){
	allowedFacings := pos.getCorrectFacings()
	if pos.direction == allowedFacings[0]{
		pos.moveFacingUp(mz)
	}else if pos.direction == allowedFacings[1]{
		pos.moveFacingRight(mz)
	}else if pos.direction == allowedFacings[2]{
		pos.moveFacingDown(mz)
	}else if pos.direction == allowedFacings[3]{
		pos.moveFacingLeft(mz)
	}
}

//check if the square in front is empty and not out of bounds
//if yes move to it
//else turn
func (pos *guardPosition) moveFacingUp(mz Maze){
	//-i,j
	newI := pos.i-1
	newJ := pos.j
	if mz.checkInBounds(newI, newJ){
		if mz.checkEmpty(newI, newJ){
			pos.i = newI
			pos.j = newJ
		}else{
			pos.direction = pos.getCorrectFacings()[1]
		}
	}else{
		pos.patrolComplete=true
	}
}

func (pos *guardPosition) moveFacingDown(mz Maze){
	//+i,j
	newI := pos.i+1
	newJ := pos.j
	if mz.checkInBounds(newI, newJ){
		if mz.checkEmpty(newI, newJ){
			pos.i = newI
			pos.j = newJ
		}else{
			pos.direction = pos.getCorrectFacings()[3]
		}
	}else{
		pos.patrolComplete=true
	}
}

func (pos *guardPosition) moveFacingLeft(mz Maze){
	//i,-j
	newI := pos.i
	newJ := pos.j-1
	if mz.checkInBounds(newI, newJ){
		if mz.checkEmpty(newI, newJ){
			pos.i = newI
			pos.j = newJ
		}else{
			pos.direction = pos.getCorrectFacings()[0]
		}
	}else{
		pos.patrolComplete=true
	}
}

func (pos *guardPosition) moveFacingRight(mz Maze){
	//i,+j
	newI := pos.i
	newJ := pos.j+1
	if mz.checkInBounds(newI, newJ){
		if mz.checkEmpty(newI, newJ){
			pos.i = newI
			pos.j = newJ
		}else{
			pos.direction = pos.getCorrectFacings()[2]
		}
	}else{
		pos.patrolComplete=true
	}
}
