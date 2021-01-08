package node

import (
	"context"

	"github.com/dgraph-io/dgo"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/infinimesh/infinimesh/pkg/node/nodepb"
)

type ObjectController struct {
	Dgraph *dgo.Dgraph
	Log    *zap.Logger

	Repo Repo
}

func (s *ObjectController) CreateObject(ctx context.Context, request *nodepb.CreateObjectRequest) (response *nodepb.Object, err error) {
	id, err := s.Repo.CreateObject(ctx, request.GetName(), request.GetParent(), request.GetKind(), request.GetNamespace())
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &nodepb.Object{Uid: id, Name: request.GetName()}, nil
}

func (s *ObjectController) DeleteObject(ctx context.Context, request *nodepb.DeleteObjectRequest) (response *nodepb.DeleteObjectResponse, err error) {
	err = s.Repo.DeleteObject(ctx, request.GetUid())
	if err != nil {
		return nil, err
	}
	return &nodepb.DeleteObjectResponse{}, nil
}

func (s *ObjectController) ListObjects(ctx context.Context, request *nodepb.ListObjectsRequest) (response *nodepb.ListObjectsResponse, err error) {
	objects, err := s.Repo.ListForAccount(ctx, request.Account, request.Namespace, request.Recurse)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &nodepb.ListObjectsResponse{
		Objects: objects,
	}, nil
}
