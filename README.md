# 服务计算作业2——selpg
这次作业的内容主要是练习和掌握go语言的相关用法。
作业的思路大概就是把文档中所给出的selpg.c文件翻译成go语言的形式，然而TA建议不要亲自解析命令，提倡用flag。
我这里实现的还是纯翻译（亲自解析命令型），打印机的部分还是参考同学的做法。



-------------------

#测试

**linux环境**

初始化运行界面：
```
retoruto@ubuntu:~/Desktop/Selpg_2$ go build selpg.go
retoruto@ubuntu:~/Desktop/Selpg_2$ ./selpg
./selpg: not enough arguments
Usage of ./selpg:

        selpg -s<Number> -e<Number> [options] [filename]

        -s<Number>      Start from Page <number>.
        -e<Number>      End to Page <number>.
        -l<Number>      [options]Specify the number of line per page.Default is5.
        -f              [options]Specify that the pages are sperated by \f.
        [filename]      [options]Read input from the file.

If no file specified, ./selpg will read input from stdin. Control-D to end.

retoruto@ubuntu:~/Desktop/Selpg_2$
```

input_file.txt中的内容为：

```
1
2
3
4
5
6
7
8
9
10
11
12


```
这里先声明为12行数据，便于测试这里把默认的行数调为“5”

下面为普通指令的测试：

```
retoruto@ubuntu:~/Desktop/Selpg_2$ ./selpg -s1 -e1 input_file.txt
1
2
3
4
5
```

```
retoruto@ubuntu:~/Desktop/Selpg_2$ ./selpg -s1 -e1 < input_file.txt
1
2
3
4
5
```

```
retoruto@ubuntu:~/Desktop/Selpg_2$ cat input_file.txt | ./selpg -s1 -e1
1
2
3
4
5
```

```
./selpg -s1 -e1 input_file.txt >output_file.txt
```
output_file.txt写入成功：
![这里写图片描述](http://img.blog.csdn.net/20171016191322659?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvcXFfMzY4MTY5MTI=/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast)

在测试错误信息的重定向的时候，参考了一位同学的做法：
在progress_input()函数的最下方添加一个测试错误的信息（使用参数os.Stderr）
![这里写图片描述](http://img.blog.csdn.net/20171016192033574?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvcXFfMzY4MTY5MTI=/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast)
在进行测试：
未定向前：

```
retoruto@ubuntu:~/Desktop/Selpg_2$ ./selpg -s1 -e1 input_file.txt
1
2
3
4
5
test error
```
定向之后：

```
retoruto@ubuntu:~/Desktop/Selpg_2$ ./selpg -s1 -e1 input_file.txt 2>error.txt
1
2
3
4
5
```
在屏幕没有显示出错误信息，而在error.txt中出现了：
![这里写图片描述](http://img.blog.csdn.net/20171016192545963?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvcXFfMzY4MTY5MTI=/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast)

```
重定向：
retoruto@ubuntu:~/Desktop/Selpg_2$ ./selpg -s1 -e1 input_file.txt >output_file.txt 2>error.txt
丢弃错误信息：
retoruto@ubuntu:~/Desktop/Selpg_2$ ./selpg -s1 -e1 input_file.txt >output_file.txt 2>/dev/null
输出被丢弃，输出错误信息：
retoruto@ubuntu:~/Desktop/Selpg_2$ ./selpg -s1 -e1 input_file.txt >/dev/null
test error
%!(EXTRA string=./selpg)
```

测试另一个 selpg -s1 -e1 input_file | other_command
```
retoruto@ubuntu:~/Desktop/Selpg_2$ ./selpg -s1 -e1 input_file.txt | wc
test error
%!(EXTRA string=./selpg)      5       5      10

重定向错误信息的输出

retoruto@ubuntu:~/Desktop/Selpg_2$ ./selpg -s1 -e1 input_file.txt 2>error.txt | wc
      5       5      10
```
测试行数设置：

```
未设置行数之前：
retoruto@ubuntu:~/Desktop/Selpg_2$ ./selpg -s1 -e2 input_file.txt
1
2
3
4
5
6
7
8
9
10
设置行数之后：
retoruto@ubuntu:~/Desktop/Selpg_2$ ./selpg -s1 -e2 -l1 input_file.txt
1
2
```
测试-f的用法：
这里用程序来写入一个带'\f'的文件：

![这里写图片描述](http://img.blog.csdn.net/20171016200113652?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvcXFfMzY4MTY5MTI=/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast)
```
retoruto@ubuntu:~/Desktop/Selpg_2$ ./selpg -s1 -e2 -f input_file.txt
test
1 page
      test
2 page
retoruto@ubuntu:~/Desktop/Selpg_2$ ./selpg -s1 -e1 -f input_file.txt
test
1 page
```
测试成功。

打印机功能的老师说暂时不实现，要实现的可以去参考第一位同学的代码。
我这里参考的同学的是通过自己写一个手动管道来实现的。

测试后台运行（参考了ps指令）

```
retoruto@ubuntu:~/Desktop/Selpg_2$ ./selpg -s1 -e2 input_file.txt > output_file.txt 2>error.txt &
[1] 13275
retoruto@ubuntu:~/Desktop/Selpg_2$ ps
   PID TTY          TIME CMD
 12985 pts/12   00:00:00 bash
 13289 pts/12   00:00:00 ps
[1]+  Done                    ./selpg -s1 -e2 input_file.txt > output_file.txt 2> error.txt
```
成功在后台运行