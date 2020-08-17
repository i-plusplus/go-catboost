package condition

import "strconv"

type FloatCondition struct {
	FeatureName string
	Border float64
}

func (fc FloatCondition) IsLeft(input map[string]string) bool{
	featureValue, e := strconv.ParseFloat(input[fc.FeatureName], 64);
	if e != nil{
		panic("float value not parseable " + input[fc.FeatureName])
	}
	if(featureValue > fc.Border ){
		return true;
	}
	return false;
}

