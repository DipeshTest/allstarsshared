/*
package awsrds is a bussiness logic package for FLOGO AWS RDS connectors developed by AllStars team,
this package has the functions StartRdsInstance,StopRdsInstance,RebootRdsInstance,DeleteRdsInstance and CreateRdsInstance,
StartRdsInstance can be used to start specified AWS RDS DB Instance,
StopRdsInstance can be used to stop specified AWS RDS DB Instance,
RebootRdsInstance can be used to reboot specified AWS RDS DB Instance,
DeleteRdsInstance can be used to delete specified AWS RDS DB Instance,
CreateRdsInstance can be used to create specified new AWS RDS DB Instance with specified DB instance specification.
*/
package awsrds

import (
	"encoding/json"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/rds"
)

// Creates AWS Client using accessKeyId,secretAccessKey and region specified in the input in order to connect to AWS as IAM user
func getAwsClient(accessKeyId, secretAccessKey, region string) rds.RDS {
	awsConfig := &aws.Config{
		Credentials: credentials.NewStaticCredentials(accessKeyId, secretAccessKey, ""),

		Region: aws.String(region),
	}

	newSession := session.New(awsConfig)

	rdsClient := rds.New(newSession)

	return *rdsClient
}

//StartRdsInstance can be used to start specified AWS RDS DB Instance.
func StartRdsInstance(accessKeyId, secretAccessKey, region string, startDbInstnceRequest *rds.StartDBInstanceInput) (code int, messgae string) {
	awsClient := getAwsClient(accessKeyId, secretAccessKey, region)

	result, err := awsClient.StartDBInstance(startDbInstnceRequest)

	if err != nil {
		// Here we parse error response into the type RequestFailure first if success the corresponding error will be returned.
		// otherwise it will be parsed  into generic Error type
		// One parsing statement can be avoided thats why we have done the error parsing as follows.
		if aerr, ok := err.(awserr.RequestFailure); ok {
			return aerr.StatusCode(), aerr.Error()
		} else if aerr, ok := err.(awserr.Error); ok {
			return 106, aerr.Error()
		}

	}
	resp, _ := json.Marshal(result)
	return 200, string(resp)

}

//StopRdsInstance can be used to stop specified AWS RDS DB Instance.
func StopRdsInstance(accessKeyId, secretAccessKey, region string, stopDbInstnceRequest *rds.StopDBInstanceInput) (code int, message string) {
	awsClient := getAwsClient(accessKeyId, secretAccessKey, region)

	result, err := awsClient.StopDBInstance(stopDbInstnceRequest)

	if err != nil {
		// Here we parse error response into the type RequestFailure first if success the corresponding error will be returned.
		// otherwise it will be parsed  into generic Error type
		// One parsing statement can be avoided thats why we have done the error parsing as follows.
		if aerr, ok := err.(awserr.RequestFailure); ok {
			return aerr.StatusCode(), aerr.Error()
		} else if aerr, ok := err.(awserr.Error); ok {
			return 106, aerr.Error()
		}

	}
	resp, _ := json.Marshal(result)
	return 200, string(resp)

}

//RebootRdsInstance can be used to reboot specified AWS RDS DB Instance.
func RebootRdsInstance(accessKeyId, secretAccessKey, region string, rebootDbInstnceRequest *rds.RebootDBInstanceInput) (code int, message string) {
	awsClient := getAwsClient(accessKeyId, secretAccessKey, region)

	result, err := awsClient.RebootDBInstance(rebootDbInstnceRequest)

	if err != nil {
		// Here we parse error response into the type RequestFailure first if success the corresponding error will be returned.
		// otherwise it will be parsed  into generic Error type
		// One parsing statement can be avoided thats why we have done the error parsing as follows.
		if aerr, ok := err.(awserr.RequestFailure); ok {
			return aerr.StatusCode(), aerr.Error()
		} else if aerr, ok := err.(awserr.Error); ok {
			return 106, aerr.Error()
		}

	}
	resp, _ := json.Marshal(result)
	return 200, string(resp)

}

// DeleteRdsInstance can be used to delete specified AWS RDS DB Instance.
func DeleteRdsInstance(accessKeyId, secretAccessKey, region string, deleteDbInstnceRequest *rds.DeleteDBInstanceInput) (code int, message string) {
	awsClient := getAwsClient(accessKeyId, secretAccessKey, region)

	result, err := awsClient.DeleteDBInstance(deleteDbInstnceRequest)

	if err != nil {
		// Here we parse error response into the type RequestFailure first if success the corresponding error will be returned.
		// otherwise it will be parsed  into generic Error type
		// One parsing statement can be avoided thats why we have done the error parsing as follows.
		if aerr, ok := err.(awserr.RequestFailure); ok {
			return aerr.StatusCode(), aerr.Error()
		} else if aerr, ok := err.(awserr.Error); ok {
			return 106, aerr.Error()
		}

	}
	resp, _ := json.Marshal(result)
	return 200, string(resp)

}

// CreateRdsInstance can be used to create specified new AWS RDS DB Instance with specified DB instance specification.
func CreateRdsInstance(accessKeyId, secretAccessKey, region string, createDbInstnceRequest *rds.CreateDBInstanceInput) (code int, message string) {

	awsClient := getAwsClient(accessKeyId, secretAccessKey, region)

	result, err := awsClient.CreateDBInstance(createDbInstnceRequest)

	if err != nil {
		// Here we parse error response into the type RequestFailure first if success the corresponding error will be returned.
		// otherwise it will be parsed  into generic Error type
		// One parsing statement can be avoided thats why we have done the error parsing as follows.
		if aerr, ok := err.(awserr.RequestFailure); ok {
			return aerr.StatusCode(), aerr.Error()
		} else if aerr, ok := err.(awserr.Error); ok {
			return 106, aerr.Error()
		}

	}
	resp, _ := json.Marshal(result)
	return 200, string(resp)

}
