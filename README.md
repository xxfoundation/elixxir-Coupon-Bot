# Coupons Bot

## Build commands
```
# Linux 64 bit binary
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags '-w -s' -o coupons.linux64 main.go
# Windows 64 bit binary
GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go build -ldflags '-w -s' -o coupons.win64 main.go
# Windows 32 big binary
GOOS=windows GOARCH=386 CGO_ENABLED=0 go build -ldflags '-w -s' -o coupons.win32 main.go
# Mac OSX 64 bit binary (intel)
GOOS=darwin GOARCH=amd64 CGO_ENABLED=0 go build -ldflags '-w -s' -o coupons.darwin64 main.go
```

## Config

```yaml
# Client information
sessionPath: "/path/to/session"
sessionPass: "session password"
networkFollowerTimeout: 60
ndf: "/path/to/ndf.json"

# QR options
qrSize: 512
qrLevel: 1
qrPath: /cmix/qr.png

# Database connection information
dbUsername: "cmix"
dbPassword: ""
dbName: "cmix_server"
dbAddress: ""
```
