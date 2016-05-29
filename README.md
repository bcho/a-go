# a

A simple file organizer.

## Usage

### Dry run by default

```shell
$ a $HOME/Desktop -fmt year-month
Will move `Screen Shot 2016-04-29 at 8.04.33 PM.png` to `2016-04/`
Will move `Screen Shot 2016-05-29 at 8.04.33 PM.png` to `2016-05/`
```

### Run

```shell
$ a $HOME/Desktop -fmt year-month -x
Moving `Screen Shot 2016-04-29 at 8.04.33 PM.png` to `2016-04/`
Moving `Screen Shot 2016-05-29 at 8.04.33 PM.png` to `2016-05/`
```
