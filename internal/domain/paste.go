package domain

import (
	"context"

	"github.com/AllPaste/paste/internal/entity"
)

type Paste interface {
	All(ctx context.Context) ([]*entity.Paste, error)
	Remove(ctx context.Context, id int64) error
	Update(ctx context.Context, paste *entity.Paste) error
	FetchOne(ctx context.Context, id int64) (*entity.Paste, error)
	Insert(ctx context.Context, paste *entity.Paste) error
}
