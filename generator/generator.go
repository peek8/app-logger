/*
 * Copyright (c) 2026 peek8.io
 *
 * Created Date: Friday, April 3rd 2026, 3:15:01 pm
 * Author: Md. Asraful Haque
 *
 */

package generator

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"

	"github.com/peek8/app-logger/logs"
	"github.com/peek8/app-logger/model"
)

// -------------------- HELPERS --------------------

func genTraceID() string {
	return fmt.Sprintf("trace-%d", time.Now().UnixNano())
}

func withTS(l model.Log) model.Log {
	l.Timestamp = time.Now().Format(time.RFC3339)
	return l
}

func outputLog(l model.Log) {
	b, _ := json.Marshal(l)
	fmt.Println(string(b))
}

// -------------------- NORMAL LOGS --------------------

func GenerateNormalLog(traceID string) model.Log {
	logs := []model.Log{
		logs.SampleInfoLogs[rand.Intn(len(logs.SampleInfoLogs))],
		logs.SampleDebugLogs[rand.Intn(len(logs.SampleDebugLogs))],
	}

	return logs[rand.Intn(len(logs))].WithTraceID(traceID)
}

func randomWarn(traceID string) model.Log {
	return logs.SampleWarnLogs[rand.Intn(len(logs.SampleWarnLogs))].WithTraceID(traceID)
}

func randomError(traceID string) model.Log {
	return logs.SampleErrorLogs[rand.Intn(len(logs.SampleErrorLogs))].WithTraceID(traceID)
}

// -------------------- RECOVERY --------------------

func recoveryLog(traceID string) model.Log {
	return model.Log{
		Level:   "INFO",
		Service: "system",
		TraceID: traceID,
		Message: "service recovered",
	}
}
