package gocards

import (
    "math/rand"
    "strings"
    "github.com/wpsmith/go-cards/suits"
    "github.com/wpsmith/go-cards/cards"
    "fmt"
)

// Deck
type Deck struct {
    //Ranks []string
    cards cards.Cards
}

// Shows the deck of cards in a terminal
func (d *Deck) Show() string {
    return d.ToStringBy("ansi")
}

// Show hand with symbols
func (d *Deck) ToStringByType(stringType string) string {
    var cardsSymbol []string
    for _, card := range d.cards {
        cardsSymbol = append(cardsSymbol, card.GetSymbolBy(stringType))
    }
    return strings.Join(cardsSymbol, " ")
}

// Show hand with symbols
func (d *Deck) ToString() string {
    var cardsSymbol []string
    for _, card := range d.cards {
        cardsSymbol = append(cardsSymbol, card.ToString())
    }
    return strings.Join(cardsSymbol, " ")
}

// Show hand with symbols
func (d *Deck) ToStringBy(stringType string) string {
    var cardsSymbol []string
    for _, card := range d.cards {
        var cardStr string
        switch strings.ToLower(stringType) {
        case "ansi":
            cardStr = card.ToANSI()
        case "htmldecimal":
            cardStr = card.ToHTMLDecimal()
        case "html":
            cardStr = card.ToHTML()
        case "htmlhex":
            cardStr = card.ToHTMLHex()
        case "glyph":
            cardStr = card.ToGlyph()
        default:
            cardStr = card.ToString()
        }
        cardsSymbol = append(cardsSymbol, cardStr)
    }
    return strings.Join(cardsSymbol, " ")
}

// Deal by Hand
func (d *Deck) DealByHand(numOfHands int, handSize int) ([]Hand, error) {
    if len(d.cards) == 0 {
        return nil, eDECK_HAS_NO_CARDS
    }

    // Create player hands
    playerHands := make([]Hand, numOfHands)

    for i := 0; i < numOfHands; i++ {
        // Draw from Deck
        handCards := d.cards.DrawX(handSize)

        // Create player hand
        playerHands[i].Deck = Deck{handCards}
    }

    return playerHands, nil
}

// Deal by Player
func (d *Deck) DealByPlayer(players int, handSize int) ([]Hand, error) {
    if len(d.cards) == 0 {
        return nil, eDECK_HAS_NO_CARDS
    }

    // Create player hands
    playerHands := make([]Hand, players)

    // Cycle through deck by number of cards in a hand
    for j := 0; j < handSize; j++ {
        // Cycle through deck by number of players
        for i := 0; i < players; i++ {
            playerHands[i].cards = append(playerHands[i].cards, *d.cards.Draw())
        }
    }
    return playerHands, nil
}

// Gets the cards
func (d *Deck) GetCards() cards.Cards {
    return d.cards
}

// Shuffle uses Knuth shuffle algorithm to randomize the deck in O(n) time
// sourced from https://gist.github.com/quux00/8258425
func (d *Deck) Shuffle() {
    N := len(d.cards)
    for i := 0; i < N; i++ {
        r := i + rand.Intn(N - i)
        d.cards[r], d.cards[i] = d.cards[i], d.cards[r]
    }
}

// Gets the cards
func (d *Deck) Duplicate() {
    d.cards = append(d.cards, d.cards...)
}

// Deck Options
type DeckOptions struct {
    Ranks []string
    Suits []suits.Suit
}

// Creates a new single Deck of cards with default suits and ranks
func NewSingleDeck(opts DeckOptions) (*Deck, error) {
    var err error
    d := new(Deck)

    ranksSlice := opts.Ranks
    if ranksSlice == nil {
        ranksSlice = cards.GetDefaultRanks()
    }

    suitsSlice := opts.Suits
    if suitsSlice == nil {
        suitsSlice, err = suits.GetSuits()
        if err != nil {
            return nil, err
        }
    }

    if err == nil {
        for _, suit := range suitsSlice {
            for _, rank := range ranksSlice {
                card, _ := cards.NewCard(suit, rank)
                d.cards.Append(card)
            }
        }
    } else {
        return nil, eDECK_WAS_NOT_CREATED
    }

    return d, nil
}

// Create a deck of multiple decks
func NewMultiDeck(numOfDecks int, opts DeckOptions) (*Deck, error) {
    var (
        e error
        d *Deck
    )

    d, e = newSingleDeck(opts)
    if e != nil {
        fmt.Println("Got error")
    }

    for i := 1; i <= numOfDecks; i++ {
        d.Duplicate()
    }

    return d, nil
}


// Simple Deck creation
func NewDeck(theCards cards.Cards) (*Deck, error) {
    d := new(Deck)
    d.cards = theCards
    return d, nil
}

// Initialize by creating the suits
func Initialize(opts suits.GlobalSuitsOptions) {
    _, err := suits.CreateSuits(opts)
    if err != nil {
        panic("Initialization Error! " + err.Error())
    }
}

// Hand
type Hand struct {
    Deck
}

// Multiple Decks
type Decks struct {
    Deck
    numberOfDecks int
}