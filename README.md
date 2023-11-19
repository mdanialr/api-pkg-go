[![Programming Language](https://img.shields.io/badge/language-GO-blue.svg)](https://shields.io/)
[![Go version](https://img.shields.io/badge/Go-v1.21-blue)](https://img.shields.io/)
[![CI Status](https://github.com/mdanialr/api-pkg-go/workflows/CI/badge.svg)](https://github.com/mdanialr/api-pkg-go/actions/workflows/on_push_pr.yml)
[![Code Coverage](https://github.com/mdanialr/api-pkg-go/wiki/coverage.svg)](https://github.com/mdanialr/api-pkg-go/wiki/coverage.svg)

# API Pkg Go
Useful collection of reusable packages for Go

- `Log`: logging pkg that support write logs to multiple output target such as `console`, `file` (with logrotate), `newrelic` and `platform log` at the same time.

## Log
There are two main parts in `log` which are __Logger__ and __Writer__. `frontend` is the API provided by `Logger` interface and `backend` is any pkg/lib that implement `Logger`

- __Logger__: Main actor that will decide where, how and whether it should write the logs or not based on the log level defined in each `Writer`.
  You may call this as the `frontend`, since you will and should only interact with the provided API from `Logger` interface.
- __Writer__: Decide where the logs passed from `Logger` should be written to. Is it to terminal, file or whatever this `Writer` will decide that.
  We already provide pre-defined `Writer` implementer namely `console`, `file`, `newrelic`, `platform`.

For now, we can support two `backend` which are [zap](https://github.com/uber-go/zap) & [slog](https://pkg.go.dev/golang.org/x/exp/slog).
Use `NewZapLogger` to use `zap` as the logger backend or `NewSlogLogger` to use `slog` instead.

### Getting Started
```go
// set console/terminal writer to listen to all logs level DEBUG, INFO, WARNING, ERROR
//  console writer, write logs to local terminal/console
cns := log.NewConsoleWriter(log.DebugLevel)	

// use predefined zap as the backend
wr := log.NewZapLogger(cns)
//  or use this if you want to use slog under the hood instead
//    wr := logger.NewSlogLogger(cns)

// call Init before using any other API
wr.Init(3 * time.Second) // you may give longer or shorter timeout/deadline

// info level log message that include contextual data 'hello':'world'
wr.Inf("INFO message", log.String("hello", "world"))
//  terminal: 2023-09-22T13:38:39.784+0700    INFO    INFO message    {"hello": "world"}
//  json: {"level":"INFO","time":"2023-09-22T13:38:39.784+0700","msg":"INFO message","hello":"world"}

// debug level log message
wr.Dbg("DEBUG message")
//  terminal: 2023-09-22T13:38:39.784+0700    DEBUG   DEBUG message
//  json: {"level":"DEBUG","time":"2023-09-22T13:38:39.784+0700","msg":"DEBUG message"}

// SHOULD be called before program exit to make sure any pending logs in buffer properly flushed by each Writer
wr.Flush(2 * time.Second) // you may give longer or shorter timeout/deadline
```

### Contextual Data
```go
// give contextual data that will be passed down to subsequent call
wr = wr.With(log.String("app_env", "local"))
wr.Wrn("warning log")
//  terminal: 2023-09-22T13:38:39.784+0700    WARN    warning log     {"app_env": "local"}
//  json: {"level":"WARN","time":"2023-09-22T13:38:39.784+0700","msg":"warning log","app_env":"local"}

wr = wr.With(log.Num("ram", 2)) // this will also accumulate previous contextual data
wr.Inf("look how many ram i have")
//  terminal: 2023-09-22T13:38:39.784+0700    INFO    look how many ram i have        {"app_env": "local", "ram": 2}
//  json: {"level":"INFO","time":"2023-09-22T13:38:39.784+0700","msg":"look how many ram i have","app_env":"local","ram":2}
```
**Note**: if the contextual data **should** change in each function call, then make sure to create new variable in each
`.With()` call.

Example
```go
func writeGoodLog(wr log.Logger) {
    // create new variable instead of replacing the old wr variable
    newWr := wr.With(
        log.String("x-request-id", uuid.NewString()),
    )

    newWr.Inf("info message")
}

func writeBadLog(wr log.Logger) {
    wr = wr.With(
        log.String("x-request-id", uuid.NewString()),
    )
	
    wr.Inf("info message")
}

func main() {
    // setup myLog that's type of log.Logger
    myLog

    writeBadLog(myLog)
    // output: {"msg":"info message", "x-request-id": "c684f881-07a5-45e6-97fd-cb2af8ad7c4e"}
    writeBadLog(myLog)
    // output: {"msg":"info message", "x-request-id": "c684f881-07a5-45e6-97fd-cb2af8ad7c4e"}
	
    // NOTE that each call on writeBadLog they will generate exactly same x-request-id, this is because the children
    //  and parent always affecting each other, if the children generate new x-request-id then the parent will also
    //   change their x-request-id and make that contextual data duplicate

	
    writeGoodLog(myLog)
    // output: {"msg":"info message", "x-request-id": "4014d36a-8f34-4b26-b91a-12480605033d"}
    writeGoodLog(myLog)
    // output: {"msg":"info message", "x-request-id": "2935ee81-a4a3-4586-b7f0-95d26473a5ac"}
	
    // This time the x-request-id will always generate new string without affecting the log.Logger from param, because
    //  it always generates new log.Logger on each writeGoodLog() instead of replace and reusing log.Logger that come from param
}
```

### Leveled Log
```go
cns := log.NewConsoleWriter(log.ErrorLevel)
wr := log.NewZapLogger(cns)
wr.Init(3 * time.Second)

// won't print, less than error
wr.Dbg("DEBUG message")

// won't print, less than error
wr.Inf("INFO message")

// won't print, less than error
wr.Wrn("warning log")

// printed, higher or equal than error
wr.Err("oops!!")
//  terminal: 2023-09-22T13:38:39.784+0700    ERROR   oops!!
//  json: {"level":"ERROR","time":"2023-09-22T13:38:39.784+0700","msg":"oops!!"}

// never forget to flush before exit
wr.Flush(1 * time.Second)
```
Log is prioritized in these order:
1. Error `Err`: (Error) print only in log level Error
2. Warning `Wrn`: (Warning, Error) print in log level Warning, Error
3. Info `Inf`: (Info, Warning, Error) print in log level Info, Warning, Error
4. Debug `Dbg`: (Debug, Info, Warning, Error) print in all log level

### Logger with Context
```go
// put the logger wr to context with 'log.WithCtx'
ctx := log.WithCtx(context.Background(), wr)
printFromCtx(ctx) // pass logger contained context

func printFromCtx(ctx context.Context) {
    // grab logger from context
    wr := log.FromCtx(ctx)
    // note that even if there is no logger inside context
    // this won't cause any panic and just return nop-logger
    // that will never write any logs
	
    wr.Inf("my information")
    //  terminal: 2023-09-22T13:38:39.784+0700    INFO    my information
    //  json: {"level":"INFO","time":"2023-09-22T13:38:39.784+0700","msg":"my information"}
}
```