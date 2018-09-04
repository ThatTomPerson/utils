GO = go
PACKAGES = $(shell $(GO) list ./...)
PACKAGE = github.com/ThatTomPerson/utils

all: $(PACKAGES);

$(PACKAGES): PKG=$(subst $(PACKAGE)/,,$@)
$(PACKAGES): URL=$(subst /$(PKG),,$@)
$(PACKAGES):
	godoc2md $@ > ./$(PKG)/README.md
.PHONY: $(PACKAGES)