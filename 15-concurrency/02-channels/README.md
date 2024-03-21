# Channels

Channels in Go are a powerful feature that allows goroutines to communicate and synchronize their execution. They provide a safe and efficient way to pass data between goroutines and coordinate their actions.

Through channels we can send and receive values with the channel operator `<-`. Channels enable safe communication and coordination between goroutines, preventing race conditions and other concurrency hazards.

- holds data
- thread-safe
- listens for data

Channels are created using the `make` function with `chan` keyword followed by it's type:

```go
ch := make(chan int) // Unbuffered channel of integers
```

We can also create a buffered channel by specifying the buffer length:

```go
ch := make(chan int, 100) // Buffered channel with capacity 100
```

We send a value into a channel using the `<-` operator:

```go
ch <- x  // Send x to the channel ch
```

Similarly, we receive a value from a channel using the `<-` operator:

```go
x := <-ch // Receive a value from ch and assign it to x
```

- Unbuffered channels block on send and receive operations until another goroutine is ready to receive/send.
- Buffered channels block on send when the buffer is full, and block on receive when the buffer is empty.

We can close a channel to indicate no more values will be sent. This is useful for signaling completion to the receivers.

```go
close(ch) // Close the channel
```

We can use range to iterate over values received from a channel until it's closed:

```go
for x := range ch {
    // Do something with x
}
```
