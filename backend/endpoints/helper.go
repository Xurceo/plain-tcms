package endpoints

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
)

func handleGetAll[T any](c *gin.Context, fetch func() ([]T, error)) {
	items, err := fetch()
	if err != nil {
		slog.Error("failed to get all", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, items)
}

func handleGetByID[T any](c *gin.Context, fetch func(string) (T, error)) {
	id := c.Param("id")
	item, err := fetch(id)
	if err != nil {
		slog.Error("failed to get by id", "error", err, "id", id)
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}
	c.JSON(http.StatusOK, item)
}

func handleCreate[Req any, Res any](c *gin.Context, create func(Req) (Res, error)) {
	req, ok := bindJSON[Req](c)
	if !ok {
		return
	}
	item, err := create(req)
	if err != nil {
		slog.Error("failed to create", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, item)
}

func handleDelete(c *gin.Context, delete func(string) error) {
	id := c.Param("id")
	if err := delete(id); err != nil {
		slog.Error("failed to delete", "error", err, "id", id)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}

func bindJSON[T any](c *gin.Context) (T, bool) {
	var req T
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return req, false
	}
	return req, true
}
