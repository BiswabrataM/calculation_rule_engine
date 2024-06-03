package eng

import (
	"fmt"
	"strconv"

	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type Action struct {
	Operator string
	Value    string
}

type Condition struct {
	Parameter string
	Operator  string
	Value     Action
}

type Reward struct {
	ID         string
	Label      string
	Param      string // New field to specify the parameter to be modified
	ActionType string
	Value      Action
}

type Eligibility struct {
	ID           string
	Label        string
	Param        string // New field to specify the parameter to be modified
	OperatorType string
	Value        Action
}

type Promotion struct {
	ID           string
	Label        string
	Description  string
	Elgibilities []Eligibility
	Rewards      []Reward
}

type Rate struct {
	RuleID       string
	Label        string
	Description  string
	Elgibilities []Eligibility
	Rewards      []Reward
}

func applyAction(source string, action Action) float64 {
	sourceVal, _ := strconv.ParseFloat(source, 64)
	actionVal, _ := strconv.ParseFloat(action.Value, 64)

	switch action.Operator {
	case "ADD":
		return sourceVal + actionVal
	case "SUB":
		return sourceVal - actionVal
	case "MULPER":
		return sourceVal * actionVal
	case "OVERRIDE":
		return actionVal
	default:
		return sourceVal
	}
}

func applyTnC(source string, operator string, target string) bool {
	sourceStr := fmt.Sprintf("%s", source)
	targetStr := fmt.Sprintf("%s", target)

	sourceVal, errSource := strconv.ParseFloat(sourceStr, 64)
	targetVal, errTarget := strconv.ParseFloat(targetStr, 64)

	if operator == "LTE" || operator == "GTE" {
		if errSource != nil || errTarget != nil {
			return false
		}

		switch operator {
		case "LTE":
			return sourceVal <= targetVal
		case "GTE":
			return sourceVal >= targetVal
		}
	} else if operator == "EQ" {
		return sourceStr == targetStr
	}
	return false
}

func init() {
	gjson.AddModifier("min", func(json, arg string) string {
		values := gjson.Parse(json).Array()
		if len(values) == 0 {
			return ""
		}

		minValue := values[0].Float()
		minIdx := 0
		for idx, value := range values[1:] {
			if value.Float() < minValue {
				minValue = value.Float()
				minIdx = idx + 1
			}
		}
		return fmt.Sprintf("%d", minIdx)
	})
}

func ExecuteRateRules(inputJSON string, rates []Rate) string {
	for _, rate := range rates {
		eligible := true
		for _, elg := range rate.Elgibilities {
			paramResult := gjson.Get(inputJSON, elg.Param)
			param_val_is_array := paramResult.IsArray()
			if param_val_is_array {
				for _, val := range paramResult.Array() {
					if !applyTnC(val.String(), elg.OperatorType, elg.Value.Value) {
						eligible = false
						break
					}
				}
			} else {
				if !applyTnC(paramResult.String(), elg.OperatorType, elg.Value.Value) {
					eligible = false
					break
				}
			}
		}
		if eligible {
			for _, reward := range rate.Rewards {
				paramResult := gjson.Get(inputJSON, reward.Param)
				newValue := applyAction(paramResult.String(), reward.Value)
				inputJSON, _ = sjson.Set(inputJSON, reward.Param, newValue)
			}
			fmt.Printf("\nRate rule eligible and applied: ruleId: %s [label: %s]", rate.RuleID, rate.Label)

		} else {
			fmt.Printf("\nRate rule skipped: ruleId: %s [label: %s]", rate.RuleID, rate.Label)
		}
	}

	return string(inputJSON)

}

func ExecuteCuoponRules(inputJSON string, promo Promotion) (bool, string) {
	eligible := true
	for _, elg := range promo.Elgibilities {
		paramResult := gjson.Get(inputJSON, elg.Param)
		param_val_is_array := paramResult.IsArray()
		if param_val_is_array {
			for _, val := range paramResult.Array() {
				if !applyTnC(val.String(), elg.OperatorType, elg.Value.Value) {
					eligible = false
					break
				}
			}
		} else {
			if !applyTnC(paramResult.String(), elg.OperatorType, elg.Value.Value) {
				eligible = false
				break
			}
		}
	}

	if eligible {
		for _, reward := range promo.Rewards {
			paramValue := gjson.Get(inputJSON, reward.Param)
			newValue := applyAction(paramValue.String(), reward.Value)
			inputJSON, _ = sjson.Set(inputJSON, reward.Param, newValue)
		}

	}

	return eligible, string(inputJSON)
}
