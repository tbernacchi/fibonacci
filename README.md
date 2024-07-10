# fibonacci
Fibonacci made in Go that apparentely was made using AI.

Using any language and libraries you choose, write a HTTP service with a single endpoint
that calculates n-th fibonacci sequence (1, 1, 2, 3, 5, 8, 13 ...) number. E.g. if the server is
running on http://localhost:8000 it should behave like this:

```
> curl <http://localhost:8000/fib?n=1>
1
> curl <http://localhost:8000/fib?n=10>
55
> curl <http://localhost:8000/fib?n=72>
498454011879264
```
<<<<<<< HEAD
# Locally
=======
# Go
>>>>>>> 93b08b8 (All files)

```
mkdir -p ~/fibonacci
cd ~fibonacci
mkdir server

fibonacci|⇒ tree .
.
├── fibonacci-test.go
├── fibonacci.go
└── server
    └── main.go

fibonacci|⇒ go mod init fibonacci
fibonacci|⇒ go run server/main.go
```
<<<<<<< HEAD



=======
>>>>>>> 93b08b8 (All files)
