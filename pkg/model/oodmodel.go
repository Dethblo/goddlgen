package model

import (
	"fmt"
	"strings"

	"example.com/goddlgen/pkg/logger"
)

type DataOrganization struct {
	TopLevelClasses map[string]*ClassDef
	Enumerations    map[string]*ClassDef
}

type ClassDef struct {
	ClassName  *string
	ClassData  *ClassFile
	Subclasses map[string]*ClassDef
	ParentName *string
	Parent     *ClassDef
	Fields     map[string]*ClassDef
}

func (cd *ClassDef) PkgFqn() string {
	return fmt.Sprintf("%s.%s.%s", cd.ClassData.Ext.NamespacePrefix, cd.ClassData.Ext.NamespaceURI, cd.SimpleName())
}

func (cd *ClassDef) SimpleName() string {
	nameParts := strings.Split(cd.ClassData.Name, ".")
	return nameParts[len(nameParts)-1]
}

func NewClassDef(srcData *ClassFile) *ClassDef {
	return &ClassDef{
		ClassName:  &srcData.Name,
		ClassData:  srcData,
		Subclasses: make(map[string]*ClassDef),
		ParentName: &srcData.SuperType,
		Parent:     nil,
		Fields:     make(map[string]*ClassDef),
	}
}

func OrganizeJson(classFiles []*ClassFile) *DataOrganization {
	log := logger.Get()

	// working copy of the class files
	var activeMap = make(map[string]*ClassFile)
	for _, classFile := range classFiles {
		activeMap[classFile.Name] = classFile
	}

	// find enumerations and remove them from working copy
	enumsMap := make(map[string]*ClassDef)
	for _, classFile := range classFiles {
		if classFile.ClassType == "Enumeration" {
			enumsMap[classFile.Name] = NewClassDef(classFile)
			delete(activeMap, classFile.Name)
		}
	}

	log.Info().Msgf("Active Map has %d classes", len(activeMap))

	// attach the inheritance hierarchy
	tree := BuildTree(classFiles)

	log.Info().Msgf("Tree has %d classes", len(tree))
	log.Info().Msgf("Organizing complete!")

	return &DataOrganization{
		TopLevelClasses: tree,
		Enumerations:    enumsMap,
	}
}

// BuildTree construct a tree from the JSON class files
func BuildTree(classFiles []*ClassFile) map[string]*ClassDef {
	// Create a map of all nodes by Name for easy lookup
	nodeMap := make(map[string]*ClassDef)

	for _, entry := range classFiles {
		nodeMap[entry.Name] = &ClassDef{
			ClassName:  &entry.Name,
			ClassData:  entry,
			Subclasses: make(map[string]*ClassDef),
			ParentName: &entry.SuperType,
			Parent:     nil,
			Fields:     make(map[string]*ClassDef),
		}
	}

	// Build the tree by connecting children to parents
	roots := make(map[string]*ClassDef)

	for _, entry := range classFiles {
		node := nodeMap[entry.Name]

		if entry.SuperType == "" {
			// This is a root node
			roots[entry.Name] = node
		} else {
			// Find the parent and add this node as a child
			parent, exists := nodeMap[entry.SuperType]
			if exists {
				parent.Subclasses[entry.Name] = node
				node.Parent = parent
			} else {
				// Parent doesn't exist, treat as orphan/root
				roots[entry.Name] = node
			}
		}
	}

	return roots
}
