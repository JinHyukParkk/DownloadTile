all: clean build

clean:
	@echo "### Clean result directory"
	@rm -rf ./result/
	@echo "### Clean DownloadTile execFile"
	@rm DownloadTile

build:
	@echo "### Build tile source"
	@go build .
