package handlers

import (
	"community-funds/config"
	"community-funds/pkg/services"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/sirupsen/logrus"
	"github.com/stripe/stripe-go/v81"
)

type StripeHandler struct {
	log     *logrus.Logger
	service *services.StripeService
}

func NewStripeHandler(log *logrus.Logger, service *services.StripeService, cfg *config.Config) *StripeHandler {
	stripe.Key = cfg.StripeKey
	return &StripeHandler{
		log:     log,
		service: service,
	}
}

func (h *StripeHandler) CreateCheckoutSession(c *fiber.Ctx) error {
	var req struct {
		Account string `json:"account"`
	}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	clientSecret, err := h.service.CreateCheckoutSession(req.Account)
	if err != nil {
		log.Debugf("An error occurred when calling the Stripe API to create an account session: %v", err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"client_secret": clientSecret})
}

func (h *StripeHandler) CreateAccountLink(c *fiber.Ctx) error {
	var req struct {
		Account string `json:"account"`
	}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	clientSecret, err := h.service.CreateAccountLink(req.Account)
	if err != nil {
		log.Debugf("An error occurred when calling the Stripe API to create an account session: %v", err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"client_secret": clientSecret})
}

func (h *StripeHandler) CreateAccount(c *fiber.Ctx) error {
	accountID, err := h.service.CreateAccount()

	if err != nil {
		log.Debugf("An error occurred when calling the Stripe API to create an account: %v", err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"account": accountID})
}

func (h *StripeHandler) CreateAccountSession(c *fiber.Ctx) error {
	var req struct {
		Account string `json:"account"`
	}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	clientSecret, err := h.service.CreateAccountSession(req.Account)
	if err != nil {
		log.Debugf("An error occurred when calling the Stripe API to create an account session: %v", err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"client_secret": clientSecret})
}
