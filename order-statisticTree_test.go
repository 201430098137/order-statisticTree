package redBlackTree

import (
	"fmt"
)

/**
 * @Author : 庄广
 * @File : order-statisticTree_test
 * @Time : 2021/6/5 20:12
 * @Software: GoLand
 */

func ExampleOrderStatisticTree_Rank() {
	tree := NewOrderStatisticTree()
	nums := []int{
		1,2,9,5,7,15,7,14,8,2,3,10,11,23,
	}
	nodes := make([]*Node, 0, len(nums))
	for i:=0; i<len(nums); i++ {
		node := &Node{
			Val:Val(nums[i]),
		}

		tree.Insert(node)
		nodes = append(nodes, node)
	}

	tree.Delete(nodes[9])
	tree.Delete(nodes[7])

	for i:=0; i<len(nums); i++ {
		r := tree.Rank(nodes[i])
		fmt.Println(r)
	}

	// OUTPUT:
	// 1
	// 2
	// 8
	// 4
	// 6
	// 11
	// 5
	// 5
	// 7
	// 1
	// 3
	// 9
	// 10
	// 12
}

type Val int

func (item Val) LessThan(other interface{}) bool {
	o := int(other.(Val))
	if int(item) <= o {
		return true
	}
	return false
}
