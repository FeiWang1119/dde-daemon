// Code generated by "dbusutil-gen em -type DisplayCfg"; DO NOT EDIT.

package displaycfg

import (
	"pkg.deepin.io/lib/dbusutil"
)

func (v *DisplayCfg) GetExportedMethods() dbusutil.ExportedMethods {
	return dbusutil.ExportedMethods{
		{
			Name:    "Get",
			Fn:      v.Get,
			OutArgs: []string{"cfgStr"},
		},
		{
			Name:   "Set",
			Fn:     v.Set,
			InArgs: []string{"cfgStr"},
		},
	}
}
