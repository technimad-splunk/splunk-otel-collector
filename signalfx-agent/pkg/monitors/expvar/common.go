// Copyright  Splunk, Inc.
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

package expvar

import (
	"errors"
	"regexp"
	"strings"
)

var capRegexp = regexp.MustCompile("(^[^A-Z]*|[A-Z]*)([A-Z][^A-Z]+|$)")

func snakeCaseSlice(slice []string) []string {
	var words []string
	for _, s := range slice {
		var capWords []string
		for _, matchedCapWords := range capRegexp.FindAllStringSubmatch(s, -1) {
			for _, matchedCapWord := range matchedCapWords[1:] {
				if matchedCapWord = strings.TrimSpace(matchedCapWord); matchedCapWord != "" {
					capWords = append(capWords, matchedCapWord)
				}
			}
		}
		words = append(words, strings.ToLower(strings.Join(capWords, "_")))
	}
	return words
}

var wordRegexp = regexp.MustCompile("[a-zA-Z]+")

func joinWords(slice []string, sep string) string {
	var words []string
	for _, s := range slice {
		if wordRegexp.MatchString(s) {
			words = append(words, s)
		}
	}
	return strings.Join(words, sep)
}

var slashLastRegexp = regexp.MustCompile(`[^\/]*$`)

func getApplicationName(values map[string]interface{}) (string, error) {
	if cmdline, ok := values["cmdline"].([]interface{}); ok && len(cmdline) > 0 {
		name, ok := cmdline[0].(string)
		if !ok {
			return "", errors.New("unable to obtain app name")
		}
		if applicationName := strings.TrimSpace(slashLastRegexp.FindStringSubmatch(name)[0]); applicationName != "" {
			return applicationName, nil
		}
	}
	return "", errors.New("cmdline map not defined")
}
