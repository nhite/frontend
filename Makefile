BUCKET = app.nhite.com
SOURCES := $(shell find $(SOURCEDIR) -name '*.go')
APP = nhiteApp/www/js/app.js

VERSION = 0.0.1
BUILD_TIME = `date +%FT%T%z`

all: $(APP)

$(APP): $(SOURCES)
	go generate ./...

deploy: $(APP)
	s3cmd sync --guess-mime-type --delete-removed nhiteApp/www/  s3://$(BUCKET)

test: $(APP)
	cd nhiteApp ; cordova run browser
