package roll

// This inlines the cmp package from Go 1.21 while we still support Go 1.19 and 1.20.
// Once we drop support for Go 1.19 and 1.20, we can remove this file and use the cmp package directly.

type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64 |
		~string
}

func isNaN[T Ordered](x T) bool {
	return x != x
}
func Compare[T Ordered](x, y T) int {
	xNaN := isNaN(x)
	yNaN := isNaN(y)
	if xNaN && yNaN {
		return 0
	}
	if xNaN || x < y {
		return -1
	}
	if yNaN || x > y {
		return +1
	}
	return 0
}
