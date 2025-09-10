package fs

type LiveProjections struct {
	Symlinks map[string]string `json:"symlinks"` // target -> source
}

// Add a symlink
func (lp *LiveProjections) Project(source, target string) error {
	// os.Symlink(source, target)
	lp.Symlinks[target] = source
	return nil
}

// Remove a symlink
func (lp *LiveProjections) Unproject(target string) error {
	// os.Remove(target)
	delete(lp.Symlinks, target)
	return nil
}
