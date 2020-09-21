package main

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	// "io/ioutil"
	// "log"
	"net/http"
	// "os"
)

// snippet-end:[dynamodb.go.create_item.struct]
// A Response struct to map the Entire Response
type Response struct {
	Success   string `json:"success"`
	Terms     string `json:"terms"`
	Privacy   string `json:"privacy"`
	Timestamp int    `json:"timestamp"`
	Target    string `json:"target"`
	Rates     Rates  `json:"rates"`
}
type Rates struct {
	// EntryNo int            `json:"entry_number"`
	ABC  ABC  `json:"ABC"`
	AGI  AGI  `json:"AGI"`
	ETH  ETH  `json:"ETH"`
	BTC  BTC  `json:"BTC"`
	USDT USDT `json:"USDT"`
	BCH  BCH  `json:"BCH"`
}

// A struct to map our Pokemon's Species which includes it's name
type ABC struct {
	Rate   float64 `json:"rate"`
	High   float64 `json:"high"`
	Low    float64 `json:"low"`
	Volume float64 `json:"vol"`
	Change float64 `json:"change"`
}
type AGI struct {
	Rate   float64 `json:"rate"`
	High   float64 `json:"high"`
	Low    float64 `json:"low"`
	Volume float64 `json:"vol"`
	Change float64 `json:"change"`
}
type ETH struct {
	Rate   float64 `json:"rate"`
	High   float64 `json:"high"`
	Low    float64 `json:"low"`
	Volume float64 `json:"vol"`
	Change float64 `json:"change"`
}
type BTC struct {
	Rate   float64 `json:"rate"`
	High   float64 `json:"high"`
	Low    float64 `json:"low"`
	Volume float64 `json:"vol"`
	Change float64 `json:"change"`
}
type USDT struct {
	Rate   float64 `json:"rate"`
	High   float64 `json:"high"`
	Low    float64 `json:"low"`
	Volume float64 `json:"vol"`
	Change float64 `json:"change"`
}
type BCH struct {
	Rate   float64 `json:"rate"`
	High   float64 `json:"high"`
	Low    float64 `json:"low"`
	Volume float64 `json:"vol"`
	Change float64 `json:"change"`
}

func main() {

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// Create DynamoDB client
	svc := dynamodb.New(sess)

	//code from previous file
	response, err := http.Get("http://api.coinlayer.com/live?access_key=10027564b8637892ec3cf5dccaa1c812&target=INR&expand=1")
	if err != nil {
		// fmt.Printf("%#v\n", Response{})

		// return , err
	}
	// fmt.Printf("%#v\n", response)

	var ip_response Response
	defer response.Body.Close()
	err = json.NewDecoder(response.Body).Decode(&ip_response)
	if err != nil {
		// fmt.Printf("%#v\n", Response{})
	}
	// body, err := ioutil.ReadAll(response.Body)
	// fmt.Printf("%#v\n", string(body))
	// fmt.Printf("%#v\n", ip_response)
	fmt.Printf("%#v\n", ip_response.Rates.ABC.High)

	// response, err := http.Get("http://api.coinlayer.com/live?access_key=10027564b8637892ec3cf5dccaa1c812&target=INR&expand=1")

	// if err != nil {
	// 	fmt.Print(err.Error())
	// 	os.Exit(1)
	// }

	// responseData, err := ioutil.ReadAll(response.Body)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// // fmt.Println(string(responseData))

	// var responseObject Response
	// json.Unmarshal([]byte(responseData), &responseObject)

	// fmt.Println(responseObject)
	// fmt.Println(len(responseObject.Rates))

	// fmt.Println(string(responseData))
	tableName := "CurrencyLayerFXTable"

	p := fmt.Sprintf("%f", ip_response.Rates.ABC.Rate)
	q := fmt.Sprintf("%f", ip_response.Rates.ABC.High)
	r := fmt.Sprintf("%f", ip_response.Rates.ABC.Low)
	s := fmt.Sprintf("%f", ip_response.Rates.ABC.Volume)
	t := fmt.Sprintf("%f", ip_response.Rates.ABC.Change)
	currencyname := "ABC"
	currvalue1 := p
	currvalue2 := q
	currvalue3 := r
	currvalue4 := s
	currvalue5 := t

	input := &dynamodb.UpdateItemInput{
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":a": {
				N: aws.String(currvalue1),
			},
			":b": {
				N: aws.String(currvalue2),
			},
			":c": {
				N: aws.String(currvalue3),
			},
			":d": {
				N: aws.String(currvalue4),
			},
			":e": {
				N: aws.String(currvalue5),
			},
		},
		TableName: aws.String(tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"CurrencySymbol": {
				S: aws.String(currencyname),
			},
		},
		ReturnValues:     aws.String("UPDATED_NEW"),
		UpdateExpression: aws.String("set Rate = :a, High = :b, Low = :c, Volume = :d, Change= :e"),
	}

	svc.UpdateItem(input)

	Agirate := fmt.Sprintf("%f", ip_response.Rates.AGI.Rate)
	Agihigh := fmt.Sprintf("%f", ip_response.Rates.AGI.High)
	Agilow := fmt.Sprintf("%f", ip_response.Rates.AGI.Low)
	Agivol := fmt.Sprintf("%f", ip_response.Rates.AGI.Volume)
	Agichange := fmt.Sprintf("%f", ip_response.Rates.AGI.Change)
	AGI := "AGI"
	AGIRate := Agirate
	AGIHigh := Agihigh
	AGILow := Agilow
	AGIVol := Agivol
	AGIChange := Agichange

	AGIinput := &dynamodb.UpdateItemInput{
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":a": {
				N: aws.String(AGIRate),
			},
			":b": {
				N: aws.String(AGIHigh),
			},
			":c": {
				N: aws.String(AGILow),
			},
			":d": {
				N: aws.String(AGIVol),
			},
			":e": {
				N: aws.String(AGIChange),
			},
		},
		TableName: aws.String(tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"CurrencySymbol": {
				S: aws.String(AGI),
			},
		},
		ReturnValues:     aws.String("UPDATED_NEW"),
		UpdateExpression: aws.String("set Rate = :a, High = :b, Low = :c, Volume = :d, Change =:e"),
	}

	svc.UpdateItem(AGIinput)

	//Etherum

	ETHrate := fmt.Sprintf("%f", ip_response.Rates.ETH.Rate)
	ETHhigh := fmt.Sprintf("%f", ip_response.Rates.ETH.High)
	ETHlow := fmt.Sprintf("%f", ip_response.Rates.ETH.Low)
	ETHvol := fmt.Sprintf("%f", ip_response.Rates.ETH.Volume)
	ETHchange := fmt.Sprintf("%f", ip_response.Rates.ETH.Change)
	ETH := "ETH"

	ETHinput := &dynamodb.UpdateItemInput{
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":a": {
				N: aws.String(ETHrate),
			},
			":b": {
				N: aws.String(ETHhigh),
			},
			":c": {
				N: aws.String(ETHlow),
			},
			":d": {
				N: aws.String(ETHvol),
			},
			":e": {
				N: aws.String(ETHchange),
			},
		},
		TableName: aws.String(tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"CurrencySymbol": {
				S: aws.String(ETH),
			},
		},
		ReturnValues:     aws.String("UPDATED_NEW"),
		UpdateExpression: aws.String("set Rate = :a, High = :b, Low = :c, Volume = :d, Change =:e"),
	}

	svc.UpdateItem(ETHinput)

	//BTC

	BTCrate := fmt.Sprintf("%f", ip_response.Rates.BTC.Rate)
	BTChigh := fmt.Sprintf("%f", ip_response.Rates.BTC.High)
	BTClow := fmt.Sprintf("%f", ip_response.Rates.BTC.Low)
	BTCvol := fmt.Sprintf("%f", ip_response.Rates.BTC.Volume)
	BTCchange := fmt.Sprintf("%f", ip_response.Rates.BTC.Change)
	BTC := "BTC"

	BTCinput := &dynamodb.UpdateItemInput{
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":a": {
				N: aws.String(BTCrate),
			},
			":b": {
				N: aws.String(BTChigh),
			},
			":c": {
				N: aws.String(BTClow),
			},
			":d": {
				N: aws.String(BTCvol),
			},
			":e": {
				N: aws.String(BTCchange),
			},
		},
		TableName: aws.String(tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"CurrencySymbol": {
				S: aws.String(BTC),
			},
		},
		ReturnValues:     aws.String("UPDATED_NEW"),
		UpdateExpression: aws.String("set Rate = :a, High = :b, Low = :c, Volume = :d, Change =:e"),
	}

	svc.UpdateItem(BTCinput)

	//BCH
	BCHrate := fmt.Sprintf("%f", ip_response.Rates.BCH.Rate)
	BCHhigh := fmt.Sprintf("%f", ip_response.Rates.BCH.High)
	BCHlow := fmt.Sprintf("%f", ip_response.Rates.BCH.Low)
	BCHvol := fmt.Sprintf("%f", ip_response.Rates.BCH.Volume)
	BCHchange := fmt.Sprintf("%f", ip_response.Rates.BCH.Change)
	BCH := "BCH"

	BCHinput := &dynamodb.UpdateItemInput{
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":a": {
				N: aws.String(BCHrate),
			},
			":b": {
				N: aws.String(BCHhigh),
			},
			":c": {
				N: aws.String(BCHlow),
			},
			":d": {
				N: aws.String(BCHvol),
			},
			":e": {
				N: aws.String(BCHchange),
			},
		},
		TableName: aws.String(tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"CurrencySymbol": {
				S: aws.String(BCH),
			},
		},
		ReturnValues:     aws.String("UPDATED_NEW"),
		UpdateExpression: aws.String("set Rate = :a, High = :b, Low = :c, Volume = :d, Change = :e"),
	}

	svc.UpdateItem(BCHinput)

	//USDT
	USDTrate := fmt.Sprintf("%f", ip_response.Rates.USDT.Rate)
	USDThigh := fmt.Sprintf("%f", ip_response.Rates.USDT.High)
	USDTlow := fmt.Sprintf("%f", ip_response.Rates.USDT.Low)
	USDTvol := fmt.Sprintf("%f", ip_response.Rates.USDT.Volume)
	USDTchange := fmt.Sprintf("%f", ip_response.Rates.USDT.Change)
	USDT := "USDT"

	USDTinput := &dynamodb.UpdateItemInput{
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":a": {
				N: aws.String(USDTrate),
			},
			":b": {
				N: aws.String(USDThigh),
			},
			":c": {
				N: aws.String(USDTlow),
			},
			":d": {
				N: aws.String(USDTvol),
			},
			":e": {
				N: aws.String(USDTchange),
			},
		},
		TableName: aws.String(tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"CurrencySymbol": {
				S: aws.String(USDT),
			},
		},
		ReturnValues:     aws.String("UPDATED_NEW"),
		UpdateExpression: aws.String("set Rate = :a, High = :b, Low = :c, Volume = :d, Change = :e"),
	}

	svc.UpdateItem(USDTinput)

}
