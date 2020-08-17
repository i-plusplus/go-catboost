package condition

import "go-catboost/beans"

type BorderCategoricalCondition struct {

}

func (bcc BorderCategoricalCondition)  getNumerator(categoricalStats beans.CategoricalStats, priorNumerator float64) float64{
	return categoricalStats.GetNumerator() + priorNumerator
}

func (bcc BorderCategoricalCondition) getDenominator(categoricalStats beans.CategoricalStats, priorDenominator float64) float64{
	return categoricalStats.GetDenominator() + categoricalStats.GetNumerator() + priorDenominator;
}


