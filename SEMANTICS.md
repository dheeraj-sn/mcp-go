# MCP Semantics

## Tools
- create_limit
- get_limit
- delete_limit

## Error Codes
- INVALID_INPUT (no retry)
- NOT_FOUND (no retry)
- INTERNAL (retryable)

## Idempotency
- create_limit is idempotent on key

## Context
- user_id required
- budget decremented per call
