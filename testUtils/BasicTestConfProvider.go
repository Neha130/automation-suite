package testUtils

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/caarlos0/env"
	"github.com/go-resty/resty/v2"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

const createSessionApiUrl string = "/orchestrator/api/v1/session"

type LogInResult struct {
	Token string `json:"token"`
}

type LogInResponse struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Result LogInResult `json:"result"`
}

type ApiErrorDto struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Errors []struct {
		Code            string `json:"code"`
		InternalMessage string `json:"internalMessage"`
		UserMessage     string `json:"userMessage"`
	} `json:"errors"`
}

type EnvironmentConfig struct {
	BaseServerUrl   string `env:"BASE_SERVER_URL" envDefault:""`
	LogInUserName   string `env:"LOGIN_USERNAME" envDefault:""`
	LogInUserPwd    string `env:"LOGIN_PASSWORD" envDefault:""`
	SSOClientSecret string `env:"CLIENT_SECRET" envDefault:""`
}

func getRestyClient() *resty.Client {
	envConf, _ := GetEnvironmentConfig()
	client := resty.New()
	client.SetBaseURL(envConf.BaseServerUrl)
	return client
}

// MakeApiCall make the api call to the requested url based on http method requested
func MakeApiCall(apiUrl string, method string, body string, queryParams map[string]string, authToken string) (*resty.Response, error) {
	var resp *resty.Response
	var err error
	switch method {
	case "GET":
		if queryParams != nil {
			return getRestyClient().SetCookie(&http.Cookie{Name: "argocd.token", Value: authToken}).R().SetQueryParams(queryParams).Get(apiUrl)
		}
		return getRestyClient().SetCookie(&http.Cookie{Name: "argocd.token", Value: authToken}).R().Get(apiUrl)
	case "POST":
		return getRestyClient().SetCookie(&http.Cookie{Name: "argocd.token", Value: authToken}).R().SetBody(body).Post(apiUrl)
	case "PUT":
		return getRestyClient().SetCookie(&http.Cookie{Name: "argocd.token", Value: authToken}).R().SetBody(body).Put(apiUrl)
	case "DELETE":
		return getRestyClient().SetCookie(&http.Cookie{Name: "argocd.token", Value: authToken}).R().SetBody(body).Delete(apiUrl)
	}
	return resp, err
}

// HandleError Log the error and return boolean value indicating whether error occurred or not
func HandleError(err error, testName string) {
	if nil != err {
		log.Println("Error occurred while invoking api for test:"+testName, "err", err)
	}
}

func GetByteArrayOfGivenJsonFile(filePath string) ([]byte, error) {
	testDataJsonFile, err := os.Open(filePath)
	if nil != err {
		log.Println("Unable to open the file. Error occurred !!", "err", err)
	}
	log.Println("Opened the given json file successfully !!!")
	defer testDataJsonFile.Close()

	byteValue, err := ioutil.ReadAll(testDataJsonFile)
	return byteValue, err
}

//support function to return auth token after log in
func GetAuthToken() string {
	envConf, _ := GetEnvironmentConfig()
	jsonString := fmt.Sprintf(`{"username": "%s", "password": "%s"}`, envConf.LogInUserName, envConf.LogInUserPwd)
	resp, err := MakeApiCall(createSessionApiUrl, http.MethodPost, jsonString, nil, "")
	HandleError(err, "getAuthToken")
	var logInResponse LogInResponse
	json.Unmarshal(resp.Body(), &logInResponse)
	return logInResponse.Result.Token
}

func GetEnvironmentConfig() (*EnvironmentConfig, error) {
	cfg := &EnvironmentConfig{}
	err := env.Parse(cfg)
	if err != nil {
		return nil, errors.New("could not get config from environment")
	}
	return cfg, err
}

const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func GetRandomStringOfGivenLength(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func GetRandomNumberOf9Digit() int {
	return 100000000 + rand.Intn(999999999-100000000)
}
