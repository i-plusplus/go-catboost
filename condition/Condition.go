package condition

type Condition interface {
	IsLeft(input map[string]string) bool
}
