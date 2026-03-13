package zap

// Field представляет ключ-значение для логирования.
type Field struct {
	Key     string
	Type    fieldType
	Integer int64
	String  string
	// ... другие поля можно добавить при необходимости
}

type fieldType int

// Конструкторы полей — возвращают заглушки, достаточные для проверки типов.

func String(key, val string) Field {
	return Field{Key: key, Type: stringType, String: val}
}

func Int(key string, val int) Field {
	return Field{Key: key, Type: intType, Integer: int64(val)}
}

func Int64(key string, val int64) Field {
	return Field{Key: key, Type: intType, Integer: val}
}

func Bool(key string, val bool) Field {
	return Field{Key: key, Type: boolType, Integer: map[bool]int64{false: 0, true: 1}[val]}
}

func Float64(key string, val float64) Field {
	// Для простоты храним как bits, но в стабе это не критично
	_ = val
	return Field{Key: key, Type: floatType}
}

func Error(err error) Field {
	if err == nil {
		return Skip()
	}
	return String("error", err.Error())
}

func Any(key string, value interface{}) Field {
	// В стабе можно вернуть заглушку — тип проверяется статически
	_ = value
	return Field{Key: key}
}

func Skip() Field {
	return Field{Key: "", Type: skipType}
}

// Типы полей (для внутренней логики, в стабе не используются, но полезны для читаемости)
const (
	unknownType fieldType = iota
	stringType
	intType
	boolType
	floatType
	skipType
)
