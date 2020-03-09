package json

import "github.com/Jeffail/gabs"

const (
	OptionReplace = MergeOption(0)
	OptionSkip    = MergeOption(1)
)

type MergeOption int

func Merge(old, new []byte, opt MergeOption) ([]byte, error) {
	rawOld, err := gabs.ParseJSON(old)
	if err != nil {
		return nil, err
	}
	rawNew, err := gabs.ParseJSON(new)
	if err != nil {
		return nil, err
	}
	switch opt {
	case OptionSkip:
		err = rawOld.MergeFn(rawNew, skipFn)
	case OptionReplace:
		err = rawOld.MergeFn(rawNew, replaceFn)
	}
	if err != nil {
		return nil, err
	}
	return []byte(rawOld.String()), nil
}

//replaceFn will apply new changes if conflict happen
func replaceFn(_, new interface{}) interface{} {
	return new
}

//skipFn will discard new changes if conflict happen
func skipFn(old, _ interface{}) interface{} {
	return old
}
