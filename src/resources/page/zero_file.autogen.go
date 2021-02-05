// Copyright 2019 The Hugo Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This file is autogenerated.

package page

import (
	"github.com/gothamhq/gotham/helpers"
	"github.com/gothamhq/gotham/source"
	"github.com/strawberryssg/strawberry-v0/hugofs"
)

// ZeroFile represents a zero value of source.File with warnings if invoked.
type zeroFile struct {
	log *helpers.DistinctLogger
}

func NewZeroFile(log *helpers.DistinctLogger) source.File {
	return zeroFile{log: log}
}

func (zeroFile) IsZero() bool {
	return true
}

func (z zeroFile) Path() (o0 string) {
	z.log.Println(".File.Path on zero object. Wrap it in if or with: {{ with .File }}{{ .Path }}{{ end }}")
	return
}
func (z zeroFile) Section() (o0 string) {
	z.log.Println(".File.Section on zero object. Wrap it in if or with: {{ with .File }}{{ .Section }}{{ end }}")
	return
}
func (z zeroFile) Lang() (o0 string) {
	z.log.Println(".File.Lang on zero object. Wrap it in if or with: {{ with .File }}{{ .Lang }}{{ end }}")
	return
}
func (z zeroFile) Filename() (o0 string) {
	z.log.Println(".File.Filename on zero object. Wrap it in if or with: {{ with .File }}{{ .Filename }}{{ end }}")
	return
}
func (z zeroFile) Dir() (o0 string) {
	z.log.Println(".File.Dir on zero object. Wrap it in if or with: {{ with .File }}{{ .Dir }}{{ end }}")
	return
}
func (z zeroFile) Extension() (o0 string) {
	z.log.Println(".File.Extension on zero object. Wrap it in if or with: {{ with .File }}{{ .Extension }}{{ end }}")
	return
}
func (z zeroFile) Ext() (o0 string) {
	z.log.Println(".File.Ext on zero object. Wrap it in if or with: {{ with .File }}{{ .Ext }}{{ end }}")
	return
}
func (z zeroFile) LogicalName() (o0 string) {
	z.log.Println(".File.LogicalName on zero object. Wrap it in if or with: {{ with .File }}{{ .LogicalName }}{{ end }}")
	return
}
func (z zeroFile) BaseFileName() (o0 string) {
	z.log.Println(".File.BaseFileName on zero object. Wrap it in if or with: {{ with .File }}{{ .BaseFileName }}{{ end }}")
	return
}
func (z zeroFile) TranslationBaseName() (o0 string) {
	z.log.Println(".File.TranslationBaseName on zero object. Wrap it in if or with: {{ with .File }}{{ .TranslationBaseName }}{{ end }}")
	return
}
func (z zeroFile) ContentBaseName() (o0 string) {
	z.log.Println(".File.ContentBaseName on zero object. Wrap it in if or with: {{ with .File }}{{ .ContentBaseName }}{{ end }}")
	return
}
func (z zeroFile) UniqueID() (o0 string) {
	z.log.Println(".File.UniqueID on zero object. Wrap it in if or with: {{ with .File }}{{ .UniqueID }}{{ end }}")
	return
}
func (z zeroFile) FileInfo() (o0 hugofs.FileMetaInfo) {
	z.log.Println(".File.FileInfo on zero object. Wrap it in if or with: {{ with .File }}{{ .FileInfo }}{{ end }}")
	return
}
