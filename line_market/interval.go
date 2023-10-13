package line_market

import (
	"fmt"
	"math"
)

const (
	MinInterval = 0
	MaxInterval = math.MaxInt64
	Infinity    = MaxInterval - 1
)

// Interval is an interval with integer bounds [min, max). By convention, the interval contains the lower bound but excludes the highest bound
type Interval struct {
	min int64
	max int64
}

// NewInterval creates a new interval with min and max options. It always returns a valid interval
func NewInterval(min int64, max int64) *Interval {
	in := &Interval{
		min: min,
		max: max,
	}
	if min < MinInterval {
		in.min = MinInterval
	}
	if max < in.min {
		in.max = in.min
	}
	return in
}

// Empty returns true when the interval is empty
func (in *Interval) Empty() bool {
	return in.min == in.max
}

func (in *Interval) Min() int64 {
	return in.min
}

func (in *Interval) Max() int64 {
	return in.max - 1
}

// Intersection returns the intersection of the intervals
func (in *Interval) Intersection(other *Interval) *Interval {
	if in.Empty() || other.Empty() {
		return NewInterval(0, 0)
	}
	lowestMax := in.max
	highestMin := in.min
	if lowestMax > other.max {
		lowestMax = other.max
	}
	if highestMin < other.min {
		highestMin = other.min
	}
	return NewInterval(highestMin, lowestMax)
}

func (in *Interval) String() string {
	upper := "+"
	if in.Empty() {
		return "[]"
	}
	if in.max < MaxInterval {
		upper = fmt.Sprintf("%d]", in.Max())
	}
	return fmt.Sprintf("[%d,%s", in.min, upper)
}

func (in *Interval) Contains(value int64) bool {
	return value >= in.min && value < in.max
}
