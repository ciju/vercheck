// A bad, but sufficient implementation of version check. Checks if
// version on the repo is same as on the local.  Might not work if
// version numbers go beyond single digit. Probably it should also skip minor version updates.
// package vercheck
package vercheck

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
)

// check if its a http url, otherwise read from filesystem.

func VerArr(v string) []int {
	var err error
	ver := []int{0, 0, 0}
	split := strings.Split(v, ".")
	for i := 0; i < len(ver); i++ {
		ver[i], err = strconv.Atoi(split[i])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Can't convert version")
		}
	}
	return ver
}

func CompareVer(v1, v2 string, l int) bool {
	v1a, v2a := VerArr(v1), VerArr(v2)

	if l == 0 {
		l = len(v1a)
	}

	for i := 0; i < l; i++ {
		if v1a[i] > v2a[i] {
			return true
		} else if v1a[i] < v2a[i] {
			break
		}
	}
	return false
}

func getVersionFile(u string) string {
	ver, err := ioutil.ReadFile(u)
	if err != nil {
		return "0.0.0"
	}

	return strings.TrimSpace(string(ver))
}

func getVersionURL(u string) string {
	resp, err := http.Get(u)
	if err != nil || resp.StatusCode != http.StatusOK {
		return "0.0.0"
	}
	defer resp.Body.Close()

	ver, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "0.0.0"
	}
	return strings.TrimSpace(string(ver))
}

func getVersion(us string) string {
	u, err := url.Parse(us)
	if err != nil {
		return "0.0.0"
	}

	if u.IsAbs() {
		return getVersionURL(us)
	}
	return getVersionFile(us)
}

// this will even get a patch update ex: 1.2.3 - 1.2.4
func HasUpdate(rem, loc string) bool {
	return CompareVer(getVersion(rem), getVersion(loc), 0)
}

func HasMinorUpdate(rem, loc string) bool {
	return CompareVer(getVersion(rem), getVersion(loc), 2)
}

func HasMajorUpdate(rem, loc string) bool {
	return CompareVer(getVersion(rem), getVersion(loc), 1)
}
