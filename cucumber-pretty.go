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

type test_step_finished struct {
	Status string `json:"status"`
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
	stepStatuses         map[string]int
}

func main() {
	rep := &reporter{}
	rep.stepStatuses = make(map[string]int)

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
	case "TestCaseStarted":
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
	case "TestStepMatched":
		r.currentStep = &step{}
		if err := json.Unmarshal([]byte(eventJSON), r.currentStep); err != nil {
			return err
		}
	case "TestStepStarted":
		r.totalSteps++
	case "TestStepFinished":
		test_step_finished := &test_step_finished{}
		if err := json.Unmarshal([]byte(eventJSON), test_step_finished); err != nil {
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

		r.stepStatuses[test_step_finished.Status]++

		if test_step_finished.Status == "failed" {
			fmt.Fprintln(os.Stdout, "      "+test_step_finished.Summary)
			for _, detailsLine := range strings.Split(test_step_finished.Details, "\n") {
				fmt.Fprintln(os.Stdout, "      "+detailsLine)
			}
		}
	case "TestingHasFinished":
		stats := &stats{}
		if err := json.Unmarshal([]byte(eventJSON), stats); err != nil {
			return err
		}

		fmt.Fprintln(os.Stdout)
		fmt.Fprintln(os.Stdout, fmt.Sprintf("%d scenarios (%d passed, %d failed)", r.totalScenarios, r.totalScenariosPassed, r.totalScenariosFailed))
		fmt.Fprintln(os.Stdout, fmt.Sprintf("%d steps (%d passed, %d failed, %d skipped)", r.totalSteps, r.stepStatuses["passed"], r.stepStatuses["failed"], r.stepStatuses["skipped"]))

		fmt.Fprintln(os.Stdout, fmt.Sprintf("%s (%s)", stats.Time, stats.Memory))
	}
	return nil
}
