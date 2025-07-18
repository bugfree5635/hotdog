package service

import (
	"context"
	"os"

	openai "github.com/sashabaranov/go-openai"
)

var client *openai.Client

func Init() {
	// 1. 阿里云百炼（通义千问）
	// baseURL := "https://dashscope.aliyuncs.com/compatible-mode/v1"
	// apiKey  := os.Getenv("DASHSCOPE_API_KEY")

	// 2. 智谱 GLM（取消下面两行注释即可）
	baseURL := "https://open.bigmodel.cn/api/paas/v4"
	apiKey := os.Getenv("ZHIPU_API_KEY")

	// 3. 原生 OpenAI（取消下面两行注释即可）
	// baseURL = "https://api.openai.com/v1"
	// apiKey  = os.Getenv("OPENAI_API_KEY")

	if apiKey == "" {
		panic("missing API_KEY env")
	}
	cfg := openai.DefaultConfig(apiKey)
	cfg.BaseURL = baseURL
	client = openai.NewClientWithConfig(cfg)
}

func Chat(ctx context.Context, prompt string) (string, error) {
	resp, err := client.CreateChatCompletion(ctx, openai.ChatCompletionRequest{
		Model:       "glm-4-air", // 或 glm-4 / gpt-4o / gpt-3.5-turbo
		Temperature: 0.3,
		Messages: []openai.ChatCompletionMessage{
			{Role: openai.ChatMessageRoleUser, Content: prompt},
		},
	})
	if err != nil {
		return "", err
	}
	return resp.Choices[0].Message.Content, nil
}
