package main

import (
	"strconv"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	l1tmp := l1
	var n1 string
	for l1tmp != nil {
		n1 = strconv.Itoa(l1tmp.Val) + n1
		l1tmp = l1tmp.Next
	}
	l2tmp := l2
	var n2 string
	for l2tmp != nil {
		n2 = strconv.Itoa(l2tmp.Val) + n2
		l2tmp = l2tmp.Next
	}

	head := new(ListNode)
	tmp := head
	kuriagari := 0
	var i, j int
	for i, j = len(n1)-1, len(n2)-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		num1, _ := strconv.Atoi(string(n1[i]))
		num2, _ := strconv.Atoi(string(n2[j]))
		sum := num1 + num2 + kuriagari
		var newNode *ListNode

		if sum >= 10 {
			newNode = &ListNode{Val: sum - 10}
			kuriagari = 1
		} else {
			newNode = &ListNode{Val: sum}
			kuriagari = 0
		}
		tmp.Next = newNode
		tmp = tmp.Next
	}
	for j >= 0 {
		var newNode *ListNode
		num2, _ := strconv.Atoi(string(n2[j]))
		sum := kuriagari + num2
		if sum >= 10 {
			newNode = &ListNode{Val: sum - 10}
			kuriagari = 1
		} else {
			newNode = &ListNode{Val: sum}
			kuriagari = 0
		}
		tmp.Next = newNode
		tmp = tmp.Next
		j--
	}
	for i >= 0 {
		var newNode *ListNode
		num1, _ := strconv.Atoi(string(n1[i]))
		sum := kuriagari + num1
		if sum >= 10 {
			newNode = &ListNode{Val: sum - 10}
			kuriagari = 1
		} else {
			newNode = &ListNode{Val: sum}
			kuriagari = 0
		}
		tmp.Next = newNode
		tmp = tmp.Next
		i--
	}
	if kuriagari == 1 {
		tmp.Next = &ListNode{Val: 1}
	}
	return head.Next
}
