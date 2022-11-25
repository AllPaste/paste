package repository

import (
	"context"
	"log"

	"github.com/AllPaste/paste/internal/data/ent/paste"

	"github.com/AllPaste/paste/internal/data/ent"
	"github.com/AllPaste/paste/internal/domain"
	"github.com/AllPaste/paste/internal/entity"
)

type PasteRepo struct {
	db     *ent.Client
	logger *log.Logger
}

// NewPasteRepo paste CURD Client
func NewPasteRepo(client *ent.Client, logger *log.Logger) domain.Paste {
	return &PasteRepo{
		db:     client,
		logger: logger,
	}
}

func (p *PasteRepo) All(ctx context.Context) ([]*entity.Paste, error) {
	pastes, err := p.db.Paste.Query().All(ctx)
	if err != nil {
		return nil, err
	}

	var entities []*entity.Paste
	for _, p := range pastes {
		entities = append(entities, entity.NewPasteWithPO(p))
	}
	return entities, nil
}

func (p *PasteRepo) Remove(ctx context.Context, id int64) error {
	if err := p.db.Paste.DeleteOneID(id).Exec(ctx); err != nil {
		return err
	}
	return nil
}

func (p *PasteRepo) Update(ctx context.Context, entityPaste *entity.Paste) error {
	if err := p.db.Paste.Update().
		Where(paste.ID(entityPaste.ID)).
		SetContent(entityPaste.Content).
		Exec(ctx); err != nil {
		return err
	}
	return nil
}

func (p *PasteRepo) FetchOne(ctx context.Context, id int64) (*entity.Paste, error) {
	pst, err := p.db.Paste.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return entity.NewPasteWithPO(pst), nil
}

func (p *PasteRepo) Insert(ctx context.Context, paste *entity.Paste) error {
	if err := p.db.Paste.Create().
		SetCreator(paste.Name).
		SetContent(paste.Content).
		Exec(ctx); err != nil {
		return err
	}
	return nil
}
