package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"regexp"
	"strconv"
	"strings"
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

func (m Moon) totalEnergy() int {
	return (m.Abs(m.x) + m.Abs(m.y) + m.Abs(m.z)) * (m.Abs(m.velocity.x) + m.Abs(m.velocity.y) + m.Abs(m.velocity.z))
}

func (m Moon) Abs(x int) int {
	return int(math.Abs(float64(x)))
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

	for {
		for _, pair := range pairs {
			applyGravity(&moons[pair[0]], &moons[pair[1]])
		}
		for i, _ := range moons {
			moons[i].applyVelocity()
		}
		timeStep++

		if timeStep == 1000 {
			break
		}
	}

	totalEnergy := calculateTotalEnergy(moons)

	fmt.Printf("Part 1 result: %v\n", totalEnergy)
}

func calculateTotalEnergy(moons []Moon) int {
	result := 0
	for _, moon := range moons {
		result += moon.totalEnergy()
	}
	return result
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
