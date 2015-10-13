// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package cloudsearchdomain

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/defaults"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/service"
	"github.com/aws/aws-sdk-go/aws/service/serviceinfo"
	"github.com/aws/aws-sdk-go/internal/protocol/restjson"
	"github.com/aws/aws-sdk-go/internal/signer/v4"
)

// You use the AmazonCloudSearch2013 API to upload documents to a search domain
// and search those documents.
//
// The endpoints for submitting UploadDocuments, Search, and Suggest requests
// are domain-specific. To get the endpoints for your domain, use the Amazon
// CloudSearch configuration service DescribeDomains action. The domain endpoints
// are also displayed on the domain dashboard in the Amazon CloudSearch console.
// You submit suggest requests to the search endpoint.
//
// For more information, see the Amazon CloudSearch Developer Guide (http://docs.aws.amazon.com/cloudsearch/latest/developerguide).
type CloudSearchDomain struct {
	*service.Service
}

// Used for custom service initialization logic
var initService func(*service.Service)

// Used for custom request initialization logic
var initRequest func(*request.Request)

// New returns a new CloudSearchDomain client.
func New(config *aws.Config) *CloudSearchDomain {
	service := &service.Service{
		ServiceInfo: serviceinfo.ServiceInfo{
			Config:      defaults.DefaultConfig.Merge(config),
			ServiceName: "cloudsearchdomain",
			SigningName: "cloudsearch",
			APIVersion:  "2013-01-01",
		},
	}
	service.Initialize()

	// Handlers
	service.Handlers.Sign.PushBack(v4.Sign)
	service.Handlers.Build.PushBack(restjson.Build)
	service.Handlers.Unmarshal.PushBack(restjson.Unmarshal)
	service.Handlers.UnmarshalMeta.PushBack(restjson.UnmarshalMeta)
	service.Handlers.UnmarshalError.PushBack(restjson.UnmarshalError)

	// Run custom service initialization if present
	if initService != nil {
		initService(service)
	}

	return &CloudSearchDomain{service}
}

// newRequest creates a new request for a CloudSearchDomain operation and runs any
// custom request initialization.
func (c *CloudSearchDomain) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	// Run custom request initialization if present
	if initRequest != nil {
		initRequest(req)
	}

	return req
}