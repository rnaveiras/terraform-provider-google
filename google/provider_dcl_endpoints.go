// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package google

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var EventarcEndpointEntryKey = "eventarc_custom_endpoint"
var EventarcEndpointEntry = &schema.Schema{
	Type:     schema.TypeString,
	Optional: true,
	DefaultFunc: schema.MultiEnvDefaultFunc([]string{
		"GOOGLE_EVENTARC_CUSTOM_ENDPOINT",
	}, ""),
}

//Add new values to config.go.erb config object declaration
// clientDataprocDCL *dataprocDcl.Client
//EventarcBasePath string
// clientEventarcDCL *eventarcDcl.Client

//Add new values to config.go.erb object initialization
// c.clientDataprocDCL = dataprocDcl.NewClient(dcl.NewConfig(dclClientOptions, dclUserAgentOptions,dclLoggerOptions, dcl.WithBasePath(c.DataprocBasePath)))
// c.clientEventarcDCL = eventarcDcl.NewClient(dcl.NewConfig(dclClientOptions, dclUserAgentOptions,dclLoggerOptions, dcl.WithBasePath(c.EventarcBasePath)))

//Add new values to provider.go.erb schema initialization
// EventarcEndpointEntryKey:               EventarcEndpointEntry,

//Add new values to provider.go.erb - provider block read
// config.EventarcBasePath = d.Get(EventarcEndpointEntryKey).(string)
