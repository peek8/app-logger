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
	"github.com/peek8/app-logger/logs"
	"github.com/peek8/app-logger/model"
	"github.com/samber/lo"
)

// -------------------- MAIN --------------------

func main() {
	cfg := model.Config{}

	versionCMD := flag.NewFlagSet("version", flag.ExitOnError)

	if len(os.Args) > 1 && os.Args[1] == "version" {
		versionCMD.Parse(os.Args[2:])
		fmt.Println("app-logger version 0.1.1")
		return
	}

	flag.StringVar(&cfg.Mode, "mode", "random", "")
	flag.StringVar(&cfg.Scenario, "scenario", "", "")
	flag.IntVar(&cfg.Rate, "rate", 10, "")
	flag.StringVar(&cfg.BurstScenario, "burst", "", "")
	flag.DurationVar(&cfg.BurstInterval, "burst-interval", 30*time.Second, "")
	flag.IntVar(&cfg.BurstSize, "burst-size", 5, "")
	flag.Float64Var(&cfg.ChaosIntensity, "chaos-intensity", 1.0, "")

	flag.Parse()

	validateCfg(cfg)

	generator.RunEngine(cfg)
}

func validateCfg(cfg model.Config) {
	if cfg.Rate <= 0 {
		fmt.Println("rate must be > 0")
		os.Exit(1)
	}

	if cfg.Scenario != "" {
		// TODO check if scenario exist at map
		if _, ok := logs.ScenarioMap[cfg.Scenario]; !ok {
			fmt.Printf("scenario '%s' not found\n", cfg.Scenario)
			fmt.Printf("Must be a value from %s", lo.Keys(logs.ScenarioMap))

			os.Exit(1)
		}

	}

}
