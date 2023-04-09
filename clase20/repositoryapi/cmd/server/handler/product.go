package handler

import (
	"net/http"
	"repositoryapi/internal/product"
	"strconv"

	"repositoryapi/pkg/web"
	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	ProductService product.IService 
}

func (h *ProductHandler) GetById(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, web.NewBadRequestApiError("Invalid ID"))
		return
	}

	prodFounded, err := h.ProductService.GetProductBy(id)
	if err != nil {
		if errApi, ok := err.(*web.ErrorApi); ok {
			ctx.AbortWithStatusJSON(errApi.Status, errApi)
			return
		}
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, &prodFounded)
}

