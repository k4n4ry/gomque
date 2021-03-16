# gomque
## usage
### enqueue
```
go run main.go -c amqp://guest:guest@localhost:5672/ -q task_queue hoge
```

or 
```
go run main.go -u guest -p guest -h localhost -P 5672 -q task_queue hoge
```

### dequeue
just developing