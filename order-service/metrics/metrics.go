package metrics

import "github.com/prometheus/client_golang/prometheus"

// --- Core HTTP metrics (vectors with service label) ---
var RequestCount = prometheus.NewCounterVec(
    prometheus.CounterOpts{
        Name: "http_requests_total",
        Help: "Total HTTP requests received",
    },
    []string{"service", "method", "path", "status"},
)

var RequestDuration = prometheus.NewHistogramVec(
    prometheus.HistogramOpts{
        Name:    "http_request_duration_seconds",
        Help:    "Request duration in seconds",
        Buckets: prometheus.DefBuckets,
    },
    []string{"service", "method", "path"},
)

var ActiveConnections = prometheus.NewGaugeVec(
    prometheus.GaugeOpts{
        Name: "http_active_connections",
        Help: "Active HTTP connections",
    },
    []string{"service"},
)

// --- Business metrics (per-service label) ---
var OrdersCreated = prometheus.NewCounterVec(
    prometheus.CounterOpts{
        Name: "orders_created_total",
        Help: "Total number of successfully created orders",
    },
    []string{"service"},
)

var OrdersFailed = prometheus.NewCounterVec(
    prometheus.CounterOpts{
        Name: "orders_failed_total",
        Help: "Total number of failed order attempts",
    },
    []string{"service"},
)

// Use a Counter for revenue (Add supported). Labelled by service.
var OrdersRevenue = prometheus.NewCounterVec(
    prometheus.CounterOpts{
        Name: "orders_revenue_total",
        Help: "Total revenue generated from orders",
    },
    []string{"service"},
)

var OrderCreationDuration = prometheus.NewHistogramVec(
    prometheus.HistogramOpts{
        Name:    "order_creation_duration_seconds",
        Help:    "Time taken to create an order",
        Buckets: prometheus.DefBuckets,
    },
    []string{"service"},
)

// Init registers everything and initializes gauges
func Init(service string) {
    prometheus.MustRegister(RequestCount)
    prometheus.MustRegister(RequestDuration)
    prometheus.MustRegister(ActiveConnections)

    prometheus.MustRegister(OrdersCreated)
    prometheus.MustRegister(OrdersFailed)
    prometheus.MustRegister(OrdersRevenue)
    prometheus.MustRegister(OrderCreationDuration)

    ActiveConnections.WithLabelValues(service).Set(0)
}

