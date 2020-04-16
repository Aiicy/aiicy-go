//
// Copyright (c) 2019-present Codist <countstarlight@gmail.com>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that can
// be found in the LICENSE file.
// Written by Codist <countstarlight@gmail.com>, December 2019
//

package logger

// LogConfig for logging
type Config struct {
	Path       string `yaml:"path" json:"path"`
	Level      string `yaml:"level" json:"level" default:"info" validate:"regexp=^(info|debug|warn|error)$"`
	Encoding   string `yaml:"encoding" json:"encoding" default:"json" validate:"regexp=^(json|console)$"`
	Format     string `yaml:"format" json:"format" default:"text" validate:"regexp=^(text|json)$"`
	MaxAge     int    `yaml:"maxAge" json:"maxAge" default:"15" validate:"min=1"`   // days
	MaxSize    int    `yaml:"maxSize" json:"maxSize" default:"50" validate:"min=1"` // MB
	MaxBackups int    `yaml:"maxBackups" json:"maxBackups" default:"15" validate:"min=1"`
}
