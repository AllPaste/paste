package entity

import (
	"github.com/AllPaste/paste/internal/data/ent"
	cpb "github.com/AllPaste/sdk/core/v1"
)

type Paste struct {
	ID        int64  `json:"id,omitempty"`
	CreatedAt int64  `json:"createdAt,omitempty"`
	UpdatedAt int64  `json:"updatedAt,omitempty"`
	Name      string `json:"name,omitempty"`
	Content   string `json:"content,omitempty"`
}

func NewPasteWithPO(po *ent.Paste) *Paste {
	if po == nil {
		return nil
	}
	return &Paste{
		ID:        po.ID,
		CreatedAt: po.CreatedAt,
		UpdatedAt: po.UpdatedAt,
		Name:      po.Creator,
		Content:   po.Content,
	}
}

func (p *Paste) ToPB() *cpb.Paste {
	if p == nil {
		return nil
	}

	return &cpb.Paste{
		Id:        p.ID,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
		Name:      p.Name,
		Content:   []byte(p.Content),
	}
}
