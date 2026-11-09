# Go review instructions

- Tests are table-driven and named "when [scenario], and [context], it should [expectation]".
- Every exported function needs a test case for its edge inputs: empty input, zero, negative numbers.
- Flag any slice expression that can panic on a negative or out-of-range index.
- Prefer returning an error over panicking in library code.
