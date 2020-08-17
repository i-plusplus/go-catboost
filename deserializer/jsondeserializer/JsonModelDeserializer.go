package jsondeserializer

import (
	"go-catboost/beans"
	"go-catboost/condition"
	"go-catboost/model"
	"go-catboost/tree"
)

type JsonModelDeserializer struct {
	
}



func (jmd JsonModelDeserializer) deserialize(jsonModel map[string]interface{}, featureNames map[int]string, hashNotPresent uint32) model.Model{
	featureInfos := jsonModel["features_info"].(map[string]interface{})
	floatFeatures := featureInfos["float_features"].([]interface{})
	numberOfNumericalFeatures := len(floatFeatures)
	hashes := (&CatHashDeserializer{}).Deserialize(featureInfos)
    var map2 map[string]map[uint64]beans.CategoricalStats =
	CTRDataDeserializer{}.Deserilize(jsonModel["ctr_data"].(map[string]interface{}), numberOfNumericalFeatures, featureNames, hashes, hashNotPresent)
	var conditionMap map[int]condition.Condition = ConditionSerializer{}.Serialize(jsonModel["features_info"].(map[string]interface{}), jsonModel["oblivious_trees"].([]interface{}), featureNames ,map2, hashes, hashNotPresent)
    var nodes []tree.TreeNode = TreeSerializer{}.Deserialize(jsonModel["oblivious_trees"].([]interface{}), conditionMap);

    scale, bias  := ScaleBiasDesieralizer{}.Deserialize(jsonModel);
	return model.Model{nodes, bias, scale}
}

func (jmd JsonModelDeserializer) Deserialize(jsonModel map[string]interface{}) model.Model{
	return jmd.deserialize(jsonModel, FeatureNamesDeserialzier{}.Deserializer(jsonModel["features_info"].(map[string]interface{})), uint32(2^32-1));
}