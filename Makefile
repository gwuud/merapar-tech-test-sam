.PHONY: init test-func test-api build deploy delete

init:
	sam init --runtime go1.x

test-func:
	sam invoke start-lambda

test-api:
	sam invoke start-api -p 8080

build:
	sam build -t template.yaml

deploy:
	sam deploy

delete:
	sam delete
