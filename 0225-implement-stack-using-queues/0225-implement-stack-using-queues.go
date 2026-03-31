type Node struct {
	value int
	next  *Node
}
type Queue struct {
	head   *Node
	tail   *Node
	length int
}

type MyStack struct {
	queue1 *Queue
	queue2 *Queue
}

func (q *Queue) Push(value int) {
	n := &Node{value: value}

	if q.head == nil {
		q.head = n
		q.tail = n
	} else {
		q.tail.next = n
		q.tail = n
	}
	q.length++

}

func (q *Queue) Pop() int {
	if q.head == nil {
		q.length = 0
		return 0
	}

	value := q.head
	q.head = q.head.next
	q.length--

	if q.length == 0 {
		q.head = nil
		q.tail = nil
	}

	return value.value
}

func Constructor() MyStack {
	return MyStack{queue1: &Queue{}, queue2: &Queue{}}
}

func (this *MyStack) Push(x int) {
	this.queue1.Push(x)
}

func (this *MyStack) Pop() int {
	c := this.queue1.length
	for i := 0; i < c-1; i++ {
		this.queue2.Push(this.queue1.Pop())
	}
	value := this.queue1.Pop()
	this.queue1 = this.queue2
	this.queue2 = &Queue{}

	return value
}

func (this *MyStack) Top() int {
	c := this.queue1.length
	for i := 0; i < c-1; i++ {
		this.queue2.Push(this.queue1.Pop())
	}
	value := this.queue1.Pop()
	this.queue2.Push(value)

	this.queue1 = this.queue2
	this.queue2 = &Queue{}

	return value
}

func (this *MyStack) Empty() bool {
	return this.queue1.length == 0
}


/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */