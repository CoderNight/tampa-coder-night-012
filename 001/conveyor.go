package robotsVlasers

type Belt struct {
	segments []BeltSegment
	start    int
}

func (b *Belt) AddSegment(seg BeltSegment) {
	b.segments = append(b.segments, seg)
	if seg.isStart {
		b.start = len(b.segments) - 1
	}
}

func NewBelt(s string) Belt {
	bi := NewBeltDescription(s).Iter()
	belt := Belt{}
	for bi.Next() {
		belt.AddSegment(bi.Segment())
	}
	return belt
}

/*\
 | This is for a discussion point:
 |
\*/

func (sArray BeltDescription) ColIter() <-chan string {
	ch := make(chan string)
	go func() {
		for colIdx := range sArray[0] {
			col := ""
			for s := range sArray {
				col += string(sArray[s][colIdx])
			}
			ch <- col
		}
		close(ch)
	}()
	return ch
}

func NewBeltFromChannel(s string) Belt {
	data := NewBeltDescription(s)
	belt := Belt{}
	for col := range data.ColIter() {
		belt.AddSegment(SegmentFromColString(col))
	}
	return belt
}
