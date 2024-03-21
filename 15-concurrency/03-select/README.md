# Select Statement

- lets a goroutine wait on multiple communication operations
- blocks until one of its cases can run, then it executes that case. It chooses one at random if multiple are ready.

Sometimes we have a single goroutine listening to multiple channels and want to process data in the order it comes through each channel.

A `select` statement is used to listen to multiple channels at the same time. It is similar to a `switch` statement but for channels.

```go
select {
case i, ok := <- chInts:
    fmt.Println(i)
case s, ok := <- chStrings:
    fmt.Println(s)
}
```

The first channel with a value ready to be received will fire and its body will execute. If multiple channels are ready at the same time one is chosen randomly. The `ok` variable in the example above refers to whether or not the channel has been closed by the sender yet.

The default case in a select statement executes immediately if no other channel has a value ready. A default case stops the select statement from blocking.

```go
select {
  case v := <-ch:
    // use v
  default:
    // receiving from ch would block
    // so do something else
}
```
