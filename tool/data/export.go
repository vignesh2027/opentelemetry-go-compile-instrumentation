// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0
package data

import (
	"bytes"
	_ "embed"
)

//go:embed otelc-bundle.tgz
var bundle []byte

//go:embed instrumentation-manifest.json
var manifestJSON []byte

// GetBundleReader returns a bytes.Reader for the embedded bundle
func GetBundleReader() *bytes.Reader {
	return bytes.NewReader(bundle)
}

// GetManifestJSON returns the embedded instrumentation manifest, the contents
// of instrumentation-manifest.json. The result is a copy, so a caller may hold
// on to it or modify it without disturbing the embedded bytes that every other
// caller shares.
func GetManifestJSON() []byte {
	return bytes.Clone(manifestJSON)
}
