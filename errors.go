package cards

import (
    "errors"
)

// Private Pseudo-Constants
var (
    eDECK_HAS_NO_CARDS = errors.New("Deck has no cards")
    eDECK_WAS_NOT_CREATED = errors.New("Deck was not created")
)

// Error for Suit property already being created
func errorCardPropertyAlreadySet(property string) error {
    return errors.New("Card " + property + " already set")
}

// Error for Suit property already being created
func errorCardMissingRequiredProperties(property string) error {
    return errors.New("Suit missing required properties (" + property + ").")
}