.PHONY: deps

FEATURE_FILES = $(shell find ./tests -name "*.feature")
EVENT_FILES = $(patsubst %.feature,%.json,$(FEATURE_FILES))
PROGRESS_FMT_FILES = $(patsubst %.feature,%.progress,$(FEATURE_FILES))

all: clean cunicorn $(PROGRESS_FMT_FILES)

%.progress: %.json
	./cunicorn -f progress --no-colors < $< > $@
	diff --unified $@.expected $@

%.json:
	godog -f events $(subst .json,.feature,$@) > $@

cunicorn:
	cd cmd/cunicorn && go build --tags=testing -o ../../cunicorn

deps:
	go get github.com/DATA-DOG/godog/cmd/godog

clean:
	rm -f cunicorn
	rm -f $(PROGRESS_FMT_FILES)
