// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by aliasgen. DO NOT EDIT.

// Package analyticshub aliases all exported identifiers in package
// "cloud.google.com/go/bigquery/analyticshub/apiv1/analyticshubpb".
//
// Deprecated: Please use types in: cloud.google.com/go/bigquery/analyticshub/apiv1/analyticshubpb.
// Please read https://github.com/googleapis/google-cloud-go/blob/main/migration.md
// for more details.
package analyticshub

import (
	src "cloud.google.com/go/bigquery/analyticshub/apiv1/analyticshubpb"
	grpc "google.golang.org/grpc"
)

// Deprecated: Please use consts in: cloud.google.com/go/bigquery/analyticshub/apiv1/analyticshubpb
const (
	Listing_ACTIVE                                = src.Listing_ACTIVE
	Listing_CATEGORY_ADVERTISING_AND_MARKETING    = src.Listing_CATEGORY_ADVERTISING_AND_MARKETING
	Listing_CATEGORY_CLIMATE_AND_ENVIRONMENT      = src.Listing_CATEGORY_CLIMATE_AND_ENVIRONMENT
	Listing_CATEGORY_COMMERCE                     = src.Listing_CATEGORY_COMMERCE
	Listing_CATEGORY_DEMOGRAPHICS                 = src.Listing_CATEGORY_DEMOGRAPHICS
	Listing_CATEGORY_ECONOMICS                    = src.Listing_CATEGORY_ECONOMICS
	Listing_CATEGORY_EDUCATION                    = src.Listing_CATEGORY_EDUCATION
	Listing_CATEGORY_ENERGY                       = src.Listing_CATEGORY_ENERGY
	Listing_CATEGORY_FINANCIAL                    = src.Listing_CATEGORY_FINANCIAL
	Listing_CATEGORY_GAMING                       = src.Listing_CATEGORY_GAMING
	Listing_CATEGORY_GEOSPATIAL                   = src.Listing_CATEGORY_GEOSPATIAL
	Listing_CATEGORY_HEALTHCARE_AND_LIFE_SCIENCE  = src.Listing_CATEGORY_HEALTHCARE_AND_LIFE_SCIENCE
	Listing_CATEGORY_MEDIA                        = src.Listing_CATEGORY_MEDIA
	Listing_CATEGORY_OTHERS                       = src.Listing_CATEGORY_OTHERS
	Listing_CATEGORY_PUBLIC_SECTOR                = src.Listing_CATEGORY_PUBLIC_SECTOR
	Listing_CATEGORY_RETAIL                       = src.Listing_CATEGORY_RETAIL
	Listing_CATEGORY_SCIENCE_AND_RESEARCH         = src.Listing_CATEGORY_SCIENCE_AND_RESEARCH
	Listing_CATEGORY_SPORTS                       = src.Listing_CATEGORY_SPORTS
	Listing_CATEGORY_TRANSPORTATION_AND_LOGISTICS = src.Listing_CATEGORY_TRANSPORTATION_AND_LOGISTICS
	Listing_CATEGORY_TRAVEL_AND_TOURISM           = src.Listing_CATEGORY_TRAVEL_AND_TOURISM
	Listing_CATEGORY_UNSPECIFIED                  = src.Listing_CATEGORY_UNSPECIFIED
	Listing_STATE_UNSPECIFIED                     = src.Listing_STATE_UNSPECIFIED
)

// Deprecated: Please use vars in: cloud.google.com/go/bigquery/analyticshub/apiv1/analyticshubpb
var (
	File_google_cloud_bigquery_analyticshub_v1_analyticshub_proto = src.File_google_cloud_bigquery_analyticshub_v1_analyticshub_proto
	Listing_Category_name                                         = src.Listing_Category_name
	Listing_Category_value                                        = src.Listing_Category_value
	Listing_State_name                                            = src.Listing_State_name
	Listing_State_value                                           = src.Listing_State_value
)

// AnalyticsHubServiceClient is the client API for AnalyticsHubService
// service. For semantics around ctx use and closing/ending streaming RPCs,
// please refer to
// https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
//
// Deprecated: Please use types in: cloud.google.com/go/bigquery/analyticshub/apiv1/analyticshubpb
type AnalyticsHubServiceClient = src.AnalyticsHubServiceClient

// AnalyticsHubServiceServer is the server API for AnalyticsHubService
// service.
//
// Deprecated: Please use types in: cloud.google.com/go/bigquery/analyticshub/apiv1/analyticshubpb
type AnalyticsHubServiceServer = src.AnalyticsHubServiceServer

// Message for creating a data exchange.
//
// Deprecated: Please use types in: cloud.google.com/go/bigquery/analyticshub/apiv1/analyticshubpb
type CreateDataExchangeRequest = src.CreateDataExchangeRequest

// Message for creating a listing.
//
// Deprecated: Please use types in: cloud.google.com/go/bigquery/analyticshub/apiv1/analyticshubpb
type CreateListingRequest = src.CreateListingRequest

// A data exchange is a container that lets you share data. Along with the
// descriptive information about the data exchange, it contains listings that
// reference shared datasets.
//
// Deprecated: Please use types in: cloud.google.com/go/bigquery/analyticshub/apiv1/analyticshubpb
type DataExchange = src.DataExchange

// Contains details of the data provider.
//
// Deprecated: Please use types in: cloud.google.com/go/bigquery/analyticshub/apiv1/analyticshubpb
type DataProvider = src.DataProvider

// Message for deleting a data exchange.
//
// Deprecated: Please use types in: cloud.google.com/go/bigquery/analyticshub/apiv1/analyticshubpb
type DeleteDataExchangeRequest = src.DeleteDataExchangeRequest

// Message for deleting a listing.
//
// Deprecated: Please use types in: cloud.google.com/go/bigquery/analyticshub/apiv1/analyticshubpb
type DeleteListingRequest = src.DeleteListingRequest

// Defines the destination bigquery dataset.
//
// Deprecated: Please use types in: cloud.google.com/go/bigquery/analyticshub/apiv1/analyticshubpb
type DestinationDataset = src.DestinationDataset

// Contains the reference that identifies a destination bigquery dataset.
//
// Deprecated: Please use types in: cloud.google.com/go/bigquery/analyticshub/apiv1/analyticshubpb
type DestinationDatasetReference = src.DestinationDatasetReference

// Message for getting a data exchange.
//
// Deprecated: Please use types in: cloud.google.com/go/bigquery/analyticshub/apiv1/analyticshubpb
type GetDataExchangeRequest = src.GetDataExchangeRequest

// Message for getting a listing.
//
// Deprecated: Please use types in: cloud.google.com/go/bigquery/analyticshub/apiv1/analyticshubpb
type GetListingRequest = src.GetListingRequest

// Message for requesting the list of data exchanges.
//
// Deprecated: Please use types in: cloud.google.com/go/bigquery/analyticshub/apiv1/analyticshubpb
type ListDataExchangesRequest = src.ListDataExchangesRequest

// Message for response to the list of data exchanges.
//
// Deprecated: Please use types in: cloud.google.com/go/bigquery/analyticshub/apiv1/analyticshubpb
type ListDataExchangesResponse = src.ListDataExchangesResponse

// Message for requesting the list of listings.
//
// Deprecated: Please use types in: cloud.google.com/go/bigquery/analyticshub/apiv1/analyticshubpb
type ListListingsRequest = src.ListListingsRequest

// Message for response to the list of Listings.
//
// Deprecated: Please use types in: cloud.google.com/go/bigquery/analyticshub/apiv1/analyticshubpb
type ListListingsResponse = src.ListListingsResponse

// Message for requesting the list of data exchanges from projects in an
// organization and location.
//
// Deprecated: Please use types in: cloud.google.com/go/bigquery/analyticshub/apiv1/analyticshubpb
type ListOrgDataExchangesRequest = src.ListOrgDataExchangesRequest

// Message for response to listing data exchanges in an organization and
// location.
//
// Deprecated: Please use types in: cloud.google.com/go/bigquery/analyticshub/apiv1/analyticshubpb
type ListOrgDataExchangesResponse = src.ListOrgDataExchangesResponse

// A listing is what gets published into a data exchange that a subscriber can
// subscribe to. It contains a reference to the data source along with
// descriptive information that will help subscribers find and subscribe the
// data.
//
// Deprecated: Please use types in: cloud.google.com/go/bigquery/analyticshub/apiv1/analyticshubpb
type Listing = src.Listing

// A reference to a shared dataset. It is an existing BigQuery dataset with a
// collection of objects such as tables and views that you want to share with
// subscribers. When subscriber's subscribe to a listing, Analytics Hub creates
// a linked dataset in the subscriber's project. A Linked dataset is an opaque,
// read-only BigQuery dataset that serves as a _symbolic link_ to a shared
// dataset.
//
// Deprecated: Please use types in: cloud.google.com/go/bigquery/analyticshub/apiv1/analyticshubpb
type Listing_BigQueryDatasetSource = src.Listing_BigQueryDatasetSource
type Listing_BigqueryDataset = src.Listing_BigqueryDataset

// Listing categories.
//
// Deprecated: Please use types in: cloud.google.com/go/bigquery/analyticshub/apiv1/analyticshubpb
type Listing_Category = src.Listing_Category

// State of the listing.
//
// Deprecated: Please use types in: cloud.google.com/go/bigquery/analyticshub/apiv1/analyticshubpb
type Listing_State = src.Listing_State

// Contains details of the listing publisher.
//
// Deprecated: Please use types in: cloud.google.com/go/bigquery/analyticshub/apiv1/analyticshubpb
type Publisher = src.Publisher

// Message for subscribing to a listing.
//
// Deprecated: Please use types in: cloud.google.com/go/bigquery/analyticshub/apiv1/analyticshubpb
type SubscribeListingRequest = src.SubscribeListingRequest
type SubscribeListingRequest_DestinationDataset = src.SubscribeListingRequest_DestinationDataset

// Message for response when you subscribe to a listing.
//
// Deprecated: Please use types in: cloud.google.com/go/bigquery/analyticshub/apiv1/analyticshubpb
type SubscribeListingResponse = src.SubscribeListingResponse

// UnimplementedAnalyticsHubServiceServer can be embedded to have forward
// compatible implementations.
//
// Deprecated: Please use types in: cloud.google.com/go/bigquery/analyticshub/apiv1/analyticshubpb
type UnimplementedAnalyticsHubServiceServer = src.UnimplementedAnalyticsHubServiceServer

// Message for updating a data exchange.
//
// Deprecated: Please use types in: cloud.google.com/go/bigquery/analyticshub/apiv1/analyticshubpb
type UpdateDataExchangeRequest = src.UpdateDataExchangeRequest

// Message for updating a Listing.
//
// Deprecated: Please use types in: cloud.google.com/go/bigquery/analyticshub/apiv1/analyticshubpb
type UpdateListingRequest = src.UpdateListingRequest

// Deprecated: Please use funcs in: cloud.google.com/go/bigquery/analyticshub/apiv1/analyticshubpb
func NewAnalyticsHubServiceClient(cc grpc.ClientConnInterface) AnalyticsHubServiceClient {
	return src.NewAnalyticsHubServiceClient(cc)
}

// Deprecated: Please use funcs in: cloud.google.com/go/bigquery/analyticshub/apiv1/analyticshubpb
func RegisterAnalyticsHubServiceServer(s *grpc.Server, srv AnalyticsHubServiceServer) {
	src.RegisterAnalyticsHubServiceServer(s, srv)
}