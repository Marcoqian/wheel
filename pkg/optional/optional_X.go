package optional

// Map 类型转换
// 注意：mapper 返回值不可为 nil
func Map[T, U any](o Optional[T], mapper func(v T) U) Optional[U] {
	if o.IsNone() {
		return None[U]()
	}
	return Of[U](mapper(o[value]))
}

func MapE[T, U any](o Optional[T], mapper func(v T) (U, error)) (Optional[U], error) {
	if o.IsNone() {
		return None[U](), nil
	}
	u, err := mapper(o[value])
	if err != nil {
		return None[U](), err
	}
	return Of[U](u), nil
}

func MapOrElse[T, U any](o Optional[T], defaultValue U, mapper func(v T) U) U {
	if o.IsNone() {
		return defaultValue
	}
	return mapper(o[value])
}

func MapOrElseE[T, U any](o Optional[T], defaultValue U, mapper func(v T) (U, error)) (U, error) {
	if o.IsNone() {
		return defaultValue, nil
	}
	return mapper(o[value])
}
