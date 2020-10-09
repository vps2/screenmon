# Screenmon 

**screenmon** - утилита для отслеживания изменений на экране и звуковом оповещении об этом событии.

```sh
screenmon.exe -h
Usage of screenmon.exe:
  -d, --display int        number of the display to track changes on. It may not match the number in the system settings. (default 1)
  -t, --timeout duration   timeout between screen capture (default 15s)
```

### Пример запуска:

```sh
screenmon.exe -d 2 -t 30s
```