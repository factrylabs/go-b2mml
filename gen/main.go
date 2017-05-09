package main

import (
	"bytes"
	"fmt"
	"go/format"
	"io/ioutil"
	"regexp"
	"strings"
)

func main() {

	// Load the source file.

	src, err := ioutil.ReadFile("../vendor/B2MML.NET/b2mml.cs")
	if err != nil {
		panic(err)
	}

	// Create a buffer for storing the generated code.

	var out bytes.Buffer

	// Print a warning that the package was auto generated.

	fmt.Fprintf(&out, "// WARNING: This package was auto generated.\n")
	fmt.Fprintf(&out, "// Changes to this file may cause incorrect behavior and\n")
	fmt.Fprintf(&out, "// will be lost if the code is regenerated.\n\n")

	// Print the package name.

	fmt.Fprintf(&out, "package b2mml\n\n")

	// Print the necessary imports.

	fmt.Fprintf(&out, "import (\n  \"encoding/xml\"\n  \"time\"\n)\n\n")

	// Print the generic XML node type.

	fmt.Fprintf(&out, "type XmlNode struct {\n")
	fmt.Fprintf(&out, "\tXmlName xml.Name\n")
	fmt.Fprintf(&out, "\tContent []byte `xml:\",innerxml\"`\n")
	fmt.Fprintf(&out, "\tNodes   []XmlNode `xml:\",any\"`\n")
	fmt.Fprintf(&out, "}\n\n")

	// Define Regex for the c# enums.

	enumEx := regexp.MustCompile(`public\senum\s(\S+)[\s\S]*?\n}`)

	// Find and loop over all c# enums. Create simple string types for them.
	// TODO: What is a better way to handle this?

	enums := enumEx.FindAllStringSubmatch(string(src), -1)
	for _, enum := range enums {
		fmt.Fprintf(&out, "type %v string\n\n", enum[1])
	}

	// Define regexes for the c# classes and their public fields.

	classEx := regexp.MustCompile(`public\spartial\sclass\s([\S]+)[\s:\s]*?([\S]*?)\s{[\s\S]*?\n}`)
	fieldEx := regexp.MustCompile(`(\[\S+\])*\s*?public\s(\S+)\s(\S+)\s+{`)

	// Find and loop over all c# classes and their public fields, create a
	// b2mmlType from them and print it as a correct go struct.

	classes := classEx.FindAllStringSubmatch(string(src), -1)
	for _, class := range classes {
		b := b2mmlType{
			name:     class[1],
			baseType: class[2],
		}
		fields := fieldEx.FindAllStringSubmatch(class[0], -1)
		for _, field := range fields {
			b.fields = append(b.fields, b2mmlField{field[3], field[2], field[1]})
		}
		fmt.Fprintf(&out, "%v\n\n", b.toGoStruct())
	}

	// Format the generated code using the Go format library.

	fout, err := format.Source(out.Bytes())
	if err != nil {
		panic(err)
	}

	// Write to the main b2mml.go file

	ioutil.WriteFile("../b2mml.go", fout, 0644)

}

// b2mmlType represents a type from the B2MML schema specification.
type b2mmlType struct {
	name     string
	baseType string
	fields   []b2mmlField
}

// toGoStruct takes the b2mml type and returns a valid Go type as a string.
func (b *b2mmlType) toGoStruct() (out string) {
	out += fmt.Sprintf("// %v was auto-generated. Do not change.\n", b.name)
	out += fmt.Sprintf("type %v struct {\n  %v\n", b.name, b.baseType)
	for _, field := range b.fields {
		out += fmt.Sprintf("%v %v %v\n", field.goName(),
			field.goDataType(b.name), field.goTags())
	}
	out += "}"
	return
}

// b2mmlField represents a field with its datatype and the original c# XML
// annotation.
type b2mmlField struct {
	name       string
	dataType   string
	xmlContext string
}

// goName returns the original field name with a capital first letter, as the
// xml attributes in the schema start with a lowercase letter, and Go requires
// a capital first letter Go to make the field public.
func (b *b2mmlField) goName() string {
	return strings.ToUpper(string(b.name[0])) + string(b.name[1:])
}

// goDataType returns the c# datatype converted to a valid Go datatype.
func (b *b2mmlField) goDataType(containingTypeName string) string {
	dataType := b.dataType
	if strings.Contains(b.dataType, "[]") {
		dataType = "[]" + strings.Replace(b.dataType, "[]", "", -1)
	}

	// check for:
	//   - c# datatypes that are different from Go's.
	//   - recursive datatypes. These should be pointers.
	switch dataType {
	case "[]System.Xml.XmlElement":
		return "[]XmlNode"
	case "System.DateTime":
		return "time.Time"
	case "decimal":
		return "float64"
	case "object":
		return "interface{}"
	case "[]object":
		return "[]interface{}"
	case containingTypeName:
		return "*" + dataType
	}
	return dataType
}

// goXMLTags checks the original c# XML annotation and returns necessare go
// struct tags for XML (un)marshalling and for JSON (un)marshalling. In case of
// attributes (which have a lowercase first letter), the xml name is explicitly
// mentioned in the tag.
func (b *b2mmlField) goTags() (tags string) {
	tags = "`json:\",omitempty\""
	switch {
	case strings.Contains(b.xmlContext, "XmlAttributeAttribute"):
		tags += fmt.Sprintf(" xml:\"%v,attr,omitempty\"", b.name)
	case strings.Contains(b.xmlContext, "XmlTextAttribute"):
		tags += fmt.Sprintf(" xml:\",chardata\"")
	default:
		tags += fmt.Sprintf(" xml:\",omitempty\"")
	}
	tags += "`"
	return
}
