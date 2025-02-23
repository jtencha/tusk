package tokenizer

type TokenItem struct {
	regexp    string
	tokentype string
}

var tokenlist = []TokenItem{

	{("\\;"), "terminator"}, //semicolon
	{("\\,"), "terminator"}, //comma

	/************ whitespace ************/
	{("\\n"), "newline"},       //newline
	{("\\s{1}"), "whitespace"}, //whitespace
	/************************************/

	/************ keywords ************/
	{("fn(?=(\\z|\\(|\\{|\\s+))"), "fn"},               //fn
	{("return(?![a-zA-Z$_0-9])"), "return"},            //return
	{("var(?=\\z|\\s+)"), "var"},                       //var
	{("if(?=(\\z|\\(|\\{|\\s+))"), "if"},               //if
	{("else(?=(\\z|\\(|\\{|\\s+))"), "else"},           //else
	{("while(?=(\\z|\\(|\\{|\\s+))"), "while"},         //while
	{("pub(?![a-zA-Z$_0-9])"), "pub"},                  //pub
	{("prv(?![a-zA-Z$_0-9])"), "prv"},                  //prv
	{("prt(?![a-zA-Z$_0-9])"), "prt"},                  //prt
	{("stat(?![a-zA-Z$_0-9])"), "stat"},                //stat
	{("link(?![a-zA-Z$_0-9])"), "link"},                //link
	{("construct(?=(\\z|\\(|\\{|\\s+))"), "construct"}, //construct
	{("this(?![a-zA-Z$_0-9])"), "this"},                //this
	/**********************************/

	/************ braces ************/
	{("\\("), "("}, //opening parenthesis
	{("\\)"), ")"}, //closing parenthesis
	{("\\{"), "{"}, //opening curly brace
	{("\\}"), "}"}, //closing curly brace
	/********************************/

	/************ operators ************/
	{("\\-\\>"), "operation"}, // ->
	{("\\+"), "operation"},    // +
	{("\\-"), "operation"},    // -
	{("\\*"), "operation"},    // *
	{("\\/"), "operation"},    // /
	{("\\=\\="), "operation"}, // ==
	{("\\>"), "operation"},    // >
	{("\\>\\="), "operation"}, // >=
	{("\\<"), "operation"},    // <
	{("\\<\\="), "operation"}, // <=
	{("\\="), "operation"},    // =
	{("\\:"), "operation"},    // :
	{("\\."), "operation"},    // .
	{("\\~"), "operation"},    // ~
	{("\\&"), "operation"},    // &
	{("\\|"), "operation"},    // |
	{("\\^"), "operation"},    // ^
	/***********************************/

	/************ misc ************/
	{"null(?![a-zA-Z$_0-9])", "null"},                          //null value
	{"([\"])((\\\\{2})*|(.*?[^\\\\](\\\\{2})*))\\1", "string"}, //string value https://stackoverflow.com/a/17231632/10696946
	{"([+-]*[0-9]*\\.[0-9]*)", "float"},                        //floating literal
	{"([+-]*\\d+)", "int"},                                     //integer literal
	{"([a-zA-Z$_][a-zA-Z$_0-9]*)", "varname"},                  //variable
	/******************************/
}
