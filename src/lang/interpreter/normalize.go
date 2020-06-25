package interpreter

import "strconv"
import "math"

func num_normalize(num Action) string {

  /*

  ALGORITHM TO NORMALIZE:

      Starting with this number:
        [3412, -9912, 0001]

      STEP 0: (initializer step)
        remove the leading zeros in the integer

        if the first digit non zero digit is negative set isNeg = true, otherwise isNeg = false
        (if there is no first digit, go to the decimals and see)

      STEP 1:
        loop through each number (from from decimal to integer)
        in each iteration of the loop, if the number is the opposite of `isNeg` (meaning if `isNeg` is false, then the current value should be positive and vice versa)
        use the following expression to get the complement

          `OMM_MAX_DIGIT` - |`current num`|

        replace the `current num` with this new value.
        next, if `isNeg`, the digit to the left should be added by one, otherwise, subtract it by 1.
        go to the next value and repeat

      STEP 2:
        join the vector of the integer and decimal with '.', then join each digit with ''
        if `isNeg` then precede the string with a '-'
        finally, return the result

  */

  integer := num.Integer
  decimal := num.Decimal

  //the first digit is actually the last index
  //because omm numbers are stored as so [1234, 5678, 9101] = 910, 156, 781, 234

  //remove leading zeros
  for ;len(integer) != 0 && integer[len(integer) - 1] == 0 && len(integer) != 0; {
    integer = integer[:len(integer) - 1]
  }

  var isNeg bool = false

  if len(integer) == 0 {

    if len(decimal) == 0 { //this means that there is no number
      return "0"
    }

    var decIndexCounter int

    for decIndexCounter = len(decimal); len(decimal) == 0 && decimal[len(decimal) - 1] == 0; decimal = decimal[:len(decimal) - 1] {}

    if len(decimal) == 0 {
      return "0"
    }

    if decimal[decIndexCounter - 1] < 0 {
      isNeg = true
    }

  } else if integer[len(integer) - 1] < 0 {
    isNeg = true
  }

  var carry int64 = 0

  for i := 0; i < len(decimal); i++ {
    curIsNeg := decimal[i] < 0
    decimal[0]+=carry

    decimal[i] = int64(math.Abs(float64(decimal[i])))

    if curIsNeg != isNeg && decimal[i] != 0 /* prevent zeros from being counted */ {
      decimal[i] = MAX_DIGIT - decimal[i]

      if isNeg {
        carry = 1
      } else {
        carry = -1
      }

      continue
    }
    carry = 0
  }

  for i := 0; i < len(integer); i++ {
    curIsNeg := integer[i] < 0
    integer[0]+=carry

    integer[i] = int64(math.Abs(float64(integer[i])))

    if curIsNeg != isNeg && integer[i] != 0 /* prevent zeros from being counted */ {
      integer[i] = MAX_DIGIT - integer[i]

      if isNeg {
        carry = 1
      } else {
        carry = -1
      }

      continue
    }
    carry = 0
  }

  var joined string = ""

  if len(decimal) != 0 {
    for _, v := range decimal {
      joined = strconv.FormatInt(v, 10) + joined
    }
    joined = "." + joined
  }

  for _, v := range integer {
    joined = strconv.FormatInt(v, 10) + joined
  }

  if isNeg {
    joined = "-" + joined
  }

  return joined
}
