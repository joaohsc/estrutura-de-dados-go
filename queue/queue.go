package queue

type IQueue interface {
	Enqueue(value int) error
	Dequeue() (int, error)
	Front() (int, error)
	IsEmpty() bool
	Size() int
}
