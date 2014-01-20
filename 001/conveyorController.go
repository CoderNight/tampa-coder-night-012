package robotsVlasers

type Direction int

const (
	East Direction = 1
	West Direction = -1
)

type ConveyorController struct {
	Belt
	hits     int
	position int
}

func (c *ConveyorController) CurrentSegment() BeltSegment {
	return c.segments[c.position]
}

func (c *ConveyorController) IsTurnEven() bool {
	return (c.position-c.start)%2 == 0
}

func (c *ConveyorController) LaserFired() bool {
	if c.IsTurnEven() {
		return c.CurrentSegment().laserNorth
	}
	return c.CurrentSegment().laserSouth
}

func (c *ConveyorController) Simulate() {
	if c.LaserFired() {
		c.hits++
	}
}

func (c *ConveyorController) Move(direction Direction) bool {
	c.position += int(direction)
	return c.position >= 0 && c.position < len(c.segments)
}

func (c ConveyorController) RunSimulation(direction Direction) int {
	c.position, c.hits = c.start, 0
	for c.Move(direction) {
		c.Simulate()
	}
	return c.hits
}

func (c *ConveyorController) BestDirection() string {
	if c.RunSimulation(East) < c.RunSimulation(West) {
		return "GO EAST"
	}
	return "GO WEST"
}
