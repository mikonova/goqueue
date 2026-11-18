# goqueue
### An implementation of flexible lifo/fifo queues, an alternative to channels.
All queues are safe for concurrent usage.

#### Example: creating a new channel
```go
q := goqueue.New[string](100)
```
Add the elements with Push() and remove them with either Pop() or Pull()
Also, clear the whole queue with Drain(elem int)
The queue blocks the gorogoutine if a capacity is reached and waits for something to drain at least one element.