type RecentCounter struct {
	queue []int
}

func Constructor() RecentCounter {
	return RecentCounter{queue: []int{}}
}

func (this *RecentCounter) Push(value int) {
	this.queue = append(this.queue, value)
}

func (this *RecentCounter) Pop() int {
	value := this.queue[0]
	this.queue = this.queue[1:]
	return value
}

func (this *RecentCounter) Peek() int {
	value := this.queue[0]
	return value
}

func (this *RecentCounter) Ping(t int) int {
	this.Push(t)
	old := this.Peek()
	for old < t-3000 {
		this.Pop()
		old = this.Peek()
	}

	return len(this.queue)
}

/**
 * Your RecentCounter object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Ping(t);
 */