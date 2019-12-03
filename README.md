# CourseAnalysis4CCNU

------

### 描述/Description

**CourseAnalysis4CCNU** 是基于**网络爬虫**, **关系型数据库**, **全文分词索引**的对于学校课程, 课堂的**检索**系统. 使用Go语言进行撰写, 仅使用系统包, 没有任何第三方依赖.

### 环境/Env

#### 系统环境/Operating System Env

```
Linux Ubuntu 18.04 TLS
```

#### 语言环境/Language Env

```
go version: 1.10+
```

#### 环境变量/Env Variable

```
COURSE_DB_USERNAME="YOUR DATABASE USERNAME"
COURSE_DB_NAME="YOUR DATABASE NAME"
COURSE_DB_HOST="YOUR DATABASE HOST"
COURSE_DB_PASSWORD="YOUR DATABASE PASSWORD"
```

### 运行/Run

```bash
$: make
$: ./analyst -h // 查看本应用help

CCNU Course And Class Analyst version: 0.1.0
Usage: ./analyst [-a analysis] [-d database] [-h help] [-s serve]

Options:
  -a	analysis xlsx
  -d	insert data to database
  -h	this help
  -s	serve apis at localhost
```