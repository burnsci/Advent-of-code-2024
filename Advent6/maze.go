package main

type Maze struct{
	maze [][]string
}

func(mz Maze) checkInBounds(i int, j int) bool{
	if (i<len(mz.maze) || i>=0){
		return true
	}
	if (j<len(mz.maze[i]) || j>=0) {
		return true
	}
	return false
}

func(mz Maze) checkEmpty(i int, j int) bool{
	if mz.maze[i][j] == "."{
		return true
	}
	return false
}

func(mz *Maze) updateVisit(guardPos guardPosition){
	mz.maze[guardPos.i][guardPos.j]="X"
}

func(mz Maze) countUniqueVisits() int{
	uniqueVisits :=0
	for i:=0; i<len(mz.maze); i++{
		for j:=0; j<len(mz.maze[i]); j++{
			if mz.maze[i][j] == "X"{
				uniqueVisits++
			}
		}
	}
	return uniqueVisits
}