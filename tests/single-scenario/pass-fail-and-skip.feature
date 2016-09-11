Feature: Pass fail and skip scenario
  In order to test event stream formatters
  As formatter
  I need to handle scenario with one failing step after passing and another should skip

  Scenario: Scenario
    Given passing
    When failing
    Then passing
