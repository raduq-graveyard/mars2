package handler

import "fmt"

// Walk the robot
func Walker(pl planet, rb robot, c []string) (string, error) {
	each(c, func(cm string) {
		switch cm {
		case "W":
			walk(&rb)
		case "R":
			rotate("R", &rb)
		case "L":
			rotate("L", &rb)
		}
	})

	if rb.AtX < 0 || pl.MaxX < rb.AtX {
		return "", fmt.Errorf("Trespassed X")
	}
	if rb.AtY < 0 || pl.MaxX < rb.AtY {
		return "", fmt.Errorf("Trespassed Y")
	}
	pos := fmt.Sprintf("%d%d%s", rb.AtX, rb.AtY, rb.Ori)
	return pos, nil
}

type frit func(string)

func each(arr []string, fn frit) {
	for _, cm := range arr {
		fn(cm)
	}
}

func walk(rb *robot) {
	switch rb.Ori {
	case "N":
		goNorth(rb)
	case "S":
		goSouth(rb)
	case "E":
		goEast(rb)
	case "W":
		goWest(rb)
	}
}

func rotate(d string, rb *robot) {
	switch rb.Ori {
	case "N":
		rb.Ori = _rotate(d, "E", "W")
	case "S":
		rb.Ori = _rotate(d, "W", "E")
	case "E":
		rb.Ori = _rotate(d, "S", "N")
	case "W":
		rb.Ori = _rotate(d, "N", "N")
	}
}

func goNorth(rb *robot) { rb.AtY++ }
func goSouth(rb *robot) { rb.AtY-- }
func goEast(rb *robot)  { rb.AtX++ }
func goWest(rb *robot)  { rb.AtX-- }
func _rotate(d string, dr string, dl string) string {
	if d == "R" {
		return dr
	}
	return dl
}
