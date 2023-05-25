# colorswap
# See LICENSE for copyright and license details.
GO ?= go
RM ?= rm
INSTALL ?= install
SCDOC ?= scdoc
GOFLAGS ?=
PREFIX ?= /usr/local
BINDIR ?= bin

all: colorswap

colorswap:
	$(GO) build $(GOFLAGS)

clean:
	$(RM) -f colorswap

install:
	$(INSTALL) -pm 0755 colorswap -t $(DESTDIR)$(PREFIX)/$(BINDIR)/

uninstall:
	$(RM) -f $(DESTDIR)$(PREFIX)/$(BINDIR)/colorswap

.PHONY: all colorswap clean install uninstall
