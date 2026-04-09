package main

var extCategoryRules = map[string]string{
	".jpg":  "images",
	".jpeg": "images",
	".png":  "images",
	".gif":  "images",
	".webp": "images",
	".svg":  "images",
	".bmp":  "images",
	".ico":  "images",

	".mp4":  "videos",
	".mkv":  "videos",
	".mov":  "videos",
	".avi":  "videos",
	".wmv":  "videos",
	".flv":  "videos",
	".webm": "videos",

	".mp3":  "audio",
	".wav":  "audio",
	".flac": "audio",
	".aac":  "audio",
	".ogg":  "audio",
	".m4a":  "audio",

	".pdf":  "documents",
	".doc":  "documents",
	".docx": "documents",
	".xls":  "documents",
	".xlsx": "documents",
	".ppt":  "documents",
	".pptx": "documents",
	".txt":  "documents",
	".md":   "documents",
	".csv":  "documents",

	".zip": "archives",
	".rar": "archives",
	".7z":  "archives",
	".tar": "archives",
	".gz":  "archives",
	".bz2": "archives",

	".exe": "programs",
	".msi": "programs",
	".dmg": "programs",
	".pkg": "programs",
	".deb": "programs",
	".rpm": "programs",

	".go":   "code",
	".js":   "code",
	".ts":   "code",
	".jsx":  "code",
	".tsx":  "code",
	".py":   "code",
	".java": "code",
	".c":    "code",
	".cpp":  "code",
	".h":    "code",
	".hpp":  "code",
	".json": "code",
	".xml":  "code",
	".yaml": "code",
	".yml":  "code",
	".html": "code",
	".css":  "code",
	".sh":   "code",
}

func CategoryByExt(ext string) string {
	if category, ok := extCategoryRules[ext]; ok {
		return category
	}
	return "others"
}
