package robotsVlasers

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestConveyors(t *testing.T) {
	Convey("Subject: Conveyor Methods", t, func() {
		Convey("Make a new Belt", func() {
			So(new(Belt), ShouldNotBeNil)
		})
		Convey("Parse a belt segment", func() {
			seg := SegmentFromColString("|-#")
			So(seg, ShouldNotBeNil)
			So(seg, ShouldResemble, BeltSegment{true, false, false})
			seg = SegmentFromColString("#X|")
			So(seg, ShouldResemble, BeltSegment{false, true, true})
		})
		Convey("Make a new belt", func() {
			So(NewBelt("|#\n--\n#|"), ShouldNotBeNil)
		})
		Convey("Subject: Segment iterators", func() {
			bi := BeltIterator{&BeltDescription{"abcd", "12300", "45678"}, 0}
			Convey("Length is length of first line", func() {
				So(bi.Length(), ShouldEqual, 4)
			})
			Convey("Next returns true until it reaches the end of the belt description", func() {
				bi := bi //don't update other test cases
				So(bi.Next(), ShouldBeTrue)
			})
			Convey("Next is the next column", func() {
				bi := bi //don't update other test cases
				bi.Next()
				So(bi.Value(), ShouldEqual, "b25")
			})
			Convey("Loop ranges are correct", func() {
				bi := bi //don't update other test cases
				result := bi.Value()
				bi.Next()
				result += bi.Value()
				So(result, ShouldEqual, "a14b25")
			})
		})
		Convey("Parse a series of segments into a belt", func() {
			belt := NewBelt("|#\n--\n#|")
			expectedSegments := []BeltSegment{BeltSegment{true, false, false}, BeltSegment{false, true, false}}
			So(belt, ShouldResemble, Belt{expectedSegments, 0})
		})
		Convey("Parse the starting location", func() {
			belt := NewBelt("|#|\n-X-\n#|#")
			So(belt.start, ShouldEqual, 1)
		})
	})
}

func TestController(t *testing.T) {
	Convey("Subject: Conveyor Controls", t, func() {
		ctrl := ConveyorController{}
		Convey("Decides on a direction", func() {
			So(ctrl.BestDirection(), ShouldNotBeNil)
		})
		Convey("Subject: Simulating the line", func() {
			Convey("Getting the current segment uses the starting position", func() {
				ctrl.Belt, ctrl.position = NewBelt("#|#\n-X-\n#|#"), 1
				So(ctrl.CurrentSegment(), ShouldResemble, BeltSegment{true, true, true})
			})
			Convey("North lasers are fired on even turns", func() {
				ctrl.Belt, ctrl.hits, ctrl.position = NewBelt("|\nX\n#"), 0, 0
				ctrl.Simulate()
				So(ctrl.hits, ShouldEqual, 1)
			})
			Convey("South lasers are fired on odd turns", func() {
				ctrl.Belt, ctrl.hits, ctrl.position = NewBelt("##\nX-\n#|"), 0, 1
				ctrl.Simulate()
				So(ctrl.hits, ShouldEqual, 1)
			})
			Convey("Running a simulation adds up all laser hits", func() {
				ctrl.Belt, ctrl.hits = NewBelt("#|#|#X###\n54321X123\n|#|#|X##|"), 0
				So(ctrl.RunSimulation(East), ShouldEqual, 1)
				So(ctrl.RunSimulation(West), ShouldEqual, 5)
			})
		})
		Convey("Chooses a direction with fewer laser hits", func() {
			ctrl.Belt = NewBelt("|||###\n--X---\n######")
			So(ctrl.BestDirection(), ShouldEqual, "GO EAST")
			ctrl.Belt = NewBelt("######\n--X---\n###|||")
			So(ctrl.BestDirection(), ShouldEqual, "GO WEST")
		})
	})
}
