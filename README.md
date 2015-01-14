# Concurrent REST Client Example

This project demonstrates using goroutines, go's concurrency primitive, to make HTTP calls concurrently. Each successive example increases in sophistication.

1. `rest-basic.go` demonstrates making HTTP calls with simple, serial execution
2. `rest-waitgroup.go` demonstrates using fork-like exeuction to make HTTP calls concurrently, and uses a WaitGroup to wait for all of the HTTP calls to complete
3. `rest-channels.go` demonstrates using channels to wait for all of the HTTP calls to complete

Run `make` to build and run each program and see the differences in execution time and behavior. The WaitGroup and Channel implementations behave similarly, but the implementations are very different. Examine the code to see how each works.
