package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.21.0"
	"google.golang.org/grpc"
)

func newExporter(ctx context.Context) (sdktrace.SpanExporter, error) {
	// Point to the jaeger service name from docker-compose.
	conn, err := grpc.NewClient("jaeger:14317", grpc.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("failed to create gRPC connection: %w", err)
	}
	return otlptracegrpc.New(ctx, otlptracegrpc.WithGRPCConn(conn))
}

func newTraceProvider(exp sdktrace.SpanExporter) *sdktrace.TracerProvider {
	return sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exp),
		sdktrace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String("hello-go-app"),
		)),
	)
}

func main() {
	ctx := context.Background()

	exporter, err := newExporter(ctx)
	if err != nil {
		log.Fatalf("failed to create exporter: %v", err)
	}
	tp := newTraceProvider(exporter)
	defer func() {
		if err := tp.Shutdown(ctx); err != nil {
			log.Printf("error shutting down trace provider: %v", err)
		}
	}()
	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))

	originalHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, DevOps World! This server is now instrumented.")
	})

	handler := otelhttp.NewHandler(originalHandler, "hello-go-handler")
	http.Handle("/", handler)

	fmt.Println("Server starting on port 8080...")
	fmt.Println("Jaeger UI running on http://localhost:16686")

	server := &http.Server{Addr: ":8080"}
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("server failed to start: %v", err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Printf("server shutdown error: %v", err)
	}
}
