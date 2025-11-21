package model

import (
	"fmt"
	"strings"

	"example.com/goddlgen/pkg/logger"
)

type DataOrganization struct {
	TopLevelClasses map[string]*ClassDef
	Enumerations    map[string]*ClassDef
	UnknownClasses  map[string]*ClassDef
}

type ClassDef struct {
	SrcData    *ClassFile
	Subclasses map[string]*ClassDef
	Parent     *ClassDef
	Fields     map[string]*ClassDef
}

func (cd *ClassDef) PkgFqn() string {
	return fmt.Sprintf("%s.%s.%s", cd.SrcData.Ext.NamespacePrefix, cd.SrcData.Ext.NamespaceURI, cd.SimpleName())
}

func (cd *ClassDef) Name() string {
	return cd.SrcData.Name
}

func (cd *ClassDef) SimpleName() string {
	nameParts := strings.Split(cd.SrcData.Name, ".")
	return nameParts[len(nameParts)-1]
}

func NewClassDef(srcData *ClassFile) *ClassDef {
	return &ClassDef{
		SrcData:    srcData,
		Subclasses: make(map[string]*ClassDef),
		Parent:     nil,
		Fields:     make(map[string]*ClassDef),
	}
}

func OrganizeJson(classFiles []*ClassFile) *DataOrganization {
	log := logger.Get()

	// as we organize the data, the active map will reflect what hasn't yet been processed
	var activeMap = make(map[string]*ClassDef)

	// fill the activemap with everything
	for _, classFile := range classFiles {
		activeMap[classFile.Name] = NewClassDef(classFile)
	}
	log.Info().Msgf("Active Map has %d classes", len(activeMap))

	// find enumerations, removing them from the activemap
	enumsMap := make(map[string]*ClassDef)
	for _, classFile := range classFiles {
		if classFile.ClassType == "Enumeration" {
			enumsMap[classFile.Name] = activeMap[classFile.Name]
			delete(activeMap, classFile.Name)
		}
	}

	// find the root defs, removing them from the activemap
	rootsMap := make(map[string]*ClassDef)
	for _, classFile := range classFiles {
		if classFile.SuperType == "" && enumsMap[classFile.Name] == nil {
			rootsMap[classFile.Name] = activeMap[classFile.Name]
			delete(activeMap, classFile.Name)
		}
	}
	log.Info().Msgf("Organizing %d root classes", len(rootsMap))

	// attach the inheritance hierarchy
	buildInheritanceHierarchy(activeMap, rootsMap)

	log.Info().Msgf("Organizing complete!")

	return &DataOrganization{
		TopLevelClasses: rootsMap,
		Enumerations:    enumsMap,
		UnknownClasses:  activeMap,
	}
}

// builds the class inheritance hierarchy by looking at the activeMap (things that are not roots)
// and recursively traversing the roots tree to attach the parent-child associations.  The end result
// should be an empty activeMap and a completed rootsMap.  However, if there were definitions
// in the original JSON that specified connections not known, then the activemap may end up
// containing these items.
func buildInheritanceHierarchy(activeMap map[string]*ClassDef, rootsMap map[string]*ClassDef) {
	log := logger.Get()
	looking := true
	lastActiveLen := len(activeMap)
	for looking {
		foundMap := make(map[string]*ClassDef)
		for childName, child := range activeMap {
			for rootName, root := range rootsMap {
				log.Debug().Msgf("Processing root %s", rootName)
				parent := findParentTopDown(root, child)
				if parent != nil {
					parent.Subclasses[childName] = child
					child.Parent = parent
					foundMap[childName] = child
					break
				}
			}
		}
		// remove found from active
		for foundName := range foundMap {
			delete(activeMap, foundName)
			lastActiveLen = len(activeMap)
		}

		if len(activeMap) == 0 {
			looking = false
		} else if len(activeMap) == lastActiveLen {
			looking = false
			log.Warn().Msgf("The count of classes that specify unknown parents is %d ", len(activeMap))
		}
	}
}

func findParentTopDown(treeRoot *ClassDef, child *ClassDef) *ClassDef {
	if child.SrcData.SuperType == treeRoot.SrcData.Name {
		return treeRoot
	}

	var parent *ClassDef = nil
	for _, subclass := range treeRoot.Subclasses {
		if child.SrcData.SuperType == subclass.SrcData.Name {
			parent = subclass
		} else {
			parent = findParentTopDown(subclass, child)
		}
	}
	return parent
}
