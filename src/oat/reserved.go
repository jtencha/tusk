package oat

const MAGIC = "OAT"

var reserved = map[string]rune{
	"seperate field" 		: 65536,
	"next action"           : 'అ',
	"escaper"        		: 0,
	"var"            		: 4,
	"log"            		: 5,
	"print"          		: 6,
	"if"             		: 7,
	"elif"           		: 8,
	"else"           		: 9,
	"condition"      		: 96,
	"while"          		: 11,
	"each"           		: 12,
	"function"       		: 97,
	"return"         		: 14,
	"await"          		: 15,
	"proto"          		: 16,
	"ovld"           		: 17,
	"let"            		: 18,
	"cast"           		: 19,
	"::"             		: 20,
	"+"              		: 21,
	"-"              		: 22,
	"*"              		: 23,
	"/"              		: 24,
	"%"              		: 25,
	"^"              		: 26,
	"=="             		: 27,
	"!="             		: 28,
	">"              		: 29,
	"<"              		: 30,
	">="             	    : 31,
	"<="             		: 81,
	"!"              		: 33,
	"&"              		: 34,
	"|"              		: 35,
	"=>"             		: 36,
	"<-"             		: 37,
	"<~"             		: 38,
	"++"             		: 39,
	"--"             		: 40,
	"+="             		: 41,
	"-="             		: 42, //how's "life"
	"*="             		: 43,
	"/="             		: 44,
	"%="             		: 45,
	"^="             		: 46,
	"break"          		: 47,
	"continue"       		: 48,
	"{"              		: 49,
	"("              		: 50,
	"c-hash"         		: 51,
	"r-hash"         		: 52,
	"c-array"        		: 53,
	"r-array"        		: 54,
	"string"         		: 55,
	"rune"           		: 56,
	"bool"           		: 57,
	"undef"          		: 58,
	"number"         		: 59,
	"variable"       		: 60,
	"varname start"  		: 61,
	"start multi action" 	: 62,
	"end multi action" 	    : 63,
	"hash key seperator"    : 66,
	"value seperator"       : 67,
	"make bool"             : 70,
	"make undef"            : 71,
	"make rune"             : 72,
	"make string"           : 73,
	"start number"          : 74,
	"end number"            : 75,
	"decimal spot"          : 76,
	"make c-array"          : 77,
	"make c-hash"           : 79,
	"start proto"           : 82,
	"end proto"             : 83,
	"start proto name"      : 84,
	"end proto name"        : 104,
	"start proto static"    : 85,
	"end proto static"      : 86,
	"start proto instance"  : 87,
	"end proto instance"    : 88,
	"start function"        : 89,
	"end function"          : 90,
	"seperate type-param"   : 93,
	"start params"          : 94,
	"new global"            : 10,
	"set global"            : 98,
	"start r-hash"          : 99,
	"end r-hash"            : 100,
	"start r-array"         : 101,
	"end r-array"           : 102,
	"param body split"      : 103,
}
