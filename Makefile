GOC := /usr/bin/go build
FETCHLIBS=/usr/bin/go get

BUILDDIR=$(CURDIR)/build
SRCDIR=src/kurz
EXECDIR=$(BUILDDIR)/bin
LIBDIR=$(BUILDDIR)/golibs

INSTALL=install
INSTALL_BIN=$(INSTALL) -m755
INSTALL_CONF=$(INSTALL) -m400

PREFIX?=$(DESTDIR)/usr
BINDIR?=$(PREFIX)/bin
CONFDIR?=$(DESTDIR)/etc/kurz/

all: kurz

kurz: Makefile src/kurz/main.go
	mkdir -p $(EXECDIR) && \
	mkdir -p $(LIBDIR) && \
	export GOPATH=$(LIBDIR) && \
	export GOBIN=$(EXECDIR) && \
	cd $(SRCDIR) && \
	$(FETCHLIBS) && \
	$(GOC)

clean:
	rm -fr build/
	rm -f src/kurz/kurz

install:
	mkdir -p $(BINDIR)
	mkdir -p $(CONFDIR)
	$(INSTALL_BIN) src/kurz/kurz $(BINDIR)/
	$(INSTALL_BIN) src/kurz/kurz_echo.sh $(BINDIR)/
	$(INSTALL_CONF) src/kurz/default.json $(CONFDIR)/
