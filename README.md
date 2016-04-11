# Rule them all formatter

As the title says, ze plan is to implement an all platform compatible
utility, which takes specked cucumber event stream from **stdin** and
spits output to **stdout**.

This would enable any cucumber implementation to implement a non
complicated common event stream formatter, which can be streamed to this
tool to get the actual **pretty, junit, html...** format output.

There are many other ideas how this could extend, but their are not clear
yet, like a common test suite for all implementations.

Formatter error output should be **stderr**

## Status

So far it is only a proof of concept and events have no clear
specification.

## Events

In order to be able to print results in all different kind of formats.
General purpose formatter, expects these events passed in as a stream of
**json** objects.

1. [TestRunStarted](#testrunstarted)
2. [TestSource](#testsource)
3. [StepDefinitionFound](#stepdefinitionfound)
4. [TestCaseStarted](#testcasestarted)
5. [TestStepStarted](#teststepstarted)
6. [TestStepFinished](#teststepfinished)
7. [TestCaseFinished](#testcasefinished)
8. [TestRunFinished](#testrunfinished)


### TestRunStarted

Triggers when tests are started. Specifies protocol version.

``` json
{
    "event": "TestRunStarted",
    "version": "0.1.0"
}
```

1. **event** - name of event.
2. **version** - `optional` protocol version used for events. If not
   provided, latest stable protocol version is expected.

### TestSource

When a test source is parsed, this event should be sent with plain text of
source. Currently only **gherkin** source is supported and it will be
determined by source extension found in **location**.

``` json
{
    "event": "TestSource",
    "location": "features/simple.feature:1",
    "source": "Feature:\n  Scenario: passing\n    Given passes"
}
```

1. **event** - name of event.
2. **location** - location in source file, based on pattern {path}:{line}.
3. **source** - is plain text of test source.

### StepDefinitionFound

Should fire when step regexp or other matching algorithm determines step
implementation in the source code. Note: there may be ambiguous matches.


``` json
{
    "event": "StepDefinitionFound",
    "location": "features/simple.feature:5",
    "definition_id": "FeatureContext::passing():6",
    "arguments": [
        [12, 18],
        [23, 26]
    ],
    "suite": "main"
}
```

1. **event** - name of event.
2. **location** - location in source file, based on pattern {path}:{line}.
3. **definition_id** - reference to step definition in source code.
4. **arguments** - list of positions for arguments which were matched.
   Positions are determined on **step text** step keyword should be
   omitted when calculating argument positions.
5. **suite** - `optional` may be used to distinguish test groups.

### TestCaseStarted

Should fire when starting to execute scenario or scenario outline example.

``` json
{
    "event": "TestCaseStarted",
    "location": "features/simple.feature:4",
    "suite": "main"
}
```

1. **event** - name of event.
2. **location** - location in source file, based on pattern {path}:{line}.
3. **suite** - `optional` may be used to distinguish test groups.

### TestStepStarted

Should fire right before step execution.

``` json
{
    "event": "TestStepStarted",
    "location": "features/simple.feature:5",
    "suite": "main"
}
```

1. **event** - name of event.
2. **location** - location in source file, based on pattern {path}:{line}.
3. **suite** - `optional` may be used to distinguish test groups.

### TestStepFinished

Should fire right after step has finished execution and give appropriate
status and details.

``` json
{
    "event": "TestStepFinished",
    "location": "features/simple.feature:5",
    "status": "failed",
    "summary": "error - user was not found by id: 1",
    "details": "error details\ndebug information",
    "duration": 125690,
    "suite": "main"
}
```

1. **event** - name of event.
2. **location** - location in source file, based on pattern {path}:{line}.
3. **status** - can be one of **passed, failed, skipped, undefined,
   ambiguous**.
4. **summary** - `optional` one line summary for step result.
5. **details** - `optional` multi-line detailed description of step
   result.
6. **duration** - `optional` duration in milliseconds to run step.
7. **suite** - `optional` may be used to distinguish test groups.

### TestCaseFinished

Should fire after all steps are executed for scenario or outline example.
Should provide appropriate result status.

``` json
{
    "event": "TestCaseFinished",
    "location": "features/simple.feature:5",
    "status": "failed",
    "duration": 125690,
    "suite": "main"
}
```

1. **event** - name of event.
2. **location** - location in source file, based on pattern {path}:{line}.
3. **status** - can be one of **passed, failed, skipped, undefined,
   ambiguous**.
4. **duration** - `optional` duration in milliseconds to run test case.
5. **suite** - `optional` may be used to distinguish test groups.

### TestRunFinished

Should fire after all tests are run. Or if stop on failure flag was
specified and failure occurred. It should carry the status of all tests
and resource usage summary information.

``` json
{
    "event": "TestRunFinished",
    "status": "failed",
    "duration": 125690,
    "memory": 3456765
}
```

1. **event** - name of event.
2. **status** - can be one of **passed, failed, skipped, undefined,
   ambiguous**.
3. **duration** - `optional` duration in milliseconds to run all tests
4. **memory** - `optional` memory consumption in bytes used by all tests

