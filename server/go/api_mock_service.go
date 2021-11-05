package openapi

import (
	"context"
)

// mock用の新しいservice構造体
type DefaultMockService struct {
}

// mock用サービスのコンストラクタ
func NewDefaultMockService() DefaultApiServicer {
    return &DefaultMockService{}
}

// mock用serviceのメソッド
func (s *DefaultMockService) PetsIdGet(ctx context.Context, id int64) (ImplResponse, error) {
    pet := Pets{
        Id:     id,
        Name:   "doggie",
        Status: "available",
    }
    return Response(200, pet), nil
}
