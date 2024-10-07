default: test

# Run acceptance tests
.PHONY: test
testacc:
	go test ./... -v $(TESTARGS) -timeout 12m

.PHONY: generate
generate:
	mkdir -p openapi
	mkdir -p generated
	cp ../registry/config/openapi/spec.yml openapi/spec.yml
	docker run -v ./generated:/app/output -v ./openapi/spec.yml:/app/openapi.yaml mcr.microsoft.com/openapi/kiota generate --language go -n github.com/registry-tools/rt-sdk/generated
