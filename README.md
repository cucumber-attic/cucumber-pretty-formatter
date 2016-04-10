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
2. [GherkinSourceRead](#gherkinsourceread)
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
2. **version** - protocol version used for events.

### GherkinSourceRead

Currently only **gherkin** source is supported and after a feature file is
parsed it should send this event with it's plain source.

``` json
{
    "event": "GherkinSourceRead",
    "id": "features/simple.feature:1",
    "suite": "main",
    "source": "Feature:\n  Scenario: passing\n    Given passes"
}
```

1. **event** - name of event.
2. **id** - location in feature file, based on pattern {path}:{line}.
3. **suite** - may not be provided. Otherwise it may be used as a name for
   filter groups.
4. **source** - is plain feature file source.

### StepDefinitionFound

Should fire when step regexp or other matching algorithm determines step
implementation in the source code. Note: there may be ambiguous matches.


``` json
{
    "event": "StepDefinitionFound",
    "id": "features/simple.feature:5",
    "suite": "main",
    "definition_id": "FeatureContext::passing():6",
    "arguments": [
        [12, 18],
        [23, 26]
    ]
}
```

1. **event** - name of event.
2. **id** - location in feature file, based on pattern {path}:{line}.
3. **suite** - may not be provided. Otherwise it may be used as a name for
   filter groups.
4. **definition_id** - reference to step definition in source code.
5. **arguments** - list of positions for arguments which were matched.
   Positions are determined on **step text** step keyword should be
   omitted when calculating argument positions.

### TestCaseStarted

Should fire when starting to execute scenario or scenario outline example.

``` json
{
    "event": "TestCaseStarted",
    "id": "features/simple.feature:4",
    "suite": "main"
}
```

1. **event** - name of event.
2. **id** - location in feature file, based on pattern {path}:{line}.
3. **suite** - may not be provided. Otherwise it may be used as a name for
   filter groups.

### TestStepStarted

Should fire right before step execution.

``` json
{
    "event": "TestStepStarted",
    "id": "features/simple.feature:5",
    "suite": "main"
}
```

1. **event** - name of event.
2. **id** - location in feature file, based on pattern {path}:{line}.
3. **suite** - may not be provided. Otherwise it may be used as a name for
   filter groups.

### TestStepFinished

Should fire right after step has finished execution and give appropriate
status and details.

``` json
{
    "event": "TestStepFinished",
    "id": "features/simple.feature:5",
    "suite": "main",
    "status": "failed",
    "summary": "error - user was not found by id: 1",
    "details": "error details\ndebug information"
}
```

1. **event** - name of event.
2. **id** - location in feature file, based on pattern {path}:{line}.
3. **suite** - may not be provided. Otherwise it may be used as a name for
   filter groups.
4. **status** - can be one of **passed, failed, skipped, undefined,
   ambiguous**.
5. **summary** - one line summary of message. Maybe omitted.
6. **details** - multi-line detailed description of step result.

### TestCaseFinished

Should fire after all steps are executed for scenario or outline example.
Should provide appropriate result status.

``` json
{
    "event": "TestCaseFinished",
    "id": "features/simple.feature:5",
    "suite": "main",
    "status": "failed"
}
```

1. **event** - name of event.
2. **id** - location in feature file, based on pattern {path}:{line}.
3. **suite** - may not be provided. Otherwise it may be used as a name for
   filter groups.
4. **status** - can be one of **passed, failed, skipped, undefined,
   ambiguous**.

### TestRunFinished

Should fire after all tests are run. Or if stop on failure flag was
specified and failure occurred. It should carry the status of all tests
and resource usage summary information.

``` json
{
    "event": "TestRunFinished",
    "id": "features/simple.feature:5",
    "suite": "main",
    "status": "failed",
    "time": "12m13.08977s",
    "memory": "3.45M"
}
```

1. **event** - name of event.
2. **id** - location in feature file, based on pattern {path}:{line}.
3. **suite** - may not be provided. Otherwise it may be used as a name for
   filter groups.
4. **status** - can be one of **passed, failed, skipped, undefined,
   ambiguous**.
5. **time** - how much time it took to run tests
6. **memory** - how much memory was consumed by tests

