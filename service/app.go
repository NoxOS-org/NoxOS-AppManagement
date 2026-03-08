package service

import (
	"github.com/Nox-OS/NoxOS-AppManagement/codegen"
	"github.com/Nox-OS/NoxOS-AppManagement/common"
	"github.com/Nox-OS/NoxOS-Common/utils/logger"
	"github.com/compose-spec/compose-go/loader"
	"github.com/compose-spec/compose-go/types"
)

type App types.ServiceConfig

func (a *App) StoreInfo() (codegen.AppStoreInfo, error) {
	var storeInfo codegen.AppStoreInfo

	ex, ok := a.Extensions[common.ComposeExtensionNameXNoxOS]
	if !ok {
		logger.Error("extension `x-noxos` not found")
		// return storeInfo, ErrComposeExtensionNameXNoxOSNotFound
	}

	// add image to store info for check stable version function.
	storeInfo.Image = a.Image

	if err := loader.Transform(ex, &storeInfo); err != nil {
		return storeInfo, err
	}

	return storeInfo, nil
}
