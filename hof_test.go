package hof

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestFilter(t *testing.T) {
	tests := []struct {
		name string
		src  []int
		want []int
	}{
		{
			name: "filter even numbers",
			src:  []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			want: []int{2, 4, 6, 8},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := []int{}
			Filter(tt.src, func(v int) bool { return v%2 == 0 })(func(v int) bool {
				got = append(got, v)
				return true
			})
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Fatalf("(-got +want)\n%s", diff)
			}
		})
	}
}

func TestMap(t *testing.T) {
	tests := []struct {
		name string
		src  []int
		want []string
	}{
		{
			name: "map int to string",
			src: []int{
				1, 2, 3, 4, 5,
			},
			want: []string{
				"1", "2", "3", "4", "5",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := []string{}
			Map(tt.src, func(v int) string { return fmt.Sprint(v) })(func(v string) bool {
				got = append(got, v)
				return true
			})
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Fatalf("(-got +want)\n%s", diff)
			}
		})
	}
}

func TestMap2(t *testing.T) {
	tests := []struct {
		name string
		src  map[int]string
		want map[string]int
	}{
		{
			name: "map2",
			src: map[int]string{
				1: "one",
				2: "two",
				3: "three",
			},
			want: map[string]int{
				"one":   1,
				"two":   2,
				"three": 3,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := map[string]int{}
			Map2(tt.src, func(k int, v string) string { return v })(func(k int, v string) bool {
				got[v] = k
				return true
			})
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Fatalf("(-got +want)\n%s", diff)
			}
		})
	}
}

func TestChunk(t *testing.T) {
	tests := []struct {
		name string
		src  []int
		size int
		want [][]int
	}{
		{
			name: "chunk",
			src:  []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			size: 3,
			want: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
		},
	}

	for _, tt := range tests {
		got := [][]int{}
		Chunk(tt.src, tt.size)(func(v []int) bool {
			got = append(got, v)
			return true
		})

		if diff := cmp.Diff(got, tt.want); diff != "" {
			t.Fatalf("(-got +want)\n%s", diff)
		}
	}
}

func TestReduce(t *testing.T) {
	tests := []struct {
		name string
		fn   func() any
		want any
	}{
		{
			name: "reduce[int]",
			fn: func() any {
				got := 0
				Reduce([]int{1, 2, 3, 4, 5}, 0, func(acc, v int) int { return acc + v })(func(v int) bool {
					got = v
					return true
				})
				return got
			},
			want: 15,
		},
		{
			name: "reduce[string]",
			fn: func() any {
				got := ""
				Reduce([]string{"a", "b", "c"}, "", func(acc, v string) string { return acc + v })(func(v string) bool {
					got = v
					return true
				})
				return got
			},
			want: "abc",
		},
		{
			name: "reduce[map]",
			fn: func() any {
				got := map[int]string{}
				Reduce([]int{1, 2, 3}, map[int]string{}, func(acc map[int]string, v int) map[int]string {
					acc[v] = fmt.Sprint(v)
					return acc
				})(func(v map[int]string) bool {
					got = v
					return true
				})
				return got
			},
			want: map[int]string{
				1: "1",
				2: "2",
				3: "3",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.fn()
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Fatalf("(-got +want)\n%s", diff)
			}
		})
	}
}
