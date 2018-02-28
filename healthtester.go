package main

import (
	"math/rand"

	"github.com/hootsuite/healthchecks"
)

type TestHealthChecker struct {
	Status             healthchecks.Status
	ShouldRandomlyFail bool
}

func (t TestHealthChecker) CheckStatus(name string) healthchecks.StatusList {
	if t.ShouldRandomlyFail {
		i := rand.Intn(10)
		status := t.Status
		if i <= 2 {
			status.Result = healthchecks.CRITICAL
		} else {
			status.Result = healthchecks.OK
			status.Description = ""
			status.Details = ""
		}
		return healthchecks.StatusList{StatusList: []healthchecks.Status{status}}
	} else {
		return healthchecks.StatusList{StatusList: []healthchecks.Status{t.Status}}
	}
}
