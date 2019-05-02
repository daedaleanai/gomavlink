package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"text/template"
)

var tmplfuncs = template.FuncMap{
	"lower":             strings.ToLower,
	"upper":             strings.ToUpper,
	"title":             strings.Title,
	"underscoreToCamel": underscoreToCamel,
	"gotype":            goType,
	"goscalartype":      goScalarType,
	"goarraysize":       goArraySize,
	"notabs":            notabs,
	"cscalartype":       cScalarType,
}

func underscoreToCamel(s string) string {
	parts := strings.Split(s, "_")
	for i, v := range parts {
		parts[i] = strings.Title(strings.ToLower(v))
	}
	return strings.Join(parts, "")
}

var wstospace = strings.NewReplacer("\t", " ", "\n", " ")

func notabs(s string) string { return wstospace.Replace(s) }

var reCType = regexp.MustCompile(`([a-z0-9_]+)(\[[0-9]+\])?`)

func goType(ctype string) (string, error) {
	parts := reCType.FindStringSubmatch(ctype)
	if len(parts) != 3 {
		return "", fmt.Errorf("Cannot parse %q as ctype([arraylen])", ctype)
	}
	var t string
	switch parts[1] {
	case "float":
		t = "float32"
	case "double":
		t = "float64"
	case "int8_t":
		t = "int8"
	case "char", "uint8_t", "uint8_t_mavlink_version", "array":
		t = "byte"
	case "int16_t":
		t = "int16"
	case "uint16_t":
		t = "uint16"
	case "int32_t":
		t = "int32"
	case "uint32_t":
		t = "uint32"
	case "int64_t":
		t = "int64"
	case "uint64_t":
		t = "uint64"
	default:
		return "", fmt.Errorf("Cannot parse %q as a ctype (invalid scalar part %q)", ctype, parts[1])
	}
	if parts[2] == "" {
		return t, nil
	}
	return fmt.Sprintf("%s%s", parts[2], t), nil
}

func scalarSize(ctype string) int {
	parts := reCType.FindStringSubmatch(ctype)
	if len(parts) != 3 {
		return 0
	}
	switch parts[1] {
	case "int8_t", "char", "uint8_t", "uint8_t_mavlink_version", "array":
		return 1
	case "int16_t", "uint16_t":
		return 2
	case "float", "int32_t", "uint32_t":
		return 4
	case "double", "int64_t", "uint64_t":
		return 8
	}
	return 0
}

func goScalarType(ctype string) (string, error) {
	parts := reCType.FindStringSubmatch(ctype)
	if len(parts) != 3 {
		return "", fmt.Errorf("Cannot parse %q as ctype([arraylen])", ctype)
	}
	var t string
	switch parts[1] {
	case "float":
		t = "float32"
	case "double":
		t = "float64"
	case "int8_t":
		t = "int8"
	case "char", "uint8_t", "uint8_t_mavlink_version", "array":
		t = "byte"
	case "int16_t":
		t = "int16"
	case "uint16_t":
		t = "uint16"
	case "int32_t":
		t = "int32"
	case "uint32_t":
		t = "uint32"
	case "int64_t":
		t = "int64"
	case "uint64_t":
		t = "uint64"
	default:
		return "", fmt.Errorf("Cannot parse %q as a ctype (invalid scalar part %q)", ctype, parts[1])
	}
	return t, nil
}

func goArraySize(ctype string) (int, error) {
	parts := reCType.FindStringSubmatch(ctype)
	if len(parts) != 3 {
		return 0, fmt.Errorf("Cannot parse %q as ctype([arraylen])", ctype)
	}
	if parts[2] == "" {
		return 0, nil
	}
	n, err := strconv.ParseUint(parts[2][1:len(parts[2])-1], 10, 8)
	if err != nil {
		return 0, err
	}
	return int(n), nil
}

func cScalarType(ctype string) (string, error) {
	parts := reCType.FindStringSubmatch(ctype)
	if len(parts) != 3 {
		return "", fmt.Errorf("Cannot parse %q as ctype([arraylen])", ctype)
	}
	if parts[1] == "uint8_t_mavlink_version" {
		return "uint8_t", nil
	}
	return parts[1], nil
}
