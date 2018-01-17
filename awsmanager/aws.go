package awsmanager

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
)

type AwsConfig struct {
	AwsId     string
	AwsSecret string
	AwsToken  string
	AwsRegion string
}

type Zone struct {
	Id string
}

func (z *AwsConfig) SessionInit() (err error) {
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(z.AwsRegion),
		Credentials: credentials.NewStaticCredentials(z.AwsId, z.AwsSecret, z.AwsToken),
	})
	if err != nil {
		return err
	}
	fmt.Println(sess)
	return nil
}

func (a *Zone) AwsmanagerInit() {
	fmt.Println("vim-go")
}
