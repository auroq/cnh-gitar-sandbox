---
title: "Tests ship with code"
description: "Ask for tests when Go code changes without test changes"
type: check
when: "A PR changes a non-test .go file but no _test.go file"
actions: "Post a PR comment asking for a test that covers the changed behavior"
---

# Tests ship with code

When a pull request changes Go source without touching any `_test.go` file,
post one comment naming the changed functions and asking for a table-driven test
that covers the new or changed behavior.
