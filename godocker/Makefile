VERSION     := 1.0.0
COMMIT      := 1234
DATE        := `date +%FT%T%z`
BUILD_FLAGS := "-X=main.appVersion=$(VERSION) -X=main.appCommit=$(COMMIT) -X=main.appBuilt=$(DATE)"

.PHONY: build

build:
	@echo "Compiling for DEV..."
	@go clean && go build -ldflags ${BUILD_FLAGS} -o build/godocker
	@echo "Done!"