package contract

type UUIDGenerator interface {
	GenerateV4() string
}
