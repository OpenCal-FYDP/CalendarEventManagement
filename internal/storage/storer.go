package storage

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

type EventData struct {
	CalendarEventID string
	//TimeToNotify string
	Start     int64
	End       int64
	Attendees []string
	Location  string
	Summary   string
}

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

func (s *Storage) CreateEvent(event *EventData) error {
	return s.UpdateEvent(event)
}

func (s *Storage) UpdateEvent(event *EventData) error {
	av, err := dynamodbattribute.MarshalMap(event)

	if err != nil {
		return err
	}

	// Now create put item
	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String("calEvents"),
	}

	_, err = s.client.PutItem(input)

	return err
}

func (s *Storage) GetEvent(calendarEventID string) (*EventData, error) {
	//e := &EventData{
	//	CalendarEventID: calendarEventID,
	//}
	//
	//av, err := dynamodbattribute.MarshalMap(e)
	//if err != nil {
	//	return nil, err
	//}

	// Now create put item
	input := &dynamodb.GetItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"CalendarEventID": {
				S: aws.String(calendarEventID),
			},
		},
		TableName: aws.String("calEvents"),
	}

	res, err := s.client.GetItem(input)
	if err != nil {
		return nil, err
	}

	retEvent := &EventData{}
	err = dynamodbattribute.UnmarshalMap(res.Item, retEvent)
	if err != nil {
		return nil, err
	}

	return retEvent, nil
}

func (s *Storage) DeleteEvent(calendarEventID string) error {
	input := &dynamodb.DeleteItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"CalendarEventID": {
				S: aws.String(calendarEventID),
			},
		},
		TableName: aws.String("calEvents"),
	}

	_, err := s.client.DeleteItem(input)

	return err
}
