package _2502_Mid

import "testing"

func TestAllocator_Allocate(t *testing.T) {
	al := Constructor(10)
	al.Allocate(1, 1)
	al.Allocate(1, 2)
	al.Allocate(1, 3)
	al.FreeMemory(2)
	al.Allocate(3, 4)
	al.Allocate(1, 1)
	al.Allocate(1, 1)
	al.FreeMemory(1)
	al.Allocate(10, 2)
	al.FreeMemory(7)
}
