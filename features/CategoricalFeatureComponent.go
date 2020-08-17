package features

import (
	"go-catboost/hash"
)

type CategoricalFeatureComponent struct {
	FeatureName string
	Hashes map[string]uint32
	HashNotPresent uint32
}

func (cfc CategoricalFeatureComponent) GetKey(old uint64, input map[string]string) uint64 {
	featureValue := input[cfc.FeatureName]
	a := int32(hash.CityHash{}.CalcCatFeatureHash(featureValue, cfc.Hashes, cfc.HashNotPresent))
	return hash.CityHash{}.CalcHash(old, uint64(a));
}

func (cfc CategoricalFeatureComponent) ToString() string {
	return "CFC" + cfc.FeatureName
}
