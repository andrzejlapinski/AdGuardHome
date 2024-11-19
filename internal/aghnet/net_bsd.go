//go:build darwin || freebsd || openbsd

package aghnet

import "github.com/andrzejlapinski/AdGuardHome/internal/aghos"

func canBindPrivilegedPorts() (can bool, err error) {
	return aghos.HaveAdminRights()
}
