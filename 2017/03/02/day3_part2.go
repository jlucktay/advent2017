// Package main for Advent of Code 2017, day 3, part 2
// http://adventofcode.com/2017/day/3
package main

import (
	"fmt"
	"math"
)

type direction uint8

const (
	right direction = iota
	up    direction = iota
	left  direction = iota
	down  direction = iota
)

type spiralNodeCoords struct {
	x, y int
}

type spiralNode struct {
	coords spiralNodeCoords
	next   *spiralNode
	parent *spiral
	value  uint
}

/*
root:			(pointer to) the first node
last:			(pointer to) the last node
dir:			direction in which to add the next node
firstChange:	side lengths come in pairs, so are we on the first of the two with this side length, or the second?
edgeDistance:	number of nodes already traversed on the current edge
sideLength:		the maximum number of nodes to add, before we need to change direction
*/
type spiral struct {
	root, last               *spiralNode
	dir                      direction
	firstChange              bool
	edgeDistance, sideLength uint
	nodeMap                  map[spiralNodeCoords]*spiralNode
}

func (s *spiral) Init() {
	s.root = nil
	s.last = nil
	s.dir = right
	s.firstChange = true
	s.edgeDistance = 0
	s.sideLength = 1
	s.nodeMap = make(map[spiralNodeCoords]*spiralNode)
}

func (s *spiral) Add() {
	if s.root == nil {
		newCoords := spiralNodeCoords{0, 0}
		newNode := &spiralNode{
			newCoords,
			nil,
			s,
			1,
		}

		s.root = newNode
		s.last = s.root
		s.nodeMap[newCoords] = newNode

		return
	}

	newCoords := s.last.coords.next(s.dir)
	newNode := &spiralNode{
		newCoords,
		nil,
		s,
		0,
	}

	newNode.value = newNode.sumNeighbours()

	s.last.next = newNode
	s.last = s.last.next
	s.nodeMap[newCoords] = newNode

	s.edgeDistance++

	if s.edgeDistance >= s.sideLength {
		s.edgeDistance = 0
		s.dir = s.dir.turn()

		if s.firstChange {
			s.firstChange = false
		} else {
			s.firstChange = true
			s.sideLength++
		}
	}
}

// Manhattan distance will be the sum of the absolutes of the last node's coordinates.
func (s *spiral) Manhattan() uint {
	return uint(math.Abs(float64(s.last.coords.x)) + math.Abs(float64(s.last.coords.y)))
}

func (snc spiralNodeCoords) neighbours() [8]spiralNodeCoords {
	return [...]spiralNodeCoords{
		{snc.x, snc.y + 1},
		{snc.x + 1, snc.y + 1},
		{snc.x + 1, snc.y},
		{snc.x + 1, snc.y - 1},
		{snc.x, snc.y - 1},
		{snc.x - 1, snc.y - 1},
		{snc.x - 1, snc.y},
		{snc.x - 1, snc.y + 1},
	}
}

func (sn spiralNode) sumNeighbours() (result uint) {
	for _, neighbour := range sn.coords.neighbours() {
		if node, ok := sn.parent.nodeMap[neighbour]; ok {
			result += node.value
		}
	}

	return
}

func (sn spiralNode) String() string {
	return fmt.Sprintf("%s %v", sn.coords, sn.value)
}

func (snc spiralNodeCoords) String() string {
	return fmt.Sprintf("[%d,%d]", snc.x, snc.y)
}

func (d direction) turn() direction {
	switch d {
	case right:
		return up
	case up:
		return left
	case left:
		return down
	default: //down
		return right
	}
}

func (snc spiralNodeCoords) next(dir direction) spiralNodeCoords {
	switch dir {
	case right:
		return spiralNodeCoords{snc.x + 1, snc.y}
	case up:
		return spiralNodeCoords{snc.x, snc.y + 1}
	case left:
		return spiralNodeCoords{snc.x - 1, snc.y}
	default: //down
		return spiralNodeCoords{snc.x, snc.y - 1}
	}
}
