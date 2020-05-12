.PHONY: deps clean build

# Deployment config
STAGE=dev
PROJECT_NAME=myproject
PROFILE=${PROJECT_NAME}_${STAGE}
BUCKET_NAME=cfn.${PROJECT_NAME}.${STAGE}
PARAMETERS=`cat env.${STAGE}`

# Stack config
STACK_NAME=${PROJECT_NAME}-${STAGE}
TEMPLATE_NAME=template.yaml

handlers := $(shell find handler -name \*main.go | awk -F'/' '{print $$2}')

deps:
	@echo "\nInstalling dependencies"
	go get ./...

clean:
	@echo "\nRemoving old builds"
	rm -rf bin

test:
	@echo "\nRunning unit tests"
	go test -short  ./...

local: 
	@echo "\nServing locally"
	env ${PARAMETERS} sam local start-api

build: 
	@echo "\nBuilding handlers"
	@for handler in $(handlers) ; do \
		GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o ./bin/$$handler/$$handler ./handler/$$handler || exit 1; \
	done

deploy: 
	@echo "\nPackaging AWS SAM Application"
	sam package \
		--template-file ${TEMPLATE_NAME} \
		--s3-bucket ${BUCKET_NAME} \
		--output-template-file packaged-${TEMPLATE_NAME} \
		--profile ${PROFILE}
	
	@echo "\nDeploying AWS SAM Application"
	sam deploy \
		--template-file packaged-${TEMPLATE_NAME} \
		--stack-name ${STACKNAME} \
		--capabilities CAPABILITY_NAMED_IAM \
		--profile ${PROFILE} \
		--parameter-overrides ${PARAMETERS}

describe:
	@echo "\nDescribe stack"
	aws cloudformation describe-stacks \
		--stack-name ${STACK_NAME} \
		--profile ${PROFILE} \
		--query 'Stacks[].Outputs[]'

mocks:
	@echo "\nGenerating mocks"
	mockgen -source=domain/user/repository.go -destination=domain/user/mock/repository_mock.go # -mock_names IRepository=MockRepository
	mockgen -source=domain/user/usecase.go -destination=domain/user/mock/usecase_mock.go # -mock_names IUsecase=MockUsecase

publish: clean build deploy

$(V).SILENT: