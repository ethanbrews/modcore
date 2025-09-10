package plugin

import "modcore/core/lockfile"

// Plugin represents a compiled-in module that can be imported and upgraded.
type Plugin interface {
	// ImportURL Import either from a URL or a file path
	ImportURL(url string) error
	ImportFile(path string) error

	// Upgrade the plugin to match the given lockfile state
	Upgrade(lf *lockfile.LockFile) error

	// Downgrade the plugin to match the given lockfile state
	Downgrade(lf *lockfile.LockFile) error
}
