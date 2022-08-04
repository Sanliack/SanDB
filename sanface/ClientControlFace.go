package sanface

type ClientControlFace interface {
	Set() SetControlFace
	Str() StrControlFace
}
