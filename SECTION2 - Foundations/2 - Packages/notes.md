# Notes about section 2

- To create a Go module type:
    - go mod init module

- To create a binary compiled from your Go code type:
    - go build   

- A function must have the first letter with uppercase if you want to export her, if not, use lowercase to make her private.

- To get a external package type:
    - go get <package_link> without http, example: 'go get github.com/badoux/checkmail'
- To remove unused dependencies type:
    - go mod tidy