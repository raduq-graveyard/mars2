package main

import "fmt"

// Walk the robot
func Walk(pl planet, rb robot, c []string) (string, error) {
	for _, cm := range c {
		if cm == "N" {
			rb.AtY++
		}
		if cm == "S" {
			rb.AtY--
		}
		if cm == "E" {
			rb.AtX++
		}
		if cm == "W" {
			rb.AtX--
		}
	}
	if rb.AtX < 0 || pl.MaxX < rb.AtX {
		return "", fmt.Errorf("Trespassed X")
	}
	if rb.AtY < 0 || pl.MaxX < rb.AtY {
		return "", fmt.Errorf("Trespassed Y")
	}
	pos := fmt.Sprintf("%d%d%s", rb.AtX, rb.AtY, rb.Ori)
	return pos, nil
}
