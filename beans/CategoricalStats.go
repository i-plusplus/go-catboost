package beans

type CategoricalStats struct {
	Denominator float64;
	Numerator float64;
	Location int;
}

func (cs *CategoricalStats) GetDenominator() float64{
    return cs.Denominator;
}

func (cs *CategoricalStats) GetNumerator() float64{
	return cs.Numerator;
}

func (cs *CategoricalStats) GetLocation() int{
	return cs.Location;
}

