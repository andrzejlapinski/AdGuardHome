package aghtest_test

import (
	"github.com/andrzejlapinski/AdGuardHome/internal/aghtest"
	"github.com/andrzejlapinski/AdGuardHome/internal/client"
	"github.com/andrzejlapinski/AdGuardHome/internal/dnsforward"
	"github.com/andrzejlapinski/AdGuardHome/internal/filtering"
)

// Put interface checks that cause import cycles here.

// type check
var _ filtering.Resolver = (*aghtest.Resolver)(nil)

// type check
var _ dnsforward.ClientsContainer = (*aghtest.ClientsContainer)(nil)

// type check
//
// TODO(s.chzhen):  It's here to avoid the import cycle.  Remove it.
var _ client.AddressProcessor = (*aghtest.AddressProcessor)(nil)

// type check
//
// TODO(s.chzhen):  It's here to avoid the import cycle.  Remove it.
var _ client.AddressUpdater = (*aghtest.AddressUpdater)(nil)
