package jsondeserializer

import "go-catboost/features"

type FeatureComponentDeserializer struct {
	
}

func (fcd FeatureComponentDeserializer) Deserialize(jsonObject map[string]interface{}, numberOfFloatFeatures int, featureNames map[int]string, hashes map[string]uint32, hashNotPresent uint32) features.FeatureComponent{

	if _,ok := jsonObject["cat_feature_index"];ok && jsonObject["combination_element"] == "cat_feature_exact_value"{
		return features.OneHotFeatureComponent{featureNames[numberOfFloatFeatures + int(jsonObject["cat_feature_index"].(float64))], int32(jsonObject["value"].(float64)), hashes, hashNotPresent}
	}else if _,ok := jsonObject["float_feature_index"]; ok {
		return features.FloatFeatureComponent{
			featureNames[int(jsonObject["float_feature_index"].(float64))], jsonObject["border"].(float64)}
	}
	return features.CategoricalFeatureComponent{featureNames[numberOfFloatFeatures + int(jsonObject["cat_feature_index"].(float64))], hashes, hashNotPresent}
}
