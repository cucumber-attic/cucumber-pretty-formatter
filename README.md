# Rule them all formatter

As the title says, ze plan is to implement an all platform compatible
utility, which takes specked cucumber event stream from **stdin** and
spits output to **stdout**.

This would enable any cucumber implementation to implement a non
complicated common event stream formatter, which can be streamed to this
tool to get the actual **pretty, junit...** format output.

There are many other ideas how this could extend, but their are not clear
yet, like a common test suite for all implementations.

Formatter error output should be **stderr**.

## Status

So far it is only a proof of concept and events have no clear
specification.

Currently implemented formatters:

- progress

## Integration

The formatter will be shipped as a binary for all OS architectures
separately. A cucumber implementation could advice user to install it with
a single one liner command so it is available in **`PATH`** or install it
together with a cucumber implementation.

If cucumber implementation see this binary in **`PATH`** then it could
stream events through it and output specific format to the stdout or other
output stream.

## Events

In order to be able to print results in all different kind of formats.
General purpose formatter, expects these events passed in as a stream of
**json** objects.

1. [`TestRunStarted`](#testrunstarted)
2. [`TestSource`](#testsource)
3. [`StepDefinitionFound`](#stepdefinitionfound)
4. [`TestCaseStarted`](#testcasestarted)
5. [`TestStepStarted`](#teststepstarted)
6. [`TestStepFinished`](#teststepfinished)
7. [`TestCaseFinished`](#testcasefinished)
8. [`TestRunFinished`](#testrunfinished)
9. [`TestAttachment`](#testattachment)


### `TestRunStarted`

Triggers when tests are started. Specifies protocol version.

``` json
{
    "event": "TestRunStarted",
    "version": "0.1.0",
    "timestamp": 1461436176456,
    "suite": "main"
}
```

1. **`event`** - name of event.
2. **`version`** - (optional) protocol version used for events. If not
   provided, latest stable protocol version is expected.
3. **`timestamp`** - unix timestamp in milliseconds since epoch. When the
   test run started.
4. **`suite`** - (optional) name of the test suite.

### `TestSource`

When a test source is parsed, this event should be sent with plain text of
source. It will be determined by source extension found in **`location`**.

Currently only **gherkin** source is supported. 

``` json
{
    "event": "TestSource",
    "location": "features/simple.feature:1",
    "source": "Feature:\n  Scenario: passing\n    Given passes"
}
```

1. **`event`** - name of event.
2. **`location`** - location in source file, based on pattern `{path}:{line}`.
3. **`source`** - is plain text of test source.

### `StepDefinitionFound`

Should fire when step regexp or other matching algorithm determines step
implementation in the source code. 

**Note:** There may be ambiguous matches.


``` json
{
    "event": "StepDefinitionFound",
    "location": "features/simple.feature:5",
    "definition": "FeatureContext::passing():6",
    "arguments": [
        [12, 18],
        [23, 26]
    ]
}
```

1. **`event`** - name of event.
2. **`location`** - location in source file, based on pattern `{path}:{line}`.
3. **`definition`** - reference to step definition in source code.
4. **`arguments`** - list of positions for arguments which were matched.
   Positions are determined on **step text** step keyword should be
   omitted when calculating argument positions.

### `TestCaseStarted`

Should fire when starting to execute scenario or scenario outline example.

``` json
{
    "event": "TestCaseStarted",
    "location": "features/simple.feature:4",
    "timestamp": 1461436176456
}
```

1. **`event`** - name of event.
2. **`location`** - location in source file, based on pattern `{path}:{line}`.
3. **`timestamp`** - unix timestamp in milliseconds since epoch (when the
   test case started).

### `TestStepStarted`

Should fire right before step execution.

``` json
{
    "event": "TestStepStarted",
    "location": "features/simple.feature:5",
    "timestamp": 1461436176456
}
```

1. **`event`** - name of event.
2. **`location`** - location in source file, based on pattern `{path}:{line}`.
3. **`timestamp`** - unix timestamp in milliseconds since epoch (when the
   test step started).

### `TestStepFinished`

Should fire right after step has finished execution and give appropriate
status and details.

``` json
{
    "event": "TestStepFinished",
    "location": "features/simple.feature:5",
    "status": "failed",
    "timestamp": 1461436176456,
    "summary": "error - user was not found by id: 1",
    "details": "error details\ndebug information"
}
```

1. **`event`** - name of event.
2. **`location`** - location in source file, based on pattern `{path}:{line}`.
3. **`status`** - can be one of `passed`, `failed`, `skipped`, `pending`,
   `undefined`, `ambiguous`.
4. **`timestamp`** - unix timestamp in milliseconds since epoch (when the
   test step finished).
5. **`summary`** - (optional) one line summary for step result.
6. **`details`** - (optional) multi-line detailed description of step
   result.

### `TestCaseFinished`

Should fire after all steps are executed for scenario or outline example.
Should provide appropriate result status.

``` json
{
    "event": "TestCaseFinished",
    "location": "features/simple.feature:5",
    "status": "failed",
    "timestamp": 1461436176456
}
```

1. **`event`** - name of event.
2. **`location`** - location in source file, based on pattern `{path}:{line}`.
3. **`status`** - can be one of `passed`, `failed`, `skipped`, `pending`,
   `undefined`, `ambiguous`.
4. **`timestamp`** - unix timestamp in milliseconds since epoch (when the
   test case finished).

### `TestRunFinished`

Should fire after all tests are run, **or** if “stop on failure” flag was
specified and failure occurred. It should carry the status of all tests
and resource usage summary information.

``` json
{
    "event": "TestRunFinished",
    "status": "failed",
    "timestamp": 1461436176456,
    "memory": 3456765,
    "snippets": "implement undefined steps with the following snippets:"
}
```

1. **`event`** - name of event.
2. **`status`** - can be one of `passed`, `failed`, `skipped`, `pending`,
   `undefined`, `ambiguous`.
3. **`timestamp`** - unix timestamp in milliseconds since epoch (when the
   test run finished).
4. **`memory`** - (optional) memory consumption in bytes used by all tests.
5. **`snippets`** - (optional) undefined step implementation source code
   snippets.

### `TestAttachment`

An attachment to test cases, for example a screenshot or video. Might be
exception stack traces.

``` json
{
    "event": "TestAttachment",
    "location": "features/simple.feature:5",
    "mime": "image/png",
    "data": "YWJjZGU=",
    "encoding": "base64",
    "timestamp": 1461436176456
}
```

1. **`event`** - name of event.
2. **`location`** - location in source file, based on pattern `{path}:{line}`.
3. **`mime`** - mime type of given media file.
4. **`data`** - encoded data.
5. **`encoding`** - data must be encoded to transfer with json format,
   usually base64 or base85.
6. **`timestamp`** - unix timestamp in milliseconds since epoch (when the
   attachment was created).
