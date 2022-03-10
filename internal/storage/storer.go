package storage

import (
	"github.com/OpenCal-FYDP/CalendarEventManagement/rpc"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

type Storage struct {
	client dynamodbiface.DynamoDBAPI
}

func New() *Storage {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	client := dynamodb.New(sess)

	return &Storage{client}
}

func (s *Storage) CreateEvent(req *rpc.CreateEventReq) error {
	av, err := dynamodbattribute.MarshalMap(req)

	if err != nil {
		return err
	}

	// Now create put item
	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(userTable),
	}

	// Now we push it into the database, then return
	_, err = s.client.PutItem(input)

	return err
}

func (s *Storage) UpdateEvent(req *rpc.UpdateEventReq) error {
	panic("a")
}

func (s *Storage) DeleteEvent(req *rpc.DeleteEventReq) error {
	panic("a")
}
