package awsrds

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/rds"
)

func getAwsClient(accessKeyId, secretAccessKey, region string) rds.RDS {
	awsConfig := &aws.Config{
		Credentials: credentials.NewStaticCredentials(accessKeyId, secretAccessKey, ""),

		Region: aws.String(region),
	}

	newSession := session.New(awsConfig)

	rdsClient := rds.New(newSession)

	return *rdsClient
}

func StartRdsInstance(accessKeyId, secretAccessKey, region string, startDbInstnceRequest *rds.StartDBInstanceInput) (int, string) {
	awsClient := getAwsClient(accessKeyId, secretAccessKey, region)

	result, err := awsClient.StartDBInstance(startDbInstnceRequest)

	if err != nil {

		if aerr, ok := err.(awserr.RequestFailure); ok {
			fmt.Println(aerr.Error())
			return aerr.StatusCode(), aerr.Error()

		} else if aerr, ok := err.(awserr.Error); ok {
			fmt.Println(aerr.Error())
			return 106, aerr.Error()

		}

	}
	resp, _ := json.Marshal(result)
	return 200, string(resp)

}

func StopRdsInstance(accessKeyId, secretAccessKey, region string, stopDbInstnceRequest *rds.StopDBInstanceInput) (int, string) {
	awsClient := getAwsClient(accessKeyId, secretAccessKey, region)

	result, err := awsClient.StopDBInstance(stopDbInstnceRequest)

	if err != nil {

		if aerr, ok := err.(awserr.RequestFailure); ok {
			fmt.Println(aerr.Error())
			return aerr.StatusCode(), aerr.Error()

		} else if aerr, ok := err.(awserr.Error); ok {
			fmt.Println(aerr.Error())
			return 106, aerr.Error()

		}

	}
	resp, _ := json.Marshal(result)
	return 200, string(resp)

}

func RebootRdsInstance(accessKeyId, secretAccessKey, region string, rebootDbInstnceRequest *rds.RebootDBInstanceInput) (int, string) {
	awsClient := getAwsClient(accessKeyId, secretAccessKey, region)

	result, err := awsClient.RebootDBInstance(rebootDbInstnceRequest)

	if err != nil {

		if aerr, ok := err.(awserr.RequestFailure); ok {
			fmt.Println(aerr.Error())
			return aerr.StatusCode(), aerr.Error()

		} else if aerr, ok := err.(awserr.Error); ok {
			fmt.Println(aerr.Error())
			return 106, aerr.Error()

		}

	}
	resp, _ := json.Marshal(result)
	return 200, string(resp)

}

func DeleteRdsInstance(accessKeyId, secretAccessKey, region string, deleteDbInstnceRequest *rds.DeleteDBInstanceInput) (int, string) {
	awsClient := getAwsClient(accessKeyId, secretAccessKey, region)

	result, err := awsClient.DeleteDBInstance(deleteDbInstnceRequest)

	if err != nil {

		if aerr, ok := err.(awserr.RequestFailure); ok {
			fmt.Println(aerr.Error())
			return aerr.StatusCode(), aerr.Error()

		} else if aerr, ok := err.(awserr.Error); ok {
			fmt.Println(aerr.Error())
			return 106, aerr.Error()

		}

	}
	resp, _ := json.Marshal(result)
	return 200, string(resp)

}

func CreateRdsInstance(accessKeyId, secretAccessKey, region string, createDbInstnceRequest *rds.CreateDBInstanceInput) (int, string) {

	awsClient := getAwsClient(accessKeyId, secretAccessKey, region)

	result, err := awsClient.CreateDBInstance(createDbInstnceRequest)

	if err != nil {

		if aerr, ok := err.(awserr.RequestFailure); ok {
			fmt.Println(aerr.Error())
			return aerr.StatusCode(), aerr.Error()

		} else if aerr, ok := err.(awserr.Error); ok {
			fmt.Println(aerr.Error())
			return 106, aerr.Error()

		}

	}
	resp, _ := json.Marshal(result)
	return 200, string(resp)

}
