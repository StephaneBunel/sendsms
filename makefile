PROJ = github.com/StephaneBunel/sendsms
EXE = $(notdir $(PROJ))
SRC = cmd/**/main.go \
      pkg/**/*.go

LDFLAGS = -s
GOBUILD = go build -v --ldflags="$(LDFLAGS)"

all: $(EXE)
$(EXE): $(SRC)
	$(GOBUILD) -v $(PROJ)/cmd/sendsms

vendor:
	govendor sync -v

clean:
	@rm -f $(EXE)
