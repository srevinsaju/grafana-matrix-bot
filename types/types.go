package types


type GrafanaWebhook struct {
	ImageUrl string `json:"imageUrl"`
	Message string `json:"message"`
	RuleName string `json:"ruleName"`
	RuleUrl string `json:"ruleUrl"`
	State string `json:"state"`
	Title string `json:"title"`
}