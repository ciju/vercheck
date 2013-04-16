# vercheck #

Compare two semver versions. First one is considered to be source, and
second the version you want to check with. Ex: first one could be a
link to the version file, on internal. Ane second could be a constant
string which was shipped with your binary.

## API ##
- `HasUpdate(src, local)` Checks for any change in version
- `HasMinorUpdate(src, local)` Checks for minor updates. Example will
ignore `x` in `1.2.x`.
- `HasMajorUpdate(src, local)` Checks for major updates. Example will
ignore `x` in `1.x.x`.

Both `src` and `local` can be a URL, or a local file path, or actual
version string.


## Usage ##
```go
import "vercheck"
...

if vercheck.HasMinorUpdate(
   "https://raw.github.com/ciju/devmirror/master/VERSION",
   "1.2.3"
) {
   // do something
}
```

## LICENSE ##
MIT

**Sponsored by [ActiveSphere](http://activesphere.com)**
