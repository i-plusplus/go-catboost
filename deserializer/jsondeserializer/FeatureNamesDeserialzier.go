package jsondeserializer

type FeatureNamesDeserialzier struct {

}

func (fnd FeatureNamesDeserialzier) Deserializer(features map[string]interface{}) map[int]string{
	featureNames := make(map[int]string)
	index := 0;
	if a,ok := features["float_features"]; ok && a != "null"{
		array := a.([]interface{})
		for _, jsonObject2 := range  array {
			jsonObject1 := jsonObject2.(map[string]interface{})
			if  featureName, ok := jsonObject1["feature_name"]; !ok || featureName == "null" {
				panic("feature_name is missing in float features ")
			}
			featureNames[index] = jsonObject1["feature_name"].(string);
			index++;
		}
	}

	if a,ok := features["categorical_features"]; ok && a != "null"{
		array := a.([]interface{})
		for _, jsonObject2 := range  array {
			jsonObject1 := jsonObject2.(map[string]interface{})
			if  featureName, ok := jsonObject1["feature_name"]; !ok || featureName == "null" {
				panic("feature_name is missing in categorical features ")
			}
			featureNames[index] = jsonObject1["feature_name"].(string);
			index++;
		}
	}

	return featureNames;
}
