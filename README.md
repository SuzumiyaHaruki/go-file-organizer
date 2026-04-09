# go-file-organizer

一个用 Go 编写的文件整理工具。

## 功能
- 扫描指定目录
- 按文件后缀分类
- 默认 dry-run 预览
- 使用 `--apply` 真正执行移动
- 自动处理同名冲突

## 使用方法

预览：
```bash
go run . --dir ./testdata
```

真正执行：
```bash
go run . --dir ./testdata --apply