/*
 * Copyright (c) 2026 peek8.io
 *
 * Created Date: Friday, April 3rd 2026, 11:46:45 am
 * Author: Md. Asraful Haque
 *
 */

package logs

import (
	"math/rand"

	"github.com/peek8/app-logger/model"
)

var ScenarioMap = map[string]func(string) []model.Log{
	"payment_failure":      PaymentFailure,
	"db_exhaustion":        DBExhaustion,
	"kafka_lag":            KafkaLag,
	"memory_leak":          MemoryLeak,
	"external_api_failure": ExternalAPIFailure,
	"bad_deployment":       BadDeployment,
	"dns_failure":          DNSFailure,
	"cache_stampede":       CacheStampede,
	"disk_full":            DiskFull,
	"auth_outage":          AuthOutage,
	"network_partition":    NetworkPartition,
	"config_drift":         ConfigDrift,
	"circuit_breaker_open": CircuitBreakerOpen,
	"thundering_herd":      ThunderingHerd,
	"slow_query":           SlowQuery,
	"deadlock":             Deadlock,
	"fd_leak":              FDLeak,
	"ssl_cert_expired":     SSLCertExpired,
	"time_skew":            TimeSkew,
	"message_duplication":  MessageDuplication,
	"partial_outage":       PartialOutage,
	"backup_failure":       BackupFailure,
}

func RandomScenario() func(string) []model.Log {
	keys := make([]string, 0, len(ScenarioMap))
	for k := range ScenarioMap {
		keys = append(keys, k)
	}
	return ScenarioMap[keys[rand.Intn(len(keys))]]
}

// 1. Payment Failure
func PaymentFailure(traceID string) []model.Log {
	return []model.Log{
		{Level: "INFO", Service: "auth", TraceID: traceID, Message: "user authenticated"},
		{Level: "INFO", Service: "order", TraceID: traceID, Message: "order created"},
		{Level: "DEBUG", Service: "payment", TraceID: traceID, Message: "initiating payment"},
		{Level: "WARN", Service: "payment", TraceID: traceID, Message: "retry attempt 1"},
		{Level: "WARN", Service: "payment", TraceID: traceID, Message: "retry attempt 2"},
		{Level: "ERROR", Service: "payment", TraceID: traceID, Message: "payment timeout"},
	}
}

// 2. DB Exhaustion
func DBExhaustion(traceID string) []model.Log {
	return []model.Log{
		{Level: "INFO", Service: "db", TraceID: traceID, Message: "connection pool ready"},
		{Level: "DEBUG", Service: "db", TraceID: traceID, Message: "acquiring connection"},
		{Level: "WARN", Service: "db", TraceID: traceID, Message: "pool usage 85%"},
		{Level: "WARN", Service: "db", TraceID: traceID, Message: "pool usage 95%"},
		{Level: "ERROR", Service: "db", TraceID: traceID, Message: "connection pool exhausted"},
		{Level: "ERROR", Service: "api", TraceID: traceID, Message: "database unavailable"},
	}
}

// 3. Kafka Lag
func KafkaLag(traceID string) []model.Log {
	return []model.Log{
		{Level: "INFO", Service: "producer", TraceID: traceID, Message: "event published"},
		{Level: "DEBUG", Service: "consumer", TraceID: traceID, Message: "polling messages"},
		{Level: "WARN", Service: "consumer", TraceID: traceID, Message: "lag increasing"},
		{Level: "WARN", Service: "consumer", TraceID: traceID, Message: "processing delay"},
		{Level: "ERROR", Service: "consumer", TraceID: traceID, Message: "timeout processing message"},
		{Level: "ERROR", Service: "consumer", TraceID: traceID, Message: "consumer crashed"},
	}
}

// 4. Memory Leak
func MemoryLeak(traceID string) []model.Log {
	return []model.Log{
		{Level: "INFO", Service: "worker", TraceID: traceID, Message: "job started"},
		{Level: "DEBUG", Service: "worker", TraceID: traceID, Message: "allocating memory"},
		{Level: "WARN", Service: "worker", TraceID: traceID, Message: "memory usage 70%"},
		{Level: "WARN", Service: "worker", TraceID: traceID, Message: "memory usage 90%"},
		{Level: "ERROR", Service: "worker", TraceID: traceID, Message: "out of memory"},
		{Level: "ERROR", Service: "worker", TraceID: traceID, Message: "process crashed"},
	}
}

// 5. External API Failure
func ExternalAPIFailure(traceID string) []model.Log {
	return []model.Log{
		{Level: "INFO", Service: "gateway", TraceID: traceID, Message: "calling external api"},
		{Level: "DEBUG", Service: "gateway", TraceID: traceID, Message: "building request"},
		{Level: "WARN", Service: "gateway", TraceID: traceID, Message: "slow response"},
		{Level: "WARN", Service: "gateway", TraceID: traceID, Message: "retry attempt"},
		{Level: "ERROR", Service: "gateway", TraceID: traceID, Message: "api unavailable"},
		{Level: "INFO", Service: "gateway", TraceID: traceID, Message: "fallback used"},
	}
}

// 6. Bad Deployment
func BadDeployment(traceID string) []model.Log {
	return []model.Log{
		{Level: "INFO", Service: "deploy", TraceID: traceID, Message: "deployment started"},
		{Level: "DEBUG", Service: "api", TraceID: traceID, Message: "loading config"},
		{Level: "WARN", Service: "deploy", TraceID: traceID, Message: "new pods starting"},
		{Level: "WARN", Service: "api", TraceID: traceID, Message: "config validation warning"},
		{Level: "ERROR", Service: "api", TraceID: traceID, Message: "config parse failed"},
		{Level: "INFO", Service: "deploy", TraceID: traceID, Message: "rollback initiated"},
	}
}

// 7. DNS Failure
func DNSFailure(traceID string) []model.Log {
	return []model.Log{
		{Level: "INFO", Service: "gateway", TraceID: traceID, Message: "resolving hostname"},
		{Level: "DEBUG", Service: "dns", TraceID: traceID, Message: "querying dns server"},
		{Level: "WARN", Service: "dns", TraceID: traceID, Message: "slow response"},
		{Level: "WARN", Service: "dns", TraceID: traceID, Message: "retry lookup"},
		{Level: "ERROR", Service: "dns", TraceID: traceID, Message: "resolution failed"},
		{Level: "ERROR", Service: "gateway", TraceID: traceID, Message: "external api unreachable"},
	}
}

// 8. Cache Stampede
func CacheStampede(traceID string) []model.Log {
	return []model.Log{
		{Level: "INFO", Service: "api", TraceID: traceID, Message: "incoming request"},
		{Level: "DEBUG", Service: "cache", TraceID: traceID, Message: "checking cache"},
		{Level: "WARN", Service: "cache", TraceID: traceID, Message: "cache miss"},
		{Level: "WARN", Service: "cache", TraceID: traceID, Message: "multiple rebuilds"},
		{Level: "ERROR", Service: "db", TraceID: traceID, Message: "too many connections"},
		{Level: "ERROR", Service: "api", TraceID: traceID, Message: "failed to fetch data"},
	}
}

// 9. Disk Full
func DiskFull(traceID string) []model.Log {
	return []model.Log{
		{Level: "INFO", Service: "storage", TraceID: traceID, Message: "disk usage 70%"},
		{Level: "DEBUG", Service: "storage", TraceID: traceID, Message: "writing file"},
		{Level: "WARN", Service: "storage", TraceID: traceID, Message: "disk usage 90%"},
		{Level: "WARN", Service: "storage", TraceID: traceID, Message: "cleanup recommended"},
		{Level: "ERROR", Service: "storage", TraceID: traceID, Message: "no space left"},
		{Level: "ERROR", Service: "db", TraceID: traceID, Message: "write failed"},
	}
}

// 10. Auth Outage
func AuthOutage(traceID string) []model.Log {
	return []model.Log{
		{Level: "INFO", Service: "api", TraceID: traceID, Message: "login request"},
		{Level: "DEBUG", Service: "auth", TraceID: traceID, Message: "validating token"},
		{Level: "WARN", Service: "auth", TraceID: traceID, Message: "db latency high"},
		{Level: "WARN", Service: "api", TraceID: traceID, Message: "retry auth"},
		{Level: "ERROR", Service: "auth", TraceID: traceID, Message: "db connection failed"},
		{Level: "ERROR", Service: "api", TraceID: traceID, Message: "login failed"},
	}
}

// 11. Network Partition
func NetworkPartition(traceID string) []model.Log {
	return []model.Log{
		{Level: "INFO", Service: "order", TraceID: traceID, Message: "creating order"},
		{Level: "DEBUG", Service: "network", TraceID: traceID, Message: "sending packet"},
		{Level: "WARN", Service: "network", TraceID: traceID, Message: "packet loss detected"},
		{Level: "WARN", Service: "order", TraceID: traceID, Message: "retry request"},
		{Level: "ERROR", Service: "payment", TraceID: traceID, Message: "timeout"},
		{Level: "ERROR", Service: "order", TraceID: traceID, Message: "order failed"},
	}
}

// 12. Config Drift
func ConfigDrift(traceID string) []model.Log {
	return []model.Log{
		{Level: "INFO", Service: "config", TraceID: traceID, Message: "loading config"},
		{Level: "DEBUG", Service: "config", TraceID: traceID, Message: "comparing versions"},
		{Level: "WARN", Service: "config", TraceID: traceID, Message: "version mismatch"},
		{Level: "WARN", Service: "api", TraceID: traceID, Message: "fallback config"},
		{Level: "ERROR", Service: "api", TraceID: traceID, Message: "invalid config"},
		{Level: "ERROR", Service: "worker", TraceID: traceID, Message: "job failed"},
	}
}

// 13. Circuit Breaker
func CircuitBreakerOpen(traceID string) []model.Log {
	return []model.Log{
		{Level: "INFO", Service: "gateway", TraceID: traceID, Message: "calling downstream"},
		{Level: "DEBUG", Service: "gateway", TraceID: traceID, Message: "tracking failures"},
		{Level: "WARN", Service: "gateway", TraceID: traceID, Message: "failure count high"},
		{Level: "WARN", Service: "gateway", TraceID: traceID, Message: "threshold nearing"},
		{Level: "ERROR", Service: "gateway", TraceID: traceID, Message: "circuit opened"},
		{Level: "INFO", Service: "gateway", TraceID: traceID, Message: "fallback served"},
	}
}

// 14. Thundering Herd
func ThunderingHerd(traceID string) []model.Log {
	return []model.Log{
		{Level: "INFO", Service: "api", TraceID: traceID, Message: "traffic spike"},
		{Level: "DEBUG", Service: "lb", TraceID: traceID, Message: "routing requests"},
		{Level: "WARN", Service: "api", TraceID: traceID, Message: "high concurrency"},
		{Level: "WARN", Service: "db", TraceID: traceID, Message: "query queue growing"},
		{Level: "ERROR", Service: "db", TraceID: traceID, Message: "overloaded"},
		{Level: "ERROR", Service: "api", TraceID: traceID, Message: "request failed"},
	}
}

// 15. Slow Query
func SlowQuery(traceID string) []model.Log {
	return []model.Log{
		{Level: "INFO", Service: "api", TraceID: traceID, Message: "processing request"},
		{Level: "DEBUG", Service: "db", TraceID: traceID, Message: "executing query"},
		{Level: "WARN", Service: "db", TraceID: traceID, Message: "query slow"},
		{Level: "WARN", Service: "api", TraceID: traceID, Message: "response delayed"},
		{Level: "ERROR", Service: "api", TraceID: traceID, Message: "timeout"},
		{Level: "ERROR", Service: "client", TraceID: traceID, Message: "request failed"},
	}
}

// 16. Deadlock
func Deadlock(traceID string) []model.Log {
	return []model.Log{
		{Level: "INFO", Service: "db", TraceID: traceID, Message: "transaction started"},
		{Level: "DEBUG", Service: "db", TraceID: traceID, Message: "locking rows"},
		{Level: "WARN", Service: "db", TraceID: traceID, Message: "lock contention"},
		{Level: "WARN", Service: "db", TraceID: traceID, Message: "waiting for lock"},
		{Level: "ERROR", Service: "db", TraceID: traceID, Message: "deadlock detected"},
		{Level: "ERROR", Service: "api", TraceID: traceID, Message: "transaction rolled back"},
	}
}

// 17. FD Leak
func FDLeak(traceID string) []model.Log {
	return []model.Log{
		{Level: "INFO", Service: "system", TraceID: traceID, Message: "opening files"},
		{Level: "DEBUG", Service: "system", TraceID: traceID, Message: "tracking descriptors"},
		{Level: "WARN", Service: "system", TraceID: traceID, Message: "fd usage high"},
		{Level: "WARN", Service: "system", TraceID: traceID, Message: "approaching limit"},
		{Level: "ERROR", Service: "system", TraceID: traceID, Message: "too many open files"},
		{Level: "ERROR", Service: "api", TraceID: traceID, Message: "cannot open file"},
	}
}

// 18. SSL Expired
func SSLCertExpired(traceID string) []model.Log {
	return []model.Log{
		{Level: "INFO", Service: "client", TraceID: traceID, Message: "initiating tls handshake"},
		{Level: "DEBUG", Service: "gateway", TraceID: traceID, Message: "validating cert"},
		{Level: "WARN", Service: "gateway", TraceID: traceID, Message: "certificate nearing expiry"},
		{Level: "WARN", Service: "client", TraceID: traceID, Message: "retry handshake"},
		{Level: "ERROR", Service: "gateway", TraceID: traceID, Message: "certificate expired"},
		{Level: "ERROR", Service: "client", TraceID: traceID, Message: "tls handshake failed"},
	}
}

// 19. Time Skew
func TimeSkew(traceID string) []model.Log {
	return []model.Log{
		{Level: "INFO", Service: "system", TraceID: traceID, Message: "checking time sync"},
		{Level: "DEBUG", Service: "system", TraceID: traceID, Message: "ntp drift detected"},
		{Level: "WARN", Service: "system", TraceID: traceID, Message: "clock drift 2s"},
		{Level: "WARN", Service: "auth", TraceID: traceID, Message: "token expiry mismatch"},
		{Level: "ERROR", Service: "auth", TraceID: traceID, Message: "token validation failed"},
		{Level: "ERROR", Service: "api", TraceID: traceID, Message: "request rejected"},
	}
}

// 20. Message Duplication
func MessageDuplication(traceID string) []model.Log {
	return []model.Log{
		{Level: "INFO", Service: "queue", TraceID: traceID, Message: "message received"},
		{Level: "DEBUG", Service: "worker", TraceID: traceID, Message: "processing message"},
		{Level: "WARN", Service: "queue", TraceID: traceID, Message: "duplicate detected"},
		{Level: "WARN", Service: "worker", TraceID: traceID, Message: "idempotency check"},
		{Level: "ERROR", Service: "worker", TraceID: traceID, Message: "duplicate processed"},
		{Level: "ERROR", Service: "db", TraceID: traceID, Message: "data inconsistency"},
	}
}

// 21. Partial Outage
func PartialOutage(traceID string) []model.Log {
	return []model.Log{
		{Level: "INFO", Service: "api", TraceID: traceID, Message: "handling request"},
		{Level: "DEBUG", Service: "lb", TraceID: traceID, Message: "routing to instance"},
		{Level: "WARN", Service: "api", TraceID: traceID, Message: "instance unhealthy"},
		{Level: "WARN", Service: "lb", TraceID: traceID, Message: "traffic imbalance"},
		{Level: "ERROR", Service: "api", TraceID: traceID, Message: "request failed"},
		{Level: "INFO", Service: "lb", TraceID: traceID, Message: "rerouting traffic"},
	}
}

// 22. Backup Failure
func BackupFailure(traceID string) []model.Log {
	return []model.Log{
		{Level: "INFO", Service: "backup", TraceID: traceID, Message: "backup started"},
		{Level: "DEBUG", Service: "backup", TraceID: traceID, Message: "reading data"},
		{Level: "WARN", Service: "backup", TraceID: traceID, Message: "slow storage"},
		{Level: "WARN", Service: "backup", TraceID: traceID, Message: "retry upload"},
		{Level: "ERROR", Service: "backup", TraceID: traceID, Message: "storage unavailable"},
		{Level: "ERROR", Service: "backup", TraceID: traceID, Message: "backup failed"},
	}
}
