package condition

import "go-catboost/hash"

type OneHotCondition struct {
	FeatureName string
	Hashes map[string]uint32
	Value int32
	HashNotPresent uint32
}



func (ohc OneHotCondition) IsLeft(input map[string]string) bool{
	featureValue := input[ohc.FeatureName];
	hash := hash.CityHash{}.CalcCatFeatureHash(featureValue, ohc.Hashes, ohc.HashNotPresent);
	if(hash == uint32(ohc.Value)){
		return true;
	}
	return false;
}

