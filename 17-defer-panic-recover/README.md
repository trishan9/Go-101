# defer, panic, recover

## defer

A defer statement pushes a function call onto a list. The list of saved calls is executed after the surrounding function returns. Defer is commonly used to simplify functions that perform various clean-up actions.

**Rules**

- A deferred function’s arguments are evaluated when the defer statement is evaluated.
- Deferred function calls are executed in Last In First Out order after the surrounding function returns. (Works like stack)

```go
func a() {
	i := 0
	defer fmt.Println(i)
	i++
	fmt.Println(i)
    return
}
```

```sh
Output:
1
0
```

## panic and recover

Panic is a built-in function that stops the ordinary flow of control and begins panicking. When the function F calls panic, execution of F stops, any deferred functions in F are executed normally, and then F returns to its caller. To the caller, F then behaves like a call to panic. The process continues up the stack until all functions in the current goroutine have returned, at which point the program crashes. Panics can be initiated by invoking panic directly. They can also be caused by runtime errors, such as out-of-bounds array accesses.

Recover is a built-in function that regains control of a panicking goroutine. Recover is only useful inside deferred functions. During normal execution, a call to recover will return nil and have no other effect. If the current goroutine is panicking, a call to recover will capture the value given to panic and resume normal execution.

```go
func division(num1, num2 int) {
  // execute the handlePanic even after panic occurs
  defer handlePanic()

  if num2 == 0 {
    panic("Cannot divide a number by zero") // runs after deferred function
  } else {
    result := num1 / num2
    fmt.Println("Result: ", result)
  }
}

func handlePanic() {
  // detect if panic occurs or not
  a := recover() // Recovery statements are used to come out of a panic

  if a != nil {
    fmt.Println("RECOVER", a)
  }
}
```

- We are using defer to call handlePanic(). It's because we only want to handle the panic after it occurs, so we are deferring its execution.

## In short

- We use `defer` to delay the execution of functions that might cause an error. The `panic` statement terminates the program immediately and `recover` is used to recover the message during panic.
