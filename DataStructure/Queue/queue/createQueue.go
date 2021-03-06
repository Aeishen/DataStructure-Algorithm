package Queue

type item interface {}

type ItemsQueue struct {
	Items []item
}

// 创建一个队列
func NewQueue() *ItemsQueue {
	return &ItemsQueue{[]item{}}
}

//入队
func (this *ItemsQueue)Enqueue(i item){
	this.Items = append(this.Items,i)
}

//获取首元素
func (this *ItemsQueue)GetFront() item{
	return this.Items[0]
}

//出队
func (this *ItemsQueue)Dequeue() {
	this.Items = this.Items[1:]
}

//获得队列长度
func (this *ItemsQueue)GetSzie() int{
	return len(this.Items)
}

//队列判空
func (this *ItemsQueue)IsEmoty() bool{
	return len(this.Items) <= 0
}