package robotsVlasers

type BeltSegment struct {
	laserNorth bool
	laserSouth bool
	isStart    bool
}

func SegmentFromColString(s string) BeltSegment {
	return BeltSegment{s[0] == '|', s[2] == '|', s[1] == 'X'}
}
