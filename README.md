# 📦 Application Log Generator

A **synthetic application log generator** designed to simulate real-world production systems with:

* ✅ Normal traffic
* ✅ Failure scenarios
* ✅ Burst incidents
* ✅ Chaos mode (overlapping failures + noise)
* ✅ Recovery signals

Perfect for:
* Observability testing (Loki / Grafana)
* Chaos engineering demos
* AI-based log analysis (Bifrost)

---

# Features

## Log Types

* INFO / DEBUG → normal traffic
* WARN → early signals
* ERROR → failures
* INFO (recovery) → system healing

## Modes

### 1. Random Mode (default)

Mix of normal logs + occasional failures

### 2. Fixed Scenario Mode

Run a single scenario continuously

### 3. Burst Mode

Inject periodic failure spikes

### 4. Chaos Mode

Simulates real production:

* noise + warnings
* random failures
* overlapping incidents
* recovery signals

---

# Example Log Output

```json
{
  "ts": "2026-04-03T12:00:00Z",
  "level": "ERROR",
  "service": "payment",
  "trace_id": "trace-123",
  "msg": "payment timeout"
}
```

---

# Available Scenarios

- payment_failure    
- db_exhaustion      
- kafka_lag          
- memory_leak        
- external_api_failure
- bad_deployment     
- dns_failure        
- cache_stampede     
- disk_full          
- auth_outage        
- network_partition  
- config_drift       
- circuit_breaker_open
- thundering_herd    
- slow_query         
- deadlock           
- fd_leak            
- ssl_cert_expired   
- time_skew          
- message_duplication
- partial_outage     
- backup_failure

👉 Easily extendable via `ScenarioMap`

---

# ⚙️ Installation

```bash
git clone https://github.com/peek8/app-logger.git
cd app-logger
go run main.go
```

---

# Usage

## Random Mode

```bash
go run main.go --mode=random --rate=10
```

---

## Fixed Scenario

```bash
go run main.go \
  --mode=fixed \
  --scenario=payment_failure
```

---

## Burst Mode

Inject failures periodically:

```bash
go run main.go \
  --mode=random \
  --burst=payment_failure \
  --burst-interval=20s \
  --burst-size=5
```

---

## Chaos Mode (Recommended)

```bash
go run main.go \
  --mode=chaos \
  --rate=30 \
  --burst-interval=15s \
  --burst-size=10 \
  --chaos-intensity=0.8
```

---

# CLI Flags

| Flag                | Description                | Default |
| ------------------- | -------------------------- | ------- |
| `--mode`            | random / fixed / chaos     | random  |
| `--scenario`        | scenario name (fixed mode) | ""      |
| `--rate`            | logs per second            | 10      |
| `--burst`           | scenario for burst         | ""      |
| `--burst-interval`  | burst frequency            | 30s     |
| `--burst-size`      | number of traces per burst | 5       |
| `--chaos-intensity` | chaos aggressiveness (0–1) | 1.0     |

---

# How Chaos Mode Works

Chaos Mode mixes:

* 70% → normal logs
* 15% → warnings
* 10% → random failures
* 5% → micro bursts

Plus:

* overlapping incidents
* random recovery signals

---

# Docker Usage

```dockerfile
FROM golang:1.22-alpine

WORKDIR /app
COPY . .

RUN go build -o app-logger

CMD ["./app-logger", "--mode=chaos", "--rate=20"]
```

---

# Kubernetes Deployment

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: app-logger
spec:
  replicas: 1
  selector:
    matchLabels:
      app: loggen
  template:
    metadata:
      labels:
        app: loggen
    spec:
      containers:
        - name: loggen
          image: app-logger:latest
          args:
            - "--mode=chaos"
            - "--rate=20"
```

---

# Use Cases

* 🔥 Chaos engineering demos
* 📊 Observability pipeline testing
* 🤖 Training AI for root cause detection
* 🧪 Load testing log ingestion systems

---

# Future Enhancements

* [ ] More scenarios
* [ ] Scenario tagging (for Loki labels)
* [ ] Helm chart
* [ ] Slack trigger integration
* [ ] AI-based incident evaluator

---

# Contributing

PRs are welcome!

---
# License
- Apache 2.0, see more details at [LICENSE File](./LICENSE).
