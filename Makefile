GOC := /usr/bin/go build
FETCHLIBS=/usr/bin/go get

BUILDDIR=$(CURDIR)/build
SRCDIR=src/kurz
GOBINDIR=$(BUILDDIR)/bin
GOPATHDIR=$(BUILDDIR)/golibs

INSTALL=install
INSTALL_BIN=$(INSTALL) -m755
INSTALL_LIB=$(INSTALL) -m644
INSTALL_CONF=$(INSTALL) -m400

PREFIX?=$(DESTDIR)/usr
SYSTEMDCONFDIR?=$(DESTDIR)/lib/systemd/system
BINDIR?=$(PREFIX)/bin
LIBDIR?=$(PREFIX)/lib/kurz
CONFDIR?=$(DESTDIR)/etc/kurz/


all: kurz

kurz: Makefile src/kurz/main.go
	mkdir -p $(GOPATHDIR) && \
	mkdir -p $(GOBINDIR) && \
	export GOPATH=$(GOPATHDIR) && \
	export GOBIN=$(GOBINDIR) && \
	cd $(SRCDIR) && \
	$(FETCHLIBS) && \
	$(GOC)

clean:
	rm -fr build/
	rm -f src/kurz/kurz

install:
	mkdir -p $(BINDIR)
	mkdir -p $(LIBDIR)
	mkdir -p $(CONFDIR)
	mkdir -p $(SYSTEMDCONFDIR)
	$(INSTALL_BIN) src/kurz/kurz $(BINDIR)/
	$(INSTALL_BIN) src/kurz/kurz_echo.sh $(BINDIR)/
	$(INSTALL_LIB) src/lib/socket_send.sh $(LIBDIR)/
	$(INSTALL_LIB) src/lib/socket_send.js $(LIBDIR)/
	$(INSTALL_CONF) src/kurz/default.json $(CONFDIR)/
	$(INSTALL_CONF) src/init/kurz.target $(SYSTEMDCONFDIR)/
