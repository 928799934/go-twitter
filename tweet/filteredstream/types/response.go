package types

import "github.com/928799934/go-twitter/resources"

type ListRulesOutput struct {
	Data   []resources.FilterdStreamRule `json:"data"`
	Meta   resources.ListSearchStreamRulesMeta
	Errors []resources.PartialError `json:"errors"`
}

func (r *ListRulesOutput) HasPartialError() bool {
	return len(r.Errors) != 0
}

type CreateRulesOutput struct {
	Data   []resources.FilterdStreamRule `json:"data"`
	Meta   resources.CreateSearchStreamRulesMeta
	Errors []resources.PartialError `json:"errors"`
}

func (r *CreateRulesOutput) HasPartialError() bool {
	return len(r.Errors) != 0
}

type DeleteRulesOutput struct {
	Data   []resources.FilterdStreamRule `json:"data"`
	Meta   resources.DeleteSearchStreamRulesMeta
	Errors []resources.PartialError `json:"errors"`
}

func (r *DeleteRulesOutput) HasPartialError() bool {
	return len(r.Errors) != 0
}

type SearchStreamMatchedRule struct {
	Tag *string `json:"tag"`
	ID  *string `json:"id"`
}

type SearchStreamOutput struct {
	Data     resources.Tweet `json:"data"`
	Includes struct {
		Users  []resources.User  `json:"users,omitempty"`
		Tweets []resources.Tweet `json:"tweets,omitempty"`
		Places []resources.Place `json:"places,omitempty"`
		Media  []resources.Media `json:"media,omitempty"`
		Polls  []resources.Poll  `json:"polls,omitempty"`
	} `json:"includes,omitempty"`
	MatchingRules []SearchStreamMatchedRule `json:"matching_rules,omitempty"`
	Errors        []resources.PartialError  `json:"errors,omitempty"`
}

func (r *SearchStreamOutput) HasPartialError() bool {
	return len(r.Errors) != 0
}
