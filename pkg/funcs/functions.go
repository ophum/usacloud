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

package funcs

// TODO v1へのserverコマンド移行までコメントアウト
//func getSSHPrivateKeyStorePath(serverID types.ID) (string, error) {
//	homeDir, err := homedir.Dir()
//	if err != nil {
//		return "", fmt.Errorf("getting HomeDir is failed:%s", err)
//	}
//	return filepath.Join(homeDir, ".ssh", fmt.Sprintf("sacloud_pkey_%d", serverID)), nil
//}