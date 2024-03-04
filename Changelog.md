# Changelog

## Version v1.11.2 (2024-03-04)

### Chores and tidying

- **deps:** update module github.com/stretchr/testify to v1.9.0 (#97) (e4e7797c)
- **deps:** update module github.com/99designs/gqlgen to v0.17.44 (#96) (df423d3f)
- **deps:** update golangci/golangci-lint-action action to v4 (#95) (56e3efee)

## Version v1.11.1 (2024-02-08)

### Chores and tidying

- **deps:** update module flamingo.me/flamingo/v3 to v3.8.0 (#93) (51d46de7)

## Version v1.11.0 (2024-02-06)

### Features

- **security:** add middleware to limit top level operation amount (#90) (4cb35ae3)

### Chores and tidying

- **deps:** update module github.com/99designs/gqlgen to v0.17.43 (#91) (0c43c473)
- **deps:** update module github.com/vektah/gqlparser/v2 to v2.5.11 (#92) (ebeef0c3)
- **deps:** update actions/setup-go action to v5 (#88) (2a52972d)
- **deps:** update module github.com/spf13/cobra to v1.8.0 (#86) (ed0576b0)
- **deps:** update module github.com/99designs/gqlgen to v0.17.41 (#87) (b5cd0ce8)

## Version v1.10.1 (2023-11-01)

### Ops and CI/CD

- bump golangci-lint to v1.55 (#85) (2e17f972)
- Update go version in pipeline (#77) (cdd32194)

### Chores and tidying

- **deps:** update actions/checkout digest to b4ffde6 (#82) (31b63a02)
- **deps:** bump golang.org/x/net from 0.14.0 to 0.17.0 (#83) (10153fb2)
- **deps:** update module github.com/99designs/gqlgen to v0.17.40 (#80) (8380d021)
- fix linter issues (#84) (d12d3544)
- **deps:** update actions/checkout action to v4 (#79) (a2d450f6)
- **deps:** update module flamingo.me/flamingo/v3 to v3.7.0 (#76) (42085260)

## Version v1.10.0 (2023-07-28)

### Features

- Add basic support for OpenCensus tracing (resolvers only) (#70) (a2dabe9c)

### Chores and tidying

- **deps:** update module github.com/99designs/gqlgen to v0.17.36 (#75) (abfcb3f6)
- **deps:** update module github.com/vektah/gqlparser/v2 to v2.5.8 (#74) (3016084d)
- **deps:** update module github.com/99designs/gqlgen to v0.17.35 (#73) (31a7b82d)
- Disable introspection already in the server, drop obsolete middleware (#72) (dcbd7df3)

## Version v1.9.1 (2023-07-04)

### Fixes

- **275771:** Fix wishlist limit, update flamingo-om3 (4626b870)

### Ops and CI/CD

- upgrade golangci-lint (#62) (b5f68030)

### Chores and tidying

- **deps:** Update github.com/99designs/gqlgen and github.com/vektah/gqlparser/v2 (#69) (2f089dfe)
- **deps:** update module flamingo.me/flamingo/v3 to v3.6.1 (#67) (2493e209)
- **deps:** update module flamingo.me/flamingo/v3 to v3.6.0 (#65) (1537ccfc)
- **deps:** update module github.com/99designs/gqlgen to v0.17.29 (#64) (941dc812)
- **deps:** update module flamingo.me/flamingo/v3 to v3.5.1 (#59) (4e75dcc7)
- **deps:** update module github.com/99designs/gqlgen to v0.17.28 (#58) (2ff9a666)
- **deps:** update actions/setup-go action to v4 (#60) (1793fc81)
- **deps:** update module github.com/spf13/cobra to v1.7.0 (#63) (7fcbbe8e)
- add regex to detect go run/install commands (#61) (f9781620)
- **deps:** update module github.com/99designs/gqlgen to v0.17.24 (#55) (df2245ba)
- **deps:** update module flamingo.me/flamingo/v3 to v3.5.0 (#56) (744fe71b)
- **deps:** update module github.com/99designs/gqlgen to v0.17.22 (#52) (9678ac9e)

### Other

- Bump golang.org/x/net from 0.5.0 to 0.7.0 (#57) (94c33e00)

## Version v1.9.0 (2022-12-05)

### Features

- **graphql:** allow to register additional OperationMiddlewares (#51) (34d923a2)
- **graphql:** support directives (#33) (05fb2104)

### Fixes

- **command:** remove ioutil (#41) (a253cb9c)

### Ops and CI/CD

- fix golangci-lint behaviour (9986e7fd)
- add golangci-lint (f4f1dc45)

### Chores and tidying

- **deps:** update module github.com/99designs/gqlgen to v0.17.21 (#50) (8104cfdf)
- **deps:** update module flamingo.me/dingo to v0.2.10 (#48) (a6b9ed58)
- **deps:** update module flamingo.me/flamingo/v3 to v3.4.0 (#47) (a169a567)
- **deps:** update module github.com/spf13/cobra to v1.6.1 (#46) (97c0f19b)
- **deps:** update module github.com/99designs/gqlgen to v0.17.20 (#42) (c08c3533)
- bump to go 1.19 (54836837)
- **deps:** update module github.com/99designs/gqlgen to v0.17.15 (#40) (2b6063ac)
- **deps:** update module github.com/vektah/gqlparser/v2 to v2.4.8 (#39) (5056b96f)
- **deps:** update module github.com/vektah/gqlparser/v2 to v2.4.7 (#38) (45865161)
- **deps:** update module github.com/99designs/gqlgen to v0.17.13 (#36) (982ec0ce)
- **deps:** update module github.com/spf13/cobra to v1.5.0 (bf1bf6e7)
- **deps:** update module github.com/99designs/gqlgen to v0.17.12 (61f0127a)
- **deps:** update module github.com/vektah/gqlparser/v2 to v2.4.6 (200c5c63)
- **deps:** update module github.com/99designs/gqlgen to v0.17.10 (#30) (984798fb)
- **deps:** update module github.com/vektah/gqlparser/v2 to v2.4.5 (#31) (604c232b)
- **deps:** update module github.com/99designs/gqlgen to v0.17.6 (#28) (3826079a)
- **deps:** update module github.com/vektah/gqlparser/v2 to v2.4.3 (#29) (c69f8afb)
- **deps:** update actions/checkout action to v3 (#24) (de0e0a4b)
- **deps:** update actions/setup-go action to v3 (#25) (64f0d694)

### Other

- Configure Renovate (#23) (18ad8968)

## Version v1.8.0 (2022-05-06)

### Features

- **handler:** do not apply cookies after flamingo ran (fd28798a)

### Fixes

- **x/tools:** update to v0.1.10 (7f139a5d)

### Chores and tidying

- Update dependencies, add staticcheck to ci (#22) (8b4b4647)

## Version v1.7.0 (2022-04-01)

### Features

- add the ability to define the maximum number of bytes used to parse the request body as multipart/form-data (29f0237e)

### Ops and CI/CD

- **semanticore:** add semanticore (43bcdcac)

### Other

- run goimports (b9df1937)

