<!DOCTYPE html>
<html lang="zh-CN">
  <head>
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <title>Admin | {{.AppName}}</title>
    {{stylesheetTag "libs/bootstrap"}}
    {{stylesheetTag "admin"}}
  </head>
  <body>
    {{.LayoutContent}}
    {{javascriptTag "libs/jquery"}}
    {{javascriptTag "libs/jquery-ujs"}}
    {{javascriptTag "admin"}}
  </body>
</html>