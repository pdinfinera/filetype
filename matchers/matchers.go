package matchers

import "gopkg.in/h2non/filetype.v0/types"

// Internal shortcut to NewType
var newType = types.NewType

// Matcher function interface as type alias
type Matcher func([]byte, int) bool

// Type interface to store pairs of type with its matcher function
type Map map[types.Type]Matcher

// Type specific matcher function interface
type TypeMatcher func([]byte, int) types.Type

// Store registered file type matchers
var Matchers = make(map[types.Type]TypeMatcher)

// Create and register a new type matcher function
func NewMatcher(kind types.Type, fn Matcher) TypeMatcher {
	matcher := func(buf []byte, length int) types.Type {
		if fn(buf, length) {
			return kind
		}
		return types.Unknown
	}

	Matchers[kind] = matcher
	return matcher
}

func register(matchers ...Map) {
	for _, m := range matchers {
		for kind, matcher := range m {
			NewMatcher(kind, matcher)
		}
	}
}

func init() {
	// Arguments order is intentional
	register(Image, Video, Audio, Font, Archive)
}