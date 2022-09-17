package main

import (
	"fmt"
	"log"
	"os"
	"sync"
	access "github.com/UtkarshM-hub/go-twitter-bot/Access"
	"github.com/joho/godotenv"
)

var wg=sync.WaitGroup{}

func main(){
	godotenv.Load(".env");
	fmt.Println("Go-Twitter Bot v0.01")
    creds := access.Credentials{
        AccessToken:       os.Getenv("ACCESS_TOKEN"),
        AccessTokenSecret: os.Getenv("ACCESS_TOKEN_SECRET"),
        ConsumerKey:       os.Getenv("CONSUMER_KEY"),
        ConsumerSecret:    os.Getenv("CONSUMER_SECRET"),
    }

    fmt.Printf("%+v\n", creds)

    client, err := access.GetClient(&creds)
    if err != nil {
        log.Println("Error getting Twitter Client")
        log.Println(err)
    }

	wg.Add(1)
    go client.Statuses.Update("Test",nil)
	wg.Wait()
}