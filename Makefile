ALPACAPLATFORM=alpaca64
ALPACADIR=$(USERPROFILE)\Desktop\alpaca

GOOS=windows
GOARCH=amd64
GOTAGS=$(ALPACAPLATFORM),alpacadebug

.PHONY: buildai build install runai run clean
buildai:
	@set GOOS=$(GOOS)
	@set GOARCH=$(GOARCH)
	@go build -o build/$(ALPACAPLATFORM)/alpaca_ai.exe -tags $(GOTAGS) cmd/alpaca_ai/main.go

build: buildai

install: build
	@if exist "$(ALPACADIR)\bin\$(ALPACAPLATFORM)" rmdir /S /Q "$(ALPACADIR)\bin\$(ALPACAPLATFORM)"
	@xcopy "build\$(ALPACAPLATFORM)" "$(ALPACADIR)\bin\$(ALPACAPLATFORM)\" /E /C /I >nul
	@if exist"$(ALPACADIR)\resources" rmdir /S /Q "$(ALPACADIR)\resources"
	@xcopy "resources" "$(ALPACADIR)\resources\" /E /C /I >nul

runai: buildai
	@if not exist "$(ALPACADIR)\bin\dev" mkdir "$(ALPACADIR)\bin\dev"
	@copy "build\$(ALPACAPLATFORM)\alpaca_ai.exe" "$(ALPACADIR)\bin\dev\alpaca_ai.exe" >nul
	@cd "build\$(ALPACAPLATFORM)" && .\alpaca_ai.exe

run: build
	@if exist "$(ALPACADIR)\bin\dev" rmdir /S /Q "$(ALPACADIR)\bin\dev"
	@xcopy "build\$(ALPACAPLATFORM)" "$(ALPACADIR)\bin\dev\" /E /C /I >nul
	@cd "build\$(ALPACAPLATFORM)" && .\alpaca_ai.exe

clean:
	@if exist "build" rmdir /S /Q build