package fsm

import (
	"NewYushinBot/bot/userdata"
)

const (
	None userdata.UserStatus = iota

	// show
	ShowNotifier

	// add
	AddingNotifierTitle
	AddingNotifierYaml

	// update
	UpdateNotifierTitleStart
	UpdateNotifierYamlStart
	UpdateNotifierTitle
	UpdateNotifierYaml

	// delete
	DeleteNotifier
)

func GetActionString(s userdata.UserStatus) string {
	switch s {
	case None:
		return "Nothing"
	case ShowNotifier:
		return "Send the notifier's title you want to show"
	case AddingNotifierTitle:
		return "Send a new notifier's title you want to add"
	case AddingNotifierYaml:
		return "Send a new notifier's yaml config you want to add"
	case UpdateNotifierTitleStart:
		return "Send the notifier's title you want to update title"
	case UpdateNotifierYamlStart:
		return "Send the notifier's title you want to update config"
	case UpdateNotifierTitle:
		return "Send a new notifier's title you want to update"
	case UpdateNotifierYaml:
		return "Send a new notifier's yaml config you want to update"
	case DeleteNotifier:
		return "Send the notifier's title you want to delete"
	}
	return "Unknown"
}
