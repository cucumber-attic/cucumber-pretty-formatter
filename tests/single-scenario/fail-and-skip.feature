Feature: Fail and skip scenario
  In order to test event stream formatters
  As formatter
  I need to handle scenario with one failing step and another should skip

  Scenario: Scenario
    When failing
    Then failing
