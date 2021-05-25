package matrix

import (
	"fmt"
	"github.com/srevinsaju/grafana-matrix-bot/types"
	"maunium.net/go/mautrix"
	"maunium.net/go/mautrix/event"
	"maunium.net/go/mautrix/format"
	"maunium.net/go/mautrix/id"
)

func SendMessage(client *mautrix.Client, channel string, webhook *types.GrafanaWebhook) error {
	var strMessage string
    metrics := ""
    if len(webhook.EvalMatches) > 0 {
        for i := range webhook.EvalMatches {
            metric := webhook.EvalMatches[i].Metrics
            value := webhook.EvalMatches[i].Value
            metrics = fmt.Sprintf("%s\n%s = %d", metrics, metric, value)
        }
    }

	strMessage = fmt.Sprintf(`**%s**
%s
[%s](%s)

%s
`, webhook.Title, webhook.Message, webhook.RuleName, webhook.RuleUrl, metrics )

	content := format.RenderMarkdown(strMessage, true, true)
	_, err := client.SendMessageEvent(id.RoomID(channel), event.EventMessage, &content)
	if err != nil {
		logger.Warnf("Failed to send message to matrix server")
		return err
	}
	return nil
}


func SendPhoto(client *mautrix.Client, channel string, webhook *types.GrafanaWebhook) error {
	link, err := client.UploadLink(webhook.ImageUrl)
	if err != nil {
		logger.Warnf("Couldn't upload photo to matrix homeserver")
		return err
	}

	_, err = client.SendImage(id.RoomID(channel), "", link.ContentURI)
	if err != nil {
		logger.Warnf("Failed to send image to matrix")
		return err
	}
	return nil
}
