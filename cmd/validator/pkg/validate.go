package pkg

import (
	"github.com/Nox-OS/NoxOS-AppManagement/codegen"
	"github.com/Nox-OS/NoxOS-AppManagement/common"
	"github.com/Nox-OS/NoxOS-AppManagement/service"
	"github.com/compose-spec/compose-go/loader"
)

func VaildDockerCompose(yaml []byte) (err error) {
	err = nil
	// recover
	defer func() {
		if r := recover(); r != nil {
			err = r.(error)
		}
	}()
	docker, err := service.NewComposeAppFromYAML(yaml, false, false)

	ex, ok := docker.Extensions[common.ComposeExtensionNameXNoxOS]
	if !ok {
		return service.ErrComposeExtensionNameXNoxOSNotFound
	}

	var storeInfo codegen.ComposeAppStoreInfo
	if err = loader.Transform(ex, &storeInfo); err != nil {
		return
	}

	return
}
