package main

import "strings"

func arrayContain(arr []string, sub string) bool {

  for i := 0; i < len(arr); i++ {
    if arr[i] == sub {
      return true
    }
  }

  return false;
}

func arrayContainInterface(arr []string, sub interface{}) bool {

  for i := 0; i < len(arr); i++ {
    if arr[i] == sub {
      return true
    }
  }

  return false;
}

func arrayContainInterfaceOperations(arr []string, sub interface{}) bool {

  for i := 0; i < len(arr); i++ {
    if arr[i] == sub.(string) {
      return true
    }
  }

  return false;
}

func arrayContain2Nest(arr [][]string, sub string) bool {

  for i := 0; i < len(arr); i++ {
    if arrayContain(arr[i], sub) {
      return true
    }
  }

  return false
}

func indexOf2Nest(sub string, arr [][]string) []int {
  for i := 0; i < len(arr); i++ {
    for o := 0; o < len(arr[i]); o++ {
      if arr[i][o] == sub {
        return []int{ i, o }
      }
    }
  }

  return []int{ -1, -1 }
}

func RepeatAdd(s string, times int) string {
  returner := ""

  for ;times > 0; times-- {
    returner+=s
  }

  return returner
}

func indexOf(sub string, data []string) int {
  for k, v := range data {
    if sub == v {
      return k
    }
  }
  return -1
}

func interfaceContain(inter []interface{}, sub interface{}) bool {
  for _, a := range inter {
    if a == sub {
      return true
    }
  }
  return false
}

func interfaceIndexOf(sub interface{}, inter []interface{}) int {
  for k, v := range inter {
    if sub == v {
      return k
    }
  }
  return -1
}

func interfaceContainOperations(inter []interface{}, sub interface{}) bool {

  loop:
  for _, a := range inter {

    switch a.(type) {
      case Action:
        continue loop
    }

    if a.(Lex).Name == sub {
      return true
    }
  }
  
  return false
}

func interfaceIndexOfOperations(sub interface{}, inter []interface{}) int {
  loop:
  for k, a := range inter {

    switch a.(type) {
      case Action:
        continue loop
    }

    if a.(Lex).Name == sub {
      return k
    }
  }

  return -1
}

func interfaceContainWithProcIndex(inter []interface{}, sub interface{}, indexes []int) bool {

  loop:
  for k, v := range inter {

    switch v.(type) {
      case Action:
        continue loop
    }

    if sub.(string) == v.(Lex).Name {

      for _, o := range indexes {
        if k == o {
          continue loop
        }

      }

      return true
    }
  }
  return false
}

func interfaceIndexOfWithProcIndex(sub interface{}, inter []interface{}, indexes []int) int {

  loop:
  for k, v := range inter {

    switch v.(type) {
      case Action:
        continue loop
    }

    if sub.(string) == v.(Lex).Name {

      for _, o := range indexes {
        if k == o {
          continue loop
        }

      }

      return k
    }
  }
  return -1
}

func interfaceContainForExp(inter []interface{}, _sub []string) bool {

  var sub []interface{}

  for _, v := range _sub {
    sub = append(sub, v)
  }

  cbCnt := 0
  glCnt := 0
  bCnt := 0
  pCnt := 0

  for o := 0; o < len(inter); o++ {

    v := inter[o]

    //prevent parenthesis after process declarations from being counted as expression parenthesis
    if o > 0 && (strings.HasPrefix(inter[o - 1].(Lex).Name, "$") || inter[o - 1].(Lex).Name == "]" || inter[o - 1].(Lex).Name == "process") && v.(Lex).Name == "(" {

      scbCnt := 0
      sglCnt := 0
      sbCnt := 0
      spCnt := 0

      for i := o; i < len(inter); i, o = i + 1, o + 1 {
        if inter[i].(Lex).Name == "{" {
          scbCnt++;
        }
        if inter[i].(Lex).Name == "}" {
          scbCnt--;
        }

        if inter[i].(Lex).Name == "[:" {
          sglCnt++;
        }
        if inter[i].(Lex).Name == ":]" {
          sglCnt--;
        }

        if inter[i].(Lex).Name == "[" {
          sbCnt++;
        }
        if inter[i].(Lex).Name == "]" {
          sbCnt--;
        }

        if inter[i].(Lex).Name == "(" {
          spCnt++;
        }
        if inter[i].(Lex).Name == ")" {
          spCnt--;
        }

        if scbCnt == 0 && sglCnt == 0 && sbCnt == 0 && spCnt == 0 {
          break
        }
      }

      continue
    }

    if v.(Lex).Name == "{" {
      cbCnt++;
    }
    if v.(Lex).Name == "}" {
      cbCnt--;
    }

    if v.(Lex).Name == "[:" {
      glCnt++;
    }
    if v.(Lex).Name == ":]" {
      glCnt--;
    }

    if v.(Lex).Name == "[" {
      bCnt++;
    }
    if v.(Lex).Name == "]" {
      bCnt--;
    }

    if v.(Lex).Name == "(" {
      pCnt++;
    }
    if v.(Lex).Name == ")" {
      pCnt--;
    }

    if cbCnt == 0 && glCnt == 0 && bCnt == 0 && pCnt == 0 {

      for _, i := range sub {
        if i == v.(Lex).Name {
          return true
        }
      }

    }
  }

  return false
}
