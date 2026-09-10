---
name: summarize-changes
description: Summarizes uncommitted changes and flags anything risky.
  Use when the user asks what changed, wants a commit message, or asks
  to review their diff before committing.
---

## Current changes

!`git diff HEAD`

## Instructions

Summarize the changes above in two or three bullet points. Then list any
risks: missing error handling, hardcoded values, secrets, or tests that
need updating. If the diff is empty, say there are no uncommitted changes.
