package features

type Feature struct {
	FeatureComponents []FeatureComponent
	Types string
}

func (f *Feature) ToString() string{
	var s string = ""
	for _,t := range f.FeatureComponents {
		s = s + t.ToString()
	}
	return s + f.Types
}

func (f *Feature) GetHash(input map[string]string) uint64{
	hash := uint64(0);
	for _,featureComponent := range f.FeatureComponents {
		hash = featureComponent.GetKey(hash, input);
	}
	return hash;
}




