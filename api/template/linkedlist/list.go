package linkedlist

// 单项 linked list ,双向链表 go源码有
// 定义一个元素节点,节点存储值以及下个节点的地址索引
type Node struct {
	val interface{}
	// list *LinkList  源码中用这个指代自己从属于那个链表
	next *Node
}

// 定义一个链表,链表需要一个虚拟头结点 / 哨兵/ 哑结点 ,dummyHead不计入链表中,
// 只是一种链表的优化手段(dummyHead.next可以遍历到所有元素,)
// 双向链表也可以优化
// 同时定义一个链表的长度
type LinkList struct {
	dummyHead *Node
	size      int
}

// 定义一个链表的函数
func New() *LinkList {
	return new(LinkList).InitList()
}

// init方法 ，注意不要和golang 的init冲突
func (l *LinkList) InitList() *LinkList {
	l.dummyHead.next = l.dummyHead
	l.size = 0
	return l
}

// 这里给出一种双向链表的写法,对比上面学习,这是go的struct语法糖
// // 返回List的指针
// func New() *List {
//  l := &List{}// 获取List{}的地址
//	l.length = 0// list初始长度为0
//	l.root.next = &l.root
//	l.root.prev = &l.root
//	return l
// }

// 指定index插入元素,不常用
func (l *LinkList) Add(index int, value interface{}) {
	if index > l.size {
		return
	}
	if index < 0 {
		index = 0
	}
	// 扩大尺寸
	l.size++
	// 先存储dummyHead
	prev := l.dummyHead
	// 添加元素,找到目标index 前一个
	for i := 0; i < index; i++ {
		prev = prev.next
	}
	// 原头插法
	// Node3 := &Node3{val: value, next: prev.next}
	// prev.next = Node3
	// 优化写法
	prev.next = &Node{value, prev.next}

}
func (l *LinkList) AddFirst(val int) {
	l.Add(0, val)
}
func (l *LinkList) AddTail(val int) {
	l.Add(l.size, val)
}

func (l *LinkList) Get(val int) interface{} {
	CheckErr(val, l)
	cur := l.dummyHead
	for i := 0; i < val+1; i++ {
		cur = cur.next
	}
	return cur.val
}
func (l *LinkList) Remove(index int) interface{} {
	CheckErr(index, l)
	prev := l.dummyHead
	// 寻找元素,知道找到index,前一位, prev 是待删除的元素的前一个,也就是index前一个
	for i := 0; i < index; i++ {
		prev = prev.next
	}
	// 删除节点,这里注意顺序
	// 先将
	ret := prev.next
	prev.next = ret.next
	ret.next = nil
	l.size--
	return ret.val
}

// 修改链表的第index(0-based)个位置的元素为e
// 在链表中不是一个常用的操作，练习用：）
func (l *LinkList) Set(index int, value interface{}) {
	CheckErr(index, l)

	cur := l.dummyHead.next
	for i := 0; i < index; i++ {
		cur = cur.next
	}
	cur.val = value
}

// 查找链表是否存在元素e
func (l *LinkList) Contains(value interface{}) bool {
	cur := l.dummyHead.next

	for cur != nil {
		if cur.val == value {
			return true
		}
		cur = cur.next
	}
	return false
}

func CheckErr(index int, l *LinkList) {
	if index < 0 || index > l.size {
		panic("add failed, index is out of range")
	}
}
