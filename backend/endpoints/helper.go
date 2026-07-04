/*
 * // TCMS - Test Case Management System
 * // Copyright (C) 2026 Pavlo Shnal
 * //
 * // This program is free software: you can redistribute it and/or modify
 * // it under the terms of the GNU Affero General Public License as published
 * // by the Free Software Foundation, either version 3 of the License, or
 * // (at your option) any later version.
 * //
 * // This program is distributed in the hope that it will be useful,
 * // but WITHOUT ANY WARRANTY; without even the implied warranty of
 * // MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
 * // GNU Affero General Public License for more details.
 * //
 * // You should have received a copy of the GNU Affero General Public License
 * // along with this program. If not, see <https://www.gnu.org/licenses/>.
 */

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

func handleUpdate[Req any, Res any](c *gin.Context, update func(string, Req) (Res, error)) {
	id := c.Param("id")
	req, ok := bindJSON[Req](c)
	if !ok {
		return
	}
	item, err := update(id, req)
	if err != nil {
		slog.Error("failed to update", "error", err, "id", id)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, item)
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
