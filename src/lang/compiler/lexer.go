package compiler

import "path"
import "strings"
import "unicode"
import "encoding/json"
import "regexp"

type Lex struct {
  Name   string
  Exp    string
  Line   uint64
  Type   string
  OName  string
  Dir    string
}

//length of each error expression
const EXPRESSION_LEN = 30

var keywords []map[string]string

var _ = json.Unmarshal(keywordJSON, &keywords)

func lexer(file, dirname, filename string) []Lex {

  var lex []Lex
  curExp := ""
  var line uint64 = 1

  for i := 0; i < len(file); i++ {

    if len(strings.TrimLeft(file[i:], " ")) == 0 {
      break
    }

    //detect a comment
    //single line comments are written as ;comment
    //like in assembly
    if strings.TrimLeft(file[i:], " ")[0] == ';' {

      var end = strings.Index(file[i:], "\n")

      //if there is no newline after the comment, just break the loop
      if end == -1 {
        break
      }

      i+=end
      line++
      continue
    }

    if file[i] == '\n' {
      line++
      continue
    }

    if unicode.IsSpace(rune(file[i])) { //if it is a whitespace, ignore it
      continue
    }

    //while the curExp is too long, pop the front
    for ;len(curExp) > EXPRESSION_LEN; curExp = curExp[1:] {}

    for _, v := range keywords {
      if testkey(v, file, i) {
        curExp+=v["remove"] + " "
        i+=len(v["remove"]) - 1

        lex = append(lex, Lex{
          Name: v["name"],
          Exp: curExp,
          Line: line,
          Type: v["type"],
          OName: v["remove"],
          Dir: path.Join(dirname, filename),
        })
        goto contOuter
      }

    }

    if strings.TrimLeft(file, " ")[i:][0] == '"' || strings.TrimLeft(file, " ")[i:][0] == '\'' || strings.TrimLeft(file, " ")[i:][0] == '`' { //detect a string
      qType := strings.TrimLeft(file, " ")[i:][0]
      value := ""
      escaped := false

      for i++; i < len(file); i++ {
        value+=string(strings.TrimLeft(file, " ")[i:][0])

        if !escaped && strings.TrimLeft(file, " ")[i:][0] == '\\' {
          escaped = true
          continue
        }

        if !escaped && strings.TrimLeft(file, " ")[i:][0] == qType {
          break
        }
        escaped = false
      }

      curExp+=value + " "
      line+=uint64(strings.Count(value, "\n"))
      lex = append(lex, Lex{
        Name: string(qType) + value,
        Exp: curExp,
        Line: line,
        Type: "expression value",
        OName: value,
        Dir: path.Join(dirname, filename),
      })
      goto contOuter
    } else if unicode.IsDigit(rune(strings.TrimLeft(file, " ")[i:][0])) || strings.TrimLeft(file, " ")[i:][0] == '+' || strings.TrimLeft(file, " ")[i:][0] == '-' || strings.TrimLeft(file, " ")[i:][0] == '.' {

      var positive = true

      if strings.TrimLeft(file, " ")[i:][0] == '+' {
        positive = true
        i++
      } else if strings.TrimLeft(file, " ")[i:][0] == '-' {
        positive = false
        i++
      }

      num := ""

      for o := i; unicode.IsDigit(rune(strings.TrimLeft(file, " ")[o:][0])) || strings.TrimLeft(file, " ")[o:][0] == '.'; o++ {
        num+=string(strings.TrimLeft(file, " ")[o:][0])
        if len(strings.TrimLeft(file, " ")[o + 1:]) == 0 {
          break
        }
      }
      i+=len(num) - 1

      if !positive {
        num = "-" + num
      }

      curExp+=num + " "

      lex = append(lex, Lex{
        Name: num,
        Exp: curExp,
        Line: line,
        Type: "expression value",
        OName: num,
        Dir: path.Join(dirname, filename),
      })
    } else {

      if unicode.IsSpace(rune(file[i:][0])) {
        continue
      }

      variable := ""

      for o := i; o < len(file); o++ {
        for _, v := range keywords {
          if testkey(v, file, o) || unicode.IsSpace(rune(file[o])) || file[o] == ';' /* it is a comment */ {
            goto break_var_loop
          }
        }

        variable+=string(file[o])
        i++
      }
      break_var_loop:

      curExp+=variable + " "
      variable = strings.TrimLeft(variable, " ")
      i--

      lex = append(lex, Lex{
        Name: "$" + variable,
        Exp: curExp,
        Line: line,
        Type: "expression value",
        OName: variable,
        Dir: path.Join(dirname, filename),
      })
    }

    contOuter: //continue the outer loop
  }

  //filter out newlines
  var newLex []Lex
  for _, v := range lex {
    if v.Type == "newlineN" {
      continue
    }
    if v.Name == "$false" || v.Name == "$true" || v.Name == "$undef" || v.Name == "$null" { //account for true, false, undef, and null values
      v.Name = v.Name[1:]
      newLex = append(newLex, v)
      continue
    }
    newLex = append(newLex, v)
  }
  lex = newLex
  newLex = nil

  //detect two operators back to back (which is an error)
  for k, v := range lex {
    if v.Type == "operation" && k + 1 < len(lex) && lex[k + 1].Type == "operation" {
      compilerErr("Cannot have two operations next to each other \nFound near this expression: " + lex[k + 1].Exp, path.Join(dirname, filename), lex[k + 1].Line)
    }
  }

  lex = term_inserter(tilde_inserter(insert_arrows(funcLex(lex))))

  return lex
}

func testkey(keyword map[string]string, file string, i int) bool {
  re, _ := regexp.Compile("^(" + keyword["pattern"] + ")")
  matched := re.MatchString(file[i:])

  if matched {
    if keyword["name"] == "+" || keyword["name"] == "-" {
      //because + and - can also be used as signs
      //if +/- comes after a keyword or operation, it must be a sign

      for _, v := range keywords {
        re, _ := regexp.Compile("(" + v["pattern"] + ")$")
        matched := re.MatchString(file[:i])

        if matched {
          return false
        }
      }

    }
  }

  return matched
}
