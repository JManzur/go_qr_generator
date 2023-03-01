# go_qr_generator

```bash
go run main.go -content jmanzur.com
```

```bash
GOOS=windows GOARCH=amd64 go build -o qrcode.exe
```

```powershell
New-Item "C:\qrcode" -ItemType Directory
cd C:\qrcode
Invoke-WebRequest https://github.com/JManzur/go_qr_generator/raw/main/qrcode.exe -OutFile qrcode.exe -UseBasicParsing
qrcode.exe -c https://jmanzur.com/
```