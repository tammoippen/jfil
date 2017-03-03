package lib

import (
	"os"
	"path/filepath"
	"time"
)

type file struct {
	Name string                   `json:"name"`
	Size int64                    `json:"size"`
	Type string                   `json:"type"`
	Path string                   `json:"path"`
	ModTime time.Time             `json:"modified"`
	Mode string                   `json:"mode"`
	Permissions file_permissions  `json:"permissions"`
	Sticky bool                   `json:"sticky"`
	Setuid bool                   `json:"setuid"`
	Setgid bool                   `json:"setgid"`
	Target *string                `json:"target"`
}

type file_permissions struct {
	Owner permissions  `json:"owner"`
	Group permissions  `json:"group"`
	Others permissions `json:"others"`
}

type permissions struct {
	Read bool     `json:"read"`
	Write bool    `json:"write"`
	Execute bool  `json:"execute"`
}

func (self file) Abs() string {
	res := self.Path + string(os.PathSeparator) + self.Name
	if self.Type == "dir" {
		res += string(os.PathSeparator)
	}
	return res
}

func NewFile(f_name string) *file {
	f_name_full, _ := filepath.Abs(f_name)
	f_info, _ := os.Lstat(f_name_full)
	f := new(file)

	switch mode := f_info.Mode(); {
	case mode.IsRegular():
		f.Type = "file"
	case mode.IsDir():
		f.Type = "dir"
	case mode&os.ModeSymlink != 0:
		f.Type = "symbolic link"
	case mode&os.ModeNamedPipe != 0:
		f.Type = "named pipe"
	case mode&os.ModeDevice != 0:
		f.Type = "device file"
	case mode&os.ModeDevice != 0 && mode & os.ModeCharDevice != 0:
		f.Type = "char device file"
	case mode&os.ModeSocket != 0:
		f.Type = "socket"
	}

	f.Name = f_info.Name()
	f.Path = filepath.Dir(f_name_full)
	f.Size = f_info.Size()

	f.Mode = f_info.Mode().String()
	f.ModTime = f_info.ModTime()
	f.Permissions = new_file_permissions(f_info.Mode())
	f.Sticky = f_info.Mode() & os.ModeSticky != 0
	f.Setuid = f_info.Mode() & os.ModeSetuid != 0
	f.Setgid = f_info.Mode() & os.ModeSetgid != 0

	if f.Type == "symbolic link" {
		tgt, _ := os.Readlink(f.Abs())
		f.Target = &tgt
	} else {
		f.Target = nil
	}

	return f
}

func new_file_permissions(mode os.FileMode) file_permissions {
	var p file_permissions
	p.Owner = new_permissions(uint32(mode >> 6) & 7)
	p.Group = new_permissions(uint32(mode >> 3) & 7)
	p.Others = new_permissions(uint32(mode >> 6) & 7)
	return p
}

func new_permissions(x uint32) permissions {
	var p permissions
	p.Read = x & 4 != 0
	p.Write = x & 2 != 0
	p.Execute = x & 1 != 0
	return p
}
