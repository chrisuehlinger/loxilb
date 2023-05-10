// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"
	"strings"

	"github.com/go-openapi/swag"
)

// DeleteConfigEndpointEpipaddressIPAddressProbetypeProbeTypeProbeportProbePortURL generates an URL for the delete config endpoint epipaddress IP address probetype probe type probeport probe port operation
type DeleteConfigEndpointEpipaddressIPAddressProbetypeProbeTypeProbeportProbePortURL struct {
	IPAddress string
	ProbePort float64
	ProbeType string

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *DeleteConfigEndpointEpipaddressIPAddressProbetypeProbeTypeProbeportProbePortURL) WithBasePath(bp string) *DeleteConfigEndpointEpipaddressIPAddressProbetypeProbeTypeProbeportProbePortURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *DeleteConfigEndpointEpipaddressIPAddressProbetypeProbeTypeProbeportProbePortURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *DeleteConfigEndpointEpipaddressIPAddressProbetypeProbeTypeProbeportProbePortURL) Build() (*url.URL, error) {
	var _result url.URL

	var _path = "/config/endpoint/epipaddress/{ip_address}/probetype/{probe_type}/probeport/{probe_port}"

	iPAddress := o.IPAddress
	if iPAddress != "" {
		_path = strings.Replace(_path, "{ip_address}", iPAddress, -1)
	} else {
		return nil, errors.New("ipAddress is required on DeleteConfigEndpointEpipaddressIPAddressProbetypeProbeTypeProbeportProbePortURL")
	}

	probePort := swag.FormatFloat64(o.ProbePort)
	if probePort != "" {
		_path = strings.Replace(_path, "{probe_port}", probePort, -1)
	} else {
		return nil, errors.New("probePort is required on DeleteConfigEndpointEpipaddressIPAddressProbetypeProbeTypeProbeportProbePortURL")
	}

	probeType := o.ProbeType
	if probeType != "" {
		_path = strings.Replace(_path, "{probe_type}", probeType, -1)
	} else {
		return nil, errors.New("probeType is required on DeleteConfigEndpointEpipaddressIPAddressProbetypeProbeTypeProbeportProbePortURL")
	}

	_basePath := o._basePath
	if _basePath == "" {
		_basePath = "/netlox/v1"
	}
	_result.Path = golangswaggerpaths.Join(_basePath, _path)

	return &_result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *DeleteConfigEndpointEpipaddressIPAddressProbetypeProbeTypeProbeportProbePortURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *DeleteConfigEndpointEpipaddressIPAddressProbetypeProbeTypeProbeportProbePortURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *DeleteConfigEndpointEpipaddressIPAddressProbetypeProbeTypeProbeportProbePortURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on DeleteConfigEndpointEpipaddressIPAddressProbetypeProbeTypeProbeportProbePortURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on DeleteConfigEndpointEpipaddressIPAddressProbetypeProbeTypeProbeportProbePortURL")
	}

	base, err := o.Build()
	if err != nil {
		return nil, err
	}

	base.Scheme = scheme
	base.Host = host
	return base, nil
}

// StringFull returns the string representation of a complete url
func (o *DeleteConfigEndpointEpipaddressIPAddressProbetypeProbeTypeProbeportProbePortURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}
