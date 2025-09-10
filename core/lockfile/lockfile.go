package lockfile

import (
	"time"
)

type ProjectionType string

const (
	ProjectionPermanent ProjectionType = "permanent"
	ProjectionTransient ProjectionType = "transient"
)

type Projection struct {
	SourcePath string         `json:"source"` // original plugin location
	TargetPath string         `json:"target"` // where the symlink is projected
	Type       ProjectionType `json:"type"`   // permanent vs transient
}

type LockFile struct {
	Version       string                `json:"version"`
	CreatedAt     time.Time             `json:"created_at"`
	Plugins       map[string]string     `json:"plugins"`        // plugin name -> hash of contents
	Projections   map[string]Projection `json:"projections"`    // symlink map
	UpgradeURL    string                `json:"upgrade_url"`    // optional for unattended upgrades
	TransientRoot string                `json:"transient_root"` // default directory for transient projections
	PermanentRoot string                `json:"permanent_root"` // optional root for permanent locations
}
