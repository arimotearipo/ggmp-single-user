package utils

import (
	"time"

	"github.com/atotto/clipboard"
)

// EraseClipboard clears the clipboard
// s: string to write to clipboard
// d: duration to clear clipboard
func ClipboardWriteAndErase(s string, d time.Duration) error {
	err := clipboard.WriteAll(s)
	if err != nil {
		return err
	}

	duration := d * time.Second

	time.AfterFunc(duration, func() {
		err = clipboard.WriteAll("")
		if err != nil {
			return
		}
	})

	return nil
}
