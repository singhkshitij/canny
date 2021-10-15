package domain

import (
	"canny/model"
	"canny/pkg/firebase"
)

func GetAllRules() []model.RuleStruct {
	return firebase.GetAllDocuments()
}

func EvaluateAllRules(rules []model.RuleStruct) {
	for _, rule := range rules {
		for _ , alert :=  range rule.Alerts {
			if EvaluateRule(alert) {
				print("DO SOMETHING AS CONDITION SUCCEEDED")
				//spew.Dump(rule)
			}
		}
	}
}
