# vgo-walkthrough

This is my recording of trying to introduce vgo into a hypothetical monorepo
containing multiple go projects.

# vgo version

```
vgo version
go version go1.10 darwin/amd64 vgo:2018-02-20.1
```

# Notes on api.github.com usage

One of the most unexpected things that I encountered was the following
message:

> GitHub applies fairly small rate limits to unauthenticated users, and
> you appear to be hitting them. To authenticate, please visit
> https://github.com/settings/tokens and click "Generate New Token" to
> create a Personal Access Token. The token only needs "public_repo"
> scope, but you can add "repo" if you want to access private
> repositories too.
> 
> Add the token to your $HOME/.netrc (%USERPROFILE%\_netrc on Windows):
> 
>     machine api.github.com login YOU password TOKEN
> 
> Sorry for the interruption.

This is not recorded in the history of this repo, but it will happen when you
execute vgo enough times. This never seemed to happen on other build tools,
so it was kind of frustrating that I needed to create a token just for
this purpose

# Notes on repositories other than github

Apparently the only service that's well encoded is Github (from grepping the
sources). What that means is that, while `vgo build` automatically generated
pseudo-versions for included modules, you cannot at present expect the same
treatment. For example, I tried doing this on GitLab, and I needed to figure
out the modules versions to include in the go.mod