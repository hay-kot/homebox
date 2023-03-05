// Package notifier exposes the notifier service for sending external messages.
package notifier

import "github.com/containrrr/shoutrrr"

// Notify is a proxy method for shoutrrr.Send. May be removed if additional
// functionality is _not_ needed.
func Notify(url, message string) error {
	return shoutrrr.Send(url, message)
}
