package jsondeserializer

type ScaleBiasDesieralizer struct {
	
}
func (sbd ScaleBiasDesieralizer) Deserialize(model map[string]interface{}) (float64, float64) {
	if v, ok := model["scale_and_bias"]; ok{
		array := v.([]float64);
		return array[0], array[1]
	}
	return 1,0
}