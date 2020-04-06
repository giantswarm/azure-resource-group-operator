package project

import (
	"github.com/giantswarm/versionbundle"
)

func NewVersionBundle() versionbundle.Bundle {
	return versionbundle.Bundle{
		Changelogs: []versionbundle.Changelog{
			{
				Component:   "azure-resource-group-operator",
				Description: "TODO",
				Kind:        versionbundle.KindChanged,
			},
		},
		Components: []versionbundle.Component{},
		Name:       "azure-resource-group-operator",
		Version:    BundleVersion(),
	}
}
