package handler

import (
	"fmt"
	"hotdog/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TranslateReq struct {
	Text string `json:"text" binding:"required"`
}

type SummarizeReq struct {
	Text   string `json:"text" binding:"required"`
	Lang   string `json:"lang,omitempty"`
	MaxLen int    `json:"max_len,omitempty"`
}

func Zh2En(c *gin.Context) {
	var req TranslateReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	prompt := "Translate the following Chinese text to English:\n" + req.Text
	reply, err := service.Chat(c.Request.Context(), prompt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"result": reply})
}

func En2Zh(c *gin.Context) {
	var req TranslateReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	prompt := "Translate the following English text to Chinese:\n" + req.Text
	reply, err := service.Chat(c.Request.Context(), prompt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"result": reply})
}

func Summarize(c *gin.Context) {
	var req SummarizeReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if req.MaxLen == 0 {
		req.MaxLen = 100
	}
	lang := req.Lang
	if lang == "" {
		lang = "zh"
	}
	prompt := fmt.Sprintf("请用%s语言对下面文本进行总结，控制在%d字以内。\n%s", lang, req.MaxLen, req.Text)
	reply, err := service.Chat(c.Request.Context(), prompt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"result": reply})
}
