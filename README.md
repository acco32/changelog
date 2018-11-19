# Changelog

![Licence](https://img.shields.io/badge/license-MIT-green.svg?style=for-the-badge)

A command line utility to create changelogs. A variation of the [changelog utility](https://gitlab.com/gitlab-org/gitlab-ce/blob/master/bin/changelog) used in [gitlab](https://gitlab.com). Allows for the 

## Pre-requisites

- [golang](https://golang.org/) 1.11+
- [tusk](https://github.com/rliebz/tusk)
- [dep](https://github.com/golang/dep)

## Conventions
- **`Changelog.md`** must reside in root folder and must have top level header **# Changelog**
- Must run command in same folder as **`Changelog.md`**

## Workflow

1. Create a new feature on a branch
2. Create a new changelog entry for feature 
    - Creates a folder called `unreleased` in root folder
    - Multiple people can be on their own branches with their own entries
3. Merge work into release branch when ready
4. Preview the changelog `changelog release vX.X.X --preview`
5. Update the changelog `changelog release vX.X.X` 
    - Version string can follow your preferred format
    - Will create a file called `changelog.yml` with all changes in root folder
6. Commit the changes 
    - Unreleased folder will be deleted
    - Can optionally keep the yaml files
7. Cake

## Development

Once you are satisfied with the current version of the code run the following:

```shell
tusk test
tusk release --version=X.X.X
git add -u
git tag vX.X.X
git commit -m"Latest release"
```

This performs the following:
- Merge all elements into `changelog.yml`
- Add all elements to `CHANGELOG.md`
- copy the latest version to `tools` folder
- Add all the updated file (includes the version in source code and the tool itself)
- Tag the latest commit
- Submit changes