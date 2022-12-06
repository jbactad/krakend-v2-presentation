FROM devopsfaith/krakend-plugin-builder:2.1.2 as plugin-builder

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY plugins ./plugins
COPY Makefile .

RUN make build-plugins


FROM devopsfaith/krakend:2.1.2

COPY --from=plugin-builder /src/bin/ /opt/krakend/plugins/





