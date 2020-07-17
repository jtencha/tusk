package types

import "strings"

type OmmHash struct {
  Hash    map[string]*OmmType
  Length  uint64
}

func (hash OmmHash) At(idx string) *OmmType {

  if _, exists := hash.Hash[idx]; !exists {
    var undef OmmType = OmmUndef{}
    hash.Hash[idx] = &undef
  }

  return hash.Hash[idx]
}

func (hash *OmmHash) Set(idx string, val OmmType) {

  if hash.Hash == nil {
    hash.Hash = make(map[string]*OmmType)
  }

  if _, exists := hash.Hash[idx]; !exists {
    hash.Length++
  }

  hash.Hash[idx] = &val
}

func (hash OmmHash) Exists(idx string) bool {
  _, exists := hash.Hash[idx]
  return exists
}

func (hash OmmHash) Format() string {

  return func() string {

    if len(hash.Hash) == 0 {
      return "[::]"
    }

    var formatted = "[:"

    for k, v := range hash.Hash {

      vFormatted := (*v).Format()

      switch (*v).(type) {
        case OmmHash: //if it is another hash, add the indents
          if vFormatted != "[::]" {
            newlineSplit := strings.Split(vFormatted, "\n")

            vFormatted = ""

            for _, val := range newlineSplit {
              vFormatted+=strings.Repeat(" ", 2) + val + "\n"
            }

            vFormatted = strings.TrimSpace(vFormatted) //remove the trailing \n (because an extra was added) and the leading two spaces (because it will be on the same line)
          }
      }

      formatted+="\n" + strings.Repeat(" ", 2) + k + ": " + vFormatted + ","
    }

    return formatted + "\n:]"
  }() //staring with 2
}

func (hash OmmHash) Type() string {
  return "hash"
}

func (hash OmmHash) TypeOf() string {
  return hash.Type()
}
