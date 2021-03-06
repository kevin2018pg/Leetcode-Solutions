package offer

//链表中倒数第k个节点
//输入一个链表，输出该链表中倒数第k个节点。为了符合大多数人的习惯，本题从1开始计数，即链表的尾节点是倒数第1个节点。
//例如，一个链表有6个节点，从头节点开始，它们的值依次是1、2、3、4、5、6。这个链表的倒数第3个节点是值为4的节点。

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getKthFromEnd(head *ListNode, k int) *ListNode {
	//初始化：前指针former，后latter都指向链表头结点
	former, latter := head, head
	//构建双指针距离，前指针向前走k步。结束后，两指针相距k步
	for i := 0; i < k; i++ {
		//如果k大于链表长度，则需要考虑越界问题
		if former == nil{
			return nil
		}
		former = former.Next
	}
	//循环：双指针同时向后移动，直到former跳出，latter与尾节点相距k-1，即是倒数第k个节点
	for former != nil {
		former, latter = former.Next, latter.Next
	}
	return latter
}
