package handlers

import (
	"sync"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/javed2214/arctic-wolf-assignment/internal/models"
	"github.com/javed2214/arctic-wolf-assignment/types"
)

var (
	riskData = make(map[string]models.Risk)
	mutex    = &sync.Mutex{}
)

func GetRisks(c *fiber.Ctx) error {
	mutex.Lock()
	defer mutex.Unlock()

	var risks []models.Risk
	for _, risk := range riskData {
		risks = append(risks, risk)
	}

	return c.JSON(risks)
}

func CreateRisk(c *fiber.Ctx) error {
	mutex.Lock()
	defer mutex.Unlock()

	var newRisk models.Risk
	if err := c.BodyParser(&newRisk); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid data format",
		})
	}

	if newRisk.State != types.OPEN && newRisk.State != types.CLOSED && newRisk.State != types.ACCEPTED && newRisk.State != types.INVESTIGATING {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid state value",
		})
	}
	newRisk.ID = uuid.New().String()

	riskData[newRisk.ID] = newRisk

	return c.Status(fiber.StatusCreated).JSON(newRisk)
}

func GetRiskByID(c *fiber.Ctx) error {
	mutex.Lock()
	defer mutex.Unlock()

	riskID := c.Params("id")
	if risk, exists := riskData[riskID]; exists {
		return c.JSON(risk)
	}
	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
		"error": "Risk not found",
	})
}
