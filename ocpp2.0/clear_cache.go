package ocpp2

import (
	"gopkg.in/go-playground/validator.v9"
	"reflect"
)

// -------------------- Clear Cache (CSMS -> CS) --------------------

// Status returned in response to ClearCacheRequest.
type ClearCacheStatus string

const (
	ClearCacheStatusAccepted ClearCacheStatus = "Accepted"
	ClearCacheStatusRejected ClearCacheStatus = "Rejected"
)

func isValidClearCacheStatus(fl validator.FieldLevel) bool {
	status := ClearCacheStatus(fl.Field().String())
	switch status {
	case ClearCacheStatusAccepted, ClearCacheStatusRejected:
		return true
	default:
		return false
	}
}

// The field definition of the ClearCache request payload sent by the CSMS to the Charging Station.
type ClearCacheRequest struct {
}

// This field definition of the ClearCache confirmation payload, sent by the Charging Station to the CSMS in response to a ClearCacheRequest.
// In case the request was invalid, or couldn't be processed, an error will be sent instead.
type ClearCacheConfirmation struct {
	Status ClearCacheStatus `json:"status" validate:"required,cacheStatus"`
}

// CSMS can request a Charging Station to clear its Authorization Cache.
// The CSMS SHALL send a ClearCacheRequest payload for clearing the Charging Station’s Authorization Cache.
// Upon receipt of a ClearCacheRequest, the Charging Station SHALL respond with a ClearCacheConfirmation payload.
// The response payload SHALL indicate whether the Charging Station was able to clear its Authorization Cache.
type ClearCacheFeature struct{}

func (f ClearCacheFeature) GetFeatureName() string {
	return ClearCacheFeatureName
}

func (f ClearCacheFeature) GetRequestType() reflect.Type {
	return reflect.TypeOf(ClearCacheRequest{})
}

func (f ClearCacheFeature) GetConfirmationType() reflect.Type {
	return reflect.TypeOf(ClearCacheConfirmation{})
}

func (r ClearCacheRequest) GetFeatureName() string {
	return ClearCacheFeatureName
}

func (c ClearCacheConfirmation) GetFeatureName() string {
	return ClearCacheFeatureName
}

// Creates a new ClearCacheRequest, which doesn't contain any required or optional fields.
func NewClearCacheRequest() *ClearCacheRequest {
	return &ClearCacheRequest{}
}

// Creates a new ClearCacheConfirmation, containing all required fields. There are no optional fields for this message.
func NewClearCacheConfirmation(status ClearCacheStatus) *ClearCacheConfirmation {
	return &ClearCacheConfirmation{Status: status}
}

func init() {
	_ = Validate.RegisterValidation("cacheStatus", isValidClearCacheStatus)
}