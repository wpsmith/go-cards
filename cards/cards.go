package cards

import (
    "sort"
    "github.com/wpsmith/go-cards/suits"
    "strings"
)

type Cards []Card

// Draw a single card
func (C *Cards) Draw() (*Card) {
    cards := C.remove(1)
    return &cards[0]
}

// Draw a single card
func (C *Cards) DrawFromBottom() (*Card) {
    cards := C.removeFromBottom(1)
    return &cards[0]
}

// Draw multiple cards
func (C *Cards) DrawX(x int) (Cards) {
    return C.remove(x)
}

// Draw a single card from the bottom of the deck
func (C *Cards) DrawXFromBottom(x int) (Cards) {
    cards := C.removeFromBottom(x)
    return cards
}

// Sorts the cards by ranking
func (C *Cards) Sort() {
    sort.Sort(C)
}

// Groups the deck by suit
func (C *Cards) GroupBySuit() {
    theSuits, _ := suits.GetSuits()
    sort.Sort(theSuits)

    var cards Cards
    for _, s := range theSuits {
        cards = append(cards, C.GetCardsBySuit(s.GetName())...)
    }
    C = &cards
}

// Gets the cards by a specific rank (e.g., all 2s)
func (C Cards) GetCardsByRank(rank string) Cards {
    return C.pluckByCard(func(c Card) bool {
        return rank == c.GetRank()
    })
}

// Gets the cards by a specific suit name (e.g., suits.CLUBS)
func (C Cards) GetCardsBySuit(suitName string) Cards {
    return C.pluckByCard(func(c Card) bool {
        return suitName == c.suit.GetName()
    })
}

// Groups the deck by suit
func (C Cards) GroupByRank() {
    C.Sort()
}

/** PRIVATE METHODS **/
// Plucks a slice of cards by a matching function
func (C Cards) pluckByCard(test func(Card) bool) (ret Cards) {
    for _, c := range C {
        if test(c) {
            ret = append(ret, c)
        }
    }
    return
}

// Removes/Shifts card from the top/beginning
func (C *Cards) remove(x int) (Cards) {
    var deckCards, drawnCards Cards

    deckCards = *C
    drawnCards = deckCards[0:x]
    deckCards = deckCards[x:]
    *C = deckCards

    return drawnCards
}

// Show hand with symbols
func (C Cards) ToStringByType(stringType string) string {
    var cardsSymbol []string
    for _, card := range C {
        cardsSymbol = append(cardsSymbol, card.GetSymbolBy(stringType))
    }
    return strings.Join(cardsSymbol, " ")
}

// Appends/Adds/Pushes a card to the bottom/end
func (C *Cards) Append(c *Card) {
    *C = append(*C, *c)
}

// Prepends/Adds/Unshifts a card to the top/beginning
func (C *Cards) Prepend(c *Card) {
    *C = append(Cards{*c}, *C...)
}

// Removes/Pops card from the bottom/end
func (C *Cards) removeFromBottom(x int) (Cards) {
    var deckCards, drawnCards Cards

    deckCards = *C
    length := len(deckCards)
    drawnCards = deckCards[length - x:]
    deckCards = deckCards[0:length - x]
    *C = deckCards

    return drawnCards
}


/** SORT INTERFACE METHODS **/
// Length of Slice
func (s Cards) Len() int {
    return len(s)
}

// Less Comparison
func (s Cards) Less(i, j int) bool {
    return s[i].sortValue < s[j].sortValue;
}

// Swap Operation
func (s Cards) Swap(i, j int) {
    s[i], s[j] = s[j], s[i]
}
