package rule

import "context"

type RuleManager interface {
	Detect(ctx context.Context) bool
}

var managers = make([]RuleManager, 4)

func Register(m RuleManager) {
	managers = append(managers, m)
}

func Detect(ctx context.Context) bool {
	for _, m := range managers {
		if m.Detect(ctx) {
			return true
		}
	}
	return false
}
