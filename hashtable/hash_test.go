package hashtable

func robotSim(commands []int, obstacles [][]int) int {
	if len(commands) == 0 {
		return 0
	}

	obs := make(map[Point]struct{})

	for _, obstacle := range obstacles {
		if len(obstacle) == 2 {
			point := Point{obstacle[0], obstacle[1]}
			obs[point] = struct{}{}
		}
	}

	directionX := []int{0,1, 0, -1}
	directionY := []int{1,0, -1, 0}

	dir := 0
	x, y := 0, 0
	ret := 0

	for _, command := range commands {
		if command == -1 {
			dir = (dir+1)%4
			continue
		}

		if command == -2 {
			dir = (dir+3)%4
			continue
		}

		for i := 0; i < command;i++ {
			curX := x + directionX[dir]
			curY := y + directionY[dir]

			point := Point{
				X: curX,
				Y: curY,
			}

			if _, ok := obs[point]; ok {
				break
			}

			curMaxDir := curX*curX +curY*curY
			if curMaxDir > ret {
				ret = curMaxDir
			}

			x = curX
			y = curY
		}
	}

	return ret
}

type Point struct {
	X int
	Y int
}
