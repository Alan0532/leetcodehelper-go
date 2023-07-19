package cn

import (
	"encoding/json"
	"errors"
	"fmt"
	"leetcodehelper/helper/model"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func (ln *ListNode) Convert(parameter string) (model.HelperNode, error) {
	if strings.Contains(parameter, ".") {
		return nil, errors.New("parameter may have decimal, not supported yet")
	}

	var list []int
	if err := json.Unmarshal([]byte(parameter), &list); err != nil {
		return nil, err
	}

	if len(list) == 0 {
		return nil, nil
	}

	ln = &ListNode{Val: list[0]}
	curr := ln
	for i := 1; i < len(list); i++ {
		curr.Next = &ListNode{Val: list[i]}
		curr = curr.Next
	}

	return ln, nil
}

func (ln *ListNode) String() string {
	curr := ln
	var sb strings.Builder
	sb.WriteString("[")
	for curr != nil {
		sb.WriteString(fmt.Sprintf("%d", curr.Val))
		sb.WriteString(",")
		curr = curr.Next
	}
	sb.WriteString("]")
	result := sb.String()
	if len(result) > 1 {
		result = result[:len(result)-2] + result[len(result)-1:]
	}
	return result
}
