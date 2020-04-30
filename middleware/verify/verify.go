package verify

import "strings"

type Verify struct {
	keys map[string]string
}

func New(conf string) *Verify {
	ret := &Verify{
		keys: make(map[string]string),
	}
	pairs := strings.Split(conf, ",")
	for _, v := range pairs {
		kv := strings.SplitN(v, ":", 2)
		if len(kv) == 2 {
			ret.keys[kv[0]] = kv[1]
		}
	}
	return ret
}