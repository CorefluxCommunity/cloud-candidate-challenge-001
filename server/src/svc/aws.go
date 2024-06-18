package svc

import "os"

type AwsEnv struct {
	Region          string
	AccessKeyID     string
	SecretAccessKey string
	SessionToken    string
}

func NewAwsEnv() AwsEnv {
	return AwsEnv{
		Region:          os.Getenv("AWS_REGION"),
		AccessKeyID:     os.Getenv("AWS_ACCESS_KEY_ID"),
		SecretAccessKey: os.Getenv("AWS_SECRET_ACCESS_KEY"),
		SessionToken:    os.Getenv("AWS_SESSION_TOKEN"),
	}
}
