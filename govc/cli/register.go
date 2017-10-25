/*
Copyright (c) 2014-2015 VMware, Inc. All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package cli

var commands = map[string]Command{}

var aliases = map[string]string{}

func Register(name string, c Command) {
	commands[name] = c
}

func Unregister(name string) {
	delete(commands, name)
}

func Alias(name string, alias string) {
	aliases[alias] = name
}

func Commands() map[string]Command {
	return commands
}
