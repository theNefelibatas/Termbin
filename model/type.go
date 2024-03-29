package model

var MIMEType map[string]string

func init() {
	MIMEType = make(map[string]string, 40)
	// https://developer.mozilla.org/zh-CN/docs/Web/HTTP/Basics_of_HTTP/MIME_types/Common_types
	MIMEType[".apng"] = "image/apng"
	MIMEType[".avi"] = "video/x-msvideo"
	MIMEType[".bmp"] = "image/bmp"
	MIMEType[".csv"] = "text/csv"
	MIMEType[".doc"] = "application/msword"
	MIMEType[".docx"] = "application/vnd.openxmlformats-officedocument.wordprocessingml.document"
	MIMEType[".epub"] = "application/epub+zip"
	MIMEType[".gif"] = "image/gif"
	MIMEType[".gz"] = "application/gzip"
	MIMEType[".htm"] = "text/html"
	MIMEType[".html"] = "text/html"
	MIMEType[".ico"] = "image/vnd.microsoft.icon"
	MIMEType[".jpeg"] = "image/jpeg"
	MIMEType[".jpg"] = "image/jpeg"
	MIMEType[".mpeg"] = "audio/mpeg"
	MIMEType[".mp3"] = "audio/mpeg"
	MIMEType[".mp4"] = "video/mp4"
	MIMEType[".pdf"] = "application/pdf"
	MIMEType[".png"] = "image/png"
	MIMEType[".ppt"] = "application/vnd.ms-powerpoint"
	MIMEType[".pptx"] = "application/vnd.openxmlformats-officedocument.presentationml.presentation"
	MIMEType[".rar"] = "application/vnd.rar"
	MIMEType[".svg"] = "image/svg+xml"
	MIMEType[".tif"] = "image/tiff"
	MIMEType[".tiff"] = "image/tiff"
	MIMEType[".txt"] = "text/plain"
	MIMEType[".wav"] = "audio/wav"
	MIMEType[".xls"] = "application/vnd.ms-excel"
	MIMEType[".xlsx"] = "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet"
	MIMEType[".zip"] = "application/zip"
}
