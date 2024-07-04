# CLI commands

Авторизация
```
auth sign-in -u <login> -p <password>
```

Список пользователей
```
admin user-list -a true/false
```

Создать нового пользователя
```
admin user-create -u <login> -e <email> -p <password>
```

Активация пользователя
```
admin user-active -u <login>
```

Деактивация пользователя
```
admin user-inactive -u <login>
```

Список репозиториев
```
repository search
```

Список участников репозитория
```
repository collaborator-list -o <owner> -r <repository>
```

Добавить участника к репозиторию
```
repository collaborator-add -o <owner> -r <repository> -u <member> -p <permission>
```

Удалить участника из репозитория
```
repository collaborator-delete -o <owner> -r <repository> -u <member>
```
