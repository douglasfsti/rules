ctx := context.Background()
exporter, err := otlptracegrpc.New(ctx,
    otlptracegrpc.WithInsecure(),
    otlptracegrpc.WithEndpoint(os.Getenv("OTEL_EXPORTER_OTLP_ENDPOINT")),
)
if err != nil {
    log.Fatalf("failed to create exporter: %v", err)
}

tp := sdktrace.NewTracerProvider(
    sdktrace.WithBatcher(exporter),
    sdktrace.WithResource(newResource()),
)
otel.SetTracerProvider(tp)
