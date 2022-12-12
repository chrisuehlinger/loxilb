// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
)

// NewDeleteConfigFdbMacAddressDevIfNameParams creates a new DeleteConfigFdbMacAddressDevIfNameParams object
//
// There are no default values defined in the spec.
func NewDeleteConfigFdbMacAddressDevIfNameParams() DeleteConfigFdbMacAddressDevIfNameParams {

	return DeleteConfigFdbMacAddressDevIfNameParams{}
}

// DeleteConfigFdbMacAddressDevIfNameParams contains all the bound params for the delete config fdb mac address dev if name operation
// typically these are obtained from a http.Request
//
// swagger:parameters DeleteConfigFdbMacAddressDevIfName
type DeleteConfigFdbMacAddressDevIfNameParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Attributes of the target device
	  Required: true
	  In: path
	*/
	IfName string
	/*Attributes IPv4 Address in the device
	  Required: true
	  In: path
	*/
	MacAddress string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewDeleteConfigFdbMacAddressDevIfNameParams() beforehand.
func (o *DeleteConfigFdbMacAddressDevIfNameParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rIfName, rhkIfName, _ := route.Params.GetOK("if_name")
	if err := o.bindIfName(rIfName, rhkIfName, route.Formats); err != nil {
		res = append(res, err)
	}

	rMacAddress, rhkMacAddress, _ := route.Params.GetOK("mac_address")
	if err := o.bindMacAddress(rMacAddress, rhkMacAddress, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindIfName binds and validates parameter IfName from path.
func (o *DeleteConfigFdbMacAddressDevIfNameParams) bindIfName(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route
	o.IfName = raw

	return nil
}

// bindMacAddress binds and validates parameter MacAddress from path.
func (o *DeleteConfigFdbMacAddressDevIfNameParams) bindMacAddress(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route
	o.MacAddress = raw

	return nil
}