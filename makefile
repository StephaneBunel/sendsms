PROJ = github.com/StephaneBunel/sendsms
TARGETS = sendsms
SRC = sms/*.go \
      provider/*.go provider/**/*.go \
      recipient/*.go recipient/**/*.go \
	  cmd/**/*.go

LDFLAGS = -s
GOBUILD = go build -v --ldflags="$(LDFLAGS)"

all: $(TARGETS)

sendsms: $(SRC)
	$(GOBUILD) $(PROJ)/cmd/sendsms

vendor:
	govendor sync -v

clean:
	@rm -f $(TARGETS)
