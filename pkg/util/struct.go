// Copyright 2017-2020 The Usacloud Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package util

import (
	"encoding/json"
	"errors"
	"io/ioutil"
)

func MarshalJSONFromPathOrContent(pathOrContent string, destination interface{}) error {
	// 今の所JSONのみ対応
	data, err := BytesFromPathOrContent(pathOrContent)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, destination)
}

func StringFromPathOrContent(pathOrContent string) (string, error) {
	data, err := BytesFromPathOrContent(pathOrContent)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func BytesFromPathOrContent(pathOrContent string) ([]byte, error) {
	if pathOrContent == "" {
		return nil, errors.New("pathOrContent required")
	}

	data, err := ioutil.ReadFile(pathOrContent)
	if err != nil {
		return []byte(pathOrContent), nil // ファイルを読んでみてエラーだった場合はJSONなどのコンテンツと判定する
	}
	return data, nil
}
