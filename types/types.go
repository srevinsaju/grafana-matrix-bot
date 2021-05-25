package types

type EvalMatch struct {
    Value int `json:"value"`
    Metrics string `json:"metrix"`
}

type GrafanaWebhook struct {
	ImageUrl string `json:"imageUrl"`
	Message string `json:"message"`
	RuleName string `json:"ruleName"`
	RuleUrl string `json:"ruleUrl"`
	State string `json:"state"`
	Title string `json:"title"`
    EvalMatches []EvalMatch `json:"evalMatches"`
}
