package jsondeserializer

import (
	"go-catboost/beans"
	"go-catboost/condition"
	"go-catboost/features"
)

type ConditionSerializer struct {

}
type splitDetails struct {
	value uint64
	catFeatureIndex int
	splitIndex int
}

 func (cs ConditionSerializer) Serialize( jsonObject map[string]interface{},
 											obliviousTrees []interface{},
 											featureNames map[int]string,
 											hashMap map[string]map[uint64]beans.CategoricalStats,
 											hashes map[string]uint32,
 											hashNotPresent uint32) map[int]condition.Condition {
 	var sd []splitDetails = make([]splitDetails, 0)

	for _, ot := range obliviousTrees {
		var tree = ot.(map[string]interface{})
		splits, ok := tree["splits"]
		if (!ok || splits == nil) {
			continue
		}
		splitList := splits.([]interface{})
		for _, split2:= range splitList {
			split := split2.(map[string]interface{})
			if split["split_type"] == "OneHotFeature" {
				splitD := splitDetails{

			uint64(split["value"].(float64)),
					int(split["cat_feature_index"].(float64)),
			int(split["split_index"].(float64)) }
				sd = append(sd, splitD)
			}
		}
	}

	var index int = 0;
    conditionMap := make(map[int]condition.Condition)
	numberOfNumericalFeatures := len(jsonObject["float_features"].([]interface{}))
/*----------------------float features----------------------------------------*/
    array := jsonObject["float_features"].([]interface{})
	for i := 0;i<numberOfNumericalFeatures;i++ {
		jsonObject1 := array[i].(map[string]interface{})
		featureName := featureNames[int(i)]
		if b,ok := jsonObject1["borders"];ok {
			borders := b.([]interface{})
			for j := 0; j < len(borders); j++{
				conditionMap[index] = condition.FloatCondition {
					featureName,
					borders[j].(float64)}
				index++;
			}
		}
	}
/*-----------------------------OneHotEncoding------------------------------------*/
array = jsonObject["categorical_features"].([]interface{})
for i := 0;i<len(array);i++ {
	 jsonObject1 := array[i].(map[string]interface{})

     flatFeatureIndex := int(jsonObject1["flat_feature_index"].(float64))
     featureName := featureNames[flatFeatureIndex + numberOfNumericalFeatures]
     if borders,ok := jsonObject1["values"]; ok {
     	border := borders.([]interface{})
     	for _, v := range border {
     		conditionMap[index] = condition.OneHotCondition{featureName, hashes,int32(v.(float64)), hashNotPresent}
     		index++
		}
	 }
}


	for _,v := range sd {
	if _, ok := conditionMap[v.splitIndex]; ok{
		continue;
	}
	featureName := featureNames[v.catFeatureIndex + numberOfNumericalFeatures];
	conditionMap[v.splitIndex] =  condition.OneHotCondition{
		featureName, hashes, int32(v.value), hashNotPresent};
	index++;
}

/*-----------------------------ctrs----------------------------------------------*/
	array2 := jsonObject["ctrs"].([]interface{});
	for _, v2 := range array2 {
	v:= v2.(map[string]interface{})

	var borders []float64 = make([]float64,0)
	if val , ok := v["borders"]; ok && val != "null"{
		borders2 := v["borders"].([]interface{})
		for _,b := range borders2 {
			borders = append(borders, b.(float64))
		}
	}

	var key features.Feature= FeatureDeserializer{}.DeserializeS(
v["identifier"].(string), numberOfNumericalFeatures, featureNames, hashes, hashNotPresent)
	var mp map[uint64]beans.CategoricalStats= hashMap[key.ToString()]
	for _,d := range borders {
		priorNumberator := v["prior_numerator"].(float64)
		priorDenomenator := v["prior_denomerator"].(float64)
		scale := v["scale"].(float64)
		shift := v["shift"].(float64)
		ctrType := v["ctr_type"].(string)
		if ctrType == "Borders"{
			conditionMap[index] = condition.CategoricalCondition{key, condition.BorderCategoricalCondition{}, mp, priorNumberator, priorDenomenator, scale, shift, d}
		}else {
			conditionMap[index] = condition.CategoricalCondition{key, condition.CountCategoricalCondition{},mp,  priorNumberator, priorDenomenator, scale, shift, d}
		}
		index++
	}


}
return conditionMap;



}
