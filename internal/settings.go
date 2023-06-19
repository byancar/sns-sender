package settings

import (
	"github.com/Netflix/go-env"
)

var Env settings

type settings struct {
	AWS struct {
		AccessKeyID     string `env:"AWS_ACCESS_KEY_ID,default=FAKE"`
		SecretAccessKey string `env:"AWS_SECRET_ACCESS_KEY,default=FAKE"`
		Region          string `env:"AWS_REGION,default=us-east-1"`
		DefaultRegion   string `env:"AWS_DEFAULT_REGION,default=us-east-1"`
		TopicName       string `env:"AWS_TOPIC_NAME,sns_sender_topic"`
		QueueName       string `env:"AWS_TOPIC_NAME,sns_sender_queue"`
	}
}

func init() {
	if _, err := env.UnmarshalFromEnviron(&Env); err != nil {
		panic(err)
	}
}
