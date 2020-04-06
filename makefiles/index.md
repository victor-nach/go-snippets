## What is a makefile?
- makefile is a special file, containing shell commands, that you create and name makefile (or Makefile depending upon the system). While in the directory containing this makefile, you type make and the commands in the makefile will be executed. 
- Make keeps track of the last time files (normally object files) were updated and only updates those files which are required (ones containing changes) to keep the sourcefile up-to-date.
- As a makefile is a list of shell commands, it must be written for the shell which will process the makefile. A makefile that works well in one shell may not execute properly in another shell.

## specifying a variable
- variable_name = value
- no need to add strings to the variable name
- that a variable is set in a makefile doesn't mean it's now available as an environment variable
- also all environment variables available when make is run is also available as variable

## Including a .env file
- include .env
- this includes the contents of the env file so that they are now variables in the makefile's namespace
- they can now be used as other regular variables

## specifying an environment variable
- you can add `export variable_name = value` to the make file to explicitly set an environment variable
- when you use `include .env` the variables in the env are imported as make file variables but are not set as environment variables
- you can go ahead and use the export option while referencing those variables `export varaible_name`
- you can also set those variables as part of the scripts that make runs like so `variable_name='$(variable_name)' \`
- variable name then becomes available as an environment variable
- more convineintly you can use the `.EXPORT_ALL_VARIABLES: ` option to automatically export all the available variables to all sub-make calls
- the -e flag can be added when calling make `make -e` to import environment variables from the system (it works without it). 

## using the @ symbol
- this disables printing the command line being executed
- Any output from the command itself still appears

## using the $ symbol
- the two main uses of the $ are in ${} and $()
- ### $()
    - using the bracket option expands or evalueates what is inside it
    - may be a variable or an expression
- ### ${} 
    - the curly brace acts as a dereference, just get value of a variable.

## = vs :=
= defines a recursively-expanded variable. := defines a simply-expanded variable.


## using dependencies