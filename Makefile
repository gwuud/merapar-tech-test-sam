.PHONY: test-func test-api build deploy delete

test-func:
	sam invoke start-lambda

test-api:
	sam invoke start-api -p 8080

build:
	sam build

deploy:
	sam deploy

delete:
	sam delete
