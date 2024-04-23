build:
	GOOS=linux GOARCH=arm64 go build -tags lambda.norpc -o bin/consent_creation/bootstrap handlers/consent_creation/main.go
	GOOS=linux GOARCH=arm64 go build -tags lambda.norpc -o bin/data_sharing/bootstrap handlers/data_sharing/main.go
	zip bin/consent_creation.zip bin/consent_creation/bootstrap
	zip bin/data_sharing.zip bin/data_sharing/bootstrap

clean:
	rm -rf bin

deploy: clean build
	serverless deploy