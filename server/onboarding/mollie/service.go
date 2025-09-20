package mollie

import (
	"log"
	"net/http"

	"github.com/VictorAvelar/mollie-api-go/v4/mollie"
)

func NewService() {
	// Create a configuration object with idempotency enabled.
	config := mollie.NewAPITestingConfig(true)

	// build your desired config
	client, err := mollie.NewClient(http.DefaultClient, config)
	if err != nil {
		log.Fatal(err)
	}
	// perform operations with the API.

	client.NewAPIRequest()
}
