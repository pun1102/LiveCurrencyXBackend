package main

import (
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/dynamodb"
    "github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-lambda-go/lambda"
	"fmt"
)
type Item struct {
    CurrencySymbol  string
    Rate float64
    High float64
    Low float64
    Volume float64
    Change float64
    
}

func Handle(item Item)(string,error) {
    sess := session.Must(session.NewSessionWithOptions(session.Options{
        SharedConfigState: session.SharedConfigEnable,
    }))

    // Create DynamoDB client
    svc := dynamodb.New(sess)
    tableName := "CurrencyLayerFXTable"
    curr := "USDT"

    result, err := svc.GetItem(&dynamodb.GetItemInput{
        TableName: aws.String(tableName),
        Key: map[string]*dynamodb.AttributeValue{
            "CurrencySymbol": {
                S: aws.String(curr),
            },
        },
    })
    if err != nil {
        return fmt.Sprintf(err.Error()),nil
    }
    if result.Item == nil {
        msg := "Could not find"
        return fmt.Sprintf(msg),nil
    }
        
    item = Item{}

    err = dynamodbattribute.UnmarshalMap(result.Item, &item)
    if err != nil {
        panic(fmt.Sprintf("Failed to unmarshal Record, %v", err))
    }
    
    
    return fmt.Sprintf("[{ high: %f, low: %f, rate: %f, vol: %f, change: %f}]", item.High, item.Low, item.Rate, item.Volume, item.Change),nil
    // snippet-end:[dynamodb.go.read_item.unmarshall]
}
func main(){
	lambda.Start(Handle)
}
// "[{ high: %f, low: %f, rate: %f, vol: %f}]", Item.High, Item.Low, Item.Rate, Item.Volume