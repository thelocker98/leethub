package main

type Node struct {
	Val  int
	Next *Node
}

type MyLinkedList struct {
	Head *Node
	Size int
}

func Constructor() MyLinkedList {
	return MyLinkedList{
		Head: nil,
		Size: 0,
	}
}

func (this *MyLinkedList) Get(index int) int {
	if index < 0 || index >= this.Size {
		return -1
	}

	curr := this.Head
	for i := 0; i < index; i++ {
		curr = curr.Next
	}

	return curr.Val
}

func (this *MyLinkedList) AddAtHead(val int) {
	node := &Node{Val: val}
	node.Next = this.Head
	this.Head = node
	this.Size++
}

func (this *MyLinkedList) AddAtTail(val int) {
	node := &Node{Val: val}

	if this.Head == nil {
		this.Head = node
	} else {
		curr := this.Head
		for curr.Next != nil {
			curr = curr.Next
		}
		curr.Next = node
	}

	this.Size++
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index < 0 || index > this.Size {
		return
	}

	if index == 0 {
		this.AddAtHead(val)
		return
	}

	curr := this.Head
	for i := 0; i < index-1; i++ {
		curr = curr.Next
	}

	node := &Node{Val: val}
	node.Next = curr.Next
	curr.Next = node

	this.Size++
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index >= this.Size {
		return
	}

	if index == 0 {
		this.Head = this.Head.Next
	} else {
		curr := this.Head
		for i := 0; i < index-1; i++ {
			curr = curr.Next
		}
		curr.Next = curr.Next.Next
	}

	this.Size--
}


/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */