package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/dghubble/oauth1"
	"github.com/joho/godotenv"
)

func constructClient() (client *http.Client) {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf(".env File not found")
	}

	config := oauth1.NewConfig(os.Getenv("API_KEY"), os.Getenv("API_KEY_SECRET"))
	token := oauth1.NewToken(os.Getenv("ACCESS_TOKEN"), os.Getenv("ACCESS_TOKEN_SECRET"))

	return config.Client(oauth1.NoContext, token)
}

func uploadMedia(filename string) map[string]interface{} {
	file, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	fileEncoded := base64.StdEncoding.EncodeToString(file)

	client := constructClient()
	resp, err := client.PostForm("https://upload.twitter.com/1.1/media/upload.json", url.Values{
		"media_data": {fileEncoded},
	})

	if err != nil {
		log.Fatal(err)
	}

	var res map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&res)

	return res
}

func sendTweet(text string, mediaIds []string) map[string]interface{} {
	var values map[string]interface{}

	if len(mediaIds) > 0 {
		media := map[string]interface{}{
			"media_ids": mediaIds,
		}
		
		values = map[string]interface{}{
			"text":  text,
			"media": media,
		}
	} else {
		values = map[string]interface{}{
			"text": text,
		}
	}

	json_data, err := json.Marshal(values)

	fmt.Println(string(json_data))

	if err != nil {
		log.Fatal(err)
	}

	client := constructClient()
	resp, err := client.Post("https://api.twitter.com/2/tweets", "application/json", bytes.NewBuffer(json_data))

	if err != nil {
		log.Fatal(err)
	}

	var res map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&res)

	return res
}
