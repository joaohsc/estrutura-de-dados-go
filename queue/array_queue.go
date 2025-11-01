package queue

import "errors"

type ArrayQueue struct {
	v     []int
	front int
	rear  int
}

func (q *ArrayQueue) Init(size int) error {
	if size > 0 {
		q.v = make([]int, size)
		q.front = -1
		q.rear = -1
		return nil
	} else {
		return errors.New("Size <= 0")
	}
}

func (q *ArrayQueue) Enqueue(e int) error {
	if q.Size() == len(q.v) {
		return errors.New("fila está lotada")
	}
	if q.front == -1 && q.rear == -1 {
		q.front = 0
		q.rear = 0
	} else {
		q.rear = (q.rear + 1) % len(q.v)
	}
	q.v[q.rear] = e
	return nil
}

func (q *ArrayQueue) Dequeue() (int, error) {
	if q.Size() == 0 {
		return 0, errors.New("impossível desenfileirar de fila vazia")

	} else if q.front == q.rear {
		temp := q.v[q.front]
		q.front = -1
		q.rear = -1
		return temp, nil
	} else {
		temp := q.v[q.front]
		q.front = (q.front + 1) % len(q.v)
		return temp, nil
	}
}

func (q *ArrayQueue) Front() (int, error) {
	if q.Size() == 0 {
		return 0, errors.New("impossível obter frente de fila vazia")
	} else {
		return q.v[q.front], nil
	}
}

func (queue *ArrayQueue) IsEmpty() bool {
	return queue.Size() == 0
}

func (q *ArrayQueue) Size() int {
	if q.front == -1 {
		return 0
	}
	if q.front <= q.rear {
		return q.rear - q.front + 1
	}
	return len(q.v) - q.front + q.rear + 1
}
