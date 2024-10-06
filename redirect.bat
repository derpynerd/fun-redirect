@echo off
set host=http://localhost:8080
set var=%1

set var=%var:~5,-1%

set url=%host%/%var%

start "" %url%
