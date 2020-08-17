package jsondeserializer

import (
	"encoding/json"
	"go-catboost/features"
)

type FeatureDeserializer struct {
	
}
func (fd FeatureDeserializer) Deserialize(json map[string]interface{}, numberOfFloatFeatures int, featureNames map[int]string, hashes map[string]uint32, hashNotPresent uint32) features.Feature {
	featureComponents := make([]features.FeatureComponent,0)
	identifier := json["identifier"].([]interface{})
	for _,i := range identifier {
		featureComponents = append(featureComponents, FeatureComponentDeserializer{}.Deserialize(i.(map[string]interface{}),numberOfFloatFeatures,featureNames, hashes, hashNotPresent));
	}

	return features.Feature{featureComponents, json["type"].(string)}
}


func (fd FeatureDeserializer) DeserializeS(json2 string, numberOfFloatFeatures int, featureNames map[int]string, hashes map[string]uint32, hashNotPresent uint32) features.Feature{
	obj := make(map[string]interface{})
	_ = json.Unmarshal([]byte(json2), &obj)
	return fd.Deserialize(obj, numberOfFloatFeatures, featureNames, hashes, hashNotPresent);
}