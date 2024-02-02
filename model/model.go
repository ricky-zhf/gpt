package model

type Req struct {
	Model string     `json:"model"`
	Msg   []Messages `json:"messages"`
}

type Messages struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ApifoxModel struct {
	Choices []Choice `json:"choices"`
	Created int64    `json:"created"`
	ID      string   `json:"id"`
	Object  string   `json:"object"`
	Usage   Usage    `json:"usage"`
}

type Choice struct {
	FinishReason *string  `json:"finish_reason,omitempty"`
	Index        *int64   `json:"index,omitempty"`
	Message      *Message `json:"message,omitempty"`
}

type Message struct {
	Content string `json:"content"`
	Role    string `json:"role"`
}

type Usage struct {
	CompletionTokens int64 `json:"completion_tokens"`
	PromptTokens     int64 `json:"prompt_tokens"`
	TotalTokens      int64 `json:"total_tokens"`
}
