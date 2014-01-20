package robotsVlasers

import "strings"

type BeltDescription []string
type BeltIterator struct {
	*BeltDescription
	position int
}

func (bd *BeltDescription) Iter() BeltIterator {
	return BeltIterator{bd, -1}
}

func (bd *BeltDescription) Length() int {
	return len((*bd)[0])
}

func (bi *BeltIterator) Next() bool {
	bi.position++
	return bi.position < bi.BeltDescription.Length()
}

func (bi *BeltIterator) Value() string {
	result := ""
	for _, s := range *(bi.BeltDescription) {
		result += string(s[bi.position])
	}
	return result
}

func (bi *BeltIterator) Segment() BeltSegment {
	return SegmentFromColString(bi.Value())
}

func NewBeltDescription(s string) *BeltDescription {
	bd := BeltDescription(strings.Split(s, "\n"))
	return &bd
}
