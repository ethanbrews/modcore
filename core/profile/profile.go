package profile

import "modcore/lockfile"

type GameProfile struct {
	Name      string                    `json:"name"`
	Versions  []*lockfile.LockFile      `json:"versions"` // list of versions
	Default   string                    `json:"default"`  // version to use by default ("latest")
	LiveProjs *lockfile.LiveProjections `json:"live_projections"`
}
