#selpg
输入文件为：in.txt <br>
内容为： <br>
test <br>
test <br>
test <br>
test <br>
test <br>
输出文件为：out.txt(初始为空) <br>
错误文件为：error.txt(初始为空) <br>
使用方法： <br>
        -s              Start from Page <number>. <br>
        -e              End to Page <number>. <br>
        -l              [options]Specify the number of line per page.Default is72.format:-l=number <br>
        -f              [options]Specify that the pages are sperated by \f. <br>
        [filename]      [options]Read input from the file. <br>
        If no file specified, selpg will read input from stdin. Control-D to end. <br>

测试:(按照使用selpg，测试的时候由于txt文件行数短，设置每页行数为2行，方便测试） <br>
测试 1:./selpg -s=1 -e=1 in.txt <br>
结果： <br>
test <br>
test <br>

测试2： ./selpg -s=1 -e=1 < in.txt <br>
结果：与测试1相同<br>

测试3：./selpg -s=1 -e=2  in.txt>out.txt <br>
结果：out.txt文件内容为： <br>
test <br>
test <br>
test <br>
test <br>

测试4： ./selpg -s=1 -e=2 in.txt 2> error.txt <br>
结果： <br>
test <br>
test <br>
test <br>
test <br>
error.txt 为空 <br>

测试5：./selpg -s=1 -e=2 in.txt >out.txt 2>error.txt<br>
结果：out.txt文件内容为：<br>
test<br>
test<br>
test<br>
test<br>

测试6：./selpg -s=1 -e=3 -l=1 in.txt<br>
结果:<br>
test<br>
test<br>
test<br>

测试7：



