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
	// 读取外部 JSON 文件
	jsonFile, fileErr := ioutil.ReadFile("/Users/lpl/code/go-project/generate_postman_test/src/r1/input/input.json")
	if fileErr != nil {
		log.Fatal(fileErr)
	}

	var data interface{}
	if err := json.Unmarshal(jsonFile, &data); err != nil {
		log.Fatal(err)
	}

	// 调用 generateSchema() 函数，生成 JSON Schema
	schema := generateSchema(data, nil)

	// 将 JSON Schema 转换为字符串输出
	schemaBytes, err := json.MarshalIndent(schema, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Create("/Users/lpl/code/go-project/generate_postman_test/src/r1/output/output.js")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// 写入 JavaScript 代码模板和 JSON Schema
	jsTemplate := `var schema = %s;

pm.test('Schema is valid', function() {
	var jsonData = pm.response.json();
	pm.expect(tv4.validate(jsonData, schema)).to.be.true;
});
`

	jsCode := fmt.Sprintf(jsTemplate, string(schemaBytes))

	_, err = file.WriteString(jsCode)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("JavaScript code written to file!")

	//fmt.Println(string(schemaBytes))
}

func generateSchema(data interface{}, parentFiled interface{}) interface{} {
	switch v := data.(type) {
	case map[string]interface{}:
		properties := make(map[string]interface{})
		var keys []string

		for k, vv := range v {
			keys = append(keys, k)
			properties[k] = generateSchema(vv, data)
		}

		return map[string]interface{}{
			"properties": properties,
			"type":       "object",
			"required":   keys,
		}
	case []interface{}:

		arr := data.([]interface{})
		if len(arr) == 0 {
			return map[string]interface{}{
				"type": "array",
			}
		}
		// 这里为了防止数组套数组，对上一级的类型进行判断
		switch parentFiled.(type) {
		case []interface{}:
			return map[string]interface{}{}
		default:

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
		s := fmt.Sprintf("%v", data)
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

func isStringInt(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}
