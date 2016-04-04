FEATURE_FILES = $(shell find ./testdata -name "*.feature")
EVENT_FILES = $(patsubst %.feature,%.json,$(FEATURE_FILES))
GENERATED_OUTPUT_FILES = $(patsubst ./testdata/%.feature,output/%.out,$(FEATURE_FILES))

all: $(GENERATED_OUTPUT_FILES)

output/%.out: ./testdata/%.json ./testdata/%.expected cucumber-pretty fake-cucumber
	mkdir -p $$(dirname $@)
	./fake-cucumber $(subst .json,.feature,$<) | ./cucumber-pretty > $@
	diff --unified $(subst .json,.expected,$<) $@
.DELETE_ON_ERROR: output/%.out

cucumber-pretty: cucumber-pretty.go
	go build -o $@

clean:
	rm -rf cucumber-pretty output
