package tl

import (
	"testing"

	"github.com/d2gr/tl"
	"github.com/d2gr/tl/iter"
)

func TestList(t *testing.T) {
	var list tl.List[int]

	type expectedValue struct {
		fn       func(int)
		iter     func() tl.Iter[int]
		expected []int
	}

	values := iter.ToSlice(iter.Range(0, 10, 1))

	es := []expectedValue{
		{
			fn:       list.PushBack,
			iter:     list.ForwardIter,
			expected: iter.ToSlice(iter.Range(0, 10, 1)),
		}, {
			fn:       list.PushFront,
			iter:     list.ForwardIter,
			expected: iter.ToSlice(iter.Range(9, -1, -1)),
		}, {
			fn:       list.PushBack,
			iter:     list.ReverseIter,
			expected: iter.ToSlice(iter.Range(9, -1, -1)),
		}, {
			fn:       list.PushFront,
			iter:     list.ReverseIter,
			expected: iter.ToSlice(iter.Range(0, 10, 1)),
		},
	}

	for idx, e := range es {
		for _, v := range values {
			e.fn(v)
		}

		iter := e.iter()
		for i := 0; i < len(e.expected); i++ {
			iter.Next()

			if e.expected[i] != iter.Get() {
				t.Fatalf("unexpected value on %d: %d <> %d", idx, e.expected[i], iter.Get())
			}
		}

		list.Reset()
	}
}