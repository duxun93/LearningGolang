package stackAndqueue

type Stack struct {
	A []int
}
type Queue struct {
	A []int
}

func (rer *Stack) Pop() int {
	if len(rer.A) == 0 {
		return -1
	}
	temp := rer.A[len(rer.A)-1]
	rer.A = rer.A[:len(rer.A)-1]
	return temp
}

func (rer *Stack) Push(num int) {
	rer.A = append(rer.A, num)
	return
}

func NewStack() *Stack {
	stack := new(Stack)
	stack.A = []int{}
	return stack
}

func (rer *Queue) Pop() int {
	if len(rer.A) == 0 {
		return -1
	}
	temp := rer.A[0]
	rer.A = rer.A[1:]
	return temp
}

func (rer *Queue) Push(num int) {
	rer.A = append(rer.A, num)
	return
}

func NewQueue() *Queue {
	queue := new(Queue)
	queue.A = []int{}
	return queue
}
