package main

import "fmt"

type queue struct {
	items []interface{}
}

func (q *queue) additemsrear(item interface{}) {
	q.items = append(q.items, item)
}

func (q *queue) additemfront(item interface{}) {
	itq := make([]interface{}, 0, len(q.items)+1)
	itq = append(itq, item)
	itq = append(itq, q.items...)
	q.items = itq
}

func (q *queue) display() {
	for _, item := range q.items {
		fmt.Print(item, " ")
	}
	fmt.Println()
}
func (q *queue) dequeue() {
	if len(q.items) == 0 {
		fmt.Println("Size is zero")
	} else {
		item := q.items[0]
		q.items = q.items[1:]
		fmt.Println("The  element removed", item)
	}
}

func main() {
	q := queue{}
	q.additemsrear(4)
	q.dequeue()
	q.additemfront(5)
	q.additemfront(7)
	q.display()
}
