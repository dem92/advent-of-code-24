package utils

type Direction struct {
	Dir string
	X   int
	Y   int
}

var DirNorth = Direction{Dir: "N", X: 0, Y: -1}
var DirNorthEast = Direction{Dir: "NE", X: 1, Y: -1}
var DirEast = Direction{Dir: "E", X: 1, Y: 0}
var DirSouthEast = Direction{Dir: "SE", X: 1, Y: 1}
var DirSouth = Direction{Dir: "S", X: 0, Y: 1}
var DirSoutWest = Direction{Dir: "SW", X: -1, Y: 1}
var DirWest = Direction{Dir: "W", X: -1, Y: 0}
var DireNorthWest = Direction{Dir: "NW", X: -1, Y: -1}

var Directions = []Direction{
	DirNorth,
	DirNorthEast,
	DirEast,
	DirSouthEast,
	DirSouth,
	DirSoutWest,
	DirWest,
	DireNorthWest,
}
