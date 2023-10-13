package line_market

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInterval_Intersection(t *testing.T) {
	in1 := NewInterval(0, 10)
	in2 := NewInterval(10, 20)
	inter := in1.Intersection(in2)
	fmt.Printf("intersection of %s and %s is %s\n", in1, in2, inter)
	assert.True(t, inter.Empty())

	in1 = NewInterval(0, 15)
	in2 = NewInterval(10, 20)
	inter = in1.Intersection(in2)
	fmt.Printf("intersection of %s and %s is %s\n", in1, in2, inter)
	assert.False(t, inter.Empty())
	assert.Equal(t, int64(10), inter.Min())
	assert.Equal(t, int64(14), inter.Max())

	inter = in2.Intersection(in1)
	assert.False(t, inter.Empty())
	assert.Equal(t, int64(10), inter.Min())
	assert.Equal(t, int64(14), inter.Max())
}
