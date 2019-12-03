# EATHER Command Line Tools

Create and add new features to Eather application

## Install
```
go get -u github.com/EatherGo/cmd/eather
```

## Create New Application
```
eather create -n NewApp
```

## Create New Module
Make sure that you are in your application directory.
```
eather module -n EmptyModule
```

This will create all necessary files for an empty module. New module will be stored to your env variable `CUSTOM_MODULES_DIR`

### With Controller 
```
eather module -n EmptyModule -c
```

### With Events
```
eather module -n EmptyModule -e
```

### With Model
```
eather module -n EmptyModule -m MyModel
```

### With Upgrade
```
eather module -n EmptyModule -u
```

### With Cron
```
eather module -n EmptyModule -cr
```

### With Callable Func
```
eather module -n EmptyModule -ca
```

### To create module with all predefined funcs
```
eather module -n EmptyModule -f -m MyModel
```
Where -f flag is for full module including Controller, Events, Upgrade, Cron, Callable func. Or run it calling all flags.
```
eather module -n EmptyModule -c -e -u -cr -ca -m MyModel
```



