package src

import "github.com/mitchellh/go-ps"

type DiffResult interface {
	Added() []ps.Process
	Removed() []ps.Process
}

type DiffResultOption struct {
	Prev []ps.Process
	Next []ps.Process
}

func processMatch(a ps.Process, b ps.Process) bool {
	return (a.Pid() == b.Pid()) && (a.Executable() == b.Executable())
}

func any(vs []ps.Process, f func(process ps.Process) bool) bool {
	for _, v := range vs {
		if f(v) {
			return true
		}
	}
	return false
}

func filter(vs []ps.Process, f func(process ps.Process) bool) []ps.Process {
	vsf := make([]ps.Process, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

func (o DiffResultOption) Added() []ps.Process {
	return filter(o.Next, func(newProcess ps.Process) bool {
		return !any(o.Prev, func(oldProcess ps.Process) bool {
			return processMatch(newProcess, oldProcess)
		})
	})
}

func (o DiffResultOption) Removed() []ps.Process {
	return filter(o.Prev, func(newProcess ps.Process) bool {
		return !any(o.Next, func(oldProcess ps.Process) bool {
			return processMatch(newProcess, oldProcess)
		})
	})
}

func GetChanges(prev []ps.Process, next []ps.Process) DiffResult {
	return DiffResultOption{
		Prev: prev,
		Next: next,
	}
}
