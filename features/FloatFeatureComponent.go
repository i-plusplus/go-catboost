package features

import (
	"go-catboost/hash"
	"strconv"
)

type FloatFeatureComponent struct {
	FeatureName string
	Border float64
}
func (ffc FloatFeatureComponent) GetKey(old uint64, input map[string]string) uint64{
	featureValue,_ := strconv.ParseFloat(input[ffc.FeatureName], 64)
	value := uint64(0);
	if(featureValue > ffc.Border){
		value = 1;
	}
	return hash.CityHash{}.CalcHash(old, value);
}

func (ffc FloatFeatureComponent) ToString() string {
	return "FFC" + ffc.FeatureName + strconv.Itoa(int(ffc.Border))
}
