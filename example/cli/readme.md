# cli

This application provides an example of extensible applications. There are two primary concepts, 
'App' and 'Command'.

An 'App' is defined as a function which re-writes a context.Context and is registered globally.

A 'Command' is defined as a function which takes in a Message. An 'App' is responsible
for adding 'Command's  to the context, and the CLI code delegates to the Commands based on
the current state of the context.

## commmands

global commands:

 * app $appName - switch to app
 * quit - quit cli

shell commands:
 
 * echo $content... - echo the arguments
 * ls - list the directory contents (fake)

http commands:

 * get $url - Run GET against the given URL, dumping the response headers and body
 * head $url - Run HEAD against the given URL, dumping the response headers and body
