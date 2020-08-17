package model

import (
	"go-catboost/tree"
)

type Model struct {
	Roots []tree.TreeNode
	Bias float64
	Scale float64
}

func (m *Model) Predict(input map[string]string) float64{
	result := float64(0.0)
	for _,root := range m.Roots {
		result += root.Compute(input)
	}
	return result*m.Scale + m.Bias
}

func (m *Model) PredictSubTrees(input map[string]string, startTree uint, endTree uint) float64{
	result := float64(0.0)
	for i := startTree;i<endTree;i++ {
		root := m.Roots[i];
		result += root.Compute(input)
	}
	return result*m.Scale + m.Bias;
}

