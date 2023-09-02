package types

type Id struct {
	Name          GoName
	AutoIncrement bool
	Type          GoType
	LowerName     GoName
	OriginName    string
}

func (i Id) Empty() bool {
	return i == Id{}
}

type Ids []Id

func NewId(name GoName, typ GoType, autoIncrement bool) Id {
	if typ.IsNumber() {
		typ = Int64
	}
	if typ.IsNullable() {
		return Id{
			Name:          name,
			Type:          typ.ToNotNull(),
			AutoIncrement: autoIncrement,
			LowerName:     name.ToLower(),
			OriginName:    name.originName,
		}
	}
	return Id{
		Name:          name,
		Type:          typ,
		AutoIncrement: autoIncrement,
		LowerName:     name.ToLower(),
		OriginName:    name.originName,
	}
}
