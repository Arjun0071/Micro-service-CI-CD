package metrics

import "github.com/prometheus/client_golang/prometheus"


var RequestCount = prometheus.NewCounterVec(
    prometheus.CounterOpts{
        Name: "http_requests_total",
        Help: "Total HTTP requests",
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


var UsersRegistered = prometheus.NewCounterVec(
    prometheus.CounterOpts{
        Name: "users_registered_total",
        Help: "Total number of successful user registrations",
    },
    []string{"service"},
)

var LoginAttempts = prometheus.NewCounterVec(
    prometheus.CounterOpts{
        Name: "login_attempts_total",
        Help: "Total number of login attempts",
    },
    []string{"service"},
)

var LoginFailures = prometheus.NewCounterVec(
    prometheus.CounterOpts{
        Name: "login_failures_total",
        Help: "Total number of failed login attempts",
    },
    []string{"service"},
)

var UsersUpdated = prometheus.NewCounterVec(
    prometheus.CounterOpts{
        Name: "users_updated_total",
        Help: "Total number of profile updates",
    },
    []string{"service"},
)


func Init(service string) {
    prometheus.MustRegister(RequestCount)
    prometheus.MustRegister(RequestDuration)
    prometheus.MustRegister(ActiveConnections)

    prometheus.MustRegister(UsersRegistered)
    prometheus.MustRegister(LoginAttempts)
    prometheus.MustRegister(LoginFailures)
    prometheus.MustRegister(UsersUpdated)

    ActiveConnections.WithLabelValues(service).Set(0)
}

