## Request
```
{
    "request_id": "req-123",
    "session": {
      "id": "sess-1",
      "user_id": "user-42",
      "budget": 10
    },
    "action": {
      "type": "tool",
      "tool": {
        "name": "create_limit",
        "input": {
          "key": "api_key_123",
          "limit": 100
        }
      }
    }
}
```
## Success Response
```
{
    "request_id": "req-123",
    "result": {
        "id": "limit-abc"
    },
    "meta": {
        "trace_id": "trace-xyz",
        "retryable": true
    }
}
```
## Error Response
```
{
    "request_id": "req-123",
    "error": {
        "code": "INVALID_INPUT",
        "message": "limit must be > 0"
    },
    "meta": {
        "retryable": false
    }
}
```