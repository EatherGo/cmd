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
