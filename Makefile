FEATURE_FILES = $(shell find ./testdata -name "*.feature")
EVENT_FILES = $(patsubst %.feature,%.json,$(FEATURE_FILES))
GENERATED_OUTPUT_FILES = $(patsubst ./testdata/%.feature,output/%.out,$(FEATURE_FILES))

<<<<<<< 9ec512631945db297a6c4c6bb470be3685ca757a
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
=======
actual.txt: cufmt inp.json
	cat inp.json | ./cufmt > actual.txt

cufmt: cmd/cufmt/main.go
	cd cmd/cufmt && go build -o ../../cufmt
>>>>>>> start to reorganize code for pretty formatter first
