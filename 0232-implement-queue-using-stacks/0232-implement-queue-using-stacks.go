type Node struct {
	value int
	next  *Node
}

type Stack struct {
	head   *Node
	length int
}

type MyQueue struct {
	stack1 *Stack
	stack2 *Stack
}

func (s *Stack) Push(value int) {
	n := &Node{value: value}
	if s.head == nil {
		s.head = n
	} else {
		n.next = s.head
		s.head = n
	}
	s.length++
}
func (s *Stack) Pop() int {
	var n *Node

	if s.head == nil {
		s.length = 0
		return 0
	}
	n = s.head
	s.head = s.head.next
	s.length--

	return n.value
}

func Constructor() MyQueue {
	return MyQueue{
		stack1: &Stack{length: 0},
		stack2: &Stack{length: 0},
	}
}

func (this *MyQueue) Push(x int) {
	this.stack1.Push(x)
}

func (this *MyQueue) Pop() int {
	if this.Empty() {
		return 0
	}

	v := this.stack1.length
	for i := 0; i < v-1; i++ {
		this.stack2.Push(this.stack1.Pop())
	}
	value := this.stack1.Pop()

	v = this.stack2.length
	for i := 0; i < v; i++ {
		this.stack1.Push(this.stack2.Pop())
	}

	return value
}

func (this *MyQueue) Peek() int {
	if this.Empty() {
		return 0
	}

	v := this.stack1.length
	for i := 0; i < v-1; i++ {
		this.stack2.Push(this.stack1.Pop())
	}
	value := this.stack1.Pop()
	this.stack2.Push(value)

	v = this.stack2.length
	for i := 0; i < v; i++ {
		this.stack1.Push(this.stack2.Pop())
	}

	return value
}

func (this *MyQueue) Empty() bool {
	return this.stack1.length == 0
}