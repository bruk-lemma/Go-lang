// package main

// import "fmt"

func main() {
	list := &List{}
	list.add(1)
	list.add(2)
	list.add(3)
	list.add(4)

	// fmt.Println("Initial List: ")
	// printList(list)

	// list.remove(2)
	// fmt.Println("List after removing 2: ")
	// printList(list)

	// list.remove(4)
	// fmt.Println("List after removing 4: ")
	// printList(list)

	fmt.Println("Initial List before insert: ")
	printList(list)
	fmt.Println("Initial List after insert: ")

	list.Insert(9, 1)
	printList(list)
}

type Node struct {
	data int
	next *Node
	prev *Node
}

type List struct {
	head *Node
}

func (l *List) add(value int) {
	newNode := &Node{data: value}
	if l.head == nil {
		l.head = newNode
		return
	}

	curr := l.head
	for curr.next != nil {
		curr = curr.next
	}
	curr.next = newNode
}

func (l *List) Insert(value int, index int) {
	newNode := &Node{data: value}

	if l.head == nil || index == 0 {
		newNode.next = l.head
		if l.head != nil {
			l.head.prev = newNode
		}
		l.head = newNode
		return
	}

	curr := l.head
	for i := 0; i < index-1; i++ {
		if curr.next == nil {
			break
		}
		curr = curr.next
	}

	newNode.next = curr.next
	newNode.prev = curr
	if curr.next != nil {
		curr.next.prev = newNode
	}
	curr.next = newNode
}

func (l *List) remove(value int) {
	if l.head == nil {
		return
	}

	if l.head.data == value {
		l.head = l.head.next
		return
	}

	curr := l.head
	for curr.next != nil && curr.next.data != value {
		curr = curr.next
	}
	if curr.next != nil {
		curr.next = curr.next.next
	}

}

func printList(l *List) {
	curr := l.head
	for curr != nil {
		fmt.Printf("%d ", curr.data)
		curr = curr.next
	}
	fmt.Println()
}
