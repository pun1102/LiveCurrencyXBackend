package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"os"
)

func main() {

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	svc := dynamodb.New(sess)

	table := "CurrencyLayerFXTable"

	input := &dynamodb.CreateTableInput{
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String("CurrencySymbol"),
				KeyType:       aws.String("HASH"),
			},
		},
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String("CurrencySymbol"),
				AttributeType: aws.String("S"),
			},
		},
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			WriteCapacityUnits: aws.Int64(50),
			ReadCapacityUnits:  aws.Int64(50),
		},

		TableName: aws.String(table),
	}

	_, err := svc.CreateTable(input)
	if err != nil {
		fmt.Println("Error while creating a table:")
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println("%s succesfully created", table)
}
