package suits

import (
    "testing"
    "reflect"
)

type testGlobalSuitsOptions struct {
    testName   string
    options    GlobalSuitsOptions
    collection *suitsCollection
}

// Test DestroySuits: sets suitsCol "global" variable to nil
func TestDestroySuits(t *testing.T) {
    var testOpts GlobalSuitsOptions = GlobalSuitsOptions{}

    suitsCol = nil
    _, err := CreateSuits(testOpts)
    if err != nil {
        t.Errorf("Error occurred: %s", err.Error())
    }

    DestroySuits()
    if !reflect.DeepEqual(suitsCol, &suitsCollection{}) {
        t.Errorf("For DestroySuits 1, expected %+v, but got %+v instead.", &suitsCollection{}, suitsCol)
    }
}

// Test SuitsCreated: initialized "global" variable set to true
func TestSuitsCreated(t *testing.T) {
    SuitsCreated()
    if true != initialized {
        t.Errorf("For SuitsCreated, expected initialized to be true, but got %b instead.", initialized)
    }
}

// Test CreateSuits with Suits already initialized
func TestSuitsInitializedCreateSuitsFail(t *testing.T) {
    SuitsCreated()
    _, err := CreateSuits(GlobalSuitsOptions{})
    if err == nil {
        t.Errorf("Error should have occurred: %s", eSUITS_ALREADY_INITIALIZED.Error())
    }
}

func doSuitsCreateSuits(t *testing.T, testName string, collection *suitsCollection, resultCollection *suitsCollection) {
    if !reflect.DeepEqual(collection.Suits, resultCollection.Suits) {
        //for idx, suit := range collection.Suits {
        //    fmt.Printf("\n%v: %v ? %v\n", suit.GetName(), suit.GetColor(), resultCollection.Suits[idx].GetColor())
        //}
        t.Errorf("For CreateSuits (%s: Suits), expected\n%+v\nbut got\n%+v\ninstead.", testName, resultCollection.Suits, suitsCol.Suits)
    }

    if !reflect.DeepEqual(collection.colors, resultCollection.colors) {
        t.Errorf("For CreateSuits (%s: Colors), expected\n%+v\nbut got\n%+v\ninstead.", testName, resultCollection.colors, suitsCol.colors)
    }

    if !reflect.DeepEqual(collection.options, resultCollection.options) {
        t.Errorf("For CreateSuits (%s: Options), expected\n%+v\nbut got\n%+v\ninstead.", testName, resultCollection.options, suitsCol.options)
    }

    if !reflect.DeepEqual(collection.rankings, resultCollection.rankings) {
        t.Errorf("For CreateSuits (%s: Rankings), expected\n%+v\nbut got\n%+v\ninstead.", testName, resultCollection.rankings, suitsCol.rankings)
    }
}

// Test CreateSuits: Create Default Suits
func TestSuitsCreateSuitsDefaultSuits(t *testing.T) {
    var (
        testSuitOptions = []SuitOptions{
            {
                Name: CLUBS,
                Symbol: cDefaultSuitsSymbols[CLUBS],
            },
            {
                Name: DIAMONDS,
                Symbol: cDefaultSuitsSymbols[DIAMONDS],
            },
            {
                Name: HEARTS,
                Symbol: cDefaultSuitsSymbols[HEARTS],
            },
            {
                Name: SPADES,
                Symbol: cDefaultSuitsSymbols[SPADES],
            },
        }
        testOpts = GlobalSuitsOptions{}
        testModifiedOpts = GlobalSuitsOptions{
            CreateDefaultSuits: true,
        }

        resultCollection = &suitsCollection{
            initialized: true,
            options: testModifiedOpts,
            Suits: []Suit{
                {
                    ansi: ANSI_BLACK + cDefaultSuitsSymbols[CLUBS] + ANSI_RESET,
                    color: BLACK,
                    glyph: cGLYPH[CLUBS],
                    htmlDecimal: cDECIMAL[CLUBS],
                    htmlHex: cHEX[CLUBS],
                    html: cHTML[CLUBS],
                    isTrumpSuit: false,
                    name: CLUBS,
                    options: testSuitOptions[0],
                    symbol: cDefaultSuitsSymbols[CLUBS],
                    rank: 1,
                },
                {
                    ansi: ANSI_RED + cDefaultSuitsSymbols[DIAMONDS] + ANSI_RESET,
                    color: RED,
                    glyph: cGLYPH[DIAMONDS],
                    htmlDecimal: cDECIMAL[DIAMONDS],
                    htmlHex: cHEX[DIAMONDS],
                    html: cHTML[DIAMONDS],
                    isTrumpSuit: false,
                    name: DIAMONDS,
                    options: testSuitOptions[1],
                    symbol: cDefaultSuitsSymbols[DIAMONDS],
                    rank: 1,
                },
                {
                    ansi: ANSI_RED + cDefaultSuitsSymbols[HEARTS] + ANSI_RESET,
                    color: RED,
                    glyph: cGLYPH[HEARTS],
                    htmlDecimal: cDECIMAL[HEARTS],
                    htmlHex: cHEX[HEARTS],
                    html: cHTML[HEARTS],
                    isTrumpSuit: false,
                    name: HEARTS,
                    options: testSuitOptions[2],
                    symbol: cDefaultSuitsSymbols[HEARTS],
                    rank: 1,
                },
                {
                    ansi: ANSI_BLACK + cDefaultSuitsSymbols[SPADES] + ANSI_RESET,
                    color: BLACK,
                    glyph: cGLYPH[SPADES],
                    htmlDecimal: cDECIMAL[SPADES],
                    htmlHex: cHEX[SPADES],
                    html: cHTML[SPADES],
                    isTrumpSuit: false,
                    name: SPADES,
                    options: testSuitOptions[3],
                    symbol: cDefaultSuitsSymbols[SPADES],
                    rank: 1,
                },
            },
        }

    )

    DestroySuits()
    collection, _ := CreateSuits(testOpts)
    doSuitsCreateSuits(t, "Create Default Suits", collection, resultCollection)
}

// Test CreateSuits: Create Default Suits
func TestSuitsCreateSuitsDefaultSuitsWithTrumpSuit(t *testing.T) {
    var (
        testResultSuitOptions = []SuitOptions{
            {
                Name: CLUBS,
                Symbol: cDefaultSuitsSymbols[CLUBS],
            },
            {
                Name: DIAMONDS,
                Symbol: cDefaultSuitsSymbols[DIAMONDS],
            },
            {
                Name: HEARTS,
                Symbol: cDefaultSuitsSymbols[HEARTS],
                IsTrumpSuit: true,
            },
            {
                Name: SPADES,
                Symbol: cDefaultSuitsSymbols[SPADES],
            },
        }
        testOpts = GlobalSuitsOptions{
            TrumpSuit: HEARTS,
        }
        testModifiedOpts = GlobalSuitsOptions{
            CreateDefaultSuits: true,
            TrumpSuit: HEARTS,
        }

        resultCollection = &suitsCollection{
            initialized: true,
            options: testModifiedOpts,
            Suits: []Suit{
                {
                    ansi: ANSI_BLACK + cDefaultSuitsSymbols[CLUBS] + ANSI_RESET,
                    color: BLACK,
                    glyph: cGLYPH[CLUBS],
                    htmlDecimal: cDECIMAL[CLUBS],
                    htmlHex: cHEX[CLUBS],
                    html: cHTML[CLUBS],
                    isTrumpSuit: false,
                    name: CLUBS,
                    options: testResultSuitOptions[0],
                    symbol: cDefaultSuitsSymbols[CLUBS],
                    rank: 1,
                },
                {
                    ansi: ANSI_RED + cDefaultSuitsSymbols[DIAMONDS] + ANSI_RESET,
                    color: RED,
                    glyph: cGLYPH[DIAMONDS],
                    htmlDecimal: cDECIMAL[DIAMONDS],
                    htmlHex: cHEX[DIAMONDS],
                    html: cHTML[DIAMONDS],
                    isTrumpSuit: false,
                    name: DIAMONDS,
                    options: testResultSuitOptions[1],
                    symbol: cDefaultSuitsSymbols[DIAMONDS],
                    rank: 1,
                },
                {
                    ansi: ANSI_RED + cDefaultSuitsSymbols[HEARTS] + ANSI_RESET,
                    color: RED,
                    glyph: cGLYPH[HEARTS],
                    htmlDecimal: cDECIMAL[HEARTS],
                    htmlHex: cHEX[HEARTS],
                    html: cHTML[HEARTS],
                    isTrumpSuit: true,
                    name: HEARTS,
                    options: testResultSuitOptions[2],
                    symbol: cDefaultSuitsSymbols[HEARTS],
                    rank: 1,
                },
                {
                    ansi: ANSI_BLACK + cDefaultSuitsSymbols[SPADES] + ANSI_RESET,
                    color: BLACK,
                    glyph: cGLYPH[SPADES],
                    htmlDecimal: cDECIMAL[SPADES],
                    htmlHex: cHEX[SPADES],
                    html: cHTML[SPADES],
                    isTrumpSuit: false,
                    name: SPADES,
                    options: testResultSuitOptions[3],
                    symbol: cDefaultSuitsSymbols[SPADES],
                    rank: 1,
                },
            },
        }

    )

    DestroySuits()
    collection, _ := CreateSuits(testOpts)
    doSuitsCreateSuits(t, "Create Default Suits", collection, resultCollection)
}

// Test CreateSuits: Create Default Suits with Custom Colors within []Suits
func TestSuitsCreateSuitsDefaultSuitsWithCustomColorsInSuits(t *testing.T) {
    var (
        testSuitOptions = []SuitOptions{
            {
                Color: BLACK,
                Name: CLUBS,
            },
            {
                Color: YELLOW,
                Name: DIAMONDS,
            },
            {
                Color: RED,
                Name: HEARTS,
            },
            {
                Color: GREEN,
                Name: SPADES,
            },
        }

        testOpts = GlobalSuitsOptions{
            Suits: testSuitOptions,
        }

        testModifiedOpts = GlobalSuitsOptions{
            //CreateDefaultSuits: true,
            Suits: testSuitOptions,
        }

        resultCollection = &suitsCollection{
            initialized: true,
            options: testModifiedOpts,
            Suits: []Suit{
                {
                    ansi: ANSI_BLACK + cDefaultSuitsSymbols[CLUBS] + ANSI_RESET,
                    color: BLACK,
                    glyph: cGLYPH[CLUBS],
                    htmlDecimal: cDECIMAL[CLUBS],
                    htmlHex: cHEX[CLUBS],
                    html: cHTML[CLUBS],
                    isTrumpSuit: false,
                    name: CLUBS,
                    options: testSuitOptions[0],
                    symbol: cDefaultSuitsSymbols[CLUBS],
                    rank: 1,
                },
                {
                    ansi: ANSI_YELLOW + cDefaultSuitsSymbols[DIAMONDS] + ANSI_RESET,
                    color: YELLOW,
                    glyph: cGLYPH[DIAMONDS],
                    htmlDecimal: cDECIMAL[DIAMONDS],
                    htmlHex: cHEX[DIAMONDS],
                    html: cHTML[DIAMONDS],
                    isTrumpSuit: false,
                    name: DIAMONDS,
                    options: testSuitOptions[1],
                    symbol: cDefaultSuitsSymbols[DIAMONDS],
                    rank: 1,
                },
                {
                    ansi: ANSI_RED + cDefaultSuitsSymbols[HEARTS] + ANSI_RESET,
                    color: RED,
                    glyph: cGLYPH[HEARTS],
                    htmlDecimal: cDECIMAL[HEARTS],
                    htmlHex: cHEX[HEARTS],
                    html: cHTML[HEARTS],
                    isTrumpSuit: false,
                    name: HEARTS,
                    options: testSuitOptions[2],
                    symbol: cDefaultSuitsSymbols[HEARTS],
                    rank: 1,
                },
                {
                    ansi: ANSI_GREEN + cDefaultSuitsSymbols[SPADES] + ANSI_RESET,
                    color: GREEN,
                    glyph: cGLYPH[SPADES],
                    htmlDecimal: cDECIMAL[SPADES],
                    htmlHex: cHEX[SPADES],
                    html: cHTML[SPADES],
                    isTrumpSuit: false,
                    name: SPADES,
                    options: testSuitOptions[3],
                    symbol: cDefaultSuitsSymbols[SPADES],
                    rank: 1,
                },
            },
        }

    )

    DestroySuits()
    collection, _ := CreateSuits(testOpts)
    doSuitsCreateSuits(t, "Create Default Suits with Custom Colors within []Suits as Custom Suits", collection, resultCollection)
}

// Test CreateSuits: Create Default Suits with Custom Colors Slice
func TestSuitsCreateSuitsDefaultSuitsWithCustomColors(t *testing.T) {
    var (
        testSuitOptions = []SuitOptions{
            {
                Name: CLUBS,
            },
            {
                Name: DIAMONDS,
            },
            {
                Name: HEARTS,
            },
            {
                Name: SPADES,
                IsTrumpSuit: true,
            },
        }
        testColors = map[string]string{
            CLUBS: BLACK,
            HEARTS: RED,
            DIAMONDS: YELLOW,
            SPADES: GREEN,
        }
        testOpts = GlobalSuitsOptions{
            Colors: testColors,
            CreateDefaultSuits: false,
            Suits: testSuitOptions,
            TrumpSuit: "",
        }

        // CreateSuits modifies SuitOptions to add Color if GlobalSuitsOptions.Colors exists
        testModifiedSuitOptions = []SuitOptions{
            {
                Color: BLACK,
                Name: CLUBS,
            },
            {
                Color: YELLOW,
                Name: DIAMONDS,
            },
            {
                Color: RED,
                Name: HEARTS,
            },
            {
                Color: GREEN,
                Name: SPADES,
                IsTrumpSuit: true,
            },
        }
        resultCollection = &suitsCollection{
            colors: testColors,
            initialized: true,
            options: testOpts,
            Suits: []Suit{
                {
                    ansi: ANSI_BLACK + cDefaultSuitsSymbols[CLUBS] + ANSI_RESET,
                    color: BLACK,
                    glyph: cGLYPH[CLUBS],
                    htmlDecimal: cDECIMAL[CLUBS],
                    htmlHex: cHEX[CLUBS],
                    html: cHTML[CLUBS],
                    isTrumpSuit: false,
                    name: CLUBS,
                    options: testModifiedSuitOptions[0],
                    symbol: cDefaultSuitsSymbols[CLUBS],
                    rank: 1,
                },
                {
                    ansi: ANSI_YELLOW + cDefaultSuitsSymbols[DIAMONDS] + ANSI_RESET,
                    color: YELLOW,
                    glyph: cGLYPH[DIAMONDS],
                    htmlDecimal: cDECIMAL[DIAMONDS],
                    htmlHex: cHEX[DIAMONDS],
                    html: cHTML[DIAMONDS],
                    isTrumpSuit: false,
                    name: DIAMONDS,
                    options: testModifiedSuitOptions[1],
                    symbol: cDefaultSuitsSymbols[DIAMONDS],
                    rank: 1,
                },
                {
                    ansi: ANSI_RED + cDefaultSuitsSymbols[HEARTS] + ANSI_RESET,
                    color: RED,
                    glyph: cGLYPH[HEARTS],
                    htmlDecimal: cDECIMAL[HEARTS],
                    htmlHex: cHEX[HEARTS],
                    html: cHTML[HEARTS],
                    isTrumpSuit: false,
                    name: HEARTS,
                    options: testModifiedSuitOptions[2],
                    symbol: cDefaultSuitsSymbols[HEARTS],
                    rank: 1,
                },
                {
                    ansi: ANSI_GREEN + cDefaultSuitsSymbols[SPADES] + ANSI_RESET,
                    color: GREEN,
                    glyph: cGLYPH[SPADES],
                    htmlDecimal: cDECIMAL[SPADES],
                    htmlHex: cHEX[SPADES],
                    html: cHTML[SPADES],
                    isTrumpSuit: true,
                    name: SPADES,
                    options: testModifiedSuitOptions[3],
                    symbol: cDefaultSuitsSymbols[SPADES],
                    rank: 1,
                },
            },
        }

    )

    DestroySuits()
    collection, _ := CreateSuits(testOpts)
    doSuitsCreateSuits(t, "Create Default Suits with Custom Colors Slice", collection, resultCollection)
}
