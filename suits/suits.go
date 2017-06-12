package suits

import (
    "strings"
    "sort"
)

var (
    initialized bool = false
    suitsCol *suitsCollection
)

// Global Suits Options Object
type GlobalSuitsOptions struct {
    Colors             map[string]string
    CreateDefaultSuits bool
    Suits              []SuitOptions
    TrumpSuit          string
}

type Suits []Suit

/** SORT INTERFACE METHODS **/
// Length of Slice
func (s Suits) Len() int {
    return len(s)
}

// Less Comparison
func (s Suits) Less(i, j int) bool {
    return s[i].name < s[j].name;
}

// Swap Operation
func (s Suits) Swap(i, j int) {
    s[i], s[j] = s[j], s[i]
}

// Suits Collection Object
type suitsCollection struct {
    colors      map[string]string
    initialized bool
    options     GlobalSuitsOptions
    Suits       Suits
    rankings    []string
}

/** PUBLIC METHODS **/
// Search searches for x in a sorted slice of Suits and
// returns the index  as specified by Search.
// The return value is the index to insert x if x is not
// present (it could be len(a)).
// The slice must be sorted in ascending order.
func (s *suitsCollection) Search(x string) int {
    var slice []string
    for _, suit := range s.Suits {
        slice = append(slice, suit.name)
    }
    return sliceIndex(len(s.Suits), sliceIndexFn(slice, x))
    //suits := s.Suits
    //return sort.Search(len(suits), func(i int) bool {
    //    return suits[i].GetName() >= x
    //})
}

// Gets the color scheme colors
func (s *suitsCollection) GetColors() map[string]string {
    return s.colors
}

/** PRIVATE METHODS **/
// Creates the default suits: Clubs, Diamonds, Hearts, Spades
func (s *suitsCollection) createDefaultSuits() error {
    if initialized {
        return eSUITS_ALREADY_INITIALIZED
    }

    // Default trump suit
    trumpSuit := DEFAULT_TRUMPSUIT
    for name, symbol := range cDefaultSuitsSymbols {
        trumpSuit = DEFAULT_TRUMPSUIT

        // Maybe set trump suit
        if s.options.TrumpSuit != "" && strings.ToLower(s.options.TrumpSuit) == strings.ToLower(name) {
            trumpSuit = true
        }

        // Do a built-in Color Scheme
        if len(s.colors) == 0 {
            suit, err := NewSuit(SuitOptions{
                Color: "",
                Name: name,
                Symbol: symbol,
                IsTrumpSuit: trumpSuit,
            })
            if err == nil {
                s.Suits = append(s.Suits, *suit)
            }

            // Do Custom Color Scheme
        } else {
            suit, err := NewSuit(SuitOptions{
                Color: s.colors[name],
                Name: name,
                Symbol: symbol,
                IsTrumpSuit: trumpSuit,
            })
            if err == nil {
                s.Suits = append(s.Suits, *suit)
            } else {
                return ErrorSuitsNotCreated()
            }
        }
    }
    return nil
}

// Creates a Suits Collection
func newSuitsCollection(opts GlobalSuitsOptions) *suitsCollection {
    s := new(suitsCollection)
    s.options = opts

    if len(opts.Colors) > 0 {
        s.colors = opts.Colors
    }

    if opts.CreateDefaultSuits {
        err := s.createDefaultSuits()
        if err != nil {
            panic("Default Suits Error! " + err.Error())
        }
    }

    return s
}

// Initialize Suits
func CreateSuits(opts GlobalSuitsOptions) (*suitsCollection, error) {
    if initialized {
        return nil, eSUITS_ALREADY_INITIALIZED
    }

    // Ensure Default Suits get created if no custom suits
    if len(opts.Suits) == 0 && !opts.CreateDefaultSuits {
        opts.CreateDefaultSuits = true
    }
    suitsCol = newSuitsCollection(opts)

    // Custom Suits
    if len(opts.Suits) > 0 {
        for _, suitOpts := range opts.Suits {
            if suitOpts.Color == "" && len(suitsCol.colors) > 0 {
                suitOpts.Color = suitsCol.colors[suitOpts.Name]
            }
            s, err := NewSuit(suitOpts)
            if err == nil {
                suitsCol.Suits = append(suitsCol.Suits, *s)
            }
        }
    }

    // Sort Suits
    sort.Sort(suitsCol.Suits)

    SuitsCreated()
    return suitsCol, nil
}

// Sets initialized "global" variable to close suits
// from being edited or further modified
func SuitsCreated() {
    suitsCol.initialized, initialized = true, true
}

// Sets initialized "global" variable to false allowing
// suits to be edited or further modified
func DestroySuits() {
    initialized = false
    suitsCol = new(suitsCollection)
}

// Returns the Suits Collection
func GetSuitsCollection() (*suitsCollection, error) {
    if len(suitsCol.Suits) == 0 {
        return nil, ErrorSuitsNotCreated()
    }
    return suitsCol, nil
}

// Returns the Suits from the Suits Collection
func GetSuits() ([]Suit, error) {
    if len(suitsCol.Suits) == 0 {
        return nil, ErrorSuitsNotCreated()
    }
    return suitsCol.Suits, nil
}