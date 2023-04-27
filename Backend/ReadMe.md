# Golang Backend

## packages

installation
```bat
go get -u github.com/sirupsen/logrus
go get -u github.com/spf13/viper
go get -u gorm.io/driver/mysql
go get -u gorm.io/gorm
go get -u github.com/labstack/echo/v4
go get -u github.com/swaggo/echo-swagger
go get -u github.com/go-playground/validator/v10

go get -u github.com/lib/pq
go get -u gorm.io/driver/postgres
go get -u gorm.io/driver/sqlite
```


## project

use 'ASSETS' folder to save images


## How to generate stories?

As is known to all, a story contains the following factors at least: characters, 
background, subjects, plots, conflicts, highlight, ending.
So, our language model is to generate based on the all these factors. 
Among all these factors to be chosen, some should gain from our users.

**Characters**: you should input their infos: Names, Ages, Jobs, ect.

**Backgrounds**: you should describe some basic scenes for the entire story.

**Conflicts**: you can define the most important conflicts between Characters.

**Subjects**: you can define a subjects that the story is to reflect.

**Endings**: Happy ending or Sad ending?
