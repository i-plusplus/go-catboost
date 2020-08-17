package condition

import "go-catboost/beans"

type CountCategoricalCondition struct {

}

func (ccc CountCategoricalCondition) getNumerator(categoricalStats beans.CategoricalStats, priorNumerator float64) float64{
return categoricalStats.GetNumerator() + priorNumerator
}

func (ccc CountCategoricalCondition) getDenominator(categoricalStats beans.CategoricalStats, priorDenominator float64) float64{
return categoricalStats.GetDenominator() + priorDenominator
}
