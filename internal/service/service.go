package service

import (
	"github.com/google/wire"
	"github.com/realotz/mstore/internal/service/storage"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(
	NewCronService,
	NewUserService,
	NewAuthService,
	storage.NewVolumeService,
)
