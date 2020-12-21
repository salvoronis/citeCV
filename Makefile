GO = go
SRCDIR = src
OUTDIR = build
OUTPORT = 8080
DOCPORT = 8080
OUTNAME = siteCV

all:
	$(GO) run $(SRCDIR)/*

build:
	$(GO) build -o $(OUTDIR)/$(OUTNAME) $(SRCDIR)/*

docker:
	docker run -p $(DOCPORT):$(OUTPORT) --name $(OUTNAME) --rm test

clean:
	rm $(OUTDIR)/*
