# Acervus Project Creation Guide

This guide details the step-by-step process of creating a software development project.

## 1. User Registration and Login

The first step involves user registration and login. If you are registered, you can log in. If not, you must register first.

**Registration Command**:

	    auth register -e <email> -p <password>

This command registers a user with the specified email address and password.

## 2. Adding ABI File and Generating Settings File

In this step, we add the ABI file to the system and create a settings file.

**Required Information for Settings**:
- Project Name
- Project Description
- Network
- Contract Address
- Contract Name
- ABI File Path
- Schema File Path
- Selected Events
- CRUD Events
- Start Block
- End Block

**Settings File Creation Command**:

      generate settings

The generate settings command is used to create the settings file.

## 3. Creating the Project

Now that our ABI file and settings file are ready, we can create our project.

 **Project Creation Command**:

	    projects create

This command will create the project in the specified directory or in the default (project) directory.

## 4. Reviewing and Editing the Project File

We can review the created project file and edit the handle functions in the `project.ts` file as necessary.

## 5. Testing the Project

Before deploying, we need to check that the system is working correctly with the 

	    test -i <projectID>

command.

Here, `projectID` is the identifier of the project to be tested.

## 6. Deploying the Project

If the test is successful, we are ready to deploy the project.

 **Deploy Command**:

	    deploy -i <projectID>

Here, `projectID` is the identifier of your project.

--- 

This guide explains how to create and deploy a software development project step by step. By carefully following each step, you can successfully complete your project.
