Todos example
---

# Build
```
go build cmd/todos/todos.go
```

# Run (Dev)
```
./todos 
```

##  Run (Dev + Live Reload)

Install [air](https://github.com/cosmtrek/air) and simply run `air` on the console.

This will start the app with live-reloading on source code change.

# Run (Prod)
```
GIN_MODE=release ./todos 
```
