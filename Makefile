build:
	@go build

install:
	@mkdir -p ~/.loko
	@mkdir -p ~/.loko/data
	@mkdir -p ~/.loko/data/services
	@cp ./data/services/* ~/.loko/data/services