package setup

import (
	"os"

	"github.com/ysitd-cloud/judge-go-client"
)

func SetupJudgeClient() judge.Client {
	return judge.NewClientv1(
		os.Getenv("JUDGE_ENDPOINT"),
		os.Getenv("JUDGE_SUBJECT"),
		judge.SUBJECT_APP,
		os.Getenv("JUDGE_TOKEN"),
	)
}
