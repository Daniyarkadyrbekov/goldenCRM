.DEFAULT_GOAL=all
PACKAGES_WITH_TESTS:=$(shell go list -f="{{if or .TestGoFiles .XTestGoFiles}}{{.ImportPath}}{{end}}" ./... | grep -v '/vendor/' | grep -v '/builds/')
TEST_TARGETS:=$(foreach p,${PACKAGES_WITH_TESTS},test-$(p))
TEST_OUT_DIR:=testout

CURRENT_DIR := $(shell pwd)
PROJECT := $(subst ${GOPATH}/src/,,$(shell pwd))

.PHONY: all
all: mod test build-local

.PHONY: build-local
build-local: build local

.PHONY: build
build:
	rice embed-go
	go build -o bin/goldenCRM.git -v .

.PHONY: local
local:
	heroku local web

.PHONY: deploy
deploy:
	git push heroku master


.PHONY: mod
mod:
	rm -rf vendor
	GO111MODULE=on go mod download

.PHONY: test
test:
	rm -rf ${TEST_OUT_DIR}
	mkdir -p -m 755 $(TEST_OUT_DIR)
	$(MAKE) -j 1 $(TEST_TARGETS)
	@echo "=== tests: ok ==="

.PHONY: $(TEST_TARGETS)
$(TEST_TARGETS):
	$(eval $@_package := $(subst test-,,$@))
	$(eval $@_filename := $(subst /,_,$($@_package)))

	@echo "== test directory $($@_package) =="
	@go test $($@_package) -v -race -coverprofile $(TEST_OUT_DIR)/$($@_filename)_cover.out \
    >> $(TEST_OUT_DIR)/$($@_filename).out \
   || ( echo 'fail $($@_package)' && cat $(TEST_OUT_DIR)/$($@_filename).out; exit 1);

.PHONY: static
static:
	docker run -it --rm \
	-v "$(shell pwd):/go/src/${PROJECT}" \
	-v "${GOPATH}/pkg:/go/pkg" \
	-w "/go/src/${PROJECT}" \
	-e "GOPRIVATE=${GOPRIVATE}" \
	-e "GOFLAGS=" \
	dialogs/go-tools-embedded:1.0.1 \
	sh -c '\
	(cd lib/storage/postgres/migrations/assets/; \
	 rm -fv ../migrations.go; \
	 go-bindata -pkg migrations -nomemcopy -o ../migrations.go *.sql)'