# LiveCurrencyXBackend

1.I have created a web page which will display the live stock of different cryptocurrencies i.e, it’s rates, highest value, lowest value, volume and delta (change in the value since w.r.t to last updated time). 
2.The API which is used in this project is an open source API by Coinlayer​. and typically gives hourly updates. For demonstration purposes I have taken six crypto-currencies. 
3.Created a table named as “CurrencyLayerFXTable” with primary key as “CurrencySymbol” using Go Lang and AWS-GO-SDK.  
4.A separate GoLang script to insert each of the corresponding values to the currencies into DynamoDb.  
5.A lambda function to trigger the insert script after an interval of 60 mins. For this purpose I have added an Event Bridge trigger to this lambda function. 
6.Also, for each currency I have written a separate lambda function  which will  read the value from DynamoDb and will return a response. All the six lambda functions are deployed with individual APIs using the GET method. 
7.For example: the api returning the response for Ethereum Currency is : https://zvw0lxf1ve.execute-api.us-east-2.amazonaws.com/eth 
8.On the frontend side, I have written six methods to fetch responses of each api using the ‘axios’ package and called it inside the  DidMount() function along with the interval of 60 minutes in order to recall the API and the same interval. 
9.Atlast, Deployed it on AWS Amplify using GIt repository. To check it out please click on this link.  [ https://master.d3vbzekfitg8lw.amplifyapp.com/ ] 
10.I have pushed the complete code base on Github, Attaching link here please go through it once.  a.Backend Codebase: https://github.com/pun1102/LiveCurrencyXBackend b.Frontend Codebase: https://github.com/pun1102/VXExchange 
