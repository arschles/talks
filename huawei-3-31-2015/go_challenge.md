# The Go Challenge

This is a challenge for new Go developers at the Huawei event on 3/31/2015.

# Context

We will simulate a simple [Actor](http://en.wikipedia.org/wiki/Actor_model) in Golang
for this challenge.

The [actor model](http://en.wikipedia.org/wiki/Actor_model) is a well known concurrency
pattern. In this pattern, a single process (called an actor) holds mutable state
and serializes access to it. All communication to and from the Actor is done via sending messages, and only the Actor may modify its own internal mutable state.

# The Challenge

The challenge is to build an actor in Go that manages a `map[string]string` internally. It should run inside a goroutine and the `main()` function should communicate with it to set and get 10,000 values, in serial.

## Details

The Actor will be a goroutine that has a `for` loop inside of it. Gets will happen on a `chan getcmd` and set on a `chan setcmd`.

`getcmd` looks like this:

```go
type getcmd struct {
  key string
  respChan chan string
}

//create with get := getcmd{key:"mykey", respChan:= make(chan string)}
```

and `setcmd` looks like this:

```go
type setcmd struct {
  key string
  value string
}

//create with set := setcmd{key: "mykey", val:"myval"}
```

The approximate structure of the `func` that implements the actor looks like this:

```go
func actor() {
  for {
    select {
      //listen for messages on the channels here
    }
  }
}
```

Inside the `select` statement, you should check for sends on the get channel. You can find a detailed description on how `select` statements work [in the Go specification](https://golang.org/ref/spec#Select_statements).

If you'd like an extra challenge:

- implement the ability to shut down, pause and start the actor at-will
- implement the ability to set expirations (TTLs) on keys

# Getting Help

Feel free to email arschles+gochallenge1@gmail.com with questions.
