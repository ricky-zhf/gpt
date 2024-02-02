package main

import (
	"GPT/constant"
	"GPT/model"
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	msgs := []model.Messages{}
	for {
		fmt.Print("You: ")
		scanner.Scan()
		input := scanner.Text()
		if input == "exit" {
			break
		}

		if input == "clear" {
			msgs = []model.Messages{}
		}

		msgs = append(msgs, model.Messages{
			Role:    "user",
			Content: input,
		})

		output, err := sendGPTRequest(msgs)
		if err != nil {
			fmt.Printf("Failed. %v\n", err)
			continue
		}

		fmt.Println("GPT:", output)
		fmt.Println()
	}
}

var client = &http.Client{}

func sendGPTRequest(msg []model.Messages) (string, error) {
	payload, err := json.Marshal(model.Req{
		Model: "gpt-3.5-turbo",
		Msg:   msg,
	})
	if err != nil {
		return "", fmt.Errorf("marshal failed. %v", err)
	}

	req, err := http.NewRequest("POST", constant.Url, bytes.NewBuffer(payload))
	if err != nil {
		return "", fmt.Errorf("new request failed. %v", err)
	}
	req.Header.Set("Authorization", "Bearer "+constant.ApiKey)
	req.Header.Set("Content-Type", "application/json")

	var response model.ApifoxModel
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return "", fmt.Errorf("req failed. %v", resp.Status)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("read failed. %v", err)
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		return "", fmt.Errorf("unmarshal failed. %v", err)
	}

	return response.Choices[0].Message.Content, nil
}
