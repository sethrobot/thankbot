package main

type circularBuffer struct {
	buffer []int64
	added  int
}

func newCircularBuffer(size int) *circularBuffer {
	return &circularBuffer{buffer: make([]int64, size, size)}
}

func (b *circularBuffer) add(element int64) {
	b.buffer[b.pointer()] = element
	b.added++
}

func (b *circularBuffer) pop() int64 {
	b.added--
	return b.buffer[b.pointer()]
}

func (b *circularBuffer) safeBuffer() []int64 {
	count := len(b.buffer)
	if b.added < count {
		count = b.added
	}
	var ret []int64
	for i := 0; i < count; i++ {
		ret = append(ret, b.pop())
	}
	return ret
}

func (b *circularBuffer) pointer() int {
	return b.added % len(b.buffer)
}
