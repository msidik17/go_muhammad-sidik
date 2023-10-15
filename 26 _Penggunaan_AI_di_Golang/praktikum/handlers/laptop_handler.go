package handler

import (
	"AI/usecase"
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo"
)

type LaptopHandler struct {
	LaptopUsecase usecase.LaptopUsecase
}

type LaptopResponse struct {
	Status string `json:"status"`
	Data   string `json:"data"`
}

func NewLaptopHandler(usecase usecase.LaptopUsecase) *LaptopHandler {
	return &LaptopHandler{
		LaptopUsecase: usecase,
	}
}

func (h *LaptopHandler) RecommendLaptop(c echo.Context) error {
	var requestData map[string]interface{}
	err := c.Bind(&requestData)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid JSON format")
	}

	budget, okBudget := requestData["budget"].(float64)
	purpose, okPurpose := requestData["purpose"].(string)
	brand, okBrand := requestData["brand"].(string)

	if !okBudget || !okPurpose || !okBrand {
		return c.JSON(http.StatusBadRequest, "Invalid request format")
	}

	userInput := fmt.Sprintf("Rekomendasi Laptop merk %s untuk %s dan dengan budget sebesar Rp %.0f.", brand, purpose, budget)

	answer, err := h.LaptopUsecase.RecommendLaptop(userInput, brand, os.Getenv("secretKeyAPI"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Error dalam mmebuat rekomendasi")
	}

	responseData := LaptopResponse{
		Status: "success",
		Data:   answer,
	}

	return c.JSON(http.StatusOK, responseData)
}
