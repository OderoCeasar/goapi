package handler

import (
	"net/http"
	"strconv"

	"github.com/OderoCeasar/goapi/internal/middleware"
	"github.com/OderoCeasar/goapi/pkg/response"
)

// ListUsers demostrates query string parsing
// GET /api/v1/users?page=1&limit=10
func (h *Handler) ListUsers(w http.ResponseWriter, r *http.Request) {
	page := parseIntQuery(r, "page", 1)
	limit := parseIntQuery(r, "limit", 10)

	// clamp limit to prevent abuse
	if limit > 100 {
		limit = 100
	}

	requestID := middleware.GetRequestID(r.Context())

	response.JSON(w, http.StatusOK, map[string]interface{}{
		"users":	[]string{},
		"page":		page,
		"limit":	limit,
		"request_id": requestID,
	})
}


// GetUser demonstrates path parameter extraction
// GET /api/v1/users/{id}
func (h *Handler) GetUser(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		response.Error(w, http.StatusBadRequest, "user id is required")
		return
	}

	// validate its a number
	userID, err := strconv.Atoi(id)
	if err != nil {
		response.Error(w, http.StatusBadRequest, "user id must be a valid integer")
		return
	}

	requestID := middleware.GetRequestID(r.Context())

	// stub response
	response.JSON(w, http.StatusOK, map[string]interface{}{
		"id":		userID,
		"name":		"John Doe",
		"email": 	"john@example.com",
		"request_id": requestID,
	})
}


// CreateUser
// POST /api/v1/users
func (h *Handler) CreateUser(w http.ResponseWriter, r * http.Request) {
	response.JSON(w, http.StatusCreated, map[string]interface{}{
		"message": "user created successfully",
	})
}

// parseIntQuery
func parseIntQuery(r *http.Request, key string, fallback int) int {
	val := r.URL.Query().Get(key)
	if val == "" {
		return fallback
	}
	n, err := strconv.Atoi(val)
	if err != nil || n < 1 {
		return fallback
	}
	return n
}