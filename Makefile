.PHONY: deps

FEATURE_FILES = $(shell find ./tests -name "*.feature")
EVENT_FILES = $(patsubst %.feature,%.json,$(FEATURE_FILES))
PROGRESS_FMT_FILES = $(patsubst %.feature,%.progress,$(FEATURE_FILES))

all: deps clean cunicorn $(EVENT_FILES) $(PROGRESS_FMT_FILES)
	@echo "all assertions passed.."

# run and assert the event stream through progress format
%.progress: %.json
	@./cunicorn -f progress --no-colors < $< > $@
	@diff --unified $@.expected $@

# use godog to generate the event stream
%.json: %.feature
	@godog -f events $(subst .json,.feature,$@) > $@

# build cunicorn with tag for testing
# this will nulify the resource usage for status summary
cunicorn:
	@cd cmd/cunicorn && go build --tags=testing -o ../../cunicorn

deps:
	@go get github.com/DATA-DOG/godog/cmd/godog

clean:
	@rm -f cunicorn $(PROGRESS_FMT_FILES) $(EVENT_FILES)
