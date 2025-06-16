# AWS Serverless Go Lambda Template
A boilerplate Serverless Framework  application template for an AWS Lambda function written in **Go**.  It can be used to scaffold your serverless Go Lambda functions.  The Make file is configured with the necessary commands and variables to build a custom binary as well as a cleanup script.  Additionally, a Dockerfile is included to build a custom binary, based on the Amazon Linux, should glibc compatibility issues arrise.

## It Includes:

| File             | Purpose                                            |
| ---------------- | -------------------------------------------------- |
| `main.go`        | Your Lambda function handler (written in Go)       |
| `go.mod`         | Go module setup (using `aws-lambda-go`)            |
| `build-lambda.sh`| Docker build, used if glibc compatibility is issue |
| `Makefile`       | Builds the binary and zips it for deployment       |
| `serverless.yml` | Tells Serverless Framework how to deploy to AWS    |
| `bin/handler.zip`  | The built + zipped binary to be uploaded to Lambda |


## Usage:
Clone the repo and copy the file structure into your project.  Configure the provider section of `serverless.yml` with your AWS credentials and region.

## Build Options

**Default release build:**

`make local-build` or `make docker-build` (use docker build if glibc compatibility issues occur)

**Debug build: Debug build (with no optimizations + symbols)**:


`make local-build BUILD_MODE=debug` or `make docker-build BUILD_MODE=debug`

**Custom binary name (e.g., for multi-function project)**:

`make local-build BINARY_NAME=myhandler` or `make docker-build BINARY_NAME=myhandler`

**Custom GOOS/GOARCH (e.g., for ARM64 Lambda or local test on Mac)**:

`make local-build GOARCH=arm64 or make docker-build GOARCH=arm64`

## Behavior:
Once deployed, this Lambda:

- Is triggered via HTTP (API Gateway)

- Responds to GET requests at /hello

- Returns:

```
 {"message":"Hello from Go Lambda!"}
```
