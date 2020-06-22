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
	"encoding/json"
	"github.com/bep/gitmap"
	"github.com/gothamhq/gotham/common/maps"
	"github.com/gothamhq/gotham/config"
	"github.com/gothamhq/gotham/hugofs/files"
	"github.com/gothamhq/gotham/langs"
	"github.com/gothamhq/gotham/media"
	"github.com/gothamhq/gotham/navigation"
	"github.com/gothamhq/gotham/source"
	"html/template"
	"time"
)

func MarshalPageToJSON(p Page) ([]byte, error) {
	content, err := p.Content()
	if err != nil {
		return nil, err
	}
	plain := p.Plain()
	plainWords := p.PlainWords()
	summary := p.Summary()
	truncated := p.Truncated()
	fuzzyWordCount := p.FuzzyWordCount()
	wordCount := p.WordCount()
	readingTime := p.ReadingTime()
	length := p.Len()
	tableOfContents := p.TableOfContents()
	rawContent := p.RawContent()
	resourceType := p.ResourceType()
	mediaType := p.MediaType()
	permalink := p.Permalink()
	relPermalink := p.RelPermalink()
	name := p.Name()
	title := p.Title()
	params := p.Params()
	data := p.Data()
	date := p.Date()
	lastmod := p.Lastmod()
	publishDate := p.PublishDate()
	expiryDate := p.ExpiryDate()
	aliases := p.Aliases()
	bundleType := p.BundleType()
	description := p.Description()
	draft := p.Draft()
	isHome := p.IsHome()
	keywords := p.Keywords()
	kind := p.Kind()
	layout := p.Layout()
	linkTitle := p.LinkTitle()
	isNode := p.IsNode()
	isPage := p.IsPage()
	path := p.Path()
	slug := p.Slug()
	lang := p.Lang()
	isSection := p.IsSection()
	section := p.Section()
	sectionsEntries := p.SectionsEntries()
	sectionsPath := p.SectionsPath()
	sitemap := p.Sitemap()
	typ := p.Type()
	weight := p.Weight()
	language := p.Language()
	file := p.File()
	gitInfo := p.GitInfo()
	outputFormats := p.OutputFormats()
	alternativeOutputFormats := p.AlternativeOutputFormats()
	menus := p.Menus()
	translationKey := p.TranslationKey()
	isTranslated := p.IsTranslated()
	allTranslations := p.AllTranslations()
	translations := p.Translations()

	s := struct {
		Content                  interface{}
		Plain                    string
		PlainWords               []string
		Summary                  template.HTML
		Truncated                bool
		FuzzyWordCount           int
		WordCount                int
		ReadingTime              int
		Len                      int
		TableOfContents          template.HTML
		RawContent               string
		ResourceType             string
		MediaType                media.Type
		Permalink                string
		RelPermalink             string
		Name                     string
		Title                    string
		Params                   maps.Params
		Data                     interface{}
		Date                     time.Time
		Lastmod                  time.Time
		PublishDate              time.Time
		ExpiryDate               time.Time
		Aliases                  []string
		BundleType               files.ContentClass
		Description              string
		Draft                    bool
		IsHome                   bool
		Keywords                 []string
		Kind                     string
		Layout                   string
		LinkTitle                string
		IsNode                   bool
		IsPage                   bool
		Path                     string
		Slug                     string
		Lang                     string
		IsSection                bool
		Section                  string
		SectionsEntries          []string
		SectionsPath             string
		Sitemap                  config.Sitemap
		Type                     string
		Weight                   int
		Language                 *langs.Language
		File                     source.File
		GitInfo                  *gitmap.GitInfo
		OutputFormats            OutputFormats
		AlternativeOutputFormats OutputFormats
		Menus                    navigation.PageMenus
		TranslationKey           string
		IsTranslated             bool
		AllTranslations          Pages
		Translations             Pages
	}{
		Content:                  content,
		Plain:                    plain,
		PlainWords:               plainWords,
		Summary:                  summary,
		Truncated:                truncated,
		FuzzyWordCount:           fuzzyWordCount,
		WordCount:                wordCount,
		ReadingTime:              readingTime,
		Len:                      length,
		TableOfContents:          tableOfContents,
		RawContent:               rawContent,
		ResourceType:             resourceType,
		MediaType:                mediaType,
		Permalink:                permalink,
		RelPermalink:             relPermalink,
		Name:                     name,
		Title:                    title,
		Params:                   params,
		Data:                     data,
		Date:                     date,
		Lastmod:                  lastmod,
		PublishDate:              publishDate,
		ExpiryDate:               expiryDate,
		Aliases:                  aliases,
		BundleType:               bundleType,
		Description:              description,
		Draft:                    draft,
		IsHome:                   isHome,
		Keywords:                 keywords,
		Kind:                     kind,
		Layout:                   layout,
		LinkTitle:                linkTitle,
		IsNode:                   isNode,
		IsPage:                   isPage,
		Path:                     path,
		Slug:                     slug,
		Lang:                     lang,
		IsSection:                isSection,
		Section:                  section,
		SectionsEntries:          sectionsEntries,
		SectionsPath:             sectionsPath,
		Sitemap:                  sitemap,
		Type:                     typ,
		Weight:                   weight,
		Language:                 language,
		File:                     file,
		GitInfo:                  gitInfo,
		OutputFormats:            outputFormats,
		AlternativeOutputFormats: alternativeOutputFormats,
		Menus:                    menus,
		TranslationKey:           translationKey,
		IsTranslated:             isTranslated,
		AllTranslations:          allTranslations,
		Translations:             translations,
	}

	return json.Marshal(&s)
}
