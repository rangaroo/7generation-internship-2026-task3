# Задача по разработке номер 3.

## Условие задачи

На вход программе передаётся текстовый файл (в виде аргумента командной строки).
Ваша программа должна прочитать файл и вывести 20 наиболее часто встречающихся 
слов в файле в порядке убывания частоты, вместе с количеством их вхождений.

Результат должен полностью совпадать с выводом следующего bash-скрипта:
#!/usr/bin/env bash
```
cat $1 | tr -cs 'a-zA-Z' '[\n*]' | grep -v "^$" | tr '[:upper:]' '[:lower:]'| sort | uniq -c | sort -nr | head -20
```

Ваша программа должна корректно обрабатывать бинарные файлы (например, /boot/vmlinuz) и не падать.

Языки и инструменты:
Решение должно быть написано целиком на Go.
Использование стандартных контейнерных библиотек языка запрещено (Go maps ЗАПРЕЩЕНЫ, slices в Go разрешены). Использование I/O-потоков разрешено,
использование строк Go (string) ЗАПРЕЩЕНО.
Эталонная платформа для задания - Linux 64-bit. Если у вас нет к ней доступа, вы можете
реализовать решение в среде Linux или Windows по своему выбору, однако Вы должны
быть готовы объяснить, какие изменения потребуются, чтобы запустить решение на
эталонной платформе.

## Usage

### Building the Program
```bash
go build -o wordcount main.go
```

### Running the Program
```bash
./wordcount <filename>
```

Example:
```bash
./wordcount test.txt
```

### Running Tests
```bash
go test -v
```

### Comparing with Reference Bash Script
```bash
# Run our solution
./wordcount test.txt > our_output.txt

# Run bash reference
cat test.txt | tr -cs 'a-zA-Z' '[\n*]' | grep -v "^$" | tr '[:upper:]' '[:lower:]' | sort | uniq -c | sort -nr | head -20 > bash_output.txt

# Compare outputs
diff our_output.txt bash_output.txt
```

