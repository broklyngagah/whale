package utils

import (
	"bytes"
	"compress/flate"
	"compress/gzip"
	"encoding/base64"
	"encoding/gob"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"

	"encoding/json"
	"unicode/utf8"

	"go.uber.org/zap"

	"runtime"
	"math/rand"
	"math"
	"carp.cn/whale/zaplogger"
)

func DumpsMapStringInterface(data map[string]interface{}) string {
	bytes, err := json.Marshal(&data)
	if err != nil {
		zaplogger.Error("Dumps",
			zap.Reflect("", data), zap.String(", error:", err.Error()))
		return ""
	}
	zaplogger.Debug("Dumps", zap.String("", string(bytes)))
	return string(bytes)
}

func SubString(str string, begin, length int) (substr string) {
	// 将字符串的转换成[]rune
	rs := []rune(str)
	lth := len(rs)

	// 简单的越界判断
	if begin < 0 {
		begin = 0
	}
	if begin >= lth {
		begin = lth
	}
	end := begin + length
	if end > lth {
		end = lth
	}

	// 返回子串
	return string(rs[begin:end])
}

func LoadsMapStringInterface(str string, data map[string]interface{}) error {
	err := json.Unmarshal([]byte(str), &data)
	return err
}

func Load2Obj(filePath string, obj interface{}) error {
	f, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer f.Close()
	enc := gob.NewDecoder(f)
	return enc.Decode(obj)
}

func Dump2File(filePath string, obj interface{}) error {
	os.Remove(filePath)
	f, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer f.Close()
	enc := gob.NewEncoder(f)
	return enc.Encode(obj)
}

func DecodeZip(b []byte) ([]byte, error) {
	reader := flate.NewReader(bytes.NewReader(b))
	return ioutil.ReadAll(reader)
}

func EncodeZip(b []byte) ([]byte, error) {
	var buf bytes.Buffer

	writer, err := flate.NewWriter(&buf, flate.BestSpeed)
	if err != nil {
		return b, err
	}
	defer writer.Close()
	_, err = writer.Write(b)
	if err != nil {
		return b, err
	}
	err = writer.Flush()
	if err != nil {
		return b, err
	}
	return buf.Bytes(), nil
}

func DecodeGzip(b []byte) ([]byte, error) {
	reader, err := gzip.NewReader(bytes.NewReader(b))
	if err != nil {
		return nil, err
	}

	defer reader.Close()
	return ioutil.ReadAll(reader)
}

func EncodeGzip(b []byte) ([]byte, error) {
	var buf bytes.Buffer

	writer, err := gzip.NewWriterLevel(&buf, gzip.BestSpeed)
	if err != nil {
		return b, err
	}
	defer writer.Close()
	_, err = writer.Write(b)
	if err != nil {
		return b, err
	}
	err = writer.Flush()
	if err != nil {
		return b, err
	}
	return buf.Bytes(), nil
}

func DecodeBase64(b []byte) ([]byte, error) {
	data, err := base64.StdEncoding.DecodeString(string(b))
	if err != nil {
		return data, err
	}
	return data, nil
}

func EncodeBase64(data []byte) []byte {
	str := base64.StdEncoding.EncodeToString(data)
	return []byte(str)
}

func GetSqlInStr(someList []int) string {

	str := GetStrFromIntList(someList)
	if strings.TrimSpace(str) == "" {
		return "0"
	} else {
		return str
	}
}

func GetStrIntList(str string, split ...string) []int {
	var sper string
	if len(split) == 1 {
		sper = split[0]
	} else {
		sper = ","
	}
	intList := make([]int, 0)
	for _, v := range strings.Split(str, sper) {
		intV, _ := strconv.Atoi(v)
		intList = append(intList, intV)
	}
	return intList
}

func GetStrFromIntList(data []int) string {
	strList := make([]string, 0)
	for _, v := range data {
		strList = append(strList, strconv.Itoa(v))
	}
	return strings.Join(strList, ",")
}

func GetStrFloatList(str string) []float64 {
	floatList := make([]float64, 0)
	for _, v := range strings.Split(str, ",") {
		floatV, _ := strconv.ParseFloat(v, 10)
		floatList = append(floatList, floatV)
	}
	return floatList
}

func GetIntListFromStrList(strlist []string) []int {

	intList := make([]int, 0, len(strlist))

	for _, str := range strlist {
		n, err := strconv.Atoi(str)
		if nil == err {
			intList = append(intList, n)
		}

	}

	return intList
}

func ReadFile(fileName string) ([]byte, error) {
	bytes, err := ioutil.ReadFile(fileName)
	return bytes, err

}

func WriteFile(fileName string, content []byte) error {
	return ioutil.WriteFile(fileName, content, os.ModePerm)
}

func IsInIntSlice(s []int, seach int) bool {
	for _, v := range s {
		if v == seach {
			return true
		}
	}
	return false
}

func IntSliceDelete(s []int, index int) []int {
	if index == 0 {
		return s[1:]
	}

	if index == len(s)-1 {
		return s[:index]
	}

	return append(s[:index], s[index+1:]...)
}

func FilterEmoji(content string) string {
	new_content := ""
	for _, value := range content {
		_, size := utf8.DecodeRuneInString(string(value))
		if size <= 3 {
			new_content += string(value)
		}
	}
	return new_content
}

func Struct2Map(st interface{}) map[string]interface{} {
	v := reflect.ValueOf(st)

	ma := map[string]interface{}{}

	if v.Kind() != reflect.Ptr {
		return ma
	}

	e := v.Elem()
	t := e.Type()

	for i := 0; i < e.NumField(); i++ {
		field := t.Field(i)
		key := field.Tag.Get("json")

		if len(key) <= 0 {
			key = field.Name
		}
		ma[field.Name] = e.Field(i).Interface()
	}

	return ma

}

func Struct2MapString(st interface{}) map[string]string {
	v := reflect.ValueOf(st)

	ma := map[string]string{}

	if v.Kind() != reflect.Ptr {
		return ma
	}

	e := v.Elem()
	t := e.Type()

	for i := 0; i < e.NumField(); i++ {
		field := t.Field(i)
		key := field.Tag.Get("json")

		if len(key) <= 0 {
			key = field.Name
		}
		ma[key] = formatAtom(e.Field(i))
	}

	return ma

}

// Any formats any value as a string.
func Any(value interface{}) string {
	return formatAtom(reflect.ValueOf(value))
}

// formatAtom formats a value without inspecting its internal structure.
func formatAtom(v reflect.Value) string {
	switch v.Kind() {
	case reflect.Invalid:
		return "invalid"
	case reflect.Int, reflect.Int8, reflect.Int16,
		reflect.Int32, reflect.Int64:
		return strconv.FormatInt(v.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16,
		reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return strconv.FormatUint(v.Uint(), 10)
		// ...floating-point and complex cases omitted for brevity...
	case reflect.Float32, reflect.Float64:
		return strconv.FormatFloat(v.Float(), 'f', -1, 64)
	case reflect.Bool:
		return strconv.FormatBool(v.Bool())
	case reflect.String:
		return v.String()
	case reflect.Struct:
		switch v.Type().String() {
		case "time.Time":
			return v.Interface().(time.Time).Format(time.RFC3339)
		default:
			return fmt.Sprintf("%v", v)
		}

	case reflect.Chan, reflect.Func, reflect.Ptr, reflect.Slice, reflect.Map:
		return v.Type().String() + " 0x" +
			strconv.FormatUint(uint64(v.Pointer()), 16)
	default: // reflect.Array, reflect.Struct, reflect.Interface
		return v.Type().String() + " value"
	}
}

func MapString2Struct(ma map[string]string, st interface{}) bool {
	v := reflect.ValueOf(st)

	if v.Kind() != reflect.Ptr {
		return false
	}

	e := v.Elem()
	t := e.Type()

	matchField := false

	for i := 0; i < e.NumField(); i++ {
		field := t.Field(i)

		fe := e.Field(i)
		key := field.Tag.Get("json")
		if len(key) <= 0 {
			key = field.Name
		}

		if v, found := ma[key]; found {
			matchField = true
			setField(&fe, v)
		}
	}

	return matchField
}

func setField(refV *reflect.Value, value string) {

	switch refV.Kind() {
	case reflect.Bool:
		var v bool
		v, err := strconv.ParseBool(value)
		if err == nil {
			refV.SetBool(v)
		}
	case reflect.Int, reflect.Int16, reflect.Int32, reflect.Int64:
		var v int64
		v, err := strconv.ParseInt(value, 0, refV.Type().Bits())
		if err == nil {
			refV.SetInt(v)
		}

	case reflect.Float64:
		v, err := strconv.ParseFloat(value, refV.Type().Bits())
		if err == nil {
			refV.SetFloat(v)
		}
	case reflect.String:
		refV.SetString(value)
	case reflect.Struct:
		switch refV.Type().String() {
		case "time.Time":
			v, err := time.Parse(time.RFC3339, value)

			if nil == err {
				refV.Set(reflect.ValueOf(v))
			}

		default:

		}
	}
}

func GetStringTraditionalByteNum(str string) int {

	byteNum := 0
	for _, s := range str {

		if len(string(s)) > 1 {
			byteNum += 2
		} else {
			byteNum += 1
		}
	}

	return byteNum
}

func String2MapBySep(str, sepOne string, sepTwo string) map[string]string {

	datas := strings.Split(str, sepOne)

	if len(datas) <= 0 || "" == datas[0] {
		return map[string]string{}
	}

	result := make(map[string]string, len(datas))

	for _, data := range datas {
		kvs := strings.Split(data, sepTwo)
		if len(kvs) <= 0 || "" == kvs[0] {
			continue
		}

		result[kvs[0]] = kvs[1]

	}

	return result
}

func String2IntMapBySep(str, sepOne string, sepTwo string) map[int]int {
	datas := strings.Split(str, sepOne)

	if len(datas) <= 0 || "" == datas[0] {
		return nil
	}

	result := make(map[int]int, len(datas))

	for _, data := range datas {
		kvs := strings.Split(data, sepTwo)
		if len(kvs) <= 0 || "" == kvs[0] {
			continue
		}

		kInt, _ := strconv.Atoi(kvs[0])
		vInt, _ := strconv.Atoi(kvs[1])
		result[kInt] = vInt
	}

	return result
}

func MapString2MapInt(ma map[string]string) map[int]int {

	maInt := map[int]int{}
	for k, v := range ma {

		kInt, _ := strconv.Atoi(k)
		vInt, _ := strconv.Atoi(v)

		maInt[kInt] = vInt
	}

	return maInt
}

func GetStack() (res string) {
	var (
		buf      = make([]byte, 1<<11)
		pc       uintptr
		file     string
		line     int
		funcName string
		data     = make(map[string]string)
		ok       bool
		b        []byte
		err      error
	)
	pc, file, line, ok = runtime.Caller(1)
	if !ok {
		return fmt.Sprint("runtime.Caller(1) error")
	}
	funcName = runtime.FuncForPC(pc).Name()
	runtime.Stack(buf, false)
	data["file"] = file
	data["line"] = strconv.Itoa(line)
	data["func"] = funcName
	data["stack"] = string(buf)

	b, err = json.Marshal(data)

	if err != nil {
		return fmt.Sprintf("%v", res)
	}
	return string(b)
}

func GetJson(i interface{}) string {
	b, err := json.Marshal(i)
	if err != nil {
		return fmt.Sprintf("%v", i)
	}
	return string(b)
}

func GenKey(k1, k2 int) (key uint64) {
	key |= uint64(uint32(k1))
	key |= uint64(uint64(k2) << 32)
	return key
}

func GetCallerInfo(dep int) string {
	pc, file, line, ok := runtime.Caller(dep)
	if !ok {
		return ""
	}
	funcName := runtime.FuncForPC(pc).Name()

	return fmt.Sprintf("%s in file %s, line:%d", funcName, file, line)
}

// 生成指定位数的随机数
func GenRandDigits(digit int) int {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	return int((rnd.Float32()*9 + 1) * float32(math.Pow10(digit-1)))
}

//-------------------------------------------------------------------------------------------------
func Interface2Map(src interface{}) map[string]interface{} {
	res := map[string]interface{}{}

	srcT := reflect.TypeOf(src)
	srcE := srcT.Elem()
	srcVE := reflect.ValueOf(src).Elem()

	for i := 0; i < srcE.NumField(); i++ {
		var v interface{}
		f := srcVE.Field(i)

		if !f.CanInterface() {
			continue
		}

		switch f.Type() {
		case reflect.TypeOf(time.Time{}):
			v = f.Interface().(time.Time).Unix()
		default:
			v = f.Interface()
		}
		if v == nil {
			continue
		}
		res[srcE.Field(i).Tag.Get("json")] = v
	}
	return res
}
