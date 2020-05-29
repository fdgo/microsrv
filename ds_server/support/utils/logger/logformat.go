package logger

import (
	"bytes"
	"encoding/json"
	"strings"
)

func FormatJsonStr(instr string) string {
	var out bytes.Buffer
	json.Indent(&out, []byte(instr), "", "  ")

	return "\n" + out.String() + "\n"
}

func SerializeToJson(st interface{}) string {
	ba, _ := json.Marshal(st)
	jsonstr := string(ba)

	return jsonstr
}

func FormatStruct(inst interface{}) string {
	instr := SerializeToJson(inst)
	return FormatJsonStr(instr)
}

func UnserializeFromJson(jsonstr string, st interface{}) error {
	d := json.NewDecoder(strings.NewReader(jsonstr))
	d.UseNumber()
	return d.Decode(st)
}
