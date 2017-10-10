package providers

import (
	"os"

	"github.com/tonyhhyip/go-di-container"
	"github.com/ysitd-cloud/judge-go-client"
)

type judgeServiceProvider struct {
	*container.AbstractServiceProvider
}

func (*judgeServiceProvider) Provides() []string {
	return []string{
		"judge.client",
		"judge.endpoint",
		"judge.subject",
		"judge.token",
	}
}

func (*judgeServiceProvider) Register(app container.Container) {
	app.Instance("judge.subject", os.Getenv("JUDGE_SUBJECT"))
	app.Instance("judge.instance", os.Getenv("JUDGE_ENDPOINT"))
	app.Instance("judge.token", os.Getenv("JUDGE_TOKEN"))

	app.Bind("judge.client", func(app container.Container) interface{} {
		return judge.NewClientv1(
			app.Make("judge.endpoint").(string),
			app.Make("judge.subject").(string),
			judge.SUBJECT_APP,
			app.Make("judge.token").(string),
		)
	})
}
