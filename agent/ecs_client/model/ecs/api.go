// Copyright 2014-2017 Amazon.com, Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//	http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

package ecs

import (
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awsutil"
	"github.com/aws/aws-sdk-go/aws/request"
)

const opCreateCluster = "CreateCluster"

// CreateClusterRequest generates a "aws/request.Request" representing the
// client's request for the CreateCluster operation. The "output" return
// value can be used to capture response data after the request's "Send" method
// is called.
//
// See CreateCluster for usage and error information.
//
// Creating a request object using this method should be used when you want to inject
// custom logic into the request's lifecycle using a custom handler, or if you want to
// access properties on the request object before or after sending the request. If
// you just want the service response, call the CreateCluster method directly
// instead.
//
// Note: You must call the "Send" method on the returned request object in order
// to execute the request.
//
//    // Example sending a request using the CreateClusterRequest method.
//    req, resp := client.CreateClusterRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ECS) CreateClusterRequest(input *CreateClusterInput) (req *request.Request, output *CreateClusterOutput) {
	op := &request.Operation{
		Name:       opCreateCluster,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateClusterInput{}
	}

	output = &CreateClusterOutput{}
	req = c.newRequest(op, input, output)
	return
}

// CreateCluster API operation for Amazon EC2 Container Service.
//
// Creates a new Amazon ECS cluster. By default, your account will receive a
// default cluster when you launch your first container instance. However, you
// can create your own cluster with a unique name with the CreateCluster action.
//
// Returns awserr.Error for service API and SDK errors. Use runtime type assertions
// with awserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the AWS API reference guide for Amazon EC2 Container Service's
// API operation CreateCluster for usage and error information.
//
// Returned Error Codes:
//   * ErrCodeServerException "ServerException"
//   These errors are usually caused by a server-side issue.
//
//   * ErrCodeClientException "ClientException"
//   These errors are usually caused by something the client did, such as use
//   an action or resource on behalf of a user that doesn't have permission to
//   use the action or resource, or specify an identifier that is not valid.
//
//   * ErrCodeInvalidParameterException "InvalidParameterException"
//   The specified parameter is invalid. Review the available parameters for the
//   API request.
//
func (c *ECS) CreateCluster(input *CreateClusterInput) (*CreateClusterOutput, error) {
	req, out := c.CreateClusterRequest(input)
	return out, req.Send()
}

// CreateClusterWithContext is the same as CreateCluster with the addition of
// the ability to pass a context and additional request options.
//
// See CreateCluster for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ECS) CreateClusterWithContext(ctx aws.Context, input *CreateClusterInput, opts ...request.Option) (*CreateClusterOutput, error) {
	req, out := c.CreateClusterRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opCreateService = "CreateService"

// CreateServiceRequest generates a "aws/request.Request" representing the
// client's request for the CreateService operation. The "output" return
// value can be used to capture response data after the request's "Send" method
// is called.
//
// See CreateService for usage and error information.
//
// Creating a request object using this method should be used when you want to inject
// custom logic into the request's lifecycle using a custom handler, or if you want to
// access properties on the request object before or after sending the request. If
// you just want the service response, call the CreateService method directly
// instead.
//
// Note: You must call the "Send" method on the returned request object in order
// to execute the request.
//
//    // Example sending a request using the CreateServiceRequest method.
//    req, resp := client.CreateServiceRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ECS) CreateServiceRequest(input *CreateServiceInput) (req *request.Request, output *CreateServiceOutput) {
	op := &request.Operation{
		Name:       opCreateService,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateServiceInput{}
	}

	output = &CreateServiceOutput{}
	req = c.newRequest(op, input, output)
	return
}

// CreateService API operation for Amazon EC2 Container Service.
//
// Runs and maintains a desired number of tasks from a specified task definition.
// If the number of tasks running in a service drops below desiredCount, Amazon
// ECS will spawn another instantiation of the task in the specified cluster.
//
// Returns awserr.Error for service API and SDK errors. Use runtime type assertions
// with awserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the AWS API reference guide for Amazon EC2 Container Service's
// API operation CreateService for usage and error information.
//
// Returned Error Codes:
//   * ErrCodeServerException "ServerException"
//   These errors are usually caused by a server-side issue.
//
//   * ErrCodeClientException "ClientException"
//   These errors are usually caused by something the client did, such as use
//   an action or resource on behalf of a user that doesn't have permission to
//   use the action or resource, or specify an identifier that is not valid.
//
//   * ErrCodeInvalidParameterException "InvalidParameterException"
//   The specified parameter is invalid. Review the available parameters for the
//   API request.
//
//   * ErrCodeClusterNotFoundException "ClusterNotFoundException"
//   The specified cluster could not be found. You can view your available clusters
//   with ListClusters. Amazon ECS clusters are region-specific.
//
func (c *ECS) CreateService(input *CreateServiceInput) (*CreateServiceOutput, error) {
	req, out := c.CreateServiceRequest(input)
	return out, req.Send()
}

// CreateServiceWithContext is the same as CreateService with the addition of
// the ability to pass a context and additional request options.
//
// See CreateService for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ECS) CreateServiceWithContext(ctx aws.Context, input *CreateServiceInput, opts ...request.Option) (*CreateServiceOutput, error) {
	req, out := c.CreateServiceRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDeleteAttributes = "DeleteAttributes"

// DeleteAttributesRequest generates a "aws/request.Request" representing the
// client's request for the DeleteAttributes operation. The "output" return
// value can be used to capture response data after the request's "Send" method
// is called.
//
// See DeleteAttributes for usage and error information.
//
// Creating a request object using this method should be used when you want to inject
// custom logic into the request's lifecycle using a custom handler, or if you want to
// access properties on the request object before or after sending the request. If
// you just want the service response, call the DeleteAttributes method directly
// instead.
//
// Note: You must call the "Send" method on the returned request object in order
// to execute the request.
//
//    // Example sending a request using the DeleteAttributesRequest method.
//    req, resp := client.DeleteAttributesRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ECS) DeleteAttributesRequest(input *DeleteAttributesInput) (req *request.Request, output *DeleteAttributesOutput) {
	op := &request.Operation{
		Name:       opDeleteAttributes,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteAttributesInput{}
	}

	output = &DeleteAttributesOutput{}
	req = c.newRequest(op, input, output)
	return
}

// DeleteAttributes API operation for Amazon EC2 Container Service.
//
// Returns awserr.Error for service API and SDK errors. Use runtime type assertions
// with awserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the AWS API reference guide for Amazon EC2 Container Service's
// API operation DeleteAttributes for usage and error information.
//
// Returned Error Codes:
//   * ErrCodeClusterNotFoundException "ClusterNotFoundException"
//   The specified cluster could not be found. You can view your available clusters
//   with ListClusters. Amazon ECS clusters are region-specific.
//
//   * ErrCodeTargetNotFoundException "TargetNotFoundException"
//
//   * ErrCodeInvalidParameterException "InvalidParameterException"
//   The specified parameter is invalid. Review the available parameters for the
//   API request.
//
func (c *ECS) DeleteAttributes(input *DeleteAttributesInput) (*DeleteAttributesOutput, error) {
	req, out := c.DeleteAttributesRequest(input)
	return out, req.Send()
}

// DeleteAttributesWithContext is the same as DeleteAttributes with the addition of
// the ability to pass a context and additional request options.
//
// See DeleteAttributes for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ECS) DeleteAttributesWithContext(ctx aws.Context, input *DeleteAttributesInput, opts ...request.Option) (*DeleteAttributesOutput, error) {
	req, out := c.DeleteAttributesRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDeleteCluster = "DeleteCluster"

// DeleteClusterRequest generates a "aws/request.Request" representing the
// client's request for the DeleteCluster operation. The "output" return
// value can be used to capture response data after the request's "Send" method
// is called.
//
// See DeleteCluster for usage and error information.
//
// Creating a request object using this method should be used when you want to inject
// custom logic into the request's lifecycle using a custom handler, or if you want to
// access properties on the request object before or after sending the request. If
// you just want the service response, call the DeleteCluster method directly
// instead.
//
// Note: You must call the "Send" method on the returned request object in order
// to execute the request.
//
//    // Example sending a request using the DeleteClusterRequest method.
//    req, resp := client.DeleteClusterRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ECS) DeleteClusterRequest(input *DeleteClusterInput) (req *request.Request, output *DeleteClusterOutput) {
	op := &request.Operation{
		Name:       opDeleteCluster,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteClusterInput{}
	}

	output = &DeleteClusterOutput{}
	req = c.newRequest(op, input, output)
	return
}

// DeleteCluster API operation for Amazon EC2 Container Service.
//
// Deletes the specified cluster. You must deregister all container instances
// from this cluster before you may delete it. You can list the container instances
// in a cluster with ListContainerInstances and deregister them with DeregisterContainerInstance.
//
// Returns awserr.Error for service API and SDK errors. Use runtime type assertions
// with awserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the AWS API reference guide for Amazon EC2 Container Service's
// API operation DeleteCluster for usage and error information.
//
// Returned Error Codes:
//   * ErrCodeServerException "ServerException"
//   These errors are usually caused by a server-side issue.
//
//   * ErrCodeClientException "ClientException"
//   These errors are usually caused by something the client did, such as use
//   an action or resource on behalf of a user that doesn't have permission to
//   use the action or resource, or specify an identifier that is not valid.
//
//   * ErrCodeInvalidParameterException "InvalidParameterException"
//   The specified parameter is invalid. Review the available parameters for the
//   API request.
//
//   * ErrCodeClusterNotFoundException "ClusterNotFoundException"
//   The specified cluster could not be found. You can view your available clusters
//   with ListClusters. Amazon ECS clusters are region-specific.
//
//   * ErrCodeClusterContainsContainerInstancesException "ClusterContainsContainerInstancesException"
//   You cannot delete a cluster that has registered container instances. You
//   must first deregister the container instances before you can delete the cluster.
//   For more information, see DeregisterContainerInstance.
//
//   * ErrCodeClusterContainsServicesException "ClusterContainsServicesException"
//   You cannot delete a cluster that contains services. You must first update
//   the service to reduce its desired task count to 0 and then delete the service.
//   For more information, see UpdateService and DeleteService.
//
func (c *ECS) DeleteCluster(input *DeleteClusterInput) (*DeleteClusterOutput, error) {
	req, out := c.DeleteClusterRequest(input)
	return out, req.Send()
}

// DeleteClusterWithContext is the same as DeleteCluster with the addition of
// the ability to pass a context and additional request options.
//
// See DeleteCluster for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ECS) DeleteClusterWithContext(ctx aws.Context, input *DeleteClusterInput, opts ...request.Option) (*DeleteClusterOutput, error) {
	req, out := c.DeleteClusterRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDeleteService = "DeleteService"

// DeleteServiceRequest generates a "aws/request.Request" representing the
// client's request for the DeleteService operation. The "output" return
// value can be used to capture response data after the request's "Send" method
// is called.
//
// See DeleteService for usage and error information.
//
// Creating a request object using this method should be used when you want to inject
// custom logic into the request's lifecycle using a custom handler, or if you want to
// access properties on the request object before or after sending the request. If
// you just want the service response, call the DeleteService method directly
// instead.
//
// Note: You must call the "Send" method on the returned request object in order
// to execute the request.
//
//    // Example sending a request using the DeleteServiceRequest method.
//    req, resp := client.DeleteServiceRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ECS) DeleteServiceRequest(input *DeleteServiceInput) (req *request.Request, output *DeleteServiceOutput) {
	op := &request.Operation{
		Name:       opDeleteService,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteServiceInput{}
	}

	output = &DeleteServiceOutput{}
	req = c.newRequest(op, input, output)
	return
}

// DeleteService API operation for Amazon EC2 Container Service.
//
// Deletes a specified service within a cluster.
//
// Returns awserr.Error for service API and SDK errors. Use runtime type assertions
// with awserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the AWS API reference guide for Amazon EC2 Container Service's
// API operation DeleteService for usage and error information.
//
// Returned Error Codes:
//   * ErrCodeServerException "ServerException"
//   These errors are usually caused by a server-side issue.
//
//   * ErrCodeClientException "ClientException"
//   These errors are usually caused by something the client did, such as use
//   an action or resource on behalf of a user that doesn't have permission to
//   use the action or resource, or specify an identifier that is not valid.
//
//   * ErrCodeInvalidParameterException "InvalidParameterException"
//   The specified parameter is invalid. Review the available parameters for the
//   API request.
//
//   * ErrCodeClusterNotFoundException "ClusterNotFoundException"
//   The specified cluster could not be found. You can view your available clusters
//   with ListClusters. Amazon ECS clusters are region-specific.
//
//   * ErrCodeServiceNotFoundException "ServiceNotFoundException"
//   The specified service could not be found. You can view your available services
//   with ListServices. Amazon ECS services are cluster-specific and region-specific.
//
func (c *ECS) DeleteService(input *DeleteServiceInput) (*DeleteServiceOutput, error) {
	req, out := c.DeleteServiceRequest(input)
	return out, req.Send()
}

// DeleteServiceWithContext is the same as DeleteService with the addition of
// the ability to pass a context and additional request options.
//
// See DeleteService for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ECS) DeleteServiceWithContext(ctx aws.Context, input *DeleteServiceInput, opts ...request.Option) (*DeleteServiceOutput, error) {
	req, out := c.DeleteServiceRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDeregisterContainerInstance = "DeregisterContainerInstance"

// DeregisterContainerInstanceRequest generates a "aws/request.Request" representing the
// client's request for the DeregisterContainerInstance operation. The "output" return
// value can be used to capture response data after the request's "Send" method
// is called.
//
// See DeregisterContainerInstance for usage and error information.
//
// Creating a request object using this method should be used when you want to inject
// custom logic into the request's lifecycle using a custom handler, or if you want to
// access properties on the request object before or after sending the request. If
// you just want the service response, call the DeregisterContainerInstance method directly
// instead.
//
// Note: You must call the "Send" method on the returned request object in order
// to execute the request.
//
//    // Example sending a request using the DeregisterContainerInstanceRequest method.
//    req, resp := client.DeregisterContainerInstanceRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ECS) DeregisterContainerInstanceRequest(input *DeregisterContainerInstanceInput) (req *request.Request, output *DeregisterContainerInstanceOutput) {
	op := &request.Operation{
		Name:       opDeregisterContainerInstance,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeregisterContainerInstanceInput{}
	}

	output = &DeregisterContainerInstanceOutput{}
	req = c.newRequest(op, input, output)
	return
}

// DeregisterContainerInstance API operation for Amazon EC2 Container Service.
//
// Deregisters an Amazon ECS container instance from the specified cluster.
// This instance will no longer be available to run tasks.
//
// If you intend to use the container instance for some other purpose after
// deregistration, you should stop all of the tasks running on the container
// instance before deregistration to avoid any orphaned tasks from consuming
// resources.
//
// Deregistering a container instance removes the instance from a cluster, but
// it does not terminate the EC2 instance; if you are finished using the instance,
// be sure to terminate it in the Amazon EC2 console to stop billing.
//
// When you terminate a container instance, it is automatically deregistered
// from your cluster.
//
// Returns awserr.Error for service API and SDK errors. Use runtime type assertions
// with awserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the AWS API reference guide for Amazon EC2 Container Service's
// API operation DeregisterContainerInstance for usage and error information.
//
// Returned Error Codes:
//   * ErrCodeServerException "ServerException"
//   These errors are usually caused by a server-side issue.
//
//   * ErrCodeClientException "ClientException"
//   These errors are usually caused by something the client did, such as use
//   an action or resource on behalf of a user that doesn't have permission to
//   use the action or resource, or specify an identifier that is not valid.
//
//   * ErrCodeInvalidParameterException "InvalidParameterException"
//   The specified parameter is invalid. Review the available parameters for the
//   API request.
//
//   * ErrCodeClusterNotFoundException "ClusterNotFoundException"
//   The specified cluster could not be found. You can view your available clusters
//   with ListClusters. Amazon ECS clusters are region-specific.
//
func (c *ECS) DeregisterContainerInstance(input *DeregisterContainerInstanceInput) (*DeregisterContainerInstanceOutput, error) {
	req, out := c.DeregisterContainerInstanceRequest(input)
	return out, req.Send()
}

// DeregisterContainerInstanceWithContext is the same as DeregisterContainerInstance with the addition of
// the ability to pass a context and additional request options.
//
// See DeregisterContainerInstance for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ECS) DeregisterContainerInstanceWithContext(ctx aws.Context, input *DeregisterContainerInstanceInput, opts ...request.Option) (*DeregisterContainerInstanceOutput, error) {
	req, out := c.DeregisterContainerInstanceRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDeregisterTaskDefinition = "DeregisterTaskDefinition"

// DeregisterTaskDefinitionRequest generates a "aws/request.Request" representing the
// client's request for the DeregisterTaskDefinition operation. The "output" return
// value can be used to capture response data after the request's "Send" method
// is called.
//
// See DeregisterTaskDefinition for usage and error information.
//
// Creating a request object using this method should be used when you want to inject
// custom logic into the request's lifecycle using a custom handler, or if you want to
// access properties on the request object before or after sending the request. If
// you just want the service response, call the DeregisterTaskDefinition method directly
// instead.
//
// Note: You must call the "Send" method on the returned request object in order
// to execute the request.
//
//    // Example sending a request using the DeregisterTaskDefinitionRequest method.
//    req, resp := client.DeregisterTaskDefinitionRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ECS) DeregisterTaskDefinitionRequest(input *DeregisterTaskDefinitionInput) (req *request.Request, output *DeregisterTaskDefinitionOutput) {
	op := &request.Operation{
		Name:       opDeregisterTaskDefinition,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeregisterTaskDefinitionInput{}
	}

	output = &DeregisterTaskDefinitionOutput{}
	req = c.newRequest(op, input, output)
	return
}

// DeregisterTaskDefinition API operation for Amazon EC2 Container Service.
//
// Deregisters the specified task definition by family and revision. Upon deregistration,
// the task definition is marked as INACTIVE. Existing tasks and services that
// reference an INACTIVE task definition continue to run without disruption.
// Existing services that reference an INACTIVE task definition can still scale
// up or down by modifying the service's desired count.
//
// You cannot use an INACTIVE task definition to run new tasks or create new
// services, and you cannot update an existing service to reference an INACTIVE
// task definition (although there may be up to a 10 minute window following
// deregistration where these restrictions have not yet taken effect).
//
// Returns awserr.Error for service API and SDK errors. Use runtime type assertions
// with awserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the AWS API reference guide for Amazon EC2 Container Service's
// API operation DeregisterTaskDefinition for usage and error information.
//
// Returned Error Codes:
//   * ErrCodeServerException "ServerException"
//   These errors are usually caused by a server-side issue.
//
//   * ErrCodeClientException "ClientException"
//   These errors are usually caused by something the client did, such as use
//   an action or resource on behalf of a user that doesn't have permission to
//   use the action or resource, or specify an identifier that is not valid.
//
//   * ErrCodeInvalidParameterException "InvalidParameterException"
//   The specified parameter is invalid. Review the available parameters for the
//   API request.
//
func (c *ECS) DeregisterTaskDefinition(input *DeregisterTaskDefinitionInput) (*DeregisterTaskDefinitionOutput, error) {
	req, out := c.DeregisterTaskDefinitionRequest(input)
	return out, req.Send()
}

// DeregisterTaskDefinitionWithContext is the same as DeregisterTaskDefinition with the addition of
// the ability to pass a context and additional request options.
//
// See DeregisterTaskDefinition for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ECS) DeregisterTaskDefinitionWithContext(ctx aws.Context, input *DeregisterTaskDefinitionInput, opts ...request.Option) (*DeregisterTaskDefinitionOutput, error) {
	req, out := c.DeregisterTaskDefinitionRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribeClusters = "DescribeClusters"

// DescribeClustersRequest generates a "aws/request.Request" representing the
// client's request for the DescribeClusters operation. The "output" return
// value can be used to capture response data after the request's "Send" method
// is called.
//
// See DescribeClusters for usage and error information.
//
// Creating a request object using this method should be used when you want to inject
// custom logic into the request's lifecycle using a custom handler, or if you want to
// access properties on the request object before or after sending the request. If
// you just want the service response, call the DescribeClusters method directly
// instead.
//
// Note: You must call the "Send" method on the returned request object in order
// to execute the request.
//
//    // Example sending a request using the DescribeClustersRequest method.
//    req, resp := client.DescribeClustersRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ECS) DescribeClustersRequest(input *DescribeClustersInput) (req *request.Request, output *DescribeClustersOutput) {
	op := &request.Operation{
		Name:       opDescribeClusters,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeClustersInput{}
	}

	output = &DescribeClustersOutput{}
	req = c.newRequest(op, input, output)
	return
}

// DescribeClusters API operation for Amazon EC2 Container Service.
//
// Describes one or more of your clusters.
//
// Returns awserr.Error for service API and SDK errors. Use runtime type assertions
// with awserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the AWS API reference guide for Amazon EC2 Container Service's
// API operation DescribeClusters for usage and error information.
//
// Returned Error Codes:
//   * ErrCodeServerException "ServerException"
//   These errors are usually caused by a server-side issue.
//
//   * ErrCodeClientException "ClientException"
//   These errors are usually caused by something the client did, such as use
//   an action or resource on behalf of a user that doesn't have permission to
//   use the action or resource, or specify an identifier that is not valid.
//
//   * ErrCodeInvalidParameterException "InvalidParameterException"
//   The specified parameter is invalid. Review the available parameters for the
//   API request.
//
func (c *ECS) DescribeClusters(input *DescribeClustersInput) (*DescribeClustersOutput, error) {
	req, out := c.DescribeClustersRequest(input)
	return out, req.Send()
}

// DescribeClustersWithContext is the same as DescribeClusters with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeClusters for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ECS) DescribeClustersWithContext(ctx aws.Context, input *DescribeClustersInput, opts ...request.Option) (*DescribeClustersOutput, error) {
	req, out := c.DescribeClustersRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribeContainerInstances = "DescribeContainerInstances"

// DescribeContainerInstancesRequest generates a "aws/request.Request" representing the
// client's request for the DescribeContainerInstances operation. The "output" return
// value can be used to capture response data after the request's "Send" method
// is called.
//
// See DescribeContainerInstances for usage and error information.
//
// Creating a request object using this method should be used when you want to inject
// custom logic into the request's lifecycle using a custom handler, or if you want to
// access properties on the request object before or after sending the request. If
// you just want the service response, call the DescribeContainerInstances method directly
// instead.
//
// Note: You must call the "Send" method on the returned request object in order
// to execute the request.
//
//    // Example sending a request using the DescribeContainerInstancesRequest method.
//    req, resp := client.DescribeContainerInstancesRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ECS) DescribeContainerInstancesRequest(input *DescribeContainerInstancesInput) (req *request.Request, output *DescribeContainerInstancesOutput) {
	op := &request.Operation{
		Name:       opDescribeContainerInstances,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeContainerInstancesInput{}
	}

	output = &DescribeContainerInstancesOutput{}
	req = c.newRequest(op, input, output)
	return
}

// DescribeContainerInstances API operation for Amazon EC2 Container Service.
//
// Describes Amazon EC2 Container Service container instances. Returns metadata
// about registered and remaining resources on each container instance requested.
//
// Returns awserr.Error for service API and SDK errors. Use runtime type assertions
// with awserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the AWS API reference guide for Amazon EC2 Container Service's
// API operation DescribeContainerInstances for usage and error information.
//
// Returned Error Codes:
//   * ErrCodeServerException "ServerException"
//   These errors are usually caused by a server-side issue.
//
//   * ErrCodeClientException "ClientException"
//   These errors are usually caused by something the client did, such as use
//   an action or resource on behalf of a user that doesn't have permission to
//   use the action or resource, or specify an identifier that is not valid.
//
//   * ErrCodeInvalidParameterException "InvalidParameterException"
//   The specified parameter is invalid. Review the available parameters for the
//   API request.
//
//   * ErrCodeClusterNotFoundException "ClusterNotFoundException"
//   The specified cluster could not be found. You can view your available clusters
//   with ListClusters. Amazon ECS clusters are region-specific.
//
func (c *ECS) DescribeContainerInstances(input *DescribeContainerInstancesInput) (*DescribeContainerInstancesOutput, error) {
	req, out := c.DescribeContainerInstancesRequest(input)
	return out, req.Send()
}

// DescribeContainerInstancesWithContext is the same as DescribeContainerInstances with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeContainerInstances for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ECS) DescribeContainerInstancesWithContext(ctx aws.Context, input *DescribeContainerInstancesInput, opts ...request.Option) (*DescribeContainerInstancesOutput, error) {
	req, out := c.DescribeContainerInstancesRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribeServices = "DescribeServices"

// DescribeServicesRequest generates a "aws/request.Request" representing the
// client's request for the DescribeServices operation. The "output" return
// value can be used to capture response data after the request's "Send" method
// is called.
//
// See DescribeServices for usage and error information.
//
// Creating a request object using this method should be used when you want to inject
// custom logic into the request's lifecycle using a custom handler, or if you want to
// access properties on the request object before or after sending the request. If
// you just want the service response, call the DescribeServices method directly
// instead.
//
// Note: You must call the "Send" method on the returned request object in order
// to execute the request.
//
//    // Example sending a request using the DescribeServicesRequest method.
//    req, resp := client.DescribeServicesRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ECS) DescribeServicesRequest(input *DescribeServicesInput) (req *request.Request, output *DescribeServicesOutput) {
	op := &request.Operation{
		Name:       opDescribeServices,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeServicesInput{}
	}

	output = &DescribeServicesOutput{}
	req = c.newRequest(op, input, output)
	return
}

// DescribeServices API operation for Amazon EC2 Container Service.
//
// Describes the specified services running in your cluster.
//
// Returns awserr.Error for service API and SDK errors. Use runtime type assertions
// with awserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the AWS API reference guide for Amazon EC2 Container Service's
// API operation DescribeServices for usage and error information.
//
// Returned Error Codes:
//   * ErrCodeServerException "ServerException"
//   These errors are usually caused by a server-side issue.
//
//   * ErrCodeClientException "ClientException"
//   These errors are usually caused by something the client did, such as use
//   an action or resource on behalf of a user that doesn't have permission to
//   use the action or resource, or specify an identifier that is not valid.
//
//   * ErrCodeInvalidParameterException "InvalidParameterException"
//   The specified parameter is invalid. Review the available parameters for the
//   API request.
//
//   * ErrCodeClusterNotFoundException "ClusterNotFoundException"
//   The specified cluster could not be found. You can view your available clusters
//   with ListClusters. Amazon ECS clusters are region-specific.
//
func (c *ECS) DescribeServices(input *DescribeServicesInput) (*DescribeServicesOutput, error) {
	req, out := c.DescribeServicesRequest(input)
	return out, req.Send()
}

// DescribeServicesWithContext is the same as DescribeServices with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeServices for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ECS) DescribeServicesWithContext(ctx aws.Context, input *DescribeServicesInput, opts ...request.Option) (*DescribeServicesOutput, error) {
	req, out := c.DescribeServicesRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribeTaskDefinition = "DescribeTaskDefinition"

// DescribeTaskDefinitionRequest generates a "aws/request.Request" representing the
// client's request for the DescribeTaskDefinition operation. The "output" return
// value can be used to capture response data after the request's "Send" method
// is called.
//
// See DescribeTaskDefinition for usage and error information.
//
// Creating a request object using this method should be used when you want to inject
// custom logic into the request's lifecycle using a custom handler, or if you want to
// access properties on the request object before or after sending the request. If
// you just want the service response, call the DescribeTaskDefinition method directly
// instead.
//
// Note: You must call the "Send" method on the returned request object in order
// to execute the request.
//
//    // Example sending a request using the DescribeTaskDefinitionRequest method.
//    req, resp := client.DescribeTaskDefinitionRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ECS) DescribeTaskDefinitionRequest(input *DescribeTaskDefinitionInput) (req *request.Request, output *DescribeTaskDefinitionOutput) {
	op := &request.Operation{
		Name:       opDescribeTaskDefinition,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeTaskDefinitionInput{}
	}

	output = &DescribeTaskDefinitionOutput{}
	req = c.newRequest(op, input, output)
	return
}

// DescribeTaskDefinition API operation for Amazon EC2 Container Service.
//
// Describes a task definition. You can specify a family and revision to find
// information on a specific task definition, or you can simply specify the
// family to find the latest ACTIVE revision in that family.
//
// You can only describe INACTIVE task definitions while an active task or service
// references them.
//
// Returns awserr.Error for service API and SDK errors. Use runtime type assertions
// with awserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the AWS API reference guide for Amazon EC2 Container Service's
// API operation DescribeTaskDefinition for usage and error information.
//
// Returned Error Codes:
//   * ErrCodeServerException "ServerException"
//   These errors are usually caused by a server-side issue.
//
//   * ErrCodeClientException "ClientException"
//   These errors are usually caused by something the client did, such as use
//   an action or resource on behalf of a user that doesn't have permission to
//   use the action or resource, or specify an identifier that is not valid.
//
//   * ErrCodeInvalidParameterException "InvalidParameterException"
//   The specified parameter is invalid. Review the available parameters for the
//   API request.
//
func (c *ECS) DescribeTaskDefinition(input *DescribeTaskDefinitionInput) (*DescribeTaskDefinitionOutput, error) {
	req, out := c.DescribeTaskDefinitionRequest(input)
	return out, req.Send()
}

// DescribeTaskDefinitionWithContext is the same as DescribeTaskDefinition with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeTaskDefinition for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ECS) DescribeTaskDefinitionWithContext(ctx aws.Context, input *DescribeTaskDefinitionInput, opts ...request.Option) (*DescribeTaskDefinitionOutput, error) {
	req, out := c.DescribeTaskDefinitionRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribeTasks = "DescribeTasks"

// DescribeTasksRequest generates a "aws/request.Request" representing the
// client's request for the DescribeTasks operation. The "output" return
// value can be used to capture response data after the request's "Send" method
// is called.
//
// See DescribeTasks for usage and error information.
//
// Creating a request object using this method should be used when you want to inject
// custom logic into the request's lifecycle using a custom handler, or if you want to
// access properties on the request object before or after sending the request. If
// you just want the service response, call the DescribeTasks method directly
// instead.
//
// Note: You must call the "Send" method on the returned request object in order
// to execute the request.
//
//    // Example sending a request using the DescribeTasksRequest method.
//    req, resp := client.DescribeTasksRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ECS) DescribeTasksRequest(input *DescribeTasksInput) (req *request.Request, output *DescribeTasksOutput) {
	op := &request.Operation{
		Name:       opDescribeTasks,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeTasksInput{}
	}

	output = &DescribeTasksOutput{}
	req = c.newRequest(op, input, output)
	return
}

// DescribeTasks API operation for Amazon EC2 Container Service.
//
// Describes a specified task or tasks.
//
// Returns awserr.Error for service API and SDK errors. Use runtime type assertions
// with awserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the AWS API reference guide for Amazon EC2 Container Service's
// API operation DescribeTasks for usage and error information.
//
// Returned Error Codes:
//   * ErrCodeServerException "ServerException"
//   These errors are usually caused by a server-side issue.
//
//   * ErrCodeClientException "ClientException"
//   These errors are usually caused by something the client did, such as use
//   an action or resource on behalf of a user that doesn't have permission to
//   use the action or resource, or specify an identifier that is not valid.
//
//   * ErrCodeInvalidParameterException "InvalidParameterException"
//   The specified parameter is invalid. Review the available parameters for the
//   API request.
//
//   * ErrCodeClusterNotFoundException "ClusterNotFoundException"
//   The specified cluster could not be found. You can view your available clusters
//   with ListClusters. Amazon ECS clusters are region-specific.
//
func (c *ECS) DescribeTasks(input *DescribeTasksInput) (*DescribeTasksOutput, error) {
	req, out := c.DescribeTasksRequest(input)
	return out, req.Send()
}

// DescribeTasksWithContext is the same as DescribeTasks with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeTasks for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ECS) DescribeTasksWithContext(ctx aws.Context, input *DescribeTasksInput, opts ...request.Option) (*DescribeTasksOutput, error) {
	req, out := c.DescribeTasksRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDiscoverPollEndpoint = "DiscoverPollEndpoint"

// DiscoverPollEndpointRequest generates a "aws/request.Request" representing the
// client's request for the DiscoverPollEndpoint operation. The "output" return
// value can be used to capture response data after the request's "Send" method
// is called.
//
// See DiscoverPollEndpoint for usage and error information.
//
// Creating a request object using this method should be used when you want to inject
// custom logic into the request's lifecycle using a custom handler, or if you want to
// access properties on the request object before or after sending the request. If
// you just want the service response, call the DiscoverPollEndpoint method directly
// instead.
//
// Note: You must call the "Send" method on the returned request object in order
// to execute the request.
//
//    // Example sending a request using the DiscoverPollEndpointRequest method.
//    req, resp := client.DiscoverPollEndpointRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ECS) DiscoverPollEndpointRequest(input *DiscoverPollEndpointInput) (req *request.Request, output *DiscoverPollEndpointOutput) {
	op := &request.Operation{
		Name:       opDiscoverPollEndpoint,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DiscoverPollEndpointInput{}
	}

	output = &DiscoverPollEndpointOutput{}
	req = c.newRequest(op, input, output)
	return
}

// DiscoverPollEndpoint API operation for Amazon EC2 Container Service.
//
// This action is only used by the Amazon EC2 Container Service agent, and it
// is not intended for use outside of the agent.
//
// Returns an endpoint for the Amazon EC2 Container Service agent to poll for
// updates.
//
// Returns awserr.Error for service API and SDK errors. Use runtime type assertions
// with awserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the AWS API reference guide for Amazon EC2 Container Service's
// API operation DiscoverPollEndpoint for usage and error information.
//
// Returned Error Codes:
//   * ErrCodeServerException "ServerException"
//   These errors are usually caused by a server-side issue.
//
//   * ErrCodeClientException "ClientException"
//   These errors are usually caused by something the client did, such as use
//   an action or resource on behalf of a user that doesn't have permission to
//   use the action or resource, or specify an identifier that is not valid.
//
func (c *ECS) DiscoverPollEndpoint(input *DiscoverPollEndpointInput) (*DiscoverPollEndpointOutput, error) {
	req, out := c.DiscoverPollEndpointRequest(input)
	return out, req.Send()
}

// DiscoverPollEndpointWithContext is the same as DiscoverPollEndpoint with the addition of
// the ability to pass a context and additional request options.
//
// See DiscoverPollEndpoint for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ECS) DiscoverPollEndpointWithContext(ctx aws.Context, input *DiscoverPollEndpointInput, opts ...request.Option) (*DiscoverPollEndpointOutput, error) {
	req, out := c.DiscoverPollEndpointRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opListAttributes = "ListAttributes"

// ListAttributesRequest generates a "aws/request.Request" representing the
// client's request for the ListAttributes operation. The "output" return
// value can be used to capture response data after the request's "Send" method
// is called.
//
// See ListAttributes for usage and error information.
//
// Creating a request object using this method should be used when you want to inject
// custom logic into the request's lifecycle using a custom handler, or if you want to
// access properties on the request object before or after sending the request. If
// you just want the service response, call the ListAttributes method directly
// instead.
//
// Note: You must call the "Send" method on the returned request object in order
// to execute the request.
//
//    // Example sending a request using the ListAttributesRequest method.
//    req, resp := client.ListAttributesRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ECS) ListAttributesRequest(input *ListAttributesInput) (req *request.Request, output *ListAttributesOutput) {
	op := &request.Operation{
		Name:       opListAttributes,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListAttributesInput{}
	}

	output = &ListAttributesOutput{}
	req = c.newRequest(op, input, output)
	return
}

// ListAttributes API operation for Amazon EC2 Container Service.
//
// Returns awserr.Error for service API and SDK errors. Use runtime type assertions
// with awserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the AWS API reference guide for Amazon EC2 Container Service's
// API operation ListAttributes for usage and error information.
//
// Returned Error Codes:
//   * ErrCodeClusterNotFoundException "ClusterNotFoundException"
//   The specified cluster could not be found. You can view your available clusters
//   with ListClusters. Amazon ECS clusters are region-specific.
//
//   * ErrCodeInvalidParameterException "InvalidParameterException"
//   The specified parameter is invalid. Review the available parameters for the
//   API request.
//
func (c *ECS) ListAttributes(input *ListAttributesInput) (*ListAttributesOutput, error) {
	req, out := c.ListAttributesRequest(input)
	return out, req.Send()
}

// ListAttributesWithContext is the same as ListAttributes with the addition of
// the ability to pass a context and additional request options.
//
// See ListAttributes for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ECS) ListAttributesWithContext(ctx aws.Context, input *ListAttributesInput, opts ...request.Option) (*ListAttributesOutput, error) {
	req, out := c.ListAttributesRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opListClusters = "ListClusters"

// ListClustersRequest generates a "aws/request.Request" representing the
// client's request for the ListClusters operation. The "output" return
// value can be used to capture response data after the request's "Send" method
// is called.
//
// See ListClusters for usage and error information.
//
// Creating a request object using this method should be used when you want to inject
// custom logic into the request's lifecycle using a custom handler, or if you want to
// access properties on the request object before or after sending the request. If
// you just want the service response, call the ListClusters method directly
// instead.
//
// Note: You must call the "Send" method on the returned request object in order
// to execute the request.
//
//    // Example sending a request using the ListClustersRequest method.
//    req, resp := client.ListClustersRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ECS) ListClustersRequest(input *ListClustersInput) (req *request.Request, output *ListClustersOutput) {
	op := &request.Operation{
		Name:       opListClusters,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &request.Paginator{
			InputTokens:     []string{"nextToken"},
			OutputTokens:    []string{"nextToken"},
			LimitToken:      "maxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListClustersInput{}
	}

	output = &ListClustersOutput{}
	req = c.newRequest(op, input, output)
	return
}

// ListClusters API operation for Amazon EC2 Container Service.
//
// Returns a list of existing clusters.
//
// Returns awserr.Error for service API and SDK errors. Use runtime type assertions
// with awserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the AWS API reference guide for Amazon EC2 Container Service's
// API operation ListClusters for usage and error information.
//
// Returned Error Codes:
//   * ErrCodeServerException "ServerException"
//   These errors are usually caused by a server-side issue.
//
//   * ErrCodeClientException "ClientException"
//   These errors are usually caused by something the client did, such as use
//   an action or resource on behalf of a user that doesn't have permission to
//   use the action or resource, or specify an identifier that is not valid.
//
//   * ErrCodeInvalidParameterException "InvalidParameterException"
//   The specified parameter is invalid. Review the available parameters for the
//   API request.
//
func (c *ECS) ListClusters(input *ListClustersInput) (*ListClustersOutput, error) {
	req, out := c.ListClustersRequest(input)
	return out, req.Send()
}

// ListClustersWithContext is the same as ListClusters with the addition of
// the ability to pass a context and additional request options.
//
// See ListClusters for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ECS) ListClustersWithContext(ctx aws.Context, input *ListClustersInput, opts ...request.Option) (*ListClustersOutput, error) {
	req, out := c.ListClustersRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

// ListClustersPages iterates over the pages of a ListClusters operation,
// calling the "fn" function with the response data for each page. To stop
// iterating, return false from the fn function.
//
// See ListClusters method for more information on how to use this operation.
//
// Note: This operation can generate multiple requests to a service.
//
//    // Example iterating over at most 3 pages of a ListClusters operation.
//    pageNum := 0
//    err := client.ListClustersPages(params,
//        func(page *ListClustersOutput, lastPage bool) bool {
//            pageNum++
//            fmt.Println(page)
//            return pageNum <= 3
//        })
//
func (c *ECS) ListClustersPages(input *ListClustersInput, fn func(*ListClustersOutput, bool) bool) error {
	return c.ListClustersPagesWithContext(aws.BackgroundContext(), input, fn)
}

// ListClustersPagesWithContext same as ListClustersPages except
// it takes a Context and allows setting request options on the pages.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ECS) ListClustersPagesWithContext(ctx aws.Context, input *ListClustersInput, fn func(*ListClustersOutput, bool) bool, opts ...request.Option) error {
	p := request.Pagination{
		NewRequest: func() (*request.Request, error) {
			var inCpy *ListClustersInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req, _ := c.ListClustersRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req, nil
		},
	}

	cont := true
	for p.Next() && cont {
		cont = fn(p.Page().(*ListClustersOutput), !p.HasNextPage())
	}
	return p.Err()
}

const opListContainerInstances = "ListContainerInstances"

// ListContainerInstancesRequest generates a "aws/request.Request" representing the
// client's request for the ListContainerInstances operation. The "output" return
// value can be used to capture response data after the request's "Send" method
// is called.
//
// See ListContainerInstances for usage and error information.
//
// Creating a request object using this method should be used when you want to inject
// custom logic into the request's lifecycle using a custom handler, or if you want to
// access properties on the request object before or after sending the request. If
// you just want the service response, call the ListContainerInstances method directly
// instead.
//
// Note: You must call the "Send" method on the returned request object in order
// to execute the request.
//
//    // Example sending a request using the ListContainerInstancesRequest method.
//    req, resp := client.ListContainerInstancesRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ECS) ListContainerInstancesRequest(input *ListContainerInstancesInput) (req *request.Request, output *ListContainerInstancesOutput) {
	op := &request.Operation{
		Name:       opListContainerInstances,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &request.Paginator{
			InputTokens:     []string{"nextToken"},
			OutputTokens:    []string{"nextToken"},
			LimitToken:      "maxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListContainerInstancesInput{}
	}

	output = &ListContainerInstancesOutput{}
	req = c.newRequest(op, input, output)
	return
}

// ListContainerInstances API operation for Amazon EC2 Container Service.
//
// Returns a list of container instances in a specified cluster.
//
// Returns awserr.Error for service API and SDK errors. Use runtime type assertions
// with awserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the AWS API reference guide for Amazon EC2 Container Service's
// API operation ListContainerInstances for usage and error information.
//
// Returned Error Codes:
//   * ErrCodeServerException "ServerException"
//   These errors are usually caused by a server-side issue.
//
//   * ErrCodeClientException "ClientException"
//   These errors are usually caused by something the client did, such as use
//   an action or resource on behalf of a user that doesn't have permission to
//   use the action or resource, or specify an identifier that is not valid.
//
//   * ErrCodeInvalidParameterException "InvalidParameterException"
//   The specified parameter is invalid. Review the available parameters for the
//   API request.
//
//   * ErrCodeClusterNotFoundException "ClusterNotFoundException"
//   The specified cluster could not be found. You can view your available clusters
//   with ListClusters. Amazon ECS clusters are region-specific.
//
func (c *ECS) ListContainerInstances(input *ListContainerInstancesInput) (*ListContainerInstancesOutput, error) {
	req, out := c.ListContainerInstancesRequest(input)
	return out, req.Send()
}

// ListContainerInstancesWithContext is the same as ListContainerInstances with the addition of
// the ability to pass a context and additional request options.
//
// See ListContainerInstances for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ECS) ListContainerInstancesWithContext(ctx aws.Context, input *ListContainerInstancesInput, opts ...request.Option) (*ListContainerInstancesOutput, error) {
	req, out := c.ListContainerInstancesRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

// ListContainerInstancesPages iterates over the pages of a ListContainerInstances operation,
// calling the "fn" function with the response data for each page. To stop
// iterating, return false from the fn function.
//
// See ListContainerInstances method for more information on how to use this operation.
//
// Note: This operation can generate multiple requests to a service.
//
//    // Example iterating over at most 3 pages of a ListContainerInstances operation.
//    pageNum := 0
//    err := client.ListContainerInstancesPages(params,
//        func(page *ListContainerInstancesOutput, lastPage bool) bool {
//            pageNum++
//            fmt.Println(page)
//            return pageNum <= 3
//        })
//
func (c *ECS) ListContainerInstancesPages(input *ListContainerInstancesInput, fn func(*ListContainerInstancesOutput, bool) bool) error {
	return c.ListContainerInstancesPagesWithContext(aws.BackgroundContext(), input, fn)
}

// ListContainerInstancesPagesWithContext same as ListContainerInstancesPages except
// it takes a Context and allows setting request options on the pages.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ECS) ListContainerInstancesPagesWithContext(ctx aws.Context, input *ListContainerInstancesInput, fn func(*ListContainerInstancesOutput, bool) bool, opts ...request.Option) error {
	p := request.Pagination{
		NewRequest: func() (*request.Request, error) {
			var inCpy *ListContainerInstancesInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req, _ := c.ListContainerInstancesRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req, nil
		},
	}

	cont := true
	for p.Next() && cont {
		cont = fn(p.Page().(*ListContainerInstancesOutput), !p.HasNextPage())
	}
	return p.Err()
}

const opListServices = "ListServices"

// ListServicesRequest generates a "aws/request.Request" representing the
// client's request for the ListServices operation. The "output" return
// value can be used to capture response data after the request's "Send" method
// is called.
//
// See ListServices for usage and error information.
//
// Creating a request object using this method should be used when you want to inject
// custom logic into the request's lifecycle using a custom handler, or if you want to
// access properties on the request object before or after sending the request. If
// you just want the service response, call the ListServices method directly
// instead.
//
// Note: You must call the "Send" method on the returned request object in order
// to execute the request.
//
//    // Example sending a request using the ListServicesRequest method.
//    req, resp := client.ListServicesRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ECS) ListServicesRequest(input *ListServicesInput) (req *request.Request, output *ListServicesOutput) {
	op := &request.Operation{
		Name:       opListServices,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &request.Paginator{
			InputTokens:     []string{"nextToken"},
			OutputTokens:    []string{"nextToken"},
			LimitToken:      "maxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListServicesInput{}
	}

	output = &ListServicesOutput{}
	req = c.newRequest(op, input, output)
	return
}

// ListServices API operation for Amazon EC2 Container Service.
//
// Lists the services that are running in a specified cluster.
//
// Returns awserr.Error for service API and SDK errors. Use runtime type assertions
// with awserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the AWS API reference guide for Amazon EC2 Container Service's
// API operation ListServices for usage and error information.
//
// Returned Error Codes:
//   * ErrCodeServerException "ServerException"
//   These errors are usually caused by a server-side issue.
//
//   * ErrCodeClientException "ClientException"
//   These errors are usually caused by something the client did, such as use
//   an action or resource on behalf of a user that doesn't have permission to
//   use the action or resource, or specify an identifier that is not valid.
//
//   * ErrCodeInvalidParameterException "InvalidParameterException"
//   The specified parameter is invalid. Review the available parameters for the
//   API request.
//
//   * ErrCodeClusterNotFoundException "ClusterNotFoundException"
//   The specified cluster could not be found. You can view your available clusters
//   with ListClusters. Amazon ECS clusters are region-specific.
//
func (c *ECS) ListServices(input *ListServicesInput) (*ListServicesOutput, error) {
	req, out := c.ListServicesRequest(input)
	return out, req.Send()
}

// ListServicesWithContext is the same as ListServices with the addition of
// the ability to pass a context and additional request options.
//
// See ListServices for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ECS) ListServicesWithContext(ctx aws.Context, input *ListServicesInput, opts ...request.Option) (*ListServicesOutput, error) {
	req, out := c.ListServicesRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

// ListServicesPages iterates over the pages of a ListServices operation,
// calling the "fn" function with the response data for each page. To stop
// iterating, return false from the fn function.
//
// See ListServices method for more information on how to use this operation.
//
// Note: This operation can generate multiple requests to a service.
//
//    // Example iterating over at most 3 pages of a ListServices operation.
//    pageNum := 0
//    err := client.ListServicesPages(params,
//        func(page *ListServicesOutput, lastPage bool) bool {
//            pageNum++
//            fmt.Println(page)
//            return pageNum <= 3
//        })
//
func (c *ECS) ListServicesPages(input *ListServicesInput, fn func(*ListServicesOutput, bool) bool) error {
	return c.ListServicesPagesWithContext(aws.BackgroundContext(), input, fn)
}

// ListServicesPagesWithContext same as ListServicesPages except
// it takes a Context and allows setting request options on the pages.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ECS) ListServicesPagesWithContext(ctx aws.Context, input *ListServicesInput, fn func(*ListServicesOutput, bool) bool, opts ...request.Option) error {
	p := request.Pagination{
		NewRequest: func() (*request.Request, error) {
			var inCpy *ListServicesInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req, _ := c.ListServicesRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req, nil
		},
	}

	cont := true
	for p.Next() && cont {
		cont = fn(p.Page().(*ListServicesOutput), !p.HasNextPage())
	}
	return p.Err()
}

const opListTaskDefinitionFamilies = "ListTaskDefinitionFamilies"

// ListTaskDefinitionFamiliesRequest generates a "aws/request.Request" representing the
// client's request for the ListTaskDefinitionFamilies operation. The "output" return
// value can be used to capture response data after the request's "Send" method
// is called.
//
// See ListTaskDefinitionFamilies for usage and error information.
//
// Creating a request object using this method should be used when you want to inject
// custom logic into the request's lifecycle using a custom handler, or if you want to
// access properties on the request object before or after sending the request. If
// you just want the service response, call the ListTaskDefinitionFamilies method directly
// instead.
//
// Note: You must call the "Send" method on the returned request object in order
// to execute the request.
//
//    // Example sending a request using the ListTaskDefinitionFamiliesRequest method.
//    req, resp := client.ListTaskDefinitionFamiliesRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ECS) ListTaskDefinitionFamiliesRequest(input *ListTaskDefinitionFamiliesInput) (req *request.Request, output *ListTaskDefinitionFamiliesOutput) {
	op := &request.Operation{
		Name:       opListTaskDefinitionFamilies,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &request.Paginator{
			InputTokens:     []string{"nextToken"},
			OutputTokens:    []string{"nextToken"},
			LimitToken:      "maxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListTaskDefinitionFamiliesInput{}
	}

	output = &ListTaskDefinitionFamiliesOutput{}
	req = c.newRequest(op, input, output)
	return
}

// ListTaskDefinitionFamilies API operation for Amazon EC2 Container Service.
//
// Returns a list of task definition families that are registered to your account
// (which may include task definition families that no longer have any ACTIVE
// task definitions). You can filter the results with the familyPrefix parameter.
//
// Returns awserr.Error for service API and SDK errors. Use runtime type assertions
// with awserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the AWS API reference guide for Amazon EC2 Container Service's
// API operation ListTaskDefinitionFamilies for usage and error information.
//
// Returned Error Codes:
//   * ErrCodeServerException "ServerException"
//   These errors are usually caused by a server-side issue.
//
//   * ErrCodeClientException "ClientException"
//   These errors are usually caused by something the client did, such as use
//   an action or resource on behalf of a user that doesn't have permission to
//   use the action or resource, or specify an identifier that is not valid.
//
//   * ErrCodeInvalidParameterException "InvalidParameterException"
//   The specified parameter is invalid. Review the available parameters for the
//   API request.
//
func (c *ECS) ListTaskDefinitionFamilies(input *ListTaskDefinitionFamiliesInput) (*ListTaskDefinitionFamiliesOutput, error) {
	req, out := c.ListTaskDefinitionFamiliesRequest(input)
	return out, req.Send()
}

// ListTaskDefinitionFamiliesWithContext is the same as ListTaskDefinitionFamilies with the addition of
// the ability to pass a context and additional request options.
//
// See ListTaskDefinitionFamilies for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ECS) ListTaskDefinitionFamiliesWithContext(ctx aws.Context, input *ListTaskDefinitionFamiliesInput, opts ...request.Option) (*ListTaskDefinitionFamiliesOutput, error) {
	req, out := c.ListTaskDefinitionFamiliesRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

// ListTaskDefinitionFamiliesPages iterates over the pages of a ListTaskDefinitionFamilies operation,
// calling the "fn" function with the response data for each page. To stop
// iterating, return false from the fn function.
//
// See ListTaskDefinitionFamilies method for more information on how to use this operation.
//
// Note: This operation can generate multiple requests to a service.
//
//    // Example iterating over at most 3 pages of a ListTaskDefinitionFamilies operation.
//    pageNum := 0
//    err := client.ListTaskDefinitionFamiliesPages(params,
//        func(page *ListTaskDefinitionFamiliesOutput, lastPage bool) bool {
//            pageNum++
//            fmt.Println(page)
//            return pageNum <= 3
//        })
//
func (c *ECS) ListTaskDefinitionFamiliesPages(input *ListTaskDefinitionFamiliesInput, fn func(*ListTaskDefinitionFamiliesOutput, bool) bool) error {
	return c.ListTaskDefinitionFamiliesPagesWithContext(aws.BackgroundContext(), input, fn)
}

// ListTaskDefinitionFamiliesPagesWithContext same as ListTaskDefinitionFamiliesPages except
// it takes a Context and allows setting request options on the pages.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ECS) ListTaskDefinitionFamiliesPagesWithContext(ctx aws.Context, input *ListTaskDefinitionFamiliesInput, fn func(*ListTaskDefinitionFamiliesOutput, bool) bool, opts ...request.Option) error {
	p := request.Pagination{
		NewRequest: func() (*request.Request, error) {
			var inCpy *ListTaskDefinitionFamiliesInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req, _ := c.ListTaskDefinitionFamiliesRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req, nil
		},
	}

	cont := true
	for p.Next() && cont {
		cont = fn(p.Page().(*ListTaskDefinitionFamiliesOutput), !p.HasNextPage())
	}
	return p.Err()
}

const opListTaskDefinitions = "ListTaskDefinitions"

// ListTaskDefinitionsRequest generates a "aws/request.Request" representing the
// client's request for the ListTaskDefinitions operation. The "output" return
// value can be used to capture response data after the request's "Send" method
// is called.
//
// See ListTaskDefinitions for usage and error information.
//
// Creating a request object using this method should be used when you want to inject
// custom logic into the request's lifecycle using a custom handler, or if you want to
// access properties on the request object before or after sending the request. If
// you just want the service response, call the ListTaskDefinitions method directly
// instead.
//
// Note: You must call the "Send" method on the returned request object in order
// to execute the request.
//
//    // Example sending a request using the ListTaskDefinitionsRequest method.
//    req, resp := client.ListTaskDefinitionsRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ECS) ListTaskDefinitionsRequest(input *ListTaskDefinitionsInput) (req *request.Request, output *ListTaskDefinitionsOutput) {
	op := &request.Operation{
		Name:       opListTaskDefinitions,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &request.Paginator{
			InputTokens:     []string{"nextToken"},
			OutputTokens:    []string{"nextToken"},
			LimitToken:      "maxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListTaskDefinitionsInput{}
	}

	output = &ListTaskDefinitionsOutput{}
	req = c.newRequest(op, input, output)
	return
}

// ListTaskDefinitions API operation for Amazon EC2 Container Service.
//
// Returns a list of task definitions that are registered to your account. You
// can filter the results by family name with the familyPrefix parameter or
// by status with the status parameter.
//
// Returns awserr.Error for service API and SDK errors. Use runtime type assertions
// with awserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the AWS API reference guide for Amazon EC2 Container Service's
// API operation ListTaskDefinitions for usage and error information.
//
// Returned Error Codes:
//   * ErrCodeServerException "ServerException"
//   These errors are usually caused by a server-side issue.
//
//   * ErrCodeClientException "ClientException"
//   These errors are usually caused by something the client did, such as use
//   an action or resource on behalf of a user that doesn't have permission to
//   use the action or resource, or specify an identifier that is not valid.
//
//   * ErrCodeInvalidParameterException "InvalidParameterException"
//   The specified parameter is invalid. Review the available parameters for the
//   API request.
//
func (c *ECS) ListTaskDefinitions(input *ListTaskDefinitionsInput) (*ListTaskDefinitionsOutput, error) {
	req, out := c.ListTaskDefinitionsRequest(input)
	return out, req.Send()
}

// ListTaskDefinitionsWithContext is the same as ListTaskDefinitions with the addition of
// the ability to pass a context and additional request options.
//
// See ListTaskDefinitions for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ECS) ListTaskDefinitionsWithContext(ctx aws.Context, input *ListTaskDefinitionsInput, opts ...request.Option) (*ListTaskDefinitionsOutput, error) {
	req, out := c.ListTaskDefinitionsRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

// ListTaskDefinitionsPages iterates over the pages of a ListTaskDefinitions operation,
// calling the "fn" function with the response data for each page. To stop
// iterating, return false from the fn function.
//
// See ListTaskDefinitions method for more information on how to use this operation.
//
// Note: This operation can generate multiple requests to a service.
//
//    // Example iterating over at most 3 pages of a ListTaskDefinitions operation.
//    pageNum := 0
//    err := client.ListTaskDefinitionsPages(params,
//        func(page *ListTaskDefinitionsOutput, lastPage bool) bool {
//            pageNum++
//            fmt.Println(page)
//            return pageNum <= 3
//        })
//
func (c *ECS) ListTaskDefinitionsPages(input *ListTaskDefinitionsInput, fn func(*ListTaskDefinitionsOutput, bool) bool) error {
	return c.ListTaskDefinitionsPagesWithContext(aws.BackgroundContext(), input, fn)
}

// ListTaskDefinitionsPagesWithContext same as ListTaskDefinitionsPages except
// it takes a Context and allows setting request options on the pages.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ECS) ListTaskDefinitionsPagesWithContext(ctx aws.Context, input *ListTaskDefinitionsInput, fn func(*ListTaskDefinitionsOutput, bool) bool, opts ...request.Option) error {
	p := request.Pagination{
		NewRequest: func() (*request.Request, error) {
			var inCpy *ListTaskDefinitionsInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req, _ := c.ListTaskDefinitionsRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req, nil
		},
	}

	cont := true
	for p.Next() && cont {
		cont = fn(p.Page().(*ListTaskDefinitionsOutput), !p.HasNextPage())
	}
	return p.Err()
}

const opListTasks = "ListTasks"

// ListTasksRequest generates a "aws/request.Request" representing the
// client's request for the ListTasks operation. The "output" return
// value can be used to capture response data after the request's "Send" method
// is called.
//
// See ListTasks for usage and error information.
//
// Creating a request object using this method should be used when you want to inject
// custom logic into the request's lifecycle using a custom handler, or if you want to
// access properties on the request object before or after sending the request. If
// you just want the service response, call the ListTasks method directly
// instead.
//
// Note: You must call the "Send" method on the returned request object in order
// to execute the request.
//
//    // Example sending a request using the ListTasksRequest method.
//    req, resp := client.ListTasksRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ECS) ListTasksRequest(input *ListTasksInput) (req *request.Request, output *ListTasksOutput) {
	op := &request.Operation{
		Name:       opListTasks,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &request.Paginator{
			InputTokens:     []string{"nextToken"},
			OutputTokens:    []string{"nextToken"},
			LimitToken:      "maxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListTasksInput{}
	}

	output = &ListTasksOutput{}
	req = c.newRequest(op, input, output)
	return
}

// ListTasks API operation for Amazon EC2 Container Service.
//
// Returns a list of tasks for a specified cluster. You can filter the results
// by family name, by a particular container instance, or by the desired status
// of the task with the family, containerInstance, and desiredStatus parameters.
//
// Returns awserr.Error for service API and SDK errors. Use runtime type assertions
// with awserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the AWS API reference guide for Amazon EC2 Container Service's
// API operation ListTasks for usage and error information.
//
// Returned Error Codes:
//   * ErrCodeServerException "ServerException"
//   These errors are usually caused by a server-side issue.
//
//   * ErrCodeClientException "ClientException"
//   These errors are usually caused by something the client did, such as use
//   an action or resource on behalf of a user that doesn't have permission to
//   use the action or resource, or specify an identifier that is not valid.
//
//   * ErrCodeInvalidParameterException "InvalidParameterException"
//   The specified parameter is invalid. Review the available parameters for the
//   API request.
//
//   * ErrCodeClusterNotFoundException "ClusterNotFoundException"
//   The specified cluster could not be found. You can view your available clusters
//   with ListClusters. Amazon ECS clusters are region-specific.
//
//   * ErrCodeServiceNotFoundException "ServiceNotFoundException"
//   The specified service could not be found. You can view your available services
//   with ListServices. Amazon ECS services are cluster-specific and region-specific.
//
func (c *ECS) ListTasks(input *ListTasksInput) (*ListTasksOutput, error) {
	req, out := c.ListTasksRequest(input)
	return out, req.Send()
}

// ListTasksWithContext is the same as ListTasks with the addition of
// the ability to pass a context and additional request options.
//
// See ListTasks for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ECS) ListTasksWithContext(ctx aws.Context, input *ListTasksInput, opts ...request.Option) (*ListTasksOutput, error) {
	req, out := c.ListTasksRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

// ListTasksPages iterates over the pages of a ListTasks operation,
// calling the "fn" function with the response data for each page. To stop
// iterating, return false from the fn function.
//
// See ListTasks method for more information on how to use this operation.
//
// Note: This operation can generate multiple requests to a service.
//
//    // Example iterating over at most 3 pages of a ListTasks operation.
//    pageNum := 0
//    err := client.ListTasksPages(params,
//        func(page *ListTasksOutput, lastPage bool) bool {
//            pageNum++
//            fmt.Println(page)
//            return pageNum <= 3
//        })
//
func (c *ECS) ListTasksPages(input *ListTasksInput, fn func(*ListTasksOutput, bool) bool) error {
	return c.ListTasksPagesWithContext(aws.BackgroundContext(), input, fn)
}

// ListTasksPagesWithContext same as ListTasksPages except
// it takes a Context and allows setting request options on the pages.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ECS) ListTasksPagesWithContext(ctx aws.Context, input *ListTasksInput, fn func(*ListTasksOutput, bool) bool, opts ...request.Option) error {
	p := request.Pagination{
		NewRequest: func() (*request.Request, error) {
			var inCpy *ListTasksInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req, _ := c.ListTasksRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req, nil
		},
	}

	cont := true
	for p.Next() && cont {
		cont = fn(p.Page().(*ListTasksOutput), !p.HasNextPage())
	}
	return p.Err()
}

const opPutAttributes = "PutAttributes"

// PutAttributesRequest generates a "aws/request.Request" representing the
// client's request for the PutAttributes operation. The "output" return
// value can be used to capture response data after the request's "Send" method
// is called.
//
// See PutAttributes for usage and error information.
//
// Creating a request object using this method should be used when you want to inject
// custom logic into the request's lifecycle using a custom handler, or if you want to
// access properties on the request object before or after sending the request. If
// you just want the service response, call the PutAttributes method directly
// instead.
//
// Note: You must call the "Send" method on the returned request object in order
// to execute the request.
//
//    // Example sending a request using the PutAttributesRequest method.
//    req, resp := client.PutAttributesRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ECS) PutAttributesRequest(input *PutAttributesInput) (req *request.Request, output *PutAttributesOutput) {
	op := &request.Operation{
		Name:       opPutAttributes,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &PutAttributesInput{}
	}

	output = &PutAttributesOutput{}
	req = c.newRequest(op, input, output)
	return
}

// PutAttributes API operation for Amazon EC2 Container Service.
//
// Returns awserr.Error for service API and SDK errors. Use runtime type assertions
// with awserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the AWS API reference guide for Amazon EC2 Container Service's
// API operation PutAttributes for usage and error information.
//
// Returned Error Codes:
//   * ErrCodeClusterNotFoundException "ClusterNotFoundException"
//   The specified cluster could not be found. You can view your available clusters
//   with ListClusters. Amazon ECS clusters are region-specific.
//
//   * ErrCodeTargetNotFoundException "TargetNotFoundException"
//
//   * ErrCodeAttributeLimitExceededException "AttributeLimitExceededException"
//
//   * ErrCodeInvalidParameterException "InvalidParameterException"
//   The specified parameter is invalid. Review the available parameters for the
//   API request.
//
func (c *ECS) PutAttributes(input *PutAttributesInput) (*PutAttributesOutput, error) {
	req, out := c.PutAttributesRequest(input)
	return out, req.Send()
}

// PutAttributesWithContext is the same as PutAttributes with the addition of
// the ability to pass a context and additional request options.
//
// See PutAttributes for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ECS) PutAttributesWithContext(ctx aws.Context, input *PutAttributesInput, opts ...request.Option) (*PutAttributesOutput, error) {
	req, out := c.PutAttributesRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opRegisterContainerInstance = "RegisterContainerInstance"

// RegisterContainerInstanceRequest generates a "aws/request.Request" representing the
// client's request for the RegisterContainerInstance operation. The "output" return
// value can be used to capture response data after the request's "Send" method
// is called.
//
// See RegisterContainerInstance for usage and error information.
//
// Creating a request object using this method should be used when you want to inject
// custom logic into the request's lifecycle using a custom handler, or if you want to
// access properties on the request object before or after sending the request. If
// you just want the service response, call the RegisterContainerInstance method directly
// instead.
//
// Note: You must call the "Send" method on the returned request object in order
// to execute the request.
//
//    // Example sending a request using the RegisterContainerInstanceRequest method.
//    req, resp := client.RegisterContainerInstanceRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ECS) RegisterContainerInstanceRequest(input *RegisterContainerInstanceInput) (req *request.Request, output *RegisterContainerInstanceOutput) {
	op := &request.Operation{
		Name:       opRegisterContainerInstance,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &RegisterContainerInstanceInput{}
	}

	output = &RegisterContainerInstanceOutput{}
	req = c.newRequest(op, input, output)
	return
}

// RegisterContainerInstance API operation for Amazon EC2 Container Service.
//
// This action is only used by the Amazon EC2 Container Service agent, and it
// is not intended for use outside of the agent.
//
// Registers an Amazon EC2 instance into the specified cluster. This instance
// will become available to place containers on.
//
// Returns awserr.Error for service API and SDK errors. Use runtime type assertions
// with awserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the AWS API reference guide for Amazon EC2 Container Service's
// API operation RegisterContainerInstance for usage and error information.
//
// Returned Error Codes:
//   * ErrCodeServerException "ServerException"
//   These errors are usually caused by a server-side issue.
//
//   * ErrCodeClientException "ClientException"
//   These errors are usually caused by something the client did, such as use
//   an action or resource on behalf of a user that doesn't have permission to
//   use the action or resource, or specify an identifier that is not valid.
//
func (c *ECS) RegisterContainerInstance(input *RegisterContainerInstanceInput) (*RegisterContainerInstanceOutput, error) {
	req, out := c.RegisterContainerInstanceRequest(input)
	return out, req.Send()
}

// RegisterContainerInstanceWithContext is the same as RegisterContainerInstance with the addition of
// the ability to pass a context and additional request options.
//
// See RegisterContainerInstance for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ECS) RegisterContainerInstanceWithContext(ctx aws.Context, input *RegisterContainerInstanceInput, opts ...request.Option) (*RegisterContainerInstanceOutput, error) {
	req, out := c.RegisterContainerInstanceRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opRegisterTaskDefinition = "RegisterTaskDefinition"

// RegisterTaskDefinitionRequest generates a "aws/request.Request" representing the
// client's request for the RegisterTaskDefinition operation. The "output" return
// value can be used to capture response data after the request's "Send" method
// is called.
//
// See RegisterTaskDefinition for usage and error information.
//
// Creating a request object using this method should be used when you want to inject
// custom logic into the request's lifecycle using a custom handler, or if you want to
// access properties on the request object before or after sending the request. If
// you just want the service response, call the RegisterTaskDefinition method directly
// instead.
//
// Note: You must call the "Send" method on the returned request object in order
// to execute the request.
//
//    // Example sending a request using the RegisterTaskDefinitionRequest method.
//    req, resp := client.RegisterTaskDefinitionRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ECS) RegisterTaskDefinitionRequest(input *RegisterTaskDefinitionInput) (req *request.Request, output *RegisterTaskDefinitionOutput) {
	op := &request.Operation{
		Name:       opRegisterTaskDefinition,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &RegisterTaskDefinitionInput{}
	}

	output = &RegisterTaskDefinitionOutput{}
	req = c.newRequest(op, input, output)
	return
}

// RegisterTaskDefinition API operation for Amazon EC2 Container Service.
//
// Registers a new task definition from the supplied family and containerDefinitions.
// Optionally, you can add data volumes to your containers with the volumes
// parameter. For more information on task definition parameters and defaults,
// see Amazon ECS Task Definitions (http://docs.aws.amazon.com/AmazonECS/latest/developerguide/task_defintions.html)
// in the Amazon EC2 Container Service Developer Guide.
//
// Returns awserr.Error for service API and SDK errors. Use runtime type assertions
// with awserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the AWS API reference guide for Amazon EC2 Container Service's
// API operation RegisterTaskDefinition for usage and error information.
//
// Returned Error Codes:
//   * ErrCodeServerException "ServerException"
//   These errors are usually caused by a server-side issue.
//
//   * ErrCodeClientException "ClientException"
//   These errors are usually caused by something the client did, such as use
//   an action or resource on behalf of a user that doesn't have permission to
//   use the action or resource, or specify an identifier that is not valid.
//
//   * ErrCodeInvalidParameterException "InvalidParameterException"
//   The specified parameter is invalid. Review the available parameters for the
//   API request.
//
func (c *ECS) RegisterTaskDefinition(input *RegisterTaskDefinitionInput) (*RegisterTaskDefinitionOutput, error) {
	req, out := c.RegisterTaskDefinitionRequest(input)
	return out, req.Send()
}

// RegisterTaskDefinitionWithContext is the same as RegisterTaskDefinition with the addition of
// the ability to pass a context and additional request options.
//
// See RegisterTaskDefinition for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ECS) RegisterTaskDefinitionWithContext(ctx aws.Context, input *RegisterTaskDefinitionInput, opts ...request.Option) (*RegisterTaskDefinitionOutput, error) {
	req, out := c.RegisterTaskDefinitionRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opRunTask = "RunTask"

// RunTaskRequest generates a "aws/request.Request" representing the
// client's request for the RunTask operation. The "output" return
// value can be used to capture response data after the request's "Send" method
// is called.
//
// See RunTask for usage and error information.
//
// Creating a request object using this method should be used when you want to inject
// custom logic into the request's lifecycle using a custom handler, or if you want to
// access properties on the request object before or after sending the request. If
// you just want the service response, call the RunTask method directly
// instead.
//
// Note: You must call the "Send" method on the returned request object in order
// to execute the request.
//
//    // Example sending a request using the RunTaskRequest method.
//    req, resp := client.RunTaskRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ECS) RunTaskRequest(input *RunTaskInput) (req *request.Request, output *RunTaskOutput) {
	op := &request.Operation{
		Name:       opRunTask,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &RunTaskInput{}
	}

	output = &RunTaskOutput{}
	req = c.newRequest(op, input, output)
	return
}

// RunTask API operation for Amazon EC2 Container Service.
//
// Start a task using random placement and the default Amazon ECS scheduler.
// If you want to use your own scheduler or place a task on a specific container
// instance, use StartTask instead.
//
// The count parameter is limited to 10 tasks per call.
//
// Returns awserr.Error for service API and SDK errors. Use runtime type assertions
// with awserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the AWS API reference guide for Amazon EC2 Container Service's
// API operation RunTask for usage and error information.
//
// Returned Error Codes:
//   * ErrCodeServerException "ServerException"
//   These errors are usually caused by a server-side issue.
//
//   * ErrCodeClientException "ClientException"
//   These errors are usually caused by something the client did, such as use
//   an action or resource on behalf of a user that doesn't have permission to
//   use the action or resource, or specify an identifier that is not valid.
//
//   * ErrCodeInvalidParameterException "InvalidParameterException"
//   The specified parameter is invalid. Review the available parameters for the
//   API request.
//
//   * ErrCodeClusterNotFoundException "ClusterNotFoundException"
//   The specified cluster could not be found. You can view your available clusters
//   with ListClusters. Amazon ECS clusters are region-specific.
//
func (c *ECS) RunTask(input *RunTaskInput) (*RunTaskOutput, error) {
	req, out := c.RunTaskRequest(input)
	return out, req.Send()
}

// RunTaskWithContext is the same as RunTask with the addition of
// the ability to pass a context and additional request options.
//
// See RunTask for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ECS) RunTaskWithContext(ctx aws.Context, input *RunTaskInput, opts ...request.Option) (*RunTaskOutput, error) {
	req, out := c.RunTaskRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opStartTask = "StartTask"

// StartTaskRequest generates a "aws/request.Request" representing the
// client's request for the StartTask operation. The "output" return
// value can be used to capture response data after the request's "Send" method
// is called.
//
// See StartTask for usage and error information.
//
// Creating a request object using this method should be used when you want to inject
// custom logic into the request's lifecycle using a custom handler, or if you want to
// access properties on the request object before or after sending the request. If
// you just want the service response, call the StartTask method directly
// instead.
//
// Note: You must call the "Send" method on the returned request object in order
// to execute the request.
//
//    // Example sending a request using the StartTaskRequest method.
//    req, resp := client.StartTaskRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ECS) StartTaskRequest(input *StartTaskInput) (req *request.Request, output *StartTaskOutput) {
	op := &request.Operation{
		Name:       opStartTask,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &StartTaskInput{}
	}

	output = &StartTaskOutput{}
	req = c.newRequest(op, input, output)
	return
}

// StartTask API operation for Amazon EC2 Container Service.
//
// Starts a new task from the specified task definition on the specified container
// instance or instances. If you want to use the default Amazon ECS scheduler
// to place your task, use RunTask instead.
//
// The list of container instances to start tasks on is limited to 10.
//
// Returns awserr.Error for service API and SDK errors. Use runtime type assertions
// with awserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the AWS API reference guide for Amazon EC2 Container Service's
// API operation StartTask for usage and error information.
//
// Returned Error Codes:
//   * ErrCodeServerException "ServerException"
//   These errors are usually caused by a server-side issue.
//
//   * ErrCodeClientException "ClientException"
//   These errors are usually caused by something the client did, such as use
//   an action or resource on behalf of a user that doesn't have permission to
//   use the action or resource, or specify an identifier that is not valid.
//
//   * ErrCodeInvalidParameterException "InvalidParameterException"
//   The specified parameter is invalid. Review the available parameters for the
//   API request.
//
//   * ErrCodeClusterNotFoundException "ClusterNotFoundException"
//   The specified cluster could not be found. You can view your available clusters
//   with ListClusters. Amazon ECS clusters are region-specific.
//
func (c *ECS) StartTask(input *StartTaskInput) (*StartTaskOutput, error) {
	req, out := c.StartTaskRequest(input)
	return out, req.Send()
}

// StartTaskWithContext is the same as StartTask with the addition of
// the ability to pass a context and additional request options.
//
// See StartTask for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ECS) StartTaskWithContext(ctx aws.Context, input *StartTaskInput, opts ...request.Option) (*StartTaskOutput, error) {
	req, out := c.StartTaskRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opStopTask = "StopTask"

// StopTaskRequest generates a "aws/request.Request" representing the
// client's request for the StopTask operation. The "output" return
// value can be used to capture response data after the request's "Send" method
// is called.
//
// See StopTask for usage and error information.
//
// Creating a request object using this method should be used when you want to inject
// custom logic into the request's lifecycle using a custom handler, or if you want to
// access properties on the request object before or after sending the request. If
// you just want the service response, call the StopTask method directly
// instead.
//
// Note: You must call the "Send" method on the returned request object in order
// to execute the request.
//
//    // Example sending a request using the StopTaskRequest method.
//    req, resp := client.StopTaskRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ECS) StopTaskRequest(input *StopTaskInput) (req *request.Request, output *StopTaskOutput) {
	op := &request.Operation{
		Name:       opStopTask,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &StopTaskInput{}
	}

	output = &StopTaskOutput{}
	req = c.newRequest(op, input, output)
	return
}

// StopTask API operation for Amazon EC2 Container Service.
//
// Stops a running task.
//
// When StopTask is called on a task, the equivalent of docker stop is issued
// to the containers running in the task. This results in a SIGTERM and a 30-second
// timeout, after which SIGKILL is sent and the containers are forcibly stopped.
// If the container handles the SIGTERM gracefully and exits within 30 seconds
// from receiving it, no SIGKILL is sent.
//
// Returns awserr.Error for service API and SDK errors. Use runtime type assertions
// with awserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the AWS API reference guide for Amazon EC2 Container Service's
// API operation StopTask for usage and error information.
//
// Returned Error Codes:
//   * ErrCodeServerException "ServerException"
//   These errors are usually caused by a server-side issue.
//
//   * ErrCodeClientException "ClientException"
//   These errors are usually caused by something the client did, such as use
//   an action or resource on behalf of a user that doesn't have permission to
//   use the action or resource, or specify an identifier that is not valid.
//
//   * ErrCodeInvalidParameterException "InvalidParameterException"
//   The specified parameter is invalid. Review the available parameters for the
//   API request.
//
//   * ErrCodeClusterNotFoundException "ClusterNotFoundException"
//   The specified cluster could not be found. You can view your available clusters
//   with ListClusters. Amazon ECS clusters are region-specific.
//
func (c *ECS) StopTask(input *StopTaskInput) (*StopTaskOutput, error) {
	req, out := c.StopTaskRequest(input)
	return out, req.Send()
}

// StopTaskWithContext is the same as StopTask with the addition of
// the ability to pass a context and additional request options.
//
// See StopTask for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ECS) StopTaskWithContext(ctx aws.Context, input *StopTaskInput, opts ...request.Option) (*StopTaskOutput, error) {
	req, out := c.StopTaskRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opSubmitContainerStateChange = "SubmitContainerStateChange"

// SubmitContainerStateChangeRequest generates a "aws/request.Request" representing the
// client's request for the SubmitContainerStateChange operation. The "output" return
// value can be used to capture response data after the request's "Send" method
// is called.
//
// See SubmitContainerStateChange for usage and error information.
//
// Creating a request object using this method should be used when you want to inject
// custom logic into the request's lifecycle using a custom handler, or if you want to
// access properties on the request object before or after sending the request. If
// you just want the service response, call the SubmitContainerStateChange method directly
// instead.
//
// Note: You must call the "Send" method on the returned request object in order
// to execute the request.
//
//    // Example sending a request using the SubmitContainerStateChangeRequest method.
//    req, resp := client.SubmitContainerStateChangeRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ECS) SubmitContainerStateChangeRequest(input *SubmitContainerStateChangeInput) (req *request.Request, output *SubmitContainerStateChangeOutput) {
	op := &request.Operation{
		Name:       opSubmitContainerStateChange,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &SubmitContainerStateChangeInput{}
	}

	output = &SubmitContainerStateChangeOutput{}
	req = c.newRequest(op, input, output)
	return
}

// SubmitContainerStateChange API operation for Amazon EC2 Container Service.
//
// This action is only used by the Amazon EC2 Container Service agent, and it
// is not intended for use outside of the agent.
//
// Sent to acknowledge that a container changed states.
//
// Returns awserr.Error for service API and SDK errors. Use runtime type assertions
// with awserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the AWS API reference guide for Amazon EC2 Container Service's
// API operation SubmitContainerStateChange for usage and error information.
//
// Returned Error Codes:
//   * ErrCodeServerException "ServerException"
//   These errors are usually caused by a server-side issue.
//
//   * ErrCodeClientException "ClientException"
//   These errors are usually caused by something the client did, such as use
//   an action or resource on behalf of a user that doesn't have permission to
//   use the action or resource, or specify an identifier that is not valid.
//
func (c *ECS) SubmitContainerStateChange(input *SubmitContainerStateChangeInput) (*SubmitContainerStateChangeOutput, error) {
	req, out := c.SubmitContainerStateChangeRequest(input)
	return out, req.Send()
}

// SubmitContainerStateChangeWithContext is the same as SubmitContainerStateChange with the addition of
// the ability to pass a context and additional request options.
//
// See SubmitContainerStateChange for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ECS) SubmitContainerStateChangeWithContext(ctx aws.Context, input *SubmitContainerStateChangeInput, opts ...request.Option) (*SubmitContainerStateChangeOutput, error) {
	req, out := c.SubmitContainerStateChangeRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opSubmitTaskStateChange = "SubmitTaskStateChange"

// SubmitTaskStateChangeRequest generates a "aws/request.Request" representing the
// client's request for the SubmitTaskStateChange operation. The "output" return
// value can be used to capture response data after the request's "Send" method
// is called.
//
// See SubmitTaskStateChange for usage and error information.
//
// Creating a request object using this method should be used when you want to inject
// custom logic into the request's lifecycle using a custom handler, or if you want to
// access properties on the request object before or after sending the request. If
// you just want the service response, call the SubmitTaskStateChange method directly
// instead.
//
// Note: You must call the "Send" method on the returned request object in order
// to execute the request.
//
//    // Example sending a request using the SubmitTaskStateChangeRequest method.
//    req, resp := client.SubmitTaskStateChangeRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ECS) SubmitTaskStateChangeRequest(input *SubmitTaskStateChangeInput) (req *request.Request, output *SubmitTaskStateChangeOutput) {
	op := &request.Operation{
		Name:       opSubmitTaskStateChange,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &SubmitTaskStateChangeInput{}
	}

	output = &SubmitTaskStateChangeOutput{}
	req = c.newRequest(op, input, output)
	return
}

// SubmitTaskStateChange API operation for Amazon EC2 Container Service.
//
// This action is only used by the Amazon EC2 Container Service agent, and it
// is not intended for use outside of the agent.
//
// Sent to acknowledge that a task changed states.
//
// Returns awserr.Error for service API and SDK errors. Use runtime type assertions
// with awserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the AWS API reference guide for Amazon EC2 Container Service's
// API operation SubmitTaskStateChange for usage and error information.
//
// Returned Error Codes:
//   * ErrCodeServerException "ServerException"
//   These errors are usually caused by a server-side issue.
//
//   * ErrCodeClientException "ClientException"
//   These errors are usually caused by something the client did, such as use
//   an action or resource on behalf of a user that doesn't have permission to
//   use the action or resource, or specify an identifier that is not valid.
//
func (c *ECS) SubmitTaskStateChange(input *SubmitTaskStateChangeInput) (*SubmitTaskStateChangeOutput, error) {
	req, out := c.SubmitTaskStateChangeRequest(input)
	return out, req.Send()
}

// SubmitTaskStateChangeWithContext is the same as SubmitTaskStateChange with the addition of
// the ability to pass a context and additional request options.
//
// See SubmitTaskStateChange for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ECS) SubmitTaskStateChangeWithContext(ctx aws.Context, input *SubmitTaskStateChangeInput, opts ...request.Option) (*SubmitTaskStateChangeOutput, error) {
	req, out := c.SubmitTaskStateChangeRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opUpdateContainerAgent = "UpdateContainerAgent"

// UpdateContainerAgentRequest generates a "aws/request.Request" representing the
// client's request for the UpdateContainerAgent operation. The "output" return
// value can be used to capture response data after the request's "Send" method
// is called.
//
// See UpdateContainerAgent for usage and error information.
//
// Creating a request object using this method should be used when you want to inject
// custom logic into the request's lifecycle using a custom handler, or if you want to
// access properties on the request object before or after sending the request. If
// you just want the service response, call the UpdateContainerAgent method directly
// instead.
//
// Note: You must call the "Send" method on the returned request object in order
// to execute the request.
//
//    // Example sending a request using the UpdateContainerAgentRequest method.
//    req, resp := client.UpdateContainerAgentRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ECS) UpdateContainerAgentRequest(input *UpdateContainerAgentInput) (req *request.Request, output *UpdateContainerAgentOutput) {
	op := &request.Operation{
		Name:       opUpdateContainerAgent,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateContainerAgentInput{}
	}

	output = &UpdateContainerAgentOutput{}
	req = c.newRequest(op, input, output)
	return
}

// UpdateContainerAgent API operation for Amazon EC2 Container Service.
//
// Updates the Amazon ECS container agent on a specified container instance.
// Updating the Amazon ECS container agent does not interrupt running tasks
// or services on the container instance. The process for updating the agent
// differs depending on whether your container instance was launched with the
// Amazon ECS-optimized AMI or another operating system.
//
// UpdateContainerAgent requires the Amazon ECS-optimized AMI or Amazon Linux
// with the ecs-init service installed and running. For help updating the Amazon
// ECS container agent on other operating systems, see Manually Updating the
// Amazon ECS Container Agent (http://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-agent-update.html#manually_update_agent)
// in the Amazon EC2 Container Service Developer Guide.
//
// Returns awserr.Error for service API and SDK errors. Use runtime type assertions
// with awserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the AWS API reference guide for Amazon EC2 Container Service's
// API operation UpdateContainerAgent for usage and error information.
//
// Returned Error Codes:
//   * ErrCodeServerException "ServerException"
//   These errors are usually caused by a server-side issue.
//
//   * ErrCodeClientException "ClientException"
//   These errors are usually caused by something the client did, such as use
//   an action or resource on behalf of a user that doesn't have permission to
//   use the action or resource, or specify an identifier that is not valid.
//
//   * ErrCodeInvalidParameterException "InvalidParameterException"
//   The specified parameter is invalid. Review the available parameters for the
//   API request.
//
//   * ErrCodeClusterNotFoundException "ClusterNotFoundException"
//   The specified cluster could not be found. You can view your available clusters
//   with ListClusters. Amazon ECS clusters are region-specific.
//
//   * ErrCodeUpdateInProgressException "UpdateInProgressException"
//   There is already a current Amazon ECS container agent update in progress
//   on the specified container instance. If the container agent becomes disconnected
//   while it is in a transitional stage, such as PENDING or STAGING, the update
//   process can get stuck in that state. However, when the agent reconnects,
//   it will resume where it stopped previously.
//
//   * ErrCodeNoUpdateAvailableException "NoUpdateAvailableException"
//   There is no update available for this Amazon ECS container agent. This could
//   be because the agent is already running the latest version, or it is so old
//   that there is no update path to the current version.
//
//   * ErrCodeMissingVersionException "MissingVersionException"
//   Amazon ECS is unable to determine the current version of the Amazon ECS container
//   agent on the container instance and does not have enough information to proceed
//   with an update. This could be because the agent running on the container
//   instance is an older or custom version that does not use our version information.
//
func (c *ECS) UpdateContainerAgent(input *UpdateContainerAgentInput) (*UpdateContainerAgentOutput, error) {
	req, out := c.UpdateContainerAgentRequest(input)
	return out, req.Send()
}

// UpdateContainerAgentWithContext is the same as UpdateContainerAgent with the addition of
// the ability to pass a context and additional request options.
//
// See UpdateContainerAgent for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ECS) UpdateContainerAgentWithContext(ctx aws.Context, input *UpdateContainerAgentInput, opts ...request.Option) (*UpdateContainerAgentOutput, error) {
	req, out := c.UpdateContainerAgentRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opUpdateContainerInstancesState = "UpdateContainerInstancesState"

// UpdateContainerInstancesStateRequest generates a "aws/request.Request" representing the
// client's request for the UpdateContainerInstancesState operation. The "output" return
// value can be used to capture response data after the request's "Send" method
// is called.
//
// See UpdateContainerInstancesState for usage and error information.
//
// Creating a request object using this method should be used when you want to inject
// custom logic into the request's lifecycle using a custom handler, or if you want to
// access properties on the request object before or after sending the request. If
// you just want the service response, call the UpdateContainerInstancesState method directly
// instead.
//
// Note: You must call the "Send" method on the returned request object in order
// to execute the request.
//
//    // Example sending a request using the UpdateContainerInstancesStateRequest method.
//    req, resp := client.UpdateContainerInstancesStateRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ECS) UpdateContainerInstancesStateRequest(input *UpdateContainerInstancesStateInput) (req *request.Request, output *UpdateContainerInstancesStateOutput) {
	op := &request.Operation{
		Name:       opUpdateContainerInstancesState,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateContainerInstancesStateInput{}
	}

	output = &UpdateContainerInstancesStateOutput{}
	req = c.newRequest(op, input, output)
	return
}

// UpdateContainerInstancesState API operation for Amazon EC2 Container Service.
//
// Returns awserr.Error for service API and SDK errors. Use runtime type assertions
// with awserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the AWS API reference guide for Amazon EC2 Container Service's
// API operation UpdateContainerInstancesState for usage and error information.
//
// Returned Error Codes:
//   * ErrCodeServerException "ServerException"
//   These errors are usually caused by a server-side issue.
//
//   * ErrCodeClientException "ClientException"
//   These errors are usually caused by something the client did, such as use
//   an action or resource on behalf of a user that doesn't have permission to
//   use the action or resource, or specify an identifier that is not valid.
//
//   * ErrCodeInvalidParameterException "InvalidParameterException"
//   The specified parameter is invalid. Review the available parameters for the
//   API request.
//
//   * ErrCodeClusterNotFoundException "ClusterNotFoundException"
//   The specified cluster could not be found. You can view your available clusters
//   with ListClusters. Amazon ECS clusters are region-specific.
//
func (c *ECS) UpdateContainerInstancesState(input *UpdateContainerInstancesStateInput) (*UpdateContainerInstancesStateOutput, error) {
	req, out := c.UpdateContainerInstancesStateRequest(input)
	return out, req.Send()
}

// UpdateContainerInstancesStateWithContext is the same as UpdateContainerInstancesState with the addition of
// the ability to pass a context and additional request options.
//
// See UpdateContainerInstancesState for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ECS) UpdateContainerInstancesStateWithContext(ctx aws.Context, input *UpdateContainerInstancesStateInput, opts ...request.Option) (*UpdateContainerInstancesStateOutput, error) {
	req, out := c.UpdateContainerInstancesStateRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opUpdateService = "UpdateService"

// UpdateServiceRequest generates a "aws/request.Request" representing the
// client's request for the UpdateService operation. The "output" return
// value can be used to capture response data after the request's "Send" method
// is called.
//
// See UpdateService for usage and error information.
//
// Creating a request object using this method should be used when you want to inject
// custom logic into the request's lifecycle using a custom handler, or if you want to
// access properties on the request object before or after sending the request. If
// you just want the service response, call the UpdateService method directly
// instead.
//
// Note: You must call the "Send" method on the returned request object in order
// to execute the request.
//
//    // Example sending a request using the UpdateServiceRequest method.
//    req, resp := client.UpdateServiceRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ECS) UpdateServiceRequest(input *UpdateServiceInput) (req *request.Request, output *UpdateServiceOutput) {
	op := &request.Operation{
		Name:       opUpdateService,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateServiceInput{}
	}

	output = &UpdateServiceOutput{}
	req = c.newRequest(op, input, output)
	return
}

// UpdateService API operation for Amazon EC2 Container Service.
//
// Modify the desired count or task definition used in a service.
//
// You can add to or subtract from the number of instantiations of a task definition
// in a service by specifying the cluster that the service is running in and
// a new desiredCount parameter.
//
// You can use UpdateService to modify your task definition and deploy a new
// version of your service, one task at a time. If you modify the task definition
// with UpdateService, Amazon ECS spawns a task with the new version of the
// task definition and then stops an old task after the new version is running.
// Because UpdateService starts a new version of the task before stopping an
// old version, your cluster must have capacity to support one more instantiation
// of the task when UpdateService is run. If your cluster cannot support another
// instantiation of the task used in your service, you can reduce the desired
// count of your service by one before modifying the task definition.
//
// When UpdateService replaces a task during an update, the equivalent of docker
// stop is issued to the containers running in the task. This results in a SIGTERM
// and a 30-second timeout, after which SIGKILL is sent and the containers are
// forcibly stopped. If the container handles the SIGTERM gracefully and exits
// within 30 seconds from receiving it, no SIGKILL is sent.
//
// Returns awserr.Error for service API and SDK errors. Use runtime type assertions
// with awserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the AWS API reference guide for Amazon EC2 Container Service's
// API operation UpdateService for usage and error information.
//
// Returned Error Codes:
//   * ErrCodeServerException "ServerException"
//   These errors are usually caused by a server-side issue.
//
//   * ErrCodeClientException "ClientException"
//   These errors are usually caused by something the client did, such as use
//   an action or resource on behalf of a user that doesn't have permission to
//   use the action or resource, or specify an identifier that is not valid.
//
//   * ErrCodeInvalidParameterException "InvalidParameterException"
//   The specified parameter is invalid. Review the available parameters for the
//   API request.
//
//   * ErrCodeClusterNotFoundException "ClusterNotFoundException"
//   The specified cluster could not be found. You can view your available clusters
//   with ListClusters. Amazon ECS clusters are region-specific.
//
//   * ErrCodeServiceNotFoundException "ServiceNotFoundException"
//   The specified service could not be found. You can view your available services
//   with ListServices. Amazon ECS services are cluster-specific and region-specific.
//
//   * ErrCodeServiceNotActiveException "ServiceNotActiveException"
//   The specified service is not active. You cannot update a service that is
//   not active. If you have previously deleted a service, you can recreate it
//   with CreateService.
//
func (c *ECS) UpdateService(input *UpdateServiceInput) (*UpdateServiceOutput, error) {
	req, out := c.UpdateServiceRequest(input)
	return out, req.Send()
}

// UpdateServiceWithContext is the same as UpdateService with the addition of
// the ability to pass a context and additional request options.
//
// See UpdateService for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ECS) UpdateServiceWithContext(ctx aws.Context, input *UpdateServiceInput, opts ...request.Option) (*UpdateServiceOutput, error) {
	req, out := c.UpdateServiceRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

// The attributes applicable to a container instance when it is registered.
type Attribute struct {
	_ struct{} `type:"structure"`

	// The name of the container instance attribute.
	//
	// Name is a required field
	Name *string `locationName:"name" type:"string" required:"true"`

	TargetId *string `locationName:"targetId" type:"string"`

	TargetType *string `locationName:"targetType" type:"string" enum:"TargetType"`

	// The value of the container instance attribute.
	Value *string `locationName:"value" type:"string"`
}

// String returns the string representation
func (s Attribute) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s Attribute) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *Attribute) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "Attribute"}
	if s.Name == nil {
		invalidParams.Add(request.NewErrParamRequired("Name"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetName sets the Name field's value.
func (s *Attribute) SetName(v string) *Attribute {
	s.Name = &v
	return s
}

// SetTargetId sets the TargetId field's value.
func (s *Attribute) SetTargetId(v string) *Attribute {
	s.TargetId = &v
	return s
}

// SetTargetType sets the TargetType field's value.
func (s *Attribute) SetTargetType(v string) *Attribute {
	s.TargetType = &v
	return s
}

// SetValue sets the Value field's value.
func (s *Attribute) SetValue(v string) *Attribute {
	s.Value = &v
	return s
}

// A regional grouping of one or more container instances on which you can run
// task requests. Each account receives a default cluster the first time you
// use the Amazon ECS service, but you may also create other clusters. Clusters
// may contain more than one instance type simultaneously.
type Cluster struct {
	_ struct{} `type:"structure"`

	// The number of services that are running on the cluster in an ACTIVE state.
	// You can view these services with ListServices.
	ActiveServicesCount *int64 `locationName:"activeServicesCount" type:"integer"`

	// The Amazon Resource Name (ARN) that identifies the cluster. The ARN contains
	// the arn:aws:ecs namespace, followed by the region of the cluster, the AWS
	// account ID of the cluster owner, the cluster namespace, and then the cluster
	// name. For example, arn:aws:ecs:region:012345678910:cluster/test.
	ClusterArn *string `locationName:"clusterArn" type:"string"`

	// A user-generated string that you use to identify your cluster.
	ClusterName *string `locationName:"clusterName" type:"string"`

	// The number of tasks in the cluster that are in the PENDING state.
	PendingTasksCount *int64 `locationName:"pendingTasksCount" type:"integer"`

	// The number of container instances registered into the cluster.
	RegisteredContainerInstancesCount *int64 `locationName:"registeredContainerInstancesCount" type:"integer"`

	// The number of tasks in the cluster that are in the RUNNING state.
	RunningTasksCount *int64 `locationName:"runningTasksCount" type:"integer"`

	// The status of the cluster. The valid values are ACTIVE or INACTIVE. ACTIVE
	// indicates that you can register container instances with the cluster and
	// the associated instances can accept tasks.
	Status *string `locationName:"status" type:"string"`
}

// String returns the string representation
func (s Cluster) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s Cluster) GoString() string {
	return s.String()
}

// SetActiveServicesCount sets the ActiveServicesCount field's value.
func (s *Cluster) SetActiveServicesCount(v int64) *Cluster {
	s.ActiveServicesCount = &v
	return s
}

// SetClusterArn sets the ClusterArn field's value.
func (s *Cluster) SetClusterArn(v string) *Cluster {
	s.ClusterArn = &v
	return s
}

// SetClusterName sets the ClusterName field's value.
func (s *Cluster) SetClusterName(v string) *Cluster {
	s.ClusterName = &v
	return s
}

// SetPendingTasksCount sets the PendingTasksCount field's value.
func (s *Cluster) SetPendingTasksCount(v int64) *Cluster {
	s.PendingTasksCount = &v
	return s
}

// SetRegisteredContainerInstancesCount sets the RegisteredContainerInstancesCount field's value.
func (s *Cluster) SetRegisteredContainerInstancesCount(v int64) *Cluster {
	s.RegisteredContainerInstancesCount = &v
	return s
}

// SetRunningTasksCount sets the RunningTasksCount field's value.
func (s *Cluster) SetRunningTasksCount(v int64) *Cluster {
	s.RunningTasksCount = &v
	return s
}

// SetStatus sets the Status field's value.
func (s *Cluster) SetStatus(v string) *Cluster {
	s.Status = &v
	return s
}

// A docker container that is part of a task.
type Container struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the container.
	ContainerArn *string `locationName:"containerArn" type:"string"`

	// The exit code returned from the container.
	ExitCode *int64 `locationName:"exitCode" type:"integer"`

	// The last known status of the container.
	LastStatus *string `locationName:"lastStatus" type:"string"`

	// The name of the container.
	Name *string `locationName:"name" type:"string"`

	// The network bindings associated with the container.
	NetworkBindings []*NetworkBinding `locationName:"networkBindings" type:"list"`

	// A short (255 max characters) human-readable string to provide additional
	// detail about a running or stopped container.
	Reason *string `locationName:"reason" type:"string"`

	// The Amazon Resource Name (ARN) of the task.
	TaskArn *string `locationName:"taskArn" type:"string"`
}

// String returns the string representation
func (s Container) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s Container) GoString() string {
	return s.String()
}

// SetContainerArn sets the ContainerArn field's value.
func (s *Container) SetContainerArn(v string) *Container {
	s.ContainerArn = &v
	return s
}

// SetExitCode sets the ExitCode field's value.
func (s *Container) SetExitCode(v int64) *Container {
	s.ExitCode = &v
	return s
}

// SetLastStatus sets the LastStatus field's value.
func (s *Container) SetLastStatus(v string) *Container {
	s.LastStatus = &v
	return s
}

// SetName sets the Name field's value.
func (s *Container) SetName(v string) *Container {
	s.Name = &v
	return s
}

// SetNetworkBindings sets the NetworkBindings field's value.
func (s *Container) SetNetworkBindings(v []*NetworkBinding) *Container {
	s.NetworkBindings = v
	return s
}

// SetReason sets the Reason field's value.
func (s *Container) SetReason(v string) *Container {
	s.Reason = &v
	return s
}

// SetTaskArn sets the TaskArn field's value.
func (s *Container) SetTaskArn(v string) *Container {
	s.TaskArn = &v
	return s
}

// Container definitions are used in task definitions to describe the different
// containers that are launched as part of a task.
type ContainerDefinition struct {
	_ struct{} `type:"structure"`

	// The command that is passed to the container. This parameter maps to Cmd in
	// the Create a container (https://docs.docker.com/reference/api/docker_remote_api_v1.19/#create-a-container)
	// section of the Docker Remote API (https://docs.docker.com/reference/api/docker_remote_api_v1.19/)
	// and the COMMAND parameter to docker run. For more information, see https://docs.docker.com/reference/builder/#cmd
	// (https://docs.docker.com/reference/builder/#cmd).
	Command []*string `locationName:"command" type:"list"`

	// The number of cpu units reserved for the container. A container instance
	// has 1,024 cpu units for every CPU core. This parameter specifies the minimum
	// amount of CPU to reserve for a container, and containers share unallocated
	// CPU units with other containers on the instance with the same ratio as their
	// allocated amount. This parameter maps to CpuShares in the Create a container
	// (https://docs.docker.com/reference/api/docker_remote_api_v1.19/#create-a-container)
	// section of the Docker Remote API (https://docs.docker.com/reference/api/docker_remote_api_v1.19/)
	// and the --cpu-shares option to docker run.
	//
	// You can determine the number of CPU units that are available per Amazon EC2
	// instance type by multiplying the vCPUs listed for that instance type on the
	// Amazon EC2 Instances (http://aws.amazon.com/ec2/instance-types/) detail page
	// by 1,024.
	//
	// For example, if you run a single-container task on a single-core instance
	// type with 512 CPU units specified for that container, and that is the only
	// task running on the container instance, that container could use the full
	// 1,024 CPU unit share at any given time. However, if you launched another
	// copy of the same task on that container instance, each task would be guaranteed
	// a minimum of 512 CPU units when needed, and each container could float to
	// higher CPU usage if the other container was not using it, but if both tasks
	// were 100% active all of the time, they would be limited to 512 CPU units.
	//
	// The Docker daemon on the container instance uses the CPU value to calculate
	// the relative CPU share ratios for running containers. For more information,
	// see CPU share constraint (https://docs.docker.com/reference/run/#cpu-share-constraint)
	// in the Docker documentation. The minimum valid CPU share value that the Linux
	// kernel will allow is 2; however, the CPU parameter is not required, and you
	// can use CPU values below 2 in your container definitions. For CPU values
	// below 2 (including null), the behavior varies based on your Amazon ECS container
	// agent version:
	//
	//    * Agent versions less than or equal to 1.1.0: Null and zero CPU values
	//    are passed to Docker as 0, which Docker then converts to 1,024 CPU shares.
	//    CPU values of 1 are passed to Docker as 1, which the Linux kernel converts
	//    to 2 CPU shares.
	//    * Agent versions greater than or equal to 1.2.0: Null, zero, and CPU values
	//    of 1 are passed to Docker as 2.
	Cpu *int64 `locationName:"cpu" type:"integer"`

	// When this parameter is true, networking is disabled within the container.
	// This parameter maps to NetworkDisabled in the Create a container (https://docs.docker.com/reference/api/docker_remote_api_v1.19/#create-a-container)
	// section of the Docker Remote API (https://docs.docker.com/reference/api/docker_remote_api_v1.19/).
	DisableNetworking *bool `locationName:"disableNetworking" type:"boolean"`

	// A list of DNS search domains that are presented to the container. This parameter
	// maps to DnsSearch in the Create a container (https://docs.docker.com/reference/api/docker_remote_api_v1.19/#create-a-container)
	// section of the Docker Remote API (https://docs.docker.com/reference/api/docker_remote_api_v1.19/)
	// and the --dns-search option to docker run.
	DnsSearchDomains []*string `locationName:"dnsSearchDomains" type:"list"`

	// A list of DNS servers that are presented to the container. This parameter
	// maps to Dns in the Create a container (https://docs.docker.com/reference/api/docker_remote_api_v1.19/#create-a-container)
	// section of the Docker Remote API (https://docs.docker.com/reference/api/docker_remote_api_v1.19/)
	// and the --dns option to docker run.
	DnsServers []*string `locationName:"dnsServers" type:"list"`

	// A key/value map of labels to add to the container. This parameter maps to
	// Labels in the Create a container (https://docs.docker.com/reference/api/docker_remote_api_v1.19/#create-a-container)
	// section of the Docker Remote API (https://docs.docker.com/reference/api/docker_remote_api_v1.19/)
	// and the --label option to docker run. This parameter requires version 1.18
	// of the Docker Remote API or greater on your container instance. To check
	// the Docker Remote API version on your container instance, log into your container
	// instance and run the following command: sudo docker version | grep "Server
	// API version"
	DockerLabels map[string]*string `locationName:"dockerLabels" type:"map"`

	// A list of strings to provide custom labels for SELinux and AppArmor multi-level
	// security systems. This parameter maps to SecurityOpt in the Create a container
	// (https://docs.docker.com/reference/api/docker_remote_api_v1.19/#create-a-container)
	// section of the Docker Remote API (https://docs.docker.com/reference/api/docker_remote_api_v1.19/)
	// and the --security-opt option to docker run.
	//
	// The Amazon ECS container agent running on a container instance must register
	// with the ECS_SELINUX_CAPABLE=true or ECS_APPARMOR_CAPABLE=true environment
	// variables before containers placed on that instance can use these security
	// options. For more information, see Amazon ECS Container Agent Configuration
	// (http://docs.aws.amazon.com/AmazonECS/latest/developerguide/developerguide/ecs-agent-config.html)
	// in the Amazon EC2 Container Service Developer Guide.
	DockerSecurityOptions []*string `locationName:"dockerSecurityOptions" type:"list"`

	// Early versions of the Amazon ECS container agent do not properly handle entryPoint
	// parameters. If you have problems using entryPoint, update your container
	// agent or enter your commands and arguments as command array items instead.
	//
	// The entry point that is passed to the container. This parameter maps to Entrypoint
	// in the Create a container (https://docs.docker.com/reference/api/docker_remote_api_v1.19/#create-a-container)
	// section of the Docker Remote API (https://docs.docker.com/reference/api/docker_remote_api_v1.19/)
	// and the --entrypoint option to docker run. For more information, see https://docs.docker.com/reference/builder/#entrypoint
	// (https://docs.docker.com/reference/builder/#entrypoint).
	EntryPoint []*string `locationName:"entryPoint" type:"list"`

	// The environment variables to pass to a container. This parameter maps to
	// Env in the Create a container (https://docs.docker.com/reference/api/docker_remote_api_v1.19/#create-a-container)
	// section of the Docker Remote API (https://docs.docker.com/reference/api/docker_remote_api_v1.19/)
	// and the --env option to docker run.
	Environment []*KeyValuePair `locationName:"environment" type:"list"`

	// If the essential parameter of a container is marked as true, the failure
	// of that container will stop the task. If the essential parameter of a container
	// is marked as false, then its failure will not affect the rest of the containers
	// in a task. If this parameter is omitted, a container is assumed to be essential.
	//
	// All tasks must have at least one essential container.
	Essential *bool `locationName:"essential" type:"boolean"`

	// A list of hostnames and IP address mappings to append to the /etc/hosts file
	// on the container. This parameter maps to ExtraHosts in the Create a container
	// (https://docs.docker.com/reference/api/docker_remote_api_v1.19/#create-a-container)
	// section of the Docker Remote API (https://docs.docker.com/reference/api/docker_remote_api_v1.19/)
	// and the --add-host option to docker run.
	ExtraHosts []*HostEntry `locationName:"extraHosts" type:"list"`

	// The hostname you would like to use for your container. This parameter maps
	// to Hostname in the Create a container (https://docs.docker.com/reference/api/docker_remote_api_v1.19/#create-a-container)
	// section of the Docker Remote API (https://docs.docker.com/reference/api/docker_remote_api_v1.19/)
	// and the --hostname option to docker run.
	Hostname *string `locationName:"hostname" type:"string"`

	// The image used to start a container. This string is passed directly to the
	// Docker daemon. Images in the Docker Hub registry are available by default.
	// Other repositories are specified with repository-url/image:tag. Up to 255
	// letters (uppercase and lowercase), numbers, hyphens, underscores, colons,
	// periods, forward slashes, and number signs are allowed. This parameter maps
	// to Image in the Create a container (https://docs.docker.com/reference/api/docker_remote_api_v1.19/#create-a-container)
	// section of the Docker Remote API (https://docs.docker.com/reference/api/docker_remote_api_v1.19/)
	// and the IMAGE parameter of docker run.
	Image *string `locationName:"image" type:"string"`

	// The link parameter allows containers to communicate with each other without
	// the need for port mappings, using the name parameter and optionally, an alias
	// for the link. This construct is analogous to name:alias in Docker links.
	// Up to 255 letters (uppercase and lowercase), numbers, hyphens, and underscores
	// are allowed for each name and alias. For more information on linking Docker
	// containers, see https://docs.docker.com/userguide/dockerlinks/ (https://docs.docker.com/userguide/dockerlinks/).
	// This parameter maps to PortBindings in the Create a container (https://docs.docker.com/reference/api/docker_remote_api_v1.19/#create-a-container)
	// section of the Docker Remote API (https://docs.docker.com/reference/api/docker_remote_api_v1.19/)
	// and the --link option to docker run.
	//
	// Containers that are collocated on a single container instance may be able
	// to communicate with each other without requiring links or host port mappings.
	// Network isolation is achieved on the container instance using security groups
	// and VPC settings.
	Links []*string `locationName:"links" type:"list"`

	// The log configuration specification for the container. This parameter maps
	// to LogConfig in the Create a container (https://docs.docker.com/reference/api/docker_remote_api_v1.19/#create-a-container)
	// section of the Docker Remote API (https://docs.docker.com/reference/api/docker_remote_api_v1.19/)
	// and the --log-driver option to docker run. This parameter requires version
	// 1.18 of the Docker Remote API or greater on your container instance. To check
	// the Docker Remote API version on your container instance, log into your container
	// instance and run the following command: sudo docker version | grep "Server
	// API version"
	//
	// The Amazon ECS container agent running on a container instance must register
	// the logging drivers available on that instance with the ECS_AVAILABLE_LOGGING_DRIVERS
	// environment variable before containers placed on that instance can use these
	// log configuration options. For more information, see Amazon ECS Container
	// Agent Configuration (http://docs.aws.amazon.com/AmazonECS/latest/developerguide/developerguide/ecs-agent-config.html)
	// in the Amazon EC2 Container Service Developer Guide.
	LogConfiguration *LogConfiguration `locationName:"logConfiguration" type:"structure"`

	// The number of MiB of memory reserved for the container. If your container
	// attempts to exceed the memory allocated here, the container is killed. This
	// parameter maps to Memory in the Create a container (https://docs.docker.com/reference/api/docker_remote_api_v1.19/#create-a-container)
	// section of the Docker Remote API (https://docs.docker.com/reference/api/docker_remote_api_v1.19/)
	// and the --memory option to docker run.
	Memory *int64 `locationName:"memory" type:"integer"`

	MemoryReservation *int64 `locationName:"memoryReservation" type:"integer"`

	// The mount points for data volumes in your container. This parameter maps
	// to Volumes in the Create a container (https://docs.docker.com/reference/api/docker_remote_api_v1.19/#create-a-container)
	// section of the Docker Remote API (https://docs.docker.com/reference/api/docker_remote_api_v1.19/)
	// and the --volume option to docker run.
	MountPoints []*MountPoint `locationName:"mountPoints" type:"list"`

	// The name of a container. If you are linking multiple containers together
	// in a task definition, the name of one container can be entered in the links
	// of another container to connect the containers. Up to 255 letters (uppercase
	// and lowercase), numbers, hyphens, and underscores are allowed. This parameter
	// maps to name in the Create a container (https://docs.docker.com/reference/api/docker_remote_api_v1.19/#create-a-container)
	// section of the Docker Remote API (https://docs.docker.com/reference/api/docker_remote_api_v1.19/)
	// and the --name option to docker run.
	Name *string `locationName:"name" type:"string"`

	// The list of port mappings for the container. This parameter maps to PortBindings
	// in the Create a container (https://docs.docker.com/reference/api/docker_remote_api_v1.19/#create-a-container)
	// section of the Docker Remote API (https://docs.docker.com/reference/api/docker_remote_api_v1.19/)
	// and the --publish option to docker run.
	PortMappings []*PortMapping `locationName:"portMappings" type:"list"`

	// When this parameter is true, the container is given elevated privileges on
	// the host container instance (similar to the root user). This parameter maps
	// to Privileged in the Create a container (https://docs.docker.com/reference/api/docker_remote_api_v1.19/#create-a-container)
	// section of the Docker Remote API (https://docs.docker.com/reference/api/docker_remote_api_v1.19/)
	// and the --privileged option to docker run.
	Privileged *bool `locationName:"privileged" type:"boolean"`

	// When this parameter is true, the container is given read-only access to its
	// root file system. This parameter maps to ReadonlyRootfs in the Create a container
	// (https://docs.docker.com/reference/api/docker_remote_api_v1.19/#create-a-container)
	// section of the Docker Remote API (https://docs.docker.com/reference/api/docker_remote_api_v1.19/)
	// and the --read-only option to docker run.
	ReadonlyRootFilesystem *bool `locationName:"readonlyRootFilesystem" type:"boolean"`

	// A list of ulimits to set in the container. This parameter maps to Ulimits
	// in the Create a container (https://docs.docker.com/reference/api/docker_remote_api_v1.19/#create-a-container)
	// section of the Docker Remote API (https://docs.docker.com/reference/api/docker_remote_api_v1.19/)
	// and the --ulimit option to docker run. Valid naming values are available
	// in the ulimitNameMapping variable of the pkg/ulimit/ulimit.go (https://github.com/docker/docker/blob/master/pkg/ulimit/ulimit.go)
	// file in the Docker source code. This parameter requires version 1.18 of the
	// Docker Remote API or greater on your container instance. To check the Docker
	// Remote API version on your container instance, log into your container instance
	// and run the following command: sudo docker version | grep "Server API version"
	Ulimits []*Ulimit `locationName:"ulimits" type:"list"`

	// The user name to use inside the container. This parameter maps to User in
	// the Create a container (https://docs.docker.com/reference/api/docker_remote_api_v1.19/#create-a-container)
	// section of the Docker Remote API (https://docs.docker.com/reference/api/docker_remote_api_v1.19/)
	// and the --user option to docker run.
	User *string `locationName:"user" type:"string"`

	// Data volumes to mount from another container. This parameter maps to VolumesFrom
	// in the Create a container (https://docs.docker.com/reference/api/docker_remote_api_v1.19/#create-a-container)
	// section of the Docker Remote API (https://docs.docker.com/reference/api/docker_remote_api_v1.19/)
	// and the --volumes-from option to docker run.
	VolumesFrom []*VolumeFrom `locationName:"volumesFrom" type:"list"`

	// The working directory in which to run commands inside the container. This
	// parameter maps to WorkingDir in the Create a container (https://docs.docker.com/reference/api/docker_remote_api_v1.19/#create-a-container)
	// section of the Docker Remote API (https://docs.docker.com/reference/api/docker_remote_api_v1.19/)
	// and the --workdir option to docker run.
	WorkingDirectory *string `locationName:"workingDirectory" type:"string"`
}

// String returns the string representation
func (s ContainerDefinition) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s ContainerDefinition) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ContainerDefinition) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "ContainerDefinition"}
	if s.ExtraHosts != nil {
		for i, v := range s.ExtraHosts {
			if v == nil {
				continue
			}
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "ExtraHosts", i), err.(request.ErrInvalidParams))
			}
		}
	}
	if s.LogConfiguration != nil {
		if err := s.LogConfiguration.Validate(); err != nil {
			invalidParams.AddNested("LogConfiguration", err.(request.ErrInvalidParams))
		}
	}
	if s.Ulimits != nil {
		for i, v := range s.Ulimits {
			if v == nil {
				continue
			}
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Ulimits", i), err.(request.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetCommand sets the Command field's value.
func (s *ContainerDefinition) SetCommand(v []*string) *ContainerDefinition {
	s.Command = v
	return s
}

// SetCpu sets the Cpu field's value.
func (s *ContainerDefinition) SetCpu(v int64) *ContainerDefinition {
	s.Cpu = &v
	return s
}

// SetDisableNetworking sets the DisableNetworking field's value.
func (s *ContainerDefinition) SetDisableNetworking(v bool) *ContainerDefinition {
	s.DisableNetworking = &v
	return s
}

// SetDnsSearchDomains sets the DnsSearchDomains field's value.
func (s *ContainerDefinition) SetDnsSearchDomains(v []*string) *ContainerDefinition {
	s.DnsSearchDomains = v
	return s
}

// SetDnsServers sets the DnsServers field's value.
func (s *ContainerDefinition) SetDnsServers(v []*string) *ContainerDefinition {
	s.DnsServers = v
	return s
}

// SetDockerLabels sets the DockerLabels field's value.
func (s *ContainerDefinition) SetDockerLabels(v map[string]*string) *ContainerDefinition {
	s.DockerLabels = v
	return s
}

// SetDockerSecurityOptions sets the DockerSecurityOptions field's value.
func (s *ContainerDefinition) SetDockerSecurityOptions(v []*string) *ContainerDefinition {
	s.DockerSecurityOptions = v
	return s
}

// SetEntryPoint sets the EntryPoint field's value.
func (s *ContainerDefinition) SetEntryPoint(v []*string) *ContainerDefinition {
	s.EntryPoint = v
	return s
}

// SetEnvironment sets the Environment field's value.
func (s *ContainerDefinition) SetEnvironment(v []*KeyValuePair) *ContainerDefinition {
	s.Environment = v
	return s
}

// SetEssential sets the Essential field's value.
func (s *ContainerDefinition) SetEssential(v bool) *ContainerDefinition {
	s.Essential = &v
	return s
}

// SetExtraHosts sets the ExtraHosts field's value.
func (s *ContainerDefinition) SetExtraHosts(v []*HostEntry) *ContainerDefinition {
	s.ExtraHosts = v
	return s
}

// SetHostname sets the Hostname field's value.
func (s *ContainerDefinition) SetHostname(v string) *ContainerDefinition {
	s.Hostname = &v
	return s
}

// SetImage sets the Image field's value.
func (s *ContainerDefinition) SetImage(v string) *ContainerDefinition {
	s.Image = &v
	return s
}

// SetLinks sets the Links field's value.
func (s *ContainerDefinition) SetLinks(v []*string) *ContainerDefinition {
	s.Links = v
	return s
}

// SetLogConfiguration sets the LogConfiguration field's value.
func (s *ContainerDefinition) SetLogConfiguration(v *LogConfiguration) *ContainerDefinition {
	s.LogConfiguration = v
	return s
}

// SetMemory sets the Memory field's value.
func (s *ContainerDefinition) SetMemory(v int64) *ContainerDefinition {
	s.Memory = &v
	return s
}

// SetMemoryReservation sets the MemoryReservation field's value.
func (s *ContainerDefinition) SetMemoryReservation(v int64) *ContainerDefinition {
	s.MemoryReservation = &v
	return s
}

// SetMountPoints sets the MountPoints field's value.
func (s *ContainerDefinition) SetMountPoints(v []*MountPoint) *ContainerDefinition {
	s.MountPoints = v
	return s
}

// SetName sets the Name field's value.
func (s *ContainerDefinition) SetName(v string) *ContainerDefinition {
	s.Name = &v
	return s
}

// SetPortMappings sets the PortMappings field's value.
func (s *ContainerDefinition) SetPortMappings(v []*PortMapping) *ContainerDefinition {
	s.PortMappings = v
	return s
}

// SetPrivileged sets the Privileged field's value.
func (s *ContainerDefinition) SetPrivileged(v bool) *ContainerDefinition {
	s.Privileged = &v
	return s
}

// SetReadonlyRootFilesystem sets the ReadonlyRootFilesystem field's value.
func (s *ContainerDefinition) SetReadonlyRootFilesystem(v bool) *ContainerDefinition {
	s.ReadonlyRootFilesystem = &v
	return s
}

// SetUlimits sets the Ulimits field's value.
func (s *ContainerDefinition) SetUlimits(v []*Ulimit) *ContainerDefinition {
	s.Ulimits = v
	return s
}

// SetUser sets the User field's value.
func (s *ContainerDefinition) SetUser(v string) *ContainerDefinition {
	s.User = &v
	return s
}

// SetVolumesFrom sets the VolumesFrom field's value.
func (s *ContainerDefinition) SetVolumesFrom(v []*VolumeFrom) *ContainerDefinition {
	s.VolumesFrom = v
	return s
}

// SetWorkingDirectory sets the WorkingDirectory field's value.
func (s *ContainerDefinition) SetWorkingDirectory(v string) *ContainerDefinition {
	s.WorkingDirectory = &v
	return s
}

// An Amazon EC2 instance that is running the Amazon ECS agent and has been
// registered with a cluster.
type ContainerInstance struct {
	_ struct{} `type:"structure"`

	// This parameter returns true if the agent is actually connected to Amazon
	// ECS. Registered instances with an agent that may be unhealthy or stopped
	// will return false, and instances without a connected agent cannot accept
	// placement request.
	AgentConnected *bool `locationName:"agentConnected" type:"boolean"`

	// The status of the most recent agent update. If an update has never been requested,
	// this value is NULL.
	AgentUpdateStatus *string `locationName:"agentUpdateStatus" type:"string" enum:"AgentUpdateStatus"`

	// The attributes set for the container instance by the Amazon ECS container
	// agent at instance registration.
	Attributes []*Attribute `locationName:"attributes" type:"list"`

	// The Amazon Resource Name (ARN) of the container instance. The ARN contains
	// the arn:aws:ecs namespace, followed by the region of the container instance,
	// the AWS account ID of the container instance owner, the container-instance
	// namespace, and then the container instance ID. For example, arn:aws:ecs:region:aws_account_id:container-instance/container_instance_ID.
	ContainerInstanceArn *string `locationName:"containerInstanceArn" type:"string"`

	// The Amazon EC2 instance ID of the container instance.
	Ec2InstanceId *string `locationName:"ec2InstanceId" type:"string"`

	// The number of tasks on the container instance that are in the PENDING status.
	PendingTasksCount *int64 `locationName:"pendingTasksCount" type:"integer"`

	RegisteredAt *time.Time `locationName:"registeredAt" type:"timestamp" timestampFormat:"unix"`

	// The registered resources on the container instance that are in use by current
	// tasks.
	RegisteredResources []*Resource `locationName:"registeredResources" type:"list"`

	// The remaining resources of the container instance that are available for
	// new tasks.
	RemainingResources []*Resource `locationName:"remainingResources" type:"list"`

	// The number of tasks on the container instance that are in the RUNNING status.
	RunningTasksCount *int64 `locationName:"runningTasksCount" type:"integer"`

	// The status of the container instance. The valid values are ACTIVE or INACTIVE.
	// ACTIVE indicates that the container instance can accept tasks.
	Status *string `locationName:"status" type:"string"`

	Version *int64 `locationName:"version" type:"long"`

	// The version information for the Amazon ECS container agent and Docker daemon
	// running on the container instance.
	VersionInfo *VersionInfo `locationName:"versionInfo" type:"structure"`
}

// String returns the string representation
func (s ContainerInstance) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s ContainerInstance) GoString() string {
	return s.String()
}

// SetAgentConnected sets the AgentConnected field's value.
func (s *ContainerInstance) SetAgentConnected(v bool) *ContainerInstance {
	s.AgentConnected = &v
	return s
}

// SetAgentUpdateStatus sets the AgentUpdateStatus field's value.
func (s *ContainerInstance) SetAgentUpdateStatus(v string) *ContainerInstance {
	s.AgentUpdateStatus = &v
	return s
}

// SetAttributes sets the Attributes field's value.
func (s *ContainerInstance) SetAttributes(v []*Attribute) *ContainerInstance {
	s.Attributes = v
	return s
}

// SetContainerInstanceArn sets the ContainerInstanceArn field's value.
func (s *ContainerInstance) SetContainerInstanceArn(v string) *ContainerInstance {
	s.ContainerInstanceArn = &v
	return s
}

// SetEc2InstanceId sets the Ec2InstanceId field's value.
func (s *ContainerInstance) SetEc2InstanceId(v string) *ContainerInstance {
	s.Ec2InstanceId = &v
	return s
}

// SetPendingTasksCount sets the PendingTasksCount field's value.
func (s *ContainerInstance) SetPendingTasksCount(v int64) *ContainerInstance {
	s.PendingTasksCount = &v
	return s
}

// SetRegisteredAt sets the RegisteredAt field's value.
func (s *ContainerInstance) SetRegisteredAt(v time.Time) *ContainerInstance {
	s.RegisteredAt = &v
	return s
}

// SetRegisteredResources sets the RegisteredResources field's value.
func (s *ContainerInstance) SetRegisteredResources(v []*Resource) *ContainerInstance {
	s.RegisteredResources = v
	return s
}

// SetRemainingResources sets the RemainingResources field's value.
func (s *ContainerInstance) SetRemainingResources(v []*Resource) *ContainerInstance {
	s.RemainingResources = v
	return s
}

// SetRunningTasksCount sets the RunningTasksCount field's value.
func (s *ContainerInstance) SetRunningTasksCount(v int64) *ContainerInstance {
	s.RunningTasksCount = &v
	return s
}

// SetStatus sets the Status field's value.
func (s *ContainerInstance) SetStatus(v string) *ContainerInstance {
	s.Status = &v
	return s
}

// SetVersion sets the Version field's value.
func (s *ContainerInstance) SetVersion(v int64) *ContainerInstance {
	s.Version = &v
	return s
}

// SetVersionInfo sets the VersionInfo field's value.
func (s *ContainerInstance) SetVersionInfo(v *VersionInfo) *ContainerInstance {
	s.VersionInfo = v
	return s
}

// The overrides that should be sent to a container.
type ContainerOverride struct {
	_ struct{} `type:"structure"`

	// The command to send to the container that overrides the default command from
	// the Docker image or the task definition.
	Command []*string `locationName:"command" type:"list"`

	Cpu *int64 `locationName:"cpu" type:"integer"`

	// The environment variables to send to the container. You can add new environment
	// variables, which are added to the container at launch, or you can override
	// the existing environment variables from the Docker image or the task definition.
	Environment []*KeyValuePair `locationName:"environment" type:"list"`

	Memory *int64 `locationName:"memory" type:"integer"`

	MemoryReservation *int64 `locationName:"memoryReservation" type:"integer"`

	// The name of the container that receives the override.
	Name *string `locationName:"name" type:"string"`
}

// String returns the string representation
func (s ContainerOverride) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s ContainerOverride) GoString() string {
	return s.String()
}

// SetCommand sets the Command field's value.
func (s *ContainerOverride) SetCommand(v []*string) *ContainerOverride {
	s.Command = v
	return s
}

// SetCpu sets the Cpu field's value.
func (s *ContainerOverride) SetCpu(v int64) *ContainerOverride {
	s.Cpu = &v
	return s
}

// SetEnvironment sets the Environment field's value.
func (s *ContainerOverride) SetEnvironment(v []*KeyValuePair) *ContainerOverride {
	s.Environment = v
	return s
}

// SetMemory sets the Memory field's value.
func (s *ContainerOverride) SetMemory(v int64) *ContainerOverride {
	s.Memory = &v
	return s
}

// SetMemoryReservation sets the MemoryReservation field's value.
func (s *ContainerOverride) SetMemoryReservation(v int64) *ContainerOverride {
	s.MemoryReservation = &v
	return s
}

// SetName sets the Name field's value.
func (s *ContainerOverride) SetName(v string) *ContainerOverride {
	s.Name = &v
	return s
}

type ContainerStateChange struct {
	_ struct{} `type:"structure"`

	ContainerName *string `locationName:"containerName" type:"string"`

	ExitCode *int64 `locationName:"exitCode" type:"integer"`

	NetworkBindings []*NetworkBinding `locationName:"networkBindings" type:"list"`

	Reason *string `locationName:"reason" type:"string"`

	Status *string `locationName:"status" type:"string"`
}

// String returns the string representation
func (s ContainerStateChange) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s ContainerStateChange) GoString() string {
	return s.String()
}

// SetContainerName sets the ContainerName field's value.
func (s *ContainerStateChange) SetContainerName(v string) *ContainerStateChange {
	s.ContainerName = &v
	return s
}

// SetExitCode sets the ExitCode field's value.
func (s *ContainerStateChange) SetExitCode(v int64) *ContainerStateChange {
	s.ExitCode = &v
	return s
}

// SetNetworkBindings sets the NetworkBindings field's value.
func (s *ContainerStateChange) SetNetworkBindings(v []*NetworkBinding) *ContainerStateChange {
	s.NetworkBindings = v
	return s
}

// SetReason sets the Reason field's value.
func (s *ContainerStateChange) SetReason(v string) *ContainerStateChange {
	s.Reason = &v
	return s
}

// SetStatus sets the Status field's value.
func (s *ContainerStateChange) SetStatus(v string) *ContainerStateChange {
	s.Status = &v
	return s
}

type CreateClusterInput struct {
	_ struct{} `type:"structure"`

	// The name of your cluster. If you do not specify a name for your cluster,
	// you will create a cluster named default. Up to 255 letters (uppercase and
	// lowercase), numbers, hyphens, and underscores are allowed.
	ClusterName *string `locationName:"clusterName" type:"string"`
}

// String returns the string representation
func (s CreateClusterInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateClusterInput) GoString() string {
	return s.String()
}

// SetClusterName sets the ClusterName field's value.
func (s *CreateClusterInput) SetClusterName(v string) *CreateClusterInput {
	s.ClusterName = &v
	return s
}

type CreateClusterOutput struct {
	_ struct{} `type:"structure"`

	// The full description of your new cluster.
	Cluster *Cluster `locationName:"cluster" type:"structure"`
}

// String returns the string representation
func (s CreateClusterOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateClusterOutput) GoString() string {
	return s.String()
}

// SetCluster sets the Cluster field's value.
func (s *CreateClusterOutput) SetCluster(v *Cluster) *CreateClusterOutput {
	s.Cluster = v
	return s
}

type CreateServiceInput struct {
	_ struct{} `type:"structure"`

	// Unique, case-sensitive identifier you provide to ensure the idempotency of
	// the request. Up to 32 ASCII characters are allowed.
	ClientToken *string `locationName:"clientToken" type:"string"`

	// The short name or full Amazon Resource Name (ARN) of the cluster that you
	// want to run your service on.
	Cluster *string `locationName:"cluster" type:"string"`

	DeploymentConfiguration *DeploymentConfiguration `locationName:"deploymentConfiguration" type:"structure"`

	// The number of instantiations of the specified task definition that you would
	// like to place and keep running on your cluster.
	//
	// DesiredCount is a required field
	DesiredCount *int64 `locationName:"desiredCount" type:"integer" required:"true"`

	LaunchType *string `locationName:"launchType" type:"string" enum:"LaunchType"`

	// A list of load balancer objects, containing the load balancer name, the container
	// name (as it appears in a container definition), and the container port to
	// access from the load balancer.
	LoadBalancers []*LoadBalancer `locationName:"loadBalancers" type:"list"`

	PlacementConstraints []*PlacementConstraint `locationName:"placementConstraints" type:"list"`

	PlacementStrategy []*PlacementStrategy `locationName:"placementStrategy" type:"list"`

	// The name or full Amazon Resource Name (ARN) of the IAM role that allows your
	// Amazon ECS container agent to make calls to your load balancer on your behalf.
	// This parameter is only required if you are using a load balancer with your
	// service.
	Role *string `locationName:"role" type:"string"`

	// The name of your service. Up to 255 letters (uppercase and lowercase), numbers,
	// hyphens, and underscores are allowed. Service names must be unique within
	// a cluster, but you can have similarly named services in multiple clusters
	// within a region or across multiple regions.
	//
	// ServiceName is a required field
	ServiceName *string `locationName:"serviceName" type:"string" required:"true"`

	// The family and revision (family:revision) or full Amazon Resource Name (ARN)
	// of the task definition that you want to run in your service. If a revision
	// is not specified, the latest ACTIVE revision is used.
	//
	// TaskDefinition is a required field
	TaskDefinition *string `locationName:"taskDefinition" type:"string" required:"true"`
}

// String returns the string representation
func (s CreateServiceInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateServiceInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateServiceInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "CreateServiceInput"}
	if s.DesiredCount == nil {
		invalidParams.Add(request.NewErrParamRequired("DesiredCount"))
	}
	if s.ServiceName == nil {
		invalidParams.Add(request.NewErrParamRequired("ServiceName"))
	}
	if s.TaskDefinition == nil {
		invalidParams.Add(request.NewErrParamRequired("TaskDefinition"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetClientToken sets the ClientToken field's value.
func (s *CreateServiceInput) SetClientToken(v string) *CreateServiceInput {
	s.ClientToken = &v
	return s
}

// SetCluster sets the Cluster field's value.
func (s *CreateServiceInput) SetCluster(v string) *CreateServiceInput {
	s.Cluster = &v
	return s
}

// SetDeploymentConfiguration sets the DeploymentConfiguration field's value.
func (s *CreateServiceInput) SetDeploymentConfiguration(v *DeploymentConfiguration) *CreateServiceInput {
	s.DeploymentConfiguration = v
	return s
}

// SetDesiredCount sets the DesiredCount field's value.
func (s *CreateServiceInput) SetDesiredCount(v int64) *CreateServiceInput {
	s.DesiredCount = &v
	return s
}

// SetLaunchType sets the LaunchType field's value.
func (s *CreateServiceInput) SetLaunchType(v string) *CreateServiceInput {
	s.LaunchType = &v
	return s
}

// SetLoadBalancers sets the LoadBalancers field's value.
func (s *CreateServiceInput) SetLoadBalancers(v []*LoadBalancer) *CreateServiceInput {
	s.LoadBalancers = v
	return s
}

// SetPlacementConstraints sets the PlacementConstraints field's value.
func (s *CreateServiceInput) SetPlacementConstraints(v []*PlacementConstraint) *CreateServiceInput {
	s.PlacementConstraints = v
	return s
}

// SetPlacementStrategy sets the PlacementStrategy field's value.
func (s *CreateServiceInput) SetPlacementStrategy(v []*PlacementStrategy) *CreateServiceInput {
	s.PlacementStrategy = v
	return s
}

// SetRole sets the Role field's value.
func (s *CreateServiceInput) SetRole(v string) *CreateServiceInput {
	s.Role = &v
	return s
}

// SetServiceName sets the ServiceName field's value.
func (s *CreateServiceInput) SetServiceName(v string) *CreateServiceInput {
	s.ServiceName = &v
	return s
}

// SetTaskDefinition sets the TaskDefinition field's value.
func (s *CreateServiceInput) SetTaskDefinition(v string) *CreateServiceInput {
	s.TaskDefinition = &v
	return s
}

type CreateServiceOutput struct {
	_ struct{} `type:"structure"`

	// The full description of your service following the create call.
	Service *Service `locationName:"service" type:"structure"`
}

// String returns the string representation
func (s CreateServiceOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateServiceOutput) GoString() string {
	return s.String()
}

// SetService sets the Service field's value.
func (s *CreateServiceOutput) SetService(v *Service) *CreateServiceOutput {
	s.Service = v
	return s
}

type DeleteAttributesInput struct {
	_ struct{} `type:"structure"`

	// Attributes is a required field
	Attributes []*Attribute `locationName:"attributes" type:"list" required:"true"`

	Cluster *string `locationName:"cluster" type:"string"`
}

// String returns the string representation
func (s DeleteAttributesInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteAttributesInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteAttributesInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DeleteAttributesInput"}
	if s.Attributes == nil {
		invalidParams.Add(request.NewErrParamRequired("Attributes"))
	}
	if s.Attributes != nil {
		for i, v := range s.Attributes {
			if v == nil {
				continue
			}
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Attributes", i), err.(request.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetAttributes sets the Attributes field's value.
func (s *DeleteAttributesInput) SetAttributes(v []*Attribute) *DeleteAttributesInput {
	s.Attributes = v
	return s
}

// SetCluster sets the Cluster field's value.
func (s *DeleteAttributesInput) SetCluster(v string) *DeleteAttributesInput {
	s.Cluster = &v
	return s
}

type DeleteAttributesOutput struct {
	_ struct{} `type:"structure"`

	Attributes []*Attribute `locationName:"attributes" type:"list"`
}

// String returns the string representation
func (s DeleteAttributesOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteAttributesOutput) GoString() string {
	return s.String()
}

// SetAttributes sets the Attributes field's value.
func (s *DeleteAttributesOutput) SetAttributes(v []*Attribute) *DeleteAttributesOutput {
	s.Attributes = v
	return s
}

type DeleteClusterInput struct {
	_ struct{} `type:"structure"`

	// The short name or full Amazon Resource Name (ARN) of the cluster that you
	// want to delete.
	//
	// Cluster is a required field
	Cluster *string `locationName:"cluster" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteClusterInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteClusterInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteClusterInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DeleteClusterInput"}
	if s.Cluster == nil {
		invalidParams.Add(request.NewErrParamRequired("Cluster"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetCluster sets the Cluster field's value.
func (s *DeleteClusterInput) SetCluster(v string) *DeleteClusterInput {
	s.Cluster = &v
	return s
}

type DeleteClusterOutput struct {
	_ struct{} `type:"structure"`

	// The full description of the deleted cluster.
	Cluster *Cluster `locationName:"cluster" type:"structure"`
}

// String returns the string representation
func (s DeleteClusterOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteClusterOutput) GoString() string {
	return s.String()
}

// SetCluster sets the Cluster field's value.
func (s *DeleteClusterOutput) SetCluster(v *Cluster) *DeleteClusterOutput {
	s.Cluster = v
	return s
}

type DeleteServiceInput struct {
	_ struct{} `type:"structure"`

	// The name of the cluster that hosts the service you want to delete. If you
	// do not specify a cluster, the default cluster is assumed.
	Cluster *string `locationName:"cluster" type:"string"`

	// The name of the service you want to delete.
	//
	// Service is a required field
	Service *string `locationName:"service" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteServiceInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteServiceInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteServiceInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DeleteServiceInput"}
	if s.Service == nil {
		invalidParams.Add(request.NewErrParamRequired("Service"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetCluster sets the Cluster field's value.
func (s *DeleteServiceInput) SetCluster(v string) *DeleteServiceInput {
	s.Cluster = &v
	return s
}

// SetService sets the Service field's value.
func (s *DeleteServiceInput) SetService(v string) *DeleteServiceInput {
	s.Service = &v
	return s
}

type DeleteServiceOutput struct {
	_ struct{} `type:"structure"`

	// The full description of the deleted service.
	Service *Service `locationName:"service" type:"structure"`
}

// String returns the string representation
func (s DeleteServiceOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteServiceOutput) GoString() string {
	return s.String()
}

// SetService sets the Service field's value.
func (s *DeleteServiceOutput) SetService(v *Service) *DeleteServiceOutput {
	s.Service = v
	return s
}

// The details of an Amazon ECS service deployment.
type Deployment struct {
	_ struct{} `type:"structure"`

	// The Unix time in seconds and milliseconds when the service was created.
	CreatedAt *time.Time `locationName:"createdAt" type:"timestamp" timestampFormat:"unix"`

	// The most recent desired count of tasks that was specified for the service
	// to deploy and/or maintain.
	DesiredCount *int64 `locationName:"desiredCount" type:"integer"`

	// The ID of the deployment.
	Id *string `locationName:"id" type:"string"`

	LaunchType *string `locationName:"launchType" type:"string" enum:"LaunchType"`

	// The number of tasks in the deployment that are in the PENDING status.
	PendingCount *int64 `locationName:"pendingCount" type:"integer"`

	// The number of tasks in the deployment that are in the RUNNING status.
	RunningCount *int64 `locationName:"runningCount" type:"integer"`

	// The status of the deployment. Valid values are PRIMARY (for the most recent
	// deployment), ACTIVE (for previous deployments that still have tasks running,
	// but are being replaced with the PRIMARY deployment), and INACTIVE (for deployments
	// that have been completely replaced).
	Status *string `locationName:"status" type:"string"`

	// The most recent task definition that was specified for the service to use.
	TaskDefinition *string `locationName:"taskDefinition" type:"string"`

	// The Unix time in seconds and milliseconds when the service was last updated.
	UpdatedAt *time.Time `locationName:"updatedAt" type:"timestamp" timestampFormat:"unix"`
}

// String returns the string representation
func (s Deployment) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s Deployment) GoString() string {
	return s.String()
}

// SetCreatedAt sets the CreatedAt field's value.
func (s *Deployment) SetCreatedAt(v time.Time) *Deployment {
	s.CreatedAt = &v
	return s
}

// SetDesiredCount sets the DesiredCount field's value.
func (s *Deployment) SetDesiredCount(v int64) *Deployment {
	s.DesiredCount = &v
	return s
}

// SetId sets the Id field's value.
func (s *Deployment) SetId(v string) *Deployment {
	s.Id = &v
	return s
}

// SetLaunchType sets the LaunchType field's value.
func (s *Deployment) SetLaunchType(v string) *Deployment {
	s.LaunchType = &v
	return s
}

// SetPendingCount sets the PendingCount field's value.
func (s *Deployment) SetPendingCount(v int64) *Deployment {
	s.PendingCount = &v
	return s
}

// SetRunningCount sets the RunningCount field's value.
func (s *Deployment) SetRunningCount(v int64) *Deployment {
	s.RunningCount = &v
	return s
}

// SetStatus sets the Status field's value.
func (s *Deployment) SetStatus(v string) *Deployment {
	s.Status = &v
	return s
}

// SetTaskDefinition sets the TaskDefinition field's value.
func (s *Deployment) SetTaskDefinition(v string) *Deployment {
	s.TaskDefinition = &v
	return s
}

// SetUpdatedAt sets the UpdatedAt field's value.
func (s *Deployment) SetUpdatedAt(v time.Time) *Deployment {
	s.UpdatedAt = &v
	return s
}

type DeploymentConfiguration struct {
	_ struct{} `type:"structure"`

	MaximumPercent *int64 `locationName:"maximumPercent" type:"integer"`

	MinimumHealthyPercent *int64 `locationName:"minimumHealthyPercent" type:"integer"`
}

// String returns the string representation
func (s DeploymentConfiguration) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s DeploymentConfiguration) GoString() string {
	return s.String()
}

// SetMaximumPercent sets the MaximumPercent field's value.
func (s *DeploymentConfiguration) SetMaximumPercent(v int64) *DeploymentConfiguration {
	s.MaximumPercent = &v
	return s
}

// SetMinimumHealthyPercent sets the MinimumHealthyPercent field's value.
func (s *DeploymentConfiguration) SetMinimumHealthyPercent(v int64) *DeploymentConfiguration {
	s.MinimumHealthyPercent = &v
	return s
}

type DeregisterContainerInstanceInput struct {
	_ struct{} `type:"structure"`

	// The short name or full Amazon Resource Name (ARN) of the cluster that hosts
	// the container instance you want to deregister. If you do not specify a cluster,
	// the default cluster is assumed.
	Cluster *string `locationName:"cluster" type:"string"`

	// The container instance ID or full Amazon Resource Name (ARN) of the container
	// instance you want to deregister. The ARN contains the arn:aws:ecs namespace,
	// followed by the region of the container instance, the AWS account ID of the
	// container instance owner, the container-instance namespace, and then the
	// container instance ID. For example, arn:aws:ecs:region:aws_account_id:container-instance/container_instance_ID.
	//
	// ContainerInstance is a required field
	ContainerInstance *string `locationName:"containerInstance" type:"string" required:"true"`

	// Force the deregistration of the container instance. If you have tasks running
	// on the container instance when you deregister it with the force option, these
	// tasks remain running and they will continue to pass Elastic Load Balancing
	// load balancer health checks until you terminate the instance or the tasks
	// stop through some other means, but they are orphaned (no longer monitored
	// or accounted for by Amazon ECS). If an orphaned task on your container instance
	// is part of an Amazon ECS service, then the service scheduler will start another
	// copy of that task on a different container instance if possible.
	Force *bool `locationName:"force" type:"boolean"`
}

// String returns the string representation
func (s DeregisterContainerInstanceInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s DeregisterContainerInstanceInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeregisterContainerInstanceInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DeregisterContainerInstanceInput"}
	if s.ContainerInstance == nil {
		invalidParams.Add(request.NewErrParamRequired("ContainerInstance"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetCluster sets the Cluster field's value.
func (s *DeregisterContainerInstanceInput) SetCluster(v string) *DeregisterContainerInstanceInput {
	s.Cluster = &v
	return s
}

// SetContainerInstance sets the ContainerInstance field's value.
func (s *DeregisterContainerInstanceInput) SetContainerInstance(v string) *DeregisterContainerInstanceInput {
	s.ContainerInstance = &v
	return s
}

// SetForce sets the Force field's value.
func (s *DeregisterContainerInstanceInput) SetForce(v bool) *DeregisterContainerInstanceInput {
	s.Force = &v
	return s
}

type DeregisterContainerInstanceOutput struct {
	_ struct{} `type:"structure"`

	// An Amazon EC2 instance that is running the Amazon ECS agent and has been
	// registered with a cluster.
	ContainerInstance *ContainerInstance `locationName:"containerInstance" type:"structure"`
}

// String returns the string representation
func (s DeregisterContainerInstanceOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s DeregisterContainerInstanceOutput) GoString() string {
	return s.String()
}

// SetContainerInstance sets the ContainerInstance field's value.
func (s *DeregisterContainerInstanceOutput) SetContainerInstance(v *ContainerInstance) *DeregisterContainerInstanceOutput {
	s.ContainerInstance = v
	return s
}

type DeregisterTaskDefinitionInput struct {
	_ struct{} `type:"structure"`

	// The family and revision (family:revision) or full Amazon Resource Name (ARN)
	// of the task definition that you want to deregister. You must specify a revision.
	//
	// TaskDefinition is a required field
	TaskDefinition *string `locationName:"taskDefinition" type:"string" required:"true"`
}

// String returns the string representation
func (s DeregisterTaskDefinitionInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s DeregisterTaskDefinitionInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeregisterTaskDefinitionInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DeregisterTaskDefinitionInput"}
	if s.TaskDefinition == nil {
		invalidParams.Add(request.NewErrParamRequired("TaskDefinition"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetTaskDefinition sets the TaskDefinition field's value.
func (s *DeregisterTaskDefinitionInput) SetTaskDefinition(v string) *DeregisterTaskDefinitionInput {
	s.TaskDefinition = &v
	return s
}

type DeregisterTaskDefinitionOutput struct {
	_ struct{} `type:"structure"`

	// The full description of the deregistered task.
	TaskDefinition *TaskDefinition `locationName:"taskDefinition" type:"structure"`
}

// String returns the string representation
func (s DeregisterTaskDefinitionOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s DeregisterTaskDefinitionOutput) GoString() string {
	return s.String()
}

// SetTaskDefinition sets the TaskDefinition field's value.
func (s *DeregisterTaskDefinitionOutput) SetTaskDefinition(v *TaskDefinition) *DeregisterTaskDefinitionOutput {
	s.TaskDefinition = v
	return s
}

type DescribeClustersInput struct {
	_ struct{} `type:"structure"`

	// A space-separated list of cluster names or full cluster Amazon Resource Name
	// (ARN) entries. If you do not specify a cluster, the default cluster is assumed.
	Clusters []*string `locationName:"clusters" type:"list"`
}

// String returns the string representation
func (s DescribeClustersInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeClustersInput) GoString() string {
	return s.String()
}

// SetClusters sets the Clusters field's value.
func (s *DescribeClustersInput) SetClusters(v []*string) *DescribeClustersInput {
	s.Clusters = v
	return s
}

type DescribeClustersOutput struct {
	_ struct{} `type:"structure"`

	// The list of clusters.
	Clusters []*Cluster `locationName:"clusters" type:"list"`

	// Any failures associated with the call.
	Failures []*Failure `locationName:"failures" type:"list"`
}

// String returns the string representation
func (s DescribeClustersOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeClustersOutput) GoString() string {
	return s.String()
}

// SetClusters sets the Clusters field's value.
func (s *DescribeClustersOutput) SetClusters(v []*Cluster) *DescribeClustersOutput {
	s.Clusters = v
	return s
}

// SetFailures sets the Failures field's value.
func (s *DescribeClustersOutput) SetFailures(v []*Failure) *DescribeClustersOutput {
	s.Failures = v
	return s
}

type DescribeContainerInstancesInput struct {
	_ struct{} `type:"structure"`

	// The short name or full Amazon Resource Name (ARN) of the cluster that hosts
	// the container instances you want to describe. If you do not specify a cluster,
	// the default cluster is assumed.
	Cluster *string `locationName:"cluster" type:"string"`

	// A space-separated list of container instance IDs or full Amazon Resource
	// Name (ARN) entries.
	//
	// ContainerInstances is a required field
	ContainerInstances []*string `locationName:"containerInstances" type:"list" required:"true"`
}

// String returns the string representation
func (s DescribeContainerInstancesInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeContainerInstancesInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeContainerInstancesInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DescribeContainerInstancesInput"}
	if s.ContainerInstances == nil {
		invalidParams.Add(request.NewErrParamRequired("ContainerInstances"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetCluster sets the Cluster field's value.
func (s *DescribeContainerInstancesInput) SetCluster(v string) *DescribeContainerInstancesInput {
	s.Cluster = &v
	return s
}

// SetContainerInstances sets the ContainerInstances field's value.
func (s *DescribeContainerInstancesInput) SetContainerInstances(v []*string) *DescribeContainerInstancesInput {
	s.ContainerInstances = v
	return s
}

type DescribeContainerInstancesOutput struct {
	_ struct{} `type:"structure"`

	// The list of container instances.
	ContainerInstances []*ContainerInstance `locationName:"containerInstances" type:"list"`

	// Any failures associated with the call.
	Failures []*Failure `locationName:"failures" type:"list"`
}

// String returns the string representation
func (s DescribeContainerInstancesOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeContainerInstancesOutput) GoString() string {
	return s.String()
}

// SetContainerInstances sets the ContainerInstances field's value.
func (s *DescribeContainerInstancesOutput) SetContainerInstances(v []*ContainerInstance) *DescribeContainerInstancesOutput {
	s.ContainerInstances = v
	return s
}

// SetFailures sets the Failures field's value.
func (s *DescribeContainerInstancesOutput) SetFailures(v []*Failure) *DescribeContainerInstancesOutput {
	s.Failures = v
	return s
}

type DescribeServicesInput struct {
	_ struct{} `type:"structure"`

	// The name of the cluster that hosts the service you want to describe. If you
	// do not specify a cluster, the default cluster is assumed.
	Cluster *string `locationName:"cluster" type:"string"`

	// A list of services you want to describe.
	//
	// Services is a required field
	Services []*string `locationName:"services" type:"list" required:"true"`
}

// String returns the string representation
func (s DescribeServicesInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeServicesInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeServicesInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DescribeServicesInput"}
	if s.Services == nil {
		invalidParams.Add(request.NewErrParamRequired("Services"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetCluster sets the Cluster field's value.
func (s *DescribeServicesInput) SetCluster(v string) *DescribeServicesInput {
	s.Cluster = &v
	return s
}

// SetServices sets the Services field's value.
func (s *DescribeServicesInput) SetServices(v []*string) *DescribeServicesInput {
	s.Services = v
	return s
}

type DescribeServicesOutput struct {
	_ struct{} `type:"structure"`

	// Any failures associated with the call.
	Failures []*Failure `locationName:"failures" type:"list"`

	// The list of services described.
	Services []*Service `locationName:"services" type:"list"`
}

// String returns the string representation
func (s DescribeServicesOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeServicesOutput) GoString() string {
	return s.String()
}

// SetFailures sets the Failures field's value.
func (s *DescribeServicesOutput) SetFailures(v []*Failure) *DescribeServicesOutput {
	s.Failures = v
	return s
}

// SetServices sets the Services field's value.
func (s *DescribeServicesOutput) SetServices(v []*Service) *DescribeServicesOutput {
	s.Services = v
	return s
}

type DescribeTaskDefinitionInput struct {
	_ struct{} `type:"structure"`

	// The family for the latest ACTIVE revision, family and revision (family:revision)
	// for a specific revision in the family, or full Amazon Resource Name (ARN)
	// of the task definition that you want to describe.
	//
	// TaskDefinition is a required field
	TaskDefinition *string `locationName:"taskDefinition" type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeTaskDefinitionInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeTaskDefinitionInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeTaskDefinitionInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DescribeTaskDefinitionInput"}
	if s.TaskDefinition == nil {
		invalidParams.Add(request.NewErrParamRequired("TaskDefinition"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetTaskDefinition sets the TaskDefinition field's value.
func (s *DescribeTaskDefinitionInput) SetTaskDefinition(v string) *DescribeTaskDefinitionInput {
	s.TaskDefinition = &v
	return s
}

type DescribeTaskDefinitionOutput struct {
	_ struct{} `type:"structure"`

	// The full task definition description.
	TaskDefinition *TaskDefinition `locationName:"taskDefinition" type:"structure"`
}

// String returns the string representation
func (s DescribeTaskDefinitionOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeTaskDefinitionOutput) GoString() string {
	return s.String()
}

// SetTaskDefinition sets the TaskDefinition field's value.
func (s *DescribeTaskDefinitionOutput) SetTaskDefinition(v *TaskDefinition) *DescribeTaskDefinitionOutput {
	s.TaskDefinition = v
	return s
}

type DescribeTasksInput struct {
	_ struct{} `type:"structure"`

	// The short name or full Amazon Resource Name (ARN) of the cluster that hosts
	// the task you want to describe. If you do not specify a cluster, the default
	// cluster is assumed.
	Cluster *string `locationName:"cluster" type:"string"`

	// A space-separated list of task IDs or full Amazon Resource Name (ARN) entries.
	//
	// Tasks is a required field
	Tasks []*string `locationName:"tasks" type:"list" required:"true"`
}

// String returns the string representation
func (s DescribeTasksInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeTasksInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeTasksInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DescribeTasksInput"}
	if s.Tasks == nil {
		invalidParams.Add(request.NewErrParamRequired("Tasks"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetCluster sets the Cluster field's value.
func (s *DescribeTasksInput) SetCluster(v string) *DescribeTasksInput {
	s.Cluster = &v
	return s
}

// SetTasks sets the Tasks field's value.
func (s *DescribeTasksInput) SetTasks(v []*string) *DescribeTasksInput {
	s.Tasks = v
	return s
}

type DescribeTasksOutput struct {
	_ struct{} `type:"structure"`

	// Any failures associated with the call.
	Failures []*Failure `locationName:"failures" type:"list"`

	// The list of tasks.
	Tasks []*Task `locationName:"tasks" type:"list"`
}

// String returns the string representation
func (s DescribeTasksOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeTasksOutput) GoString() string {
	return s.String()
}

// SetFailures sets the Failures field's value.
func (s *DescribeTasksOutput) SetFailures(v []*Failure) *DescribeTasksOutput {
	s.Failures = v
	return s
}

// SetTasks sets the Tasks field's value.
func (s *DescribeTasksOutput) SetTasks(v []*Task) *DescribeTasksOutput {
	s.Tasks = v
	return s
}

type DiscoverPollEndpointInput struct {
	_ struct{} `type:"structure"`

	// The cluster that the container instance belongs to.
	Cluster *string `locationName:"cluster" type:"string"`

	// The container instance ID or full Amazon Resource Name (ARN) of the container
	// instance. The ARN contains the arn:aws:ecs namespace, followed by the region
	// of the container instance, the AWS account ID of the container instance owner,
	// the container-instance namespace, and then the container instance ID. For
	// example, arn:aws:ecs:region:aws_account_id:container-instance/container_instance_ID.
	ContainerInstance *string `locationName:"containerInstance" type:"string"`
}

// String returns the string representation
func (s DiscoverPollEndpointInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s DiscoverPollEndpointInput) GoString() string {
	return s.String()
}

// SetCluster sets the Cluster field's value.
func (s *DiscoverPollEndpointInput) SetCluster(v string) *DiscoverPollEndpointInput {
	s.Cluster = &v
	return s
}

// SetContainerInstance sets the ContainerInstance field's value.
func (s *DiscoverPollEndpointInput) SetContainerInstance(v string) *DiscoverPollEndpointInput {
	s.ContainerInstance = &v
	return s
}

type DiscoverPollEndpointOutput struct {
	_ struct{} `type:"structure"`

	// The endpoint for the Amazon ECS agent to poll.
	Endpoint *string `locationName:"endpoint" type:"string"`

	// The telemetry endpoint for the Amazon ECS agent.
	TelemetryEndpoint *string `locationName:"telemetryEndpoint" type:"string"`
}

// String returns the string representation
func (s DiscoverPollEndpointOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s DiscoverPollEndpointOutput) GoString() string {
	return s.String()
}

// SetEndpoint sets the Endpoint field's value.
func (s *DiscoverPollEndpointOutput) SetEndpoint(v string) *DiscoverPollEndpointOutput {
	s.Endpoint = &v
	return s
}

// SetTelemetryEndpoint sets the TelemetryEndpoint field's value.
func (s *DiscoverPollEndpointOutput) SetTelemetryEndpoint(v string) *DiscoverPollEndpointOutput {
	s.TelemetryEndpoint = &v
	return s
}

// A failed resource.
type Failure struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the failed resource.
	Arn *string `locationName:"arn" type:"string"`

	// The reason for the failure.
	Reason *string `locationName:"reason" type:"string"`
}

// String returns the string representation
func (s Failure) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s Failure) GoString() string {
	return s.String()
}

// SetArn sets the Arn field's value.
func (s *Failure) SetArn(v string) *Failure {
	s.Arn = &v
	return s
}

// SetReason sets the Reason field's value.
func (s *Failure) SetReason(v string) *Failure {
	s.Reason = &v
	return s
}

// Hostnames and IP address entries that are added to the /etc/hosts file of
// a container via the extraHosts parameter of its ContainerDefinition.
type HostEntry struct {
	_ struct{} `type:"structure"`

	// The hostname to use in the /etc/hosts entry.
	//
	// Hostname is a required field
	Hostname *string `locationName:"hostname" type:"string" required:"true"`

	// The IP address to use in the /etc/hosts entry.
	//
	// IpAddress is a required field
	IpAddress *string `locationName:"ipAddress" type:"string" required:"true"`
}

// String returns the string representation
func (s HostEntry) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s HostEntry) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *HostEntry) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "HostEntry"}
	if s.Hostname == nil {
		invalidParams.Add(request.NewErrParamRequired("Hostname"))
	}
	if s.IpAddress == nil {
		invalidParams.Add(request.NewErrParamRequired("IpAddress"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetHostname sets the Hostname field's value.
func (s *HostEntry) SetHostname(v string) *HostEntry {
	s.Hostname = &v
	return s
}

// SetIpAddress sets the IpAddress field's value.
func (s *HostEntry) SetIpAddress(v string) *HostEntry {
	s.IpAddress = &v
	return s
}

// Details on a container instance host volume.
type HostVolumeProperties struct {
	_ struct{} `type:"structure"`

	// The path on the host container instance that is presented to the container.
	// If this parameter is empty, then the Docker daemon has assigned a host path
	// for you.
	SourcePath *string `locationName:"sourcePath" type:"string"`
}

// String returns the string representation
func (s HostVolumeProperties) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s HostVolumeProperties) GoString() string {
	return s.String()
}

// SetSourcePath sets the SourcePath field's value.
func (s *HostVolumeProperties) SetSourcePath(v string) *HostVolumeProperties {
	s.SourcePath = &v
	return s
}

// A key and value pair object.
type KeyValuePair struct {
	_ struct{} `type:"structure"`

	// The name of the key value pair. For environment variables, this is the name
	// of the environment variable.
	Name *string `locationName:"name" type:"string"`

	// The value of the key value pair. For environment variables, this is the value
	// of the environment variable.
	Value *string `locationName:"value" type:"string"`
}

// String returns the string representation
func (s KeyValuePair) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s KeyValuePair) GoString() string {
	return s.String()
}

// SetName sets the Name field's value.
func (s *KeyValuePair) SetName(v string) *KeyValuePair {
	s.Name = &v
	return s
}

// SetValue sets the Value field's value.
func (s *KeyValuePair) SetValue(v string) *KeyValuePair {
	s.Value = &v
	return s
}

type ListAttributesInput struct {
	_ struct{} `type:"structure"`

	AttributeName *string `locationName:"attributeName" type:"string"`

	AttributeValue *string `locationName:"attributeValue" type:"string"`

	Cluster *string `locationName:"cluster" type:"string"`

	MaxResults *int64 `locationName:"maxResults" type:"integer"`

	NextToken *string `locationName:"nextToken" type:"string"`

	// TargetType is a required field
	TargetType *string `locationName:"targetType" type:"string" required:"true" enum:"TargetType"`
}

// String returns the string representation
func (s ListAttributesInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s ListAttributesInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListAttributesInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "ListAttributesInput"}
	if s.TargetType == nil {
		invalidParams.Add(request.NewErrParamRequired("TargetType"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetAttributeName sets the AttributeName field's value.
func (s *ListAttributesInput) SetAttributeName(v string) *ListAttributesInput {
	s.AttributeName = &v
	return s
}

// SetAttributeValue sets the AttributeValue field's value.
func (s *ListAttributesInput) SetAttributeValue(v string) *ListAttributesInput {
	s.AttributeValue = &v
	return s
}

// SetCluster sets the Cluster field's value.
func (s *ListAttributesInput) SetCluster(v string) *ListAttributesInput {
	s.Cluster = &v
	return s
}

// SetMaxResults sets the MaxResults field's value.
func (s *ListAttributesInput) SetMaxResults(v int64) *ListAttributesInput {
	s.MaxResults = &v
	return s
}

// SetNextToken sets the NextToken field's value.
func (s *ListAttributesInput) SetNextToken(v string) *ListAttributesInput {
	s.NextToken = &v
	return s
}

// SetTargetType sets the TargetType field's value.
func (s *ListAttributesInput) SetTargetType(v string) *ListAttributesInput {
	s.TargetType = &v
	return s
}

type ListAttributesOutput struct {
	_ struct{} `type:"structure"`

	Attributes []*Attribute `locationName:"attributes" type:"list"`

	NextToken *string `locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s ListAttributesOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s ListAttributesOutput) GoString() string {
	return s.String()
}

// SetAttributes sets the Attributes field's value.
func (s *ListAttributesOutput) SetAttributes(v []*Attribute) *ListAttributesOutput {
	s.Attributes = v
	return s
}

// SetNextToken sets the NextToken field's value.
func (s *ListAttributesOutput) SetNextToken(v string) *ListAttributesOutput {
	s.NextToken = &v
	return s
}

type ListClustersInput struct {
	_ struct{} `type:"structure"`

	// The maximum number of cluster results returned by ListClusters in paginated
	// output. When this parameter is used, ListClusters only returns maxResults
	// results in a single page along with a nextToken response element. The remaining
	// results of the initial request can be seen by sending another ListClusters
	// request with the returned nextToken value. This value can be between 1 and
	// 100. If this parameter is not used, then ListClusters returns up to 100 results
	// and a nextToken value if applicable.
	MaxResults *int64 `locationName:"maxResults" type:"integer"`

	// The nextToken value returned from a previous paginated ListClusters request
	// where maxResults was used and the results exceeded the value of that parameter.
	// Pagination continues from the end of the previous results that returned the
	// nextToken value. This value is null when there are no more results to return.
	NextToken *string `locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s ListClustersInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s ListClustersInput) GoString() string {
	return s.String()
}

// SetMaxResults sets the MaxResults field's value.
func (s *ListClustersInput) SetMaxResults(v int64) *ListClustersInput {
	s.MaxResults = &v
	return s
}

// SetNextToken sets the NextToken field's value.
func (s *ListClustersInput) SetNextToken(v string) *ListClustersInput {
	s.NextToken = &v
	return s
}

type ListClustersOutput struct {
	_ struct{} `type:"structure"`

	// The list of full Amazon Resource Name (ARN) entries for each cluster associated
	// with your account.
	ClusterArns []*string `locationName:"clusterArns" type:"list"`

	// The nextToken value to include in a future ListClusters request. When the
	// results of a ListClusters request exceed maxResults, this value can be used
	// to retrieve the next page of results. This value is null when there are no
	// more results to return.
	NextToken *string `locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s ListClustersOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s ListClustersOutput) GoString() string {
	return s.String()
}

// SetClusterArns sets the ClusterArns field's value.
func (s *ListClustersOutput) SetClusterArns(v []*string) *ListClustersOutput {
	s.ClusterArns = v
	return s
}

// SetNextToken sets the NextToken field's value.
func (s *ListClustersOutput) SetNextToken(v string) *ListClustersOutput {
	s.NextToken = &v
	return s
}

type ListContainerInstancesInput struct {
	_ struct{} `type:"structure"`

	// The short name or full Amazon Resource Name (ARN) of the cluster that hosts
	// the container instances you want to list. If you do not specify a cluster,
	// the default cluster is assumed..
	Cluster *string `locationName:"cluster" type:"string"`

	Filter *string `locationName:"filter" type:"string"`

	// The maximum number of container instance results returned by ListContainerInstances
	// in paginated output. When this parameter is used, ListContainerInstances
	// only returns maxResults results in a single page along with a nextToken response
	// element. The remaining results of the initial request can be seen by sending
	// another ListContainerInstances request with the returned nextToken value.
	// This value can be between 1 and 100. If this parameter is not used, then
	// ListContainerInstances returns up to 100 results and a nextToken value if
	// applicable.
	MaxResults *int64 `locationName:"maxResults" type:"integer"`

	// The nextToken value returned from a previous paginated ListContainerInstances
	// request where maxResults was used and the results exceeded the value of that
	// parameter. Pagination continues from the end of the previous results that
	// returned the nextToken value. This value is null when there are no more results
	// to return.
	NextToken *string `locationName:"nextToken" type:"string"`

	Status *string `locationName:"status" type:"string" enum:"ContainerInstanceStatus"`
}

// String returns the string representation
func (s ListContainerInstancesInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s ListContainerInstancesInput) GoString() string {
	return s.String()
}

// SetCluster sets the Cluster field's value.
func (s *ListContainerInstancesInput) SetCluster(v string) *ListContainerInstancesInput {
	s.Cluster = &v
	return s
}

// SetFilter sets the Filter field's value.
func (s *ListContainerInstancesInput) SetFilter(v string) *ListContainerInstancesInput {
	s.Filter = &v
	return s
}

// SetMaxResults sets the MaxResults field's value.
func (s *ListContainerInstancesInput) SetMaxResults(v int64) *ListContainerInstancesInput {
	s.MaxResults = &v
	return s
}

// SetNextToken sets the NextToken field's value.
func (s *ListContainerInstancesInput) SetNextToken(v string) *ListContainerInstancesInput {
	s.NextToken = &v
	return s
}

// SetStatus sets the Status field's value.
func (s *ListContainerInstancesInput) SetStatus(v string) *ListContainerInstancesInput {
	s.Status = &v
	return s
}

type ListContainerInstancesOutput struct {
	_ struct{} `type:"structure"`

	// The list of container instance full Amazon Resource Name (ARN) entries for
	// each container instance associated with the specified cluster.
	ContainerInstanceArns []*string `locationName:"containerInstanceArns" type:"list"`

	// The nextToken value to include in a future ListContainerInstances request.
	// When the results of a ListContainerInstances request exceed maxResults, this
	// value can be used to retrieve the next page of results. This value is null
	// when there are no more results to return.
	NextToken *string `locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s ListContainerInstancesOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s ListContainerInstancesOutput) GoString() string {
	return s.String()
}

// SetContainerInstanceArns sets the ContainerInstanceArns field's value.
func (s *ListContainerInstancesOutput) SetContainerInstanceArns(v []*string) *ListContainerInstancesOutput {
	s.ContainerInstanceArns = v
	return s
}

// SetNextToken sets the NextToken field's value.
func (s *ListContainerInstancesOutput) SetNextToken(v string) *ListContainerInstancesOutput {
	s.NextToken = &v
	return s
}

type ListServicesInput struct {
	_ struct{} `type:"structure"`

	// The short name or full Amazon Resource Name (ARN) of the cluster that hosts
	// the services you want to list. If you do not specify a cluster, the default
	// cluster is assumed..
	Cluster *string `locationName:"cluster" type:"string"`

	// The maximum number of container instance results returned by ListServices
	// in paginated output. When this parameter is used, ListServices only returns
	// maxResults results in a single page along with a nextToken response element.
	// The remaining results of the initial request can be seen by sending another
	// ListServices request with the returned nextToken value. This value can be
	// between 1 and 100. If this parameter is not used, then ListServices returns
	// up to 100 results and a nextToken value if applicable.
	MaxResults *int64 `locationName:"maxResults" type:"integer"`

	// The nextToken value returned from a previous paginated ListServices request
	// where maxResults was used and the results exceeded the value of that parameter.
	// Pagination continues from the end of the previous results that returned the
	// nextToken value. This value is null when there are no more results to return.
	NextToken *string `locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s ListServicesInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s ListServicesInput) GoString() string {
	return s.String()
}

// SetCluster sets the Cluster field's value.
func (s *ListServicesInput) SetCluster(v string) *ListServicesInput {
	s.Cluster = &v
	return s
}

// SetMaxResults sets the MaxResults field's value.
func (s *ListServicesInput) SetMaxResults(v int64) *ListServicesInput {
	s.MaxResults = &v
	return s
}

// SetNextToken sets the NextToken field's value.
func (s *ListServicesInput) SetNextToken(v string) *ListServicesInput {
	s.NextToken = &v
	return s
}

type ListServicesOutput struct {
	_ struct{} `type:"structure"`

	// The nextToken value to include in a future ListServices request. When the
	// results of a ListServices request exceed maxResults, this value can be used
	// to retrieve the next page of results. This value is null when there are no
	// more results to return.
	NextToken *string `locationName:"nextToken" type:"string"`

	// The list of full Amazon Resource Name (ARN) entries for each service associated
	// with the specified cluster.
	ServiceArns []*string `locationName:"serviceArns" type:"list"`
}

// String returns the string representation
func (s ListServicesOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s ListServicesOutput) GoString() string {
	return s.String()
}

// SetNextToken sets the NextToken field's value.
func (s *ListServicesOutput) SetNextToken(v string) *ListServicesOutput {
	s.NextToken = &v
	return s
}

// SetServiceArns sets the ServiceArns field's value.
func (s *ListServicesOutput) SetServiceArns(v []*string) *ListServicesOutput {
	s.ServiceArns = v
	return s
}

type ListTaskDefinitionFamiliesInput struct {
	_ struct{} `type:"structure"`

	// The familyPrefix is a string that is used to filter the results of ListTaskDefinitionFamilies.
	// If you specify a familyPrefix, only task definition family names that begin
	// with the familyPrefix string are returned.
	FamilyPrefix *string `locationName:"familyPrefix" type:"string"`

	// The maximum number of task definition family results returned by ListTaskDefinitionFamilies
	// in paginated output. When this parameter is used, ListTaskDefinitions only
	// returns maxResults results in a single page along with a nextToken response
	// element. The remaining results of the initial request can be seen by sending
	// another ListTaskDefinitionFamilies request with the returned nextToken value.
	// This value can be between 1 and 100. If this parameter is not used, then
	// ListTaskDefinitionFamilies returns up to 100 results and a nextToken value
	// if applicable.
	MaxResults *int64 `locationName:"maxResults" type:"integer"`

	// The nextToken value returned from a previous paginated ListTaskDefinitionFamilies
	// request where maxResults was used and the results exceeded the value of that
	// parameter. Pagination continues from the end of the previous results that
	// returned the nextToken value. This value is null when there are no more results
	// to return.
	NextToken *string `locationName:"nextToken" type:"string"`

	Status *string `locationName:"status" type:"string" enum:"TaskDefinitionFamilyStatus"`
}

// String returns the string representation
func (s ListTaskDefinitionFamiliesInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s ListTaskDefinitionFamiliesInput) GoString() string {
	return s.String()
}

// SetFamilyPrefix sets the FamilyPrefix field's value.
func (s *ListTaskDefinitionFamiliesInput) SetFamilyPrefix(v string) *ListTaskDefinitionFamiliesInput {
	s.FamilyPrefix = &v
	return s
}

// SetMaxResults sets the MaxResults field's value.
func (s *ListTaskDefinitionFamiliesInput) SetMaxResults(v int64) *ListTaskDefinitionFamiliesInput {
	s.MaxResults = &v
	return s
}

// SetNextToken sets the NextToken field's value.
func (s *ListTaskDefinitionFamiliesInput) SetNextToken(v string) *ListTaskDefinitionFamiliesInput {
	s.NextToken = &v
	return s
}

// SetStatus sets the Status field's value.
func (s *ListTaskDefinitionFamiliesInput) SetStatus(v string) *ListTaskDefinitionFamiliesInput {
	s.Status = &v
	return s
}

type ListTaskDefinitionFamiliesOutput struct {
	_ struct{} `type:"structure"`

	// The list of task definition family names that match the ListTaskDefinitionFamilies
	// request.
	Families []*string `locationName:"families" type:"list"`

	// The nextToken value to include in a future ListTaskDefinitionFamilies request.
	// When the results of a ListTaskDefinitionFamilies request exceed maxResults,
	// this value can be used to retrieve the next page of results. This value is
	// null when there are no more results to return.
	NextToken *string `locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s ListTaskDefinitionFamiliesOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s ListTaskDefinitionFamiliesOutput) GoString() string {
	return s.String()
}

// SetFamilies sets the Families field's value.
func (s *ListTaskDefinitionFamiliesOutput) SetFamilies(v []*string) *ListTaskDefinitionFamiliesOutput {
	s.Families = v
	return s
}

// SetNextToken sets the NextToken field's value.
func (s *ListTaskDefinitionFamiliesOutput) SetNextToken(v string) *ListTaskDefinitionFamiliesOutput {
	s.NextToken = &v
	return s
}

type ListTaskDefinitionsInput struct {
	_ struct{} `type:"structure"`

	// The full family name that you want to filter the ListTaskDefinitions results
	// with. Specifying a familyPrefix will limit the listed task definitions to
	// task definition revisions that belong to that family.
	FamilyPrefix *string `locationName:"familyPrefix" type:"string"`

	// The maximum number of task definition results returned by ListTaskDefinitions
	// in paginated output. When this parameter is used, ListTaskDefinitions only
	// returns maxResults results in a single page along with a nextToken response
	// element. The remaining results of the initial request can be seen by sending
	// another ListTaskDefinitions request with the returned nextToken value. This
	// value can be between 1 and 100. If this parameter is not used, then ListTaskDefinitions
	// returns up to 100 results and a nextToken value if applicable.
	MaxResults *int64 `locationName:"maxResults" type:"integer"`

	// The nextToken value returned from a previous paginated ListTaskDefinitions
	// request where maxResults was used and the results exceeded the value of that
	// parameter. Pagination continues from the end of the previous results that
	// returned the nextToken value. This value is null when there are no more results
	// to return.
	NextToken *string `locationName:"nextToken" type:"string"`

	// The order in which to sort the results. Valid values are ASC and DESC. By
	// default (ASC), task definitions are listed lexicographically by family name
	// and in ascending numerical order by revision so that the newest task definitions
	// in a family are listed last. Setting this parameter to DESC reverses the
	// sort order on family name and revision so that the newest task definitions
	// in a family are listed first.
	Sort *string `locationName:"sort" type:"string" enum:"SortOrder"`

	// The task definition status that you want to filter the ListTaskDefinitions
	// results with. By default, only ACTIVE task definitions are listed. By setting
	// this parameter to INACTIVE, you can view task definitions that are INACTIVE
	// as long as an active task or service still references them. If you paginate
	// the resulting output, be sure to keep the status value constant in each subsequent
	// request.
	Status *string `locationName:"status" type:"string" enum:"TaskDefinitionStatus"`
}

// String returns the string representation
func (s ListTaskDefinitionsInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s ListTaskDefinitionsInput) GoString() string {
	return s.String()
}

// SetFamilyPrefix sets the FamilyPrefix field's value.
func (s *ListTaskDefinitionsInput) SetFamilyPrefix(v string) *ListTaskDefinitionsInput {
	s.FamilyPrefix = &v
	return s
}

// SetMaxResults sets the MaxResults field's value.
func (s *ListTaskDefinitionsInput) SetMaxResults(v int64) *ListTaskDefinitionsInput {
	s.MaxResults = &v
	return s
}

// SetNextToken sets the NextToken field's value.
func (s *ListTaskDefinitionsInput) SetNextToken(v string) *ListTaskDefinitionsInput {
	s.NextToken = &v
	return s
}

// SetSort sets the Sort field's value.
func (s *ListTaskDefinitionsInput) SetSort(v string) *ListTaskDefinitionsInput {
	s.Sort = &v
	return s
}

// SetStatus sets the Status field's value.
func (s *ListTaskDefinitionsInput) SetStatus(v string) *ListTaskDefinitionsInput {
	s.Status = &v
	return s
}

type ListTaskDefinitionsOutput struct {
	_ struct{} `type:"structure"`

	// The nextToken value to include in a future ListTaskDefinitions request. When
	// the results of a ListTaskDefinitions request exceed maxResults, this value
	// can be used to retrieve the next page of results. This value is null when
	// there are no more results to return.
	NextToken *string `locationName:"nextToken" type:"string"`

	// The list of task definition Amazon Resource Name (ARN) entries for the ListTaskDefinitions
	// request.
	TaskDefinitionArns []*string `locationName:"taskDefinitionArns" type:"list"`
}

// String returns the string representation
func (s ListTaskDefinitionsOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s ListTaskDefinitionsOutput) GoString() string {
	return s.String()
}

// SetNextToken sets the NextToken field's value.
func (s *ListTaskDefinitionsOutput) SetNextToken(v string) *ListTaskDefinitionsOutput {
	s.NextToken = &v
	return s
}

// SetTaskDefinitionArns sets the TaskDefinitionArns field's value.
func (s *ListTaskDefinitionsOutput) SetTaskDefinitionArns(v []*string) *ListTaskDefinitionsOutput {
	s.TaskDefinitionArns = v
	return s
}

type ListTasksInput struct {
	_ struct{} `type:"structure"`

	// The short name or full Amazon Resource Name (ARN) of the cluster that hosts
	// the tasks you want to list. If you do not specify a cluster, the default
	// cluster is assumed..
	Cluster *string `locationName:"cluster" type:"string"`

	// The container instance ID or full Amazon Resource Name (ARN) of the container
	// instance that you want to filter the ListTasks results with. Specifying a
	// containerInstance will limit the results to tasks that belong to that container
	// instance.
	ContainerInstance *string `locationName:"containerInstance" type:"string"`

	// The task status that you want to filter the ListTasks results with. Specifying
	// a desiredStatus of STOPPED will limit the results to tasks that are in the
	// STOPPED status, which can be useful for debugging tasks that are not starting
	// properly or have died or finished. The default status filter is RUNNING.
	DesiredStatus *string `locationName:"desiredStatus" type:"string" enum:"DesiredStatus"`

	// The name of the family that you want to filter the ListTasks results with.
	// Specifying a family will limit the results to tasks that belong to that family.
	Family *string `locationName:"family" type:"string"`

	// The maximum number of task results returned by ListTasks in paginated output.
	// When this parameter is used, ListTasks only returns maxResults results in
	// a single page along with a nextToken response element. The remaining results
	// of the initial request can be seen by sending another ListTasks request with
	// the returned nextToken value. This value can be between 1 and 100. If this
	// parameter is not used, then ListTasks returns up to 100 results and a nextToken
	// value if applicable.
	MaxResults *int64 `locationName:"maxResults" type:"integer"`

	// The nextToken value returned from a previous paginated ListTasks request
	// where maxResults was used and the results exceeded the value of that parameter.
	// Pagination continues from the end of the previous results that returned the
	// nextToken value. This value is null when there are no more results to return.
	NextToken *string `locationName:"nextToken" type:"string"`

	// The name of the service that you want to filter the ListTasks results with.
	// Specifying a serviceName will limit the results to tasks that belong to that
	// service.
	ServiceName *string `locationName:"serviceName" type:"string"`

	// The startedBy value that you want to filter the task results with. Specifying
	// a startedBy value will limit the results to tasks that were started with
	// that value.
	StartedBy *string `locationName:"startedBy" type:"string"`
}

// String returns the string representation
func (s ListTasksInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s ListTasksInput) GoString() string {
	return s.String()
}

// SetCluster sets the Cluster field's value.
func (s *ListTasksInput) SetCluster(v string) *ListTasksInput {
	s.Cluster = &v
	return s
}

// SetContainerInstance sets the ContainerInstance field's value.
func (s *ListTasksInput) SetContainerInstance(v string) *ListTasksInput {
	s.ContainerInstance = &v
	return s
}

// SetDesiredStatus sets the DesiredStatus field's value.
func (s *ListTasksInput) SetDesiredStatus(v string) *ListTasksInput {
	s.DesiredStatus = &v
	return s
}

// SetFamily sets the Family field's value.
func (s *ListTasksInput) SetFamily(v string) *ListTasksInput {
	s.Family = &v
	return s
}

// SetMaxResults sets the MaxResults field's value.
func (s *ListTasksInput) SetMaxResults(v int64) *ListTasksInput {
	s.MaxResults = &v
	return s
}

// SetNextToken sets the NextToken field's value.
func (s *ListTasksInput) SetNextToken(v string) *ListTasksInput {
	s.NextToken = &v
	return s
}

// SetServiceName sets the ServiceName field's value.
func (s *ListTasksInput) SetServiceName(v string) *ListTasksInput {
	s.ServiceName = &v
	return s
}

// SetStartedBy sets the StartedBy field's value.
func (s *ListTasksInput) SetStartedBy(v string) *ListTasksInput {
	s.StartedBy = &v
	return s
}

type ListTasksOutput struct {
	_ struct{} `type:"structure"`

	// The nextToken value to include in a future ListTasks request. When the results
	// of a ListTasks request exceed maxResults, this value can be used to retrieve
	// the next page of results. This value is null when there are no more results
	// to return.
	NextToken *string `locationName:"nextToken" type:"string"`

	// The list of task Amazon Resource Name (ARN) entries for the ListTasks request.
	TaskArns []*string `locationName:"taskArns" type:"list"`
}

// String returns the string representation
func (s ListTasksOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s ListTasksOutput) GoString() string {
	return s.String()
}

// SetNextToken sets the NextToken field's value.
func (s *ListTasksOutput) SetNextToken(v string) *ListTasksOutput {
	s.NextToken = &v
	return s
}

// SetTaskArns sets the TaskArns field's value.
func (s *ListTasksOutput) SetTaskArns(v []*string) *ListTasksOutput {
	s.TaskArns = v
	return s
}

// Details on a load balancer that is used with a service.
type LoadBalancer struct {
	_ struct{} `type:"structure"`

	// The name of the container to associate with the load balancer.
	ContainerName *string `locationName:"containerName" type:"string"`

	// The port on the container to associate with the load balancer. This port
	// must correspond to a containerPort in the service's task definition. Your
	// container instances must allow ingress traffic on the hostPort of the port
	// mapping.
	ContainerPort *int64 `locationName:"containerPort" type:"integer"`

	// The name of the load balancer.
	LoadBalancerName *string `locationName:"loadBalancerName" type:"string"`

	TargetGroupArn *string `locationName:"targetGroupArn" type:"string"`
}

// String returns the string representation
func (s LoadBalancer) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s LoadBalancer) GoString() string {
	return s.String()
}

// SetContainerName sets the ContainerName field's value.
func (s *LoadBalancer) SetContainerName(v string) *LoadBalancer {
	s.ContainerName = &v
	return s
}

// SetContainerPort sets the ContainerPort field's value.
func (s *LoadBalancer) SetContainerPort(v int64) *LoadBalancer {
	s.ContainerPort = &v
	return s
}

// SetLoadBalancerName sets the LoadBalancerName field's value.
func (s *LoadBalancer) SetLoadBalancerName(v string) *LoadBalancer {
	s.LoadBalancerName = &v
	return s
}

// SetTargetGroupArn sets the TargetGroupArn field's value.
func (s *LoadBalancer) SetTargetGroupArn(v string) *LoadBalancer {
	s.TargetGroupArn = &v
	return s
}

// Log configuration options to send to a custom log driver for the container.
type LogConfiguration struct {
	_ struct{} `type:"structure"`

	// The log driver to use for the container. This parameter requires version
	// 1.18 of the Docker Remote API or greater on your container instance. To check
	// the Docker Remote API version on your container instance, log into your container
	// instance and run the following command: sudo docker version | grep "Server
	// API version"
	//
	// LogDriver is a required field
	LogDriver *string `locationName:"logDriver" type:"string" required:"true" enum:"LogDriver"`

	// The configuration options to send to the log driver. This parameter requires
	// version 1.19 of the Docker Remote API or greater on your container instance.
	// To check the Docker Remote API version on your container instance, log into
	// your container instance and run the following command: sudo docker version
	// | grep "Server API version"
	Options map[string]*string `locationName:"options" type:"map"`
}

// String returns the string representation
func (s LogConfiguration) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s LogConfiguration) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *LogConfiguration) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "LogConfiguration"}
	if s.LogDriver == nil {
		invalidParams.Add(request.NewErrParamRequired("LogDriver"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetLogDriver sets the LogDriver field's value.
func (s *LogConfiguration) SetLogDriver(v string) *LogConfiguration {
	s.LogDriver = &v
	return s
}

// SetOptions sets the Options field's value.
func (s *LogConfiguration) SetOptions(v map[string]*string) *LogConfiguration {
	s.Options = v
	return s
}

// Details on a volume mount point that is used in a container definition.
type MountPoint struct {
	_ struct{} `type:"structure"`

	// The path on the container to mount the host volume at.
	ContainerPath *string `locationName:"containerPath" type:"string"`

	// If this value is true, the container has read-only access to the volume.
	// If this value is false, then the container can write to the volume. The default
	// value is false.
	ReadOnly *bool `locationName:"readOnly" type:"boolean"`

	// The name of the volume to mount.
	SourceVolume *string `locationName:"sourceVolume" type:"string"`
}

// String returns the string representation
func (s MountPoint) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s MountPoint) GoString() string {
	return s.String()
}

// SetContainerPath sets the ContainerPath field's value.
func (s *MountPoint) SetContainerPath(v string) *MountPoint {
	s.ContainerPath = &v
	return s
}

// SetReadOnly sets the ReadOnly field's value.
func (s *MountPoint) SetReadOnly(v bool) *MountPoint {
	s.ReadOnly = &v
	return s
}

// SetSourceVolume sets the SourceVolume field's value.
func (s *MountPoint) SetSourceVolume(v string) *MountPoint {
	s.SourceVolume = &v
	return s
}

// Details on the network bindings between a container and its host container
// instance. After a task reaches the RUNNING status, manual and automatic host
// and container port assignments are visible in the networkBindings section
// of DescribeTasks API responses.
type NetworkBinding struct {
	_ struct{} `type:"structure"`

	// The IP address that the container is bound to on the container instance.
	BindIP *string `locationName:"bindIP" type:"string"`

	// The port number on the container that is be used with the network binding.
	ContainerPort *int64 `locationName:"containerPort" type:"integer"`

	// The port number on the host that is used with the network binding.
	HostPort *int64 `locationName:"hostPort" type:"integer"`

	// The protocol used for the network binding.
	Protocol *string `locationName:"protocol" type:"string" enum:"TransportProtocol"`
}

// String returns the string representation
func (s NetworkBinding) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s NetworkBinding) GoString() string {
	return s.String()
}

// SetBindIP sets the BindIP field's value.
func (s *NetworkBinding) SetBindIP(v string) *NetworkBinding {
	s.BindIP = &v
	return s
}

// SetContainerPort sets the ContainerPort field's value.
func (s *NetworkBinding) SetContainerPort(v int64) *NetworkBinding {
	s.ContainerPort = &v
	return s
}

// SetHostPort sets the HostPort field's value.
func (s *NetworkBinding) SetHostPort(v int64) *NetworkBinding {
	s.HostPort = &v
	return s
}

// SetProtocol sets the Protocol field's value.
func (s *NetworkBinding) SetProtocol(v string) *NetworkBinding {
	s.Protocol = &v
	return s
}

type PlacementConstraint struct {
	_ struct{} `type:"structure"`

	Expression *string `locationName:"expression" type:"string"`

	Type *string `locationName:"type" type:"string" enum:"PlacementConstraintType"`
}

// String returns the string representation
func (s PlacementConstraint) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s PlacementConstraint) GoString() string {
	return s.String()
}

// SetExpression sets the Expression field's value.
func (s *PlacementConstraint) SetExpression(v string) *PlacementConstraint {
	s.Expression = &v
	return s
}

// SetType sets the Type field's value.
func (s *PlacementConstraint) SetType(v string) *PlacementConstraint {
	s.Type = &v
	return s
}

type PlacementStrategy struct {
	_ struct{} `type:"structure"`

	Field *string `locationName:"field" type:"string"`

	Type *string `locationName:"type" type:"string" enum:"PlacementStrategyType"`
}

// String returns the string representation
func (s PlacementStrategy) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s PlacementStrategy) GoString() string {
	return s.String()
}

// SetField sets the Field field's value.
func (s *PlacementStrategy) SetField(v string) *PlacementStrategy {
	s.Field = &v
	return s
}

// SetType sets the Type field's value.
func (s *PlacementStrategy) SetType(v string) *PlacementStrategy {
	s.Type = &v
	return s
}

// Port mappings allow containers to access ports on the host container instance
// to send or receive traffic. Port mappings are specified as part of the container
// definition. After a task reaches the RUNNING status, manual and automatic
// host and container port assignments are visible in the networkBindings section
// of DescribeTasks API responses.
type PortMapping struct {
	_ struct{} `type:"structure"`

	// The port number on the container that is bound to the user-specified or automatically
	// assigned host port. If you specify a container port and not a host port,
	// your container will automatically receive a host port in the ephemeral port
	// range (for more information, see hostPort).
	ContainerPort *int64 `locationName:"containerPort" type:"integer"`

	// The port number on the container instance to reserve for your container.
	// You can specify a non-reserved host port for your container port mapping,
	// or you can omit the hostPort (or set it to 0) while specifying a containerPort
	// and your container will automatically receive a port in the ephemeral port
	// range for your container instance operating system and Docker version.
	//
	// The default ephemeral port range is 49153 to 65535, and this range is used
	// for Docker versions prior to 1.6.0. For Docker version 1.6.0 and later, the
	// Docker daemon tries to read the ephemeral port range from /proc/sys/net/ipv4/ip_local_port_range;
	// if this kernel parameter is unavailable, the default ephemeral port range
	// is used. You should not attempt to specify a host port in the ephemeral port
	// range, since these are reserved for automatic assignment. In general, ports
	// below 32768 are outside of the ephemeral port range.
	//
	// The default reserved ports are 22 for SSH, the Docker ports 2375 and 2376,
	// and the Amazon ECS Container Agent port 51678. Any host port that was previously
	// specified in a running task is also reserved while the task is running (once
	// a task stops, the host port is released).The current reserved ports are displayed
	// in the remainingResources of DescribeContainerInstances output, and a container
	// instance may have up to 50 reserved ports at a time, including the default
	// reserved ports (automatically assigned ports do not count toward this limit).
	HostPort *int64 `locationName:"hostPort" type:"integer"`

	// The protocol used for the port mapping. Valid values are tcp and udp. The
	// default is tcp.
	Protocol *string `locationName:"protocol" type:"string" enum:"TransportProtocol"`
}

// String returns the string representation
func (s PortMapping) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s PortMapping) GoString() string {
	return s.String()
}

// SetContainerPort sets the ContainerPort field's value.
func (s *PortMapping) SetContainerPort(v int64) *PortMapping {
	s.ContainerPort = &v
	return s
}

// SetHostPort sets the HostPort field's value.
func (s *PortMapping) SetHostPort(v int64) *PortMapping {
	s.HostPort = &v
	return s
}

// SetProtocol sets the Protocol field's value.
func (s *PortMapping) SetProtocol(v string) *PortMapping {
	s.Protocol = &v
	return s
}

type PutAttributesInput struct {
	_ struct{} `type:"structure"`

	// Attributes is a required field
	Attributes []*Attribute `locationName:"attributes" type:"list" required:"true"`

	Cluster *string `locationName:"cluster" type:"string"`
}

// String returns the string representation
func (s PutAttributesInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s PutAttributesInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PutAttributesInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "PutAttributesInput"}
	if s.Attributes == nil {
		invalidParams.Add(request.NewErrParamRequired("Attributes"))
	}
	if s.Attributes != nil {
		for i, v := range s.Attributes {
			if v == nil {
				continue
			}
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Attributes", i), err.(request.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetAttributes sets the Attributes field's value.
func (s *PutAttributesInput) SetAttributes(v []*Attribute) *PutAttributesInput {
	s.Attributes = v
	return s
}

// SetCluster sets the Cluster field's value.
func (s *PutAttributesInput) SetCluster(v string) *PutAttributesInput {
	s.Cluster = &v
	return s
}

type PutAttributesOutput struct {
	_ struct{} `type:"structure"`

	Attributes []*Attribute `locationName:"attributes" type:"list"`
}

// String returns the string representation
func (s PutAttributesOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s PutAttributesOutput) GoString() string {
	return s.String()
}

// SetAttributes sets the Attributes field's value.
func (s *PutAttributesOutput) SetAttributes(v []*Attribute) *PutAttributesOutput {
	s.Attributes = v
	return s
}

type RegisterContainerInstanceInput struct {
	_ struct{} `type:"structure"`

	// The container instance attributes that this container instance supports.
	Attributes []*Attribute `locationName:"attributes" type:"list"`

	// The short name or full Amazon Resource Name (ARN) of the cluster that you
	// want to register your container instance with. If you do not specify a cluster,
	// the default cluster is assumed..
	Cluster *string `locationName:"cluster" type:"string"`

	// The Amazon Resource Name (ARN) of the container instance (if it was previously
	// registered).
	ContainerInstanceArn *string `locationName:"containerInstanceArn" type:"string"`

	// The instance identity document for the Amazon EC2 instance to register. This
	// document can be found by running the following command from the instance:
	// curl http://169.254.169.254/latest/dynamic/instance-identity/document/
	InstanceIdentityDocument *string `locationName:"instanceIdentityDocument" type:"string"`

	// The instance identity document signature for the Amazon EC2 instance to register.
	// This signature can be found by running the following command from the instance:
	// curl http://169.254.169.254/latest/dynamic/instance-identity/signature/
	InstanceIdentityDocumentSignature *string `locationName:"instanceIdentityDocumentSignature" type:"string"`

	// The resources available on the instance.
	TotalResources []*Resource `locationName:"totalResources" type:"list"`

	// The version information for the Amazon ECS container agent and Docker daemon
	// running on the container instance.
	VersionInfo *VersionInfo `locationName:"versionInfo" type:"structure"`
}

// String returns the string representation
func (s RegisterContainerInstanceInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s RegisterContainerInstanceInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *RegisterContainerInstanceInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "RegisterContainerInstanceInput"}
	if s.Attributes != nil {
		for i, v := range s.Attributes {
			if v == nil {
				continue
			}
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Attributes", i), err.(request.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetAttributes sets the Attributes field's value.
func (s *RegisterContainerInstanceInput) SetAttributes(v []*Attribute) *RegisterContainerInstanceInput {
	s.Attributes = v
	return s
}

// SetCluster sets the Cluster field's value.
func (s *RegisterContainerInstanceInput) SetCluster(v string) *RegisterContainerInstanceInput {
	s.Cluster = &v
	return s
}

// SetContainerInstanceArn sets the ContainerInstanceArn field's value.
func (s *RegisterContainerInstanceInput) SetContainerInstanceArn(v string) *RegisterContainerInstanceInput {
	s.ContainerInstanceArn = &v
	return s
}

// SetInstanceIdentityDocument sets the InstanceIdentityDocument field's value.
func (s *RegisterContainerInstanceInput) SetInstanceIdentityDocument(v string) *RegisterContainerInstanceInput {
	s.InstanceIdentityDocument = &v
	return s
}

// SetInstanceIdentityDocumentSignature sets the InstanceIdentityDocumentSignature field's value.
func (s *RegisterContainerInstanceInput) SetInstanceIdentityDocumentSignature(v string) *RegisterContainerInstanceInput {
	s.InstanceIdentityDocumentSignature = &v
	return s
}

// SetTotalResources sets the TotalResources field's value.
func (s *RegisterContainerInstanceInput) SetTotalResources(v []*Resource) *RegisterContainerInstanceInput {
	s.TotalResources = v
	return s
}

// SetVersionInfo sets the VersionInfo field's value.
func (s *RegisterContainerInstanceInput) SetVersionInfo(v *VersionInfo) *RegisterContainerInstanceInput {
	s.VersionInfo = v
	return s
}

type RegisterContainerInstanceOutput struct {
	_ struct{} `type:"structure"`

	// An Amazon EC2 instance that is running the Amazon ECS agent and has been
	// registered with a cluster.
	ContainerInstance *ContainerInstance `locationName:"containerInstance" type:"structure"`
}

// String returns the string representation
func (s RegisterContainerInstanceOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s RegisterContainerInstanceOutput) GoString() string {
	return s.String()
}

// SetContainerInstance sets the ContainerInstance field's value.
func (s *RegisterContainerInstanceOutput) SetContainerInstance(v *ContainerInstance) *RegisterContainerInstanceOutput {
	s.ContainerInstance = v
	return s
}

type RegisterTaskDefinitionInput struct {
	_ struct{} `type:"structure"`

	// A list of container definitions in JSON format that describe the different
	// containers that make up your task.
	//
	// ContainerDefinitions is a required field
	ContainerDefinitions []*ContainerDefinition `locationName:"containerDefinitions" type:"list" required:"true"`

	// You must specify a family for a task definition, which allows you to track
	// multiple versions of the same task definition. You can think of the family
	// as a name for your task definition. Up to 255 letters (uppercase and lowercase),
	// numbers, hyphens, and underscores are allowed.
	//
	// Family is a required field
	Family *string `locationName:"family" type:"string" required:"true"`

	NetworkMode *string `locationName:"networkMode" type:"string" enum:"NetworkMode"`

	PlacementConstraints []*TaskDefinitionPlacementConstraint `locationName:"placementConstraints" type:"list"`

	TaskRoleArn *string `locationName:"taskRoleArn" type:"string"`

	// A list of volume definitions in JSON format that containers in your task
	// may use.
	Volumes []*Volume `locationName:"volumes" type:"list"`
}

// String returns the string representation
func (s RegisterTaskDefinitionInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s RegisterTaskDefinitionInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *RegisterTaskDefinitionInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "RegisterTaskDefinitionInput"}
	if s.ContainerDefinitions == nil {
		invalidParams.Add(request.NewErrParamRequired("ContainerDefinitions"))
	}
	if s.Family == nil {
		invalidParams.Add(request.NewErrParamRequired("Family"))
	}
	if s.ContainerDefinitions != nil {
		for i, v := range s.ContainerDefinitions {
			if v == nil {
				continue
			}
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "ContainerDefinitions", i), err.(request.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetContainerDefinitions sets the ContainerDefinitions field's value.
func (s *RegisterTaskDefinitionInput) SetContainerDefinitions(v []*ContainerDefinition) *RegisterTaskDefinitionInput {
	s.ContainerDefinitions = v
	return s
}

// SetFamily sets the Family field's value.
func (s *RegisterTaskDefinitionInput) SetFamily(v string) *RegisterTaskDefinitionInput {
	s.Family = &v
	return s
}

// SetNetworkMode sets the NetworkMode field's value.
func (s *RegisterTaskDefinitionInput) SetNetworkMode(v string) *RegisterTaskDefinitionInput {
	s.NetworkMode = &v
	return s
}

// SetPlacementConstraints sets the PlacementConstraints field's value.
func (s *RegisterTaskDefinitionInput) SetPlacementConstraints(v []*TaskDefinitionPlacementConstraint) *RegisterTaskDefinitionInput {
	s.PlacementConstraints = v
	return s
}

// SetTaskRoleArn sets the TaskRoleArn field's value.
func (s *RegisterTaskDefinitionInput) SetTaskRoleArn(v string) *RegisterTaskDefinitionInput {
	s.TaskRoleArn = &v
	return s
}

// SetVolumes sets the Volumes field's value.
func (s *RegisterTaskDefinitionInput) SetVolumes(v []*Volume) *RegisterTaskDefinitionInput {
	s.Volumes = v
	return s
}

type RegisterTaskDefinitionOutput struct {
	_ struct{} `type:"structure"`

	// The full description of the registered task definition.
	TaskDefinition *TaskDefinition `locationName:"taskDefinition" type:"structure"`
}

// String returns the string representation
func (s RegisterTaskDefinitionOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s RegisterTaskDefinitionOutput) GoString() string {
	return s.String()
}

// SetTaskDefinition sets the TaskDefinition field's value.
func (s *RegisterTaskDefinitionOutput) SetTaskDefinition(v *TaskDefinition) *RegisterTaskDefinitionOutput {
	s.TaskDefinition = v
	return s
}

// Describes the resources available for a container instance.
type Resource struct {
	_ struct{} `type:"structure"`

	// When the doubleValue type is set, the value of the resource must be a double
	// precision floating-point type.
	DoubleValue *float64 `locationName:"doubleValue" type:"double"`

	// When the integerValue type is set, the value of the resource must be an integer.
	IntegerValue *int64 `locationName:"integerValue" type:"integer"`

	// When the longValue type is set, the value of the resource must be an extended
	// precision floating-point type.
	LongValue *int64 `locationName:"longValue" type:"long"`

	// The name of the resource, such as CPU, MEMORY, PORTS, or a user-defined resource.
	Name *string `locationName:"name" type:"string"`

	// When the stringSetValue type is set, the value of the resource must be a
	// string type.
	StringSetValue []*string `locationName:"stringSetValue" type:"list"`

	// The type of the resource, such as INTEGER, DOUBLE, LONG, or STRINGSET.
	Type *string `locationName:"type" type:"string"`
}

// String returns the string representation
func (s Resource) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s Resource) GoString() string {
	return s.String()
}

// SetDoubleValue sets the DoubleValue field's value.
func (s *Resource) SetDoubleValue(v float64) *Resource {
	s.DoubleValue = &v
	return s
}

// SetIntegerValue sets the IntegerValue field's value.
func (s *Resource) SetIntegerValue(v int64) *Resource {
	s.IntegerValue = &v
	return s
}

// SetLongValue sets the LongValue field's value.
func (s *Resource) SetLongValue(v int64) *Resource {
	s.LongValue = &v
	return s
}

// SetName sets the Name field's value.
func (s *Resource) SetName(v string) *Resource {
	s.Name = &v
	return s
}

// SetStringSetValue sets the StringSetValue field's value.
func (s *Resource) SetStringSetValue(v []*string) *Resource {
	s.StringSetValue = v
	return s
}

// SetType sets the Type field's value.
func (s *Resource) SetType(v string) *Resource {
	s.Type = &v
	return s
}

type RunTaskInput struct {
	_ struct{} `type:"structure"`

	// The short name or full Amazon Resource Name (ARN) of the cluster that you
	// want to run your task on. If you do not specify a cluster, the default cluster
	// is assumed..
	Cluster *string `locationName:"cluster" type:"string"`

	// The number of instantiations of the specified task that you would like to
	// place on your cluster.
	//
	// The count parameter is limited to 10 tasks per call.
	Count *int64 `locationName:"count" type:"integer"`

	Group *string `locationName:"group" type:"string"`

	LaunchType *string `locationName:"launchType" type:"string" enum:"LaunchType"`

	// A list of container overrides in JSON format that specify the name of a container
	// in the specified task definition and the overrides it should receive. You
	// can override the default command for a container (that is specified in the
	// task definition or Docker image) with a command override. You can also override
	// existing environment variables (that are specified in the task definition
	// or Docker image) on a container or add new environment variables to it with
	// an environment override.
	//
	// A total of 8192 characters are allowed for overrides. This limit includes
	// the JSON formatting characters of the override structure.
	Overrides *TaskOverride `locationName:"overrides" type:"structure"`

	PlacementConstraints []*PlacementConstraint `locationName:"placementConstraints" type:"list"`

	PlacementStrategy []*PlacementStrategy `locationName:"placementStrategy" type:"list"`

	RoleArn *string `locationName:"roleArn" type:"string"`

	// An optional tag specified when a task is started. For example if you automatically
	// trigger a task to run a batch process job, you could apply a unique identifier
	// for that job to your task with the startedBy parameter. You can then identify
	// which tasks belong to that job by filtering the results of a ListTasks call
	// with the startedBy value.
	//
	// If a task is started by an Amazon ECS service, then the startedBy parameter
	// contains the deployment ID of the service that starts it.
	StartedBy *string `locationName:"startedBy" type:"string"`

	// The family and revision (family:revision) or full Amazon Resource Name (ARN)
	// of the task definition that you want to run. If a revision is not specified,
	// the latest ACTIVE revision is used.
	//
	// TaskDefinition is a required field
	TaskDefinition *string `locationName:"taskDefinition" type:"string" required:"true"`
}

// String returns the string representation
func (s RunTaskInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s RunTaskInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *RunTaskInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "RunTaskInput"}
	if s.TaskDefinition == nil {
		invalidParams.Add(request.NewErrParamRequired("TaskDefinition"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetCluster sets the Cluster field's value.
func (s *RunTaskInput) SetCluster(v string) *RunTaskInput {
	s.Cluster = &v
	return s
}

// SetCount sets the Count field's value.
func (s *RunTaskInput) SetCount(v int64) *RunTaskInput {
	s.Count = &v
	return s
}

// SetGroup sets the Group field's value.
func (s *RunTaskInput) SetGroup(v string) *RunTaskInput {
	s.Group = &v
	return s
}

// SetLaunchType sets the LaunchType field's value.
func (s *RunTaskInput) SetLaunchType(v string) *RunTaskInput {
	s.LaunchType = &v
	return s
}

// SetOverrides sets the Overrides field's value.
func (s *RunTaskInput) SetOverrides(v *TaskOverride) *RunTaskInput {
	s.Overrides = v
	return s
}

// SetPlacementConstraints sets the PlacementConstraints field's value.
func (s *RunTaskInput) SetPlacementConstraints(v []*PlacementConstraint) *RunTaskInput {
	s.PlacementConstraints = v
	return s
}

// SetPlacementStrategy sets the PlacementStrategy field's value.
func (s *RunTaskInput) SetPlacementStrategy(v []*PlacementStrategy) *RunTaskInput {
	s.PlacementStrategy = v
	return s
}

// SetRoleArn sets the RoleArn field's value.
func (s *RunTaskInput) SetRoleArn(v string) *RunTaskInput {
	s.RoleArn = &v
	return s
}

// SetStartedBy sets the StartedBy field's value.
func (s *RunTaskInput) SetStartedBy(v string) *RunTaskInput {
	s.StartedBy = &v
	return s
}

// SetTaskDefinition sets the TaskDefinition field's value.
func (s *RunTaskInput) SetTaskDefinition(v string) *RunTaskInput {
	s.TaskDefinition = &v
	return s
}

type RunTaskOutput struct {
	_ struct{} `type:"structure"`

	// Any failures associated with the call.
	Failures []*Failure `locationName:"failures" type:"list"`

	// A full description of the tasks that were run. Each task that was successfully
	// placed on your cluster will be described here.
	Tasks []*Task `locationName:"tasks" type:"list"`
}

// String returns the string representation
func (s RunTaskOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s RunTaskOutput) GoString() string {
	return s.String()
}

// SetFailures sets the Failures field's value.
func (s *RunTaskOutput) SetFailures(v []*Failure) *RunTaskOutput {
	s.Failures = v
	return s
}

// SetTasks sets the Tasks field's value.
func (s *RunTaskOutput) SetTasks(v []*Task) *RunTaskOutput {
	s.Tasks = v
	return s
}

// Details on a service within a cluster
type Service struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the of the cluster that hosts the service.
	ClusterArn *string `locationName:"clusterArn" type:"string"`

	CreatedAt *time.Time `locationName:"createdAt" type:"timestamp" timestampFormat:"unix"`

	DeploymentConfiguration *DeploymentConfiguration `locationName:"deploymentConfiguration" type:"structure"`

	// The current state of deployments for the service.
	Deployments []*Deployment `locationName:"deployments" type:"list"`

	// The desired number of instantiations of the task definition to keep running
	// on the service. This value is specified when the service is created with
	// CreateService, and it can be modified with UpdateService.
	DesiredCount *int64 `locationName:"desiredCount" type:"integer"`

	// The event stream for your service. A maximum of 100 of the latest events
	// are displayed.
	Events []*ServiceEvent `locationName:"events" type:"list"`

	LaunchType *string `locationName:"launchType" type:"string" enum:"LaunchType"`

	// A list of load balancer objects, containing the load balancer name, the container
	// name (as it appears in a container definition), and the container port to
	// access from the load balancer.
	LoadBalancers []*LoadBalancer `locationName:"loadBalancers" type:"list"`

	// The number of tasks in the cluster that are in the PENDING state.
	PendingCount *int64 `locationName:"pendingCount" type:"integer"`

	PlacementConstraints []*PlacementConstraint `locationName:"placementConstraints" type:"list"`

	PlacementStrategy []*PlacementStrategy `locationName:"placementStrategy" type:"list"`

	// The Amazon Resource Name (ARN) of the IAM role associated with the service
	// that allows the Amazon ECS container agent to register container instances
	// with a load balancer.
	RoleArn *string `locationName:"roleArn" type:"string"`

	// The number of tasks in the cluster that are in the RUNNING state.
	RunningCount *int64 `locationName:"runningCount" type:"integer"`

	// The Amazon Resource Name (ARN) that identifies the service. The ARN contains
	// the arn:aws:ecs namespace, followed by the region of the service, the AWS
	// account ID of the service owner, the service namespace, and then the service
	// name. For example, arn:aws:ecs:region:012345678910:service/my-service.
	ServiceArn *string `locationName:"serviceArn" type:"string"`

	// A user-generated string that you can use to identify your service.
	ServiceName *string `locationName:"serviceName" type:"string"`

	// The status of the service. The valid values are ACTIVE, DRAINING, or INACTIVE.
	Status *string `locationName:"status" type:"string"`

	// The task definition to use for tasks in the service. This value is specified
	// when the service is created with CreateService, and it can be modified with
	// UpdateService.
	TaskDefinition *string `locationName:"taskDefinition" type:"string"`
}

// String returns the string representation
func (s Service) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s Service) GoString() string {
	return s.String()
}

// SetClusterArn sets the ClusterArn field's value.
func (s *Service) SetClusterArn(v string) *Service {
	s.ClusterArn = &v
	return s
}

// SetCreatedAt sets the CreatedAt field's value.
func (s *Service) SetCreatedAt(v time.Time) *Service {
	s.CreatedAt = &v
	return s
}

// SetDeploymentConfiguration sets the DeploymentConfiguration field's value.
func (s *Service) SetDeploymentConfiguration(v *DeploymentConfiguration) *Service {
	s.DeploymentConfiguration = v
	return s
}

// SetDeployments sets the Deployments field's value.
func (s *Service) SetDeployments(v []*Deployment) *Service {
	s.Deployments = v
	return s
}

// SetDesiredCount sets the DesiredCount field's value.
func (s *Service) SetDesiredCount(v int64) *Service {
	s.DesiredCount = &v
	return s
}

// SetEvents sets the Events field's value.
func (s *Service) SetEvents(v []*ServiceEvent) *Service {
	s.Events = v
	return s
}

// SetLaunchType sets the LaunchType field's value.
func (s *Service) SetLaunchType(v string) *Service {
	s.LaunchType = &v
	return s
}

// SetLoadBalancers sets the LoadBalancers field's value.
func (s *Service) SetLoadBalancers(v []*LoadBalancer) *Service {
	s.LoadBalancers = v
	return s
}

// SetPendingCount sets the PendingCount field's value.
func (s *Service) SetPendingCount(v int64) *Service {
	s.PendingCount = &v
	return s
}

// SetPlacementConstraints sets the PlacementConstraints field's value.
func (s *Service) SetPlacementConstraints(v []*PlacementConstraint) *Service {
	s.PlacementConstraints = v
	return s
}

// SetPlacementStrategy sets the PlacementStrategy field's value.
func (s *Service) SetPlacementStrategy(v []*PlacementStrategy) *Service {
	s.PlacementStrategy = v
	return s
}

// SetRoleArn sets the RoleArn field's value.
func (s *Service) SetRoleArn(v string) *Service {
	s.RoleArn = &v
	return s
}

// SetRunningCount sets the RunningCount field's value.
func (s *Service) SetRunningCount(v int64) *Service {
	s.RunningCount = &v
	return s
}

// SetServiceArn sets the ServiceArn field's value.
func (s *Service) SetServiceArn(v string) *Service {
	s.ServiceArn = &v
	return s
}

// SetServiceName sets the ServiceName field's value.
func (s *Service) SetServiceName(v string) *Service {
	s.ServiceName = &v
	return s
}

// SetStatus sets the Status field's value.
func (s *Service) SetStatus(v string) *Service {
	s.Status = &v
	return s
}

// SetTaskDefinition sets the TaskDefinition field's value.
func (s *Service) SetTaskDefinition(v string) *Service {
	s.TaskDefinition = &v
	return s
}

// Details on an event associated with a service.
type ServiceEvent struct {
	_ struct{} `type:"structure"`

	// The Unix time in seconds and milliseconds when the event was triggered.
	CreatedAt *time.Time `locationName:"createdAt" type:"timestamp" timestampFormat:"unix"`

	// The ID string of the event.
	Id *string `locationName:"id" type:"string"`

	// The event message.
	Message *string `locationName:"message" type:"string"`
}

// String returns the string representation
func (s ServiceEvent) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s ServiceEvent) GoString() string {
	return s.String()
}

// SetCreatedAt sets the CreatedAt field's value.
func (s *ServiceEvent) SetCreatedAt(v time.Time) *ServiceEvent {
	s.CreatedAt = &v
	return s
}

// SetId sets the Id field's value.
func (s *ServiceEvent) SetId(v string) *ServiceEvent {
	s.Id = &v
	return s
}

// SetMessage sets the Message field's value.
func (s *ServiceEvent) SetMessage(v string) *ServiceEvent {
	s.Message = &v
	return s
}

type StartTaskInput struct {
	_ struct{} `type:"structure"`

	// The short name or full Amazon Resource Name (ARN) of the cluster that you
	// want to start your task on. If you do not specify a cluster, the default
	// cluster is assumed..
	Cluster *string `locationName:"cluster" type:"string"`

	// The container instance IDs or full Amazon Resource Name (ARN) entries for
	// the container instances on which you would like to place your task.
	//
	// The list of container instances to start tasks on is limited to 10.
	//
	// ContainerInstances is a required field
	ContainerInstances []*string `locationName:"containerInstances" type:"list" required:"true"`

	Group *string `locationName:"group" type:"string"`

	// A list of container overrides in JSON format that specify the name of a container
	// in the specified task definition and the overrides it should receive. You
	// can override the default command for a container (that is specified in the
	// task definition or Docker image) with a command override. You can also override
	// existing environment variables (that are specified in the task definition
	// or Docker image) on a container or add new environment variables to it with
	// an environment override.
	//
	// A total of 8192 characters are allowed for overrides. This limit includes
	// the JSON formatting characters of the override structure.
	Overrides *TaskOverride `locationName:"overrides" type:"structure"`

	// An optional tag specified when a task is started. For example if you automatically
	// trigger a task to run a batch process job, you could apply a unique identifier
	// for that job to your task with the startedBy parameter. You can then identify
	// which tasks belong to that job by filtering the results of a ListTasks call
	// with the startedBy value.
	//
	// If a task is started by an Amazon ECS service, then the startedBy parameter
	// contains the deployment ID of the service that starts it.
	StartedBy *string `locationName:"startedBy" type:"string"`

	// The family and revision (family:revision) or full Amazon Resource Name (ARN)
	// of the task definition that you want to start. If a revision is not specified,
	// the latest ACTIVE revision is used.
	//
	// TaskDefinition is a required field
	TaskDefinition *string `locationName:"taskDefinition" type:"string" required:"true"`
}

// String returns the string representation
func (s StartTaskInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s StartTaskInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *StartTaskInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "StartTaskInput"}
	if s.ContainerInstances == nil {
		invalidParams.Add(request.NewErrParamRequired("ContainerInstances"))
	}
	if s.TaskDefinition == nil {
		invalidParams.Add(request.NewErrParamRequired("TaskDefinition"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetCluster sets the Cluster field's value.
func (s *StartTaskInput) SetCluster(v string) *StartTaskInput {
	s.Cluster = &v
	return s
}

// SetContainerInstances sets the ContainerInstances field's value.
func (s *StartTaskInput) SetContainerInstances(v []*string) *StartTaskInput {
	s.ContainerInstances = v
	return s
}

// SetGroup sets the Group field's value.
func (s *StartTaskInput) SetGroup(v string) *StartTaskInput {
	s.Group = &v
	return s
}

// SetOverrides sets the Overrides field's value.
func (s *StartTaskInput) SetOverrides(v *TaskOverride) *StartTaskInput {
	s.Overrides = v
	return s
}

// SetStartedBy sets the StartedBy field's value.
func (s *StartTaskInput) SetStartedBy(v string) *StartTaskInput {
	s.StartedBy = &v
	return s
}

// SetTaskDefinition sets the TaskDefinition field's value.
func (s *StartTaskInput) SetTaskDefinition(v string) *StartTaskInput {
	s.TaskDefinition = &v
	return s
}

type StartTaskOutput struct {
	_ struct{} `type:"structure"`

	// Any failures associated with the call.
	Failures []*Failure `locationName:"failures" type:"list"`

	// A full description of the tasks that were started. Each task that was successfully
	// placed on your container instances will be described here.
	Tasks []*Task `locationName:"tasks" type:"list"`
}

// String returns the string representation
func (s StartTaskOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s StartTaskOutput) GoString() string {
	return s.String()
}

// SetFailures sets the Failures field's value.
func (s *StartTaskOutput) SetFailures(v []*Failure) *StartTaskOutput {
	s.Failures = v
	return s
}

// SetTasks sets the Tasks field's value.
func (s *StartTaskOutput) SetTasks(v []*Task) *StartTaskOutput {
	s.Tasks = v
	return s
}

type StopTaskInput struct {
	_ struct{} `type:"structure"`

	// The short name or full Amazon Resource Name (ARN) of the cluster that hosts
	// the task you want to stop. If you do not specify a cluster, the default cluster
	// is assumed..
	Cluster *string `locationName:"cluster" type:"string"`

	Reason *string `locationName:"reason" type:"string"`

	// The task ID or full Amazon Resource Name (ARN) entry of the task you would
	// like to stop.
	//
	// Task is a required field
	Task *string `locationName:"task" type:"string" required:"true"`
}

// String returns the string representation
func (s StopTaskInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s StopTaskInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *StopTaskInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "StopTaskInput"}
	if s.Task == nil {
		invalidParams.Add(request.NewErrParamRequired("Task"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetCluster sets the Cluster field's value.
func (s *StopTaskInput) SetCluster(v string) *StopTaskInput {
	s.Cluster = &v
	return s
}

// SetReason sets the Reason field's value.
func (s *StopTaskInput) SetReason(v string) *StopTaskInput {
	s.Reason = &v
	return s
}

// SetTask sets the Task field's value.
func (s *StopTaskInput) SetTask(v string) *StopTaskInput {
	s.Task = &v
	return s
}

type StopTaskOutput struct {
	_ struct{} `type:"structure"`

	// Details on a task in a cluster.
	Task *Task `locationName:"task" type:"structure"`
}

// String returns the string representation
func (s StopTaskOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s StopTaskOutput) GoString() string {
	return s.String()
}

// SetTask sets the Task field's value.
func (s *StopTaskOutput) SetTask(v *Task) *StopTaskOutput {
	s.Task = v
	return s
}

type SubmitContainerStateChangeInput struct {
	_ struct{} `type:"structure"`

	// The short name or full Amazon Resource Name (ARN) of the cluster that hosts
	// the container.
	Cluster *string `locationName:"cluster" type:"string"`

	// The name of the container.
	ContainerName *string `locationName:"containerName" type:"string"`

	// The exit code returned for the state change request.
	ExitCode *int64 `locationName:"exitCode" type:"integer"`

	// The network bindings of the container.
	NetworkBindings []*NetworkBinding `locationName:"networkBindings" type:"list"`

	// The reason for the state change request.
	Reason *string `locationName:"reason" type:"string"`

	// The status of the state change request.
	Status *string `locationName:"status" type:"string"`

	// The task ID or full Amazon Resource Name (ARN) of the task that hosts the
	// container.
	Task *string `locationName:"task" type:"string"`
}

// String returns the string representation
func (s SubmitContainerStateChangeInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s SubmitContainerStateChangeInput) GoString() string {
	return s.String()
}

// SetCluster sets the Cluster field's value.
func (s *SubmitContainerStateChangeInput) SetCluster(v string) *SubmitContainerStateChangeInput {
	s.Cluster = &v
	return s
}

// SetContainerName sets the ContainerName field's value.
func (s *SubmitContainerStateChangeInput) SetContainerName(v string) *SubmitContainerStateChangeInput {
	s.ContainerName = &v
	return s
}

// SetExitCode sets the ExitCode field's value.
func (s *SubmitContainerStateChangeInput) SetExitCode(v int64) *SubmitContainerStateChangeInput {
	s.ExitCode = &v
	return s
}

// SetNetworkBindings sets the NetworkBindings field's value.
func (s *SubmitContainerStateChangeInput) SetNetworkBindings(v []*NetworkBinding) *SubmitContainerStateChangeInput {
	s.NetworkBindings = v
	return s
}

// SetReason sets the Reason field's value.
func (s *SubmitContainerStateChangeInput) SetReason(v string) *SubmitContainerStateChangeInput {
	s.Reason = &v
	return s
}

// SetStatus sets the Status field's value.
func (s *SubmitContainerStateChangeInput) SetStatus(v string) *SubmitContainerStateChangeInput {
	s.Status = &v
	return s
}

// SetTask sets the Task field's value.
func (s *SubmitContainerStateChangeInput) SetTask(v string) *SubmitContainerStateChangeInput {
	s.Task = &v
	return s
}

type SubmitContainerStateChangeOutput struct {
	_ struct{} `type:"structure"`

	// Acknowledgement of the state change.
	Acknowledgment *string `locationName:"acknowledgment" type:"string"`
}

// String returns the string representation
func (s SubmitContainerStateChangeOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s SubmitContainerStateChangeOutput) GoString() string {
	return s.String()
}

// SetAcknowledgment sets the Acknowledgment field's value.
func (s *SubmitContainerStateChangeOutput) SetAcknowledgment(v string) *SubmitContainerStateChangeOutput {
	s.Acknowledgment = &v
	return s
}

type SubmitTaskStateChangeInput struct {
	_ struct{} `type:"structure"`

	// The short name or full Amazon Resource Name (ARN) of the cluster that hosts
	// the task.
	Cluster *string `locationName:"cluster" type:"string"`

	Containers []*ContainerStateChange `locationName:"containers" type:"list"`

	// The reason for the state change request.
	Reason *string `locationName:"reason" type:"string"`

	// The status of the state change request.
	Status *string `locationName:"status" type:"string"`

	// The task ID or full Amazon Resource Name (ARN) of the task in the state change
	// request.
	Task *string `locationName:"task" type:"string"`
}

// String returns the string representation
func (s SubmitTaskStateChangeInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s SubmitTaskStateChangeInput) GoString() string {
	return s.String()
}

// SetCluster sets the Cluster field's value.
func (s *SubmitTaskStateChangeInput) SetCluster(v string) *SubmitTaskStateChangeInput {
	s.Cluster = &v
	return s
}

// SetContainers sets the Containers field's value.
func (s *SubmitTaskStateChangeInput) SetContainers(v []*ContainerStateChange) *SubmitTaskStateChangeInput {
	s.Containers = v
	return s
}

// SetReason sets the Reason field's value.
func (s *SubmitTaskStateChangeInput) SetReason(v string) *SubmitTaskStateChangeInput {
	s.Reason = &v
	return s
}

// SetStatus sets the Status field's value.
func (s *SubmitTaskStateChangeInput) SetStatus(v string) *SubmitTaskStateChangeInput {
	s.Status = &v
	return s
}

// SetTask sets the Task field's value.
func (s *SubmitTaskStateChangeInput) SetTask(v string) *SubmitTaskStateChangeInput {
	s.Task = &v
	return s
}

type SubmitTaskStateChangeOutput struct {
	_ struct{} `type:"structure"`

	// Acknowledgement of the state change.
	Acknowledgment *string `locationName:"acknowledgment" type:"string"`
}

// String returns the string representation
func (s SubmitTaskStateChangeOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s SubmitTaskStateChangeOutput) GoString() string {
	return s.String()
}

// SetAcknowledgment sets the Acknowledgment field's value.
func (s *SubmitTaskStateChangeOutput) SetAcknowledgment(v string) *SubmitTaskStateChangeOutput {
	s.Acknowledgment = &v
	return s
}

// Details on a task in a cluster.
type Task struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the of the cluster that hosts the task.
	ClusterArn *string `locationName:"clusterArn" type:"string"`

	// The Amazon Resource Name (ARN) of the container instances that host the task.
	ContainerInstanceArn *string `locationName:"containerInstanceArn" type:"string"`

	// The containers associated with the task.
	Containers []*Container `locationName:"containers" type:"list"`

	CreatedAt *time.Time `locationName:"createdAt" type:"timestamp" timestampFormat:"unix"`

	// The desired status of the task.
	DesiredStatus *string `locationName:"desiredStatus" type:"string"`

	Group *string `locationName:"group" type:"string"`

	// The last known status of the task.
	LastStatus *string `locationName:"lastStatus" type:"string"`

	LaunchType *string `locationName:"launchType" type:"string" enum:"LaunchType"`

	// One or more container overrides.
	Overrides *TaskOverride `locationName:"overrides" type:"structure"`

	RoleArn *string `locationName:"roleArn" type:"string"`

	StartedAt *time.Time `locationName:"startedAt" type:"timestamp" timestampFormat:"unix"`

	// The tag specified when a task is started. If the task is started by an Amazon
	// ECS service, then the startedBy parameter contains the deployment ID of the
	// service that starts it.
	StartedBy *string `locationName:"startedBy" type:"string"`

	StoppedAt *time.Time `locationName:"stoppedAt" type:"timestamp" timestampFormat:"unix"`

	StoppedReason *string `locationName:"stoppedReason" type:"string"`

	// The Amazon Resource Name (ARN) of the task.
	TaskArn *string `locationName:"taskArn" type:"string"`

	// The Amazon Resource Name (ARN) of the of the task definition that creates
	// the task.
	TaskDefinitionArn *string `locationName:"taskDefinitionArn" type:"string"`

	Version *int64 `locationName:"version" type:"long"`
}

// String returns the string representation
func (s Task) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s Task) GoString() string {
	return s.String()
}

// SetClusterArn sets the ClusterArn field's value.
func (s *Task) SetClusterArn(v string) *Task {
	s.ClusterArn = &v
	return s
}

// SetContainerInstanceArn sets the ContainerInstanceArn field's value.
func (s *Task) SetContainerInstanceArn(v string) *Task {
	s.ContainerInstanceArn = &v
	return s
}

// SetContainers sets the Containers field's value.
func (s *Task) SetContainers(v []*Container) *Task {
	s.Containers = v
	return s
}

// SetCreatedAt sets the CreatedAt field's value.
func (s *Task) SetCreatedAt(v time.Time) *Task {
	s.CreatedAt = &v
	return s
}

// SetDesiredStatus sets the DesiredStatus field's value.
func (s *Task) SetDesiredStatus(v string) *Task {
	s.DesiredStatus = &v
	return s
}

// SetGroup sets the Group field's value.
func (s *Task) SetGroup(v string) *Task {
	s.Group = &v
	return s
}

// SetLastStatus sets the LastStatus field's value.
func (s *Task) SetLastStatus(v string) *Task {
	s.LastStatus = &v
	return s
}

// SetLaunchType sets the LaunchType field's value.
func (s *Task) SetLaunchType(v string) *Task {
	s.LaunchType = &v
	return s
}

// SetOverrides sets the Overrides field's value.
func (s *Task) SetOverrides(v *TaskOverride) *Task {
	s.Overrides = v
	return s
}

// SetRoleArn sets the RoleArn field's value.
func (s *Task) SetRoleArn(v string) *Task {
	s.RoleArn = &v
	return s
}

// SetStartedAt sets the StartedAt field's value.
func (s *Task) SetStartedAt(v time.Time) *Task {
	s.StartedAt = &v
	return s
}

// SetStartedBy sets the StartedBy field's value.
func (s *Task) SetStartedBy(v string) *Task {
	s.StartedBy = &v
	return s
}

// SetStoppedAt sets the StoppedAt field's value.
func (s *Task) SetStoppedAt(v time.Time) *Task {
	s.StoppedAt = &v
	return s
}

// SetStoppedReason sets the StoppedReason field's value.
func (s *Task) SetStoppedReason(v string) *Task {
	s.StoppedReason = &v
	return s
}

// SetTaskArn sets the TaskArn field's value.
func (s *Task) SetTaskArn(v string) *Task {
	s.TaskArn = &v
	return s
}

// SetTaskDefinitionArn sets the TaskDefinitionArn field's value.
func (s *Task) SetTaskDefinitionArn(v string) *Task {
	s.TaskDefinitionArn = &v
	return s
}

// SetVersion sets the Version field's value.
func (s *Task) SetVersion(v int64) *Task {
	s.Version = &v
	return s
}

// Details of a task definition.
type TaskDefinition struct {
	_ struct{} `type:"structure"`

	// A list of container definitions in JSON format that describe the different
	// containers that make up your task. For more information on container definition
	// parameters and defaults, see Amazon ECS Task Definitions (http://docs.aws.amazon.com/AmazonECS/latest/developerguide/task_defintions.html)
	// in the Amazon EC2 Container Service Developer Guide.
	ContainerDefinitions []*ContainerDefinition `locationName:"containerDefinitions" type:"list"`

	// The family of your task definition. You can think of the family as the name
	// of your task definition.
	Family *string `locationName:"family" type:"string"`

	NetworkMode *string `locationName:"networkMode" type:"string" enum:"NetworkMode"`

	PlacementConstraints []*TaskDefinitionPlacementConstraint `locationName:"placementConstraints" type:"list"`

	// The container instance attributes required by your task.
	RequiresAttributes []*Attribute `locationName:"requiresAttributes" type:"list"`

	// The revision of the task in a particular family. You can think of the revision
	// as a version number of a task definition in a family. When you register a
	// task definition for the first time, the revision is 1, and each time you
	// register a new revision of a task definition in the same family, the revision
	// value always increases by one (even if you have deregistered previous revisions
	// in this family).
	Revision *int64 `locationName:"revision" type:"integer"`

	// The status of the task definition.
	Status *string `locationName:"status" type:"string" enum:"TaskDefinitionStatus"`

	// The full Amazon Resource Name (ARN) of the of the task definition.
	TaskDefinitionArn *string `locationName:"taskDefinitionArn" type:"string"`

	TaskRoleArn *string `locationName:"taskRoleArn" type:"string"`

	// The list of volumes in a task. For more information on volume definition
	// parameters and defaults, see Amazon ECS Task Definitions (http://docs.aws.amazon.com/AmazonECS/latest/developerguide/task_defintions.html)
	// in the Amazon EC2 Container Service Developer Guide.
	Volumes []*Volume `locationName:"volumes" type:"list"`
}

// String returns the string representation
func (s TaskDefinition) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s TaskDefinition) GoString() string {
	return s.String()
}

// SetContainerDefinitions sets the ContainerDefinitions field's value.
func (s *TaskDefinition) SetContainerDefinitions(v []*ContainerDefinition) *TaskDefinition {
	s.ContainerDefinitions = v
	return s
}

// SetFamily sets the Family field's value.
func (s *TaskDefinition) SetFamily(v string) *TaskDefinition {
	s.Family = &v
	return s
}

// SetNetworkMode sets the NetworkMode field's value.
func (s *TaskDefinition) SetNetworkMode(v string) *TaskDefinition {
	s.NetworkMode = &v
	return s
}

// SetPlacementConstraints sets the PlacementConstraints field's value.
func (s *TaskDefinition) SetPlacementConstraints(v []*TaskDefinitionPlacementConstraint) *TaskDefinition {
	s.PlacementConstraints = v
	return s
}

// SetRequiresAttributes sets the RequiresAttributes field's value.
func (s *TaskDefinition) SetRequiresAttributes(v []*Attribute) *TaskDefinition {
	s.RequiresAttributes = v
	return s
}

// SetRevision sets the Revision field's value.
func (s *TaskDefinition) SetRevision(v int64) *TaskDefinition {
	s.Revision = &v
	return s
}

// SetStatus sets the Status field's value.
func (s *TaskDefinition) SetStatus(v string) *TaskDefinition {
	s.Status = &v
	return s
}

// SetTaskDefinitionArn sets the TaskDefinitionArn field's value.
func (s *TaskDefinition) SetTaskDefinitionArn(v string) *TaskDefinition {
	s.TaskDefinitionArn = &v
	return s
}

// SetTaskRoleArn sets the TaskRoleArn field's value.
func (s *TaskDefinition) SetTaskRoleArn(v string) *TaskDefinition {
	s.TaskRoleArn = &v
	return s
}

// SetVolumes sets the Volumes field's value.
func (s *TaskDefinition) SetVolumes(v []*Volume) *TaskDefinition {
	s.Volumes = v
	return s
}

type TaskDefinitionPlacementConstraint struct {
	_ struct{} `type:"structure"`

	Expression *string `locationName:"expression" type:"string"`

	Type *string `locationName:"type" type:"string" enum:"TaskDefinitionPlacementConstraintType"`
}

// String returns the string representation
func (s TaskDefinitionPlacementConstraint) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s TaskDefinitionPlacementConstraint) GoString() string {
	return s.String()
}

// SetExpression sets the Expression field's value.
func (s *TaskDefinitionPlacementConstraint) SetExpression(v string) *TaskDefinitionPlacementConstraint {
	s.Expression = &v
	return s
}

// SetType sets the Type field's value.
func (s *TaskDefinitionPlacementConstraint) SetType(v string) *TaskDefinitionPlacementConstraint {
	s.Type = &v
	return s
}

// The overrides associated with a task.
type TaskOverride struct {
	_ struct{} `type:"structure"`

	// One or more container overrides sent to a task.
	ContainerOverrides []*ContainerOverride `locationName:"containerOverrides" type:"list"`

	TaskRoleArn *string `locationName:"taskRoleArn" type:"string"`
}

// String returns the string representation
func (s TaskOverride) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s TaskOverride) GoString() string {
	return s.String()
}

// SetContainerOverrides sets the ContainerOverrides field's value.
func (s *TaskOverride) SetContainerOverrides(v []*ContainerOverride) *TaskOverride {
	s.ContainerOverrides = v
	return s
}

// SetTaskRoleArn sets the TaskRoleArn field's value.
func (s *TaskOverride) SetTaskRoleArn(v string) *TaskOverride {
	s.TaskRoleArn = &v
	return s
}

// The ulimit settings to pass to the container.
type Ulimit struct {
	_ struct{} `type:"structure"`

	// The hard limit for the ulimit type.
	//
	// HardLimit is a required field
	HardLimit *int64 `locationName:"hardLimit" type:"integer" required:"true"`

	// The type of the ulimit.
	//
	// Name is a required field
	Name *string `locationName:"name" type:"string" required:"true" enum:"UlimitName"`

	// The soft limit for the ulimit type.
	//
	// SoftLimit is a required field
	SoftLimit *int64 `locationName:"softLimit" type:"integer" required:"true"`
}

// String returns the string representation
func (s Ulimit) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s Ulimit) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *Ulimit) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "Ulimit"}
	if s.HardLimit == nil {
		invalidParams.Add(request.NewErrParamRequired("HardLimit"))
	}
	if s.Name == nil {
		invalidParams.Add(request.NewErrParamRequired("Name"))
	}
	if s.SoftLimit == nil {
		invalidParams.Add(request.NewErrParamRequired("SoftLimit"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetHardLimit sets the HardLimit field's value.
func (s *Ulimit) SetHardLimit(v int64) *Ulimit {
	s.HardLimit = &v
	return s
}

// SetName sets the Name field's value.
func (s *Ulimit) SetName(v string) *Ulimit {
	s.Name = &v
	return s
}

// SetSoftLimit sets the SoftLimit field's value.
func (s *Ulimit) SetSoftLimit(v int64) *Ulimit {
	s.SoftLimit = &v
	return s
}

type UpdateContainerAgentInput struct {
	_ struct{} `type:"structure"`

	// The short name or full Amazon Resource Name (ARN) of the cluster that your
	// container instance is running on. If you do not specify a cluster, the default
	// cluster is assumed.
	Cluster *string `locationName:"cluster" type:"string"`

	// The container instance ID or full Amazon Resource Name (ARN) entries for
	// the container instance on which you would like to update the Amazon ECS container
	// agent.
	//
	// ContainerInstance is a required field
	ContainerInstance *string `locationName:"containerInstance" type:"string" required:"true"`
}

// String returns the string representation
func (s UpdateContainerAgentInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s UpdateContainerAgentInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateContainerAgentInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "UpdateContainerAgentInput"}
	if s.ContainerInstance == nil {
		invalidParams.Add(request.NewErrParamRequired("ContainerInstance"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetCluster sets the Cluster field's value.
func (s *UpdateContainerAgentInput) SetCluster(v string) *UpdateContainerAgentInput {
	s.Cluster = &v
	return s
}

// SetContainerInstance sets the ContainerInstance field's value.
func (s *UpdateContainerAgentInput) SetContainerInstance(v string) *UpdateContainerAgentInput {
	s.ContainerInstance = &v
	return s
}

type UpdateContainerAgentOutput struct {
	_ struct{} `type:"structure"`

	// An Amazon EC2 instance that is running the Amazon ECS agent and has been
	// registered with a cluster.
	ContainerInstance *ContainerInstance `locationName:"containerInstance" type:"structure"`
}

// String returns the string representation
func (s UpdateContainerAgentOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s UpdateContainerAgentOutput) GoString() string {
	return s.String()
}

// SetContainerInstance sets the ContainerInstance field's value.
func (s *UpdateContainerAgentOutput) SetContainerInstance(v *ContainerInstance) *UpdateContainerAgentOutput {
	s.ContainerInstance = v
	return s
}

type UpdateContainerInstancesStateInput struct {
	_ struct{} `type:"structure"`

	Cluster *string `locationName:"cluster" type:"string"`

	// ContainerInstances is a required field
	ContainerInstances []*string `locationName:"containerInstances" type:"list" required:"true"`

	// Status is a required field
	Status *string `locationName:"status" type:"string" required:"true" enum:"ContainerInstanceStatus"`
}

// String returns the string representation
func (s UpdateContainerInstancesStateInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s UpdateContainerInstancesStateInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateContainerInstancesStateInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "UpdateContainerInstancesStateInput"}
	if s.ContainerInstances == nil {
		invalidParams.Add(request.NewErrParamRequired("ContainerInstances"))
	}
	if s.Status == nil {
		invalidParams.Add(request.NewErrParamRequired("Status"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetCluster sets the Cluster field's value.
func (s *UpdateContainerInstancesStateInput) SetCluster(v string) *UpdateContainerInstancesStateInput {
	s.Cluster = &v
	return s
}

// SetContainerInstances sets the ContainerInstances field's value.
func (s *UpdateContainerInstancesStateInput) SetContainerInstances(v []*string) *UpdateContainerInstancesStateInput {
	s.ContainerInstances = v
	return s
}

// SetStatus sets the Status field's value.
func (s *UpdateContainerInstancesStateInput) SetStatus(v string) *UpdateContainerInstancesStateInput {
	s.Status = &v
	return s
}

type UpdateContainerInstancesStateOutput struct {
	_ struct{} `type:"structure"`

	ContainerInstances []*ContainerInstance `locationName:"containerInstances" type:"list"`

	Failures []*Failure `locationName:"failures" type:"list"`
}

// String returns the string representation
func (s UpdateContainerInstancesStateOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s UpdateContainerInstancesStateOutput) GoString() string {
	return s.String()
}

// SetContainerInstances sets the ContainerInstances field's value.
func (s *UpdateContainerInstancesStateOutput) SetContainerInstances(v []*ContainerInstance) *UpdateContainerInstancesStateOutput {
	s.ContainerInstances = v
	return s
}

// SetFailures sets the Failures field's value.
func (s *UpdateContainerInstancesStateOutput) SetFailures(v []*Failure) *UpdateContainerInstancesStateOutput {
	s.Failures = v
	return s
}

type UpdateServiceInput struct {
	_ struct{} `type:"structure"`

	// The short name or full Amazon Resource Name (ARN) of the cluster that your
	// service is running on. If you do not specify a cluster, the default cluster
	// is assumed.
	Cluster *string `locationName:"cluster" type:"string"`

	DeploymentConfiguration *DeploymentConfiguration `locationName:"deploymentConfiguration" type:"structure"`

	// The number of instantiations of the task that you would like to place and
	// keep running in your service.
	DesiredCount *int64 `locationName:"desiredCount" type:"integer"`

	// The name of the service that you want to update.
	//
	// Service is a required field
	Service *string `locationName:"service" type:"string" required:"true"`

	// The family and revision (family:revision) or full Amazon Resource Name (ARN)
	// of the task definition that you want to run in your service. If a revision
	// is not specified, the latest ACTIVE revision is used. If you modify the task
	// definition with UpdateService, Amazon ECS spawns a task with the new version
	// of the task definition and then stops an old task after the new version is
	// running.
	TaskDefinition *string `locationName:"taskDefinition" type:"string"`
}

// String returns the string representation
func (s UpdateServiceInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s UpdateServiceInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateServiceInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "UpdateServiceInput"}
	if s.Service == nil {
		invalidParams.Add(request.NewErrParamRequired("Service"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetCluster sets the Cluster field's value.
func (s *UpdateServiceInput) SetCluster(v string) *UpdateServiceInput {
	s.Cluster = &v
	return s
}

// SetDeploymentConfiguration sets the DeploymentConfiguration field's value.
func (s *UpdateServiceInput) SetDeploymentConfiguration(v *DeploymentConfiguration) *UpdateServiceInput {
	s.DeploymentConfiguration = v
	return s
}

// SetDesiredCount sets the DesiredCount field's value.
func (s *UpdateServiceInput) SetDesiredCount(v int64) *UpdateServiceInput {
	s.DesiredCount = &v
	return s
}

// SetService sets the Service field's value.
func (s *UpdateServiceInput) SetService(v string) *UpdateServiceInput {
	s.Service = &v
	return s
}

// SetTaskDefinition sets the TaskDefinition field's value.
func (s *UpdateServiceInput) SetTaskDefinition(v string) *UpdateServiceInput {
	s.TaskDefinition = &v
	return s
}

type UpdateServiceOutput struct {
	_ struct{} `type:"structure"`

	// The full description of your service following the update call.
	Service *Service `locationName:"service" type:"structure"`
}

// String returns the string representation
func (s UpdateServiceOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s UpdateServiceOutput) GoString() string {
	return s.String()
}

// SetService sets the Service field's value.
func (s *UpdateServiceOutput) SetService(v *Service) *UpdateServiceOutput {
	s.Service = v
	return s
}

// The Docker and Amazon ECS container agent version information on a container
// instance.
type VersionInfo struct {
	_ struct{} `type:"structure"`

	// The Git commit hash for the Amazon ECS container agent build on the amazon-ecs-agent
	//  (https://github.com/aws/amazon-ecs-agent/commits/master) GitHub repository.
	AgentHash *string `locationName:"agentHash" type:"string"`

	// The version number of the Amazon ECS container agent.
	AgentVersion *string `locationName:"agentVersion" type:"string"`

	// The Docker version running on the container instance.
	DockerVersion *string `locationName:"dockerVersion" type:"string"`
}

// String returns the string representation
func (s VersionInfo) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s VersionInfo) GoString() string {
	return s.String()
}

// SetAgentHash sets the AgentHash field's value.
func (s *VersionInfo) SetAgentHash(v string) *VersionInfo {
	s.AgentHash = &v
	return s
}

// SetAgentVersion sets the AgentVersion field's value.
func (s *VersionInfo) SetAgentVersion(v string) *VersionInfo {
	s.AgentVersion = &v
	return s
}

// SetDockerVersion sets the DockerVersion field's value.
func (s *VersionInfo) SetDockerVersion(v string) *VersionInfo {
	s.DockerVersion = &v
	return s
}

// A data volume used in a task definition.
type Volume struct {
	_ struct{} `type:"structure"`

	// The path on the host container instance that is presented to the containers
	// which access the volume. If this parameter is empty, then the Docker daemon
	// assigns a host path for you.
	Host *HostVolumeProperties `locationName:"host" type:"structure"`

	// The name of the volume. This name is referenced in the sourceVolume parameter
	// of container definition mountPoints.
	Name *string `locationName:"name" type:"string"`
}

// String returns the string representation
func (s Volume) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s Volume) GoString() string {
	return s.String()
}

// SetHost sets the Host field's value.
func (s *Volume) SetHost(v *HostVolumeProperties) *Volume {
	s.Host = v
	return s
}

// SetName sets the Name field's value.
func (s *Volume) SetName(v string) *Volume {
	s.Name = &v
	return s
}

// Details on a data volume from another container.
type VolumeFrom struct {
	_ struct{} `type:"structure"`

	// If this value is true, the container has read-only access to the volume.
	// If this value is false, then the container can write to the volume. The default
	// value is false.
	ReadOnly *bool `locationName:"readOnly" type:"boolean"`

	// The name of the container to mount volumes from.
	SourceContainer *string `locationName:"sourceContainer" type:"string"`
}

// String returns the string representation
func (s VolumeFrom) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s VolumeFrom) GoString() string {
	return s.String()
}

// SetReadOnly sets the ReadOnly field's value.
func (s *VolumeFrom) SetReadOnly(v bool) *VolumeFrom {
	s.ReadOnly = &v
	return s
}

// SetSourceContainer sets the SourceContainer field's value.
func (s *VolumeFrom) SetSourceContainer(v string) *VolumeFrom {
	s.SourceContainer = &v
	return s
}

const (
	// AgentUpdateStatusPending is a AgentUpdateStatus enum value
	AgentUpdateStatusPending = "PENDING"

	// AgentUpdateStatusStaging is a AgentUpdateStatus enum value
	AgentUpdateStatusStaging = "STAGING"

	// AgentUpdateStatusStaged is a AgentUpdateStatus enum value
	AgentUpdateStatusStaged = "STAGED"

	// AgentUpdateStatusUpdating is a AgentUpdateStatus enum value
	AgentUpdateStatusUpdating = "UPDATING"

	// AgentUpdateStatusUpdated is a AgentUpdateStatus enum value
	AgentUpdateStatusUpdated = "UPDATED"

	// AgentUpdateStatusFailed is a AgentUpdateStatus enum value
	AgentUpdateStatusFailed = "FAILED"
)

const (
	// ContainerInstanceStatusActive is a ContainerInstanceStatus enum value
	ContainerInstanceStatusActive = "ACTIVE"

	// ContainerInstanceStatusDraining is a ContainerInstanceStatus enum value
	ContainerInstanceStatusDraining = "DRAINING"
)

const (
	// DesiredStatusRunning is a DesiredStatus enum value
	DesiredStatusRunning = "RUNNING"

	// DesiredStatusPending is a DesiredStatus enum value
	DesiredStatusPending = "PENDING"

	// DesiredStatusStopped is a DesiredStatus enum value
	DesiredStatusStopped = "STOPPED"
)

const (
	// LaunchTypeStandard is a LaunchType enum value
	LaunchTypeStandard = "STANDARD"

	// LaunchTypeServerless is a LaunchType enum value
	LaunchTypeServerless = "SERVERLESS"
)

const (
	// LogDriverJsonFile is a LogDriver enum value
	LogDriverJsonFile = "json-file"

	// LogDriverSyslog is a LogDriver enum value
	LogDriverSyslog = "syslog"

	// LogDriverJournald is a LogDriver enum value
	LogDriverJournald = "journald"

	// LogDriverGelf is a LogDriver enum value
	LogDriverGelf = "gelf"

	// LogDriverFluentd is a LogDriver enum value
	LogDriverFluentd = "fluentd"

	// LogDriverAwslogs is a LogDriver enum value
	LogDriverAwslogs = "awslogs"

	// LogDriverSplunk is a LogDriver enum value
	LogDriverSplunk = "splunk"
)

const (
	// NetworkModeBridge is a NetworkMode enum value
	NetworkModeBridge = "bridge"

	// NetworkModeHost is a NetworkMode enum value
	NetworkModeHost = "host"

	// NetworkModeNone is a NetworkMode enum value
	NetworkModeNone = "none"
)

const (
	// PlacementConstraintTypeDistinctInstance is a PlacementConstraintType enum value
	PlacementConstraintTypeDistinctInstance = "distinctInstance"

	// PlacementConstraintTypeMemberOf is a PlacementConstraintType enum value
	PlacementConstraintTypeMemberOf = "memberOf"
)

const (
	// PlacementStrategyTypeRandom is a PlacementStrategyType enum value
	PlacementStrategyTypeRandom = "random"

	// PlacementStrategyTypeSpread is a PlacementStrategyType enum value
	PlacementStrategyTypeSpread = "spread"

	// PlacementStrategyTypeBinpack is a PlacementStrategyType enum value
	PlacementStrategyTypeBinpack = "binpack"
)

const (
	// SortOrderAsc is a SortOrder enum value
	SortOrderAsc = "ASC"

	// SortOrderDesc is a SortOrder enum value
	SortOrderDesc = "DESC"
)

const (
	// TargetTypeContainerInstance is a TargetType enum value
	TargetTypeContainerInstance = "container-instance"
)

const (
	// TaskDefinitionFamilyStatusActive is a TaskDefinitionFamilyStatus enum value
	TaskDefinitionFamilyStatusActive = "ACTIVE"

	// TaskDefinitionFamilyStatusInactive is a TaskDefinitionFamilyStatus enum value
	TaskDefinitionFamilyStatusInactive = "INACTIVE"

	// TaskDefinitionFamilyStatusAll is a TaskDefinitionFamilyStatus enum value
	TaskDefinitionFamilyStatusAll = "ALL"
)

const (
	// TaskDefinitionPlacementConstraintTypeMemberOf is a TaskDefinitionPlacementConstraintType enum value
	TaskDefinitionPlacementConstraintTypeMemberOf = "memberOf"
)

const (
	// TaskDefinitionStatusActive is a TaskDefinitionStatus enum value
	TaskDefinitionStatusActive = "ACTIVE"

	// TaskDefinitionStatusInactive is a TaskDefinitionStatus enum value
	TaskDefinitionStatusInactive = "INACTIVE"
)

const (
	// TransportProtocolTcp is a TransportProtocol enum value
	TransportProtocolTcp = "tcp"

	// TransportProtocolUdp is a TransportProtocol enum value
	TransportProtocolUdp = "udp"
)

const (
	// UlimitNameCore is a UlimitName enum value
	UlimitNameCore = "core"

	// UlimitNameCpu is a UlimitName enum value
	UlimitNameCpu = "cpu"

	// UlimitNameData is a UlimitName enum value
	UlimitNameData = "data"

	// UlimitNameFsize is a UlimitName enum value
	UlimitNameFsize = "fsize"

	// UlimitNameLocks is a UlimitName enum value
	UlimitNameLocks = "locks"

	// UlimitNameMemlock is a UlimitName enum value
	UlimitNameMemlock = "memlock"

	// UlimitNameMsgqueue is a UlimitName enum value
	UlimitNameMsgqueue = "msgqueue"

	// UlimitNameNice is a UlimitName enum value
	UlimitNameNice = "nice"

	// UlimitNameNofile is a UlimitName enum value
	UlimitNameNofile = "nofile"

	// UlimitNameNproc is a UlimitName enum value
	UlimitNameNproc = "nproc"

	// UlimitNameRss is a UlimitName enum value
	UlimitNameRss = "rss"

	// UlimitNameRtprio is a UlimitName enum value
	UlimitNameRtprio = "rtprio"

	// UlimitNameRttime is a UlimitName enum value
	UlimitNameRttime = "rttime"

	// UlimitNameSigpending is a UlimitName enum value
	UlimitNameSigpending = "sigpending"

	// UlimitNameStack is a UlimitName enum value
	UlimitNameStack = "stack"
)
