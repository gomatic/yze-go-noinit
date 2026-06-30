package noinit_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/tools/go/analysis/analysistest"

	noinit "github.com/gomatic/yze-go-noinit"
)

func TestInitFuncIsReported(t *testing.T) {
	analysistest.Run(t, analysistest.TestData(), noinit.Analyzer, "a")
}

func TestRegistrationIsWellFormed(t *testing.T) {
	assert.NoError(t, noinit.Registration.Validate())
	assert.Equal(t, "yze/noinit", noinit.Registration.RuleID())
	assert.Same(t, noinit.Analyzer, noinit.Registration.Analyzer)
}
