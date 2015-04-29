GOC := /usr/bin/go build
FETCHLIBS=/usr/bin/go get

MAKEMAN=/usr/bin/md2man-roff

BUILDDIR=$(CURDIR)/build
SRCDIR=src/kurz
MANSRCDIR=src/man
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
CONFDIR?=$(DESTDIR)/etc/kurz
MANPAGEDIR?=$(DESTDIR)/usr/share/man

all: kurz manpages

kurz: Makefile src/kurz/main.go
	mkdir -p $(GOPATHDIR) && \
	mkdir -p $(GOBINDIR) && \
	export GOPATH=$(GOPATHDIR) && \
	export GOBIN=$(GOBINDIR) && \
	cd $(SRCDIR) && \
	$(FETCHLIBS) && \
	$(GOC)

manpages:
	cd $(MANSRCDIR) && \
	$(MAKEMAN) kurz.1.md > kurz.1

clean:
	rm -fr build/
	rm -f src/kurz/kurz
	rm -f src/man/kurz.1

install:
	mkdir -p $(BINDIR)
	mkdir -p $(LIBDIR)
	mkdir -p $(CONFDIR)
	mkdir -p $(SYSTEMDCONFDIR)
	mkdir -p $(MANPAGEDIR)/man1
	$(INSTALL_BIN) src/kurz/kurz $(BINDIR)/
	$(INSTALL_BIN) src/kurz/kurz_echo.sh $(BINDIR)/
	$(INSTALL_LIB) src/lib/socket_send.sh $(LIBDIR)/
	$(INSTALL_LIB) src/lib/socket_send.js $(LIBDIR)/
	$(INSTALL_LIB) src/man/kurz.1 $(MANPAGEDIR)/man1/
	$(INSTALL_CONF) src/kurz/default.json $(CONFDIR)/
	$(INSTALL_CONF) src/init/kurz.target $(SYSTEMDCONFDIR)/
