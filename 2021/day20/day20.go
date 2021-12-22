package day20

import (
	"github.com/1e9y/adventofcode/challenge"
)

type PixelStatus int

const (
	DarkPixel     PixelStatus = 0
	LightPixel    PixelStatus = 1
	PaintingPixel PixelStatus = 2
)

type Image struct {
	grid                   map[int]map[int]PixelStatus
	xmin, xmax, ymin, ymax int
	infinite               PixelStatus
}

func (image Image) String() (result string) {
	offset := 5
	for j := image.ymin - offset; j <= image.ymax+offset; j++ {
		for i := image.xmin - offset; i <= image.xmax+offset; i++ {
			if j < image.ymin || j > image.ymax || i < image.xmin || i > image.xmax {
				if image.infinite&LightPixel > 0 {
					result += "#"
				} else {
					result += " "
				}
				continue
			}
			if i == 0 && j == 0 {
				result += "+"
			} else if image.get(i, j) == 0 {
				result += " "
			} else {
				// result += fmt.Sprintf("%v", image.get(i, j))
				if image.get(i, j)&LightPixel > 0 {
					result += "#"
				} else {
					result += "."
				}
			}
		}
		result += "\n"
	}
	return
}

func (image Image) set(x, y int, value PixelStatus) {
	if _, ok := image.grid[y]; !ok {
		image.grid[y] = make(map[int]PixelStatus)
	}
	image.grid[y][x] = value
}

func (image Image) get(x, y int) PixelStatus {
	if _, ok := image.grid[y]; !ok {
		image.grid[y] = make(map[int]PixelStatus)
	}
	return image.grid[y][x]
}

func (image Image) lit(x, y int) {
	if _, ok := image.grid[y]; !ok {
		image.grid[y] = make(map[int]PixelStatus)
	}
	image.grid[y][x] |= PaintingPixel
}

func (image Image) clear(x, y int) {
	if _, ok := image.grid[y]; !ok {
		image.grid[y] = make(map[int]PixelStatus)
	}
	image.grid[y][x] &^= PaintingPixel
}

func (image Image) process() {
	for j := range image.grid {
		for i := range (image.grid)[j] {
			image.grid[j][i] >>= 1
		}
	}
}

func (image Image) count() (n int) {
	for j := range image.grid {
		for i := range image.grid[j] {
			if image.grid[j][i]&LightPixel > 0 {
				n++
			}
		}
	}
	return
}

func (image Image) probe(x, y int) int {
	n := 0
	for j := y - 1; j <= y+1; j++ {
		for i := x - 1; i <= x+1; i++ {
			if j < image.ymin || j > image.ymax || i < image.xmin || i > image.xmax {
				if image.infinite&LightPixel > 0 {
					n |= 1
				}
			} else if image.get(i, j)&LightPixel > 0 {
				n |= 1
			}
			n <<= 1
		}
	}
	n >>= 1
	return n
}

func (image Image) enlarge() {
	image.xmin--
	image.xmax--
	image.ymin--
	image.ymax++

	image.grid[image.ymin] = make(map[int]PixelStatus)
	image.grid[image.ymax] = make(map[int]PixelStatus)
	for y := -image.ymin; y <= image.ymax; y++ {
		image.grid[y][image.xmin] = image.infinite
		image.grid[y][image.xmax] = image.infinite
	}
	for x := -image.xmin; x <= image.xmax; x++ {
		image.grid[image.ymin][x] = image.infinite
		image.grid[image.ymin][x] = image.infinite
	}
}

func enhance(algorithm string, image Image) func(int) int {
	return func(steps int) int {
		for steps > 0 {
			steps--

			image.ymin -= 1
			image.ymax += 1
			image.xmin -= 1
			image.xmax += 1
			if algorithm[0] == '#' {
				for j := image.ymin; j <= image.ymax; j++ {
					for i := image.xmin; i <= image.xmax; i++ {
						if j == image.ymin || j == image.ymax || i == image.xmin || i == image.xmax {
							if image.infinite&LightPixel > 0 {
								image.set(i, j, LightPixel)
								image.clear(i, j)
							} else {
								image.set(i, j, DarkPixel)
								image.lit(i, j)
							}
						}
					}
				}
			}
			for j := image.ymin; j <= image.ymax; j++ {
				for i := image.xmin; i <= image.xmax; i++ {
					pixel := algorithm[image.probe(i, j)]
					if pixel == '#' {
						image.lit(i, j)
					} else if pixel == '.' {
						image.clear(i, j)
					} else {
						panic("bad input: bad algorithm")
					}
				}
			}

			image.process()
			if algorithm[0] == '#' {
				image.infinite ^= LightPixel
			}
		}
		return image.count()
	}
}

func newImageFromInput(input []string) Image {
	height := len(input)
	width := len(input[0])

	image := Image{
		grid:     map[int]map[int]PixelStatus{},
		ymin:     -height / 2,
		ymax:     height / 2,
		xmin:     -width / 2,
		xmax:     width / 2,
		infinite: DarkPixel,
	}
	for j := 0; j < height; j++ {
		for i := 0; i < width; i++ {
			if input[j][i] == '#' {
				image.lit(i-width/2, j-height/2)
			}
		}
	}
	image.process()
	return image
}

func parseInput(input []string) (string, Image) {
	image := newImageFromInput(input[2:])
	return input[0], image
}
func A(input *challenge.Challenge) int {
	return enhance(parseInput(input.LineSlice()))(2)
}

func B(input *challenge.Challenge) int {
	return enhance(parseInput(input.LineSlice()))(50)
}
