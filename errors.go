package gocards

import (
    "errors"
)

// Private Pseudo-Constants
var (
    eDECK_HAS_NO_CARDS = errors.New("Deck has no cards")
    eDECK_WAS_NOT_CREATED = errors.New("Deck was not created")
)