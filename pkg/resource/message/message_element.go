package message

type MessageElementAt struct {
	Id   string
	Name string //?	收发	目标用户的名称
	Role string //?	收发	目标角色
	Type string //?	收发	特殊操作，例如 all 表示 @全体成员，here 表示 @在线成员
}

func (e *MessageElementAt) Tag() string {
	return "at"
}

type MessageElementSharp struct {
	Id   string //收发	目标频道的 ID
	Name string //?	收发	目标频道的名称
}

func (e *MessageElementSharp) Tag() string {
	return "sharp"
}

type MessageElementA struct {
	Href string
}

func (e *MessageElementA) Tag() string {
	return "a"
}

type messageResourceElement struct {
	Src     string
	Cache   bool
	Timeout string //ms
}

type MessageElementImg struct {
	*messageResourceElement
	Width  uint32
	Height uint32
}

func (e *MessageElementImg) Tag() string {
	return "img"
}

type MessageElementAudio struct {
	*messageResourceElement
}

func (e *MessageElementAudio) Tag() string {
	return "audio"
}

type MessageElementVedio struct {
	*messageResourceElement
}

func (e *MessageElementVedio) Tag() string {
	return "vedio"
}

type MessageElementFile struct {
	*messageResourceElement
}

func (e *MessageElementFile) Tag() string {
	return "file"
}
