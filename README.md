# KDS Consumer On AWS Lambda

### Native build for AWS Lambda

https://docs.aws.amazon.com/ja_jp/lambda/latest/dg/golang-package.html

```sh
$ GOOS=linux GOARCH=arm64 go build -tags lambda.norpc -o bootstrap main.go \
  && chmod +x bootstrap \
  && zip lambda.zip bootstrap \
  && rm bootstrap
```
