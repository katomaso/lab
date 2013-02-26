package gosrc

import (
	"github.com/mb0/lab/ws"
)

type Import struct {
	Path string
	Id   ws.Id
}
type File struct {
	Id   ws.Id
	Name string
	Err  error
}

type Info struct {
	Files   []File
	Imports []Import
	Uses    []ws.Id
}

func (nfo *Info) Import(path string) *Import {
	for i := range nfo.Imports {
		imprt := &nfo.Imports[i]
		if path == imprt.Path {
			return imprt
		}
	}
	return nil
}
func (nfo *Info) AddImport(path string) {
	if nfo.Import(path) == nil {
		nfo.Imports = append(nfo.Imports, Import{Path: path})
	}
}
func (nfo *Info) File(id ws.Id) *File {
	for i := range nfo.Files {
		file := &nfo.Files[i]
		if id == file.Id {
			return file
		}
	}
	return nil
}
func (nfo *Info) AddFile(id ws.Id, name string) {
	if nfo.File(id) == nil {
		nfo.Files = append(nfo.Files, File{Id: id, Name: name})
	}
}
func (nfo *Info) AddUse(id ws.Id) {
	for _, i := range nfo.Uses {
		if i == id {
			return
		}
	}
	nfo.Uses = append(nfo.Uses, id)
}
func (nfo *Info) Merge(old *Info) {
	if nfo == nil || old == nil {
		return
	}
	for i := range nfo.Imports {
		imprt := &nfo.Imports[i]
		if imprt.Id == 0 {
			with := old.Import(imprt.Path)
			if with != nil {
				imprt.Id = with.Id
			}
		}
	}
	nfo.Uses = old.Uses
	return
}