package t


import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// TypeScript lexer.
var TorqueScript = internal.Register(MustNewLexer(
	&Config{
		Name:      "TypeScript",
		Aliases:   []string{"ts", "tsx", "typescript"},
		Filenames: []string{"*.ts", "*.tsx"},
		MimeTypes: []string{"text/x-typescript"},
		DotAll:    true,
		EnsureNL:  true,
	},
	Rules{
		"commentsandwhitespace": {
			{`\s+`, Text, nil},
			{`//.*?\n`, CommentSingle, nil},
			{`/\*.*?\*/`, CommentMultiline, nil},
		},
		"root": {
			Include("commentsandwhitespace"),
			{`\+\+|--|~|&&|\?|:|\|\||\\(?=\n)|(<<|>>>?|==?|!=?|[-<>+*%&|^/])=?`, Operator, nil},
			{`[{(\[;,]`, Punctuation, nil},
			{`[})\].]`, Punctuation, nil},
			{`(for|foreach|foreach$|while|do|break|return|continue|switch|switch$|case|default|if|else|new)\b`, Keyword, nil},
			{`(datablock|singleton|function)\b`, KeywordDeclaration, nil},
			{`()\b`, KeywordReserved, nil},
			{`(true|false)\b`, KeywordConstant, nil},
			{`(%this)\b`, NameBuiltin, nil},
			{`[0-9][0-9]*\.[0-9]+([eE][0-9]+)?[fd]?`, LiteralNumberFloat, nil},
			{`0x[0-9a-fA-F]+`, LiteralNumberHex, nil},
			{`[0-9]+`, LiteralNumberInteger, nil},
		},
	},
))

