package types

type OmmString struct {
  string *string
  Length  uint64
}

func (str *OmmString) FromGoType(val string) {
  str.string = &val
  str.Length = uint64(len(val))
}

func (str OmmString) ToGoType() string {
  return *str.string
}

func (_ OmmString) ValueFunc() {}
