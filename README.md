#### 简介
jm 方便json marshal的命令行工具

#####
``` console
Usage of ./jm:
  -a, --array
    	json array
  -c, --combination
    	Combine command line and file data
  -d, --delimiter string
    	use DELIM instead of TAB for field delimiter
  -data string[]
    	json data
  -o, --object
    	json object
```

#### 示例
* 把命令行里面的key val数组转成json object
```
echo -ne ""|./jm -o -data mode word appkey test-appkey scoreCoefficient 1 audioFormat mp3 eof guo-test-end|jq
{
  "appkey": "test-appkey",
  "audioFormat": "mp3",
  "eof": "guo-test-end",
  "mode": "word",
  "scoreCoefficient": "1"
}
```

* 把命令行里面的key val数组+管道前面的数据转成json object
```
echo -ne "displayText\thello world"|./jm -o -data mode word appkey test-appkey scoreCoefficient 1 audioFormat mp3 eof guo-test-end|jq

{
  "appkey": "test-appkey",
  "audioFormat": "mp3",
  "displayText": "hello world",
  "eof": "guo-test-end",
  "mode": "word",
  "scoreCoefficient": "1"
}

```
