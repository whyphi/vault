# vault
Whyphi's Access Management System

## Instructions for Managing Access
To grant or revoke access to the Whyphi Access Management System, follow these steps:

1. Open the `vault.yaml` file.

2. Add the name and email parameters of the user you want to grant access to.

3. Please ensure that the user you are adding is using their BU (Boston University) email address, as the system only authorizes BU email domains.
Create a pull request (PR) with your changes to the vault.yaml file.

4. Have another team member review your PR to ensure that the `vault.yaml` file is correctly formatted and the correct users are being added.

5. Once the PR has been reviewed and approved, merge it into the main branch.

6. Merging the PR into main will trigger the `main.go` script, which will synchronize the data from `vault.yaml` with our MongoDB Users database.

## Why Was Vault Built?
Vault was created to simplify access management for Whyphi while avoiding unnecessary complexity:

- Simplified User Management: Managing user access for Whyphi can be a challenging task. While we could have developed a server with APIs to add or delete users, this approach may have introduced overengineering and added more projects to maintain.

- Integration with GitHub: Since the majority of Whyphi's codebase resides on GitHub, having a "database" in the form of vault.yaml on GitHub provides a convenient and easily modifiable solution.

- User-Friendly YAML Format: The choice of YAML format for vault.yaml makes it easy for newcomers to understand and edit. This simplicity encourages team members to add their information to vault.yaml with ease.

- Learning Opportunity: Vault also serves as an opportunity for team members to learn a new programming language, Go, and work with new technologies, contributing to personal and professional growth.

By following the provided instructions, you can efficiently manage access to Whyphi's Access Management System and help maintain an organized and up-to-date user database.
