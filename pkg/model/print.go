package model

import (
	"fmt"
	"sort"

	"github.com/xlab/treeprint"
)

func PrintClassHierarchy(rootMap map[string]*ClassDef) {
	tree := treeprint.New()

	// sort by keys
	keys := make([]string, 0)
	for key := range rootMap {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	for _, className := range keys {
		branch := tree.AddBranch(getPrefix(rootMap[className]) + className)
		addChildren(rootMap[className].Subclasses, branch)
	}

	fmt.Println(tree.String())
}

func addChildren(childMap map[string]*ClassDef, parent treeprint.Tree) {
	if len(childMap) == 0 {
		return
	}

	// sort by keys
	keys := make([]string, 0)
	for key := range childMap {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	for _, childName := range keys {
		branch := parent.AddBranch(getPrefix(childMap[childName]) + childName)
		addChildren(childMap[childName].Subclasses, branch)
	}
}

func getPrefix(cf *ClassDef) string {
	prefix := ""
	if cf.ClassData.ClassType == "AbstractClass" {
		prefix = "(A) "
	} else if cf.ClassData.ClassType == "NormalClass" {
		prefix = "(N) "
	}
	return prefix
}
