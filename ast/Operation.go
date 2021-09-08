package ast

import (
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/types"
	"github.com/tusklang/tusk/data"
	"github.com/tusklang/tusk/tokenizer"
)

type Operation struct {
	OpType string
	Token  *tokenizer.Token
}

func (o *Operation) Parse(lex []tokenizer.Token, i *int) error {

	o.Token = &lex[*i]
	o.OpType = lex[*i].Name

	return nil
}

func (o *Operation) Compile(compiler *Compiler, class *types.StructType, node *ASTNode, block *ir.Block) data.Value {

	var (
		lc = node.Left[0].Group.Compile(compiler, class, node.Left[0], block)
		rc = node.Right[0].Group.Compile(compiler, class, node.Right[0], block)
	)

	return compiler.OperationStore.RunOperation(lc, rc, o.OpType, block)
}
