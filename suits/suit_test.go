package suits

import (
    "testing"
)

/** Test NewSuit **/
type testNewSuit struct {
    opts SuitOptions
    suit *Suit
}

var (
    // Suit Options
    newSuitClubsSuitOptions SuitOptions = SuitOptions{
        Name: CLUBS,
        //Symbol: cDefaultSuitsSymbols[CLUBS],
    }
    newSuitDiamondsSuitOptions SuitOptions = SuitOptions{
        Name: DIAMONDS,
        //Symbol: cDefaultSuitsSymbols[DIAMONDS],
    }
    newSuitHeartsSuitOptions SuitOptions = SuitOptions{
        Name: HEARTS,
        //Symbol: cDefaultSuitsSymbols[HEARTS],
    }
    newSuitSpadesSuitOptions SuitOptions = SuitOptions{
        Color: GREEN,
        Name: SPADES,
        //Symbol: cDefaultSuitsSymbols[SPADES],
        IsTrumpSuit: true,
    }
    newSuitEaglesSuitOptions SuitOptions = SuitOptions{
        Color: GREEN,
        Name: "Eagles",
        Symbol: "?",
        Glyph: "\U0001F985",
        HTMLDecimal: "&#129413",
        HTMLHex: "&#x1f985",
        Rank: 1,
    }
    newSuitFaultyEaglesSuitOptions SuitOptions = SuitOptions{
        Color: GREEN,
        Name: "Eagles",
        Glyph: "\U0001F985",
        HTMLDecimal: "&#129413",
    }

    // Suits
    clubsSuit *Suit = &Suit{
        ansi: ANSI_BLACK + cDefaultSuitsSymbols[CLUBS] + ANSI_RESET,
        color: BLACK,
        glyph: cGLYPH[CLUBS],
        htmlDecimal: cDECIMAL[CLUBS],
        htmlHex: cHEX[CLUBS],
        html: cHTML[CLUBS],
        isTrumpSuit: false,
        name: CLUBS,
        options: newSuitClubsSuitOptions,
        symbol: cDefaultSuitsSymbols[CLUBS],
    }
    diamondsSuit *Suit = &Suit{
        ansi: ANSI_RED + cDefaultSuitsSymbols[DIAMONDS] + ANSI_RESET,
        color: RED,
        glyph: cGLYPH[DIAMONDS],
        htmlDecimal: cDECIMAL[DIAMONDS],
        htmlHex: cHEX[DIAMONDS],
        html: cHTML[DIAMONDS],
        isTrumpSuit: false,
        name: DIAMONDS,
        options: newSuitDiamondsSuitOptions,
        symbol: cDefaultSuitsSymbols[DIAMONDS],
    }
    heartsSuit *Suit = &Suit{
        ansi: ANSI_RED + cDefaultSuitsSymbols[HEARTS] + ANSI_RESET,
        color: RED,
        glyph: cGLYPH[HEARTS],
        htmlDecimal: cDECIMAL[HEARTS],
        htmlHex: cHEX[HEARTS],
        html: cHTML[HEARTS],
        isTrumpSuit: false,
        name: HEARTS,
        options: newSuitHeartsSuitOptions,
        symbol: cDefaultSuitsSymbols[HEARTS],
    }
    spadesSuit *Suit = &Suit{
        ansi: ANSI_GREEN + cDefaultSuitsSymbols[SPADES] + ANSI_RESET,
        color: GREEN,
        glyph: cGLYPH[SPADES],
        htmlDecimal: cDECIMAL[SPADES],
        htmlHex: cHEX[SPADES],
        html: cHTML[SPADES],
        isTrumpSuit: true,
        name: SPADES,
        options: newSuitSpadesSuitOptions,
        symbol: cDefaultSuitsSymbols[SPADES],
    }
    eaglesSuit *Suit = &Suit{
        ansi: ANSI_GREEN + "?" + ANSI_RESET,
        color: GREEN,
        glyph: "\U0001F985",
        htmlDecimal: "&#129413",
        htmlHex: "&#x1f985",
        html: "&#x1f985",
        isTrumpSuit: false,
        name: "Eagles",
        options: newSuitEaglesSuitOptions,
        symbol: "?",
    }


    // Test Array
    testsNewSuit = []testNewSuit{
        {newSuitClubsSuitOptions, clubsSuit},
        {newSuitDiamondsSuitOptions, diamondsSuit},
        {newSuitHeartsSuitOptions, heartsSuit},
        {newSuitSpadesSuitOptions, spadesSuit},
        {newSuitEaglesSuitOptions, eaglesSuit},
    }
)

// Test the Error return for NewSuit
func TestNewSuitError(t *testing.T) {
    var err error
    _, err = NewSuit(newSuitFaultyEaglesSuitOptions)
    if err == nil {
        t.Errorf("For NewSuit error testing, expected " + errorSuitMissingRequiredProperties("Color, Glyph, HTMLHex, HTMLDecimal, Name, Symbol").Error() + ", but got <nil> instead.")
    }

    _, err = NewSuit(SuitOptions{})
    if err == nil {
        t.Errorf("For NewSuit error testing, expected " + errorSuitMissingRequiredProperties("Color, Glyph, HTMLHex, HTMLDecimal, Name, Symbol").Error() + ", but got <nil> instead.")
    }
}

// Test helper function
func doNewSuitTestError(t *testing.T, testName string, valType string, expectedVal interface{}, val interface{}) {
    v, ok := val.(string)
    if ok {
        val = getEmptyStringMaybe(v)
    }
    t.Errorf("For %s, expected " + valType + ", but got " + valType + " instead.", testName, expectedVal, val)
}

// Test helper function
func doSuitTestResults(t *testing.T, pairSuit *Suit, v *Suit, opts SuitOptions, pair testNewSuit) {
    // @todo Check to see if this errors!
    //if !reflect.DeepEqual(v.options, opts) {
    //    t.Errorf("For %s, expected %+v, but got %+v instead.", opts.Name + " Options", opts, pairSuit, v)
    //    //t.Errorf("For %s and %+v, expected %+v, but got %+v instead.", opts.Name + " Options", opts, suit, v)
    //}

    // Check Public Methods
    if pairSuit.name != v.GetName() {
        doNewSuitTestError(t, opts.Name + " GetName", "%s", pairSuit.name, v.GetName())
    }
    if pairSuit.symbol != v.GetSymbol() {
        doNewSuitTestError(t, opts.Name + " GetSymbol", "%s", pairSuit.symbol, v.GetSymbol())
    }
    if pairSuit.color != v.GetColor() {
        doNewSuitTestError(t, opts.Name + " GetColor", "%s", pairSuit.color, v.GetColor())
    }
    if pairSuit.isTrumpSuit != v.IsTrumpSuit() {
        doNewSuitTestError(t, opts.Name + " IsTrumpSuit", "%b", pairSuit.isTrumpSuit, v.IsTrumpSuit())
        //t.Errorf("For %s and %+v, expected %b, but got %b instead.", opts.Name, opts, pairSuit.isTrumpSuit, v.IsTrumpSuit())
    }
    if pairSuit.symbol != v.ToString() {
        doNewSuitTestError(t, opts.Name + " ToString", "%s", pairSuit.symbol, v.ToString())
    }
    if pairSuit.glyph != v.ToGlyph() {
        doNewSuitTestError(t, opts.Name + " ToGlyph", "%s", pairSuit.glyph, v.ToGlyph())
    }
    if pairSuit.htmlDecimal != v.ToHTMLDecimal() {
        doNewSuitTestError(t, opts.Name + " ToHTMLDecimal", "%s", pairSuit.htmlDecimal, v.ToHTMLDecimal())
    }
    if pairSuit.htmlHex != v.ToHTMLHex() {
        doNewSuitTestError(t, opts.Name + " ToHTMLHex", "%s", pairSuit.htmlHex, v.ToHTMLHex())
    }
    if pairSuit.html != v.ToHTML() {
        doNewSuitTestError(t, opts.Name + " ToHTML", "%s", pairSuit.html, v.ToHTML())
    }
    if pairSuit.ansi != v.ToANSI() {
        doNewSuitTestError(t, opts.Name + " ToANSI", "%s", pairSuit.ansi, v.ToANSI())
    }
}

func TestNewSuit(t *testing.T) {
    for _, pair := range testsNewSuit {
        suit, err := NewSuit(pair.opts)
        if err != nil {
            t.Errorf("For NewSuit, expected received error %#v\n", err.Error())
        }
        doSuitTestResults(t, pair.suit, suit, pair.opts, pair)
    }
}


func TestNewSuitDefaultSuitColor(t *testing.T) {
    for _, pair := range testsNewSuit {
        suit, _ := NewSuit(pair.opts)
        doSuitTestResults(t, pair.suit, suit, pair.opts, pair)
    }
}

func TestNewSuitRanking(t *testing.T) {

}

func TestNewSuitSpadesTrumpSuit(t *testing.T) {

}