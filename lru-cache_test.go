package leetcode_solutions_golang

import (
	"container/list"
	"testing"
)

func TestNewLruCache(t *testing.T) {
	type args struct {
		capacity int
	}
	tests := []struct {
		name string
		args args
		want LRUCache
	}{
		{
			name: "test case 1",
			args: args{
				capacity: 3,
			},
			want: LRUCache{
				q:    list.New(),
				used: make(map[int]*list.Element, 3),
				data: make(map[int]int, 3),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewLruCache(tt.args.capacity)
			if got.q.Len() != tt.want.q.Len() {
				t.Errorf("NewLruCache() = %v, want %v", got, tt.want)
			}
			if len(got.used) != len(tt.want.used) {
				t.Errorf("NewLruCache() = %v, want %v", got, tt.want)
			}
			if len(got.data) != len(tt.want.data) {
				t.Errorf("NewLruCache() = %v, want %v", got, tt.want)
			}
		})
	}
}
