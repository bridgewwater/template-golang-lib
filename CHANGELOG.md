# Changelog

All notable changes to this project will be documented in this file. See [convention-change-log](https://github.com/convention-change/convention-change-log) for commit guidelines.

## [1.1.0](https://github.com/bridgewwater/template-golang-lib/compare/1.0.0...v1.1.0) (2024-04-28)

### 🐛 Bug Fixes

* update temp-golang-lib for fast add lib ([d0c21384](https://github.com/bridgewwater/template-golang-lib/commit/d0c21384e8c992b7352787c001ae085e5fd6690c))

### ✨ Features

* update full docker build pipeline of this project ([30cfaf5d](https://github.com/bridgewwater/template-golang-lib/commit/30cfaf5da7f1b8cbad43aa2ab2d1e33144efb6a0))

* update github.com/sinlov-go/unittest-kit v1.1.0 ([8f24849e](https://github.com/bridgewwater/template-golang-lib/commit/8f24849e54a897196f053ac5f4f25d962a9e4b39))

* use github.com/sinlov-go/unittest-kit ([21907873](https://github.com/bridgewwater/template-golang-lib/commit/21907873a97c0f54b5580c6c8736593671649aa0))

* change env to CI_DEBUG to change unit test case flag ([8a5f3961](https://github.com/bridgewwater/template-golang-lib/commit/8a5f3961c190aac778eff8379ed94b7b76a13b44))

### ♻ Refactor

* remove useless example_test/TestMain.go ([5493d8db](https://github.com/bridgewwater/template-golang-lib/commit/5493d8dbe57133ae8f14ac0cea1f4d31dcf6f29c))

* change example test case kit for fast use ([e5dbc431](https://github.com/bridgewwater/template-golang-lib/commit/e5dbc4315152f7279e5ba1a9a1763e85243a8d00))

* change example_test to different file ([f4d52850](https://github.com/bridgewwater/template-golang-lib/commit/f4d528509e8ff5c560bb0ee7247561a33af17988))

### 👷‍ Build System

* bump github.com/stretchr/testify from 1.8.4 to 1.9.0 ([3eefca96](https://github.com/bridgewwater/template-golang-lib/commit/3eefca963c6c837aab17467d2bb4bf1862d102c9))

* z-MakefileUtils/MakeGoMod.mk go vet use -tags ([77be0cbf](https://github.com/bridgewwater/template-golang-lib/commit/77be0cbf00f5a556b6a9ac756da06d048ffc1ec3))

* update local make kit ([5466b928](https://github.com/bridgewwater/template-golang-lib/commit/5466b92833d66b047ffa31e676d36927fd632bfc))

## 1.0.0 (2024-01-22)

### 🐛 Bug Fixes

* change go-release-platform defaults working-directory not settings ([61a855b0](https://github.com/bridgewwater/template-golang-lib/commit/61a855b080fa032ca09176dfdd8b633ca95929ef))

* fix tools of shell may parse error not show help ([84127074](https://github.com/bridgewwater/template-golang-lib/commit/841270747684ddb3509ce04cf5e5f51b207fe53e))

### ✨ Features

* change to go1.19.10 for build ([e5f7efb7](https://github.com/bridgewwater/template-golang-lib/commit/e5f7efb7835fcff0e7e214a8759388ada928cf49))

* update build template config ([1c2e262a](https://github.com/bridgewwater/template-golang-lib/commit/1c2e262a79031dc9f7a4e056e7e4adee3ef6a015))

* remove cross build ([30bfaa6e](https://github.com/bridgewwater/template-golang-lib/commit/30bfaa6e48a9cfbd025a85ae8d3fb2871cf82209))

* add go-release-platform defaults settings ([4c3a23c2](https://github.com/bridgewwater/template-golang-lib/commit/4c3a23c2a1edee72d059f6d444dbca1fda56a4bd))

* let go-build-platform to test build ([4c6a6738](https://github.com/bridgewwater/template-golang-lib/commit/4c6a6738f337f5fd4318a416896608922ade9497))

* update ci.yml use go-release-platform.yml ([a666f44e](https://github.com/bridgewwater/template-golang-lib/commit/a666f44ea88daede06218c6555ab0da5c79260ac))

* add workflows-template of golang-codecov.yml ([c3e881dd](https://github.com/bridgewwater/template-golang-lib/commit/c3e881dd29d4fec5fc8546cfe8344b1c04b02db3))

* add github.com/sebdah/goldie at example_test ([e2d2fdb7](https://github.com/bridgewwater/template-golang-lib/commit/e2d2fdb733390d1e49a8a6e7344e0cd3ed37be70))

* minimum go verison go1.18 ([a8908d46](https://github.com/bridgewwater/template-golang-lib/commit/a8908d4616a227d0e0757537f309695fc16bcb8f))

* change build ENTRANCE and update shell kit ([6cbb53a0](https://github.com/bridgewwater/template-golang-lib/commit/6cbb53a00ea26b19d280eb02b74988020bdffcce))

* change template-golang-lib update license by git info ([f79d6366](https://github.com/bridgewwater/template-golang-lib/commit/f79d6366d6680e1695a5efb95541fd38b89767b4))

### ♻ Refactor

* update temp-golang-lib ([e4a317e9](https://github.com/bridgewwater/template-golang-lib/commit/e4a317e9d1a05b2c38c227a13868e6426cac55cb))

### 👷‍ Build System

* support actions/upload-artifact/tree/v4 ([74dc23eb](https://github.com/bridgewwater/template-golang-lib/commit/74dc23eb432eaed4467665852f7b3b16440beb18))

* bump actions/download-artifact from 3 to 4 ([dc0d8fa1](https://github.com/bridgewwater/template-golang-lib/commit/dc0d8fa12f8958f21d6a8c3c81f48ec8fbc24d47))

* bump actions/upload-artifact from 3 to 4 ([7af875de](https://github.com/bridgewwater/template-golang-lib/commit/7af875de8f5138ccca92eeff8dc4d1fdb0b5a263))

* bump actions/setup-go from 4 to 5 ([8bb1f05b](https://github.com/bridgewwater/template-golang-lib/commit/8bb1f05b3993b422d0510ade88d50936544ad4cd))

* change golangci/golangci-lint-action use version latest ([26962fdb](https://github.com/bridgewwater/template-golang-lib/commit/26962fdb0e0ecf5c43c426bef457b18820ab175f))

* fix name of go-release-platform.yml ([3b72ac10](https://github.com/bridgewwater/template-golang-lib/commit/3b72ac10190735faae589886414523be5163ea52))

* bump actions/checkout from 3 to 4 ([f8859add](https://github.com/bridgewwater/template-golang-lib/commit/f8859addd5d0d7d4be6e25e47d3573213ad21085))

* update ci build with tags ([886f9452](https://github.com/bridgewwater/template-golang-lib/commit/886f9452d22904c87c90a9920e3ff67cc0a7ace7))

* update temp-golang-lib and remove modVendor at make task dep ([9eb5f330](https://github.com/bridgewwater/template-golang-lib/commit/9eb5f330835ef4df00c882c7bb61d51eb3416712))

* update temp-golang-lib ([3a02a426](https://github.com/bridgewwater/template-golang-lib/commit/3a02a426e3b18e16ce1a1dcb9f042a41c203a6d7))

* change version go1.19.12 to build ([8d775a07](https://github.com/bridgewwater/template-golang-lib/commit/8d775a07653a4696aaf897941c210b21fd2b7d15))

* update one ci for build ([7a259f51](https://github.com/bridgewwater/template-golang-lib/commit/7a259f51ba72802d2b9a27fbb64826da2d1549e4))

* update template-golang-lib ([9162a44a](https://github.com/bridgewwater/template-golang-lib/commit/9162a44a4233908f93d56fdc5b6c65e87bdf59fa))

* update template-golang-lib ([a47c0189](https://github.com/bridgewwater/template-golang-lib/commit/a47c018901b2b1b87d07fac4cfb3914521c76e8f))

* support windows busybox ([1242af08](https://github.com/bridgewwater/template-golang-lib/commit/1242af0880689cff3d6d6343733ed982f4b58815))

* fix docker build path error ([b8be5f08](https://github.com/bridgewwater/template-golang-lib/commit/b8be5f08f81a37967fd09b32510869e6f0457eb1))

* remove go tidy flag -x ([117635ed](https://github.com/bridgewwater/template-golang-lib/commit/117635ede5e70661ec64428f33d99ccddde7f109))

* change build macOSArm64 ([29803a5e](https://github.com/bridgewwater/template-golang-lib/commit/29803a5e7c2f79fe09ebac3d4c34001581ed6670))

* change MakeGoDist to support GNU/Make 4.4+ for windows ([cfc84876](https://github.com/bridgewwater/template-golang-lib/commit/cfc84876ce0d8587ad470966ab61a8ab75cca90f))

* make style task for local check ([891e4c4b](https://github.com/bridgewwater/template-golang-lib/commit/891e4c4be39fd6837a4d077605f1adf715822b1b))

* change golang-codecov.md only tag to send ([12c20903](https://github.com/bridgewwater/template-golang-lib/commit/12c2090368a5d7fab70fe3a7fb6e440783971383))

* update golang-codecov.md ([45003855](https://github.com/bridgewwater/template-golang-lib/commit/45003855f629487579844e5e80e91e23f7b86d0c))
