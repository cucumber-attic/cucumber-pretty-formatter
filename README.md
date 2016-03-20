# Rule them all formatter

As the title says, ze plan is to implement an all platform compatible
utility, which takes specked cucumber event stream from **stdin** and
spits output to **stdout**.

This would enable any cucumber implementation to implement a non
complicated common event stream formatter, which can be streamed to this
tool to get the actual **pretty, junit, html...** format output.

There are many other ideas how this could extend, but their are not clear
yet, like a common test suite for all implementations.

## Status

So far it is only a proof of concept and events have no clear
specification.

    ├── actual.txt    # actual output result
    ├── expected.txt  # expected output
    ├── inp.json      # input of event stream
    ├── main.go       # source
    ├── Makefile      # test case automation
    ├── my.feature    # feature file being tested
    ├── README.md
