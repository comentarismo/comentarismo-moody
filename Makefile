all: deps
		@echo "--> Building"
		@godep go build main.go

deps:
		@echo "--> Installing build dependencies"
		@godep save
