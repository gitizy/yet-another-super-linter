# Copyright (C) 2021 VERDOÏA Laurent
#
# This program is free software: you can redistribute it and/or modify
# it under the terms of the GNU General Public License as published by
# the Free Software Foundation, either version 3 of the License, or
# (at your option) any later version.
#
# This program is distributed in the hope that it will be useful,
# but WITHOUT ANY WARRANTY; without even the implied warranty of
# MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
# GNU General Public License for more details.
#
# You should have received a copy of the GNU General Public License
# along with this program.  If not, see <https://www.gnu.org/licenses/>.

FROM golangci/golangci-lint:v1.40.1 as golangci-lint
FROM hadolint/hadolint:v2.4.1-alpine as hadolint
FROM koalaman/shellcheck:v0.7.2 as shellcheck
FROM mvdan/shfmt:v3.2.4 as shfmt

FROM golang:1.15.12-alpine3.13 AS yasl

COPY . /opt/yasl

WORKDIR /opt/yasl

RUN go build .

FROM alpine:3.13.5

RUN set -eux; \
    apk add --no-cache \
        git=2.30.2-r0 \
        go=1.15.12-r0 \
        npm=14.16.1-r1 \
    ; \
    npm install \
        --global \
        --no-audit \
        --no-fund \
        --no-progress \
        --no-update-notifier \
        dockerfilelint@1.8.0 \
        jscpd@3.3.25 \
        markdownlint-cli@0.27.1

COPY --from=golangci-lint /usr/bin/golangci-lint /usr/bin/
COPY --from=hadolint /bin/hadolint /usr/bin/hadolint
COPY --from=shellcheck /bin/shellcheck /usr/bin/
COPY --from=shfmt /bin/shfmt /usr/bin/

COPY --from=yasl /opt/yasl/yasl /usr/local/bin/
COPY docker/config /opt/yasl/
COPY docker/bin /usr/local/bin/

CMD ["yasl", "lint"]
