How to find types
https://play.golang.org/p/VaAPodMY0Yf

scope
https://play.golang.org/p/cJD1oyOIp4P

multivariate return
https://play.golang.org/p/ejypDEJr7ud

variadic functions (more than one parameters), lexical elements
https://play.golang.org/p/SrGEoZMwxUw

unfurling variadic parameters (makes into slices (slice like python's list))
https://play.golang.org/p/jBBdAxUhJEG

pointers - dereference, (Golang is pass by value)
https://play.golang.org/p/RPB-W4aAdog
https://play.golang.org/p/ZDNjy9CMqCT
*see example from "new"

defer
https://play.golang.org/p/d6VjndBRrnT

panic (error handling)
https://play.golang.org/p/k3aF1hesZBj

methods
https://play.golang.org/p/mrEuwnziJt8

structs
https://play.golang.org/p/G2tPT3uoUjC

array
https://play.golang.org/p/Vtc5db-WSTA

maps (dictionary)
https://play.golang.org/p/taMgJsRbIJ2

new
https://play.golang.org/p/-8EKAqzF7lP

numeric types
byte is alias for uint8
rune is alias for int32
https://golang.org/ref/spec#Numeric_types

struct
https://play.golang.org/p/ba8rFIR8s8v

::: MODULE :::

runtime
https://play.golang.org/p/nf9UQp9aov_f
https://golang.org/pkg/runtime/

reflect
https://play.golang.org/p/VaAPodMY0Yf

buffer:
  A buffer is a continuous region of memory. It's a term that's used a lot so it can be used kind of imprecisely, but a context I use it frequently is "allocate a buffer of 100 bytes" which requests the operating system/platform to provide you with a continuous region of 100 bytes.
  Doing many small writes can hurt performance. Each write is ultimately a syscall and if doing frequently can put burden on the CPU. Devices like disks work better dealing with block-aligned data. To avoid the overhead of many small write operations Golang is shipped with bufio.Writer.
  This might be most analogous to a list in python; as you get closer to the hardware you might care about how that list is stored in memory. One way is to allocate a buffer which has all its elements in memory sequentially, so you know you can go from one element to another in the array by incrementing through the actual memory address you have (memory address being an index into the buffer).













Caleb Doxsey's An Introduction to Programming in Go
https://www.golang-book.com/books/intro
https://www.doxsey.net/
