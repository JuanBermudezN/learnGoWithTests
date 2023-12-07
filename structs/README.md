# Structs section

## Overview

This README provides insights into key concepts and best practices applied in the codebase.

## Methods in Go

- **Functions vs. Methods:** Functions are ubiquitous, while methods are functions with receivers, invoked on specific type instances.
- **Method Declaration:** Associates methods with receiver base types, linking functionality to types.

## Formatting Precision

- **Format Strings:** Transition from 'f' to 'g' for precise decimal numbers in error messages.
- **Example:** Demonstrates 'g' usage for accurate results in a circle area calculation.

## Table Driven Tests

- **Purpose:** Table-driven tests ensure consistent testing across multiple cases.
- **Challenge:** Identifying failures in large tables and developer inconvenience.
- **Improved Error Messages:** Utilize `%#v` for structured error messages, renaming fields for clarity. g. The %#v format string will print out our struct with the values in its field, so the developer can see at a glance the properties that are being tested.

- **Test Organization:** Adopt `t.Run` and naming conventions for clearer output and targeted test execution.

## Testing Specific Cases

- **Tip:** Use `go test -run` with specific test case names (e.g., `TestArea/Rectangle`) for focused testing.

## Conclusion

By adhering to these practices, developers enhance code readability, simplify debugging, and elevate the testing experience, ensuring robust and maintainable Go code.

