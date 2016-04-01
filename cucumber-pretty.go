package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type step struct {
	DefId string `json:"step_definition_id"`
	SrcId string `json:"step_source_id"`
}

type failure struct {
	Summary string `json:"error_summary"`
	Details string `json:"error_details"`
}

type stats struct {
	Time   string `json:"time"`
	Memory string `json:"memory"`
}

type reporter struct {
	Source     string `json:"source"`
	cursor      int
	currentStep *step

	totalScenarios       int
	totalSteps           int
	totalScenariosPassed int
	totalScenariosFailed int
	totalStepsPassed     int
	totalStepsFailed     int
	totalStepsSkipped    int
}

func main() {
	rep := &reporter{}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if err := rep.handleEvent(scanner.Text()); err != nil {
			log.Fatalf("failed to process event: %s", err)
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("failed to read stdin: %s", err)
	}
}

func (r *reporter) eventType(eventJSON string) (string, error) {
	typ := struct{ Event string }{}
	return typ.Event, json.Unmarshal([]byte(eventJSON), &typ)
}

func (r *reporter) handleEvent(eventJSON string) error {
	typ, err := r.eventType(eventJSON)
	if err != nil {
		return err
	}

	switch typ {
	case "SourceRead":
		return json.Unmarshal([]byte(eventJSON), r)
		// case "StartedTestingFeature":
		// 	log.Println("started feat")
	case "StartedTestingScenario":
		data := struct {
			SourceId string `json:"scenario_source_id"`
		}{}
		err := json.Unmarshal([]byte(eventJSON), &data)
		if err != nil {
			return err
		}
		line := data.SourceId[strings.Index(data.SourceId, ":")+1:]
		printTo, err := strconv.Atoi(line)
		if err != nil {
			return err
		}

		printLines := strings.Split(r.Source, "\n")[r.cursor:printTo]

		fmt.Fprintln(os.Stdout, strings.Join(printLines, "\n")+" # "+data.SourceId)
		r.cursor = printTo

		r.totalScenarios++
	case "ScenarioHasPassed":
		r.totalScenariosPassed++
	case "ScenarioHasFailed":
		r.totalScenariosFailed++
	case "FoundStepDefinition":
		r.currentStep = &step{}
		if err := json.Unmarshal([]byte(eventJSON), r.currentStep); err != nil {
			return err
		}
	case "StartedExecutingStep":
		r.totalSteps++
	case "StepHasPassed":
		line := r.currentStep.SrcId[strings.Index(r.currentStep.SrcId, ":")+1:]
		printTo, err := strconv.Atoi(line)
		if err != nil {
			return err
		}
		printLines := strings.Split(r.Source, "\n")[r.cursor:printTo]
		fmt.Fprintln(os.Stdout, strings.Join(printLines, "\n")+" # "+r.currentStep.DefId)
		r.cursor = printTo
		r.totalStepsPassed++
	case "SkippedStep":
		line := r.currentStep.SrcId[strings.Index(r.currentStep.SrcId, ":")+1:]
		printTo, err := strconv.Atoi(line)
		if err != nil {
			return err
		}
		printLines := strings.Split(r.Source, "\n")[r.cursor:printTo]
		fmt.Fprintln(os.Stdout, strings.Join(printLines, "\n")+" # "+r.currentStep.DefId)
		r.cursor = printTo
		r.totalStepsSkipped++
	case "StepHasFailed":
		failure := &failure{}
		if err := json.Unmarshal([]byte(eventJSON), failure); err != nil {
			return err
		}
		line := r.currentStep.SrcId[strings.Index(r.currentStep.SrcId, ":")+1:]
		printTo, err := strconv.Atoi(line)
		if err != nil {
			return err
		}
		printLines := strings.Split(r.Source, "\n")[r.cursor:printTo]
		fmt.Fprintln(os.Stdout, strings.Join(printLines, "\n")+" # "+r.currentStep.DefId)
		r.cursor = printTo

		fmt.Fprintln(os.Stdout, "      "+failure.Summary)
		for _, detailsLine := range strings.Split(failure.Details, "\n") {
			fmt.Fprintln(os.Stdout, "      "+detailsLine)
		}
		r.totalStepsFailed++
	case "TestingHasFinished":
		stats := &stats{}
		if err := json.Unmarshal([]byte(eventJSON), stats); err != nil {
			return err
		}

		fmt.Fprintln(os.Stdout)
		fmt.Fprintln(os.Stdout, fmt.Sprintf("%d scenarios (%d passed, %d failed)", r.totalScenarios, r.totalScenariosPassed, r.totalScenariosFailed))
		fmt.Fprintln(os.Stdout, fmt.Sprintf("%d steps (%d passed, %d failed, %d skipped)", r.totalSteps, r.totalStepsPassed, r.totalStepsFailed, r.totalStepsSkipped))

		fmt.Fprintln(os.Stdout, fmt.Sprintf("%s (%s)", stats.Time, stats.Memory))
	}
	return nil
}
