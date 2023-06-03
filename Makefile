.PHONY all install test build

all: install

install:
	./scripts/install.sh

test:
	./scripts/test.sh

build:
	./scripts/build.sh

publish:
	./scripts/publish.sh