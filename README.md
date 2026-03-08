# NoxOS-AppManagement

[![Go Reference](https://pkg.go.dev/badge/github.com/Nox-OS/NoxOS-AppManagement.svg)](https://pkg.go.dev/github.com/Nox-OS/NoxOS-AppManagement)
[![Go Report Card](https://goreportcard.com/badge/github.com/Nox-OS/NoxOS-AppManagement)](https://goreportcard.com/report/github.com/Nox-OS/NoxOS-AppManagement)
[![goreleaser](https://github.com/Nox-OS/NoxOS-AppManagement/actions/workflows/release.yml/badge.svg)](https://github.com/Nox-OS/NoxOS-AppManagement/actions/workflows/release.yml)
[![codecov](https://codecov.io/gh/Nox-OS/NoxOS-AppManagement/branch/main/graph/badge.svg?token=ZCWZOFKXJT)](https://codecov.io/gh/Nox-OS/NoxOS-AppManagement)
[![Vulnerabilities](https://sonarcloud.io/api/project_badges/measure?project=Nox-OS_NoxOS-AppManagement&metric=vulnerabilities)](https://sonarcloud.io/summary/new_code?id=Nox-OS_NoxOS-AppManagement)
[![Bugs](https://sonarcloud.io/api/project_badges/measure?project=Nox-OS_NoxOS-AppManagement&metric=bugs)](https://sonarcloud.io/summary/new_code?id=Nox-OS_NoxOS-AppManagement)
[![Code Smells](https://sonarcloud.io/api/project_badges/measure?project=Nox-OS_NoxOS-AppManagement&metric=code_smells)](https://sonarcloud.io/summary/new_code?id=Nox-OS_NoxOS-AppManagement)
[![Lines of Code](https://sonarcloud.io/api/project_badges/measure?project=Nox-OS_NoxOS-AppManagement&metric=ncloc)](https://sonarcloud.io/summary/new_code?id=Nox-OS_NoxOS-AppManagement)
[![Duplicated Lines (%)](https://sonarcloud.io/api/project_badges/measure?project=Nox-OS_NoxOS-AppManagement&metric=duplicated_lines_density)](https://sonarcloud.io/summary/new_code?id=Nox-OS_NoxOS-AppManagement)

App Management service handles the full lifecycle of NoxOS applications — installation, updates, removal, and runtime state.

## Overview

NoxOS-AppManagement orchestrates Docker Compose-based apps from the NoxOS AppStore and user-defined custom apps. It provides the backend for app operations triggered through the UI or CLI.

**Key responsibilities:**

- Install, start, stop, restart, and remove apps
- Pull and sync app definitions from the NoxOS AppStore
- Track app status and resource usage
- Register app routes with NoxOS-Gateway
- Emit lifecycle events to NoxOS-MessageBus

## API

API specification is available in the [`api/`](./api/) directory.

## Development

### Prerequisites

- Go 1.21+
- Docker and Docker Compose

### Build

```bash
go build ./...
```

### Test

```bash
go test ./...
```

### Run

```bash
go run main.go
```

## Contributing

Issues and pull requests are welcome.

- [NoxOS GitHub](https://github.com/Nox-OS)
- [Open an issue](https://github.com/Nox-OS/NoxOS-AppManagement/issues)

## License

[Apache 2.0](LICENSE)
