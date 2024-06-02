package Leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func listLen(list *ListNode) int {

	var listTemp = list
	var leng = 0

	for listTemp != nil {
		leng++
		listTemp = listTemp.Next
	}
	return leng
}

func calculateDigit(number int) int {
	if number == 0 {
		return 1
	}
	var digit int
	for number > 0 {
		digit++
		number /= 10
	}
	return digit
}

func listExplorer(list *ListNode) int {
	var multiply = 1
	var number = 0

	for list != nil {
		number += list.Val * multiply
		multiply *= 10
		list = list.Next
	}
	return number
}

func addNode(head *ListNode, number int) {
	newNode := &ListNode{Val: number}
	temp := head
	for temp.Next != nil {
		temp = temp.Next
	}
	temp.Next = newNode
}
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	var number1 = listExplorer(l1)
	var number2 = listExplorer(l2)
	var total = number1 + number2
	var digit = calculateDigit(total)
	var answer ListNode
	answer.Val = total % 10
	total /= 10
	for i := 0; i < digit-1; i++ {
		var temp = total % 10
		addNode(&answer, temp)
		total /= 10
	}
	return &answer
}
