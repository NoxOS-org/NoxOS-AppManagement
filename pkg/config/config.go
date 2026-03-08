package config

import (
	"path/filepath"

	"github.com/Nox-OS/NoxOS-Common/utils/constants"
)

var (
	AppManagementConfigFilePath    = filepath.Join(constants.DefaultConfigPath, "app-management.conf")
	AppManagementGlobalEnvFilePath = filepath.Join(constants.DefaultConfigPath, "env")
	RemoveRuntimeIfNoNvidiaGPUFlag = false
)
