package handlers

import (
	"net/http"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type ActuatorHandler struct {
	db  *gorm.DB
	log *logrus.Logger
}

func NewActuatorHandler(log *logrus.Logger, db *gorm.DB) *ActuatorHandler {
	return &ActuatorHandler{
		db:  db,
		log: log,
	}
}

// Health Check - Basic API status
// @Summary Health Check
// @Description Returns API health status
// @Tags actuator
// @Produce json
// @Success 200 {object} map[string]string
// @Router /health [get]
func (h *ActuatorHandler) HealthCheck(c *fiber.Ctx) error {
	return c.Status(http.StatusOK).JSON(fiber.Map{"status": "UP"})
}

// Readiness Probe - Checks if dependencies (db) are healthy
// @Summary Readiness Probe
// @Description Returns readiness status based on database connectivity
// @Tags actuator
// @Produce json
// @Success 200 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /ready [get]
func (h *ActuatorHandler) ReadinessCheck(c *fiber.Ctx) error {
	sqlDB, err := h.db.DB()
	if err != nil || sqlDB.Ping() != nil {
		h.log.Warn("Readiness check failed: Database not available")
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"status": "NOT READY", "error": "Database connection failed"})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{"status": "READY"})
}

// Shutdown - Gracefully shuts down the API
// @Summary Shutdown API
// @Description Terminates the application gracefully
// @Tags actuator
// @Produce json
// @Success 200 {object} map[string]string
// @Router /shutdown [post]
func (h *ActuatorHandler) Shutdown(c *fiber.Ctx) error {
	h.log.Warn("Shutdown signal received, terminating service...")

	go func() {
		time.Sleep(1 * time.Second) // Allow time for the response to be sent
		os.Exit(0)
	}()

	return c.Status(http.StatusOK).JSON(fiber.Map{"status": "Shutting down..."})
}

// Restart - Simulated restart by shutting down and relying on Render auto-restart
// @Summary Restart API
// @Description Simulates an API restart
// @Tags actuator
// @Produce json
// @Success 200 {object} map[string]string
// @Router /restart [post]
func (h *ActuatorHandler) Restart(c *fiber.Ctx) error {
	h.log.Warn("Restart signal received, restarting service...")

	go func() {
		time.Sleep(1 * time.Second)
		os.Exit(1) // Exiting with a non-zero status prompts a restart in managed environments
	}()

	return c.Status(http.StatusOK).JSON(fiber.Map{"status": "Restarting..."})
}
