# Acervus CLI

This document serves as a guide for users interacting with Acervus Cloud via Acervus Command Line Interface (CLI). It provides detailed instructions on the usage of various commands and project management within Acervus ecosystem.

## Overview

Acervus CLI allows users to create, manage, and interact with their projects on Acervus Cloud. For a comprehensive list of available commands and their functions, execute the following command in your [CLI](https://google.com) environment:

```
acervus help
```

Using `--help` flag with commands will display additional information about that command.

## Installation

The installation process is currently being developed and will be provided in the future.

## Usage Instructions

### Create Account

To utilize Acervus Cloud services, you must first create an account. Create your account using the command below:

```
acervus auth register -e <your_email_address> -p <your_password>
```

This command automatically logs you in after creating your account. It also creates and stores your authentication tokens inside of `./credentials.json` file.

### Renewing Refresh Tokens

If you lost your refresh tokens or want to regenerate them (maybe they are expired), you can do so by simply logging in with your account. To login:

```
acervus auth login -e <your_email_address> -p <your_password>
```

### Managing Multiple Accounts

To manage multiple accounts, register and login with new emails. However, before using the CLI, select the account you want to use (your last login is automatically selected as your currently active account). To list and select your currently active account from your `./credentials.json`, run:


```
acervus auth
```

This command will display your saved accounts, allowing you to select one as your currently active account.

### Creating and Developing a New Project

To start a new project, first generate a settings file:

```
acervus generate settings
```

You will need an `abi.json` file to start. This command will ask for required information and then generate a `./settings.yaml` file.

To actually create your project on server:

```
acervus projects create
```

This command initiates the creation of a new project on the server and provides you with a `projectID` upon successful completion.

While the following commands are executed automatically during the create process, they are also available for standalone use should you need to regenerate your project files.

To generate a GraphQL schema for your project:

```
acervus generate graphql -i <project_id>
```

This command generates GraphQL entities based on your ABI. You can manually modify the generated `./schema.graphql` file and regenerate it with this command. Notice that a GraphQL file is necessary to run the project.

To generate a boilerplate project:

```
acervus generate boilerplate -i <project_id>
```

This will create all necessary files in `./project/<project_id>`. Optionally, use `-d <project_dir>` flag to specify a different folder.

You can then open `./project/<project_id>/project.ts` and begin coding your application.

### Testing

Test your project on the cloud before deployment:

```
acervus test -p <path_to_project_file> -i <project_id> -s <settings_file>
```

Default `path_to_project_file` is `./project/<project_id>/project.ts`, and default `settings_file` is `./settings.yaml`. 

This command runs your script in the cloud with a mock event and returns results or error messages.

Be careful, you can not modify your project once it is deployed.

### Deploying Your Project

To deploy your project:

```
acervus deploy -p <path_to_project_file> -i <project_id> -s <settings_file>
```

Once deployed, your project starts running, and you can use the GraphQL API to query results.

Note that authentication is not required to create a project, but it is required for deployment and testing of a project.

### Managing Multiple Projects

You can create multiple projects on your local environment using `generate` commands. To manage multiple projects:

```
acervus projects
```

This lists your cloud-based projects. 

### Exporting Project Data

To export your project data:

```
acervus projects export -i <project_id> 
```

This sends a download link to your account email.

### Migrating from Other Platforms

To migrate an existing project to Acervus Cloud:

```
acervus migrate -d <target_dir> -s <source_platform>
```

Currently supported platform for `-s` flag are `Subgraph`.


### Pausing / Resuming a Project

To pause a project temporarily:

```
acervus projects pause -i <project_id>
```

To resume a paused project:

```
acervus projects resume -i <project_id>
```

### Querying data and logs

To get resulting data:

```
acervus query data -i <project_id>
```

To get project logs:

```
acervus query logs -i <project_id>
```