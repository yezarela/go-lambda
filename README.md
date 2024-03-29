# Go Lambda

Go architecture for AWS Lambda 

## Requirements

- [AWS CLI](https://docs.aws.amazon.com/cli/latest/userguide/cli-chap-install.html) 
- [AWS SAM CLI](https://docs.aws.amazon.com/serverless-application-model/latest/developerguide/serverless-sam-cli-install.html)
- [Docker installed](https://www.docker.com/community-edition)
- [Golang](https://golang.org)

This is the explanation of our file structure:

```bash
.
├── domain                           <-- All domains/entities belong here
├── user                             <-- Module directory
│   └── handler                     
│       └── user-get                 <-- Source code for a lambda function
│           ├── main.go              <-- Lambda function code
│           └── main_test.go         <-- Unit tests
│   └── repository 
│       ├── mysql_repository.go      <-- Repository for our db
│       └── mysql_repository_test.go <-- Unit tests
│   └── usecase 
│       ├── user_usecase.go          <-- Business rules
│       └── user_usecase_test.go     <-- Unit tests
│
├── infra                            <-- All external framework/drivers 
├── utils                            <-- Shared utilities
├── Makefile                         <-- Make to automate build
└── template.yaml                    <-- Cloudformation template
```

## Setup process

### Installing dependencies

```shell
make deps
docker-compose up -d
```

### Building

```shell
make build
```

### Local development

**Invoking function locally through local API Gateway**

```bash
make local
```

If the previous command ran successfully you should now be able to hit the following local endpoint to invoke your function `http://localhost:3000/hello`

## Packaging and deployment

The following command will package your lambda functions, create a Cloudformation Stack and deploy your SAM resources:

```bash
make deploy
```

After deployment is complete you can run the following command to retrieve the API Gateway Endpoint URL:

```bash
make describe
```

### Testing

We use `testing` package that is built-in in Golang and you can simply run the following command to run our tests:

```shell
go test -v ./handlers/hello-world/
```
For mocking purpose, we use `gomock` from Golang to help us to generate mock files quickly.

## Inspiration

- [https://github.com/awslabs/serverless-application-model](https://github.com/awslabs/serverless-application-model)
- [https://github.com/bxcodec/go-clean-arch](https://github.com/bxcodec/go-clean-arch)

## TODO
- Resource tagging
