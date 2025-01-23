# headerparser_usage

> https://github.com/google/syzkaller/blob/master/docs/headerparser_usage.md

**说明**

辅助解析头文件的工具

**使用**

```bash
$ python headerparser.py --filenames=./test_headers/th_b.h
B {
          B1     len|fileoff|flags|intN     #(unsigned long)
          B2     len|fileoff|flags|intN     #(unsigned long)
}
struct_containing_union {
          something          len|fileoff|flags|int32                   #(int)
          a_union.a_char     ptr[in|out, string]|ptr[in, filename]     #(char*)
          a_union.B_ptr      ptr|buffer|array                          #(struct B*)
}
```

**一些问题**

Let us try parsing `test_headers/th_a.h` header file to generate argument structs.

```
$ python headerparser.py --filenames=./test_headers/th_a.h
ERROR:root:HeaderFilePreprocessorException: /tmp/tmpW8xzty/source.o:36:2: before: some_type

$ python headerparser.py --filenames=./test_headers/th_a.h --debug
DEBUG:GlobalHierarchy:load_header_files : ['./test_headers/th_a.h']
DEBUG:HeaderFilePreprocessor:HeaderFilePreprocessor._mktempfiles: sourcefile=/tmp/tmpbBQYhR/source.cobjectfile=/tmp/tmpbBQYhR/source.o
DEBUG:HeaderFilePreprocessor:HeaderFilePreprocessor.execute: cp ./test_headers/th_a.h /tmp/tmpbBQYhR
DEBUG:HeaderFilePreprocessor:HeaderFilePreprocessor.execute: gcc -I. -E -P -c /tmp/tmpbBQYhR/source.c > /tmp/tmpbBQYhR/source.o
ERROR:root:HeaderFilePreprocessorException: /tmp/tmpbBQYhR/source.o:36:2: before: some_type
```

From the error message, we can see that the error occurs as pycparser is not aware of the type `some_type`. We can resolve this by making pycparser aware of the unknown type. In order to do this, we supply headerparser with a include file that contains C declarations and includes that can fix the parse error.

```
$ cat > include_file
typedef int some_type;
$ python headerparser.py --filenames=./test_headers/th_a.h --include=./include_file
A {
          B_item              ptr|buffer|array                          #(struct B*)
          char_ptr            ptr[in|out, string]|ptr[in, filename]     #(char*)
          an_unsigned_int     len|fileoff|int32                         #(unsigned int)
          a_bool              _Bool                                     #(_Bool)
          another_bool        _Bool                                     #(_Bool)
          var                 some_type                                 #(some_type)
}
```

