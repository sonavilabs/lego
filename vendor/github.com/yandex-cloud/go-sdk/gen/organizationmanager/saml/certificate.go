// Code generated by sdkgen. DO NOT EDIT.

//nolint
package saml

import (
	"context"

	"google.golang.org/grpc"

	"github.com/yandex-cloud/go-genproto/yandex/cloud/operation"
	saml "github.com/yandex-cloud/go-genproto/yandex/cloud/organizationmanager/v1/saml"
)

//revive:disable

// CertificateServiceClient is a saml.CertificateServiceClient with
// lazy GRPC connection initialization.
type CertificateServiceClient struct {
	getConn func(ctx context.Context) (*grpc.ClientConn, error)
}

// Create implements saml.CertificateServiceClient
func (c *CertificateServiceClient) Create(ctx context.Context, in *saml.CreateCertificateRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return saml.NewCertificateServiceClient(conn).Create(ctx, in, opts...)
}

// Delete implements saml.CertificateServiceClient
func (c *CertificateServiceClient) Delete(ctx context.Context, in *saml.DeleteCertificateRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return saml.NewCertificateServiceClient(conn).Delete(ctx, in, opts...)
}

// Get implements saml.CertificateServiceClient
func (c *CertificateServiceClient) Get(ctx context.Context, in *saml.GetCertificateRequest, opts ...grpc.CallOption) (*saml.Certificate, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return saml.NewCertificateServiceClient(conn).Get(ctx, in, opts...)
}

// List implements saml.CertificateServiceClient
func (c *CertificateServiceClient) List(ctx context.Context, in *saml.ListCertificatesRequest, opts ...grpc.CallOption) (*saml.ListCertificatesResponse, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return saml.NewCertificateServiceClient(conn).List(ctx, in, opts...)
}

type CertificateIterator struct {
	ctx  context.Context
	opts []grpc.CallOption

	err           error
	started       bool
	requestedSize int64
	pageSize      int64

	client  *CertificateServiceClient
	request *saml.ListCertificatesRequest

	items []*saml.Certificate
}

func (c *CertificateServiceClient) CertificateIterator(ctx context.Context, req *saml.ListCertificatesRequest, opts ...grpc.CallOption) *CertificateIterator {
	var pageSize int64
	const defaultPageSize = 1000
	pageSize = req.PageSize
	if pageSize == 0 {
		pageSize = defaultPageSize
	}
	return &CertificateIterator{
		ctx:      ctx,
		opts:     opts,
		client:   c,
		request:  req,
		pageSize: pageSize,
	}
}

func (it *CertificateIterator) Next() bool {
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

	it.items = response.Certificates
	it.request.PageToken = response.NextPageToken
	return len(it.items) > 0
}

func (it *CertificateIterator) Take(size int64) ([]*saml.Certificate, error) {
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

	var result []*saml.Certificate

	for it.requestedSize > 0 && it.Next() {
		it.requestedSize--
		result = append(result, it.Value())
	}

	if it.err != nil {
		return nil, it.err
	}

	return result, nil
}

func (it *CertificateIterator) TakeAll() ([]*saml.Certificate, error) {
	return it.Take(0)
}

func (it *CertificateIterator) Value() *saml.Certificate {
	if len(it.items) == 0 {
		panic("calling Value on empty iterator")
	}
	return it.items[0]
}

func (it *CertificateIterator) Error() error {
	return it.err
}

// ListOperations implements saml.CertificateServiceClient
func (c *CertificateServiceClient) ListOperations(ctx context.Context, in *saml.ListCertificateOperationsRequest, opts ...grpc.CallOption) (*saml.ListCertificateOperationsResponse, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return saml.NewCertificateServiceClient(conn).ListOperations(ctx, in, opts...)
}

type CertificateOperationsIterator struct {
	ctx  context.Context
	opts []grpc.CallOption

	err           error
	started       bool
	requestedSize int64
	pageSize      int64

	client  *CertificateServiceClient
	request *saml.ListCertificateOperationsRequest

	items []*operation.Operation
}

func (c *CertificateServiceClient) CertificateOperationsIterator(ctx context.Context, req *saml.ListCertificateOperationsRequest, opts ...grpc.CallOption) *CertificateOperationsIterator {
	var pageSize int64
	const defaultPageSize = 1000
	pageSize = req.PageSize
	if pageSize == 0 {
		pageSize = defaultPageSize
	}
	return &CertificateOperationsIterator{
		ctx:      ctx,
		opts:     opts,
		client:   c,
		request:  req,
		pageSize: pageSize,
	}
}

func (it *CertificateOperationsIterator) Next() bool {
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

func (it *CertificateOperationsIterator) Take(size int64) ([]*operation.Operation, error) {
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

func (it *CertificateOperationsIterator) TakeAll() ([]*operation.Operation, error) {
	return it.Take(0)
}

func (it *CertificateOperationsIterator) Value() *operation.Operation {
	if len(it.items) == 0 {
		panic("calling Value on empty iterator")
	}
	return it.items[0]
}

func (it *CertificateOperationsIterator) Error() error {
	return it.err
}

// Update implements saml.CertificateServiceClient
func (c *CertificateServiceClient) Update(ctx context.Context, in *saml.UpdateCertificateRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return saml.NewCertificateServiceClient(conn).Update(ctx, in, opts...)
}
