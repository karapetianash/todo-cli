## Introduction
This CLI is a command-line tool that allows you to create and manage your To-Do lists efficiently. This tool provides a simple and intuitive way to keep track of your tasks and stay organized.

## Supported methods
- Add:
    Creates a new to-do item and appends it to the list.
- Complete:
    Marks a to-do item as completed.
- Delete:
    Deletes a to-do item from the list.
- Save:
    Saves the list of items to a file using the JSON format.
- Get:
    Obtains a list of items from a saved JSON file.

## Usage
To build the app, run the following command in the root folder:

```
> go build .
```
Above command will generate todo-cli file. This name is defined in the go.mod file and it will be the initialized module name.

You can set an environment variable TODO_FILENAME to name the file where your tasks will be stored.

```
> set TODO_FILENAME=someName
```
After that you can run the file using the cmd and pass the task:

```
> .\todo-cli.exe -add New Task
```