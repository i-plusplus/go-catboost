package condition

import "go-catboost/beans"

type StatsCalculator interface {
 getNumerator(categoricalStats beans.CategoricalStats, priorNumerator float64) float64
 getDenominator(categoricalStats beans.CategoricalStats, priorDenominator float64) float64
}
