# Concurrent REST Client Example

In this project we demonstrate using goroutines, go's concurrency primitive, to make REST calls concurrently. Each successive example increases in sophistication.

1. `rest-basic.go` demonstrates making HTTP calls with simple, serial execution
2. `rest-goroutines.go` demonstrates using fork-like exeuction to make HTTP calls concurrently
3. `rest-channels.go` demonstrates using channels to make HTTP calls concurrently

Run `make` to build and run each program and see the differences in execution time and behavior.
