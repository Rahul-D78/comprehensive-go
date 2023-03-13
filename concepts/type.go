package concepts

// Understanding Type definations
// A type defination creates a new distinct type with the same underlying type and operations as the given
// type and binds an identifier, the type name to it.

// A type determines a set of values together with operations and methods specific to those values.
// A type may be denoted by a type name
// More about types - https://go.dev/ref/spec#Types
// More about StructTag - https://pkg.go.dev/reflect#StructTag

// Type interface - https://pkg.go.dev/reflect#Type
// type is the representation of a Go type.

// A defined type may have methods associated with it. It does not inherit any methods bound to the given type.

// A Mutex is a data type with two methods, Lock and Unlock.
// type Mutex struct         { /* Mutex fields */ }
// func (m *Mutex) Lock()    { /* Lock implementation */ }
// func (m *Mutex) Unlock()  { /* Unlock implementation */ }

// NewMutex has the same composition as Mutex but its method set is empty.
// type NewMutex Mutex

// Type definitions may be used to define different boolean, numeric, or string types and associate methods with them:
// type TimeZone int
// const (
// EST TimeZone = -(5 + iota)
// )
// func (tz TimeZone) String() string {
// return fmt.Sprintf("GMT%+dh", tz)
// }
