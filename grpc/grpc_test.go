package grpc

import (
	"context"
	"fmt"
	"testing"

	"github.com/pkg/errors"
	"github.com/realab/thunder/graphql/schemabuilder"
	"github.com/realab/thunder/grpc/api"
)

type testServer struct {
	authors []*api.Author
}

func (ts *testServer) GetAuthor(_ context.Context, req *api.GetAuthorReq) (*api.Author, error) {
	for _, i := range ts.authors {
		if i.Id == req.Id {
			return i, nil
		}
	}
	return nil, errors.Errorf("Not exist: %d", req.Id)
}

func TestProtoRegister(t *testing.T) {
	server := &testServer{
		authors: []*api.Author{
			{Id: 1, Name: "name_1"},
			{Id: 2, Name: "name_2"},
			{Id: 3, Name: "name_3"},
		},
	}

	builder := schemabuilder.NewSchema()
	builder.Object("author", api.Author{})
	obj := builder.Query()
	obj.FieldFunc("authors", func(ctx context.Context, req api.GetAuthorReq) (*api.Author, error) {
		return server.GetAuthor(ctx, &req)
	})

	schema, err := builder.Build()
	if err != nil {
		t.Logf("Error: %+v\n", err)
	}
	fmt.Printf("AAAAA: %+v\n", schema)
}
