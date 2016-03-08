package main

type circularBuffer struct {
	buffer  []int64
	pointer int
	added   int
}

func newCircularBuffer(size int) *circularBuffer {
	return &circularBuffer{buffer: make([]int64, size, size)}
}

func (b *circularBuffer) add(element int64) {
	b.buffer[b.pointer] = element
	b.pointer = (b.pointer + 1) % len(b.buffer)
	b.added++
}

func (b *circularBuffer) safeBuffer() []int64 {
	if b.added < len(b.buffer) {
		return b.buffer[0:b.added] // buffer not full
	}
	return b.buffer
}
