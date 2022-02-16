THIS_DIR := $(dir $(abspath $(firstword $(MAKEFILE_LIST))))

install-cli-tool:
		@echo "========================== Installing Fare Estimation CLI tool ==============================="
		go build
		go install github.com/iliaskaras/fare-estimation

run-tests:
		go test -count=1 ${THIS_DIR}/app/distances/
		go test -count=1 ${THIS_DIR}/app/fares/
		go test -count=1 ${THIS_DIR}/app/files/
		go test -count=1 ${THIS_DIR}/app/rides/
