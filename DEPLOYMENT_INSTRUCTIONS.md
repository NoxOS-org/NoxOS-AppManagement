# CasaOS-AppManagement Binary Build & Deployment

## What Was Done

The original `IceWhaleTech/CasaOS-AppManagement` code already uses Docker API version negotiation correctly (`client.WithAPIVersionNegotiation()`), so **no code changes were needed**. The issue was likely that your deployed binary was compiled with an older version.

## Build Details

- **Binary**: `dist/casaos-app-management`
- **Architecture**: Linux amd64
- **Size**: 64 MB
- **Build Date**: 2026-02-08
- **MD5**: `2bca1a766dac1504728d5d5ca91a312d`
- **Build Flags**: `-ldflags="-s -w"` (stripped and optimized)
- **Docker SDK Version**: v24.0.7
- **API Negotiation**: ✅ Enabled (will negotiate with Docker daemon automatically)

## Deployment Instructions

### 1. Back Up Current Binary

On your server:
```bash
sudo cp /usr/bin/casaos-app-management /usr/bin/casaos-app-management.bak
```

### 2. Transfer New Binary

From your development machine (where you built the binary):
```bash
# Option A: Using SCP
scp /config/workspace/casaos101/CasaOS-AppManagement/dist/casaos-app-management user@your-server:/tmp/

# Option B: Using rsync
rsync -avz /config/workspace/casaos101/CasaOS-AppManagement/dist/casaos-app-management user@your-server:/tmp/
```

### 3. Replace Binary on Server

On your server:
```bash
# Stop the service
sudo systemctl stop casaos-app-management

# Replace the binary
sudo cp /tmp/casaos-app-management /usr/bin/casaos-app-management

# Ensure correct permissions
sudo chmod +x /usr/bin/casaos-app-management
sudo chown root:root /usr/bin/casaos-app-management

# Start the service
sudo systemctl start casaos-app-management
```

### 4. Verify

```bash
# Check service status
sudo systemctl status casaos-app-management

# Check logs for API version errors (should be gone)
sudo journalctl -u casaos-app-management -f --since "1 minute ago"
```

You should **no longer see** the error:
```
client version 1.43 is too old. Minimum supported API version is 1.44
```

## Rollback Instructions

If something goes wrong:
```bash
sudo systemctl stop casaos-app-management
sudo cp /usr/bin/casaos-app-management.bak /usr/bin/casaos-app-management
sudo systemctl start casaos-app-management
```

## Technical Details

### Why This Fixes the Issue

The original code uses `client.WithAPIVersionNegotiation()` which tells the Docker client to:
1. Query the Docker daemon for its supported API version
2. Use the highest mutually supported version
3. This allows the client to work with both old and new Docker daemons

Your old binary likely had a hardcoded fallback to API 1.43, but this freshly built binary will properly negotiate with your Docker 29.x daemon (which requires API 1.44+).

### Verification Commands

After deployment, you can verify the API negotiation is working:
```bash
# Check Docker daemon API version
docker version --format '{{.Server.APIVersion}}'

# Check what version the app-management is using (in logs)
sudo journalctl -u casaos-app-management -n 100 | grep -i "api\|version"
```

## Notes

- DNS issues on your server won't affect the binary since all dependencies are statically compiled
- The binary is CGO-disabled, so it has no external C library dependencies
- The build uses Go 1.21 (original project requirement)

