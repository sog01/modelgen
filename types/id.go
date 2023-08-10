package types

type Id struct {
	Name GoName
	Type GoType
}

func NewId(name GoName, typ GoType) Id {
	if typ.IsNullable() {
		return Id{
			Name: name,
			Type: typ.ToNotNull(),
		}
	}
	return Id{
		Name: name,
		Type: typ,
	}
}
