package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getHomepage(c *gin.Context) {
	headMeta := Metas{
		Author:      "https://tradingviet.com",
		Title:       "TRADING VIỆT - CHUYÊN TRANG TIN TỨC TRADING TỔNG HỢP",
		Description: "Trading Việt – Chuyên trang tin tức trading tổng hợp Tin tức, nhận định, kiến thức, phân tích, đầu tư thị trường tài chính Forex, Vàng Phái sinh, Hàng hóa, Ngoại hối, Chứng khoán, Tiền điện tử Bitcoin",
		KeyWord:     "tin tức ngoại hối, thông tin ngoại hối, tin tức ngoại hối hôm nay, tin tức thị trường ngoại hối, tin ngoại hối, trang tin tức ngoại hối, nhận định thị trường forex, tin tức forex, tin tức forex hôm nay, tin forex, tin tức thị trường forex, tin tức quan trọng forex, giao dịch forex theo tin tức, trang tin tức forex, tin tức ngoại hối vàng, tin tức ngoại tệ, tin tức tài chính ngoại hối, nên chơi sàn forex nào, các sàn ngoại hối uy tín, trang forex uy tín, phân tích forex, phân tích kỹ thuật forex, phân tích cơ bản forex, phân tích forex hàng ngày, phân tích dự báo forex, traderviet, forex factory, forexfactory, investing, tin forex factory, lịch kinh tế forex, tin tức forex hàng ngày, tin tức forex mới nhất, tin tức forexfactory, tin tức tài chính forex, tin tức kinh tế forex, tin tức forex nhanh nhất, robot forex, bot forex, bot trade forex, robot trade forex, robot trade fx, bot trade fx, tạo bot trade forex, tạo bot trade fx",
		Url:         "https://tradingviet.com",
		Image:       "../dist/assets/logo/tv.png",
		Favicon:     "../dist/assets/favicon.ico",
		Owner:       1,
	}
	c.HTML(http.StatusOK, "index.html", gin.H{
		"head": headMeta,
	})
}
func getNews(c *gin.Context) {
	headMeta := Metas{
		Author:      "https://tradingviet.com",
		Title:       "TRADING VIỆT - CHUYÊN TRANG TIN TỨC TRADING TỔNG HỢP",
		Description: "Trading Việt – Chuyên trang tin tức trading tổng hợp Tin tức, nhận định, kiến thức, phân tích, đầu tư thị trường tài chính Forex, Vàng Phái sinh, Hàng hóa, Ngoại hối, Chứng khoán, Tiền điện tử Bitcoin",
		KeyWord:     "tin tức ngoại hối, thông tin ngoại hối, tin tức ngoại hối hôm nay, tin tức thị trường ngoại hối, tin ngoại hối, trang tin tức ngoại hối, nhận định thị trường forex, tin tức forex, tin tức forex hôm nay, tin forex, tin tức thị trường forex, tin tức quan trọng forex, giao dịch forex theo tin tức, trang tin tức forex, tin tức ngoại hối vàng, tin tức ngoại tệ, tin tức tài chính ngoại hối, nên chơi sàn forex nào, các sàn ngoại hối uy tín, trang forex uy tín, phân tích forex, phân tích kỹ thuật forex, phân tích cơ bản forex, phân tích forex hàng ngày, phân tích dự báo forex, traderviet, forex factory, forexfactory, investing, tin forex factory, lịch kinh tế forex, tin tức forex hàng ngày, tin tức forex mới nhất, tin tức forexfactory, tin tức tài chính forex, tin tức kinh tế forex, tin tức forex nhanh nhất, robot forex, bot forex, bot trade forex, robot trade forex, robot trade fx, bot trade fx, tạo bot trade forex, tạo bot trade fx",
		Url:         "https://tradingviet.com",
		Image:       "../dist/assets/logo/tv.png",
		Favicon:     "../dist/assets/favicon.ico",
		Owner:       1,
	}
	c.HTML(http.StatusOK, "detail.html", gin.H{
		"head": headMeta,
	})
}
