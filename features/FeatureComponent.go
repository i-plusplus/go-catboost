package features

type FeatureComponent interface {
	GetKey(old uint64, input map[string]string) uint64
	ToString() string
}
