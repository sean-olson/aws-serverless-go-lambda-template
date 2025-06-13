# aws-serverless-go-lambda-template
A boilerplate Serverless Framework  application template for an AWS Lambda function

### It Includes:

| File             | Purpose                                            |
| ---------------- | -------------------------------------------------- |
| `main.go`        | Your Lambda function handler (written in Go)       |
| `go.mod`         | Go module setup (using `aws-lambda-go`)            |
| `Makefile`       | Builds the binary and zips it for deployment       |
| `serverless.yml` | Tells Serverless Framework how to deploy to AWS    |
| `bin/hello.zip`  | The built + zipped binary to be uploaded to Lambda |


### Usage:
clone the repo and copy the file structure into your project.  Configure the provider section of `serverless.yml` with your AWS credentials and region.

### Behavior:
Once deployed, this Lambda:

- Is triggered via HTTP (API Gateway)

- Responds to GET requests at /hello

- Returns:

```
 Hello, /hello!
```
