package load

import (
	"github.com/hootuu/utils/errors"
	"strings"
)

type ByCode = string

type By struct {
	handlers map[ByCode]func() *errors.Error
}

func NewBy() *By {
	return &By{handlers: make(map[ByCode]func() *errors.Error)}
}

func (by *By) Verify(code ByCode) *errors.Error {
	_, ok := by.handlers[code]
	if !ok {
		return errors.Verify("can not use by: " + code)
	}
	return nil
}

func (by *By) Register(code ByCode, handle func() *errors.Error) *By {
	by.handlers[code] = handle
	return by
}

func (by *By) Do(byCode ByCode) *errors.Error {
	call, ok := by.handlers[byCode]
	if !ok {
		return errors.Verify("[by] must be: [" + strings.Join(by.AllByCode(), "|") + "], but: [" + byCode + "]")
	}
	return call()
}

func (by *By) AllByCode() []string {
	all := make([]string, len(by.handlers))
	i := 0
	for k, _ := range by.handlers {
		all[i] = k
		i += 1
	}
	return all
}
