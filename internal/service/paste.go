package service

import (
	"context"

	"github.com/AllPaste/paste/pkg/log"

	"github.com/AllPaste/paste/internal/entity"

	"github.com/AllPaste/paste/internal/domain"
	cpb "github.com/AllPaste/sdk/core/v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

type PasteService struct {
	logger *log.Logger
	paste  domain.Paste

	cpb.UnimplementedPasteServiceServer
}

func NewPasteService(paste domain.Paste, logger *log.Logger) *PasteService {
	return &PasteService{
		paste:  paste,
		logger: logger,
	}
}

func (p PasteService) List(
	ctx context.Context,
	in *cpb.ListPasteRequest,
) (*cpb.ListPasteReply, error) {
	// TODO: 分页等参数
	return nil, nil
}

func (p PasteService) Create(
	ctx context.Context,
	in *cpb.CreatePasteRequest,
) (*emptypb.Empty, error) {
	var empty *emptypb.Empty

	if err := p.paste.Insert(ctx, &entity.Paste{
		Name:    in.Name,
		Content: string(in.Content),
	}); err != nil {
		return nil, err
	}

	return empty, nil
}

func (p PasteService) Delete(
	ctx context.Context,
	in *cpb.ByIDPasteRequest,
) (*emptypb.Empty, error) {
	var empty *emptypb.Empty

	if err := p.paste.Remove(ctx, in.Id); err != nil {
		return nil, err
	}
	return empty, nil
}

func (p PasteService) Put(
	ctx context.Context,
	in *cpb.PutPasteRequest,
) (*emptypb.Empty, error) {
	var empty *emptypb.Empty

	if err := p.paste.Update(ctx, &entity.Paste{
		Content: string(in.Content),
	}); err != nil {
	}
	return empty, nil
}

func (p PasteService) Get(
	ctx context.Context,
	in *cpb.ByIDPasteRequest,
) (*cpb.Paste, error) {
	paste, err := p.paste.FetchOne(ctx, in.Id)
	if err != nil {
		return nil, err
	}
	return paste.ToPB(), nil
}
