# Changelog

## Version v1.12.1 (2025-02-17)

### Chores and tidying

- **deps:** update module flamingo.me/flamingo/v3 to v3.13.0 (#141) (16fac4c2)
- **deps:** update module github.com/spf13/cobra to v1.9.1 (#144) (a25eb68d)
- **deps:** bump go version and adjust linter (#145) (6e8bd771)
- **deps:** update module github.com/vektah/gqlparser/v2 to v2.5.21 (#140) (045ce12f)
- **deps:** update module github.com/99designs/gqlgen to v0.17.60 (#136) (01e19a8e)
- **deps:** update dependency go to v1.23.4 (#134) (879aacf8)
- **deps:** update module flamingo.me/dingo to v0.3.0 (#138) (2fc3f168)
- **deps:** update module github.com/vektah/gqlparser/v2 to v2.5.20 (#135) (f9f5fa75)
- **deps:** update actions/checkout digest to 11bd719 (#133) (0c0dd5a9)
- **deps:** update actions/checkout digest to eef6144 (#131) (53d57a7f)
- **deps:** update dependency go to v1.23.2 (#129) (b28969e4)
- **deps:** update module github.com/99designs/gqlgen to v0.17.55 (#130) (1af1723b)
- **deps:** update module github.com/vektah/gqlparser/v2 to v2.5.18 (#132) (77c5226e)
- **deps:** update module github.com/vektah/gqlparser/v2 to v2.5.17 (#127) (497c1c62)
- **deps:** update module github.com/99designs/gqlgen to v0.17.54 (#126) (7b9c181f)
- **deps:** update module flamingo.me/flamingo/v3 to v3.11.0 (#125) (81fb4e73)

## Version v1.12.0 (2024-09-18)

### Features

- **#119:** mute type hints in GraphQL error messages (#123) (32e0eb3f)

### Chores and tidying

- **deps:** update module github.com/99designs/gqlgen to v0.17.53 (#124) (69c4f731)
- **deps:** update dependency go to v1.23.1 (#120) (a164bfd8)
- **deps:** update module flamingo.me/flamingo/v3 to v3.10.1 (#122) (6b9b55e6)
- **deps:** update module flamingo.me/flamingo/v3 to v3.10.0 (#121) (104a8bf4)
- **deps:** update dependency go to v1.22.6 (#118) (38f71f93)
- **deps:** update dependency go to v1.22.5 (#115) (be2c76a5)
- **deps:** update module flamingo.me/flamingo/v3 to v3.9.0 (#116) (9f32919b)

## Version v1.11.3 (2024-06-17)

### Chores and tidying

- **deps:** update actions/checkout digest to 692973e (#105) (704adc4d)
- **deps:** update dependency go to v1.22.4 (#111) (b3410fcb)
- **deps:** update module github.com/spf13/cobra to v1.8.1 (#114) (32224faa)
- **deps:** update module github.com/99designs/gqlgen to v0.17.49 (#109) (57c3cd2f)
- **deps:** update golangci/golangci-lint-action action to v6 (#108) (8d17a220)
- **deps:** update dependency go to v1.22.3 (#104) (945832b0)
- **deps:** update module flamingo.me/flamingo/v3 to v3.8.1 (#107) (e253d0c3)
- **deps:** bump golang.org/x/net from 0.22.0 to 0.23.0 (#103) (1163d35c)
- **deps:** update module github.com/99designs/gqlgen to v0.17.45 (#100) (24372995)
- **deps:** bump google.golang.org/protobuf from 1.32.0 to 1.33.0 (#101) (c47f7c03)

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

