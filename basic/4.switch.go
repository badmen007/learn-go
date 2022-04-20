package main

import "fmt"

// switch 不用 break
func grade(score int) string {
	g := ""
	switch {
	case score < 0 || score > 100:
		panic(fmt.Sprintf(
			"Wrong score: %d", score))
	case score < 60:
		g = "F"
	case score < 80:
		g = "c"
	case score < 90:
		g = "b"
	case score <= 100:
		g = "A"
	}
	return g
}

func main() {
	fmt.Println(
		grade(70),
		grade(59),
		grade(60),
		grade(82),
		//grade(101),
	)
}
