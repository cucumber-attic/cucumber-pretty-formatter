## Questions for general behavior

Printing output

- if second scenario background step fails, where to print an error?
- if scenarios does not have any steps, but there is background, should
  background steps execute, skip?
- how to treat scenarios without steps? Should they be undefined?

Suites

- guess as an event **ID** we should use **suite name** and **feature
  location**.

Steps

- How do we determine step arguments from events and print them?
