package main

type building interface {
	getType() string
	accept(visitor)
}

type visitor interface {
	visitForChurch(*church)
	visitForHotel(*hotel)
	visitForStadium(*stadium)
}

type church struct {
	sideA int
	sideB int
}

func (c *church) getType() string {
	return "Church"
}

func (c *church) accept(v visitor) {
	v.visitForChurch(c)
}

type stadium struct {
	radius int
}

func (s *stadium) getType() string {
	return "Stadium"
}

func (s *stadium) accept(v visitor) {
	v.visitForStadium(s)
}

type hotel struct {
	l int
	b int
}

func (h *hotel) getType() string {
	return "Hotel"
}

func (h *hotel) accept(v visitor) {
	v.visitForHotel(h)
}

func main() {

}
