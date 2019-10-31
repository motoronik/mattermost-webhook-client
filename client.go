package mattermost_webhook_client

import (
	"fmt"
	"net/http"
	"time"
	"io/ioutil"
	"encoding/json"
	"bytes"
)

type WebhookSettings struct {
	url string
	user string
	channel string
}

type WebhookClient struct {
	settings WebhookSettings
	timeout time.Duration
}

func CreateWebhookClient(url string, user string, channel string) *WebhookClient {
	return &WebhookClient{
		WebhookSettings{url, user, channel},
		time.Duration(5),
	}
}

func (client WebhookClient) Post (message string) (body string, code int) {
	httpClient := http.Client{
		Timeout: time.Second * client.timeout,
	}

	payload := map[string] string {
		"channel" : client.settings.channel,
		"username" : client.settings.user,
		"text" : message,
	}

	payloadJson, _ := json.Marshal(payload)

	response, errorHttpClient := httpClient.Post(
		client.settings.url,
		"application/json",
		bytes.NewBuffer(payloadJson),
		)

	if errorHttpClient != nil {
		fmt.Print(errorHttpClient)
	}

	buffer, _ := ioutil.ReadAll(response.Body)

	body = string(buffer)
	code = response.StatusCode
	return
}