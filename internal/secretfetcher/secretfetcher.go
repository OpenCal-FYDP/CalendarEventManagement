package secretfetcher

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
)

func fetchOauthSecret() (secString string, secretBinary string, err error) {
	secretName := "ServerSideOauthClient"
	region := "us-east-1"

	//Create a Secrets Manager client
	sess, err := session.NewSession()
	if err != nil {
		// Handle session creation error
		fmt.Println(err.Error())
		return
	}
	svc := secretsmanager.New(sess,
		aws.NewConfig().WithRegion(region))
	input := &secretsmanager.GetSecretValueInput{
		SecretId:     aws.String(secretName),
		VersionStage: aws.String("AWSCURRENT"), // VersionStage defaults to AWSCURRENT if unspecified
	}

	result, err := svc.GetSecretValue(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case secretsmanager.ErrCodeDecryptionFailure:
				// Secrets Manager can't decrypt the protected secret text using the provided KMS key.
				return "", "", fmt.Errorf("%s:%s", secretsmanager.ErrCodeDecryptionFailure, aerr.Error())

			case secretsmanager.ErrCodeInternalServiceError:
				// An error occurred on the server side.
				return "", "", fmt.Errorf("%s:%s", secretsmanager.ErrCodeInternalServiceError, aerr.Error())

			case secretsmanager.ErrCodeInvalidParameterException:
				// You provided an invalid value for a parameter.
				return "", "", fmt.Errorf("%s:%s", secretsmanager.ErrCodeInvalidParameterException, aerr.Error())

			case secretsmanager.ErrCodeInvalidRequestException:
				// You provided a parameter value that is not valid for the current state of the resource.
				return "", "", fmt.Errorf("%s:%s", secretsmanager.ErrCodeInvalidRequestException, aerr.Error())

			case secretsmanager.ErrCodeResourceNotFoundException:
				// We can't find the resource that you asked for.
				return "", "", fmt.Errorf("%s:%s", secretsmanager.ErrCodeResourceNotFoundException, aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			return "", "", err
		}
		return
	}

	// Decrypts secret using the associated KMS key.
	// Depending on whether the secret is a string or binary, one of these fields will be populated.
	var secretString, decodedBinarySecret string
	if result.SecretString != nil {
		secretString = *result.SecretString
	} else {
		decodedBinarySecretBytes := make([]byte, base64.StdEncoding.DecodedLen(len(result.SecretBinary)))
		l, err := base64.StdEncoding.Decode(decodedBinarySecretBytes, result.SecretBinary)
		if err != nil {
			return "", "", err
		}
		decodedBinarySecret = string(decodedBinarySecretBytes[:l])
	}

	return secretString, decodedBinarySecret, nil
}

type OauthClientSecret struct {
	OAuthClientID string
	ClientSecret  string
}

func parseSecretString(secretString string) (*OauthClientSecret, error) {
	// Declared an empty map interface
	var result map[string]string
	// Unmarshal or Decode the JSON to the interface.
	err := json.Unmarshal([]byte(secretString), &result)
	if err != nil {
		return nil, err
	}

	res := &OauthClientSecret{
		OAuthClientID: result["OAuthClientID"],
		ClientSecret:  result["ClientSecret"],
	}

	return res, nil
}

func GetOauthConfig() (*OauthClientSecret, error) {
	secString, _, err := fetchOauthSecret()
	if err != nil {
		return nil, err
	}
	return parseSecretString(secString)
}
