EVENT_FILES = $(shell find ./testdata -name "*.json")
GENERATED_OUTPUT_FILES = $(patsubst ./testdata/%.json,output/%.out,$(EVENT_FILES))

all: $(GENERATED_OUTPUT_FILES)

output/%.out: ./testdata/%.json ./testdata/%.json.expected streamer
	cat $< | ./streamer > $@
	diff --unified $<.expected $@
.DELETE_ON_ERROR: output/%.out

streamer: main.go
	go build -o streamer

clean:
	rm -rf output/*
