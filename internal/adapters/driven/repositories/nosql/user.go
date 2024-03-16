package reponosql

import (
	"fiap-hf-auth/internal/core/db"
	"fiap-hf-auth/internal/core/domain/entity/dto"
	"fiap-hf-auth/internal/core/domain/repository"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

var _ repository.UserRepositoryNoSQL = (*userDB)(nil)

type userDB struct {
	Database  db.NoSQLDatabase
	tableName string
}

func NewUserDynamoDB(database db.NoSQLDatabase, tableName string) *userDB {
	return &userDB{Database: database, tableName: tableName}
}

func (c *userDB) GetUserByCPF(cpf string) (*dto.UserNoSQL, error) {
	filter := "cpf = :value"
	attrSearch := map[string]*dynamodb.AttributeValue{
		":value": {
			S: aws.String(cpf),
		},
	}

	input := &dynamodb.ScanInput{
		TableName:                 aws.String(c.tableName),
		FilterExpression:          aws.String(filter),
		ExpressionAttributeValues: attrSearch,
	}

	result, err := c.Database.Scan(input)
	if err != nil {
		return nil, err
	}

	var userList = make([]dto.UserNoSQL, 0)
	for _, item := range result.Items {
		var c dto.UserNoSQL
		if err := dynamodbattribute.UnmarshalMap(item, &c); err != nil {
			return nil, err
		}
		userList = append(userList, c)
	}

	if len(userList) > 0 {
		return &userList[0], nil
	}

	return nil, nil
}

func (c *userDB) GetUserByEmail(email string) (*dto.UserNoSQL, error) {

	filter := "email = :value"
	attrSearch := map[string]*dynamodb.AttributeValue{
		":value": {
			S: aws.String(email),
		},
	}

	input := &dynamodb.ScanInput{
		TableName:                 aws.String(c.tableName),
		FilterExpression:          aws.String(filter),
		ExpressionAttributeValues: attrSearch,
	}

	result, err := c.Database.Scan(input)
	if err != nil {
		return nil, err
	}

	var userList = make([]dto.UserNoSQL, 0)
	for _, item := range result.Items {
		var c dto.UserNoSQL
		if err := dynamodbattribute.UnmarshalMap(item, &c); err != nil {
			return nil, err
		}
		userList = append(userList, c)
	}

	if len(userList) > 0 {
		return &userList[0], nil
	}

	return nil, nil
}

func (c *userDB) SaveUser(user dto.UserNoSQL) (*dto.UserNoSQL, error) {

	putItem := map[string]*dynamodb.AttributeValue{
		"uuid": {
			S: aws.String(user.UUID),
		},
		"name": {
			S: aws.String(user.Name),
		},
		"username": {
			S: aws.String(user.Username),
		},
		"password": {
			S: aws.String(user.Password),
		},
		"cpf": {
			S: aws.String(user.CPF),
		},
		"email": {
			S: aws.String(user.Email),
		},
		"created_at": {
			S: aws.String(user.CreatedAt),
		},
	}

	inputPutItem := &dynamodb.PutItemInput{
		Item:      putItem,
		TableName: aws.String(c.tableName),
	}

	putOut, err := c.Database.PutItem(inputPutItem)

	if err != nil {
		return nil, err
	}

	var out *dto.UserNoSQL

	if err := dynamodbattribute.UnmarshalMap(putOut.Attributes, &out); err != nil {
		return nil, err
	}

	return out, nil
}
