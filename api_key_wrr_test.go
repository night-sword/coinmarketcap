package coinmarketcap

import (
	"testing"
)

func TestApiKeyWRR_GetSet(t *testing.T) {
	inst := newApiKeyWRR([]string{})

	k1, k2, k3, k4 := "aaa", "bbb", "ccc", "ddd"

	t.Run("not init, got empty key", func(t *testing.T) {
		k := inst.getKey()

		if k != "" {
			t.Errorf("key got %s, want %s", k, "")
			return
		}
	})

	t.Run("robin count test", func(t *testing.T) {
		keys := []string{k1, k2, k3}

		inst.SetKeys(keys)
		ret := testApiKeyWRR_runTimes(t, keys, inst, 10000)

		count1, count2, count3 := 3334, 3333, 3333
		if ret[k1] != count1 || ret[k2] != count2 || ret[k3] != count3 {
			t.Errorf("ret count want a=%d & b=%d & c=%d but got not a=%d & b=%d & c=%d", count1, count2, count3, ret[k1], ret[k2], ret[k3])
			return
		}
	})

	t.Run("new robin count test", func(t *testing.T) {
		olds := []string{k1, k2, k3}
		news := []string{k1, k2, k4}

		inst.SetKeys(olds)
		ret := testApiKeyWRR_runTimes(t, olds, inst, 10000)

		inst.SetKeys(news)
		ret = testApiKeyWRR_runTimes(t, news, inst, 10000)

		count1, count2, count3 := 3334, 3333, 3333
		if ret[k1] != count1 || ret[k2] != count2 || ret[k4] != count3 {
			t.Errorf("ret count want a=%d & b=%d & c=%d but got not a=%d & b=%d & c=%d", count1, count2, count3, ret[k1], ret[k2], ret[k4])
			return
		}
	})
}

func testApiKeyWRR_runTimes(t *testing.T, keys []string, wrr *apiKeyWRR, times int) (ret map[string]int) {
	ret = make(map[string]int, len(keys))
	for _, k := range keys {
		ret[k] = 0
	}

	for i := 0; i < times; i++ {
		p := wrr.getKey()
		ret[p] = ret[p] + 1
	}

	return
}

func TestProviderWRR_isChange(t *testing.T) {
	k1, k2, k3, k4 := "aaa", "bbb", "ccc", "ddd"
	keys := []string{k1, k2, k3, k4}

	tests := []struct {
		name    string
		current []string
		new     []string
		want    bool
	}{
		{"case.1", []string{}, []string{}, false},
		{"case.2", []string{k1}, []string{k1}, false},
		{"case.3", []string{k1, k2}, []string{k2, k1}, false},
		{"case.4", []string{k1, k2}, []string{k1}, true},
		{"case.5", []string{k1}, []string{k1, k2}, true},
		{"case.6", []string{k1}, []string{k2}, true},
	}

	inst := newApiKeyWRR(keys)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := inst.isChange(tt.current, tt.new); got != tt.want {
				t.Errorf("isChange() = %v, want %v", got, tt.want)
			}
		})
	}
}
