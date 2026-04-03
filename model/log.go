/*
 * Copyright (c) 2026 peek8.io
 *
 * Created Date: Friday, April 3rd 2026, 11:45:42 am
 * Author: Md. Asraful Haque
 *
 */
package model

type Log struct {
	Timestamp string            `json:"ts"`
	Level     string            `json:"level"`
	Service   string            `json:"service"`
	TraceID   string            `json:"trace_id"`
	Message   string            `json:"msg"`
	Labels    map[string]string `json:"labels,omitempty"`
}

func (l Log) WithTraceID(traceID string) Log {
	l.TraceID = traceID
	return l
}

func (l Log) WithTimestamp(timestamp string) Log {
	l.Timestamp = timestamp

	return l
}
