package coinmarketcap

import (
	"slices"
	"sync"

	"github.com/mr-karan/balance"
)

// api-key weighted-round-robin
type apiKeyWRR struct {
	mutex   *sync.RWMutex
	keys    []string
	balance *balance.Balance
}

func newApiKeyWRR(keys []string) *apiKeyWRR {
	b := balance.NewBalance()
	for _, key := range keys {
		if key == "" {
			continue
		}
		_ = b.Add(key, 1)
	}

	return &apiKeyWRR{
		mutex:   new(sync.RWMutex),
		keys:    keys,
		balance: b,
	}
}

func (inst *apiKeyWRR) getKey() (key string) {
	inst.mutex.RLock()
	defer inst.mutex.RUnlock()

	key = inst.balance.Get()
	return
}

func (inst *apiKeyWRR) SetKeys(keys []string) {
	inst.mutex.RLock()
	current := inst.keys
	inst.mutex.RUnlock()

	if !inst.isChange(current, keys) {
		return
	}

	b := balance.NewBalance()
	for _, key := range keys {
		_ = b.Add(key, 1)
	}

	inst.mutex.Lock()
	defer inst.mutex.Unlock()

	inst.keys = keys
	inst.balance = b

	return
}

// Determine if the elements and weights involved in the round-robin have changed.
func (inst *apiKeyWRR) isChange(current, new []string) bool {
	if len(current) != len(new) {
		return true
	}

	for _, c := range current {
		if !slices.Contains(new, c) {
			return true
		}
	}

	return false
}
