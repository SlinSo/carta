# go option
GO        ?= go
TAGS      :=
TESTS     := .
TESTFLAGS :=
LDFLAGS   := #-w -s
GOFLAGS   :=
BINDIR    := $(CURDIR)/bin

# tools
CP := cp -u -v

# Required for globs to work correctly
SHELL=/usr/bin/env sh

# github.com/jessfraz/junk/sembump download to gopath externally
.PHONY: bump
BUMP := patch
bump:
	@echo "(${BUMP})ing"

	$(eval VERSION_FILE := version.txt)
	$(eval VERSION := $(shell cat ${VERSION_FILE}))
	$(eval NEW_VERSION = $(shell sembump --kind $(BUMP) $(VERSION)))

	@echo "Bumping version.txt from $(VERSION) to $(NEW_VERSION)"
	@echo $(NEW_VERSION) > ${VERSION_FILE}
	git add ${VERSION_FILE}
	git commit -vsam "bump 'carta' version to $(NEW_VERSION)"
	git tag -a $(NEW_VERSION) -m "$(NEW_VERSION)"
	git push origin $(NEW_VERSION)
