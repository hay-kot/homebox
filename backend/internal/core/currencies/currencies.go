// Package currencies provides a shared definition of currencies. This uses a global
// variable to hold the currencies.
package currencies

import (
	"bytes"
	_ "embed"
	"encoding/json"
	"io"
	"slices"
	"strings"
	"sync"
)

//go:embed currencies.json
var defaults []byte

type CollectorFunc func() ([]Currency, error)

func CollectJSON(reader io.Reader) CollectorFunc {
	return func() ([]Currency, error) {
		var currencies []Currency
		err := json.NewDecoder(reader).Decode(&currencies)
		if err != nil {
			return nil, err
		}

		return currencies, nil
	}
}

func CollectDefaults() CollectorFunc {
	return CollectJSON(bytes.NewReader(defaults))
}

func CollectionCurrencies(collectors ...CollectorFunc) ([]Currency, error) {
	out := make([]Currency, 0, len(collectors))
	for i := range collectors {
		c, err := collectors[i]()
		if err != nil {
			return nil, err
		}

		out = append(out, c...)
	}

	return out, nil
}

type Currency struct {
	Name   string `json:"name"`
	Code   string `json:"code"`
	Local  string `json:"local"`
	Symbol string `json:"symbol"`
}

type CurrencyRegistry struct {
	mu       sync.RWMutex
	registry map[string]Currency
}

func NewCurrencyService(currencies []Currency) *CurrencyRegistry {
	registry := make(map[string]Currency, len(currencies))
	for i := range currencies {
		registry[currencies[i].Code] = currencies[i]
	}

	return &CurrencyRegistry{
		registry: registry,
	}
}

func (cs *CurrencyRegistry) Slice() []Currency {
	cs.mu.RLock()
	defer cs.mu.RUnlock()

	keys := make([]string, 0, len(cs.registry))
	for key := range cs.registry {
		keys = append(keys, key)
	}

	slices.Sort(keys)

	out := make([]Currency, 0, len(cs.registry))
	for i := range keys {
		out = append(out, cs.registry[keys[i]])
	}

	return out
}

func (cs *CurrencyRegistry) IsSupported(code string) bool {
	lower := strings.ToLower(code)

	cs.mu.RLock()
	defer cs.mu.RUnlock()
	_, ok := cs.registry[lower]
	return ok
}
