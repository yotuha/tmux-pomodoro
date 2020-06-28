## COMMAND

```
tmux-pomodoro <command> [<args>]
command: start, stop, help
```

## CONFIG EXAMPLE

Priority: args > config

.pomodoro.yaml

```
set: 1
work: 1500
rest: 300
afterWorkRunCmd: - "terminal-notifier -message \"Done pomodoro\" -title \"pomodoro\"" - "afplay /System/Library/Sounds/Purr.aiff"
afterRestRunCmd: - "terminal-notifier -message \"Finish rest\" -title \"pomodoro\"" - "afplay /System/Library/Sounds/Purr.aiff"
```

.tmux.conf

```
tm_pomodoro="#(cat ~/.pomodoro)"
set -g status-right \$tm_pomodoro
set -g status-interval 1
```
