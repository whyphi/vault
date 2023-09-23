# vault
Whyphi's Access Management System

## Instructions
To give/remove access to Whyphi, edit `vault.yaml` by adding `name` and `email` parameter of the user. 

- Ensure that the user that is beind added is using their BU email as the client will not allow other email domains to be authorized
- Make a pull request with changes to `vault.yaml`
- Get the PR reviewed by another person, ensuring that `vault.yaml` is formatted correctly
- Merge the PR to `main`, in which it will trigger the `main.go` and synchronize `vault.yaml` with our MongoDB `Users` database

## Why was Vault build?
- Managing users for whyphi is very tough. We could create a server with APIs to add or delete users, but in essence, overengineering could happen and there would be more projects to maintain
- As majority of whyphi's codebase is in GitHub, having a "database" on GitHub in which it could be modified easily seems like an appropriate solution
- A YAML file is easy to read and edit, making it really easy for newcomers to add their information to `vault.yaml` very easily
- and... to learn a new language (Go) and work with new technologies!
