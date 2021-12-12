package day11

import (
	"github.com/1e9y/adventofcode/challenge"
	"github.com/1e9y/adventofcode/util"
)

const MaximumEnergy = 9

type Octopus struct {
	energy     int
	discharged bool
	neighbours map[*Octopus]*Octopus
}

func (octopus *Octopus) bind(other *Octopus) {
	if octopus == other {
		return
	}

	octopus.neighbours[other] = other
	other.neighbours[octopus] = octopus
}

func (octopus *Octopus) increase() {
	octopus.energy += 1
}

func (octopus *Octopus) reset() {
	octopus.energy = 0
	octopus.discharged = false
}

func (octopus *Octopus) discharge() (affected []*Octopus) {
	if octopus.discharged {
		return
	}

	octopus.discharged = true
	for _, octopus := range octopus.neighbours {
		if !octopus.discharged {
			octopus.increase()
			if octopus.energy > MaximumEnergy {
				affected = append(affected, octopus)
			}
		}
	}
	return
}

type Swarm struct {
	octopuses     []*Octopus
	width, height int
}

func newSwarm(width, height int) (swarm *Swarm) {
	swarm = &Swarm{
		octopuses: make([]*Octopus, width*height),
		width:     width,
		height:    height,
	}

	for i := range swarm.octopuses {
		swarm.octopuses[i] = &Octopus{
			neighbours: make(map[*Octopus]*Octopus),
		}
	}

	for i := range swarm.octopuses {
		x := i % width
		y := i / height

		octopus := swarm.octopuses[i]
		for dx := -1; dx <= 1; dx++ {
			for dy := -1; dy <= 1; dy++ {
				sx := x + dx
				sy := y + dy
				if sx < 0 || sy < 0 || sx >= width || sy >= height {
					continue
				}
				other := swarm.octopuses[sy*width+sx]
				octopus.bind(other)
			}
		}
	}

	return
}

func (swarm *Swarm) energize(input []string) {
	width := len(input[0])
	for j, row := range input {
		for i, column := range row {
			swarm.octopuses[j*width+i].energy = util.MustAtoi(string(column))
		}
	}
}

func (swarm *Swarm) synchonized() bool {
	energy := swarm.octopuses[0].energy
	for _, octopus := range swarm.octopuses {
		if octopus.energy != energy {
			return false
		}
	}
	return true
}

func (swarm *Swarm) flash() (total int) {
	discharging := make([]*Octopus, 0)
	for _, octopus := range swarm.octopuses {
		octopus.increase()
		if octopus.energy > MaximumEnergy {
			discharging = append(discharging, octopus)
		}
	}

	for {
		if len(discharging) == 0 {
			break
		}

		octopus := discharging[0]
		discharging = discharging[1:]
		flashed := octopus.discharge()
		discharging = append(discharging, flashed...)
	}

	for _, octopus := range swarm.octopuses {
		if octopus.discharged {
			octopus.reset()
			total++
		}
	}

	return
}

func flashes(input []string, steps int) (result int) {
	swarm := newSwarm(len(input[0]), len(input))
	swarm.energize(input)

	for s := 0; s < steps; s++ {
		result += swarm.flash()
	}

	return
}

func synchonization(input []string) (result int) {
	swarm := newSwarm(len(input[0]), len(input))
	swarm.energize(input)

	for !swarm.synchonized() {
		_ = swarm.flash()
		result++
	}

	return
}

func A(input *challenge.Challenge) int {
	return flashes(input.LineSlice(), 100)
}

func B(input *challenge.Challenge) int {
	return synchonization(input.LineSlice())
}
