@echo on
xcopy E:\net\test\dist\static\*.* H:\go\mygo\src\test\static /e /H /R /y
copy /y E:\net\test\dist\index.html H:\go\mygo\src\test\views\