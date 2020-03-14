package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"net/http"
	"os"
)

type Handler struct {
	store Store
}

type Response struct {
	Error bool `json:"error"`
	Message string `json:"message"`
	Data interface{} `json:"data"`
}

func newHandler(s Store) *Handler {
	return &Handler{store: s}
}

func (handler *Handler) handleSet(ctx *gin.Context) {
	key, value := ctx.Query("key"), ctx.Query("value")
	if key == "" || value == "" {
		ctx.JSON(http.StatusBadRequest, &Response{Error: true, Message: "bad request: key or value is missing"})
		return
	}
	if err := handler.store.Set(key, value); err != nil {
		ctx.JSON(http.StatusInternalServerError, &Response{Error: true, Message: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, &Response{Error: false, Message: "success"})
}

func (handler *Handler) handleGet(ctx *gin.Context) {
	key := ctx.Query("key")
	if key == "" {
		ctx.JSON(http.StatusBadRequest, &Response{Error: true, Message: "bad request: key is missing"})
		return
	}
	value, err := handler.store.Get(key)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &Response{Error: true, Message: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, &Response{Error: false, Message: "success", Data: value})
}

func (handler *Handler) handleStatus(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, &Response{Error: false, Message: "we're okay!"})
}

func Run() error {
	redisUri := os.Getenv("REDIS_HOST")
	redisClient := redis.NewClient(&redis.Options{
		Addr: redisUri,
		DB: 0,
	})
	if err := redisClient.Ping().Err(); err != nil {
		return err
	}
	s := newRedisStore(redisClient)
	handler := newHandler(s)

	router := gin.Default()
	router.GET("/set", handler.handleSet)
	router.GET("/get", handler.handleGet)
	router.GET("/status", handler.handleStatus)

	addr := os.Getenv("PORT")
	if addr == "" {
		addr = "7002"
	}
	addr = fmt.Sprintf(":%s", addr)
	if err := router.Run(addr); err != nil {
		return err
	}
	return nil
}