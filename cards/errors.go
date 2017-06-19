package cards

import (
    "errors"
)

// Error for Suit property already being created
func errorCardPropertyAlreadySet(property string) error {
    return errors.New("Card " + property + " already set")
}

// Error for Suit property already being created
func errorCardMissingRequiredProperties(property string) error {
    return errors.New("Suit missing required properties (" + property + ").")
}