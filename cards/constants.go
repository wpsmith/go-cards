package cards

var (
    // Card Ranks/Values
    cDEFAULT_RANKS []string = []string{"2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K", "A"}
    cDEFAULT_CARD_VALUES map[string]int = map[string]int{
        "2": 2,
        "3": 3,
        "4": 4,
        "5": 5,
        "6": 6,
        "7": 7,
        "8": 8,
        "9": 9,
        "10": 10,
        "J": 11, // Jack
        "C": 12, // Knight
        "Q": 13, // Queen
        "K": 14, // King
        "A": 15, // Ace
        "F": 16, // Fool
        "T": 17, // Trump
        "R": 18, // Joker

    }
    NUMERIC_CARDS []string = []string{"2", "3", "4", "5", "6", "7", "8", "9", "10"}
    FACE_CARDS []string = []string{"J", "Q", "K", "C"}
    ACE_CARDS []string = []string{"A"}
    JOKER_CARDS []string = []string{"T", "F", "R"}

    // Display Values
    // @see http://www.fileformat.info/info/unicode/block/playing_cards/images.htm
    cGLYPH map[string]string = map[string]string{
        // SPADES            HEARTS               DIAMONDS             CLUBS
        "A♠":  "\U0001f0a1", "A♥":  "\U0001f0b1", "A♦":  "\U0001f0c1", "A♣":  "\U0001f0d1",
        "2♠":  "\U0001f0a2", "2♥":  "\U0001f0b2", "2♦":  "\U0001f0c2", "2♣":  "\U0001f0d2",
        "3♠":  "\U0001f0a3", "3♥":  "\U0001f0b3", "3♦":  "\U0001f0c3", "3♣":  "\U0001f0d3",
        "4♠":  "\U0001f0a4", "4♥":  "\U0001f0b4", "4♦":  "\U0001f0c4", "4♣":  "\U0001f0d4",
        "5♠":  "\U0001f0a5", "5♥":  "\U0001f0b5", "5♦":  "\U0001f0c5", "5♣":  "\U0001f0d5",
        "6♠":  "\U0001f0a6", "6♥":  "\U0001f0b6", "6♦":  "\U0001f0c6", "6♣":  "\U0001f0d6",
        "7♠":  "\U0001f0a7", "7♥":  "\U0001f0b7", "7♦":  "\U0001f0c7", "7♣":  "\U0001f0d7",
        "8♠":  "\U0001f0a8", "8♥":  "\U0001f0b8", "8♦":  "\U0001f0c8", "8♣":  "\U0001f0d8",
        "9♠":  "\U0001f0a9", "9♥":  "\U0001f0b9", "9♦":  "\U0001f0c9", "9♣":  "\U0001f0d9",
        "10♠": "\U0001f0aa", "10♥": "\U0001f0ba", "10♦": "\U0001f0ca", "10♣": "\U0001f0da",
        "J♠":  "\U0001f0ab", "J♥":  "\U0001f0bb", "J♦":  "\U0001f0cb", "J♣":  "\U0001f0db",
        "C♠":  "\U0001f0ac", "C♥":  "\U0001f0bc", "C♦":  "\U0001f0cc", "C♣":  "\U0001f0dc",
        "Q♠":  "\U0001f0ad", "Q♥":  "\U0001f0bd", "Q♦":  "\U0001f0cd", "Q♣":  "\U0001f0dd",
        "K♠":  "\U0001f0ae", "K♥":  "\U0001f0be", "K♦":  "\U0001f0ce", "K♣":  "\U0001f0de",

        // Other Cards
        "back": "\U0001f0a0",

        // Other Cards, Jokers, Fool, and Trumps Cards
        "redjoker":   "\U0001f0bf",
        "blackjoker": "\U0001f0cf",
        "whitejoker": "\U0001f0df",
        "fool":       "\U0001f0e0",
        "trump1":     "\U0001f0e1",
        "trump2":     "\U0001f0e2",
        "trump3":     "\U0001f0e3",
        "trump4":     "\U0001f0e4",
        "trump5":     "\U0001f0e5",
        "trump6":     "\U0001f0e6",
        "trump7":     "\U0001f0e1",
        "trump8":     "\U0001f0e8",
        "trump9":     "\U0001f0e9",
        "trump10":    "\U0001f0ea",
        "trump11":    "\U0001f0eb",
        "trump12":    "\U0001f0ec",
        "trump13":    "\U0001f0ed",
        "trump14":    "\U0001f0ee",
        "trump15":    "\U0001f0ef",
        "trump16":    "\U0001f0f0",
        "trump17":    "\U0001f0f1",
        "trump18":    "\U0001f0f2",
        "trump19":    "\U0001f0f3",
        "trump20":    "\U0001f0f4",
        "trump21":    "\U0001f0f5",
    }
    cHEX map[string]string = map[string]string{
        // SPADES            HEARTS               DIAMONDS             CLUBS
        "A♠":  "&#x1f0a1", "A♥":  "&#x1f0b1", "A♦":  "&#x1f0c1", "A♣":  "&#x1f0d1",
        "2♠":  "&#x1f0a2", "2♥":  "&#x1f0b2", "2♦":  "&#x1f0c2", "2♣":  "&#x1f0d2",
        "3♠":  "&#x1f0a3", "3♥":  "&#x1f0b3", "3♦":  "&#x1f0c3", "3♣":  "&#x1f0d3",
        "4♠":  "&#x1f0a4", "4♥":  "&#x1f0b4", "4♦":  "&#x1f0c4", "4♣":  "&#x1f0d4",
        "5♠":  "&#x1f0a5", "5♥":  "&#x1f0b5", "5♦":  "&#x1f0c5", "5♣":  "&#x1f0d5",
        "6♠":  "&#x1f0a6", "6♥":  "&#x1f0b6", "6♦":  "&#x1f0c6", "6♣":  "&#x1f0d6",
        "7♠":  "&#x1f0a7", "7♥":  "&#x1f0b7", "7♦":  "&#x1f0c7", "7♣":  "&#x1f0d7",
        "8♠":  "&#x1f0a8", "8♥":  "&#x1f0b8", "8♦":  "&#x1f0c8", "8♣":  "&#x1f0d8",
        "9♠":  "&#x1f0a9", "9♥":  "&#x1f0b9", "9♦":  "&#x1f0c9", "9♣":  "&#x1f0d9",
        "10♠": "&#x1f0aa", "10♥": "&#x1f0ba", "10♦": "&#x1f0ca", "10♣": "&#x1f0da",
        "J♠":  "&#x1f0ab", "J♥":  "&#x1f0bb", "J♦":  "&#x1f0cb", "J♣":  "&#x1f0db",
        "C♠":  "&#x1f0ac", "C♥":  "&#x1f0bc", "C♦":  "&#x1f0cc", "C♣":  "&#x1f0dc",
        "Q♠":  "&#x1f0ad", "Q♥":  "&#x1f0bd", "Q♦":  "&#x1f0cd", "Q♣":  "&#x1f0dd",
        "K♠":  "&#x1f0ae", "K♥":  "&#x1f0be", "K♦":  "&#x1f0ce", "K♣":  "&#x1f0de",

        // Other Cards, Jokers, Fool, and Trumps Cards
        "back":       "&#x1f0a0",
        "redjoker":   "&#x1f0bf",
        "blackjoker": "&#x1f0cf",
        "whitejoker": "&#x1f0df",
        "fool":       "&#x1f0e0",
        "trump1":     "&#x1f0e1",
        "trump2":     "&#x1f0e2",
        "trump3":     "&#x1f0e3",
        "trump4":     "&#x1f0e4",
        "trump5":     "&#x1f0e5",
        "trump6":     "&#x1f0e6",
        "trump7":     "&#x1f0e7",
        "trump8":     "&#x1f0e8",
        "trump9":     "&#x1f0e9",
        "trump10":    "&#x1f0ea",
        "trump11":    "&#x1f0eb",
        "trump12":    "&#x1f0ec",
        "trump13":    "&#x1f0ed",
        "trump14":    "&#x1f0ee",
        "trump15":    "&#x1f0ef",
        "trump16":    "&#x1f0f0",
        "trump17":    "&#x1f0f1",
        "trump18":    "&#x1f0f2",
        "trump19":    "&#x1f0f3",
        "trump20":    "&#x1f0f4",
        "trump21":    "&#x1f0f5",
    }
    cDECIMAL map[string]string = map[string]string{
        // SPADES          HEARTS             DIAMONDS           CLUBS
        "A♠":  "&#127137", "A♥":  "&#127153", "A♦":  "&#127169", "A♣":  "&#127185",
        "2♠":  "&#127138", "2♥":  "&#127154", "2♦":  "&#127170", "2♣":  "&#127186",
        "3♠":  "&#127139", "3♥":  "&#127155", "3♦":  "&#127171", "3♣":  "&#127187",
        "4♠":  "&#127140", "4♥":  "&#127156", "4♦":  "&#127172", "4♣":  "&#127188",
        "5♠":  "&#127141", "5♥":  "&#127157", "5♦":  "&#127173", "5♣":  "&#127189",
        "6♠":  "&#127142", "6♥":  "&#127158", "6♦":  "&#127174", "6♣":  "&#127190",
        "7♠":  "&#127143", "7♥":  "&#127159", "7♦":  "&#127175", "7♣":  "&#127191",
        "8♠":  "&#127144", "8♥":  "&#127160", "8♦":  "&#127176", "8♣":  "&#127192",
        "9♠":  "&#127145", "9♥":  "&#127161", "9♦":  "&#127177", "9♣":  "&#127193",
        "10♠": "&#127146", "10♥": "&#127162", "10♦": "&#127178", "10♣": "&#127194",
        "J♠":  "&#127147", "J♥":  "&#127163", "J♦":  "&#127179", "J♣":  "&#127195",
        "C♠":  "&#127148", "C♥":  "&#127164", "C♦":  "&#127180", "C♣":  "&#127196",
        "Q♠":  "&#127149", "Q♥":  "&#127165", "Q♦":  "&#127181", "Q♣":  "&#127197",
        "K♠":  "&#127150", "K♥":  "&#127166", "K♦":  "&#127182", "K♣":  "&#127198",

        // Other Cards, Jokers, Fool, and Trumps Cards
        "back":       "&#127136",
        "redjoker":   "&#127167",
        "blackjoker": "&#127183",
        "whitejoker": "&#127199",
        "fool":       "&#127200",
        "trump1":     "&#127201",
        "trump2":     "&#127202",
        "trump3":     "&#127203",
        "trump4":     "&#127204",
        "trump5":     "&#127205",
        "trump6":     "&#127206",
        "trump7":     "&#127207",
        "trump8":     "&#127208",
        "trump9":     "&#127209",
        "trump10":    "&#127210",
        "trump11":    "&#127211",
        "trump12":    "&#127212",
        "trump13":    "&#127213",
        "trump14":    "&#127214",
        "trump15":    "&#127215",
        "trump16":    "&#127216",
        "trump17":    "&#127217",
        "trump18":    "&#127218",
        "trump19":    "&#127219",
        "trump20":    "&#127220",
        "trump21":    "&#127221",
    }
)

// CardType the type of card
type CardType int

const (
    _ CardType = iota
    // NumberCard a number card type
    NUMBER_CARD
    // FaceCard a face card type
    FACE_CARD
    // AceCard an ace card type
    ACE_CARD
    // JokerCard a joker card type
    JOKER_CARD
)

// Make Pseudo-Constants Available
func GetDefaultRanks() []string {
    return cDEFAULT_RANKS
}

func GetCardsGlyphs() map[string]string {
    return cGLYPH
}

func GetCardsHTMLHex() map[string]string {
    return cHEX
}

func GetCardsHTMLDecimal() map[string]string {
    return cDECIMAL
}

func GetCardsHTML() map[string]string {
    return GetCardsHTMLHex()
}

func GetCardGlyph(cardSymbol string) string {
    return cGLYPH[cardSymbol]
}

func GetCardHTMLHex(cardSymbol string) string {
    return cHEX[cardSymbol]
}

func GetCardHTMLDecimal(cardSymbol string) string {
    return cDECIMAL[cardSymbol]
}

func GetCardHTML(cardSymbol string) string {
    return GetCardHTMLHex(cardSymbol)
}