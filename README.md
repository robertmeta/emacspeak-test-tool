# Emacspeak Test Tool

## Quickstart

```python3 ett.py -p <server_binary> -s <test_script>```

Example:

```python3 ett.py -p ~/projects/robertmeta/swiftmac/.build/debug/swiftmac -s swiftmac.etts```

## Script Language

The etts scripts attempt to be as close to the log output of 
emacspeak logs as possible, they only add a handful of 
features

Expansions: 

- $ETT_SD: will be expanded to the directory the script lives
in, useful for referencing the ogg files for testing. 


Commands (go on their own line):

- DELAY <seconds>: can be decimals, like 1.1 or .2
- RESTART: restart the server (useful to read startup ENV)
- ENV <name> <value>: set a variable in env (useful with restart)
- ENVCLEAR: clear the environment
- START_SKIP: when it sees a start skip it will stop processing
until you do ```END_SKIP```, this lets you focus on one test 
or one set without making a new file
- END_SKIP: stops skipping, continues regular processing


## Goals

The goal of this tiny python script is to let you take emacspeak
log files generated by the log-{server} variants and create 
reproducable sharable scritps that will show any bugs / issues 
you encounter. 

I will also be using it to write complete test scritps for 
swiftmac and SharpWin servers.

It is written in standard python3.


