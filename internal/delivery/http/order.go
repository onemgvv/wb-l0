package http

import (
	"github.com/gin-gonic/gin"
	"github.com/onemgvv/wb-l0/internal/domain"
	"net/http"
)

func (h *Handler) GetByID(ctx *gin.Context) {
	orderId := ctx.Param("id")

	one, err := h.service.Orders.GetById(orderId)
	if err != nil {
		ctx.JSON(http.StatusNotFound, NotFoundResponse(&ResponseInput{
			Message: OrderNotFound, Data: err.Error(),
		}))
	}

	ctx.JSON(http.StatusOK, OkResponse(&ResponseInput{
		Message: OrderFound, Data: map[string]domain.Order{"order": one},
	}))
}
