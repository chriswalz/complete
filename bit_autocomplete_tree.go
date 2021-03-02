package complete

func getBitTree() *CompTree {
	var bitCompleteTree = &b
	for k, v := range descriptionMap {
		if bitCompleteTree.Sub[k] == nil {
			continue
		}
		bitCompleteTree.Sub[k].Desc = v
	}
	return bitCompleteTree
}

var descriptionMap = map[string]string{
	"add":             "Add file contents to the index",
	"am":              "Apply a series of patches from a mailbox",
	"archive":         "Create an archive of files from a named tree",
	"branch":          "List, create, or delete branches",
	"bisect":          "Use binary search to find the commit that introduced a bug",
	"bundle":          "Move objects and refs by archive",
	"commit":          "Record changes to the repository",
	"clone":           "Clone a repository into a new directory",
	"checkout":        "Switch branches or restore working tree files",
	"co":              "Switch branches or restore working tree files",
	"fetch":           "Download objects and refs from another repository",
	"diff":            "Show changes between commits, commit and working tree, etc",
	"cherry-pick":     "Apply the changes introduced by some existing commits",
	"citool":          "Graphical alternative to git-commit",
	"clean":           "Remove untracked files from the working tree",
	"describe":        "Give an object a human readable name based on an available ref",
	"format-patch":    "Prepare patches for e-mail submission",
	"gc":              "Cleanup unnecessary files and optimize the local repository",
	"gitk":            "The Git repository browser",
	"grep":            "Print lines matching a pattern",
	"gui":             "A portable graphical interface to Git",
	"init":            "Create an empty Git repository or reinitialize an existing one",
	"log":             "Show commit logs",
	"merge":           "Join two or more development histories together",
	"mv":              "Move or rename a file, a directory, or a symlink",
	"notes":           "Add or inspect object notes",
	"pull":            "Fetch from and integrate with another repository or a local branch",
	"push":            "Update remote refs along with associated objects",
	"range-diff":      "Compare two commit ranges (e.g. two versions of a branch)",
	"rebase":          "Reapply commits on top of another base tip",
	"reset":           "Reset current HEAD to the specified state",
	"restore":         "Restore working tree files",
	"revert":          "Revert some existing commits",
	"rm":              "Remove files from the working tree and from the index",
	"show":            "Show various types of objects",
	"stash":           "Stash the changes in a dirty working directory away",
	"shortlog":        "Summarize 'git log' output",
	"status":          "Show the working tree status",
	"submodule":       "Initialize, update or inspect submodules",
	"switch":          "Switch branches",
	"tag":             "Create, list, delete or verify a tag object signed with GPG",
	"worktree":        "Manage multiple working trees",
	"config":          "Get and set repository or global options",
	"fast-import":     "Backend for fast Git data importers",
	"filter-branch":   "Rewrite branches",
	"mergetool":       "Run merge conflict resolution tools to resolve merge conflicts",
	"pack-refs":       "Pack heads and tags for efficient repository access",
	"prune":           "Prune all unreachable objects from the object database",
	"reflog":          "Manage reflog information",
	"remote":          "Manage set of tracked repositories",
	"rename":          "",
	"remove":          "",
	"set-head":        "",
	"repack":          "Pack unpacked objects in a repository",
	"replace":         "Create, list, delete refs to replace objects",
	"annotate":        "Annotate file lines with commit information",
	"blame":           "Show what revision and author last modified each line of a file",
	"count-objects":   "Count unpacked number of objects and their disk consumption",
	"difftool":        "Show changes using common diff tools",
	"fsck":            "Verifies the connectivity and validity of the objects in the database",
	"gitweb":          "Git web interface (web frontend to Git repositories)",
	"help":            "Display help information about Git",
	"instaweb":        "Instantly browse your working repository in gitweb",
	"merge-tree":      "Show three-way merge without touching index",
	"rerere":          "Reuse recorded resolution of conflicted merges",
	"show-branch":     "Show branches and their commits",
	"verify-commit":   "Check the GPG signature of commits",
	"verify-tag":      "Check the GPG signature of tags",
	"whatchanged":     "Show logs with difference each commit introduces",
	"archimport":      "Import a GNU Arch repository into Git",
	"cvsexportcommit": "Export a single commit to a CVS checkout",
	"cvsimport":       "Salvage your data out of another SCM people love to hate",
	"cvsserver":       "A CVS server emulator for Git",
	"imap-send":       "Send a collection of patches from stdin to an IMAP folder",
	"p4":              "Import from and submit to Perforce repositories",
	"fast-export":     "Git data exporter",
	"version":         "Print bit and git version",
	"--version":       "Print bit and git version",
	"release":         "Commit unstaged changes, bump minor tag, push",
	"pr":              "Check out a pull request from Github (requires GH CLI)",
	"info":            "Get general information about the status of your repository",
	"gitmoji":         "(Pre-alpha) Commit using gitmojis",
	"save":            "Save your changes to your current branch",
	"update":          "Updates bit to the latest or specified version",
	"complete":        "Add classical tab completion to bit",
	"sync":            "Synchronizes local changes with changes on origin or specified branch",
}
