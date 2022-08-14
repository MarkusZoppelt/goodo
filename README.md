# GooDo

A simple command-line ToDo app written in Go.

    Usage:
      goodo [flags]
      goodo [command]
    
    Available Commands:
      addTask     Add a new Task to a given ToDo
      addTodo     Add a new ToDo
      completion  Generate the autocompletion script for the specified shell
      deleteTask  Delete a Task from a ToDo
      deleteTodo  Delete a ToDo
      help        Help about any command
      list        Prints all ToDos and Tasks
      updateTask  Update a ToDo's Task with a new name
      updateTodo  Update a ToDo's name and description
    
    Flags:
      -D, --DELETE   Deletes everything in the database.
      -h, --help     help for goodo

## Data Model

A TODO has a name and description.
Each TODO can have a list of TASKS associated with the TODO. 
Each TASK has only a name.

## Features:

 - create, read, update and delete TODOs.
 - remove TASKS from TODOs
 - update individual TASKS
