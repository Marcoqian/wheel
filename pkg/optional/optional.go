package optional

import (
	"fmt"
)

// Optional 灵感来自于 Java 的空指针处理
// Go 中没有很好的空指针处理工具，因此实现一个简易版的
type Optional[T any] []T

var ErrValueIsNone = fmt.Errorf("value is none")

const value = 0

// OfNullable 可处理 nil 值
func OfNullable[T any](v *T) Optional[*T] {
	if v == nil {
		return None[*T]()
	}
	return Of[*T](v)
}

// Of 仅可用于非 nil 值
func Of[T any](v T) Optional[T] {
	return Optional[T]{value: v}
}

// None 空 optional
func None[T any]() Optional[T] {
	return nil
}

func (o Optional[T]) IsNone() bool {
	return o == nil
}

func (o Optional[T]) IsPresent() bool {
	return o != nil
}

func (o Optional[T]) Get() (T, error) {
	if o.IsNone() {
		var zero T
		return zero, ErrValueIsNone
	}
	return o[value], nil
}

func (o Optional[T]) OrElse(defaultValue T) T {
	if o.IsNone() {
		return defaultValue
	}
	return o[value]
}

func (o Optional[T]) OrElseGet(supplier func() T) T {
	if o.IsNone() {
		return supplier()
	}
	return o[value]
}

func (o Optional[T]) OrElseGetE(supplier func() (T, error)) (T, error) {
	if o.IsNone() {
		return supplier()
	}
	return o[value], nil
}

func (o Optional[T]) OrErr(e error) error {
	if o.IsNone() {
		return e
	}
	return nil
}

func (o Optional[T]) Filter(predicate func(v T) bool) Optional[T] {
	if o.IsNone() || !predicate(o[value]) {
		return None[T]()
	}
	return o
}

func (o Optional[T]) IfPresent(consumer func(v T)) {
	if o.IsPresent() {
		consumer(o[value])
	}
}

func (o Optional[T]) IfPresentE(consumer func(v T) error) error {
	if o.IsPresent() {
		return consumer(o[value])
	}
	return nil
}

func (o Optional[T]) IfNone(consumer func()) {
	if o.IsNone() {
		consumer()
	}
}

func (o Optional[T]) IfNoneE(consumer func() error) error {
	if o.IsNone() {
		return consumer()
	}
	return nil
}

func (o Optional[T]) Determine(consumerOfPresent func(v T), consumerOfNone func()) {
	if o.IsPresent() {
		consumerOfPresent(o[value])
		return
	}
	consumerOfNone()
}

func (o Optional[T]) DetermineE(consumerOfPresent func(v T) error, consumerOfNone func() error) error {
	if o.IsPresent() {
		return consumerOfPresent(o[value])
	}
	return consumerOfNone()
}

// Map ... Go 的泛型不算强大，结构体泛型无法推导出其他类型，但是保留此方法仍然是有意义的
// 更泛用的 Map 参考 optional_X.go 中的 Map 方法（虽然通用，但是不能链式调用，用起来也不是很优雅）
// 注意此处 apply 返回值不能为 nil
func (o Optional[T]) Map(apply func(v T) T) Optional[T] {
	if o.IsNone() {
		return None[T]()
	}
	return Of(apply(o[value]))
}

func (o Optional[T]) MapE(apply func(v T) (T, error)) (Optional[T], error) {
	if o.IsNone() {
		return None[T](), nil
	}
	v, err := apply(o[value])
	if err != nil {
		return None[T](), err
	}
	return Of(v), nil
}
