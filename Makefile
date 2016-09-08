.PHONY: deps

FEATURE_FILES = $(shell find ./tests -name "*.feature")
EVENT_FILES = $(patsubst %.feature,%.json,$(FEATURE_FILES))
PROGRESS_FMT_FILES = $(patsubst %.feature,%.progress,$(FEATURE_FILES))

all: clean deps $(PROGRESS_FMT_FILES)

%.progress: %.json
	cunicorn -f progress --no-colors < $< > $@

%.json:
	godog -f events $(subst .json,.feature,$@) > $@

deps:
	go get github.com/DATA-DOG/godog/cmd/godog
	go install ./...

clean:
	rm -f $(PROGRESS_FMT_FILES)
	rm -f $(EVENT_FILES)
