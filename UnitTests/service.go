package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const JokeApiEndpoint string = "https://v2.jokeapi.dev/joke/Any?blacklistFlags=nsfw,religious,political,racist,sexist,explicit"

type JokeResponse struct {
	Err      bool            `json:"error"`
	Category string          `json:"category"`
	JokeType string          `json:"type"`
	Flags    map[string]bool `json:"flags"`
	Safe     bool            `json:"safe"`
	Id       float64         `json:"id"`
	Lang     string          `json:"lang"`
	Joke     string          `json:"joke"`
	Setup    string          `json:"setup"`
	Delivery string          `json:"delivery"`
}

func GetJoke() JokeResponse {
	response, err := http.Get(JokeApiEndpoint)
	if err != nil {
		return JokeResponse{}
	}

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		return JokeResponse{}
	}

	var responseObject JokeResponse
	json.Unmarshal(responseData, &responseObject)

	if responseObject.JokeType == "single" {
		fmt.Println(responseObject.Joke)
	} else {
		fmt.Println(responseObject.Setup)
		fmt.Println(responseObject.Delivery)
	}

	return responseObject
}
