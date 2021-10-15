package domain

import (
	"canny/model"
	"canny/pkg/firebase"
	"github.com/davecgh/go-spew/spew"
)

func GetAllRules() []model.RuleStruct {
	return firebase.GetAllDocuments()
}

func EvaluateAllRules(rules []model.RuleStruct) {
	for _, rule := range rules {
		for _ , alert :=  range rule.Alerts {
			if EvaluateRule(alert) {
				print("DO SOMETHING AS CONDITION SUCCEEDED")
				spew.Dump(alert, rule.Owner)
				//TODO send email
			}
		}
	}
}
