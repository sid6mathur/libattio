# Regenerates the Go package in attio/ from Attio's public OpenAPI spec.
# Needs curl, jq, Java 11+ and Go. `make` runs every step.

SPEC_URL := https://api.attio.com/openapi/api
SPEC     := spec/openapi.json
PKG      := attio
MAVEN    := https://repo1.maven.org/maven2/org/openapitools/openapi-generator-cli

# Defaults to the newest openapi-generator release; pin with GENERATOR_VERSION=7.23.0.
ifndef GENERATOR_VERSION
GENERATOR_VERSION := $(shell curl -fsSL $(MAVEN)/maven-metadata.xml | sed -n 's:.*<release>\(.*\)</release>.*:\1:p')
endif
GENERATOR_JAR := .cache/openapi-generator-cli-$(GENERATOR_VERSION).jar

.PHONY: all spec generate check
all: spec generate check

# Download the spec and apply the fixes in spec/patch.jq.
spec:
	curl -fsSL $(SPEC_URL) | jq -f spec/patch.jq > $(SPEC).tmp
	mv $(SPEC).tmp $(SPEC)

generate: $(GENERATOR_JAR)
	@# Delete the previous output first so schemas removed upstream disappear too.
	cd $(PKG) && sed -n '/\.go$$/p' .openapi-generator/FILES 2>/dev/null | xargs -r rm -f
	java -jar $(GENERATOR_JAR) generate -g go -i $(SPEC) -o $(PKG) \
	  --package-name attio --git-user-id sid6mathur --git-repo-id libattio \
	  --global-property apiDocs=false,modelDocs=false,apiTests=false,modelTests=false \
	  --additional-properties withGoMod=false,hideGenerationTimestamp=true,generateInterfaces=true,disallowAdditionalPropertiesIfNotPresent=false
	go run golang.org/x/tools/cmd/goimports@latest -w $(PKG)
	go mod tidy

check:
	go vet ./...
	go test ./...

$(GENERATOR_JAR):
	mkdir -p $(@D)
	curl -fsSL -o $@ $(MAVEN)/$(GENERATOR_VERSION)/openapi-generator-cli-$(GENERATOR_VERSION).jar
