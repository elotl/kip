package util

import (
	"reflect"
	"sort"
	"strings"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func StringInSlice(s string, slice []string) bool {
	for _, x := range slice {
		if s == x {
			return true
		}
	}
	return false
}

func StringMapIntersection(a, b map[string]string) map[string]string {
	// create a new map of the intersection
	c := make(map[string]string)
	for k, v1 := range a {
		v2, ok := b[k]
		if ok && v1 == v2 {
			c[k] = v1
		}
	}
	return c
}

func StringMapUnion(a, b map[string]string) map[string]string {
	// create a new map of the intersection
	c := make(map[string]string)
	for k, v := range a {
		c[k] = v
	}
	for k, v := range b {
		c[k] = v
	}
	return c
}

func StringMapDifference(a, b map[string]string) map[string]string {
	c := make(map[string]string)
	for k, v := range a {
		_, ok := b[k]
		if !ok {
			c[k] = v
		}
	}
	return c
}

func StringSliceIntersection(a, b []string) []string {
	c := make([]string, 0, max(len(a), len(b)))
	sort.Strings(a)
	sort.Strings(b)
	for _, v := range a {
		if sort.SearchStrings(b, v) < len(a) {
			c = append(c, v)
		}
	}
	return c
}

func StringSliceUnion(a, b []string) []string {
	c := make([]string, 0, max(len(a), len(b)))
	d := make(map[string]struct{})
	for _, v := range a {
		d[v] = struct{}{}
	}
	for _, v := range b {
		d[v] = struct{}{}
	}
	for k, _ := range d {
		c = append(c, k)
	}
	sort.Strings(c)
	return c
}

func StringSliceDifference(a, b []string) []string {
	c := make([]string, 0, len(a))
	sort.Strings(a)
	sort.Strings(b)
	for _, v := range a {
		if sort.SearchStrings(b, v) == len(a) {
			c = append(c, v)
		}
	}
	return c
}

func StringSliceEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	sort.Strings(a)
	sort.Strings(b)
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func StringMapKeys(a map[string]string) []string {
	c := make([]string, len(a))
	i := 0
	for k, _ := range a {
		c[i] = k
		i++
	}
	return c
}

func StringSliceMapKeys(a map[string][]string) []string {
	c := make([]string, len(a))
	i := 0
	for k, _ := range a {
		c[i] = k
		i++
	}
	return c
}

func StringMapValues(a map[string]string) []string {
	c := make([]string, len(a))
	i := 0
	for _, v := range a {
		c[i] = v
		i++
	}
	return c
}

func StringSliceUnique(a []string) []string {
	c := make([]string, 0, len(a))
	d := make(map[string]struct{})
	for _, v := range a {
		d[v] = struct{}{}
	}
	for k, _ := range d {
		c = append(c, k)
	}
	sort.Strings(c)
	return c
}

func StringSetKeys(a map[string]struct{}) []string {
	c := make([]string, len(a))
	i := 0
	for k, _ := range a {
		c[i] = k
		i++
	}
	return c
}

func StringSliceRemove(s []string, r string) []string {
	for i, v := range s {
		if v == r {
			return append(s[:i], s[i+1:]...)
		}
	}
	return s
}

func getStringField(obj interface{}, field string) string {
	r := reflect.ValueOf(obj)
	f := reflect.Indirect(r).FieldByName(field)
	return f.String()
}

func GetStructStringField(objs []interface{}, field string) []string {
	ss := make([]string, len(objs))
	for i := 0; i < len(objs); i++ {
		ss[i] = getStringField(objs[i], field)
	}
	return ss
}

// If we cared, would be a good place to use a prefix tree...
func FilterKeysWithPrefix(m map[string]string, prefixes []string) map[string]string {
	filtered := make(map[string]string)
	for k, v := range m {
		addLabel := true
		for _, lab := range prefixes {
			if strings.HasPrefix(k, lab) {
				addLabel = false
				break
			}
		}
		if addLabel {
			filtered[k] = v
		}
	}
	return filtered
}
