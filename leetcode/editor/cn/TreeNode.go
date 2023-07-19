package cn

import (
	"errors"
	"leetcodehelper/helper/model"
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (node *TreeNode) Convert(parameter string) (model.HelperNode, error) {
	if strings.Contains(parameter, ".") {
		return nil, errors.New("parameter may have decimal, not supported yet")
	}

	list := strings.Split(parameter[1:len(parameter)-1], ",")
	if len(list) == 0 {
		return nil, nil
	}

	// 构造根节点
	node = &TreeNode{}
	head := node
	rootVal, err := strconv.Atoi(list[0])
	if err != nil {
		return nil, err
	}
	node.Val = rootVal

	list = list[1:]

	queue := make([]*TreeNode, 0)
	poll := false

	for _, val := range list {
		if val != "null" {
			newNode := &TreeNode{}
			nodeVal, err := strconv.Atoi(val)
			if err != nil {
				return nil, err
			}
			newNode.Val = nodeVal
			if poll {
				node.Right = newNode
			} else {
				node.Left = newNode
			}
			queue = append(queue, newNode)
		}
		if poll {
			node = queue[0]
			queue = queue[1:]
		}
		poll = !poll
	}

	return head, nil
}

func (node *TreeNode) String() string {
	if node == nil {
		return "null"
	}

	var sb strings.Builder
	sb.WriteString("[")
	queue := []*TreeNode{node}

	nullSize := 0

	for len(queue) > 0 && len(queue) != nullSize {
		node = queue[0]
		queue = queue[1:]

		if node != nil {
			sb.WriteString(strconv.Itoa(node.Val))
			sb.WriteString(",")
			queue = append(queue, node.Left)
			queue = append(queue, node.Right)

			if node.Left == nil {
				nullSize++
			}
			if node.Right == nil {
				nullSize++
			}
		} else {
			sb.WriteString("null,")
			nullSize--
		}
	}

	sb.WriteString("]")
	result := sb.String()
	if len(result) > 1 {
		result = result[:len(result)-2] + result[len(result)-1:]
	}
	return result
}
