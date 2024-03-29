package model

var CodeFileExt map[string]string

func init() {
	CodeFileExt = make(map[string]string, 30)
	CodeFileExt[".c"] = "c"
	CodeFileExt[".cpp"] = "cpp"
	CodeFileExt[".cs"] = "csharp"
	CodeFileExt[".dart"] = "dart"
	CodeFileExt[".go"] = "go"
	CodeFileExt[".hs"] = "haskell"
	CodeFileExt[".ini"] = "ini"
	CodeFileExt[".java"] = "java"
	CodeFileExt[".js"] = "javascript"
	CodeFileExt[".json"] = "json"
	CodeFileExt[".kt"] = "kotlin"
	CodeFileExt[".lua"] = "lua"
	CodeFileExt[".ml"] = "ocaml"
	CodeFileExt[".pl"] = "perl"
	CodeFileExt[".ps1"] = "powershell"
	CodeFileExt[".py"] = "python"
	CodeFileExt[".rb"] = "ruby"
	CodeFileExt[".rs"] = "rust"
	CodeFileExt[".sol"] = "solidity"
	CodeFileExt[".swift"] = "swift"
	CodeFileExt[".ts"] = "typescript"
	CodeFileExt[".v"] = "verilog"
	CodeFileExt[".vue"] = "vue"
	CodeFileExt[".xml"] = "xml"
	CodeFileExt[".yaml"] = "yaml"
	CodeFileExt[".yml"] = "yaml"
	CodeFileExt[".zig"] = "zig"
}
