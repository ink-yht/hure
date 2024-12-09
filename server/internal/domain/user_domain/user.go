package user_domain

// GetSexText 响应给前端：在返回数据时，将数值转为文本
func GetSexText(sex int8) string {
	switch sex {
	case 1:
		return "男"
	case 2:
		return "女"
	default:
		return "未知"
	}
}

// GetSexValue 后端接受处理时解析为数值存储
func GetSexValue(sexText string) int8 {
	switch sexText {
	case "男":
		return 1
	case "女":
		return 2
	default:
		return 99 // 未知或未指定
	}
}
