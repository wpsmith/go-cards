package cards

import (
    "github.com/wpsmith/go-cards/suits"
    "strings"
)

// CardElement
// Create defined interface for Card Elements
type CardElement interface {
    GetSymbol() string
    GetValue() int
    ToString() string
    ToGlyph() string
    ToHTML() string
    ToANSI() string
    setSymbol() error
    setGlyph() error
    setHTML() error
    setANSI() error
}

// Card
type Card struct {
    ansi        string
    cardType    CardType
    color       string
    glyph       string
    htmlDecimal string
    htmlHex     string
    html        string
    initialized bool
    name        string
    rank        string
    sortValue   int
    suit        suits.Suit
    suitFirst   bool
    symbol      string
    value       int
}

/** PUBLIC METHODS **/
// Search searches for x in a sorted slice of Suits and
// returns the index  as specified by Search.
// The return value is the index to insert x if x is not
// present (it could be len(a)).
// The slice must be sorted in ascending order.
func (c *Card) Search(x string) int {
    return c.SearchBy(cDEFAULT_RANKS, x)
}

// Search by
func (c *Card) SearchBy(dict []string, x string) int {
    return suits.SliceIndex(len(dict), suits.SliceIndexFn(dict, x))
}

// Returns the rank value for the card based on index in slice.
func (c *Card) RankValue() int {
    rankValue := c.Search(c.rank)

    if len(cDEFAULT_RANKS) == rankValue {
        return 0
    }

    return rankValue + 1
}

// Returns the rank value for the card based on index in slice.
func (c *Card) GetRank() string {
    return c.rank
}

// Returns the rank value for the suit of the card.
func (c *Card) SuitRankValue() int {
    suitsCol, _ := suits.GetSuitsCollection()
    return suitsCol.Search(c.suit.GetName())
}

// Gets the card symbol dynamically
func (c *Card) GetSymbol() string {
    if c.suitFirst {
        return c.getSymbolSymbolFirst()
    }
    return c.getSymbolRankFirst()
}

// Gets the card symbol dynamically by type
func (c *Card) GetSymbolBy(stringType string) string {
    switch strings.ToLower(stringType) {
    case "html":
        return c.ToHTMLHex()
    case "hex":
        return c.ToHTMLHex()
    case "decimal":
        return c.ToHTMLDecimal()
    case "ansi":
        return c.ToANSI()
    }
    return c.ToString()
}

// Gets the card value
func (c *Card) GetValue() int {
    return c.value;
}

// Show card with symbol
func (c *Card) ToString() string {
    return c.GetSymbol()
}

// Show card glyph with symbol
func (c *Card) ToGlyph() string {
    return c.glyph
    //return cGLYPH[c.Symbol()]
}

// Show card HTML Deximal with symbol
func (c *Card) ToHTMLDecimal() string {
    return c.htmlDecimal
    //return cDECIMAL[c.Symbol()]
}

// Show card HTML HEX with symbol
func (c *Card) ToHTMLHex() string {
    return c.htmlHex
    //return cHEX[c.Symbol()]
}

// Show card HTML with symbol (Alias for ToHTMLHex)
func (c *Card) ToHTML() string {
    return c.html
    //return c.ToHTMLHex()
}

// Show card in ANSI with symbol
func (c *Card) ToANSI() string {
    var output string
    //switch c.suit.GetColor() {
    switch c.color {
    case suits.BLACK:
        output = suits.ANSI_BLACK
    case suits.RED:
        output = suits.ANSI_RED
    case suits.BLUE:
        output = suits.ANSI_BLUE
    case suits.GREEN:
        output = suits.ANSI_GREEN
    case suits.YELLOW:
        output = suits.ANSI_YELLOW
    }
    output = output + c.GetSymbol() + suits.ANSI_RESET
    return output
}

/** PRIVATE METHODS **/
// Gets card symbol with rank first
func (c *Card) getSymbolRankFirst() string {
    return c.rank + c.suit.GetSymbol()
}

// Gets card symbol with rank symbol
func (c *Card) getSymbolSymbolFirst() string {
    return c.suit.GetSymbol() + c.rank
}

// Sets the card symbol
// @uses errorCardPropertyAlreadySet Returns error for property being set
// @see errors.go
// @uses cDefaultSuitsSymbols
// @see constants.go
func (c *Card) setSymbol(opts CardOptions) error {
    if c.symbol != "" {
        return errorCardPropertyAlreadySet("symbol")
    }

    // Use options symbol setting
    if opts.Symbol != "" {
        c.symbol = opts.Symbol
    } else {
        c.symbol = c.GetSymbol()
    }

    return nil
}

// Sets the card glyph
// @uses errorCardPropertyAlreadySet Returns error for property being set
// @see errors.go
// @uses cHEX
// @see constants.go
func (c *Card) setGlyph(opts CardOptions) error {
    if c.glyph != "" {
        return errorCardPropertyAlreadySet("glyph")
    }

    // Use options symbol setting
    if opts.Glyph != "" {
        c.glyph = opts.Glyph
    } else {
        c.glyph = cGLYPH[c.getSymbolRankFirst()]
    }

    return nil
}

// Sets the card HTML hex
// @uses errorCardPropertyAlreadySet Returns error for property being set
// @see errors.go
// @uses cHEX
// @see constants.go
func (c *Card) setHTMLHex(opts CardOptions) error {
    if c.htmlHex != "" {
        return errorCardPropertyAlreadySet("htmlHex")
    }

    // Use options symbol setting
    if opts.HTMLHex != "" {
        c.htmlHex = opts.HTMLHex
    } else {
        c.htmlHex = cHEX[c.getSymbolRankFirst()]
    }

    if c.htmlHex == "" {
        return errorCardMissingRequiredProperties("HTMLHex")
    }

    return nil
}

// Sets the card HTML Decimal
// @uses errorCardPropertyAlreadySet Returns error for property being set
// @see errors.go
// @uses cHTML
// @see constants.go
func (c *Card) setHTMLDecimal(opts CardOptions) error {
    if c.htmlDecimal != "" {
        return errorCardPropertyAlreadySet("htmlDecimal")
    }

    // Use options symbol setting
    if opts.HTMLDecimal != "" {
        c.htmlDecimal = opts.HTMLDecimal
    } else {
        c.htmlDecimal = cDECIMAL[c.getSymbolRankFirst()]
    }

    if c.htmlDecimal == "" {
        return errorCardMissingRequiredProperties("HTMLDecimal")
    }

    return nil
}

// Sets the card HTML
// @uses errorCardPropertyAlreadySet Returns error for property being set
// @see errors.go
// @uses cHTML
// @see constants.go
func (c *Card) setHTML(opts CardOptions) error {
    if c.html != "" {
        return errorCardPropertyAlreadySet("html")
    }

    // Use options symbol setting
    if opts.HTML != "" {
        c.html = opts.HTML
    } else {
        if c.htmlHex != "" {
            c.html = c.htmlHex
        } else if c.htmlDecimal != "" {
            c.html = c.htmlDecimal
        }
    }

    if c.html == "" {
        return errorCardMissingRequiredProperties("HTML, HTMLHex, HTMLDecimal")
    }

    return nil
}

// Sets the card color
// @uses errorCardPropertyAlreadySet Returns error for property being set
// @see errors.go
func (c *Card) setColor(opts CardOptions) error {
    if c.color != "" {
        return errorCardPropertyAlreadySet("color")
    }

    c.color = c.suit.GetColor()
    if opts.Color != "" {
        c.color = opts.Color
    }

    return nil
}

// Sets the card value
// @uses errorCardPropertyAlreadySet Returns error for property being set
// @see errors.go
func (c *Card) setSortValue() error {
    if c.sortValue != 0 {
        return errorCardPropertyAlreadySet("sortValue")
    }

    c.sortValue = cDEFAULT_CARD_VALUES[c.rank] + c.suit.GetRank()
    return nil
}

// Sets the card value
// @uses errorCardPropertyAlreadySet Returns error for property being set
// @see errors.go
func (c *Card) setValue(opts CardOptions) error {
    if c.value != 0 {
        return errorCardPropertyAlreadySet("value")
    }

    if c.sortValue == 0 {
        c.setSortValue()
    }

    c.value = cDEFAULT_CARD_VALUES[c.rank]
    if opts.Value != 0 {
        c.value = opts.Value
    }

    return nil
}

// Sets the card ANSI
// @uses errorCardPropertyAlreadySet Returns error for property being set
// @see errors.go
func (c *Card) setANSI(opts CardOptions) error {
    if c.ansi != "" {
        return errorCardPropertyAlreadySet("ansi")
    }

    // Use options symbol setting
    if opts.ANSI != "" {
        c.ansi = opts.ANSI
    } else {
        var output string
        switch c.color {
        case suits.BLACK:
            output = suits.ANSI_BLACK
        case suits.RED:
            output = suits.ANSI_RED
        case suits.BLUE:
            output = suits.ANSI_BLUE
        case suits.GREEN:
            output = suits.ANSI_GREEN
        case suits.YELLOW:
            output = suits.ANSI_YELLOW
        case suits.WHITE:
            output = suits.ANSI_WHITE
        case suits.MAGENTA:
            output = suits.ANSI_MAGENTA
        case suits.CYAN:
            output = suits.ANSI_CYAN
        }
        output = output + c.GetSymbol() + suits.ANSI_RESET
        c.ansi = output
    }

    return nil
}

// Determines whether the card belongs to a specific set of card ranks
func (c *Card) isCard( cards []string, rank string) bool {
    return (len(cards) != c.SearchBy(cards, rank))
}

// Suit Options object
type CardOptions struct {
    ANSI        string
    Color       string
    Glyph       string
    HTMLDecimal string
    HTMLHex     string
    HTML        string
    Rank        string
    Suit        suits.Suit
    SuitFirst   bool
    Symbol      string
    Type        CardType
    Value       int
}

func NewCustomCard(opts CardOptions) (*Card, error) {
    card := new(Card)

    if opts.Suit.GetName() == "" {
        return nil, errorCardMissingRequiredProperties("Suit")
    }
    card.suit = opts.Suit

    if opts.Rank == "" {
        return nil, errorCardMissingRequiredProperties("Rank")
    }
    card.rank = opts.Rank

    // Set Suit First
    card.suitFirst = opts.SuitFirst

    // Set Everything
    card.setSortValue()
    card.setValue(opts)
    card.setSymbol(opts)
    card.setColor(opts)
    card.setGlyph(opts)
    card.setHTMLHex(opts)
    card.setHTMLDecimal(opts)
    card.setHTML(opts)
    card.setANSI(opts)

    // Set Card Type
    // Numeric Cards
    if card.isCard(NUMERIC_CARDS, opts.Rank) {
        card.cardType = NUMBER_CARD

        // Face Cards
    } else if card.isCard(FACE_CARDS, opts.Rank) {
        card.cardType = FACE_CARD

        // Ace Card
    } else if card.isCard(ACE_CARDS, opts.Rank) {
        card.cardType = ACE_CARD

        // Joker Cards
    } else if card.isCard(JOKER_CARDS, opts.Rank) {
        card.cardType = JOKER_CARD
    }

    card.initialized = true
    return card, nil
}

func NewCard(suit suits.Suit, rank string) (*Card, error) {
    return NewCustomCard(CardOptions{
        Suit: suit,
        Rank: rank,
    })
}
