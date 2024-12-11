package harperdb

import (
	"fmt"
	"testing"
)

const functionContent = "'use strict';\n\nconst https = require('https');\n\nconst authRequest = (options) => {\n  return new Promise((resolve, reject) => {\n    const req = https.request(options, (res) => {\n      res.setEncoding('utf8');\n      let responseBody = '';\n\n      res.on('data', (chunk) => {\n        responseBody += chunk;\n      });\n\n      res.on('end', () => {\n        resolve(JSON.parse(responseBody));\n      });\n    });\n\n    req.on('error', (err) => {\n      reject(err);\n    });\n\n    req.end();\n  });\n};\n\nconst customValidation = async (request,logger) => {\n  const options = {\n    hostname: 'jsonplaceholder.typicode.com',\n    port: 443,\n    path: '/todos/1',\n    method: 'GET',\n    headers: { authorization: request.headers.authorization },\n  };\n\n  const result = await authRequest(options);\n\n  /*\n   *  throw an authentication error based on the response body or statusCode\n   */\n  if (result.error) {\n    const errorString = result.error || 'Sorry, there was an error authenticating your request';\n    logger.error(errorString);\n    throw new Error(errorString);\n  }\n  return request;\n};\n\nmodule.exports = customValidation;\n"

func TestCustomFunctionsStatus(t *testing.T) {
	if _, err := c.CustomFunctionStatus(); err != nil {
		t.Fatal(err)
	}
}

func TestEndToEndCustomFunctions(t *testing.T) {
	project := randomID()
	if _, err := c.AddCustomFunctionProject(project); err != nil {
		t.Fatal(err)
	}

	if _, err := c.SetCustomFunction(project, "helpers", "example", functionContent); err != nil {
		t.Fatal(err)
	}

	if _, err := c.GetCustomFunctions(); err != nil {
		t.Fatal(err)
	}

	resp, err := c.GetCustomFunction(project, "helpers", "example")

	if err != nil {
		t.Fatal(err)
	}

	if resp.Message != functionContent {
		t.Fatal(fmt.Errorf("Expected contents to match defined const"))
	}

	if _, err := c.PackageCustomFunctionProject(project, false); err != nil {
		t.Fatal(err)
	}

	if _, err := c.DropCustomFunction(project, "helpers", "example"); err != nil {
		t.Fatal(err)
	}

	if _, err := c.DropCustomFunctionProject(project); err != nil {
		t.Fatal(err)
	}
}
