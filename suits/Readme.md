# Suits

## Creating Suits
To create suits, call `suits.CreateSuits`. It accepts an options argument with these properties:
* *Rankings* (`[]string`): Array of strings ordering the Suits from _Lowest_ to _Highest_ (e.g., `cDefaultRankings []string = []string{CLUBS, DIAMONDS, HEARTS, SPADES}`)
* *Colors* (`map[string]string`): A key-value (suit:color) map of suits to colors.<br/> `GetSuitsDefaultColorScheme()` returns:
    Default:
    ```golang
     map[string]string = map[string]string{
       CLUBS: BLACK,
       DIAMONDS: RED,
       HEARTS: RED,
       SPADES: BLACK,
    }
    ```
    Custom example:
    ```golang
    MyColorScheme map[string]string = map[string]string{
        CLUBS: GREEN,
        DIAMONDS: BLUE,
        HEARTS: BLUE,
        SPADES: GREEN,
    }
    ```
* *CreateDefaultSuits* (`bool`): A boolean value representing whether the default suits (Clubs, Diamonds, Hards, and Spades) should be created. (Default: true)
* *Suits* (`[]SuitOptions`): An array of custom suits (i.e., suit options). This is what you would use to create your own custom suit(s).
* *TrumpSuit* (`string`): The suit that is the trump suit.

### Creating Default Suits
Using the `CreateDefaultSuits` property will automagically create the default suits.
```golang
suits.CreateSuits(suits.GlobalSuitsOptions{
    CreateDefaultSuits: true,
})
```

To prevent the creation of default suits:
```golang
suits.CreateSuits(suits.GlobalSuitsOptions{
    CreateDefaultSuits: false,
})
```

### Creating a Custom Suit
You can create a custom suit in one of two ways:

1. Creating custom suit via `suits.CreateSuits` method invocation. `suits.CreateSuits` only returns an error if suits have already been created.

1. Creating a custom suit using `NewSuit`. All suits are created ultimately with this method. `NewSuit` returns a `Suit` object.

`suits.CreateSuits` accepts `GlobalSuitsOptions` which has a property to create custom Suits accepting `SuitOptions` (e.g., `suits.CreateSuits(GlobalSuitsOptions{Suits: []Suit{SuitOptions{}}})`). `suits.CreateSuits` loops through the `Suits` property and uses `suits.NewSuit()` to create the custom suit. `suits.NewSuit()` accepts the `SuitOptions` object.

The only requirement in over-writing a default suit by creating a default suit yourself is the Name. If you are creating a custom suit, then all elements are required (except `Rank`).

#### SuitOptions Object
* ANSI (`string`): ANSI compatible display of the suit. These ANSI constants are available for use: ANSI_RESET, ANSI_BLACK, ANSI_RED, ANSI_GREEN, ANSI_YELLOW, ANSI_BLUE.
* Color (`string`): Color of the suit. The color constants are available for use: BLACK, RED, GREEN, BLUE, YELLOW.
* Glyph (`string`): Unicode symbol for displaying the suit symbol.
* HTMLDecimal (`string`): HTML decimal entity.
* HTMLHex (`string`): HTML hex entity.
* HTML (`string`): HTML popular entity (named entity, e.g. `&clubsuit`)
* IsTrumpSuit bool`): Whether the suit is a trump suit.
* Name (`string`): Name of the suit. Make the first letter uppercase for consistency. See Suit Name constants below.
* Rank (`int`):
* Symbol (`string`): Symbol to be used.

### Displays
Currently, unicode glyphs, HTML entities (HEX, DECIMAL, & SIMPLE) are available via constants.

### Color Scheme
The traditional coloring of suits are:
* HEARTS: <span style="color:red">RED</span>
* DIAMONDS: <span style="color:red">RED</span>
* SPADES: BLACK
* CLUBS: BLACK

You can set a *Custom* Color Scheme where you define the color scheme however you'd like. To use your custom color scheme:
```golang
suits.CreateSuits(suits.GlobalSuitsOptions{
    Colors: map[string]string{
        suits.CLUBS: suits.BLUE,
        suits.DIAMONDS: suits.GREEN,
        suits.HEARTS: suits.GREEN,
        suits.SPADES: suits.BLUE,
    },
})
```

Or, you can create the Suits yourself:
```golang
suits.CreateSuits(suits.GlobalSuitsOptions{
        Suits: []suits.SuitOptions{
            {
                Color: suits.BLACK,
                Name: suits.CLUBS,
            },
            {
                Color: suits.RED,
                Name: suits.HEARTS,
            },
            {
                Color: suits.YELLOW,
                Name: suits.DIAMONDS,
            },
            {
                Color: suits.GREEN,
                Name: suits.SPADES,
                IsTrumpSuit: true,
            },
        },
    })

```

## Pre-defined Settings

### Constants
These constants are available for your use:
* Default Suit Names
  * CLUBS = "Clubs"
  * DIAMONDS = "Diamonds"
  * HEARTS = "Hearts"
  * SPADES = "Spades"
* Colors
  * BLACK = "black"
  * RED = "red"
  * GREEN = "green"
  * BLUE = "blue"
  * YELLOW = "yellow"
  * WHITE = "white"
  * MAGENTA = "magenta"
  * PURPLE = MAGENTA
  * CYAN = "cyan"
  * BLUEGREEN = CYAN
  * ANSI_RESET = "\u001b[0m"
  * ANSI_BLACK = "\u001b[30m"
  * ANSI_RED = "\u001b[31m"
  * ANSI_GREEN = "\u001b[32m"
  * ANSI_YELLOW = "\u001b[33m"
  * ANSI_BLUE = "\u001b[34m"
  * ANSI_WHITE = "\u001b[37m"
  * ANSI_MAGENTA = "\u001b[35m"
  * ANSI_PURPLE = ANSI_MAGENTA
  * ANSI_CYAN = "\u001b[36m"
  * ANSI_BLUEGREEN = ANSI_CYAN
* Defaults
  * DEFAULT_TRUMPSUIT = false

#### Pseudo-Constants Available via Get* Functions
These constant maps are available for your use:
* Default Color Scheme: `GetSuitsDefaultColorScheme()`
* Default Suit Symbols: `GetSuitsDefaultSymbols()`
* Default Glyphs: `GetSuitsGlyphs()`
* Default HTML Decimal entity: `GetSuitsHTMLDecimal()`
* Default HTML Hex entity: `GetSuitsHTMLHex()`
* Default HTML entity (popular): `GetSuitsHTML()`

You can also access these maps specifically for each default suit:
* Default Color Scheme: `GetSuitDefaultColor()`
* Default Suit Symbols: `GetSuitDefaultSymbol()`
* Default Glyphs: `GetSuitGlyph()`
* Default HTML Decimal entity: `GetSuitHTMLDecimal()`
* Default HTML Hex entity: `GetSuitHTMLHex()`
* Default HTML entity (popular): `GetSuitHTML()`


## Errors
Currently, this library does not allow you to make any changes to Suits once they have been initialized/created.