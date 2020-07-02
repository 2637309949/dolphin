package srv

import (
	pApp "github.com/2637309949/dolphin/platform/app"
	pModel "github.com/2637309949/dolphin/platform/model"
)

func init() {
	jobName := "hello"
	worker := pApp.App.Manager.Worker()
	worker.AddJobHandler(jobName, func(w pModel.Worker) (interface{}, error) {
		return map[string]interface{}{
			"mean":  78,
			"score": 87,
		}, nil
	})
}
