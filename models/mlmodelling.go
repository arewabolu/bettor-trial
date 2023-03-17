package models

import (
	"math"

	"github.com/sajari/regression"
	"gonum.org/v1/gonum/stat"
)

func TrainAndTest(TD *TeamData) (regression.Regression, float64) {
	trainingNum := (4 * len(TD.GoalFor)) / 5
	testNum := len(TD.GoalFor) / 5
	var r regression.Regression
	r.SetObserved("goals")
	r.SetVar(0, "status")
	r.SetVar(1, "goalsAgainst")
	for i := 0; i < trainingNum; i++ {
		r.Train(regression.DataPoint(TD.GoalFor[i], []float64{TD.Status[i], TD.GoalAgainst[i]}))
	}
	// Train/fit the regression model.
	r.Run()
	var mAE float64
	for i := trainingNum; i < trainingNum+testNum; i++ {
		// Predict y with our trained model.
		GFPredicted, _ := r.Predict([]float64{TD.Status[i], TD.GoalAgainst[i]})
		//

		// Add the to the mean absolute error.
		mAE += math.Abs(TD.GoalFor[i]-GFPredicted) / float64(len(TD.GoalFor))
	}
	//
	return r, mAE
}

func AveragexGFCalc(status float64, r regression.Regression, tD *TeamData, size int) float64 {
	predictions := make([]float64, 0, len(tD.Status))
	for i := len(tD.GoalFor) - 1; len(predictions) <= size; i-- {
		if status == tD.Status[i] {
			GFPredicted, _ := r.Predict([]float64{tD.Status[i], tD.GoalAgainst[i]})
			predictions = append(predictions, GFPredicted)
		}
	}
	return stat.Mean(predictions, nil)
}
