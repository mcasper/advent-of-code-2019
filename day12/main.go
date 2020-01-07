package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type Moon struct {
	x        int
	y        int
	z        int
	velocity Velocity
}

func (m *Moon) applyVelocity() {
	m.x = m.x + m.velocity.x
	m.y = m.y + m.velocity.y
	m.z = m.z + m.velocity.z
}

func (m Moon) Id() string {
	return fmt.Sprintf("%v|%v|%v|%v|%v|%v", m.x, m.y, m.z, m.velocity.x, m.velocity.y, m.velocity.z)
}

func (m Moon) Dimension(i int) int {
	if i == 0 {
		return m.x
	} else if i == 1 {
		return m.y
	} else {
		return m.z
	}
}

func moonFromString(str string) Moon {
	re := regexp.MustCompile(`=([\d\-_]+)`)
	matches := re.FindAllStringSubmatch(str, 3)

	x, err := strconv.Atoi(matches[0][1])
	if err != nil {
		log.Fatal(err)
	}

	y, err := strconv.Atoi(matches[1][1])
	if err != nil {
		log.Fatal(err)
	}

	z, err := strconv.Atoi(matches[2][1])
	if err != nil {
		log.Fatal(err)
	}

	return Moon{
		x: x,
		y: y,
		z: z,
		velocity: Velocity{
			x: 0,
			y: 0,
			z: 0,
		},
	}
}

type Velocity struct {
	x int
	y int
	z int
}

func main() {
	input, err := ioutil.ReadFile("part1.txt")
	split := strings.Split(strings.Trim(string(input), "\n"), "\n")

	if err != nil {
		log.Fatal(err)
	}

	var moons []Moon
	for _, line := range split {
		moons = append(moons, moonFromString(line))
	}

	pairs := moonsToPairs(moons)
	timeStep := 0

	start := time.Now().Unix()

	xPeriod := 0
	xHashes := make(map[string]int)
	xT := 0
	yPeriod := 0
	yHashes := make(map[string]int)
	yT := 0
	zPeriod := 0
	zHashes := make(map[string]int)
	zT := 0

	for {
		for _, pair := range pairs {
			applyGravity(&moons[pair[0]], &moons[pair[1]])
		}
		for i := range moons {
			moons[i].applyVelocity()
		}

		if xPeriod == 0 {
			if val, ok := xHashes[dimensionHash(moons, 0)]; ok {
				if timeStep-val == 1 {
					if xT > 0 {
						fmt.Println("Found x period!")
						xPeriod = timeStep - xT
					} else {
						xT = timeStep - xT
					}
				}
			} else {
				xHashes[dimensionHash(moons, 0)] = timeStep
			}
		}

		if yPeriod == 0 {
			if val, ok := yHashes[dimensionHash(moons, 1)]; ok {
				if timeStep-val == 1 {
					if yT > 0 {
						fmt.Println("Found y period!")
						yPeriod = timeStep - yT
					} else {
						yT = timeStep - yT
					}
				}
			} else {
				yHashes[dimensionHash(moons, 1)] = timeStep
			}
		}

		if zPeriod == 0 {
			if val, ok := zHashes[dimensionHash(moons, 2)]; ok {
				if timeStep-val == 1 {
					if zT > 0 {
						fmt.Println("Found z period!")
						zPeriod = timeStep - zT
					} else {
						zT = timeStep - zT
					}
				}
			} else {
				zHashes[dimensionHash(moons, 2)] = timeStep
			}
		}

		if xPeriod != 0 && yPeriod != 0 && zPeriod != 0 {
			fmt.Printf("%v, %v, %v\n", xPeriod, yPeriod, zPeriod)
			fmt.Printf("%v\n", LCM(xPeriod, yPeriod, zPeriod)*2)
			break
		}

		if timeStep%1_000_000 == 0 {
			fmt.Printf("[%vs] On time step %v\n", time.Now().Unix()-start, timeStep)
		}

		timeStep++
	}
}

func applyGravity(moon1, moon2 *Moon) {
	if moon1.x > moon2.x {
		moon1.velocity.x--
		moon2.velocity.x++
	} else if moon2.x > moon1.x {
		moon1.velocity.x++
		moon2.velocity.x--
	}

	if moon1.y > moon2.y {
		moon1.velocity.y--
		moon2.velocity.y++
	} else if moon2.y > moon1.y {
		moon1.velocity.y++
		moon2.velocity.y--
	}

	if moon1.z > moon2.z {
		moon1.velocity.z--
		moon2.velocity.z++
	} else if moon2.z > moon1.z {
		moon1.velocity.z++
		moon2.velocity.z--
	}
}

func moonsToPairs(moons []Moon) [][]int {
	var result [][]int
	for i1, moon1 := range moons {
		for i2, moon2 := range moons {
			if moon1 == moon2 {
				continue
			}
			if i2 < i1 {
				continue
			}
			result = append(result, []int{i1, i2})
		}
	}
	return result
}

func dimensionHash(moons []Moon, dim int) string {
	return fmt.Sprintf("%v|%v|%v|%v", moons[0].Dimension(dim), moons[1].Dimension(dim), moons[2].Dimension(dim), moons[3].Dimension(dim))
}

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}
