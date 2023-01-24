- http://localhost:8080/
- http://localhost:8080/foo
- http://localhost:8080/favicon.ico
- http://localhost:8080/public/assets/images/example.png

# Embedding Files

The go command now supports including static files and file trees as part of the final executable, using the new `//go:embed` directive. See the documentation for the new [embed](https://tip.golang.org/pkg/embed/) package for details.

See: [Go 1.16 Release Notes](https://tip.golang.org/doc/go1.16#embed)
