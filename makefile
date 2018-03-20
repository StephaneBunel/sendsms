EXE = sendsms
PROJ = github.com/StephaneBunel/sendsms
TARGETS = sendsms
SRC = domain/*.go service/*.go infra/*.go infra/**/*.go main.go

LDFLAGS = -s
GOBUILD = go build -v --ldflags="$(LDFLAGS)"

all: $(TARGETS)

$(EXE): $(SRC)
	$(GOBUILD) -v $(PROJ)

vendor:
	govendor sync -v

clean:
	@rm -f $(TARGETS) $(PROJ)/$(EXE)
