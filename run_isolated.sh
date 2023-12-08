#!/bin/bash

# Проверяем, передан ли путь к исполняемому файлу
if [ "$#" -ne 1 ]; then
    echo "Использование: $0 <путь_к_программе>"
    exit 1
fi

executable=$1

# Запускаем программу в новом пространстве имен сети и PID
unshare --net --pid --fork --ipc --uts --mount-proc $executable
