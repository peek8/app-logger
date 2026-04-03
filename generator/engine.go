/*
 * Copyright (c) 2026 peek8.io
 *
 * Created Date: Friday, April 3rd 2026, 3:29:11 pm
 * Author: Md. Asraful Haque
 *
 */

package generator

import (
	"math/rand"
	"time"

	"github.com/peek8/app-logger/logs"
	"github.com/peek8/app-logger/model"
)

func RunEngine(cfg model.Config) {
	ticker := time.NewTicker(time.Second / time.Duration(cfg.Rate))
	defer ticker.Stop()

	burstTicker := time.NewTicker(cfg.BurstInterval)
	defer burstTicker.Stop()

	for {
		select {

		// MAIN FLOW
		case <-ticker.C:
			traceID := genTraceID()

			switch cfg.Mode {

			case "fixed":
				runScenario(cfg.Scenario, traceID)

			case "chaos":
				r := rand.Float64()

				switch {
				case r < 0.7:
					outputLog(withTS(GenerateNormalLog(traceID)))
				case r < 0.85:
					outputLog(withTS(randomWarn(traceID)))
				case r < 0.95:
					go runScenarioRandom(traceID)
				default:
					go triggerBurst(2, "")
				}

				// recovery signal
				if rand.Float64() < 0.1 {
					outputLog(withTS(recoveryLog(traceID)))
				}

			default: // random mode
				r := rand.Float64()

				if r < 0.1 {
					go runScenarioRandom(traceID)
				} else {
					outputLog(withTS(GenerateNormalLog(traceID)))
				}
			}

		// BURST MODE
		case <-burstTicker.C:
			if cfg.BurstSize > 0 {
				size := int(float64(cfg.BurstSize) * cfg.ChaosIntensity)
				if size < 1 {
					size = 1
				}
				go triggerBurst(size, cfg.BurstScenario)
			}
		}
	}
}

// -------------------- SCENARIO EXECUTION --------------------

func runScenario(name string, traceID string) {
	scenario := logs.ScenarioMap[name]
	logs := scenario(traceID)

	for _, l := range logs {
		outputLog(withTS(l))
		time.Sleep(time.Duration(rand.Intn(80)) * time.Millisecond)
	}
}

func runScenarioRandom(traceID string) {
	keys := make([]string, 0, len(logs.ScenarioMap))
	for k := range logs.ScenarioMap {
		keys = append(keys, k)
	}
	runScenario(keys[rand.Intn(len(keys))], traceID)
}

// -------------------- BURST --------------------

func triggerBurst(size int, scenario string) {
	for i := 0; i < size; i++ {
		traceID := genTraceID()

		go func(tid string) {
			if scenario != "" {
				runScenario(scenario, tid)
			} else {
				runScenarioRandom(tid)
			}
		}(traceID)
	}
}
