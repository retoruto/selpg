输入文件为：in.txt \n
内容为：
test
test
test
test
test
输出文件为：out.txt(初始为空)
错误文件为：error.txt(初始为空)
使用方法：
        -s              Start from Page <number>.
        -e              End to Page <number>.
        -l              [options]Specify the number of line per page.Default is72.format:-l=number
        -f              [options]Specify that the pages are sperated by \f.
        [filename]      [options]Read input from the file.
        If no file specified, selpg will read input from stdin. Control-D to end.

测试:(按照使用selpg，测试的时候由于txt文件行数短，设置每页行数为2行，方便测试）
测试 1:./selpg -s=1 -e=1 in.txt
结果：
test
test

测试2： ./selpg -s=1 -e=1 < in.txt
结果：与测试1相同

测试3：./selpg -s=1 -e=2  in.txt>out.txt
结果：out.txt文件内容为：
test
test
test
test

测试4： ./selpg -s=1 -e=2 in.txt 2> error.txt
结果：
test
test
test
test
error.txt为空

测试5：./selpg -s=1 -e=2 in.txt >out.txt 2>error.txt
结果：out.txt文件内容为：
test
test
test
test

测试6：./selpg -s=1 -e=3 -l=1 in.txt
结果:
test
test
test

测试7：



