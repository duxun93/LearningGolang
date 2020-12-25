//使用两个队列实现一个栈
package stackAndqueue

type StackByQueue struct {
	que1 Queue
	que2 Queue
}

func (rer *StackByQueue) Pop() (num int) {
	if len(rer.que1.A) == 0 && len(rer.que2.A) == 0 {
		return -1
	}
	if len(rer.que1.A) == 0 {
		for len(rer.que2.A) > 1 {
			temp := rer.que2.Pop()
			rer.que1.Push(temp)
		}
		return rer.que2.Pop()
	} else {
		for len(rer.que1.A) > 1 {
			temp := rer.que1.Pop()
			rer.que2.Push(temp)
		}
		return rer.que1.Pop()
	}
}

func (rer *StackByQueue) Push(num int) {
	if len(rer.que1.A) == 0 && len(rer.que2.A) == 0 {
		rer.que1.Push(num)
	} else if len(rer.que1.A) == 0 {
		rer.que2.Push(num)
	} else {
		rer.que1.Push(num)
	}
	return
}

func NewStackByQueue() *StackByQueue {
	result := new(StackByQueue)
	result.que1 = *NewQueue()
	result.que2 = *NewQueue()
	return result
}
