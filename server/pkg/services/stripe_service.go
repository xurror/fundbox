package services

import (
	"community-funds/config"
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/stripe/stripe-go/v81"
	"github.com/stripe/stripe-go/v81/account"
	"github.com/stripe/stripe-go/v81/accountlink"
	"github.com/stripe/stripe-go/v81/accountsession"
	"github.com/stripe/stripe-go/v81/checkout/session"
)

type StripeService struct {
	log *logrus.Logger
}

func NewStripeService(log *logrus.Logger, cfg *config.Config) *StripeService {
	stripe.Key = cfg.StripeKey
	return &StripeService{
		log: log,
	}
}

func (h *StripeService) CreateCheckoutSession(accountID string) (*string, error) {
	params := &stripe.CheckoutSessionParams{
		LineItems: []*stripe.CheckoutSessionLineItemParams{
			{
				PriceData: &stripe.CheckoutSessionLineItemPriceDataParams{
					Currency: stripe.String(string(stripe.CurrencyUSD)),
					ProductData: &stripe.CheckoutSessionLineItemPriceDataProductDataParams{
						Name: stripe.String("T-shirt"),
					},
					UnitAmount: stripe.Int64(1000),
				},
				Quantity: stripe.Int64(1),
			},
		},
		PaymentIntentData: &stripe.CheckoutSessionPaymentIntentDataParams{
			ApplicationFeeAmount: stripe.Int64(123),
		},
		Mode:      stripe.String(string(stripe.CheckoutSessionModePayment)),
		UIMode:    stripe.String(string(stripe.CheckoutSessionUIModeEmbedded)),
		ReturnURL: stripe.String("https://example.com/checkout/return?session_id={CHECKOUT_SESSION_ID}"),
	}
	params.SetStripeAccount(accountID)

	session, err := session.New(params)
	if err != nil {
		return nil, err
	}
	return &session.ClientSecret, nil
}

func (h *StripeService) CreateAccountLink(account string) (*string, error) {
	params := &stripe.AccountLinkParams{
		Account:    stripe.String(account),
		ReturnURL:  stripe.String(fmt.Sprintf("http://localhost:4242/return/%s", account)),
		RefreshURL: stripe.String(fmt.Sprintf("http://localhost:4242/refresh/%s", account)),
		Type:       stripe.String("account_onboarding"),
	}
	accountLink, err := accountlink.New(params)
	if err != nil {
		return nil, err
	}

	return &accountLink.URL, nil
}

func (h *StripeService) CreateAccount() (*string, error) {
	params := &stripe.AccountParams{
		Controller: &stripe.AccountControllerParams{
			StripeDashboard: &stripe.AccountControllerStripeDashboardParams{
				Type: stripe.String("none"),
			},
			Fees: &stripe.AccountControllerFeesParams{
				Payer: stripe.String("application"),
			},
			Losses: &stripe.AccountControllerLossesParams{
				Payments: stripe.String("application"),
			},
			RequirementCollection: stripe.String("application"),
		},
		Capabilities: &stripe.AccountCapabilitiesParams{
			CardPayments: &stripe.AccountCapabilitiesCardPaymentsParams{
				Requested: stripe.Bool(true),
			},
			Transfers: &stripe.AccountCapabilitiesTransfersParams{
				Requested: stripe.Bool(true),
			},
		},
		Country: stripe.String("US"),
	}

	account, err := account.New(params)
	if err != nil {
		return nil, err
	}

	return &account.ID, nil
}

func (h *StripeService) CreateAccountSession(account string) (*string, error) {
	params := &stripe.AccountSessionParams{
		Account: stripe.String(account),
		Components: &stripe.AccountSessionComponentsParams{
			AccountOnboarding: &stripe.AccountSessionComponentsAccountOnboardingParams{
				Enabled: stripe.Bool(true),
			},
		},
	}

	accountSession, err := accountsession.New(params)
	if err != nil {
		return nil, err
	}

	return &accountSession.ClientSecret, nil
}
