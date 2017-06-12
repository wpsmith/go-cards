package suits

// @todo Add additional suits: stars (mix red/black), red crosses, black bullets, green leaves, green eagle, green castles; see https://en.wikipedia.org/wiki/Suit_(cards)#Four-color_suits

// Actual REAL constants
const (
    // Color Constants
    BLACK = "black"
    RED = "red"
    GREEN = "green"
    BLUE = "blue"
    YELLOW = "yellow"
    WHITE = "white"
    MAGENTA = "magenta"
    PURPLE = MAGENTA
    CYAN = "cyan"
    BLUEGREEN = CYAN

    // Defaults
    DEFAULT_TRUMPSUIT = false

    // Default Suit Names
    CLUBS = "Clubs"
    DIAMONDS = "Diamonds"
    HEARTS = "Hearts"
    SPADES = "Spades"

    // @see http://www.lihaoyi.com/post/BuildyourownCommandLinewithANSIescapecodes.html
    ANSI_RESET = "\u001b[0m"
    ANSI_BLACK = "\u001b[30m"
    ANSI_RED = "\u001b[31m"
    ANSI_GREEN = "\u001b[32m"
    ANSI_YELLOW = "\u001b[33m"
    ANSI_BLUE = "\u001b[34m"
    ANSI_WHITE = "\u001b[37m"
    ANSI_MAGENTA = "\u001b[35m"
    ANSI_PURPLE = ANSI_MAGENTA
    ANSI_CYAN = "\u001b[36m"
    ANSI_BLUEGREEN = ANSI_CYAN
)

// Private Pseudo-Constants used by the suits package
var (
    // Default color scheme
    cDefaultColorScheme map[string]string = map[string]string{
        CLUBS: BLACK,
        DIAMONDS: RED,
        HEARTS: RED,
        SPADES: BLACK,
    }

    // Suits
    // @see http://www.fileformat.info/info/unicode/block/miscellaneous_symbols/list.htm
    cDefaultSuitsSymbols map[string]string = map[string]string{
        CLUBS: "♣",
        DIAMONDS: "♦",
        HEARTS: "♥",
        SPADES: "♠",
    }

    // Glyphs, HTML representations
    cGLYPH map[string]string = map[string]string{
        SPADES: "\u2660",
        CLUBS: "\u2663",
        HEARTS: "\u2665",
        DIAMONDS: "\u2666",
    }

    cHEX map[string]string = map[string]string{
        SPADES: "&#x02660",
        CLUBS: "&#x02663",
        HEARTS: "&#x02665",
        DIAMONDS: "&#x02666",
    }

    cDECIMAL map[string]string = map[string]string{
        SPADES: "&#9824",
        CLUBS: "&#9827",
        HEARTS: "&#9829",
        DIAMONDS: "&#9830",
    }

    cHTML map[string]string = map[string]string{
        SPADES: "&spadesuit",
        CLUBS: "&clubsuit",
        HEARTS: "&heartsuit",
        DIAMONDS: "&diamondsuit",
    }
)

// Make Pseudo-Constants Available
func GetSuitsDefaultColorScheme() map[string]string {
    return cDefaultColorScheme
}

func GetSuitsDefaultSymbols() map[string]string {
    return cDefaultSuitsSymbols
}

func GetSuitsGlyphs() map[string]string {
    return cGLYPH
}

func GetSuitsHTMLHex() map[string]string {
    return cHEX
}

func GetSuitsHTMLDecimal() map[string]string {
    return cDECIMAL
}

func GetSuitsHTML() map[string]string {
    return cHTML
}

func GetSuitDefaultColor(suit string) string {
    if suit == SPADES || suit == CLUBS || suit == HEARTS || suit == DIAMONDS {
        return cDefaultColorScheme[suit]
    }
    return ""
}

func GetSuitDefaultSymbol(suit string) string {
    if suit == SPADES || suit == CLUBS || suit == HEARTS || suit == DIAMONDS {
        return cDefaultSuitsSymbols[suit]
    }
    return ""
}

func GetSuitGlyph(suit string) string {
    if suit == SPADES || suit == CLUBS || suit == HEARTS || suit == DIAMONDS {
        return cGLYPH[suit]
    }
    return ""
}

func GetSuitHTMLHex(suit string) string {
    if suit == SPADES || suit == CLUBS || suit == HEARTS || suit == DIAMONDS {
        return cHEX[suit]
    }
    return ""
}

func GetSuitHTMLDecimal(suit string) string {
    if suit == SPADES || suit == CLUBS || suit == HEARTS || suit == DIAMONDS {
        return cDECIMAL[suit]
    }
    return ""
}

func GetSuitHTML(suit string) string {
    if suit == SPADES || suit == CLUBS || suit == HEARTS || suit == DIAMONDS {
        return cHTML[suit]
    }
    return ""
}