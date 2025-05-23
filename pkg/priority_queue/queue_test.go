package priority_queue_test

import (
	"ca/pkg/priority_queue"
	"fmt"
	"testing"
)

func TestQueue(t *testing.T) {
	q := priority_queue.New()
	q.Push(&priority_queue.Item{
		Value:    nil,
		Priority: 1,
	})
	q.Push(&priority_queue.Item{
		Value:    nil,
		Priority: 3,
	})
	i := &priority_queue.Item{
		Value:    nil,
		Priority: 2,
	}
	q.Push(i)

	fmt.Println(q.Pop())
	i.Priority = 5
	q.Update(i)

	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
}
