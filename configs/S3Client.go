package config

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

var Svc *s3.S3

func init() {
	Svc = s3.New(session.New(), aws.NewConfig().WithRegion("us-east-1"))
}
