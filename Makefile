EVENT_FILES = $(shell find ./testdata -name "*.json")
GENERATED_OUTPUT_FILES = $(patsubst ./testdata/%.json,output/%.out,$(EVENT_FILES))

all: $(GENERATED_OUTPUT_FILES)

output/%.out: ./testdata/%.json ./testdata/%.json.expected cucumber-pretty
	cat $< | ./cucumber-pretty > $@
	diff --unified $<.expected $@
.DELETE_ON_ERROR: output/%.out

cucumber-pretty: cucumber-pretty.go
	go build -o $@

clean:
	rm -rf cucumber-pretty output/*
