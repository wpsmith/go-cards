package cards

import (
    "testing"
    "github.com/wpsmith/go-cards/suits"
    "reflect"
)


// Test the Error return for NewCard
func TestNewCardNoSuitError(t *testing.T) {
    var err error
    _, err = NewCustomCard(CardOptions{
        Rank: "2",
    })
    if err == nil {
        t.Errorf("For NewCard error testing, expected " + errorCardMissingRequiredProperties("Suit").Error() + ", but got <nil> instead.")
    }

    suitsCol, _ := suits.CreateSuits(suits.GlobalSuitsOptions{})
    _, err = NewCustomCard(CardOptions{
        Suit: *suitsCol.Suits.GetSuit(suits.CLUBS),
    })
    if err == nil {
        t.Errorf("For NewSuit error testing, expected " + errorCardMissingRequiredProperties("Rank").Error() + ", but got <nil> instead.")
    }
}

/** Test NewSuit **/
type testNewCustomCard struct {
    opts CardOptions
    card *Card
}

func TestNewCustomCard(t *testing.T) {
    var (
        //RANKS []string = []string{"2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K", "A", "RJ", "BJ"}
        testNewCustomCardCard = testNewCustomCard{
            CardOptions{
                Rank: "BJ",
                Glyph: "\U0001f0bf",
                HTMLDecimal: "&#127167",
                HTMLHex: "&#x1f0bf",
                HTML: "&#x1f0bf",
            },
            &Card{
                ansi: suits.ANSI_BLACK + "JBJ" + suits.ANSI_RESET,
                cardType: JOKER_CARD,
                color: suits.BLACK,
                glyph: cGLYPH["BJB"],
                htmlDecimal: cDECIMAL["BJB"],
                htmlHex: cHEX["BJB"],
                html: cHEX["BJB"],
                initialized: true,
                rank: "BJ",
                sortValue: 5 + cDEFAULT_CARD_VALUES["BJ"],
                suit: suits.Suit{},
                suitFirst: false,
                symbol: "BJB",
                value: cDEFAULT_CARD_VALUES["BJ"],
            },
        }
    )

    // Clean Slate
    suits.DestroySuits()
    var opts suits.SuitOptions = suits.SuitOptions{
        ANSI: suits.ANSI_BLACK,
        Color: suits.BLACK,
        Glyph: cGLYPH["BJB"],
        HTMLDecimal: cDECIMAL["BJB"],
        HTMLHex: cHEX["BJB"],
        HTML: cHEX["BJB"],
        IsTrumpSuit: true,
        Name: suits.BLACK,
        Rank: 5,
        Symbol: "B",
    }
    blackSuit, _ := suits.NewSuit(opts)

    // Start from scratch
    testNewCustomCardCard.opts.Suit = *blackSuit
    testNewCustomCardCard.card.suit = *blackSuit
    card, err := NewCustomCard(testNewCustomCardCard.opts)

    if err != nil {
        t.Errorf("For NewCustomCard, expected no error, received error: %#v\n", err.Error())
    }

    // Not sure why this fails??
    if !reflect.DeepEqual(card, testNewCustomCardCard.card) {
        t.Errorf("For NewCustomCard, expected\n%#v\nreceived \n%#v\n", testNewCustomCardCard.card, card)
    }
}


type testNewCard struct {
    suit suits.Suit
    rank string
    card *Card
}
func TestNewCard(t *testing.T) {
    var (
        testNewCardCard = testNewCard{
            suits.Suit{},
            "2",
            &Card{
                ansi: suits.ANSI_BLACK + "2♣" + suits.ANSI_RESET,
                cardType: NUMBER_CARD,
                color: suits.BLACK,
                glyph: "\U0001f0d2",
                htmlDecimal: "&#127186",
                htmlHex: "&#x1f0d2",
                html: "&#x1f0d2",
                initialized: true,
                rank: "2",
                sortValue: 3,
                suit: suits.Suit{},
                suitFirst: false,
                symbol: "2♣",
                value: 2,
            },
        }
    )

    // Clean Slate
    suits.DestroySuits()

    // Start from scratch
    s, _ := suits.CreateSuits(suits.GlobalSuitsOptions{})
    clubsSuit := *s.GetSuit(suits.CLUBS)
    testNewCardCard.suit = clubsSuit
    testNewCardCard.card.suit = clubsSuit

    card, err := NewCard(testNewCardCard.suit, testNewCardCard.rank)
    if err != nil {
        t.Errorf("For NewCard, expected no error, received error: %#v\n", err.Error())
    }

    if !reflect.DeepEqual(card, testNewCardCard.card) {
        t.Errorf("For NewCard, expected\n%#v\nreceived\n%#v\n", testNewCardCard.card, card)
    }
}

type testCardRanks struct {
    val  string
    card *Card
    rank int
}

func TestCardRankValue(t *testing.T) {
    var testsCardRanks []testCardRanks

    // Clean Slate
    suits.DestroySuits()

    // Start from scratch
    s, _ := suits.CreateSuits(suits.GlobalSuitsOptions{})
    suit := s.GetSuit(suits.DIAMONDS)

    // Create testing
    for index, rank := range cDEFAULT_RANKS {
        c, _ := NewCard(*suit, rank)
        testsCardRanks = append(testsCardRanks, testCardRanks{
            val: rank,
            card: c,
            rank: index + 1,
        })
    }

    // Do Testing
    for _, testCardRank := range testsCardRanks {
        if testCardRank.card.RankValue() != testCardRank.rank {
            t.Errorf("For %s card, expected %d received %d\n", testCardRank.val, testCardRank.rank, testCardRank.card.RankValue())
        }
    }
}