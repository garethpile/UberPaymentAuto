package main

import (
	"context"
	"fmt"
	"os"
	"time"
	"strings"

	"io/ioutil"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"

)

func main() {
	lambda.Start(handler)
}

func handler(ctx context.Context, s3Event events.S3Event) {
	fmt.Printf("Starting handler \n")
	
	
	for _, record := range s3Event.Records {
		s3 := record.S3
		fmt.Printf("[%s - %s] Bucket = %s, Key = %s \n", record.EventSource, record.EventTime, s3.Bucket.Name, s3.Object.Key)
		fileContent := getDataFromS3File(s3.Bucket.Name, s3.Object.Key)
		dataExtracted := extractData(fileContent)
		insertIntoDynamoDB(dataExtracted, s3.Object.Key)
		fmt.Printf("Finished handler")
	}

}

func getDataFromS3File(bucket string, s3File string) string {
	fmt.Printf("Starting getDataFromS#File\n")
	//the only writable directory in the lambda is /tmp
	file, err := os.Create("/tmp/" + s3File)
	if err != nil {
		exitErrorf("Unable to open file %q, %v", s3File, err)
	}

	defer file.Close()

	// replace with your bucket region
	sess, _ := session.NewSession(&aws.Config{
		Region: aws.String("eu-west-1")},
	)

	downloader := s3manager.NewDownloader(sess)

	_, err = downloader.Download(file,
		&s3.GetObjectInput{
			Bucket: aws.String(bucket),
			Key:    aws.String(s3File),
		})
	if err != nil {
		exitErrorf("Unable to download s3File %q, %v", s3File, err)
	}

	dat, err := ioutil.ReadFile(file.Name())

	if err != nil {
		exitErrorf("Cannot read the file", err)
	}

	fmt.Printf("Finished getDataFromFile\n")

	return string(dat)

}

func extractData(data string) []string {

	fmt.Printf("Starting extractData\n")

	lines := strings.Split(data, "\n")

	var csvlines []string

	for _, currentLine := range lines {
	
		csvrow := strings.Split(currentLine,"\n")[0]
		csvlines = append(csvlines , csvrow)
	
	}

	fmt.Printf("Finished extractData\n")

	return csvlines
}

func insertIntoDynamoDB(dataToInsert []string, fileName string) {

	fmt.Printf("Starting insertIntoDynamoDB Round 2000\n")


	type MyDataFromS3 struct {
		Id string
		UberDriverId string
		FirstName    string
		LastName    string
		Total    string
	}

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	svc := dynamodb.New(sess)

	for currentIndex, currentItem := range dataToInsert {
        
		guid := time.Now()
		sguid := guid.String()

		stringcurrentIndex := string(currentIndex)
		fmt.Printf(stringcurrentIndex) 

		uberDriverId := strings.Split(currentItem,",")[0]
		firstName := strings.Split(currentItem,",")[1]
		lastName := strings.Split(currentItem,",")[2]
		total := strings.Split(currentItem,",")[3]

		uberDriverId = uberDriverId[1 : len(uberDriverId)-1]
		firstName = firstName[1 : len(firstName)-1]
		lastName = lastName[1 : len(lastName)-1]
		total = total[1 : len(total)-1]

		data := MyDataFromS3{

			Id: sguid,
			UberDriverId: uberDriverId,
			FirstName: firstName,
            LastName: lastName,
			Total: total,
		}

		av, err := dynamodbattribute.MarshalMap(data)
		if err != nil {
			exitErrorf("Got error marshalling new movie item:", av, err)
		}

		tableName := "UberPaymentTransactions"
		input := &dynamodb.PutItemInput{
			Item:      av,
			TableName: aws.String(tableName),
		}

		_, err = svc.PutItem(input)
		if err != nil {
			exitErrorf("Got error calling PutItem:", err)
		}
	}

	fmt.Printf("Finished insertIntoDynamoDB\n")
}

func exitErrorf(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, msg+"\n", args...)
	os.Exit(1)
}
