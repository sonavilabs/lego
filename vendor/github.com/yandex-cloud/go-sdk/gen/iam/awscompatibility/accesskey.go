// Code generated by sdkgen. DO NOT EDIT.

//nolint
package awscompatibility

import (
	"context"

	"google.golang.org/grpc"

	awscompatibility "github.com/yandex-cloud/go-genproto/yandex/cloud/iam/v1/awscompatibility"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/operation"
)

//revive:disable

// AccessKeyServiceClient is a awscompatibility.AccessKeyServiceClient with
// lazy GRPC connection initialization.
type AccessKeyServiceClient struct {
	getConn func(ctx context.Context) (*grpc.ClientConn, error)
}

// Create implements awscompatibility.AccessKeyServiceClient
func (c *AccessKeyServiceClient) Create(ctx context.Context, in *awscompatibility.CreateAccessKeyRequest, opts ...grpc.CallOption) (*awscompatibility.CreateAccessKeyResponse, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return awscompatibility.NewAccessKeyServiceClient(conn).Create(ctx, in, opts...)
}

// Delete implements awscompatibility.AccessKeyServiceClient
func (c *AccessKeyServiceClient) Delete(ctx context.Context, in *awscompatibility.DeleteAccessKeyRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return awscompatibility.NewAccessKeyServiceClient(conn).Delete(ctx, in, opts...)
}

// Get implements awscompatibility.AccessKeyServiceClient
func (c *AccessKeyServiceClient) Get(ctx context.Context, in *awscompatibility.GetAccessKeyRequest, opts ...grpc.CallOption) (*awscompatibility.AccessKey, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return awscompatibility.NewAccessKeyServiceClient(conn).Get(ctx, in, opts...)
}

// List implements awscompatibility.AccessKeyServiceClient
func (c *AccessKeyServiceClient) List(ctx context.Context, in *awscompatibility.ListAccessKeysRequest, opts ...grpc.CallOption) (*awscompatibility.ListAccessKeysResponse, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return awscompatibility.NewAccessKeyServiceClient(conn).List(ctx, in, opts...)
}

type AccessKeyIterator struct {
	ctx  context.Context
	opts []grpc.CallOption

	err           error
	started       bool
	requestedSize int64
	pageSize      int64

	client  *AccessKeyServiceClient
	request *awscompatibility.ListAccessKeysRequest

	items []*awscompatibility.AccessKey
}

func (c *AccessKeyServiceClient) AccessKeyIterator(ctx context.Context, req *awscompatibility.ListAccessKeysRequest, opts ...grpc.CallOption) *AccessKeyIterator {
	var pageSize int64
	const defaultPageSize = 1000
	pageSize = req.PageSize
	if pageSize == 0 {
		pageSize = defaultPageSize
	}
	return &AccessKeyIterator{
		ctx:      ctx,
		opts:     opts,
		client:   c,
		request:  req,
		pageSize: pageSize,
	}
}

func (it *AccessKeyIterator) Next() bool {
	if it.err != nil {
		return false
	}
	if len(it.items) > 1 {
		it.items[0] = nil
		it.items = it.items[1:]
		return true
	}
	it.items = nil // consume last item, if any

	if it.started && it.request.PageToken == "" {
		return false
	}
	it.started = true

	if it.requestedSize == 0 || it.requestedSize > it.pageSize {
		it.request.PageSize = it.pageSize
	} else {
		it.request.PageSize = it.requestedSize
	}

	response, err := it.client.List(it.ctx, it.request, it.opts...)
	it.err = err
	if err != nil {
		return false
	}

	it.items = response.AccessKeys
	it.request.PageToken = response.NextPageToken
	return len(it.items) > 0
}

func (it *AccessKeyIterator) Take(size int64) ([]*awscompatibility.AccessKey, error) {
	if it.err != nil {
		return nil, it.err
	}

	if size == 0 {
		size = 1 << 32 // something insanely large
	}
	it.requestedSize = size
	defer func() {
		// reset iterator for future calls.
		it.requestedSize = 0
	}()

	var result []*awscompatibility.AccessKey

	for it.requestedSize > 0 && it.Next() {
		it.requestedSize--
		result = append(result, it.Value())
	}

	if it.err != nil {
		return nil, it.err
	}

	return result, nil
}

func (it *AccessKeyIterator) TakeAll() ([]*awscompatibility.AccessKey, error) {
	return it.Take(0)
}

func (it *AccessKeyIterator) Value() *awscompatibility.AccessKey {
	if len(it.items) == 0 {
		panic("calling Value on empty iterator")
	}
	return it.items[0]
}

func (it *AccessKeyIterator) Error() error {
	return it.err
}

// ListOperations implements awscompatibility.AccessKeyServiceClient
func (c *AccessKeyServiceClient) ListOperations(ctx context.Context, in *awscompatibility.ListAccessKeyOperationsRequest, opts ...grpc.CallOption) (*awscompatibility.ListAccessKeyOperationsResponse, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return awscompatibility.NewAccessKeyServiceClient(conn).ListOperations(ctx, in, opts...)
}

type AccessKeyOperationsIterator struct {
	ctx  context.Context
	opts []grpc.CallOption

	err           error
	started       bool
	requestedSize int64
	pageSize      int64

	client  *AccessKeyServiceClient
	request *awscompatibility.ListAccessKeyOperationsRequest

	items []*operation.Operation
}

func (c *AccessKeyServiceClient) AccessKeyOperationsIterator(ctx context.Context, req *awscompatibility.ListAccessKeyOperationsRequest, opts ...grpc.CallOption) *AccessKeyOperationsIterator {
	var pageSize int64
	const defaultPageSize = 1000
	pageSize = req.PageSize
	if pageSize == 0 {
		pageSize = defaultPageSize
	}
	return &AccessKeyOperationsIterator{
		ctx:      ctx,
		opts:     opts,
		client:   c,
		request:  req,
		pageSize: pageSize,
	}
}

func (it *AccessKeyOperationsIterator) Next() bool {
	if it.err != nil {
		return false
	}
	if len(it.items) > 1 {
		it.items[0] = nil
		it.items = it.items[1:]
		return true
	}
	it.items = nil // consume last item, if any

	if it.started && it.request.PageToken == "" {
		return false
	}
	it.started = true

	if it.requestedSize == 0 || it.requestedSize > it.pageSize {
		it.request.PageSize = it.pageSize
	} else {
		it.request.PageSize = it.requestedSize
	}

	response, err := it.client.ListOperations(it.ctx, it.request, it.opts...)
	it.err = err
	if err != nil {
		return false
	}

	it.items = response.Operations
	it.request.PageToken = response.NextPageToken
	return len(it.items) > 0
}

func (it *AccessKeyOperationsIterator) Take(size int64) ([]*operation.Operation, error) {
	if it.err != nil {
		return nil, it.err
	}

	if size == 0 {
		size = 1 << 32 // something insanely large
	}
	it.requestedSize = size
	defer func() {
		// reset iterator for future calls.
		it.requestedSize = 0
	}()

	var result []*operation.Operation

	for it.requestedSize > 0 && it.Next() {
		it.requestedSize--
		result = append(result, it.Value())
	}

	if it.err != nil {
		return nil, it.err
	}

	return result, nil
}

func (it *AccessKeyOperationsIterator) TakeAll() ([]*operation.Operation, error) {
	return it.Take(0)
}

func (it *AccessKeyOperationsIterator) Value() *operation.Operation {
	if len(it.items) == 0 {
		panic("calling Value on empty iterator")
	}
	return it.items[0]
}

func (it *AccessKeyOperationsIterator) Error() error {
	return it.err
}

// Update implements awscompatibility.AccessKeyServiceClient
func (c *AccessKeyServiceClient) Update(ctx context.Context, in *awscompatibility.UpdateAccessKeyRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return awscompatibility.NewAccessKeyServiceClient(conn).Update(ctx, in, opts...)
}
