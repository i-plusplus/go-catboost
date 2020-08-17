package tree

import (
	"go-catboost/condition"
)

type TreeNode struct {
	IsLeaf bool
	LeafValue float64
	LeafIndex int
	Condition condition.Condition
	Left *TreeNode
	Right *TreeNode
}

func (tn *TreeNode) Compute(input map[string]string) float64 {
	for !tn.IsLeaf {
		if(tn.Condition.IsLeft(input)){
			tn = tn.Right
		}else{
			tn = tn.Left
		}
	}
	return tn.LeafValue;
}
