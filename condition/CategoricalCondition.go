package condition

import (
	"go-catboost/beans"
	"go-catboost/features"
)

type ICategoricalCondition interface {
	getPriorDenominator() float64
	getPriorNumerator() float64
	getScale() float64
	getShift() float64
	getBorder() float64
	IsLeft(input map[string]string) bool
}
type CategoricalCondition struct {
	Feature features.Feature
	StatsC StatsCalculator
	Stats map[uint64]beans.CategoricalStats
	PriorNumerator float64
	PriorDenominator float64
	Scale float64
	Shift float64
	Border float64
}

func (cc CategoricalCondition) getPriorDenominator() float64 {
	return cc.PriorDenominator;
}

func (cc CategoricalCondition) getPriorNumerator() float64{
	return cc.PriorNumerator;
}

func (cc CategoricalCondition) getScale() float64{
	return cc.Scale;
}

func (cc CategoricalCondition) getShift() float64{
	return cc.Shift;
}

func (cc CategoricalCondition) getBorder() float64{
	return cc.Border;
}


func (cc CategoricalCondition)  IsLeft(input map[string]string) bool{
	hashValue := cc.Feature.GetHash(input);

	categoricalStats,ok := cc.Stats[hashValue];
	if(!ok){
		categoricalStats = beans.CategoricalStats{float64(0.0),float64(0),-1};
	}

	denominator := cc.StatsC.getDenominator(categoricalStats, cc.getPriorDenominator())
	numenator := cc.StatsC.getNumerator(categoricalStats, cc.getPriorNumerator());

	nodeValue := ((numenator/denominator) + cc.getShift()) * cc.getScale();

	if(nodeValue > cc.getBorder()){
		return true;
	}
	return false;
}

