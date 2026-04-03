/*
 * Copyright (c) 2026 peek8.io
 *
 * Created Date: Friday, April 3rd 2026, 3:24:06 pm
 * Author: Md. Asraful Haque
 *
 */
package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/peek8/app-logger/generator"
	"github.com/peek8/app-logger/model"
)

// -------------------- MAIN --------------------

func main() {
	cfg := model.Config{}

	flag.StringVar(&cfg.Mode, "mode", "random", "")
	flag.StringVar(&cfg.Scenario, "scenario", "", "")
	flag.IntVar(&cfg.Rate, "rate", 10, "")
	flag.StringVar(&cfg.BurstScenario, "burst", "", "")
	flag.DurationVar(&cfg.BurstInterval, "burst-interval", 30*time.Second, "")
	flag.IntVar(&cfg.BurstSize, "burst-size", 5, "")
	flag.Float64Var(&cfg.ChaosIntensity, "chaos-intensity", 1.0, "")

	flag.Parse()

	if cfg.Rate <= 0 {
		fmt.Println("rate must be > 0")
		os.Exit(1)
	}

	generator.RunEngine(cfg)
}
