main:
* contruct the app logger
* run func com logger como parametro

run:
* Configura CPU Quota
* Configuration (flag > env > code)
* App Starting - expvar, defer log shutdown, exibi configs
* Inicializa authentication support - construct a key store based on the key files sotred in the specified directory
* Start Database - defer stopping database support
* Start Tracing Support - OT/Zipkin - ver compatibilidade com opentelemetry
* Start Debug Service - mux + listenAndServe
* Start API Service
  * make a shutdown channel that listen from SIGINT SIGTERM
  * construct the API mux <-shutdown, log, metrics.New(), auth, db
  *  construct a server to service the requests against the mux (addr, handler, read/write/idle timeout, errorlog-zap.NewStdLog(log.Desugar()))
  * make a serverErrors channel that listen for errors comming from the listener - use a buffered channel so the goroutine can exit if we don't collect this error
  *  start the service listening for api requests. go func() { serverErrors <- api.ListenAndServe() }()
* Shutdown

```go
select {
case err := <-serverErrors:
    return fmt.Errorf("server error: %w", err)

case sig := <-shutdown:
    log.Infow("shutdown", "status", "shutdown started", "signal", sig)
    defer log.Infow("shutdown", "status", "shutdown complete", "signal", sig)

    // Give outstanding requests a deadline for completion.
    ctx, cancel := context.WithTimeout(context.Background(), cfg.Web.ShutdownTimeout)
    defer cancel()

    // Asking listener to shutdown and shed load.
    if err := api.Shutdown(ctx); err != nil {
        api.Close()
        return fmt.Errorf("could not stop server gracefully: %w", err)
    }
}
```