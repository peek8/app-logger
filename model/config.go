/*
 * Copyright (c) 2026 peek8.io
 *
 * Created Date: Friday, April 3rd 2026, 3:19:27 pm
 * Author: Md. Asraful Haque
 *
 */
package model

import "time"

type Config struct {
	Mode           string
	Scenario       string
	Rate           int
	BurstScenario  string
	BurstInterval  time.Duration
	BurstSize      int
	ChaosIntensity float64
}
