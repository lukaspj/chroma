package t


import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// TorqueScript lexer.
var TorqueScript = internal.Register(MustNewLexer(
	&Config{
		Name:      "TorqueScript",
		Aliases:   []string{"tscript", "torquescript"},
		Filenames: []string{"*.tscript", "*.mis", "*.gui"},
		MimeTypes: []string{"text/x-torquescript"},
		DotAll:    true,
		EnsureNL:  true,
	},
	Rules{
		"commentsandwhitespace": {
			{`\s+`, Text, nil},
			{`//.*?\n`, CommentSingle, nil},
			{`/\*.*?\*/`, CommentMultiline, nil},
		},
		"paramlist": {
			Include("commentsandwhitespace"),
			{`%this`, NameBuiltin, nil},
			{`%(?!this)\w+`, NameVariable, nil},
			{`[(),]`, Punctuation, nil},
		},
		"arglist": {
			{`,`, Punctuation, nil},
			{`\)`, Punctuation, Pop(1)},
			Include("expressions"),
		},
		"objname": {
			Include("commentsandwhitespace"),
			{`[a-zA-Z_0-9]+`, NameEntity, nil},
			{`[():]`, Punctuation, nil},
		},
		"entity": {
			{`%this`, NameBuiltin, nil},
			{`%(?!this)\w+`, NameVariable, nil},
			{`\$[\w:\[\]]+`, NameVariableGlobal, nil},
			{`[a-zA-Z_0-9:]+`, NameEntity, nil},
		},
		"arrayaccessor": {
			{`\]`, Punctuation, Pop(1)},
			Include("expressions"),
		},
		"accessors": {
			{`\(`, Punctuation, Push("arglist")},
			{`(\.)([a-zA-Z0-9_]+)`, ByGroups(Punctuation, NameAttribute), nil},
			{`\[`, Punctuation, Push("arrayaccessor")},
			{`(-?->)([a-zA-Z0-9_]+)`, ByGroups(Punctuation, NameAttribute), nil},
			Default(Pop(1)),
		},
		"string": {
			{`"`, LiteralString, Pop(1)},
			{`\\([\\abfnrtv"\']|x[a-fA-F0-9]{2,4}|u[a-fA-F0-9]{4}|U[a-fA-F0-9]{8}|[0-7]{1,3})`, LiteralStringEscape, nil},
			{`[^\\"\n]+`, LiteralString, nil},
			{`\\\n`, LiteralString, nil},
			{`\\`, LiteralString, nil},
		},
		"tag": {
			{`'`, LiteralStringSingle, Pop(1)},
			{`[^\\'\n]+`, LiteralStringSingle, nil},
		},
		"expressions": {
			{`(true|false)`, NameBuiltin, nil},
			{`(SPC|NL|TAB|@)`, OperatorWord, nil},
			{`"`, LiteralString, Push("string")},
			{`'`, LiteralStringSingle, Push("tag")},
			{`([%\$]?[a-zA-Z0-9_:]+)(\s*)(\()`, ByGroups(NameFunction, Text, Punctuation), Push("accessors", "arglist")},
			{`([%\$]?[a-zA-Z0-9_:]+)`, UsingSelf("entity"), Push("accessors")},

			{`(==|!=|>=|<=|&&|\|\||::|--|\+\+|\$=|!\$=|<<|>>|\+=|-=|\*=|/=|%=|&=|^=|\|=|<<=|>>=|->|-->|\?|\+|-|\*|/|<|>|\||!|&|%|\^|~|=)`, Operator, nil},
			// {`(==|!=|>=|<=|&&|\|\||::|--|\+\+|\$=|!\$=|<<|>>|\+=|-=|\*=|/=|%=|&=|^=|\|=|<<=|>>=|->|-->|\?|\+|-|\*|/|<|>|\||!|&|%|^|~)`, Operator, nil},
			// {`[~!%^&*+=|?:<>/-$]`, Operator, nil},

			{`\s+`, Text, nil},
			{`(\d+\.\d*|\.\d+|\d+)[eE][+-]?\d+[LlUu]*`, LiteralNumberFloat, nil},
			{`(\d+\.\d*|\.\d+|\d+[fF])[fF]?`, LiteralNumberFloat, nil},
			{`0x[0-9a-fA-F]+[LlUu]*`, LiteralNumberHex, nil},
			{`0[0-7]+[LlUu]*`, LiteralNumberOct, nil},
			{`\d+[LlUu]*`, LiteralNumberInteger, nil},
		},
		"declaration": {
			{`(datablock|singleton)(\s*)([a-zA-Z0-9:_]+)(\s*\([^)]*\)[^{]*)(\{)`, ByGroups(KeywordDeclaration, Text, NameEntity, UsingSelf("objname"), Punctuation), Push("declarationbody")},
		},
		"logic-statement-condition": {
			Include("commentsandwhitespace"),
			{`\(`, Punctuation, nil},
			{`\)`, Punctuation, Pop(1)},
			{`;`, Punctuation, nil},
			Include("expressions"),
		},
		"statement": {
			Include("declaration"),
			Include("commentsandwhitespace"),
			{`(if|else|switch|switch\$|for|foreach|foreach\$)`, Keyword, Push("logic-statement-condition")},
			{`\{`, Punctuation, Push("functionbody")},
			{`\}`, Punctuation, Pop(1)},
			Include("expressions"),
			{`;`, Punctuation, Pop(1)},
		},
		"declarationbody": {
			Include("commentsandwhitespace"),
			Include("declaration"),
			{`([a-zA-Z0-9_\[\]]+)(\s*=\s*[^;]+)(;)`, ByGroups(NameAttribute, UsingSelf("expressions"), Punctuation), nil},
			{`\}`, Punctuation, Pop(1)},
		},
		"functionbody": {
			Include("commentsandwhitespace"),
			Include("statement"),
			Default(Pop(1)),
		},
		"root": {
			Include("commentsandwhitespace"),
			{`(function)(\s*)([a-zA-Z_0-9:]+)(\s*\([^)]*\))([^{]*)(\{)`, ByGroups(KeywordDeclaration, Text, NameFunction, UsingSelf("paramlist"), UsingSelf("root"), Punctuation), Push("functionbody")},
			Default(Push("statement")),
		},
	},
))

