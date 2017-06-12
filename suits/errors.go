package suits

import (
    "errors"
    "fmt"
)

// Private Pseudo-Constants
var (
    eSUITS_ALREADY_INITIALIZED = errors.New("Suits have already been initialized")

)

// Error for Suit property already being created
func errorSuitPropertyAlreadySet(property string) error {
    return errors.New("Suit " + property + " already set")
}

// Error for Suit property already being created
func errorSuitMissingRequiredProperties(property string) error {
    return errors.New("Suit missing required properties (" + property + ").")
}

// Error for Suits not being created yet
func ErrorSuitsNotCreated() error {
    msg := fmt.Sprintf("Suits not created yet %#v", suitsCol)
    return errors.New(msg)
}