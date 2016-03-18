package main

import "testing"

func TestCircularBuffer(t *testing.T) {
	buffer := newCircularBuffer(3)
	for _, i := range []int64{1, 2, 3, 4, 5, 6} {
		buffer.add(i)
	}
	var pop int64
	pop = buffer.pop()
	if pop != 6 {
		t.Fatalf("1st pop = %d, expected 6", pop)
	}
	pop = buffer.pop()
	if pop != 5 {
		t.Fatalf("2nd pop = %d to equal 5", pop)
	}
	pop = buffer.pop()
	if pop != 4 {
		t.Fatalf("3rd pop = %d to equal 4", pop)
	}
}

func TestSafeBuffer(t *testing.T) {
	buffer := newCircularBuffer(3)
	for _, i := range []int64{1, 2} {
		buffer.add(i)
	}
	ret := buffer.safeBuffer()
	if len(ret) != 2 {
		t.Fatalf("safeBuffer length = %d, expected 2", len(ret))
	}
	if ret[0] != 2 {
		t.Fatalf("safeBuffer[0] = %d, expected 2", ret[0])
	}
	if ret[1] != 1 {
		t.Fatalf("safeBuffer[1] = %d, expected 1", ret[1])
	}
}
