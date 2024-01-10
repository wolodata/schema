package constant

import (
	"fmt"
)

const (
	ShortIdWhiteHouse      = "WhiteHouse"
	ShortIdDOJ             = "DOJ"
	ShortIdUSTR            = "USTR"
	ShortIdStateDepartment = "StateDepartment"
)

var (
	sourceMap map[string]*Source
)

type Source struct {
	ShortId     string
	NameChinese string
	NameEnglish string
	IsChinese   bool
}

func registerSource(s Source) {
	_, ok := sourceMap[s.ShortId]
	if ok {
		panic(fmt.Sprintf("duplicate source: %v", s))
	}

	if sourceMap == nil {
		sourceMap = make(map[string]*Source)
	}
	sourceMap[s.ShortId] = &s
}

func GetSourceByShortId(shortId string) (*Source, error) {
	has, ok := sourceMap[shortId]
	if !ok {
		return nil, fmt.Errorf("source not found, shortId: %s", shortId)
	}

	return has, nil
}

func init() {
	registerSource(Source{
		ShortId:     ShortIdWhiteHouse,
		NameChinese: "白宫",
		NameEnglish: "whitehouse.gov",
		IsChinese:   false,
	})
	registerSource(Source{
		ShortId:     ShortIdDOJ,
		NameChinese: "司法部",
		NameEnglish: "justice.gov",
		IsChinese:   false,
	})
	registerSource(Source{
		ShortId:     ShortIdUSTR,
		NameChinese: "贸易代表署",
		NameEnglish: "ustr.gov",
		IsChinese:   false,
	})
	registerSource(Source{
		ShortId:     ShortIdStateDepartment,
		NameChinese: "国务院",
		NameEnglish: "state.gov",
		IsChinese:   false,
	})
}
