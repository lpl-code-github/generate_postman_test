package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

func main() {
	// 读取外部 JSON 文件（input）
	jsonFile, fileErr := ioutil.ReadFile("D:\\miaomiao\\go-project\\generate_postman_test\\src\\r1\\input\\input.json")
	if fileErr != nil {
		log.Fatal(fileErr)
	}

	var data interface{}
	if err := json.Unmarshal(jsonFile, &data); err != nil { // 解析 JSON 文件到数据结构中
		log.Fatal(err)
	}

	// 调用 generateSchema() 函数，生成 JSON Schema
	schema := generateSchema(data, nil)

	// 将 JSON Schema 对象格式化成 []byte 类型的 JSON 数据，并进行缩进
	schemaBytes, err := json.MarshalIndent(schema, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	// 创建并打开输出文件
	file, err := os.Create("D:\\miaomiao\\go-project\\generate_postman_test\\src\\r1\\output\\output.js")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// 写入 JavaScript 代码模板和 JSON Schema
	jsCode := getJsonSchemaTest(schemaBytes)
	// 将 JavaScript 代码模板和 JSON Schema 写入输出文件
	_, err = file.WriteString(jsCode)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("JavaScript code written to file!")
}

// 生成 JSON Schema 的函数
func generateSchema(data interface{}, parentFiled interface{}) interface{} {
	switch v := data.(type) {
	case map[string]interface{}:
		properties := make(map[string]interface{})
		var keys []string

		for k, vv := range v { // 遍历对象中的键值对
			keys = append(keys, k)                   // 记录键名到键名字符串数组中
			properties[k] = generateSchema(vv, data) // 递归解析当前键值并赋值给属性集合
		}

		return map[string]interface{}{ // 返回包含属性集合和所属类型、必需字段的对象
			"properties": properties,
			"type":       "object",
			"required":   keys,
		}
	case []interface{}:
		arr := data.([]interface{})
		if len(arr) == 0 { // 如果数组为空，则返回类型为数组的对象
			return map[string]interface{}{
				"type": "array",
			}
		}

		// 这里为了防止数组套数组，对上一级的类型进行判断,如果上一级是也是数组，item就返回空
		switch parentFiled.(type) { // 判断父级对象的类型
		case []interface{}: // 如果父级对象是数组类型，则返回空的 json object
			return map[string]interface{}{}
		default: // 否则返回包含数组项类型和所属数组类型的对象
			return map[string]interface{}{
				"items": generateSchema(arr[0], data),
				"type":  "array",
			}
		}

	case float64:
		return map[string]interface{}{
			"type": "integer",
		}
	case bool:
		return map[string]interface{}{
			"type": "boolean",
		}
	case string:
		s := fmt.Sprintf("%v", data) // 定义临时字符串变量并将数据转换为字符串类型
		// 有些时候 postman 会把数值型的字符串认为是 integer，这里多加一手判断
		if isStringInt(s) {
			return map[string]interface{}{
				"type": []interface{}{"string", "integer"},
			}
		}
		return map[string]interface{}{
			"type": "string",
		}
	default:
		return map[string]interface{}{
			"type": "null",
		}
	}
}

// 判断字符串是否可以成功解析为整型
func isStringInt(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

// 返回JsonSchemaTest测试模板
func getJsonSchemaTest(schemaBytes []byte) string {
	// 写入 JavaScript 代码模板和 JSON Schema
	jsTemplate := `const schema = %s;

pm.test('Schema is valid', function() {
	var jsonData = pm.response.json();
	pm.expect(tv4.validate(jsonData, schema)).to.be.true;
});
`

	jsCode := fmt.Sprintf(jsTemplate, string(schemaBytes)) // 把json格式的schema写入js模版
	return jsCode
}
