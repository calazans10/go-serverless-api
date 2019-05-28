package api

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/gofrs/uuid"
)

var (
	awsRegion = os.Getenv("REGION")
	tableName = os.Getenv("DYNAMODB_TABLE")
	db        = dynamodb.New(session.New(), aws.NewConfig().WithRegion(awsRegion))
)

// GetUsers retrieves all the users from the DB
func GetUsers() ([]User, error) {
	input := &dynamodb.ScanInput{
		TableName: aws.String(tableName),
	}

	result, err := db.Scan(input)
	if err != nil {
		return []User{}, err
	}

	if len(result.Items) == 0 {
		return []User{}, err
	}

	var users []User
	err = dynamodbattribute.UnmarshalListOfMaps(result.Items, &users)
	if err != nil {
		return []User{}, err
	}

	return users, nil
}

// CreateUser inserts a new User item to the table.
func CreateUser(user User) error {
	uuid, err := uuid.NewV4()
	if err != nil {
		return err
	}
	user.ID = uuid.String()

	av, err := dynamodbattribute.MarshalMap(user)
	if err != nil {
		return err
	}

	input := &dynamodb.PutItemInput{
		TableName: aws.String(tableName),
		Item:      av,
	}

	_, err = db.PutItem(input)
	return err
}
