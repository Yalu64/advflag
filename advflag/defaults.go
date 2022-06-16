package advflag

// defaultArgSym defines the character/s where an argument starts
// -hello 123
// -> type [123]
var defaultArgSym = "-"

// defaultBoolSym defines the character/s where a bool argument starts
// --hello
// -> true
var defaultBoolSym = "--"

// Maximum length an argument can have
// ex: -123456789 test123
// would not exceed this limit because the characters after Prefix do not exceed defaultArgMaxCharLen
var defaultArgMaxCharLen = 1 << 16

// Maximum length an argument prefix can have
// ex: -123456789 test123
// would not exceed this limit because the characters after defaultArgSym do not exceed defaultArgPrefixMaxCharLen
var defaultArgPrefixMaxCharLen = 1 << 8

var AppName = FilePath()
var PrintSyntax bool = true
