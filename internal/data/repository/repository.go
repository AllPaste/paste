package repository

import "github.com/google/wire"

// ProviderSet is Repo providers.
var ProviderSet = wire.NewSet(NewPasteRepo)
