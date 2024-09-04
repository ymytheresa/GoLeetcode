/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    dummy := &ListNode{}
    recur(l1, l2, 0, dummy)
    return dummy.Next
}

func recur(l1 *ListNode, l2 *ListNode, carry int, head *ListNode){
    if l1 == nil && l2 == nil && carry == 0{
        return 
    }
    
    l1v, l2v := 0, 0
    if l1 != nil{
        l1v = l1.Val
        l1 = l1.Next
    }
    if l2 != nil{
        l2v = l2.Val
        l2 = l2.Next
    }
    sum := carry + l1v + l2v
    carry = sum/10
    remainder := sum%10
    head.Next = &ListNode{Val: remainder}
    recur(l1, l2, carry, head.Next)
}   