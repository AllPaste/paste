// Code generated by ent, DO NOT EDIT.

package paste

import (
	"entgo.io/ent/dialect/sql"
	"github.com/AllPaste/paste/internal/data/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Paste {
	return predicate.Paste(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Paste {
	return predicate.Paste(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Paste {
	return predicate.Paste(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Paste {
	return predicate.Paste(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Paste {
	return predicate.Paste(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Paste {
	return predicate.Paste(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Paste {
	return predicate.Paste(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Paste {
	return predicate.Paste(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Paste {
	return predicate.Paste(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Content applies equality check predicate on the "content" field. It's identical to ContentEQ.
func Content(v string) predicate.Paste {
	return predicate.Paste(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldContent), v))
	})
}

// Creator applies equality check predicate on the "creator" field. It's identical to CreatorEQ.
func Creator(v string) predicate.Paste {
	return predicate.Paste(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreator), v))
	})
}

// ContentEQ applies the EQ predicate on the "content" field.
func ContentEQ(v string) predicate.Paste {
	return predicate.Paste(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldContent), v))
	})
}

// ContentNEQ applies the NEQ predicate on the "content" field.
func ContentNEQ(v string) predicate.Paste {
	return predicate.Paste(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldContent), v))
	})
}

// ContentIn applies the In predicate on the "content" field.
func ContentIn(vs ...string) predicate.Paste {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Paste(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldContent), v...))
	})
}

// ContentNotIn applies the NotIn predicate on the "content" field.
func ContentNotIn(vs ...string) predicate.Paste {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Paste(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldContent), v...))
	})
}

// ContentGT applies the GT predicate on the "content" field.
func ContentGT(v string) predicate.Paste {
	return predicate.Paste(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldContent), v))
	})
}

// ContentGTE applies the GTE predicate on the "content" field.
func ContentGTE(v string) predicate.Paste {
	return predicate.Paste(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldContent), v))
	})
}

// ContentLT applies the LT predicate on the "content" field.
func ContentLT(v string) predicate.Paste {
	return predicate.Paste(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldContent), v))
	})
}

// ContentLTE applies the LTE predicate on the "content" field.
func ContentLTE(v string) predicate.Paste {
	return predicate.Paste(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldContent), v))
	})
}

// ContentContains applies the Contains predicate on the "content" field.
func ContentContains(v string) predicate.Paste {
	return predicate.Paste(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldContent), v))
	})
}

// ContentHasPrefix applies the HasPrefix predicate on the "content" field.
func ContentHasPrefix(v string) predicate.Paste {
	return predicate.Paste(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldContent), v))
	})
}

// ContentHasSuffix applies the HasSuffix predicate on the "content" field.
func ContentHasSuffix(v string) predicate.Paste {
	return predicate.Paste(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldContent), v))
	})
}

// ContentEqualFold applies the EqualFold predicate on the "content" field.
func ContentEqualFold(v string) predicate.Paste {
	return predicate.Paste(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldContent), v))
	})
}

// ContentContainsFold applies the ContainsFold predicate on the "content" field.
func ContentContainsFold(v string) predicate.Paste {
	return predicate.Paste(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldContent), v))
	})
}

// CreatorEQ applies the EQ predicate on the "creator" field.
func CreatorEQ(v string) predicate.Paste {
	return predicate.Paste(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreator), v))
	})
}

// CreatorNEQ applies the NEQ predicate on the "creator" field.
func CreatorNEQ(v string) predicate.Paste {
	return predicate.Paste(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreator), v))
	})
}

// CreatorIn applies the In predicate on the "creator" field.
func CreatorIn(vs ...string) predicate.Paste {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Paste(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCreator), v...))
	})
}

// CreatorNotIn applies the NotIn predicate on the "creator" field.
func CreatorNotIn(vs ...string) predicate.Paste {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Paste(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCreator), v...))
	})
}

// CreatorGT applies the GT predicate on the "creator" field.
func CreatorGT(v string) predicate.Paste {
	return predicate.Paste(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreator), v))
	})
}

// CreatorGTE applies the GTE predicate on the "creator" field.
func CreatorGTE(v string) predicate.Paste {
	return predicate.Paste(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreator), v))
	})
}

// CreatorLT applies the LT predicate on the "creator" field.
func CreatorLT(v string) predicate.Paste {
	return predicate.Paste(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreator), v))
	})
}

// CreatorLTE applies the LTE predicate on the "creator" field.
func CreatorLTE(v string) predicate.Paste {
	return predicate.Paste(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreator), v))
	})
}

// CreatorContains applies the Contains predicate on the "creator" field.
func CreatorContains(v string) predicate.Paste {
	return predicate.Paste(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldCreator), v))
	})
}

// CreatorHasPrefix applies the HasPrefix predicate on the "creator" field.
func CreatorHasPrefix(v string) predicate.Paste {
	return predicate.Paste(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldCreator), v))
	})
}

// CreatorHasSuffix applies the HasSuffix predicate on the "creator" field.
func CreatorHasSuffix(v string) predicate.Paste {
	return predicate.Paste(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldCreator), v))
	})
}

// CreatorEqualFold applies the EqualFold predicate on the "creator" field.
func CreatorEqualFold(v string) predicate.Paste {
	return predicate.Paste(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldCreator), v))
	})
}

// CreatorContainsFold applies the ContainsFold predicate on the "creator" field.
func CreatorContainsFold(v string) predicate.Paste {
	return predicate.Paste(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldCreator), v))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Paste) predicate.Paste {
	return predicate.Paste(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Paste) predicate.Paste {
	return predicate.Paste(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Paste) predicate.Paste {
	return predicate.Paste(func(s *sql.Selector) {
		p(s.Not())
	})
}
