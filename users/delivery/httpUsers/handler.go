package httpUsers

import (
	"ChatAPI/models"
	"ChatAPI/users"
	"net/http"
	"strings"

	"github.com/prometheus/common/log"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	UseCase users.UseCaseUsers
}

func NewHandlerUsers(useCase users.UseCaseUsers) *Handler {
	return &Handler{
		UseCase: useCase,
	}
}

func (h *Handler) AddUser(c *gin.Context) {
	var json models.User

	// unmarshal
	err := c.ShouldBindJSON(&json)
	if err != nil {
		log.Error(err)
		c.String(http.StatusBadRequest, "invalid json")
	}

	// validation data
	if strings.TrimSpace(json.Username) == "" || json.Id != 0 {
		log.Error(err)
		c.String(http.StatusBadRequest, err.Error())
	}

	// add user -> UseCase
	user, err := h.UseCase.AddUser(c, &json)
	if err != nil {
		log.Error(err)
		c.String(http.StatusBadRequest, err.Error())
	}

	c.JSON(http.StatusOK, user)
}
