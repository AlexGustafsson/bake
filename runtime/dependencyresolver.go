package runtime

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	git "github.com/go-git/go-git/v5"
	log "github.com/sirupsen/logrus"
)

type DependencyResolver struct {
	// root is the root path on disk where dependencies are stored
	root string
	// importedPaths are the set of imported paths that have been resolved
	importedPaths map[string]bool
	// ResolvedPackages are all of the resolved and parsed packages, by their import path
	ResolvedPackages map[string]*Package
}

type Import interface {
	// Path is the full path of the import, as specified in the imports declaration
	Path() string
	// Name is the name of the imported package
	Name() string
	// Root is the absolute path to the resolved package on disk
	Root() string
	// Fetch fetches the package, ensuring it's stored on disk
	Fetch() error
	// Exists checks whether or not the package exists on disk
	Exists() (bool, error)
}

type Package struct {
	Modules []*Module
}

type BaseImport struct {
	path string
	name string
	root string
}

type GitHubImport struct {
	BaseImport
	Author         string
	Repository     string
	PackagePath    string
	RepositoryRoot string
}

func (baseImport *BaseImport) Name() string {
	return baseImport.name
}

func (baseImport *BaseImport) Path() string {
	return baseImport.path
}

func (baseImport *BaseImport) Root() string {
	return baseImport.root
}

func CreateDependencyResolver(root string) *DependencyResolver {
	return &DependencyResolver{
		root:             root,
		importedPaths:    make(map[string]bool),
		ResolvedPackages: make(map[string]*Package),
	}
}

func (resolver *DependencyResolver) Resolve(imports []string) []error {
	errs := make([]error, 0)

	for len(imports) > 0 {
		// Pop the top path
		path := imports[0]
		imports = imports[1:]

		// Resolve type of import
		var resolvedImport Import
		var err error
		if strings.HasPrefix(path, "github.com/") {
			resolvedImport, err = CreateGitHubImport(path, resolver.root)
		} else {
			err = fmt.Errorf("invalid import '%s'", path)
		}
		// Failed to resolve import
		if err != nil {
			errs = append(errs, err)
			continue
		}

		// Ignore already resolved paths
		if _, ok := resolver.importedPaths[resolvedImport.Path()]; ok {
			continue
		}
		resolver.importedPaths[resolvedImport.Path()] = true

		// Check if it already exists
		exists, err := resolvedImport.Exists()
		if err != nil {
			errs = append(errs, err)
			continue
		}

		// Fetch
		if !exists {
			log.Infof("fetching repository for package '%s' ('%s')", resolvedImport.Name(), resolvedImport.Path())
			err = resolvedImport.Fetch()
			// Failed to fetch
			if err != nil {
				errs = append(errs, err)
				continue
			}
		}

		stat, err := os.Stat(resolvedImport.Root())
		if err != nil {
			errs = append(errs, err)
			continue
		}

		if !stat.IsDir() {
			errs = append(errs, fmt.Errorf("incorrectly resolved dependency '%s' - not a directory", resolvedImport.Root()))
			continue
		}

		entries, err := os.ReadDir(resolvedImport.Root())
		if err != nil {
			errs = append(errs, err)
			continue
		}

		resolvedPackage := &Package{
			Modules: make([]*Module, 0),
		}

		// Parse and resolve each module in the package
		for _, entry := range entries {
			if entry.Type().IsRegular() {
				log.Debugf("resolving new file '%s'", entry.Name())

				// Read the file
				file, err := os.Open(resolvedImport.Root() + "/" + entry.Name())
				if err != nil {
					errs = append(errs, err)
					continue
				}
				inputBytes, err := ioutil.ReadAll(file)
				if err != nil {
					errs = append(errs, err)
					continue
				}
				input := string(inputBytes)

				// Parse the module
				module := CreateModule(input)
				err = module.Parse()
				if err != nil {
					errs = append(errs, err)
					continue
				}

				// Add the imports to the paths to resolve
				imports = append(imports, module.Imports()...)

				// Add the module to the resolved package
				resolvedPackage.Modules = append(resolvedPackage.Modules, module)
			}
		}

		resolver.ResolvedPackages[resolvedImport.Path()] = resolvedPackage
	}

	return errs
}

func CreateGitHubImport(path string, root string) (*GitHubImport, error) {
	parts := strings.Split(path, "/")
	if len(parts) < 3 {
		// The URL is not in the format of "github.com/" author "/" repo [ "/path/to/package" ]
		return nil, fmt.Errorf("invalid import '%s'", path)
	}

	author := parts[1]
	repository := parts[2]
	packagePath := strings.Join(parts[3:], "/")
	packageName := parts[len(parts)-1]

	// TODO: Safe and actual path resolve
	repositoryRoot := fmt.Sprintf("%s/%s/%s", root, author, repository)
	packageRoot := fmt.Sprintf("%s/%s", repositoryRoot, packagePath)

	return &GitHubImport{
		BaseImport: BaseImport{
			name: packageName,
			path: path,
			root: packageRoot,
		},
		Author:         author,
		Repository:     repository,
		PackagePath:    packagePath,
		RepositoryRoot: repositoryRoot,
	}, nil
}

func (gitHubImport *GitHubImport) Exists() (bool, error) {
	stat, err := os.Stat(gitHubImport.root)
	if err != nil {
		return false, err
	}

	return stat.IsDir(), nil
}

func (gitHubImport *GitHubImport) Fetch() error {
	if stat, err := os.Stat(gitHubImport.root); err == nil && stat.IsDir() {
		log.Infof("skipping '%s' ('%s') - already exists", gitHubImport.Name(), gitHubImport.Path())
		return nil
	}

	repositoryURL := fmt.Sprintf("https://github.com/%s/%s", gitHubImport.Author, gitHubImport.Repository)

	// TODO: Sparse checkout https://github.com/go-git/go-git/issues/90
	_, err := git.PlainClone(gitHubImport.RepositoryRoot, false, &git.CloneOptions{
		URL:   repositoryURL,
		Depth: 1,
	})
	if err != nil {
		return err
	}

	return nil
}
