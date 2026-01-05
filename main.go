package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type Person struct {
	Name    string  `json:"name"`
	Email   string  `json:"email"`
	Phone   string  `json:"phone"`
	Balance float64 `json:"balance"`
}

func main() {

	bucket := "amazn-s3-demo-project-golang"
	key := "sample2.json"

	ctx := context.Background()

	cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion("us-east-2"))

	if err != nil {
		log.Fatalf("unable to load AWS config: %v", err)
	}

	client := s3.NewFromConfig(cfg)

	obj, err := client.GetObject(ctx, &s3.GetObjectInput{
		Bucket: &bucket,
		Key:    &key,
	})
	if err != nil {
		log.Fatalf("failed to get object: %v", err)
	}
	defer obj.Body.Close()

	data, err := io.ReadAll(obj.Body)
	if err != nil {
		log.Fatalf("failed to read object body: %v", err)
	}

	var people []Person

	if err := json.Unmarshal(data, &people); err != nil {
		log.Fatalf("JSON must be an array of Person objects: %v", err)
	}

	for _, p := range people {
		if p.Balance < 500 {
			fmt.Println("LOW BALANCE:", p.Name, p.Balance)
		} else {
			fmt.Println("OK:", p.Name, p.Balance)
		}
	}

}
