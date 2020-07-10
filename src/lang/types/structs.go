package types

//export Action
type Action struct {
  Type            string
  Name            string
  Value           OmmType
  ExpAct        []Action
  Indexes     [][]Action

  //stuff for operations

  First         []Action
  Second        []Action
  Degree        []Action

  //////////////////////

  Access          string

  //stuff to panic errors and give stack

  File            string
  Line            uint64

  //////////////////////////////////////
}

type Variable struct {
  Type      string
  Value     OmmType
  GoProc    func(args [][]Action, cli_params CliParams, vars map[string]Variable, this_vals []Action, dir string) OmmType
}

type Returner struct {
  Variables map[string]Variable
  Exp       OmmType
  Type      string
}
