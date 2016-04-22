package vsolver

import (
	"os"
	"path"
)

type Result interface {
	Lock
	Attempts() int
}

type result struct {
	// A list of the projects selected by the solver.
	p []LockedProject

	// The number of solutions that were attempted
	att int

	// The hash digest of the input opts
	hd []byte
}

func CreateVendorTree(basedir string, l Lock, sm SourceManager) error {
	err := os.MkdirAll(basedir, 0777)
	if err != nil {
		return err
	}

	// TODO parallelize
	for _, p := range l.Projects() {
		to := path.Join(basedir, string(p.n))
		os.MkdirAll(to, 0777)
		err := sm.ExportAtomTo(p.toAtom(), to)
		if err != nil {
			os.RemoveAll(basedir)
			return err
		}
		// TODO dump version metadata file
	}

	return nil
}

func (r result) Projects() []LockedProject {
	return r.p
}

func (r result) Attempts() int {
	return r.att
}

func (r result) InputHash() []byte {
	return r.hd
}