package jsondeserializer

import (
	"go-catboost/beans"
	"strconv"
)

type CTRDataDeserializer struct {

}

func (cdd CTRDataDeserializer) Deserilize(jsonObject map[string]interface{}, numberOfNumericalFeatures int , featureNames map[int]string, hashes map[string]uint32, hashNotPresent uint32) map[string]map[uint64]beans.CategoricalStats {
	 map2 := make(map[string]map[uint64]beans.CategoricalStats)
	for k,v := range jsonObject {
		key := FeatureDeserializer{}.DeserializeS(k, numberOfNumericalFeatures, featureNames, hashes, hashNotPresent);

		value := make(map[uint64]beans.CategoricalStats)
		obj := v.(map[string]interface{})
		counterDenominator := int(0)
		if val,ok := obj["counter_denominator"]; ok {
			counterDenominator = int(val.(float64))
		}
		hashStride := int(1)
		if val,ok := obj["hash_stride"]; ok {
			hashStride = int(val.(float64))
		}
		jsonArray := obj["hash_map"].([]interface{})
		for i:=0;i<len(jsonArray);i+=hashStride {
			hash,_ := strconv.ParseUint(jsonArray[i].(string), 10, 64)
			var categoricalStats beans.CategoricalStats
			if hashStride == 3 {
				categoricalStats = beans.CategoricalStats{jsonArray[i+1].(float64), jsonArray[i+2].(float64), i/3.0}
			}else {
				categoricalStats = beans.CategoricalStats{float64(counterDenominator),jsonArray[i+1].(float64),i/3.0 }
			}
			value[hash] =categoricalStats
		}
		map2[key.ToString()] = value
	}
	return map2
}

