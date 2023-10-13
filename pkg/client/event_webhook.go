package client

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"io"
	"net/http"

	"github.com/dezhishen/satori-sdk-go/pkg/config"
	"github.com/gin-gonic/gin"
)

type HttpReverseEventChannel struct {
	handler func(message []byte) error
	svc     *http.Server
	addr    string
}

func NewHttpReverseEventChannel(conf *config.SatoriEventConfig) (EventChannel, error) {
	router := gin.Default()
	if conf.AccessToken != "" {
		router.Use(func(c *gin.Context) {
			if c.GetHeader("Authorization") != "Bearer "+conf.AccessToken {
				c.AbortWithStatus(http.StatusUnauthorized)
				return
			}
			c.Next()
		})
	}
	/*
		X-Signature的内容中，sha1的值的计算方式为：
		获取请求体的内容，即 requestBody
		使用密钥为 secret的HMAC算法加密 requestBody
		得到的HMAC加密值以Hex格式输出为字符串
	*/
	router.Use(func(c *gin.Context) {
		hSignature := c.GetHeader("X-Signature")
		if hSignature != "" {
			if conf.Secret == "" {
				c.AbortWithStatus(http.StatusUnauthorized)
				return
			}
			// 读取 request body
			body, err := io.ReadAll(c.Request.Body)
			if err != nil {
				c.AbortWithStatus(http.StatusInternalServerError)
				return
			}
			// gin框架中request body只能读取一次，需要复写 request body
			c.Request.Body = io.NopCloser(bytes.NewBuffer(body))
			// 使用密钥计算request body的 hmac码
			mac := hmac.New(sha1.New, []byte(conf.Secret))
			if _, err := mac.Write(body); err != nil {
				c.AbortWithStatus(http.StatusInternalServerError)
				return
			}
			// 校验hmac签名
			if "sha1="+hex.EncodeToString(mac.Sum(nil)) != hSignature {
				c.AbortWithStatus(http.StatusUnauthorized)
				return
			}
		}
		c.Next()
	})

	result := &HttpReverseEventChannel{
		addr: conf.Addr,
	}
	router.POST("/event", func(c *gin.Context) {
		var body []byte
		var err error
		if body, err = io.ReadAll(c.Request.Body); err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		if err = result.handler(body); err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		c.Status(http.StatusOK)
	})
	result.svc = &http.Server{
		Addr:    conf.Addr,
		Handler: router,
	}
	return result, nil
}

func (cli *HttpReverseEventChannel) StartListen(ctx context.Context, handler func(message []byte) error) error {
	errChan := make(chan error)
	cli.handler = handler
	// 异常通过chan 传递
	go func() {
		if err := cli.svc.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			// panic(err)
			errChan <- err
		}
	}()
	select {
	case <-ctx.Done():
		return cli.svc.Shutdown(ctx)
	case err := <-errChan:
		return err
	}
}
