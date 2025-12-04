package metrics

import (
    "github.com/prometheus/client_golang/prometheus"
)

// Core HTTP Metrics //

var RequestCount = prometheus.NewCounterVec(
    prometheus.CounterOpts{
        Name: "http_requests_total",
        Help: "Total number of HTTP requests",
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
        Help: "Number of active HTTP connections",
    },
    []string{"service"},
)

// Business (Inventory) Metrics //

// Counts how many books were created
var BooksCreated = prometheus.NewCounterVec(
    prometheus.CounterOpts{
        Name: "books_created_total",
        Help: "Total number of books created",
    },
    []string{"service"},
)

// Counts how many books were updated
var BooksUpdated = prometheus.NewCounterVec(
    prometheus.CounterOpts{
        Name: "books_updated_total",
        Help: "Total number of books updated",
    },
    []string{"service"},
)

// Counts how many books were deleted
var BooksDeleted = prometheus.NewCounterVec(
    prometheus.CounterOpts{
        Name: "books_deleted_total",
        Help: "Total number of books deleted",
    },
    []string{"service"},
)

// Tracks current stock per book
var BookStock = prometheus.NewGaugeVec(
    prometheus.GaugeOpts{
        Name: "book_stock",
        Help: "Stock level per book by title",
    },
    []string{"service", "title"},
)

// Initialization//

func Init(service string) {
    // Core metrics
    prometheus.MustRegister(RequestCount)
    prometheus.MustRegister(RequestDuration)
    prometheus.MustRegister(ActiveConnections)

    // Business metrics
    prometheus.MustRegister(BooksCreated)
    prometheus.MustRegister(BooksUpdated)
    prometheus.MustRegister(BooksDeleted)
    prometheus.MustRegister(BookStock)

    // Initialize gauges
    ActiveConnections.WithLabelValues(service).Set(0)
}

