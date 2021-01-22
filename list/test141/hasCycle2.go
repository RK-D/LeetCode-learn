package test141

// 双指针
// 龟兔赛跑思想
// 种植终止条件很简单，因为没有环，
// 所以会存在r 先到达 nil 尾指针 为空，这时可以直接终止，
// 否则反复转圈赛跑看谁快
func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	l, r := head, head.Next
	for l != r {
		if r == nil || r.Next == nil {
			return false
		}
		l = l.Next
		r = r.Next.Next
	}
	return true
}
