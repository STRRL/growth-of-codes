package scc

import (
	"bytes"
	"encoding/json"
	"github.com/boyter/scc/v3/processor"
	"github.com/pkg/errors"
	"os/exec"
)

type Adaptor interface {
	Analyze(path string) ([]processor.LanguageSummary, error)
}

type CommandLineSCC struct {
}

func NewCommandLineSCC() *CommandLineSCC {
	return &CommandLineSCC{}
}

func (it *CommandLineSCC) Analyze(path string) ([]processor.LanguageSummary, error) {
	command := exec.Command("scc", "--by-file", "-f", "json", path)
	buf := bytes.NewBuffer(nil)
	command.Stdout = buf

	err := command.Run()
	if err != nil {
		return nil, errors.Wrapf(err, "scc analyze, execute scc command")
	}

	var sccOutput []processor.LanguageSummary
	err = json.Unmarshal(buf.Bytes(), &sccOutput)
	if err != nil {
		return nil, errors.Wrapf(err, "scc analyze, unmarshal scc output")
	}
	return sccOutput, nil
}
