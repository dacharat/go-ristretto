package handler

import (
	"net/http"
	"time"

	"github.com/dacharat/go-ristretto/pkg/caches"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	c caches.ICache
}

func NewHandler(c caches.ICache) Handler {
	return Handler{
		c: c,
	}
}

func (h *Handler) GetByKey(c *gin.Context) {
	key := c.Param("key")
	ctx := c.Request.Context()
	value, err := h.c.Get(ctx, key)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"value": value,
	})
}

func (h *Handler) SetByKey(c *gin.Context) {
	key := c.Param("key")
	ctx := c.Request.Context()

	var req Reqest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}
	duration, err := time.ParseDuration(req.TTL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}

	err = h.c.Set(ctx, key, req.Value, duration)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}
