package network

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"net/http"
)

// LoadBalancerLoadBalancingRulesClient is the network Client
type LoadBalancerLoadBalancingRulesClient struct {
	ManagementClient
}

// NewLoadBalancerLoadBalancingRulesClient creates an instance of the LoadBalancerLoadBalancingRulesClient client.
func NewLoadBalancerLoadBalancingRulesClient(subscriptionID string) LoadBalancerLoadBalancingRulesClient {
	return NewLoadBalancerLoadBalancingRulesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewLoadBalancerLoadBalancingRulesClientWithBaseURI creates an instance of the LoadBalancerLoadBalancingRulesClient
// client.
func NewLoadBalancerLoadBalancingRulesClientWithBaseURI(baseURI string, subscriptionID string) LoadBalancerLoadBalancingRulesClient {
	return LoadBalancerLoadBalancingRulesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Get gets the specified load balancer load balancing rule.
//
// resourceGroupName is the name of the resource group. loadBalancerName is the name of the load balancer.
// loadBalancingRuleName is the name of the load balancing rule.
func (client LoadBalancerLoadBalancingRulesClient) Get(resourceGroupName string, loadBalancerName string, loadBalancingRuleName string) (result LoadBalancingRule, err error) {
	req, err := client.GetPreparer(resourceGroupName, loadBalancerName, loadBalancingRuleName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.LoadBalancerLoadBalancingRulesClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.LoadBalancerLoadBalancingRulesClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.LoadBalancerLoadBalancingRulesClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client LoadBalancerLoadBalancingRulesClient) GetPreparer(resourceGroupName string, loadBalancerName string, loadBalancingRuleName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"loadBalancerName":      autorest.Encode("path", loadBalancerName),
		"loadBalancingRuleName": autorest.Encode("path", loadBalancingRuleName),
		"resourceGroupName":     autorest.Encode("path", resourceGroupName),
		"subscriptionId":        autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers/{loadBalancerName}/loadBalancingRules/{loadBalancingRuleName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client LoadBalancerLoadBalancingRulesClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client LoadBalancerLoadBalancingRulesClient) GetResponder(resp *http.Response) (result LoadBalancingRule, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List gets all the load balancing rules in a load balancer.
//
// resourceGroupName is the name of the resource group. loadBalancerName is the name of the load balancer.
func (client LoadBalancerLoadBalancingRulesClient) List(resourceGroupName string, loadBalancerName string) (result LoadBalancerLoadBalancingRuleListResult, err error) {
	req, err := client.ListPreparer(resourceGroupName, loadBalancerName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.LoadBalancerLoadBalancingRulesClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.LoadBalancerLoadBalancingRulesClient", "List", resp, "Failure sending request")
		return
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.LoadBalancerLoadBalancingRulesClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client LoadBalancerLoadBalancingRulesClient) ListPreparer(resourceGroupName string, loadBalancerName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"loadBalancerName":  autorest.Encode("path", loadBalancerName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers/{loadBalancerName}/loadBalancingRules", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client LoadBalancerLoadBalancingRulesClient) ListSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client LoadBalancerLoadBalancingRulesClient) ListResponder(resp *http.Response) (result LoadBalancerLoadBalancingRuleListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListNextResults retrieves the next set of results, if any.
func (client LoadBalancerLoadBalancingRulesClient) ListNextResults(lastResults LoadBalancerLoadBalancingRuleListResult) (result LoadBalancerLoadBalancingRuleListResult, err error) {
	req, err := lastResults.LoadBalancerLoadBalancingRuleListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network.LoadBalancerLoadBalancingRulesClient", "List", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network.LoadBalancerLoadBalancingRulesClient", "List", resp, "Failure sending next results request")
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.LoadBalancerLoadBalancingRulesClient", "List", resp, "Failure responding to next results request")
	}

	return
}

// ListComplete gets all elements from the list without paging.
func (client LoadBalancerLoadBalancingRulesClient) ListComplete(resourceGroupName string, loadBalancerName string, cancel <-chan struct{}) (<-chan LoadBalancingRule, <-chan error) {
	resultChan := make(chan LoadBalancingRule)
	errChan := make(chan error, 1)
	go func() {
		defer func() {
			close(resultChan)
			close(errChan)
		}()
		list, err := client.List(resourceGroupName, loadBalancerName)
		if err != nil {
			errChan <- err
			return
		}
		if list.Value != nil {
			for _, item := range *list.Value {
				select {
				case <-cancel:
					return
				case resultChan <- item:
					// Intentionally left blank
				}
			}
		}
		for list.NextLink != nil {
			list, err = client.ListNextResults(list)
			if err != nil {
				errChan <- err
				return
			}
			if list.Value != nil {
				for _, item := range *list.Value {
					select {
					case <-cancel:
						return
					case resultChan <- item:
						// Intentionally left blank
					}
				}
			}
		}
	}()
	return resultChan, errChan
}
