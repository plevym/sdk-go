package point

import (
	"context"
	"net/http"

	"github.com/plevym/sdk-go/pkg/config"
	"github.com/plevym/sdk-go/pkg/internal/httpclient"
)

const (
	urlBase                = "https://api.mercadopago.com/point"
	urlDevices             = urlBase + "/integration-api/devices"
	urlPaymentIntent       = urlDevices + "/{device_id}/payment-intents"
	urlPaymentIntentGet    = urlBase + "/integration-api/payment-intents/{payment_intent_id}"
	urlPaymentIntentCancel = urlDevices + "/{device_id}/payment-intents/{payment_intent_id}"
	urlDevicesWithID       = urlDevices + "/{device_id}"
)

// client is the implementation of Client.
type client struct {
	cfg *config.Config
}

// Client contains the methods to interact with the Point API.
type Client interface {
	// Create a point payment intent.
	// It is a post request to the endpoint: https://api.mercadopago.com/point/integration-api/devices/{device_id}/payment-intents
	// Reference: https://www.mercadopago.com/developers/en/reference/integrations_api_paymentintent_mlb/_point_integration-api_devices_deviceid_payment-intents/post
	Create(ctx context.Context, deviceID string, request Request) (*Response, error)

	// Get a point payment intent.
	// It is a get request to the endpoint: https://api.mercadopago.com/point/integration-api/payment-intents/{payment_intent_id}
	// Reference: https://www.mercadopago.com/developers/en/reference/integrations_api/_point_integration-api_payment-intents_paymentintentid/get
	Get(ctx context.Context, paymentIntentID string) (*Response, error)

	// Cancel a point payment intent.
	// It is a cancel request to the endpoint: https://api.mercadopago.com/point/integration-api/devices/{device_id}/payment-intents/{payment_intent_id}
	// Reference: https://www.mercadopago.com/developers/en/reference/integrations_api/_point_integration-api_devices_deviceid_payment-intents_paymentintentid/delete
	Cancel(ctx context.Context, deviceID, paymentIntentID string) (*CancelResponse, error)

	// ListDevices retrieve devices.
	// It is a get request to the endpoint: https://api.mercadopago.com/point/integration-api/devices
	// Reference: https://www.mercadopago.com/developers/en/reference/integrations_api/_point_integration-api_devices/get
	ListDevices(ctx context.Context) (*DevicesResponse, error)

	// UpdateOperatingMode update operating mode from device.
	// It is a patch request to the endpoint: https://api.mercadopago.com/point/integration-api/devices/{device-id}
	// Reference: https://www.mercadopago.com/developers/en/reference/integrations_api/_point_integration-api_devices_device-id/patch
	UpdateOperatingMode(ctx context.Context, deviceID, operatingMode string) (*OperatingModeResponse, error)
}

// NewClient returns a new Point Client.
func NewClient(c *config.Config) Client {
	return &client{cfg: c}
}

func (c *client) Create(ctx context.Context, deviceID string, request Request) (*Response, error) {
	pathParams := map[string]string{
		"device_id": deviceID,
	}

	requestData := httpclient.RequestData{
		Body:       request,
		Method:     http.MethodPost,
		URL:        urlPaymentIntent,
		PathParams: pathParams,
	}
	resource, err := httpclient.DoRequest[*Response](ctx, c.cfg, requestData)
	if err != nil {
		return nil, err
	}

	return resource, nil
}

func (c *client) Get(ctx context.Context, paymentIntentID string) (*Response, error) {
	pathParams := map[string]string{
		"payment_intent_id": paymentIntentID,
	}

	requestData := httpclient.RequestData{
		Method:     http.MethodGet,
		URL:        urlPaymentIntentGet,
		PathParams: pathParams,
	}
	resource, err := httpclient.DoRequest[*Response](ctx, c.cfg, requestData)
	if err != nil {
		return nil, err
	}

	return resource, nil
}

func (c *client) Cancel(ctx context.Context, deviceID string, paymentIntentID string) (*CancelResponse, error) {
	pathParams := map[string]string{
		"device_id":         deviceID,
		"payment_intent_id": paymentIntentID,
	}

	requestData := httpclient.RequestData{
		Method:     http.MethodDelete,
		URL:        urlPaymentIntentCancel,
		PathParams: pathParams,
	}
	resource, err := httpclient.DoRequest[*CancelResponse](ctx, c.cfg, requestData)
	if err != nil {
		return nil, err
	}

	return resource, nil
}

func (c *client) ListDevices(ctx context.Context) (*DevicesResponse, error) {
	requestData := httpclient.RequestData{
		Method: http.MethodGet,
		URL:    urlDevices,
	}
	resource, err := httpclient.DoRequest[*DevicesResponse](ctx, c.cfg, requestData)
	if err != nil {
		return nil, err
	}

	return resource, nil
}

func (c *client) UpdateOperatingMode(ctx context.Context, deviceID, operatingMode string) (*OperatingModeResponse, error) {
	request := &OperatingModeRequest{OperatingMode: operatingMode}

	pathParams := map[string]string{
		"device_id": deviceID,
	}

	requestData := httpclient.RequestData{
		Body:       request,
		Method:     http.MethodPatch,
		URL:        urlDevicesWithID,
		PathParams: pathParams,
	}
	resource, err := httpclient.DoRequest[*OperatingModeResponse](ctx, c.cfg, requestData)
	if err != nil {
		return nil, err
	}

	return resource, nil
}
