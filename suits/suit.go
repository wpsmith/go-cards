package suits

// Suit Type
type Suit struct {
    ansi        string
    color       string      // Color
    glyph       string
    htmlDecimal string
    htmlHex     string
    html        string
    isTrumpSuit bool        // Trump?
    name        string      // Suit name
    options     SuitOptions // Raw options
    rank        int         // Suit rank
    symbol      string      // Unicode symbol
}

/** PUBLIC SUIT METHODS **/
// Gets the name from private property
func (s *Suit) GetName() string {
    return s.name
}

// Gets the symbol from private property
func (s *Suit) GetSymbol() string {
    return s.symbol
}

// Gets the color from private property
// @uses maybeSetDefaultColor
func (s *Suit) GetColor() string {
    return s.color
}

// Gets the rank from private property
// Sets rank to 1 if rank is not set
func (s *Suit) GetRank() int {
    return s.rank
}

// Gets whether trump suit from private property
func (s *Suit) IsTrumpSuit() bool {
    return s.isTrumpSuit
}

// To String
func (s *Suit) ToString() string {
    return s.symbol
}

// Show card glyph with symbol
// @uses cGLYPH Constant Slice
// @see constants.go
func (s *Suit) ToGlyph() string {
    return s.glyph
}

// Show card HTML Decimal with symbol
// @uses cDECIMAL Constant Slice
// @see constants.go
func (s *Suit) ToHTMLDecimal() string {
    return s.htmlDecimal
}

// Show card HTML HEX with symbol
// @uses cHEX Constant Slice
// @see constants.go
func (s *Suit) ToHTMLHex() string {
    return s.htmlHex
}

// Show card HTML HEX with symbol
// @uses cHTML Constant Slice
// @see constants.go
func (s *Suit) ToHTML() string {
    return s.html
}

// Show card in ANSI with symbol
func (s *Suit) ToANSI() string {
    return s.ansi
}

/** PRIVATE SUIT METHODS **/
// Sets the suit color
// @uses errorSuitPropertyAlreadySet Returns error for property being set
// @see errors.go
// @uses cDefaultColorScheme
// @see constants.go
func (s *Suit) setDefaultColor() error {
    if s.color != "" {
        return errorSuitPropertyAlreadySet("color")
    }

    // Use options color setting
    if s.options.Color != "" {
        s.color = s.options.Color
    } else {
        s.color = cDefaultColorScheme[s.name]
    }

    return nil
}

// Sets the suit symbol
// @uses errorSuitPropertyAlreadySet Returns error for property being set
// @see errors.go
// @uses cDefaultSuitsSymbols
// @see constants.go
func (s *Suit) setSymbol() error {
    if s.symbol != "" {
        return errorSuitPropertyAlreadySet("symbol")
    }

    // Use options symbol setting
    if s.options.Symbol != "" {
        s.symbol = s.options.Symbol
    } else {
        s.symbol = cDefaultSuitsSymbols[s.name]
    }

    return nil
}

// Sets the suit glyph
// @uses errorSuitPropertyAlreadySet Returns error for property being set
// @see errors.go
// @uses cHEX
// @see constants.go
func (s *Suit) setGlyph() error {
    if s.glyph != "" {
        return errorSuitPropertyAlreadySet("glyph")
    }

    // Use options symbol setting
    if s.options.Glyph != "" {
        s.glyph = s.options.Glyph
    } else {
        s.glyph = cGLYPH[s.name]
    }

    return nil
}

// Sets the suit HTML hex
// @uses errorSuitPropertyAlreadySet Returns error for property being set
// @see errors.go
// @uses cHEX
// @see constants.go
func (s *Suit) setHTMLHex() error {
    if s.htmlHex != "" {
        return errorSuitPropertyAlreadySet("htmlHex")
    }

    // Use options symbol setting
    if s.options.HTMLHex != "" {
        s.htmlHex = s.options.HTMLHex
    } else {
        s.htmlHex = cHEX[s.name]
    }

    return nil
}

// Sets the suit HTML Decimal
// @uses errorSuitPropertyAlreadySet Returns error for property being set
// @see errors.go
// @uses cHTML
// @see constants.go
func (s *Suit) setHTMLDecimal() error {
    if s.htmlDecimal != "" {
        return errorSuitPropertyAlreadySet("htmlDecimal")
    }

    // Use options symbol setting
    if s.options.HTMLDecimal != "" {
        s.htmlDecimal = s.options.HTMLDecimal
    } else {
        s.htmlDecimal = cDECIMAL[s.name]
    }

    return nil
}

// Sets the suit HTML
// @uses errorSuitPropertyAlreadySet Returns error for property being set
// @see errors.go
// @uses cHTML
// @see constants.go
func (s *Suit) setHTML() error {
    if s.html != "" {
        return errorSuitPropertyAlreadySet("html")
    }

    // Use options symbol setting
    if s.options.HTML != "" {
        s.html = s.options.HTML
    } else {
        s.html = cHTML[s.name]

        if s.html == "" {
            if s.htmlHex != "" {
                s.html = s.ToHTMLHex()
            } else if s.htmlDecimal != "" {
                s.html = s.ToHTMLDecimal()
            }
        }
    }

    return nil
}

// Sets the suit ANSI
// @uses errorSuitPropertyAlreadySet Returns error for property being set
// @see errors.go
func (s *Suit) setANSI() error {
    if s.ansi != "" {
        return errorSuitPropertyAlreadySet("ansi")
    }

    // Use options symbol setting
    if s.options.ANSI != "" {
        s.ansi = s.options.ANSI
    } else {
        var output string
        switch s.color {
        case BLACK:
            output = ANSI_BLACK
        case RED:
            output = ANSI_RED
        case BLUE:
            output = ANSI_BLUE
        case GREEN:
            output = ANSI_GREEN
        case YELLOW:
            output = ANSI_YELLOW
        case WHITE:
            output = ANSI_WHITE
        case MAGENTA:
            output = ANSI_MAGENTA
        case CYAN:
            output = ANSI_CYAN
        }
        output = output + s.symbol + ANSI_RESET
        s.ansi = output
    }

    return nil
}

// Sets the suit rank
// @uses errorSuitPropertyAlreadySet Returns error for property being set
// @see errors.go
func (s *Suit) setRanking() error {
    if s.rank != 0 {
        return errorSuitPropertyAlreadySet("rank")
    }

    if s.options.Rank != 0 {
        s.rank = s.options.Rank
    } else if s.options.Rank == 0 {
        s.rank = 1
    }

    return nil
}

// Sets the suit isTrumpSuit
// @uses DEFAULT_TRUMPSUIT
// @see constants.go
func (s *Suit) setTrumpSuit() error {
    s.isTrumpSuit = DEFAULT_TRUMPSUIT
    if DEFAULT_TRUMPSUIT != s.options.IsTrumpSuit {
        s.isTrumpSuit = s.options.IsTrumpSuit
    }

    return nil
}

// Suit Options object
type SuitOptions struct {
    ANSI        string
    Color       string
    Glyph       string
    HTMLDecimal string
    HTMLHex     string
    HTML        string
    IsTrumpSuit bool
    Name        string
    Rank        int
    Symbol      string
}

// Create the Suit
func newSuit(opts SuitOptions) (*Suit, error) {
    if opts.Name == "" {
        return nil, errorSuitMissingRequiredProperties("Name")
    }

    s := new(Suit)

    // Set Options
    s.options = opts

    // Set Name
    s.name = opts.Name

    // Set Properties
    s.setSymbol()
    s.setGlyph()
    s.setDefaultColor()
    s.setHTMLHex()
    s.setHTMLDecimal()
    s.setHTML()
    s.setANSI()
    s.setRanking()
    s.setTrumpSuit()

    return s, nil
}

// Creates a new Suit based on options
func NewSuit(opts SuitOptions) (*Suit, error) {
    if initialized {
        return nil, eSUITS_ALREADY_INITIALIZED
    }

    switch opts.Name {
    case "":
        return nil, errorSuitMissingRequiredProperties("Name")
    case CLUBS, DIAMONDS, HEARTS, SPADES:
        return newSuit(opts)
    default:
        if (opts.Name == "" || opts.Color == "" || opts.Glyph == "" || opts.HTMLDecimal == "" || opts.HTMLHex == "" || opts.Symbol == "") {
            return nil, errorSuitMissingRequiredProperties("Color, Glyph, HTMLHex, HTMLDecimal, Name, Symbol")
        }
        return newSuit(opts)
    }
}