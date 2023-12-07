# Test Rules

## Overview

This document outlines the rules and best practices for writing tests in Go, specifically focusing on the conventions and requirements associated with test functions.

## Test File Naming

- The test file should be named in the format `xxx_test.go`.
  
## Test Function Requirements

- The test function must begin with the word `Test`.
- It should take only one argument, which is `t *testing.T`.
- Import the "testing" package to use the *testing.T type, similar to how "fmt" is imported in other files.

## Test Driven Development (TDD)

- Emphasizes capturing requirements in tests before implementation to ensure the test covers the intended functionality.
- Group tests around a specific "thing" and use subtests to describe different scenarios.

## Helper Functions

- For helper functions, accept a `testing.TB`, an interface satisfied by both *testing.T and *testing.B.
- Allows calling helper functions from tests or benchmarks.
- `t.Helper()` indicates that a method is a helper, aiding in error tracking by reporting line numbers in the function call.

## TDD Cycle

1. **Write a Test:** Define the test cases first.
2. **Make the Compiler Pass:** Implement the minimum code required for compilation.
3. **Run the Test:** Verify that the test fails and check the error message for clarity.
4. **Write Code to Pass Test:** Develop enough code to make the test cases pass.
5. **Refactor:** Optimize and improve code without changing its functionality.