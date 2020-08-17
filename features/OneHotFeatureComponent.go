package features

import (
	"go-catboost/hash"
	"strconv"
)

type OneHotFeatureComponent struct {
	FeatureName string
	Value int32
	Hashes map[string]uint32
	HashNotPresent uint32


}
func (ohfc OneHotFeatureComponent) GetKey(old uint64, input map[string]string) uint64{
	featureValue := input[ohfc.FeatureName];
	fvalue := hash.CityHash{}.CalcCatFeatureHash(featureValue, ohfc.Hashes, ohfc.HashNotPresent);
	h := uint64(0);
	if(fvalue == uint32(ohfc.Value)){
		h = 1;
	}
	old = hash.CityHash{}.CalcHash(old, h);
	return old;
}

func (ohfc OneHotFeatureComponent) ToString() string{
	return "OHE" + ohfc.FeatureName + strconv.Itoa(int(ohfc.Value))
}

