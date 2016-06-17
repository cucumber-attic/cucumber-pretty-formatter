FEATURES = $(shell find . -name "*.feature")
EVENTS = $(patsubst %.feature,%.json,$(FEATURES))
RESULTS = $(patsubst %.feature,%.out,$(FEATURES))

all: $(RESULTS)

# NOTE: godog needs to be installed: go get github.com/DATA-DOG/godog/cmd/godog
# and $PATH should have $GOPATH/bin exported
%.out: ./%.json ./%.expected cunicorn godog
	mkdir -p $$(dirname $@)
	cat $< | ./cunicorn > $@
	diff --unified $(subst .json,.expected,$<) $@
# .DELETE_ON_ERROR: %.out

%.json: ./%.feature
	godog -f events $(subst .json,.feature,$<) > $@

cunicorn: ../formatter.go
	@command -v go >/dev/null 2>&1 || (echo "go needs to be installed and available in your PATH"; exit 1)
	cd ../cmd/cunicorn && go build -o ../../fixtures/$@

godog:
	@command -v godog >/dev/null 2>&1 || go get github.com/DATA-DOG/godog/cmd/godog

clean:
	rm -rf cunicorn *.out *.json

