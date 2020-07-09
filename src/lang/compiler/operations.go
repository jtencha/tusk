package compiler

type Operation struct {
  Type        string
  Line        uint64
  Left       *Operation
  Right      *Operation
  Degree     *Operation
  Item        Item //in case there is no operation
}

var operations []map[string]func(exp []Item, index int, opType string)Operation

func operationIncludes(group []Item) bool {
  for _, v := range operations {
    for k := range v {
      for _, i := range group {
        if i.Token.Name == k {
          return true
        }
      }
    }
  }

  return false
}

func indexOper(group []Item, opers []string) (int, int) {

  for i := len(group) - 1; i >= 0; i-- {
    for k, v := range opers {
      if group[i].Token.Name == v {
        return i, k
      }
    }
  }

  return -1, -1
}

//operation function used for most operators (except assignment, not gate, similarity, function calls, etc...)
func normalOpFunc(exp []Item, index int, opType string) Operation {

  var (
    left = exp[:index]
    right = exp[index + 1:]
  )

  return Operation{
    Type: opType,
    Line: exp[index].Line,
    Left: &makeOperations([][]Item{ left })[0],
    Right: &makeOperations([][]Item{ right })[0],
  }
}

//operation function for both similarity operators (~~ and ~~~)
func similarityOpFunc(exp []Item, index int, opType string) Operation {

  hasColon := false //if it has a colon (to indicate a degree)
  var degExp []Item //expression of the degree
  var i int

  for i = index + 1; i < len(exp); i++ {
    if exp[i].Token.Name == ":" {
      hasColon = true
      break
    }
    degExp = append(degExp, exp[i])
  }

  if !hasColon {
    return normalOpFunc(exp, index, opType)
  }

  left := exp[:index]
  right := exp[index + i + 1:]

  return Operation{
    Type: opType,
    Line: exp[index].Line,
    Left: &makeOperations([][]Item{ left })[0],
    Right: &makeOperations([][]Item{ right })[0],
    Degree: &makeOperations([][]Item{ degExp })[0],
  }
}

//ODO is
// statement operator (~)
// assigner (:)
// boolean operations (except not gate)
// comparisons
// exponentiation
// mult, div, modulo
// add, subtract
// index operations and function calls
// not gate
// assignment operations

func makeOperations(groups [][]Item) []Operation {

  operations = []map[string]func(exp []Item, index int, opType string) Operation {
    map[string]func(exp []Item, index int, opType string) Operation {
      "~": normalOpFunc,
      "?": normalOpFunc,
      "cb-ob": normalOpFunc, //operator to connect a close brace to open brace (like this: while (true) <need operator here> {})
    },
    map[string]func(exp []Item, index int, opType string) Operation {
      ":": normalOpFunc,
    },
    map[string]func(exp []Item, index int, opType string) Operation {
      "&": normalOpFunc,
      "|": normalOpFunc,
    },
    map[string]func(exp []Item, index int, opType string) Operation {
      "=": normalOpFunc,
      "!=": normalOpFunc,
      ">": normalOpFunc,
      ">=": normalOpFunc,
      "<": normalOpFunc,
      "<=": normalOpFunc,
      "~~": similarityOpFunc,
      "~~~": similarityOpFunc,
    },
    map[string]func(exp []Item, index int, opType string) Operation {
      "^": normalOpFunc,
    },
    map[string]func(exp []Item, index int, opType string) Operation {
      "*": normalOpFunc,
      "/": normalOpFunc,
      "%": normalOpFunc,
    },
    map[string]func(exp []Item, index int, opType string) Operation {
      "+": normalOpFunc,
      "-": normalOpFunc,
    },
    map[string]func(exp []Item, index int, opType string) Operation {
      "::": normalOpFunc,
      "sync": normalOpFunc,
      "async": normalOpFunc,
    },
    map[string]func(exp []Item, index int, opType string) Operation {
      "!": func(exp []Item, index int, opType string) Operation {
        return Operation{
          Type: opType,
          Line: exp[index].Line,
          Right: &makeOperations([][]Item{ exp[index + 1:] })[0],
        }
      },
    },
    map[string]func(exp []Item, index int, opType string) Operation {
      "++": func(exp []Item, index int, opType string) Operation {
        return Operation{
          Type: opType,
          Line: exp[index].Line,
          Left: &makeOperations([][]Item{ exp[:index] })[0],
        }
      },
      "--": func(exp []Item, index int, opType string) Operation {
        return Operation{
          Type: opType,
          Line: exp[index].Line,
          Left: &makeOperations([][]Item{ exp[:index] })[0],
        }
      },
      "+=": normalOpFunc,
      "-=": normalOpFunc,
      "*=": normalOpFunc,
      "/=": normalOpFunc,
      "%=": normalOpFunc,
      "^=": normalOpFunc,
    },
  }

  var newGroups []Operation

  for _, v := range groups {

    if !operationIncludes(v) {
      newGroups = append(newGroups, Operation{
        Type: "none",
        Line: v[0].Line,
        Item: v[0],
      })
    }

    for _, val := range operations {

      var opers []string
      var funcs []func(exp []Item, index int, opType string) Operation

      for oper, function := range val {
        opers = append(opers, oper)
        funcs = append(funcs, function)
      }

      indexOfOper, operNum := indexOper(v, opers)
      if indexOfOper != -1 {
        newGroups = append(newGroups, funcs[operNum](v, indexOfOper, opers[operNum]))
        goto breakOuter
      }

    }

    breakOuter:
  }

  return newGroups
}
