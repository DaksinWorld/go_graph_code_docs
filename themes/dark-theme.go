package themes

import "github.com/DaksinWorld/go_graph_code_docs/structs"

var NodeTemplate = []structs.Attr{
	{Key: "height", Value: "2"},
	{Key: "shape", Value: "box"},
	{Key: "style", Value: "rounded,filled"},
	{Key: "fillcolor", Value: "#1F1F1F"},
	{Key: "fontcolor", Value: "#FFFFFF"},
	{Key: "fontname", Value: "Montserrat"},
	{Key: "fontsize", Value: "10"},
	{Key: "color", Value: "black"},
	{Key: "width", Value: "2"},
}
var EdgeTemplate = []structs.Attr{
	{Key: "color", Value: "#CCCCCC"},
	{Key: "fontcolor", Value: "#CCCCCC"},
	{Key: "fontname", Value: "Montserrat"},
	{Key: "fontweight", Value: "700"},
	{Key: "fontsize", Value: "10"},
}

var DarkTheme = structs.Theme{Name: "DarkTheme", CollOfAttrs: [][]structs.Attr{NodeTemplate, EdgeTemplate}}
