# Releasing zibble

`zibble` is the root Go module:

```text
github.com/xuan-cao-swi/sample-go-package
```

## Release checklist

1. Choose the next semantic version, such as `0.2.0`.
2. Update `Version` in [version.go](version.go) to the same version without the `v` prefix.
3. Run the tests and tidy the module:

	```shell
	go test ./...
	go mod tidy
	```

4. Commit the version and dependency changes, then create and merge the pull request.
5. After the pull request is merged, tag the release from the merge commit and push the tag:

	```shell
	git tag v0.2.0
	git push origin v0.2.0
	```

6. Create a GitHub Release for the tag at <https://github.com/xuan-cao-swi/sample-go-package/releases/new>.

The Go toolchain uses the Git tag as the module version. `Version` is a package-level release marker and should stay aligned with the tag.

## Releasing swozibble

`swozibble` is a separate module under `instrumentation/` and must be tagged independently. Keep its release version aligned with the root `zibble` release.

Publish the root `zibble` tag first. Then update the `require github.com/xuan-cao-swi/sample-go-package` version in `instrumentation/github.com/xuancao/zibble/swozibble/go.mod` to the new root version and remove any local `replace` directive before releasing `swozibble`.

Before tagging it, run the checks from its directory:

```shell
cd instrumentation/github.com/xuancao/zibble/swozibble
go test ./...
go mod tidy
```

While the root module is only available in the local checkout, use a temporary replacement before running those commands:

```shell
go mod edit -replace github.com/xuan-cao-swi/sample-go-package=../../../../..
go mod tidy
go test ./...
```

Remove that replacement and require the published root version before creating the release tag.

From the repository root, tag the module using its directory path followed by the version:

```shell
git tag instrumentation/github.com/xuancao/zibble/swozibble/v0.2.0
git push origin instrumentation/github.com/xuancao/zibble/swozibble/v0.2.0
```

Clients can then add the integration with:

```shell
	go get github.com/xuan-cao-swi/sample-go-package/instrumentation/github.com/xuancao/zibble/swozibble@v0.2.0
```

## Major version changes

Per Go module versioning rules, a `v2.0.0` or later release requires `/v2` in the module path and imports. Update `go.mod`, package imports, and the release tags together before publishing a major version.
