package jsondeserializer

import (
	"fmt"
	"go-catboost/condition"
	"go-catboost/tree"
)

type TreeSerializer struct {
	
}


func (ts TreeSerializer) Deserialize(jsonArray []interface{}, conditionMap map[int]condition.Condition) []tree.TreeNode{

	nodes := make([]tree.TreeNode,0)
	for i,v := range jsonArray {
		if i == 8 {
			fmt.Println(v)
		}
		tn := ts.deserialize(v.(map[string]interface{}), 0, 0, conditionMap, i)
		nodes = append(nodes, tn);

	}
	return nodes;
}

func (ts TreeSerializer) deserialize(jsonObject map[string]interface{}, i int,nodeLocation int, conditionMap map[int]condition.Condition, tn int )  tree.TreeNode {
	var  array []interface{} = make([]interface{},0)

	if v, ok := jsonObject["splits"]; !ok || v == nil{

		array = make([]interface{},0)
	} else{

		array = jsonObject["splits"].([]interface{})
	}
	var treen tree.TreeNode;
	if len(array) == i {
		val := jsonObject["leaf_values"].([]interface{})[nodeLocation].(float64)
		treen = tree.TreeNode{true,
			val,
			nodeLocation,
			nil, nil, nil}
	}else{
		left := (ts.deserialize(jsonObject, i + 1, nodeLocation << 1, conditionMap, tn))
		right := (ts.deserialize(jsonObject, i + 1, (nodeLocation << 1) + 1, conditionMap, tn))
		treen = tree.TreeNode{
			false,
			0,
			0,
			conditionMap[int(array[len(array)-1-i].(map[string]interface{})["split_index"].(float64))],
		&left,
	&right};

	}

	if treen.Condition == nil && treen.IsLeaf == false{
		fmt.Print("e")
	}


	return  treen;
}
