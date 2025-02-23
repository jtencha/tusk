package ast

import (
	"github.com/llir/llvm/ir"
	"github.com/tusklang/tusk/data"
	"github.com/tusklang/tusk/tokenizer"
)

type WhileStatement struct {
	Condition []*ASTNode
	Body      []*ASTNode
}

func (ws *WhileStatement) Parse(lex []tokenizer.Token, i *int) error {
	return ifwhileParse(ws, lex, i)
}

func (ws *WhileStatement) SetCond(g []*ASTNode) {
	ws.Condition = g
}

func (ws *WhileStatement) SetBody(g []*ASTNode) {
	ws.Body = g
}

func (ws *WhileStatement) Type() string {
	return "while"
}

func (ws *WhileStatement) Compile(compiler *Compiler, class *data.Class, node *ASTNode, function *data.Function) data.Value {

	wscond := function.LLFunc.NewBlock("") //create a block to determine if the loop should continue (condition)
	function.ActiveBlock.NewBr(wscond)
	function.ActiveBlock = wscond

	wsbod := function.LLFunc.NewBlock("") //block to store the body of the while loop
	wsbod.NewBr(wscond)
	rest := function.LLFunc.NewBlock("") //block to store the rest of the code (after this while statement)

	cond := ws.Condition[0].Group.Compile(compiler, class, ws.Condition[0], function)
	wscond.NewCondBr(cond.LLVal(wscond), wsbod, rest)

	function.ActiveBlock = wsbod

	gotoCond := ir.NewBr(wscond)
	function.PushTermStack(gotoCond)

	ws.Body[0].Group.Compile(compiler, class, ws.Body[0], function)

	function.ActiveBlock = rest

	//if the pushed `goto` to the term stack was not used, pop it still
	if function.LastTermStack() == gotoCond {
		function.PopTermStack()
	}

	if val := function.PopTermStack(); val != nil {
		function.ActiveBlock.Term = val
	}

	return nil
}
