package main

import (
	"GPT/model"
	"GPT/tphub"
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	messages := []model.Messages{}
	for {
		fmt.Print("You: ")
		scanner.Scan()
		input := scanner.Text()
		if input == "exit" {
			break
		}

		if input == "clear" {
			messages = []model.Messages{}
		}

		messages = append(messages, model.Messages{
			Role:    "user",
			Content: input,
		})

		output, err := tphub.SendGPTRequest(messages)
		if err != nil {
			fmt.Printf("Failed. %v\n", err)
			continue
		}

		fmt.Println("GPT:", output)
		fmt.Println()
	}
}
