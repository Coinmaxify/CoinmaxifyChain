package ibftswitch

import (
	"bytes"
	"fmt"
	"github.com/Coinmaxify/Coinmaxify/command/helper"
	"github.com/Coinmaxify/Coinmaxify/consensus/ibft"
	"github.com/Coinmaxify/Coinmaxify/helper/common"
)

type IBFTSwitchResult struct {
	Chain             string             `json:"chain"`
	Type              ibft.MechanismType `json:"type"`
	From              common.JSONNumber  `json:"from"`
	Deployment        *common.JSONNumber `json:"deployment,omitempty"`
	MaxValidatorCount common.JSONNumber  `json:"maxValidatorCount"`
	MinValidatorCount common.JSONNumber  `json:"minValidatorCount"`
}

func (r *IBFTSwitchResult) GetOutput() string {
	var buffer bytes.Buffer

	buffer.WriteString("\n[NEW IBFT FORK]\n")

	outputs := []string{
		fmt.Sprintf("Chain|%s", r.Chain),
		fmt.Sprintf("Type|%s", r.Type),
	}
	if r.Deployment != nil {
		outputs = append(outputs, fmt.Sprintf("Deployment|%d", r.Deployment.Value))
	}

	outputs = append(outputs, fmt.Sprintf("From|%d", r.From.Value))
	outputs = append(outputs, fmt.Sprintf("MaxValidatorCount|%d", r.MaxValidatorCount.Value))
	outputs = append(outputs, fmt.Sprintf("MinValidatorCount|%d", r.MinValidatorCount.Value))

	buffer.WriteString(helper.FormatKV(outputs))
	buffer.WriteString("\n")

	return buffer.String()
}
