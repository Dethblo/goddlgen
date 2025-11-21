package model

import (
	"fmt"

	"github.com/xlab/treeprint"
)

func PrintClassHierarchy(rootMap map[string]*ClassDef) {
	tree := treeprint.New()

	for _, class := range rootMap {
		branch := tree.AddBranch(class.Name())
		addChildren(class.Subclasses, branch)
	}

	fmt.Println(tree.String())
}

func addChildren(childMap map[string]*ClassDef, parent treeprint.Tree) {
	if len(childMap) == 0 {
		return
	}

	for _, child := range childMap {
		branch := parent.AddBranch(child.Name())
		addChildren(child.Subclasses, branch)
	}
}
