package day19

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/1e9y/adventofcode/challenge"
	"github.com/1e9y/adventofcode/util"
)

func Signs(max int) <-chan []int {
	ch := make(chan []int)
	var generate func([]int)
	generate = func(set []int) {
		if len(set) == max {
			ch <- set
			return
		}
		generate(append(set, -1))
		generate(append(set, 1))
	}
	go func() {
		generate([]int{})
		close(ch)
	}()
	return ch
}

func Permutations() <-chan []int {
	ch := make(chan []int)
	go func() {
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				if i == j {
					continue
				}
				for k := 0; k < 3; k++ {
					if i == k || j == k {
						continue
					}
					ch <- []int{i, j, k}
				}
			}
		}
		close(ch)
	}()
	return ch
}

// var signs, permutations []int

// func init() {
// 	for s := range Signs(3) {
// 		sign
// 	}
// var signs = Signs(3)
// var permutations = Permutations()
// }

type Beacon struct {
	position  []int
	relatives map[*Beacon]string
	distances map[int]*Beacon
}

func (beacon *Beacon) String() string {
	x, y, z := beacon.position[0], beacon.position[1], beacon.position[2]
	return fmt.Sprintf("%d,%d,%d", x, y, z)
}

func (beacon Beacon) compare(other *Beacon) []*Beacon {
	result := []*Beacon{}
	for d := range beacon.distances {
		if otherBeacon := other.distances[d]; otherBeacon != nil {
			result = append(result, otherBeacon)
		}
	}
	return result
}

func distance(a, b *Beacon) int {
	x1, y1, z1 := a.position[0], a.position[1], a.position[2]
	x2, y2, z2 := b.position[0], b.position[1], b.position[2]
	return (x1-x2)*(x1-x2) + (y1-y2)*(y1-y2) + (z1-z2)*(z1-z2)
}

type ScannerPosition struct {
	x, y, z int
}

type Scanner struct {
	id        int
	beacons   []*Beacon
	distances map[int][]*Beacon
	position  ScannerPosition
	locked    bool
}

func (scanner Scanner) String() (result string) {
	result = fmt.Sprintf("--- scanner %d --- (%d,%d,%d)\n", scanner.id, scanner.position.x, scanner.position.y, scanner.position.z)
	for _, beacon := range scanner.beacons {
		result += fmt.Sprintln(beacon)
	}
	return
}

func newScannerFromInput(id int, input <-chan string) *Scanner {
	scanner := &Scanner{
		id:        id,
		beacons:   make([]*Beacon, 0),
		distances: make(map[int][]*Beacon),
		position:  ScannerPosition{},
		locked:    false,
	}

	for line := range input {
		if line == "" {
			break
			// return scanner
		}

		parts := strings.Split(line, ",")
		if len(parts) != 3 {
			panic(fmt.Errorf("bad input: bad beacon: %s", line))
		}

		beacon := &Beacon{
			position: []int{
				util.MustAtoi(parts[0]),
				util.MustAtoi(parts[1]),
				util.MustAtoi(parts[2]),
			},
			relatives: make(map[*Beacon]string),
			distances: make(map[int]*Beacon),
		}

		var d int
		for i := 0; i < len(scanner.beacons); i++ {
			// TODO: Check if two different pairs have same distance
			d = distance(beacon, scanner.beacons[i])
			scanner.distances[d] = []*Beacon{beacon, scanner.beacons[i]}
		}

		scanner.beacons = append(scanner.beacons, beacon)
	}

	var d int
	for _, beacon := range scanner.beacons {
		for _, otherBeacon := range scanner.beacons {
			if &beacon == &otherBeacon {
				continue
			}
			d = distance(beacon, otherBeacon)
			beacon.distances[d] = otherBeacon
			otherBeacon.distances[d] = beacon

			dx := util.AbsInt(beacon.position[0] - otherBeacon.position[0])
			dy := util.AbsInt(beacon.position[1] - otherBeacon.position[1])
			dz := util.AbsInt(beacon.position[2] - otherBeacon.position[2])
			fingerprint := fmt.Sprintf("%d,%d,%d", d,
				util.MinInt(dx, dy, dz),
				util.MaxInt(dx, dy, dz),
			)
			beacon.relatives[otherBeacon] = fingerprint
			otherBeacon.relatives[beacon] = fingerprint
		}
	}

	return scanner
}

func (scanner *Scanner) lock(position ScannerPosition) {
	scanner.position = position
	scanner.locked = true
}

func (scanner *Scanner) repostion(dx, dy, dz int, perm, sign []int) {
	// if scanner.id == 4 {
	// 	// fmt.Println(scanner)
	// 	fmt.Println(sign, perm)
	// }
	scanner.position = ScannerPosition{
		x:/* scanner.position.x */ -dx,
		y:/* scanner.position.y */ -dy,
		z:/* scanner.position.z */ -dz,
	}
	for _, beacon := range scanner.beacons {
		// if scanner.id == 4 {
		// 	fmt.Println(beacon, dx, dy, dz)
		// }
		newx := sign[0]*beacon.position[perm[0]] /* - dx */ + scanner.position.x
		newy := sign[1]*beacon.position[perm[1]] /* - dy */ + scanner.position.y
		newz := sign[2]*beacon.position[perm[2]] /* - dz */ + scanner.position.z
		beacon.position[0] = newx
		beacon.position[1] = newy
		beacon.position[2] = newz
		// if scanner.id == 4 {
		// 	fmt.Println(beacon)
		// }
	}
	scanner.locked = true
	// println("repositioned", scanner.id)
	// if scanner.id == 4 {
	// 	fmt.Println(scanner)
	// }
}

func (scanner *Scanner) compare(other *Scanner) map[*Beacon]*Beacon {
	index := map[*Beacon]*Beacon{}
	for _, thisBeacon := range scanner.beacons {
		for _, otherBeacon := range other.beacons {
			distances := thisBeacon.compare(otherBeacon)
			if len(distances) >= 12 {
				index[thisBeacon] = otherBeacon
			}
		}
	}
	return index
}

func (scanner *Scanner) align(other *Scanner, overlaps map[*Beacon]*Beacon) bool {
	// if scanner.id == 4 && other.id == 2 {
	// 	for k, v := range overlaps {
	// 		fmt.Println(k, " :", v)
	// 	}
	// }

	//
	// 408,-1815,803  : -258,-428,682
	//                 0      1    2
	//                  -428 682  258
	//                   1    2   0

	// 1105,-1205,1229
	// 432,-2009,1493  : 673,-379,-804
	//                 0      1    2
	//                  646 498  828
	//                   1    2   0

	for perm := range Permutations() {
		i, j, k := perm[0], perm[1], perm[2]
		for sign := range Signs(3) {
			sx, sy, sz := sign[0], sign[1], sign[2]

			// first pair
			var dx, dy, dz int
			var x1, y1, z1 int
			var x2, y2, z2 int
			var x1p, y1p, z1p int
			var x2p, y2p, z2p int
			for thatBeacon, otherBeacon := range overlaps {
				// x1, y1, z1 = (*thatBeacon).position[i], (*thatBeacon).position[j], (*thatBeacon).position[k]
				x1, y1, z1 = (*thatBeacon).position[0], (*thatBeacon).position[1], (*thatBeacon).position[2]
				x2, y2, z2 = (*otherBeacon).position[i], (*otherBeacon).position[j], (*otherBeacon).position[k]
				x2, y2, z2 = sx*x2, sy*y2, sz*z2

				dx = x2 - x1
				dy = y2 - y1
				dz = z2 - z1
				if scanner.id == 1 && other.id == 4 {
					// fmt.Println("TRYING", dx, dy, dz)
					// fmt.Println("TRYING", x2, y2, z2)
				}
				confirmations := 0
				for thatBeaconPrime, otherBeaconPrime := range overlaps {
					if thatBeaconPrime == thatBeacon {
						continue
					}
					// if dx == 0 && dy == 0 && dz == 0 {
					// fmt.Println(other)
					// fmt.Println("  ", x2, y2, z2)
					// dx = x2 - x1
					// dy = y2 - y1
					// dz = z2 - z1
					// if scanner.id == 1 && other.id == 4 {
					// fmt.Println(k, " :", v)
					// fmt.Println("TRYING", dx, dy, dz)
					// } else {

					x1p, y1p, z1p = (*thatBeaconPrime).position[0], (*thatBeaconPrime).position[1], (*thatBeaconPrime).position[2]
					x2p, y2p, z2p = (*otherBeaconPrime).position[i], (*otherBeaconPrime).position[j], (*otherBeaconPrime).position[k]
					x2p, y2p, z2p = sx*x2p, sy*y2p, sz*z2p
					if util.AbsInt(x1p) != util.AbsInt(x2p-dx) || util.AbsInt(y1p) != util.AbsInt(y2p-dy) || util.AbsInt(z1p) != util.AbsInt(z2p-dz) {
						continue
					} else if confirmations >= 3 {
						// fmt.Println("YAY", x2, y2, z2)
						// fmt.Println("FOUND", dx, dy, dz)
						other.repostion(dx, dy, dz, perm, sign)

						// fmt.Println(other)
						return true
					} else {
						confirmations++
					}
					// }
				}
			}
		}
	}
	return false
}

var scannerRe = regexp.MustCompile(`--- scanner (\d+) ---`)

func parseInput(input <-chan string) (scanners []*Scanner) {
	for line := range input {
		if match := scannerRe.FindStringSubmatch(line); len(match) != 0 {
			id := util.MustAtoi(match[1])
			scanner := newScannerFromInput(id, input)
			scanners = append(scanners, scanner)
		}
	}
	return
}

func beacons(scanners []*Scanner) (int, int) {
	locked := map[int]bool{
		0: true,
	}
	scanners[0].lock(ScannerPosition{0, 0, 0})
	// outer:
	for len(locked) < len(scanners) {
		for i := range scanners {
			for j := range scanners {
				if i == j || !locked[i] || locked[j] {
					continue
				}
				overlaps := scanners[i].compare(scanners[j])
				if len(overlaps) >= 12 {
					// fmt.Println(len(overlaps))
					// fmt.Printf("scanner %v (%v) overlaps with scanner %v (%v)\n", i, scanners[i].locked, j, scanners[j].locked)
					if scanners[i].align(scanners[j], overlaps) {
						locked[j] = true
					}
					// fmt.Println(scanners[j].position)
					// fmt.Println(scanners[j])
				}
			}
		}
	}

	// plainOverlaps := make([][][]Beacon, 0)
	// for _, v := range overlaps {
	// 	plainOverlaps = append(plainOverlaps, v)
	// }

	// findCoordinatesTranslation(plainOverlaps)
	// findCoordinatesTranslationFromPair(plainOverlaps[0])

	unique := map[string]bool{}
	// xmax, ymax, zmax := math.MinInt32, math.MinInt32, math.MinInt32
	// xmin, ymin, zmin := math.MaxInt32, math.MaxInt32, math.MaxInt32
	for _, scanner := range scanners {
		// xmax = util.MaxInt(xmax, scanner.position.x)
		// ymax = util.MaxInt(ymax, scanner.position.y)
		// zmax = util.MaxInt(zmax, scanner.position.z)

		// xmin = util.MinInt(xmin, scanner.position.x)
		// ymin = util.MinInt(ymin, scanner.position.y)
		// zmin = util.MinInt(zmin, scanner.position.z)

		for _, beacon := range scanner.beacons {
			unique[beacon.String()] = true
		}
	}

	var max int
	for i := range scanners {
		for j := i + 1; j < len(scanners); j++ {
			dist := util.AbsInt(scanners[i].position.x-scanners[j].position.x) +
				util.AbsInt(scanners[i].position.y-scanners[j].position.y) +
				util.AbsInt(scanners[i].position.z-scanners[j].position.z)
			max = util.MaxInt(max, dist)
		}
	}

	// fmt.Println(xmax, ymax, zmax, xmin, ymin, zmin)

	return len(unique), max
}

func A(input *challenge.Challenge) int {
	answer, _ := beacons(parseInput(input.Lines()))
	return answer
}

func B(input *challenge.Challenge) int {
	_, answer := beacons(parseInput(input.Lines()))
	return answer
	// return secondHomework(input.LineSlice())
}
