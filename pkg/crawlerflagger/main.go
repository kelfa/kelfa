// DO NOT EDIT: Auto generated
package crawlerflagger

type Crawler struct {
	Pattern      string
	Instances    []string
	URL          *string
	Description  *string
	AdditionDate *string
	DependsOn    []string
}

// ExactMatch allows you to identify if and which kind of crawler a User Agent belongs to
// The data are taken from the https://github.com/monperrus/crawler-user-agents/ project
func ExactMatch(ua string) *Crawler {
	switch ua {

	case "Googlebot/2.1 (+http://www.google.com/bot.html)",
        "Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)",
        "Mozilla/5.0 (iPhone; CPU iPhone OS 6_0 like Mac OS X) AppleWebKit/536.26 (KHTML, like Gecko) Version/6.0 Mobile/10A5376e Safari/8536.25 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)",
        "Mozilla/5.0 (iPhone; CPU iPhone OS 8_3 like Mac OS X) AppleWebKit/537.36 (KHTML, like Gecko) Version/8.0 Mobile/12F70 Safari/600.1.4 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)",
        "Mozilla/5.0 (iPhone; CPU iPhone OS 8_3 like Mac OS X) AppleWebKit/600.1.4 (KHTML, like Gecko) Version/8.0 Mobile/12F70 Safari/600.1.4 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)",
        "Mozilla/5.0 (Linux; Android 6.0.1; Nexus 5X Build/MMB29P) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/27.0.1453 Mobile Safari/537.36 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)",
        "Mozilla/5.0 (Linux; Android 6.0.1; Nexus 5X Build/MMB29P) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/41.0.2272.96 Mobile Safari/537.36 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)",
        "Mozilla/5.0 AppleWebKit/537.36 (KHTML, like Gecko; compatible; Googlebot/2.1; +http://www.google.com/bot.html) Safari/537.36":
		return &Crawler{
			Pattern:      `Googlebot\/`,
			Instances:    []string{ "Googlebot/2.1 (+http://www.google.com/bot.html)", "Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)", "Mozilla/5.0 (iPhone; CPU iPhone OS 6_0 like Mac OS X) AppleWebKit/536.26 (KHTML, like Gecko) Version/6.0 Mobile/10A5376e Safari/8536.25 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)", "Mozilla/5.0 (iPhone; CPU iPhone OS 8_3 like Mac OS X) AppleWebKit/537.36 (KHTML, like Gecko) Version/8.0 Mobile/12F70 Safari/600.1.4 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)", "Mozilla/5.0 (iPhone; CPU iPhone OS 8_3 like Mac OS X) AppleWebKit/600.1.4 (KHTML, like Gecko) Version/8.0 Mobile/12F70 Safari/600.1.4 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)", "Mozilla/5.0 (Linux; Android 6.0.1; Nexus 5X Build/MMB29P) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/27.0.1453 Mobile Safari/537.36 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)", "Mozilla/5.0 (Linux; Android 6.0.1; Nexus 5X Build/MMB29P) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/41.0.2272.96 Mobile Safari/537.36 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)", "Mozilla/5.0 AppleWebKit/537.36 (KHTML, like Gecko; compatible; Googlebot/2.1; +http://www.google.com/bot.html) Safari/537.36" },
			URL:          stringer("http://www.google.com/bot.html"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "DoCoMo/2.0 N905i(c100;TB;W24H16) (compatible; Googlebot-Mobile/2.1; +http://www.google.com/bot.html)",
        "Mozilla/5.0 (iPhone; CPU iPhone OS 6_0 like Mac OS X) AppleWebKit/536.26 (KHTML, like Gecko) Version/6.0 Mobile/10A5376e Safari/8536.25 (compatible; Googlebot-Mobile/2.1; +http://www.google.com/bot.html)",
        "Mozilla/5.0 (iPhone; U; CPU iPhone OS 4_1 like Mac OS X; en-us) AppleWebKit/532.9 (KHTML, like Gecko) Version/4.0.5 Mobile/8B117 Safari/6531.22.7 (compatible; Googlebot-Mobile/2.1; +http://www.google.com/bot.html)",
        "Nokia6820/2.0 (4.83) Profile/MIDP-1.0 Configuration/CLDC-1.0 (compatible; Googlebot-Mobile/2.1; +http://www.google.com/bot.html)",
        "SAMSUNG-SGH-E250/1.0 Profile/MIDP-2.0 Configuration/CLDC-1.1 UP.Browser/6.2.3.3.c.1.101 (GUI) MMP/2.0 (compatible; Googlebot-Mobile/2.1; +http://www.google.com/bot.html)":
		return &Crawler{
			Pattern:      `Googlebot-Mobile`,
			Instances:    []string{ "DoCoMo/2.0 N905i(c100;TB;W24H16) (compatible; Googlebot-Mobile/2.1; +http://www.google.com/bot.html)", "Mozilla/5.0 (iPhone; CPU iPhone OS 6_0 like Mac OS X) AppleWebKit/536.26 (KHTML, like Gecko) Version/6.0 Mobile/10A5376e Safari/8536.25 (compatible; Googlebot-Mobile/2.1; +http://www.google.com/bot.html)", "Mozilla/5.0 (iPhone; U; CPU iPhone OS 4_1 like Mac OS X; en-us) AppleWebKit/532.9 (KHTML, like Gecko) Version/4.0.5 Mobile/8B117 Safari/6531.22.7 (compatible; Googlebot-Mobile/2.1; +http://www.google.com/bot.html)", "Nokia6820/2.0 (4.83) Profile/MIDP-1.0 Configuration/CLDC-1.0 (compatible; Googlebot-Mobile/2.1; +http://www.google.com/bot.html)", "SAMSUNG-SGH-E250/1.0 Profile/MIDP-2.0 Configuration/CLDC-1.1 UP.Browser/6.2.3.3.c.1.101 (GUI) MMP/2.0 (compatible; Googlebot-Mobile/2.1; +http://www.google.com/bot.html)" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Googlebot-Image/1.0":
		return &Crawler{
			Pattern:      `Googlebot-Image`,
			Instances:    []string{ "Googlebot-Image/1.0" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Googlebot-News":
		return &Crawler{
			Pattern:      `Googlebot-News`,
			Instances:    []string{ "Googlebot-News" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Googlebot-Video/1.0":
		return &Crawler{
			Pattern:      `Googlebot-Video`,
			Instances:    []string{ "Googlebot-Video/1.0" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "AdsBot-Google (+http://www.google.com/adsbot.html)":
		return &Crawler{
			Pattern:      `AdsBot-Google([^-]|$)`,
			Instances:    []string{ "AdsBot-Google (+http://www.google.com/adsbot.html)" },
			URL:          stringer("https://support.google.com/webmasters/answer/1061943?hl=en"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "AdsBot-Google-Mobile-Apps",
        "Mozilla/5.0 (Linux; Android 5.0; SM-G920A) AppleWebKit (KHTML, like Gecko) Chrome Mobile Safari (compatible; AdsBot-Google-Mobile; +http://www.google.com/mobile/adsbot.html)",
        "Mozilla/5.0 (iPhone; CPU iPhone OS 9_1 like Mac OS X) AppleWebKit/601.1.46 (KHTML, like Gecko) Version/9.0 Mobile/13B143 Safari/601.1 (compatible; AdsBot-Google-Mobile; +http://www.google.com/mobile/adsbot.html)":
		return &Crawler{
			Pattern:      `AdsBot-Google-Mobile`,
			Instances:    []string{ "AdsBot-Google-Mobile-Apps", "Mozilla/5.0 (Linux; Android 5.0; SM-G920A) AppleWebKit (KHTML, like Gecko) Chrome Mobile Safari (compatible; AdsBot-Google-Mobile; +http://www.google.com/mobile/adsbot.html)", "Mozilla/5.0 (iPhone; CPU iPhone OS 9_1 like Mac OS X) AppleWebKit/601.1.46 (KHTML, like Gecko) Version/9.0 Mobile/13B143 Safari/601.1 (compatible; AdsBot-Google-Mobile; +http://www.google.com/mobile/adsbot.html)" },
			URL:          stringer("https://support.google.com/adwords/answer/2404197"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Feedfetcher-Google; (+http://www.google.com/feedfetcher.html; 1 subscribers; feed-id=728742641706423)":
		return &Crawler{
			Pattern:      `Feedfetcher-Google`,
			Instances:    []string{ "Feedfetcher-Google; (+http://www.google.com/feedfetcher.html; 1 subscribers; feed-id=728742641706423)" },
			URL:          stringer("https://support.google.com/webmasters/answer/178852"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Mediapartners-Google",
        "Mozilla/5.0 (compatible; MSIE or Firefox mutant; not on Windows server;) Daumoa/4.0 (Following Mediapartners-Google)",
        "Mozilla/5.0 (iPhone; U; CPU iPhone OS 10_0 like Mac OS X; en-us) AppleWebKit/602.1.38 (KHTML, like Gecko) Version/10.0 Mobile/14A5297c Safari/602.1 (compatible; Mediapartners-Google/2.1; +http://www.google.com/bot.html)",
        "Mozilla/5.0 (iPhone; U; CPU iPhone OS 4_1 like Mac OS X; en-us) AppleWebKit/532.9 (KHTML, like Gecko) Version/4.0.5 Mobile/8B117 Safari/6531.22.7 (compatible; Mediapartners-Google/2.1; +http://www.google.com/bot.html)":
		return &Crawler{
			Pattern:      `Mediapartners-Google`,
			Instances:    []string{ "Mediapartners-Google", "Mozilla/5.0 (compatible; MSIE or Firefox mutant; not on Windows server;) Daumoa/4.0 (Following Mediapartners-Google)", "Mozilla/5.0 (iPhone; U; CPU iPhone OS 10_0 like Mac OS X; en-us) AppleWebKit/602.1.38 (KHTML, like Gecko) Version/10.0 Mobile/14A5297c Safari/602.1 (compatible; Mediapartners-Google/2.1; +http://www.google.com/bot.html)", "Mozilla/5.0 (iPhone; U; CPU iPhone OS 4_1 like Mac OS X; en-us) AppleWebKit/532.9 (KHTML, like Gecko) Version/4.0.5 Mobile/8B117 Safari/6531.22.7 (compatible; Mediapartners-Google/2.1; +http://www.google.com/bot.html)" },
			URL:          stringer("https://support.google.com/webmasters/answer/1061943?hl=en"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "APIs-Google (+https://developers.google.com/webmasters/APIs-Google.html)":
		return &Crawler{
			Pattern:      `APIs-Google`,
			Instances:    []string{ "APIs-Google (+https://developers.google.com/webmasters/APIs-Google.html)" },
			URL:          stringer("https://support.google.com/webmasters/answer/1061943?hl=en"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Mozilla/5.0 (Windows Phone 8.1; ARM; Trident/7.0; Touch; rv:11.0; IEMobile/11.0; NOKIA; Lumia 530) like Gecko (compatible; adidxbot/2.0; +http://www.bing.com/bingbot.htm)",
        "Mozilla/5.0 (compatible; adidxbot/2.0;  http://www.bing.com/bingbot.htm)",
        "Mozilla/5.0 (compatible; adidxbot/2.0; +http://www.bing.com/bingbot.htm)",
        "Mozilla/5.0 (compatible; bingbot/2.0;  http://www.bing.com/bingbot.htm)",
        "Mozilla/5.0 (compatible; bingbot/2.0; +http://www.bing.com/bingbot.htm",
        "Mozilla/5.0 (compatible; bingbot/2.0; +http://www.bing.com/bingbot.htm)",
        "Mozilla/5.0 (compatible; bingbot/2.0; +http://www.bing.com/bingbot.htm) SitemapProbe",
        "Mozilla/5.0 (iPhone; CPU iPhone OS 7_0 like Mac OS X) AppleWebKit/537.51.1 (KHTML, like Gecko) Version/7.0 Mobile/11A465 Safari/9537.53 (compatible; adidxbot/2.0;  http://www.bing.com/bingbot.htm)",
        "Mozilla/5.0 (iPhone; CPU iPhone OS 7_0 like Mac OS X) AppleWebKit/537.51.1 (KHTML, like Gecko) Version/7.0 Mobile/11A465 Safari/9537.53 (compatible; adidxbot/2.0; +http://www.bing.com/bingbot.htm)",
        "Mozilla/5.0 (iPhone; CPU iPhone OS 7_0 like Mac OS X) AppleWebKit/537.51.1 (KHTML, like Gecko) Version/7.0 Mobile/11A465 Safari/9537.53 (compatible; bingbot/2.0;  http://www.bing.com/bingbot.htm)",
        "Mozilla/5.0 (iPhone; CPU iPhone OS 7_0 like Mac OS X) AppleWebKit/537.51.1 (KHTML, like Gecko) Version/7.0 Mobile/11A465 Safari/9537.53 (compatible; bingbot/2.0; +http://www.bing.com/bingbot.htm)",
        "Mozilla/5.0 (seoanalyzer; compatible; bingbot/2.0; +http://www.bing.com/bingbot.htm)":
		return &Crawler{
			Pattern:      `bingbot`,
			Instances:    []string{ "Mozilla/5.0 (Windows Phone 8.1; ARM; Trident/7.0; Touch; rv:11.0; IEMobile/11.0; NOKIA; Lumia 530) like Gecko (compatible; adidxbot/2.0; +http://www.bing.com/bingbot.htm)", "Mozilla/5.0 (compatible; adidxbot/2.0;  http://www.bing.com/bingbot.htm)", "Mozilla/5.0 (compatible; adidxbot/2.0; +http://www.bing.com/bingbot.htm)", "Mozilla/5.0 (compatible; bingbot/2.0;  http://www.bing.com/bingbot.htm)", "Mozilla/5.0 (compatible; bingbot/2.0; +http://www.bing.com/bingbot.htm", "Mozilla/5.0 (compatible; bingbot/2.0; +http://www.bing.com/bingbot.htm)", "Mozilla/5.0 (compatible; bingbot/2.0; +http://www.bing.com/bingbot.htm) SitemapProbe", "Mozilla/5.0 (iPhone; CPU iPhone OS 7_0 like Mac OS X) AppleWebKit/537.51.1 (KHTML, like Gecko) Version/7.0 Mobile/11A465 Safari/9537.53 (compatible; adidxbot/2.0;  http://www.bing.com/bingbot.htm)", "Mozilla/5.0 (iPhone; CPU iPhone OS 7_0 like Mac OS X) AppleWebKit/537.51.1 (KHTML, like Gecko) Version/7.0 Mobile/11A465 Safari/9537.53 (compatible; adidxbot/2.0; +http://www.bing.com/bingbot.htm)", "Mozilla/5.0 (iPhone; CPU iPhone OS 7_0 like Mac OS X) AppleWebKit/537.51.1 (KHTML, like Gecko) Version/7.0 Mobile/11A465 Safari/9537.53 (compatible; bingbot/2.0;  http://www.bing.com/bingbot.htm)", "Mozilla/5.0 (iPhone; CPU iPhone OS 7_0 like Mac OS X) AppleWebKit/537.51.1 (KHTML, like Gecko) Version/7.0 Mobile/11A465 Safari/9537.53 (compatible; bingbot/2.0; +http://www.bing.com/bingbot.htm)", "Mozilla/5.0 (seoanalyzer; compatible; bingbot/2.0; +http://www.bing.com/bingbot.htm)" },
			URL:          stringer("http://www.bing.com/bingbot.htm"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Mozilla/5.0 (compatible; Yahoo! Slurp/3.0; http://help.yahoo.com/help/us/ysearch/slurp)",
        "Mozilla/5.0 (compatible; Yahoo! Slurp; http://help.yahoo.com/help/us/ysearch/slurp)",
        "Mozilla/5.0 (compatible; Yahoo! Slurp China; http://misc.yahoo.com.cn/help.html)":
		return &Crawler{
			Pattern:      `Slurp`,
			Instances:    []string{ "Mozilla/5.0 (compatible; Yahoo! Slurp/3.0; http://help.yahoo.com/help/us/ysearch/slurp)", "Mozilla/5.0 (compatible; Yahoo! Slurp; http://help.yahoo.com/help/us/ysearch/slurp)", "Mozilla/5.0 (compatible; Yahoo! Slurp China; http://misc.yahoo.com.cn/help.html)" },
			URL:          stringer("http://help.yahoo.com/help/us/ysearch/slurp"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "WGETbot/1.0 (+http://wget.alanreed.org)",
        "Wget/1.14 (linux-gnu)":
		return &Crawler{
			Pattern:      `[wW]get`,
			Instances:    []string{ "WGETbot/1.0 (+http://wget.alanreed.org)", "Wget/1.14 (linux-gnu)" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "eCairn-Grabber/1.0 (+http://ecairn.com/grabber) curl/7.15":
		return &Crawler{
			Pattern:      `curl`,
			Instances:    []string{ "eCairn-Grabber/1.0 (+http://ecairn.com/grabber) curl/7.15" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "LinkedInBot/1.0 (compatible; Mozilla/5.0; Jakarta Commons-HttpClient/3.1 +http://www.linkedin.com)",
        "LinkedInBot/1.0 (compatible; Mozilla/5.0; Jakarta Commons-HttpClient/4.3 +http://www.linkedin.com)":
		return &Crawler{
			Pattern:      `LinkedInBot`,
			Instances:    []string{ "LinkedInBot/1.0 (compatible; Mozilla/5.0; Jakarta Commons-HttpClient/3.1 +http://www.linkedin.com)", "LinkedInBot/1.0 (compatible; Mozilla/5.0; Jakarta Commons-HttpClient/4.3 +http://www.linkedin.com)" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Python-urllib/2.5",
        "Python-urllib/2.6",
        "Python-urllib/2.7",
        "Python-urllib/3.1",
        "Python-urllib/3.2",
        "Python-urllib/3.3",
        "Python-urllib/3.4",
        "Python-urllib/3.5",
        "Python-urllib/3.6":
		return &Crawler{
			Pattern:      `Python-urllib`,
			Instances:    []string{ "Python-urllib/2.5", "Python-urllib/2.6", "Python-urllib/2.7", "Python-urllib/3.1", "Python-urllib/3.2", "Python-urllib/3.3", "Python-urllib/3.4", "Python-urllib/3.5", "Python-urllib/3.6" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "python-requests/2.18.4":
		return &Crawler{
			Pattern:      `python-requests`,
			Instances:    []string{ "python-requests/2.18.4" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "2Bone_LinkChecker/1.0 libwww-perl/6.03",
        "2Bone_LinkChkr/1.0 libwww-perl/6.03",
        "amibot - http://www.amidalla.de - tech@amidalla.com libwww-perl/5.831":
		return &Crawler{
			Pattern:      `libwww-perl`,
			Instances:    []string{ "2Bone_LinkChecker/1.0 libwww-perl/6.03", "2Bone_LinkChkr/1.0 libwww-perl/6.03", "amibot - http://www.amidalla.de - tech@amidalla.com libwww-perl/5.831" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "httpunit/1.x":
		return &Crawler{
			Pattern:      `httpunit`,
			Instances:    []string{ "httpunit/1.x" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "NutchCVS/0.7.1 (Nutch; http://lucene.apache.org/nutch/bot.html; nutch-agent@lucene.apache.org)",
        "istellabot-nutch/Nutch-1.10":
		return &Crawler{
			Pattern:      `nutch`,
			Instances:    []string{ "NutchCVS/0.7.1 (Nutch; http://lucene.apache.org/nutch/bot.html; nutch-agent@lucene.apache.org)", "istellabot-nutch/Nutch-1.10" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Go-http-client/1.1":
		return &Crawler{
			Pattern:      `Go-http-client`,
			Instances:    []string{ "Go-http-client/1.1" },
			URL:          stringer("https://golang.org/pkg/net/http/"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "phpcrawl":
		return &Crawler{
			Pattern:      `phpcrawl`,
			Instances:    []string{ "phpcrawl" },
			URL:          stringer("http://phpcrawl.cuab.de/"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "adidxbot/1.1 (+http://search.msn.com/msnbot.htm)",
        "adidxbot/2.0 (+http://search.msn.com/msnbot.htm)",
        "librabot/1.0 (+http://search.msn.com/msnbot.htm)",
        "librabot/2.0 (+http://search.msn.com/msnbot.htm)",
        "msnbot-NewsBlogs/2.0b (+http://search.msn.com/msnbot.htm)",
        "msnbot-UDiscovery/2.0b (+http://search.msn.com/msnbot.htm)",
        "msnbot-media/1.0 (+http://search.msn.com/msnbot.htm)",
        "msnbot-media/1.1 (+http://search.msn.com/msnbot.htm)",
        "msnbot-media/2.0b (+http://search.msn.com/msnbot.htm)",
        "msnbot/1.0 (+http://search.msn.com/msnbot.htm)",
        "msnbot/1.1 (+http://search.msn.com/msnbot.htm)",
        "msnbot/2.0b (+http://search.msn.com/msnbot.htm)",
        "msnbot/2.0b (+http://search.msn.com/msnbot.htm).",
        "msnbot/2.0b (+http://search.msn.com/msnbot.htm)._":
		return &Crawler{
			Pattern:      `msnbot`,
			Instances:    []string{ "adidxbot/1.1 (+http://search.msn.com/msnbot.htm)", "adidxbot/2.0 (+http://search.msn.com/msnbot.htm)", "librabot/1.0 (+http://search.msn.com/msnbot.htm)", "librabot/2.0 (+http://search.msn.com/msnbot.htm)", "msnbot-NewsBlogs/2.0b (+http://search.msn.com/msnbot.htm)", "msnbot-UDiscovery/2.0b (+http://search.msn.com/msnbot.htm)", "msnbot-media/1.0 (+http://search.msn.com/msnbot.htm)", "msnbot-media/1.1 (+http://search.msn.com/msnbot.htm)", "msnbot-media/2.0b (+http://search.msn.com/msnbot.htm)", "msnbot/1.0 (+http://search.msn.com/msnbot.htm)", "msnbot/1.1 (+http://search.msn.com/msnbot.htm)", "msnbot/2.0b (+http://search.msn.com/msnbot.htm)", "msnbot/2.0b (+http://search.msn.com/msnbot.htm).", "msnbot/2.0b (+http://search.msn.com/msnbot.htm)._" },
			URL:          stringer("http://search.msn.com/msnbot.htm"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "FAST-WebCrawler/3.6/FirstPage (atw-crawler at fast dot no;http://fast.no/support/crawler.asp)",
        "FAST-WebCrawler/3.7 (atw-crawler at fast dot no; http://fast.no/support/crawler.asp)",
        "FAST-WebCrawler/3.7/FirstPage (atw-crawler at fast dot no;http://fast.no/support/crawler.asp)",
        "FAST-WebCrawler/3.8":
		return &Crawler{
			Pattern:      `FAST-WebCrawler`,
			Instances:    []string{ "FAST-WebCrawler/3.6/FirstPage (atw-crawler at fast dot no;http://fast.no/support/crawler.asp)", "FAST-WebCrawler/3.7 (atw-crawler at fast dot no; http://fast.no/support/crawler.asp)", "FAST-WebCrawler/3.7/FirstPage (atw-crawler at fast dot no;http://fast.no/support/crawler.asp)", "FAST-WebCrawler/3.8" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "FAST Enterprise Crawler 6 / Scirus scirus-crawler@fast.no; http://www.scirus.com/srsapp/contactus/",
        "FAST Enterprise Crawler 6 used by Schibsted (webcrawl@schibstedsok.no)":
		return &Crawler{
			Pattern:      `FAST Enterprise Crawler`,
			Instances:    []string{ "FAST Enterprise Crawler 6 / Scirus scirus-crawler@fast.no; http://www.scirus.com/srsapp/contactus/", "FAST Enterprise Crawler 6 used by Schibsted (webcrawl@schibstedsok.no)" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "BIGLOTRON (Beta 2;GNU/Linux)":
		return &Crawler{
			Pattern:      `BIGLOTRON`,
			Instances:    []string{ "BIGLOTRON (Beta 2;GNU/Linux)" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Mozilla/2.0 (compatible; Ask Jeeves/Teoma; +http://sp.ask.com/docs/about/tech_crawling.html)",
        "Mozilla/2.0 (compatible; Ask Jeeves/Teoma; +http://about.ask.com/en/docs/about/webmasters.shtml)":
		return &Crawler{
			Pattern:      `Teoma`,
			Instances:    []string{ "Mozilla/2.0 (compatible; Ask Jeeves/Teoma; +http://sp.ask.com/docs/about/tech_crawling.html)", "Mozilla/2.0 (compatible; Ask Jeeves/Teoma; +http://about.ask.com/en/docs/about/webmasters.shtml)" },
			URL:          stringer("http://about.ask.com/en/docs/about/webmasters.shtml"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "ConveraCrawler/0.9e (+http://ews.converasearch.com/crawl.htm)":
		return &Crawler{
			Pattern:      `convera`,
			Instances:    []string{ "ConveraCrawler/0.9e (+http://ews.converasearch.com/crawl.htm)" },
			URL:          stringer("http://ews.converasearch.com/crawl.htm"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Seekbot/1.0 (http://www.seekbot.net/bot.html) RobotsTxtFetcher/1.2":
		return &Crawler{
			Pattern:      `seekbot`,
			Instances:    []string{ "Seekbot/1.0 (http://www.seekbot.net/bot.html) RobotsTxtFetcher/1.2" },
			URL:          stringer("http://www.seekbot.net/bot.html"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Gigabot/1.0",
        "Gigabot/2.0 (http://www.gigablast.com/spider.html)":
		return &Crawler{
			Pattern:      `Gigabot`,
			Instances:    []string{ "Gigabot/1.0", "Gigabot/2.0 (http://www.gigablast.com/spider.html)" },
			URL:          stringer("http://www.gigablast.com/spider.html"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "GigablastOpenSource/1.0":
		return &Crawler{
			Pattern:      `Gigablast`,
			Instances:    []string{ "GigablastOpenSource/1.0" },
			URL:          stringer("https://github.com/gigablast/open-source-search-engine"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Mozilla/5.0 (compatible; Alexabot/1.0; +http://www.alexa.com/help/certifyscan; certifyscan@alexa.com)",
        "Mozilla/5.0 (compatible; Exabot PyExalead/3.0; +http://www.exabot.com/go/robot)",
        "Mozilla/5.0 (compatible; Exabot-Images/3.0; +http://www.exabot.com/go/robot)",
        "Mozilla/5.0 (compatible; Exabot/3.0 (BiggerBetter); +http://www.exabot.com/go/robot)",
        "Mozilla/5.0 (compatible; Exabot/3.0; +http://www.exabot.com/go/robot)",
        "Mozilla/5.0 (compatible; Exabot/3.0;  http://www.exabot.com/go/robot)":
		return &Crawler{
			Pattern:      `exabot`,
			Instances:    []string{ "Mozilla/5.0 (compatible; Alexabot/1.0; +http://www.alexa.com/help/certifyscan; certifyscan@alexa.com)", "Mozilla/5.0 (compatible; Exabot PyExalead/3.0; +http://www.exabot.com/go/robot)", "Mozilla/5.0 (compatible; Exabot-Images/3.0; +http://www.exabot.com/go/robot)", "Mozilla/5.0 (compatible; Exabot/3.0 (BiggerBetter); +http://www.exabot.com/go/robot)", "Mozilla/5.0 (compatible; Exabot/3.0; +http://www.exabot.com/go/robot)", "Mozilla/5.0 (compatible; Exabot/3.0;  http://www.exabot.com/go/robot)" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "ia_archiver (+http://www.alexa.com/site/help/webmasters; crawler@alexa.com)",
        "ia_archiver-web.archive.org":
		return &Crawler{
			Pattern:      `ia_archiver`,
			Instances:    []string{ "ia_archiver (+http://www.alexa.com/site/help/webmasters; crawler@alexa.com)", "ia_archiver-web.archive.org" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "GingerCrawler/1.0 (Language Assistant for Dyslexics; www.gingersoftware.com/crawler_agent.htm; support at ginger software dot com)":
		return &Crawler{
			Pattern:      `GingerCrawler`,
			Instances:    []string{ "GingerCrawler/1.0 (Language Assistant for Dyslexics; www.gingersoftware.com/crawler_agent.htm; support at ginger software dot com)" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Mozilla/4.5 (compatible; HTTrack 3.0x; Windows 98)":
		return &Crawler{
			Pattern:      `HTTrack`,
			Instances:    []string{ "Mozilla/4.5 (compatible; HTTrack 3.0x; Windows 98)" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Mozilla/4.0 (compatible; grub-client-0.3.0; Crawl your own stuff with http://grub.org)",
        "Mozilla/4.0 (compatible; grub-client-1.0.4; Crawl your own stuff with http://grub.org)",
        "Mozilla/4.0 (compatible; grub-client-1.0.5; Crawl your own stuff with http://grub.org)",
        "Mozilla/4.0 (compatible; grub-client-1.0.6; Crawl your own stuff with http://grub.org)",
        "Mozilla/4.0 (compatible; grub-client-1.0.7; Crawl your own stuff with http://grub.org)",
        "Mozilla/4.0 (compatible; grub-client-1.1.1; Crawl your own stuff with http://grub.org)",
        "Mozilla/4.0 (compatible; grub-client-1.2.1; Crawl your own stuff with http://grub.org)",
        "Mozilla/4.0 (compatible; grub-client-1.3.1; Crawl your own stuff with http://grub.org)",
        "Mozilla/4.0 (compatible; grub-client-1.3.7; Crawl your own stuff with http://grub.org)",
        "Mozilla/4.0 (compatible; grub-client-1.4.3; Crawl your own stuff with http://grub.org)",
        "Mozilla/4.0 (compatible; grub-client-1.5.3; Crawl your own stuff with http://grub.org)":
		return &Crawler{
			Pattern:      `grub.org`,
			Instances:    []string{ "Mozilla/4.0 (compatible; grub-client-0.3.0; Crawl your own stuff with http://grub.org)", "Mozilla/4.0 (compatible; grub-client-1.0.4; Crawl your own stuff with http://grub.org)", "Mozilla/4.0 (compatible; grub-client-1.0.5; Crawl your own stuff with http://grub.org)", "Mozilla/4.0 (compatible; grub-client-1.0.6; Crawl your own stuff with http://grub.org)", "Mozilla/4.0 (compatible; grub-client-1.0.7; Crawl your own stuff with http://grub.org)", "Mozilla/4.0 (compatible; grub-client-1.1.1; Crawl your own stuff with http://grub.org)", "Mozilla/4.0 (compatible; grub-client-1.2.1; Crawl your own stuff with http://grub.org)", "Mozilla/4.0 (compatible; grub-client-1.3.1; Crawl your own stuff with http://grub.org)", "Mozilla/4.0 (compatible; grub-client-1.3.7; Crawl your own stuff with http://grub.org)", "Mozilla/4.0 (compatible; grub-client-1.4.3; Crawl your own stuff with http://grub.org)", "Mozilla/4.0 (compatible; grub-client-1.5.3; Crawl your own stuff with http://grub.org)" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Mozilla/5.0 (Windows; U; Windows NT 5.1; en-US) Speedy Spider (http://www.entireweb.com/about/search_tech/speedy_spider/)",
        "Mozilla/5.0 (Windows; U; Windows NT 5.1; en-US) Speedy Spider for SpeedyAds (http://www.entireweb.com/about/search_tech/speedy_spider/)",
        "Mozilla/5.0 (compatible; Speedy Spider; http://www.entireweb.com/about/search_tech/speedy_spider/)",
        "Speedy Spider (Entireweb; Beta/1.2; http://www.entireweb.com/about/search_tech/speedyspider/)",
        "Speedy Spider (http://www.entireweb.com/about/search_tech/speedy_spider/)":
		return &Crawler{
			Pattern:      `speedy`,
			Instances:    []string{ "Mozilla/5.0 (Windows; U; Windows NT 5.1; en-US) Speedy Spider (http://www.entireweb.com/about/search_tech/speedy_spider/)", "Mozilla/5.0 (Windows; U; Windows NT 5.1; en-US) Speedy Spider for SpeedyAds (http://www.entireweb.com/about/search_tech/speedy_spider/)", "Mozilla/5.0 (compatible; Speedy Spider; http://www.entireweb.com/about/search_tech/speedy_spider/)", "Speedy Spider (Entireweb; Beta/1.2; http://www.entireweb.com/about/search_tech/speedyspider/)", "Speedy Spider (http://www.entireweb.com/about/search_tech/speedy_spider/)" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "findlinks/1.0 (+http://wortschatz.uni-leipzig.de/findlinks/)",
        "findlinks/1.1.3-beta8 (+http://wortschatz.uni-leipzig.de/findlinks/)",
        "findlinks/1.1.3-beta9 (+http://wortschatz.uni-leipzig.de/findlinks/)",
        "findlinks/1.1.5-beta7 (+http://wortschatz.uni-leipzig.de/findlinks/)",
        "findlinks/1.1.6-beta1 (+http://wortschatz.uni-leipzig.de/findlinks/)",
        "findlinks/1.1.6-beta1 (+http://wortschatz.uni-leipzig.de/findlinks/; YaCy 0.1; yacy.net)",
        "findlinks/1.1.6-beta2 (+http://wortschatz.uni-leipzig.de/findlinks/)",
        "findlinks/1.1.6-beta3 (+http://wortschatz.uni-leipzig.de/findlinks/)",
        "findlinks/1.1.6-beta4 (+http://wortschatz.uni-leipzig.de/findlinks/)",
        "findlinks/1.1.6-beta5 (+http://wortschatz.uni-leipzig.de/findlinks/)",
        "findlinks/1.1.6-beta6 (+http://wortschatz.uni-leipzig.de/findlinks/)",
        "findlinks/2.0 (+http://wortschatz.uni-leipzig.de/findlinks/)",
        "findlinks/2.0.1 (+http://wortschatz.uni-leipzig.de/findlinks/)",
        "findlinks/2.0.2 (+http://wortschatz.uni-leipzig.de/findlinks/)",
        "findlinks/2.0.4 (+http://wortschatz.uni-leipzig.de/findlinks/)",
        "findlinks/2.0.5 (+http://wortschatz.uni-leipzig.de/findlinks/)",
        "findlinks/2.0.9 (+http://wortschatz.uni-leipzig.de/findlinks/)",
        "findlinks/2.1 (+http://wortschatz.uni-leipzig.de/findlinks/)",
        "findlinks/2.1.3 (+http://wortschatz.uni-leipzig.de/findlinks/)",
        "findlinks/2.1.5 (+http://wortschatz.uni-leipzig.de/findlinks/)",
        "findlinks/2.2 (+http://wortschatz.uni-leipzig.de/findlinks/)",
        "findlinks/2.5 (+http://wortschatz.uni-leipzig.de/findlinks/)",
        "findlinks/2.6 (+http://wortschatz.uni-leipzig.de/findlinks/)":
		return &Crawler{
			Pattern:      `findlink`,
			Instances:    []string{ "findlinks/1.0 (+http://wortschatz.uni-leipzig.de/findlinks/)", "findlinks/1.1.3-beta8 (+http://wortschatz.uni-leipzig.de/findlinks/)", "findlinks/1.1.3-beta9 (+http://wortschatz.uni-leipzig.de/findlinks/)", "findlinks/1.1.5-beta7 (+http://wortschatz.uni-leipzig.de/findlinks/)", "findlinks/1.1.6-beta1 (+http://wortschatz.uni-leipzig.de/findlinks/)", "findlinks/1.1.6-beta1 (+http://wortschatz.uni-leipzig.de/findlinks/; YaCy 0.1; yacy.net)", "findlinks/1.1.6-beta2 (+http://wortschatz.uni-leipzig.de/findlinks/)", "findlinks/1.1.6-beta3 (+http://wortschatz.uni-leipzig.de/findlinks/)", "findlinks/1.1.6-beta4 (+http://wortschatz.uni-leipzig.de/findlinks/)", "findlinks/1.1.6-beta5 (+http://wortschatz.uni-leipzig.de/findlinks/)", "findlinks/1.1.6-beta6 (+http://wortschatz.uni-leipzig.de/findlinks/)", "findlinks/2.0 (+http://wortschatz.uni-leipzig.de/findlinks/)", "findlinks/2.0.1 (+http://wortschatz.uni-leipzig.de/findlinks/)", "findlinks/2.0.2 (+http://wortschatz.uni-leipzig.de/findlinks/)", "findlinks/2.0.4 (+http://wortschatz.uni-leipzig.de/findlinks/)", "findlinks/2.0.5 (+http://wortschatz.uni-leipzig.de/findlinks/)", "findlinks/2.0.9 (+http://wortschatz.uni-leipzig.de/findlinks/)", "findlinks/2.1 (+http://wortschatz.uni-leipzig.de/findlinks/)", "findlinks/2.1.3 (+http://wortschatz.uni-leipzig.de/findlinks/)", "findlinks/2.1.5 (+http://wortschatz.uni-leipzig.de/findlinks/)", "findlinks/2.2 (+http://wortschatz.uni-leipzig.de/findlinks/)", "findlinks/2.5 (+http://wortschatz.uni-leipzig.de/findlinks/)", "findlinks/2.6 (+http://wortschatz.uni-leipzig.de/findlinks/)" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "panscient.com":
		return &Crawler{
			Pattern:      `panscient`,
			Instances:    []string{ "panscient.com" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "yacybot (/global; amd64 FreeBSD 10.3-RELEASE; java 1.8.0_77; GMT/en) http://yacy.net/bot.html",
        "yacybot (/global; amd64 FreeBSD 10.3-RELEASE-p7; java 1.7.0_95; GMT/en) http://yacy.net/bot.html",
        "yacybot (-global; amd64 FreeBSD 9.2-RELEASE-p10; java 1.7.0_65; Europe/en) http://yacy.net/bot.html",
        "yacybot (/global; amd64 Linux 2.6.32-042stab093.4; java 1.7.0_65; Etc/en) http://yacy.net/bot.html",
        "yacybot (/global; amd64 Linux 2.6.32-042stab094.8; java 1.7.0_79; America/en) http://yacy.net/bot.html",
        "yacybot (/global; amd64 Linux 2.6.32-042stab108.8; java 1.7.0_91; America/en) http://yacy.net/bot.html",
        "yacybot (-global; amd64 Linux 2.6.32-042stab111.11; java 1.7.0_79; Europe/en) http://yacy.net/bot.html",
        "yacybot (-global; amd64 Linux 2.6.32-042stab116.1; java 1.7.0_79; Europe/en) http://yacy.net/bot.html",
        "yacybot (/global; amd64 Linux 2.6.32-573.3.1.el6.x86_64; java 1.7.0_85; Europe/en) http://yacy.net/bot.html",
        "yacybot (-global; amd64 Linux 3.10.0-229.4.2.el7.x86_64; java 1.7.0_79; Europe/en) http://yacy.net/bot.html",
        "yacybot (-global; amd64 Linux 3.10.0-229.4.2.el7.x86_64; java 1.8.0_45; Europe/en) http://yacy.net/bot.html",
        "yacybot (/global; amd64 Linux 3.10.0-229.7.2.el7.x86_64; java 1.8.0_45; Europe/en) http://yacy.net/bot.html",
        "yacybot (/global; amd64 Linux 3.10.0-327.22.2.el7.x86_64; java 1.7.0_101; Etc/en) http://yacy.net/bot.html",
        "yacybot (/global; amd64 Linux 3.11.10-21-desktop; java 1.7.0_51; America/en) http://yacy.net/bot.html",
        "yacybot (/global; amd64 Linux 3.12.1; java 1.7.0_65; Europe/en) http://yacy.net/bot.html",
        "yacybot (/global; amd64 Linux 3.13.0-042stab093.4; java 1.7.0_79; Europe/de) http://yacy.net/bot.html",
        "yacybot (/global; amd64 Linux 3.13.0-042stab093.4; java 1.7.0_79; Europe/en) http://yacy.net/bot.html",
        "yacybot (/global; amd64 Linux 3.13.0-45-generic; java 1.7.0_75; Europe/en) http://yacy.net/bot.html",
        "yacybot (-global; amd64 Linux 3.13.0-61-generic; java 1.7.0_79; Europe/en) http://yacy.net/bot.html",
        "yacybot (/global; amd64 Linux 3.13.0-74-generic; java 1.7.0_91; Europe/en) http://yacy.net/bot.html",
        "yacybot (/global; amd64 Linux 3.13.0-83-generic; java 1.7.0_95; Europe/de) http://yacy.net/bot.html",
        "yacybot (/global; amd64 Linux 3.13.0-83-generic; java 1.7.0_95; Europe/en) http://yacy.net/bot.html",
        "yacybot (/global; amd64 Linux 3.13.0-85-generic; java 1.7.0_101; Europe/en) http://yacy.net/bot.html",
        "yacybot (/global; amd64 Linux 3.13.0-85-generic; java 1.7.0_95; Europe/en) http://yacy.net/bot.html",
        "yacybot (/global; amd64 Linux 3.13.0-88-generic; java 1.7.0_101; Europe/en) http://yacy.net/bot.html",
        "yacybot (/global; amd64 Linux 3.14-0.bpo.1-amd64; java 1.7.0_55; Europe/de) http://yacy.net/bot.html",
        "yacybot (/global; amd64 Linux 3.14.32-xxxx-grs-ipv6-64; java 1.7.0_75; Europe/en) http://yacy.net/bot.html",
        "yacybot (-global; amd64 Linux 3.14.32-xxxx-grs-ipv6-64; java 1.8.0_111; Europe/de) http://yacy.net/bot.html",
        "yacybot (/global; amd64 Linux 3.16.0-4-amd64; java 1.7.0_111; Europe/de) http://yacy.net/bot.html",
        "yacybot (/global; amd64 Linux 3.16.0-4-amd64; java 1.7.0_75; America/en) http://yacy.net/bot.html",
        "yacybot (-global; amd64 Linux 3.16.0-4-amd64; java 1.7.0_75; Europe/en) http://yacy.net/bot.html",
        "yacybot (/global; amd64 Linux 3.16.0-4-amd64; java 1.7.0_75; Europe/en) http://yacy.net/bot.html",
        "yacybot (/global; amd64 Linux 3.16.0-4-amd64; java 1.7.0_79; Europe/de) http://yacy.net/bot.html",
        "yacybot (/global; amd64 Linux 3.16.0-4-amd64; java 1.7.0_79; Europe/en) http://yacy.net/bot.html",
        "yacybot (/global; amd64 Linux 3.16.0-4-amd64; java 1.7.0_91; Europe/de) http://yacy.net/bot.html",
        "yacybot (/global; amd64 Linux 3.16.0-4-amd64; java 1.7.0_95; Europe/en) http://yacy.net/bot.html",
        "yacybot (/global; amd64 Linux 3.16.0-4-amd64; java 1.8.0_111; Europe/en) http://yacy.net/bot.html",
        "yacybot (/global; amd64 Linux 3.16-0.bpo.2-amd64; java 1.7.0_65; Europe/en) http://yacy.net/bot.html",
        "yacybot (-global; amd64 Linux 3.19.0-15-generic; java 1.8.0_45-internal; Europe/de) http://yacy.net/bot.html",
        "yacybot (-global; amd64 Linux 3.2.0-4-amd64; java 1.7.0_65; Europe/en) http://yacy.net/bot.html",
        "yacybot (-global; amd64 Linux 3.2.0-4-amd64; java 1.7.0_67; Europe/en) http://yacy.net/bot.html",
        "yacybot (-global; amd64 Linux 4.4.0-57-generic; java 9-internal; Europe/en) http://yacy.net/bot.html",
        "yacybot (-global; amd64 Windows 8.1 6.3; java 1.7.0_55; Europe/de) http://yacy.net/bot.html",
        "yacybot (-global; amd64 Windows 8 6.2; java 1.7.0_55; Europe/de) http://yacy.net/bot.html":
		return &Crawler{
			Pattern:      `yacybot`,
			Instances:    []string{ "yacybot (/global; amd64 FreeBSD 10.3-RELEASE; java 1.8.0_77; GMT/en) http://yacy.net/bot.html", "yacybot (/global; amd64 FreeBSD 10.3-RELEASE-p7; java 1.7.0_95; GMT/en) http://yacy.net/bot.html", "yacybot (-global; amd64 FreeBSD 9.2-RELEASE-p10; java 1.7.0_65; Europe/en) http://yacy.net/bot.html", "yacybot (/global; amd64 Linux 2.6.32-042stab093.4; java 1.7.0_65; Etc/en) http://yacy.net/bot.html", "yacybot (/global; amd64 Linux 2.6.32-042stab094.8; java 1.7.0_79; America/en) http://yacy.net/bot.html", "yacybot (/global; amd64 Linux 2.6.32-042stab108.8; java 1.7.0_91; America/en) http://yacy.net/bot.html", "yacybot (-global; amd64 Linux 2.6.32-042stab111.11; java 1.7.0_79; Europe/en) http://yacy.net/bot.html", "yacybot (-global; amd64 Linux 2.6.32-042stab116.1; java 1.7.0_79; Europe/en) http://yacy.net/bot.html", "yacybot (/global; amd64 Linux 2.6.32-573.3.1.el6.x86_64; java 1.7.0_85; Europe/en) http://yacy.net/bot.html", "yacybot (-global; amd64 Linux 3.10.0-229.4.2.el7.x86_64; java 1.7.0_79; Europe/en) http://yacy.net/bot.html", "yacybot (-global; amd64 Linux 3.10.0-229.4.2.el7.x86_64; java 1.8.0_45; Europe/en) http://yacy.net/bot.html", "yacybot (/global; amd64 Linux 3.10.0-229.7.2.el7.x86_64; java 1.8.0_45; Europe/en) http://yacy.net/bot.html", "yacybot (/global; amd64 Linux 3.10.0-327.22.2.el7.x86_64; java 1.7.0_101; Etc/en) http://yacy.net/bot.html", "yacybot (/global; amd64 Linux 3.11.10-21-desktop; java 1.7.0_51; America/en) http://yacy.net/bot.html", "yacybot (/global; amd64 Linux 3.12.1; java 1.7.0_65; Europe/en) http://yacy.net/bot.html", "yacybot (/global; amd64 Linux 3.13.0-042stab093.4; java 1.7.0_79; Europe/de) http://yacy.net/bot.html", "yacybot (/global; amd64 Linux 3.13.0-042stab093.4; java 1.7.0_79; Europe/en) http://yacy.net/bot.html", "yacybot (/global; amd64 Linux 3.13.0-45-generic; java 1.7.0_75; Europe/en) http://yacy.net/bot.html", "yacybot (-global; amd64 Linux 3.13.0-61-generic; java 1.7.0_79; Europe/en) http://yacy.net/bot.html", "yacybot (/global; amd64 Linux 3.13.0-74-generic; java 1.7.0_91; Europe/en) http://yacy.net/bot.html", "yacybot (/global; amd64 Linux 3.13.0-83-generic; java 1.7.0_95; Europe/de) http://yacy.net/bot.html", "yacybot (/global; amd64 Linux 3.13.0-83-generic; java 1.7.0_95; Europe/en) http://yacy.net/bot.html", "yacybot (/global; amd64 Linux 3.13.0-85-generic; java 1.7.0_101; Europe/en) http://yacy.net/bot.html", "yacybot (/global; amd64 Linux 3.13.0-85-generic; java 1.7.0_95; Europe/en) http://yacy.net/bot.html", "yacybot (/global; amd64 Linux 3.13.0-88-generic; java 1.7.0_101; Europe/en) http://yacy.net/bot.html", "yacybot (/global; amd64 Linux 3.14-0.bpo.1-amd64; java 1.7.0_55; Europe/de) http://yacy.net/bot.html", "yacybot (/global; amd64 Linux 3.14.32-xxxx-grs-ipv6-64; java 1.7.0_75; Europe/en) http://yacy.net/bot.html", "yacybot (-global; amd64 Linux 3.14.32-xxxx-grs-ipv6-64; java 1.8.0_111; Europe/de) http://yacy.net/bot.html", "yacybot (/global; amd64 Linux 3.16.0-4-amd64; java 1.7.0_111; Europe/de) http://yacy.net/bot.html", "yacybot (/global; amd64 Linux 3.16.0-4-amd64; java 1.7.0_75; America/en) http://yacy.net/bot.html", "yacybot (-global; amd64 Linux 3.16.0-4-amd64; java 1.7.0_75; Europe/en) http://yacy.net/bot.html", "yacybot (/global; amd64 Linux 3.16.0-4-amd64; java 1.7.0_75; Europe/en) http://yacy.net/bot.html", "yacybot (/global; amd64 Linux 3.16.0-4-amd64; java 1.7.0_79; Europe/de) http://yacy.net/bot.html", "yacybot (/global; amd64 Linux 3.16.0-4-amd64; java 1.7.0_79; Europe/en) http://yacy.net/bot.html", "yacybot (/global; amd64 Linux 3.16.0-4-amd64; java 1.7.0_91; Europe/de) http://yacy.net/bot.html", "yacybot (/global; amd64 Linux 3.16.0-4-amd64; java 1.7.0_95; Europe/en) http://yacy.net/bot.html", "yacybot (/global; amd64 Linux 3.16.0-4-amd64; java 1.8.0_111; Europe/en) http://yacy.net/bot.html", "yacybot (/global; amd64 Linux 3.16-0.bpo.2-amd64; java 1.7.0_65; Europe/en) http://yacy.net/bot.html", "yacybot (-global; amd64 Linux 3.19.0-15-generic; java 1.8.0_45-internal; Europe/de) http://yacy.net/bot.html", "yacybot (-global; amd64 Linux 3.2.0-4-amd64; java 1.7.0_65; Europe/en) http://yacy.net/bot.html", "yacybot (-global; amd64 Linux 3.2.0-4-amd64; java 1.7.0_67; Europe/en) http://yacy.net/bot.html", "yacybot (-global; amd64 Linux 4.4.0-57-generic; java 9-internal; Europe/en) http://yacy.net/bot.html", "yacybot (-global; amd64 Windows 8.1 6.3; java 1.7.0_55; Europe/de) http://yacy.net/bot.html", "yacybot (-global; amd64 Windows 8 6.2; java 1.7.0_55; Europe/de) http://yacy.net/bot.html" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "BlackBerry9000/4.6.0.167 Profile/MIDP-2.0 Configuration/CLDC-1.1 VendorID/102 ips-agent",
        "Mozilla/5.0 (X11; U; Linux i686; en-US; rv:1.7.12; ips-agent) Gecko/20050922 Fedora/1.0.7-1.1.fc4 Firefox/1.0.7",
        "Mozilla/5.0 (X11; U; Linux i686; en-US; rv:1.9.1.3; ips-agent) Gecko/20090824 Fedora/1.0.7-1.1.fc4  Firefox/3.5.3",
        "Mozilla/5.0 (X11; U; Linux i686; en-US; rv:1.9.2.24; ips-agent) Gecko/20111107 Ubuntu/10.04 (lucid) Firefox/3.6.24",
        "Mozilla/5.0 (X11; Ubuntu; Linux i686; rv:14.0; ips-agent) Gecko/20100101 Firefox/14.0.1":
		return &Crawler{
			Pattern:      `ips-agent`,
			Instances:    []string{ "BlackBerry9000/4.6.0.167 Profile/MIDP-2.0 Configuration/CLDC-1.1 VendorID/102 ips-agent", "Mozilla/5.0 (X11; U; Linux i686; en-US; rv:1.7.12; ips-agent) Gecko/20050922 Fedora/1.0.7-1.1.fc4 Firefox/1.0.7", "Mozilla/5.0 (X11; U; Linux i686; en-US; rv:1.9.1.3; ips-agent) Gecko/20090824 Fedora/1.0.7-1.1.fc4  Firefox/3.5.3", "Mozilla/5.0 (X11; U; Linux i686; en-US; rv:1.9.2.24; ips-agent) Gecko/20111107 Ubuntu/10.04 (lucid) Firefox/3.6.24", "Mozilla/5.0 (X11; Ubuntu; Linux i686; rv:14.0; ips-agent) Gecko/20100101 Firefox/14.0.1" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "MJ12bot/v1.2.0 (http://majestic12.co.uk/bot.php?+)",
        "Mozilla/5.0 (compatible; MJ12bot/v1.2.1; http://www.majestic12.co.uk/bot.php?+)",
        "Mozilla/5.0 (compatible; MJ12bot/v1.2.3; http://www.majestic12.co.uk/bot.php?+)",
        "Mozilla/5.0 (compatible; MJ12bot/v1.2.4; http://www.majestic12.co.uk/bot.php?+)",
        "Mozilla/5.0 (compatible; MJ12bot/v1.2.5; http://www.majestic12.co.uk/bot.php?+)",
        "Mozilla/5.0 (compatible; MJ12bot/v1.3.0; http://www.majestic12.co.uk/bot.php?+)",
        "Mozilla/5.0 (compatible; MJ12bot/v1.3.1; http://www.majestic12.co.uk/bot.php?+)",
        "Mozilla/5.0 (compatible; MJ12bot/v1.3.2; http://www.majestic12.co.uk/bot.php?+)",
        "Mozilla/5.0 (compatible; MJ12bot/v1.3.3; http://www.majestic12.co.uk/bot.php?+)",
        "Mozilla/5.0 (compatible; MJ12bot/v1.4.0; http://www.majestic12.co.uk/bot.php?+)",
        "Mozilla/5.0 (compatible; MJ12bot/v1.4.1; http://www.majestic12.co.uk/bot.php?+)",
        "Mozilla/5.0 (compatible; MJ12bot/v1.4.2; http://www.majestic12.co.uk/bot.php?+)",
        "Mozilla/5.0 (compatible; MJ12bot/v1.4.3; http://www.majestic12.co.uk/bot.php?+)",
        "Mozilla/5.0 (compatible; MJ12bot/v1.4.4 (domain ownership verifier); http://www.majestic12.co.uk/bot.php?+)",
        "Mozilla/5.0 (compatible; MJ12bot/v1.4.4; http://www.majestic12.co.uk/bot.php?+)",
        "Mozilla/5.0 (compatible; MJ12bot/v1.4.5; http://www.majestic12.co.uk/bot.php?+)",
        "Mozilla/5.0 (compatible; MJ12bot/v1.4.6; http://mj12bot.com/)",
        "Mozilla/5.0 (compatible; MJ12bot/v1.4.7; http://mj12bot.com/)",
        "Mozilla/5.0 (compatible; MJ12bot/v1.4.7; http://www.majestic12.co.uk/bot.php?+)",
        "Mozilla/5.0 (compatible; MJ12bot/v1.4.8; http://mj12bot.com/)":
		return &Crawler{
			Pattern:      `MJ12bot`,
			Instances:    []string{ "MJ12bot/v1.2.0 (http://majestic12.co.uk/bot.php?+)", "Mozilla/5.0 (compatible; MJ12bot/v1.2.1; http://www.majestic12.co.uk/bot.php?+)", "Mozilla/5.0 (compatible; MJ12bot/v1.2.3; http://www.majestic12.co.uk/bot.php?+)", "Mozilla/5.0 (compatible; MJ12bot/v1.2.4; http://www.majestic12.co.uk/bot.php?+)", "Mozilla/5.0 (compatible; MJ12bot/v1.2.5; http://www.majestic12.co.uk/bot.php?+)", "Mozilla/5.0 (compatible; MJ12bot/v1.3.0; http://www.majestic12.co.uk/bot.php?+)", "Mozilla/5.0 (compatible; MJ12bot/v1.3.1; http://www.majestic12.co.uk/bot.php?+)", "Mozilla/5.0 (compatible; MJ12bot/v1.3.2; http://www.majestic12.co.uk/bot.php?+)", "Mozilla/5.0 (compatible; MJ12bot/v1.3.3; http://www.majestic12.co.uk/bot.php?+)", "Mozilla/5.0 (compatible; MJ12bot/v1.4.0; http://www.majestic12.co.uk/bot.php?+)", "Mozilla/5.0 (compatible; MJ12bot/v1.4.1; http://www.majestic12.co.uk/bot.php?+)", "Mozilla/5.0 (compatible; MJ12bot/v1.4.2; http://www.majestic12.co.uk/bot.php?+)", "Mozilla/5.0 (compatible; MJ12bot/v1.4.3; http://www.majestic12.co.uk/bot.php?+)", "Mozilla/5.0 (compatible; MJ12bot/v1.4.4 (domain ownership verifier); http://www.majestic12.co.uk/bot.php?+)", "Mozilla/5.0 (compatible; MJ12bot/v1.4.4; http://www.majestic12.co.uk/bot.php?+)", "Mozilla/5.0 (compatible; MJ12bot/v1.4.5; http://www.majestic12.co.uk/bot.php?+)", "Mozilla/5.0 (compatible; MJ12bot/v1.4.6; http://mj12bot.com/)", "Mozilla/5.0 (compatible; MJ12bot/v1.4.7; http://mj12bot.com/)", "Mozilla/5.0 (compatible; MJ12bot/v1.4.7; http://www.majestic12.co.uk/bot.php?+)", "Mozilla/5.0 (compatible; MJ12bot/v1.4.8; http://mj12bot.com/)" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Mozilla/5.0 (compatible; woriobot +http://worio.com)",
        "Mozilla/5.0 (compatible; woriobot support [at] zite [dot] com +http://zite.com)":
		return &Crawler{
			Pattern:      `woriobot`,
			Instances:    []string{ "Mozilla/5.0 (compatible; woriobot +http://worio.com)", "Mozilla/5.0 (compatible; woriobot support [at] zite [dot] com +http://zite.com)" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Yanga WorldSearch Bot v1.1/beta (http://www.yanga.co.uk/)":
		return &Crawler{
			Pattern:      `yanga`,
			Instances:    []string{ "Yanga WorldSearch Bot v1.1/beta (http://www.yanga.co.uk/)" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Buzzbot/1.0 (Buzzbot; http://www.buzzstream.com; buzzbot@buzzstream.com)":
		return &Crawler{
			Pattern:      `buzzbot`,
			Instances:    []string{ "Buzzbot/1.0 (Buzzbot; http://www.buzzstream.com; buzzbot@buzzstream.com)" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "MLBot (www.metadatalabs.com/mlbot)":
		return &Crawler{
			Pattern:      `mlbot`,
			Instances:    []string{ "MLBot (www.metadatalabs.com/mlbot)" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Mozilla/5.0 (compatible; YandexBot/3.0; +http://yandex.com/bots)":
		return &Crawler{
			Pattern:      `YandexBot`,
			Instances:    []string{ "Mozilla/5.0 (compatible; YandexBot/3.0; +http://yandex.com/bots)" },
			URL:          stringer("http://yandex.com/bots"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Mozilla/5.0 (compatible; YandexImages/3.0; +http://yandex.com/bots)":
		return &Crawler{
			Pattern:      `YandexImages`,
			Instances:    []string{ "Mozilla/5.0 (compatible; YandexImages/3.0; +http://yandex.com/bots)" },
			URL:          stringer("http://yandex.com/bots"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Mozilla/5.0 (compatible; YandexAccessibilityBot/3.0; +http://yandex.com/bots":
		return &Crawler{
			Pattern:      `YandexAccessibilityBot`,
			Instances:    []string{ "Mozilla/5.0 (compatible; YandexAccessibilityBot/3.0; +http://yandex.com/bots" },
			URL:          stringer("http://yandex.com/bots"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Mozilla/5.0 (iPhone; CPU iPhone OS 8_1 like Mac OS X) AppleWebKit/600.1.4 (KHTML, like Gecko) Version/8.0 Mobile/12B411 Safari/600.1.4 (compatible; YandexMobileBot/3.0; +http://yandex.com/bots)":
		return &Crawler{
			Pattern:      `YandexMobileBot`,
			Instances:    []string{ "Mozilla/5.0 (iPhone; CPU iPhone OS 8_1 like Mac OS X) AppleWebKit/600.1.4 (KHTML, like Gecko) Version/8.0 Mobile/12B411 Safari/600.1.4 (compatible; YandexMobileBot/3.0; +http://yandex.com/bots)" },
			URL:          stringer("https://yandex.com/support/webmaster/robot-workings/check-yandex-robots.xml#robot-in-logs"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Linguee Bot (http://www.linguee.com/bot)",
        "Linguee Bot (http://www.linguee.com/bot; bot@linguee.com)":
		return &Crawler{
			Pattern:      `Linguee Bot`,
			Instances:    []string{ "Linguee Bot (http://www.linguee.com/bot)", "Linguee Bot (http://www.linguee.com/bot; bot@linguee.com)" },
			URL:          stringer("http://www.linguee.com/bot"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "CyberPatrol SiteCat Webbot (http://www.cyberpatrol.com/cyberpatrolcrawler.asp)":
		return &Crawler{
			Pattern:      `CyberPatrol`,
			Instances:    []string{ "CyberPatrol SiteCat Webbot (http://www.cyberpatrol.com/cyberpatrolcrawler.asp)" },
			URL:          stringer("http://www.cyberpatrol.com/cyberpatrolcrawler.asp"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Mozilla/5.0 (Windows NT 5.1; U; Win64; fr; rv:1.8.1) VoilaBot BETA 1.2 (support.voilabot@orange-ftgroup.com)",
        "Mozilla/5.0 (Windows; U; Windows NT 5.1; fr; rv:1.8.1) VoilaBot BETA 1.2 (support.voilabot@orange-ftgroup.com)":
		return &Crawler{
			Pattern:      `voilabot`,
			Instances:    []string{ "Mozilla/5.0 (Windows NT 5.1; U; Win64; fr; rv:1.8.1) VoilaBot BETA 1.2 (support.voilabot@orange-ftgroup.com)", "Mozilla/5.0 (Windows; U; Windows NT 5.1; fr; rv:1.8.1) VoilaBot BETA 1.2 (support.voilabot@orange-ftgroup.com)" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Mozilla/5.0 (compatible; Baiduspider/2.0; +http://www.baidu.com/search/spider.html)",
        "Mozilla/5.0 (compatible; Baiduspider-render/2.0; +http://www.baidu.com/search/spider.html)":
		return &Crawler{
			Pattern:      `Baiduspider`,
			Instances:    []string{ "Mozilla/5.0 (compatible; Baiduspider/2.0; +http://www.baidu.com/search/spider.html)", "Mozilla/5.0 (compatible; Baiduspider-render/2.0; +http://www.baidu.com/search/spider.html)" },
			URL:          stringer("http://www.baidu.jp/spider/"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Mozilla/5.0 (compatible; spbot/1.0; +http://www.seoprofiler.com/bot/ )",
        "Mozilla/5.0 (compatible; spbot/1.1; +http://www.seoprofiler.com/bot/ )",
        "Mozilla/5.0 (compatible; spbot/1.2; +http://www.seoprofiler.com/bot/ )",
        "Mozilla/5.0 (compatible; spbot/2.0.1; +http://www.seoprofiler.com/bot/ )",
        "Mozilla/5.0 (compatible; spbot/2.0.2; +http://www.seoprofiler.com/bot/ )",
        "Mozilla/5.0 (compatible; spbot/2.0.3; +http://www.seoprofiler.com/bot/ )",
        "Mozilla/5.0 (compatible; spbot/2.0.4; +http://www.seoprofiler.com/bot )",
        "Mozilla/5.0 (compatible; spbot/2.0; +http://www.seoprofiler.com/bot/ )",
        "Mozilla/5.0 (compatible; spbot/2.1; +http://www.seoprofiler.com/bot )",
        "Mozilla/5.0 (compatible; spbot/3.0; +http://www.seoprofiler.com/bot )",
        "Mozilla/5.0 (compatible; spbot/3.1; +http://www.seoprofiler.com/bot )",
        "Mozilla/5.0 (compatible; spbot/4.0.1; +http://www.seoprofiler.com/bot )",
        "Mozilla/5.0 (compatible; spbot/4.0.2; +http://www.seoprofiler.com/bot )",
        "Mozilla/5.0 (compatible; spbot/4.0.3; +http://www.seoprofiler.com/bot )",
        "Mozilla/5.0 (compatible; spbot/4.0.4; +http://www.seoprofiler.com/bot )",
        "Mozilla/5.0 (compatible; spbot/4.0.5; +http://www.seoprofiler.com/bot )",
        "Mozilla/5.0 (compatible; spbot/4.0.6; +http://www.seoprofiler.com/bot )",
        "Mozilla/5.0 (compatible; spbot/4.0.7; +http://OpenLinkProfiler.org/bot )",
        "Mozilla/5.0 (compatible; spbot/4.0.7; +https://www.seoprofiler.com/bot )",
        "Mozilla/5.0 (compatible; spbot/4.0.8; +http://OpenLinkProfiler.org/bot )",
        "Mozilla/5.0 (compatible; spbot/4.0.9; +http://OpenLinkProfiler.org/bot )",
        "Mozilla/5.0 (compatible; spbot/4.0; +http://www.seoprofiler.com/bot )",
        "Mozilla/5.0 (compatible; spbot/4.0a; +http://www.seoprofiler.com/bot )",
        "Mozilla/5.0 (compatible; spbot/4.0b; +http://www.seoprofiler.com/bot )",
        "Mozilla/5.0 (compatible; spbot/4.1.0; +http://OpenLinkProfiler.org/bot )",
        "Mozilla/5.0 (compatible; spbot/4.2.0; +http://OpenLinkProfiler.org/bot )",
        "Mozilla/5.0 (compatible; spbot/4.3.0; +http://OpenLinkProfiler.org/bot )",
        "Mozilla/5.0 (compatible; spbot/4.4.0; +http://OpenLinkProfiler.org/bot )",
        "Mozilla/5.0 (compatible; spbot/4.4.1; +http://OpenLinkProfiler.org/bot )",
        "Mozilla/5.0 (compatible; spbot/4.4.2; +http://OpenLinkProfiler.org/bot )",
        "Mozilla/5.0 (compatible; spbot/5.0.1; +http://OpenLinkProfiler.org/bot )",
        "Mozilla/5.0 (compatible; spbot/5.0.2; +http://OpenLinkProfiler.org/bot )",
        "Mozilla/5.0 (compatible; spbot/5.0.3; +http://OpenLinkProfiler.org/bot )",
        "Mozilla/5.0 (compatible; spbot/5.0; +http://OpenLinkProfiler.org/bot )":
		return &Crawler{
			Pattern:      `spbot`,
			Instances:    []string{ "Mozilla/5.0 (compatible; spbot/1.0; +http://www.seoprofiler.com/bot/ )", "Mozilla/5.0 (compatible; spbot/1.1; +http://www.seoprofiler.com/bot/ )", "Mozilla/5.0 (compatible; spbot/1.2; +http://www.seoprofiler.com/bot/ )", "Mozilla/5.0 (compatible; spbot/2.0.1; +http://www.seoprofiler.com/bot/ )", "Mozilla/5.0 (compatible; spbot/2.0.2; +http://www.seoprofiler.com/bot/ )", "Mozilla/5.0 (compatible; spbot/2.0.3; +http://www.seoprofiler.com/bot/ )", "Mozilla/5.0 (compatible; spbot/2.0.4; +http://www.seoprofiler.com/bot )", "Mozilla/5.0 (compatible; spbot/2.0; +http://www.seoprofiler.com/bot/ )", "Mozilla/5.0 (compatible; spbot/2.1; +http://www.seoprofiler.com/bot )", "Mozilla/5.0 (compatible; spbot/3.0; +http://www.seoprofiler.com/bot )", "Mozilla/5.0 (compatible; spbot/3.1; +http://www.seoprofiler.com/bot )", "Mozilla/5.0 (compatible; spbot/4.0.1; +http://www.seoprofiler.com/bot )", "Mozilla/5.0 (compatible; spbot/4.0.2; +http://www.seoprofiler.com/bot )", "Mozilla/5.0 (compatible; spbot/4.0.3; +http://www.seoprofiler.com/bot )", "Mozilla/5.0 (compatible; spbot/4.0.4; +http://www.seoprofiler.com/bot )", "Mozilla/5.0 (compatible; spbot/4.0.5; +http://www.seoprofiler.com/bot )", "Mozilla/5.0 (compatible; spbot/4.0.6; +http://www.seoprofiler.com/bot )", "Mozilla/5.0 (compatible; spbot/4.0.7; +http://OpenLinkProfiler.org/bot )", "Mozilla/5.0 (compatible; spbot/4.0.7; +https://www.seoprofiler.com/bot )", "Mozilla/5.0 (compatible; spbot/4.0.8; +http://OpenLinkProfiler.org/bot )", "Mozilla/5.0 (compatible; spbot/4.0.9; +http://OpenLinkProfiler.org/bot )", "Mozilla/5.0 (compatible; spbot/4.0; +http://www.seoprofiler.com/bot )", "Mozilla/5.0 (compatible; spbot/4.0a; +http://www.seoprofiler.com/bot )", "Mozilla/5.0 (compatible; spbot/4.0b; +http://www.seoprofiler.com/bot )", "Mozilla/5.0 (compatible; spbot/4.1.0; +http://OpenLinkProfiler.org/bot )", "Mozilla/5.0 (compatible; spbot/4.2.0; +http://OpenLinkProfiler.org/bot )", "Mozilla/5.0 (compatible; spbot/4.3.0; +http://OpenLinkProfiler.org/bot )", "Mozilla/5.0 (compatible; spbot/4.4.0; +http://OpenLinkProfiler.org/bot )", "Mozilla/5.0 (compatible; spbot/4.4.1; +http://OpenLinkProfiler.org/bot )", "Mozilla/5.0 (compatible; spbot/4.4.2; +http://OpenLinkProfiler.org/bot )", "Mozilla/5.0 (compatible; spbot/5.0.1; +http://OpenLinkProfiler.org/bot )", "Mozilla/5.0 (compatible; spbot/5.0.2; +http://OpenLinkProfiler.org/bot )", "Mozilla/5.0 (compatible; spbot/5.0.3; +http://OpenLinkProfiler.org/bot )", "Mozilla/5.0 (compatible; spbot/5.0; +http://OpenLinkProfiler.org/bot )" },
			URL:          stringer("http://www.seoprofiler.com/bot"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "PostRank/2.0 (postrank.com)",
        "PostRank/2.0 (postrank.com; 1 subscribers)":
		return &Crawler{
			Pattern:      `postrank`,
			Instances:    []string{ "PostRank/2.0 (postrank.com)", "PostRank/2.0 (postrank.com; 1 subscribers)" },
			URL:          stringer("http://www.postrank.com"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "TurnitinBot (https://turnitin.com/robot/crawlerinfo.html)":
		return &Crawler{
			Pattern:      `TurnitinBot`,
			Instances:    []string{ "TurnitinBot (https://turnitin.com/robot/crawlerinfo.html)" },
			URL:          stringer("http://www.turnitin.com"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Mozilla/5.0 (compatible;  Page2RSS/0.7; +http://page2rss.com/)":
		return &Crawler{
			Pattern:      `page2rss`,
			Instances:    []string{ "Mozilla/5.0 (compatible;  Page2RSS/0.7; +http://page2rss.com/)" },
			URL:          stringer("http://www.page2rss.com"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Mozilla/5.0 (compatible; Whoiswebsitebot/0.1; +http://www.whoiswebsite.net)":
		return &Crawler{
			Pattern:      `sitebot`,
			Instances:    []string{ "Mozilla/5.0 (compatible; Whoiswebsitebot/0.1; +http://www.whoiswebsite.net)" },
			URL:          stringer("http://www.sitebot.org"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Mozilla/5.0 (compatible; linkdexbot/2.0; +http://www.linkdex.com/about/bots/)",
        "Mozilla/5.0 (compatible; linkdexbot/2.0; +http://www.linkdex.com/bots/)",
        "Mozilla/5.0 (compatible; linkdexbot/2.1; +http://www.linkdex.com/about/bots/)",
        "Mozilla/5.0 (compatible; linkdexbot/2.1; +http://www.linkdex.com/bots/)",
        "Mozilla/5.0 (compatible; linkdexbot/2.2; +http://www.linkdex.com/bots/)",
        "linkdex.com/v2.0",
        "linkdexbot/Nutch-1.0-dev (http://www.linkdex.com/; crawl at linkdex dot com)":
		return &Crawler{
			Pattern:      `linkdex`,
			Instances:    []string{ "Mozilla/5.0 (compatible; linkdexbot/2.0; +http://www.linkdex.com/about/bots/)", "Mozilla/5.0 (compatible; linkdexbot/2.0; +http://www.linkdex.com/bots/)", "Mozilla/5.0 (compatible; linkdexbot/2.1; +http://www.linkdex.com/about/bots/)", "Mozilla/5.0 (compatible; linkdexbot/2.1; +http://www.linkdex.com/bots/)", "Mozilla/5.0 (compatible; linkdexbot/2.2; +http://www.linkdex.com/bots/)", "linkdex.com/v2.0", "linkdexbot/Nutch-1.0-dev (http://www.linkdex.com/; crawl at linkdex dot com)" },
			URL:          stringer("http://www.linkdex.com"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Mozilla/5.0 (compatible; Ezooms/1.0; ezooms.bot@gmail.com)":
		return &Crawler{
			Pattern:      `ezooms`,
			Instances:    []string{ "Mozilla/5.0 (compatible; Ezooms/1.0; ezooms.bot@gmail.com)" },
			URL:          stringer("http://www.phpbb.com/community/viewtopic.php?f=64&t=935605&start=450#p12948289"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Mozilla/5.0 (compatible; DotBot/1.1; http://www.opensiteexplorer.org/dotbot, help@moz.com)",
        "dotbot":
		return &Crawler{
			Pattern:      `dotbot`,
			Instances:    []string{ "Mozilla/5.0 (compatible; DotBot/1.1; http://www.opensiteexplorer.org/dotbot, help@moz.com)", "dotbot" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Mozilla/5.0 (compatible; Linux x86_64; Mail.RU_Bot/2.0; +http://go.mail.ru/help/robots)",
        "Mozilla/5.0 (compatible; Linux x86_64; Mail.RU_Bot/2.0; +http://go.mail.ru/",
        "Mozilla/5.0 (compatible; Mail.RU_Bot/2.0; +http://go.mail.ru/",
        "Mozilla/5.0 (compatible; Linux x86_64; Mail.RU_Bot/Robots/2.0; +http://go.mail.ru/help/robots)":
		return &Crawler{
			Pattern:      `Mail.RU_Bot`,
			Instances:    []string{ "Mozilla/5.0 (compatible; Linux x86_64; Mail.RU_Bot/2.0; +http://go.mail.ru/help/robots)", "Mozilla/5.0 (compatible; Linux x86_64; Mail.RU_Bot/2.0; +http://go.mail.ru/", "Mozilla/5.0 (compatible; Mail.RU_Bot/2.0; +http://go.mail.ru/", "Mozilla/5.0 (compatible; Linux x86_64; Mail.RU_Bot/Robots/2.0; +http://go.mail.ru/help/robots)" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Mozilla/5.0 (compatible; discobot/1.0; +http://discoveryengine.com/discobot.html)",
        "Mozilla/5.0 (compatible; discobot/2.0; +http://discoveryengine.com/discobot.html)",
        "mozilla/5.0 (compatible; discobot/1.1; +http://discoveryengine.com/discobot.html)":
		return &Crawler{
			Pattern:      `discobot`,
			Instances:    []string{ "Mozilla/5.0 (compatible; discobot/1.0; +http://discoveryengine.com/discobot.html)", "Mozilla/5.0 (compatible; discobot/2.0; +http://discoveryengine.com/discobot.html)", "mozilla/5.0 (compatible; discobot/1.1; +http://discoveryengine.com/discobot.html)" },
			URL:          stringer("http://discoveryengine.com/discobot.html"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Mozilla/5.0 (compatible; heritrix/1.12.1 +http://www.webarchiv.cz)",
        "Mozilla/5.0 (compatible; heritrix/1.12.1b +http://netarkivet.dk/website/info.html)",
        "Mozilla/5.0 (compatible; heritrix/1.14.2 +http://rjpower.org)",
        "Mozilla/5.0 (compatible; heritrix/1.14.2 +http://www.webarchiv.cz)",
        "Mozilla/5.0 (compatible; heritrix/1.14.3 +http://archive.org)",
        "Mozilla/5.0 (compatible; heritrix/1.14.3 +http://www.accelobot.com)",
        "Mozilla/5.0 (compatible; heritrix/1.14.3 +http://www.webarchiv.cz)",
        "Mozilla/5.0 (compatible; heritrix/1.14.3.r6601 +http://www.buddybuzz.net/yptrino)",
        "Mozilla/5.0 (compatible; heritrix/1.14.4 +http://parsijoo.ir)",
        "Mozilla/5.0 (compatible; heritrix/1.14.4 +http://www.exif-search.com)",
        "Mozilla/5.0 (compatible; heritrix/2.0.2 +http://aihit.com)",
        "Mozilla/5.0 (compatible; heritrix/2.0.2 +http://seekda.com)",
        "Mozilla/5.0 (compatible; heritrix/3.0.0-SNAPSHOT-20091120.021634 +http://crawler.archive.org)",
        "Mozilla/5.0 (compatible; heritrix/3.1.0-RC1 +http://boston.lti.cs.cmu.edu/crawler_12/)",
        "Mozilla/5.0 (compatible; heritrix/3.1.1 +http://places.tomtom.com/crawlerinfo)",
        "Mozilla/5.0 (compatible; heritrix/3.1.1 +http://www.mixdata.com)",
        "Mozilla/5.0 (compatible; heritrix/3.1.1; UniLeipzigASV +http://corpora.informatik.uni-leipzig.de/crawler_faq.html)",
        "Mozilla/5.0 (compatible; heritrix/3.2.0 +http://www.crim.ca)",
        "Mozilla/5.0 (compatible; heritrix/3.2.0 +http://www.exif-search.com)",
        "Mozilla/5.0 (compatible; heritrix/3.2.0 +http://www.mixdata.com)",
        "Mozilla/5.0 (compatible; heritrix/3.3.0-SNAPSHOT-20160309-0050; UniLeipzigASV +http://corpora.informatik.uni-leipzig.de/crawler_faq.html)",
        "Mozilla/5.0 (compatible; sukibot_heritrix/3.1.1 +http://suki.ling.helsinki.fi/eng/webmasters.html)":
		return &Crawler{
			Pattern:      `heritrix`,
			Instances:    []string{ "Mozilla/5.0 (compatible; heritrix/1.12.1 +http://www.webarchiv.cz)", "Mozilla/5.0 (compatible; heritrix/1.12.1b +http://netarkivet.dk/website/info.html)", "Mozilla/5.0 (compatible; heritrix/1.14.2 +http://rjpower.org)", "Mozilla/5.0 (compatible; heritrix/1.14.2 +http://www.webarchiv.cz)", "Mozilla/5.0 (compatible; heritrix/1.14.3 +http://archive.org)", "Mozilla/5.0 (compatible; heritrix/1.14.3 +http://www.accelobot.com)", "Mozilla/5.0 (compatible; heritrix/1.14.3 +http://www.webarchiv.cz)", "Mozilla/5.0 (compatible; heritrix/1.14.3.r6601 +http://www.buddybuzz.net/yptrino)", "Mozilla/5.0 (compatible; heritrix/1.14.4 +http://parsijoo.ir)", "Mozilla/5.0 (compatible; heritrix/1.14.4 +http://www.exif-search.com)", "Mozilla/5.0 (compatible; heritrix/2.0.2 +http://aihit.com)", "Mozilla/5.0 (compatible; heritrix/2.0.2 +http://seekda.com)", "Mozilla/5.0 (compatible; heritrix/3.0.0-SNAPSHOT-20091120.021634 +http://crawler.archive.org)", "Mozilla/5.0 (compatible; heritrix/3.1.0-RC1 +http://boston.lti.cs.cmu.edu/crawler_12/)", "Mozilla/5.0 (compatible; heritrix/3.1.1 +http://places.tomtom.com/crawlerinfo)", "Mozilla/5.0 (compatible; heritrix/3.1.1 +http://www.mixdata.com)", "Mozilla/5.0 (compatible; heritrix/3.1.1; UniLeipzigASV +http://corpora.informatik.uni-leipzig.de/crawler_faq.html)", "Mozilla/5.0 (compatible; heritrix/3.2.0 +http://www.crim.ca)", "Mozilla/5.0 (compatible; heritrix/3.2.0 +http://www.exif-search.com)", "Mozilla/5.0 (compatible; heritrix/3.2.0 +http://www.mixdata.com)", "Mozilla/5.0 (compatible; heritrix/3.3.0-SNAPSHOT-20160309-0050; UniLeipzigASV +http://corpora.informatik.uni-leipzig.de/crawler_faq.html)", "Mozilla/5.0 (compatible; sukibot_heritrix/3.1.1 +http://suki.ling.helsinki.fi/eng/webmasters.html)" },
			URL:          stringer("https://github.com/internetarchive/heritrix3/wiki"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Mozilla/5.0 (compatible; MSIE 7.0 +http://www.europarchive.org)":
		return &Crawler{
			Pattern:      `europarchive.org`,
			Instances:    []string{ "Mozilla/5.0 (compatible; MSIE 7.0 +http://www.europarchive.org)" },
			URL:          stringer(""),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Mozilla/5.0 (compatible; NerdByNature.Bot; http://www.nerdbynature.net/bot)":
		return &Crawler{
			Pattern:      `NerdByNature.Bot`,
			Instances:    []string{ "Mozilla/5.0 (compatible; NerdByNature.Bot; http://www.nerdbynature.net/bot)" },
			URL:          stringer("http://www.nerdbynature.net/bot"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Mozilla/5.0 (compatible; AhrefsBot/6.1; +http://ahrefs.com/robot/)",
        "Mozilla/5.0 (compatible; AhrefsSiteAudit/6.1; +http://ahrefs.com/robot/)",
        "Mozilla/5.0 (compatible; AhrefsBot/5.2; News; +http://ahrefs.com/robot/)",
        "Mozilla/5.0 (compatible; AhrefsBot/5.2; +http://ahrefs.com/robot/)",
        "Mozilla/5.0 (compatible; AhrefsSiteAudit/5.2; +http://ahrefs.com/robot/)",
        "Mozilla/5.0 (compatible; AhrefsBot/6.1; News; +http://ahrefs.com/robot/)":
		return &Crawler{
			Pattern:      `Ahrefs(Bot|SiteAudit)`,
			Instances:    []string{ "Mozilla/5.0 (compatible; AhrefsBot/6.1; +http://ahrefs.com/robot/)", "Mozilla/5.0 (compatible; AhrefsSiteAudit/6.1; +http://ahrefs.com/robot/)", "Mozilla/5.0 (compatible; AhrefsBot/5.2; News; +http://ahrefs.com/robot/)", "Mozilla/5.0 (compatible; AhrefsBot/5.2; +http://ahrefs.com/robot/)", "Mozilla/5.0 (compatible; AhrefsSiteAudit/5.2; +http://ahrefs.com/robot/)", "Mozilla/5.0 (compatible; AhrefsBot/6.1; News; +http://ahrefs.com/robot/)" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "fuelbot":
		return &Crawler{
			Pattern:      `fuelbot`,
			Instances:    []string{ "fuelbot" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "CrunchBot/1.0 (+http://www.leadcrunch.com/crunchbot)":
		return &Crawler{
			Pattern:      `CrunchBot`,
			Instances:    []string{ "CrunchBot/1.0 (+http://www.leadcrunch.com/crunchbot)" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Mozilla/5.0 (Windows NT 6.1; rv:38.0) Gecko/20100101 Firefox/38.0 (IndeedBot 1.1)":
		return &Crawler{
			Pattern:      `IndeedBot`,
			Instances:    []string{ "Mozilla/5.0 (Windows NT 6.1; rv:38.0) Gecko/20100101 Firefox/38.0 (IndeedBot 1.1)" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Mozilla/5.0 (compatible; Mappy/1.0; +http://mappydata.net/bot/)":
		return &Crawler{
			Pattern:      `mappydata`,
			Instances:    []string{ "Mozilla/5.0 (compatible; Mappy/1.0; +http://mappydata.net/bot/)" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "woobot":
		return &Crawler{
			Pattern:      `woobot`,
			Instances:    []string{ "woobot" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "ZoominfoBot (zoominfobot at zoominfo dot com)":
		return &Crawler{
			Pattern:      `ZoominfoBot`,
			Instances:    []string{ "ZoominfoBot (zoominfobot at zoominfo dot com)" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Mozilla/5.0 (compatible; PrivacyAwareBot/1.1; +http://www.privacyaware.org)":
		return &Crawler{
			Pattern:      `PrivacyAwareBot`,
			Instances:    []string{ "Mozilla/5.0 (compatible; PrivacyAwareBot/1.1; +http://www.privacyaware.org)" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Mozilla/5.0 (compatible; MSIE 9.0; Windows NT 6.1; Multiviewbot":
		return &Crawler{
			Pattern:      `Multiviewbot`,
			Instances:    []string{ "Mozilla/5.0 (compatible; MSIE 9.0; Windows NT 6.1; Multiviewbot" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/45.0.2454.101 Safari/537.36 SWIMGBot":
		return &Crawler{
			Pattern:      `SWIMGBot`,
			Instances:    []string{ "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/45.0.2454.101 Safari/537.36 SWIMGBot" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Mozilla/5.0 (compatible; Grobbot/2.2; +https://grob.it)":
		return &Crawler{
			Pattern:      `Grobbot`,
			Instances:    []string{ "Mozilla/5.0 (compatible; Grobbot/2.2; +https://grob.it)" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Mozilla/5.0 (compatible; eright/1.0; +bot@eright.com)":
		return &Crawler{
			Pattern:      `eright`,
			Instances:    []string{ "Mozilla/5.0 (compatible; eright/1.0; +bot@eright.com)" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Mozilla/5.0 (compatible; Apercite; +http://www.apercite.fr/robot/index.html)":
		return &Crawler{
			Pattern:      `Apercite`,
			Instances:    []string{ "Mozilla/5.0 (compatible; Apercite; +http://www.apercite.fr/robot/index.html)" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "semanticbot",
        "semanticbot (info@semanticaudience.com)":
		return &Crawler{
			Pattern:      `semanticbot`,
			Instances:    []string{ "semanticbot", "semanticbot (info@semanticaudience.com)" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Aboundex/0.2 (http://www.aboundex.com/crawler/)",
        "Aboundex/0.3 (http://www.aboundex.com/crawler/)":
		return &Crawler{
			Pattern:      `Aboundex`,
			Instances:    []string{ "Aboundex/0.2 (http://www.aboundex.com/crawler/)", "Aboundex/0.3 (http://www.aboundex.com/crawler/)" },
			URL:          stringer("http://www.aboundex.com/crawler/"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "CipaCrawler/3.0 (info@domaincrawler.com; http://www.domaincrawler.com/www.example.com)":
		return &Crawler{
			Pattern:      `domaincrawler`,
			Instances:    []string{ "CipaCrawler/3.0 (info@domaincrawler.com; http://www.domaincrawler.com/www.example.com)" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Summify (Summify/1.0.1; +http://summify.com)":
		return &Crawler{
			Pattern:      `summify`,
			Instances:    []string{ "Summify (Summify/1.0.1; +http://summify.com)" },
			URL:          stringer("http://summify.com"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "CCBot/2.0 (http://commoncrawl.org/faq/)",
        "CCBot/2.0 (https://commoncrawl.org/faq/)":
		return &Crawler{
			Pattern:      `CCBot`,
			Instances:    []string{ "CCBot/2.0 (http://commoncrawl.org/faq/)", "CCBot/2.0 (https://commoncrawl.org/faq/)" },
			URL:          stringer("http://www.commoncrawl.org/bot.html"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Mozilla/5.0 (compatible; SeznamBot/3.2-test1-1; +http://napoveda.seznam.cz/en/seznambot-intro/)",
        "Mozilla/5.0 (compatible; SeznamBot/3.2-test1; +http://napoveda.seznam.cz/en/seznambot-intro/)",
        "Mozilla/5.0 (compatible; SeznamBot/3.2-test2; +http://napoveda.seznam.cz/en/seznambot-intro/)",
        "Mozilla/5.0 (compatible; SeznamBot/3.2-test4; +http://napoveda.seznam.cz/en/seznambot-intro/)",
        "Mozilla/5.0 (compatible; SeznamBot/3.2; +http://napoveda.seznam.cz/en/seznambot-intro/)":
		return &Crawler{
			Pattern:      `seznambot`,
			Instances:    []string{ "Mozilla/5.0 (compatible; SeznamBot/3.2-test1-1; +http://napoveda.seznam.cz/en/seznambot-intro/)", "Mozilla/5.0 (compatible; SeznamBot/3.2-test1; +http://napoveda.seznam.cz/en/seznambot-intro/)", "Mozilla/5.0 (compatible; SeznamBot/3.2-test2; +http://napoveda.seznam.cz/en/seznambot-intro/)", "Mozilla/5.0 (compatible; SeznamBot/3.2-test4; +http://napoveda.seznam.cz/en/seznambot-intro/)", "Mozilla/5.0 (compatible; SeznamBot/3.2; +http://napoveda.seznam.cz/en/seznambot-intro/)" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "ec2linkfinder":
		return &Crawler{
			Pattern:      `ec2linkfinder`,
			Instances:    []string{ "ec2linkfinder" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Mozilla/5.0 (compatible; aiHitBot/2.9; +https://www.aihitdata.com/about)":
		return &Crawler{
			Pattern:      `aiHitBot`,
			Instances:    []string{ "Mozilla/5.0 (compatible; aiHitBot/2.9; +https://www.aihitdata.com/about)" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "facebookexternalhit/1.0 (+http://www.facebook.com/externalhit_uatext.php)",
        "facebookexternalhit/1.1",
        "facebookexternalhit/1.1 (+http://www.facebook.com/externalhit_uatext.php)":
		return &Crawler{
			Pattern:      `facebookexternalhit`,
			Instances:    []string{ "facebookexternalhit/1.0 (+http://www.facebook.com/externalhit_uatext.php)", "facebookexternalhit/1.1", "facebookexternalhit/1.1 (+http://www.facebook.com/externalhit_uatext.php)" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Mozilla/5.0 (compatible; Yeti/1.1; +http://naver.me/bot)":
		return &Crawler{
			Pattern:      `Yeti`,
			Instances:    []string{ "Mozilla/5.0 (compatible; Yeti/1.1; +http://naver.me/bot)" },
			URL:          stringer("http://naver.me/bot"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 5.1; RetrevoPageAnalyzer; +http://www.retrevo.com/content/about-us)":
		return &Crawler{
			Pattern:      `RetrevoPageAnalyzer`,
			Instances:    []string{ "Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 5.1; RetrevoPageAnalyzer; +http://www.retrevo.com/content/about-us)" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Sogou News Spider/4.0(+http://www.sogou.com/docs/help/webmasters.htm#07)",
        "Sogou Pic Spider/3.0(+http://www.sogou.com/docs/help/webmasters.htm#07)",
        "Sogou web spider/4.0(+http://www.sogou.com/docs/help/webmasters.htm#07)":
		return &Crawler{
			Pattern:      `Sogou`,
			Instances:    []string{ "Sogou News Spider/4.0(+http://www.sogou.com/docs/help/webmasters.htm#07)", "Sogou Pic Spider/3.0(+http://www.sogou.com/docs/help/webmasters.htm#07)", "Sogou web spider/4.0(+http://www.sogou.com/docs/help/webmasters.htm#07)" },
			URL:          stringer("http://www.sogou.com/docs/help/webmasters.htm#07"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Wotbox/2.0 (bot@wotbox.com; http://www.wotbox.com)",
        "Wotbox/2.01 (+http://www.wotbox.com/bot/)":
		return &Crawler{
			Pattern:      `wotbox`,
			Instances:    []string{ "Wotbox/2.0 (bot@wotbox.com; http://www.wotbox.com)", "Wotbox/2.01 (+http://www.wotbox.com/bot/)" },
			URL:          stringer("http://www.wotbox.com"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "DoCoMo/2.0 P900i(c100;TB;W24H11) (compatible; ichiro/mobile goo; +http://help.goo.ne.jp/help/article/1142/)",
        "DoCoMo/2.0 P900i(c100;TB;W24H11) (compatible; ichiro/mobile goo; +http://search.goo.ne.jp/option/use/sub4/sub4-1/)",
        "DoCoMo/2.0 P900i(c100;TB;W24H11) (compatible; ichiro/mobile goo;+http://search.goo.ne.jp/option/use/sub4/sub4-1/)",
        "DoCoMo/2.0 P900i(c100;TB;W24H11)(compatible; ichiro/mobile goo;+http://help.goo.ne.jp/door/crawler.html)",
        "DoCoMo/2.0 P901i(c100;TB;W24H11) (compatible; ichiro/mobile goo; +http://help.goo.ne.jp/door/crawler.html)",
        "KDDI-CA31 UP.Browser/6.2.0.7.3.129 (GUI) MMP/2.0 (compatible; ichiro/mobile goo; +http://help.goo.ne.jp/help/article/1142/)",
        "KDDI-CA31 UP.Browser/6.2.0.7.3.129 (GUI) MMP/2.0 (compatible; ichiro/mobile goo; +http://search.goo.ne.jp/option/use/sub4/sub4-1/)",
        "KDDI-CA31 UP.Browser/6.2.0.7.3.129 (GUI) MMP/2.0 (compatible; ichiro/mobile goo;+http://search.goo.ne.jp/option/use/sub4/sub4-1/)",
        "ichiro/2.0 (http://help.goo.ne.jp/door/crawler.html)",
        "ichiro/2.0 (ichiro@nttr.co.jp)",
        "ichiro/3.0 (http://help.goo.ne.jp/door/crawler.html)",
        "ichiro/3.0 (http://help.goo.ne.jp/help/article/1142)",
        "ichiro/3.0 (http://search.goo.ne.jp/option/use/sub4/sub4-1/)",
        "ichiro/4.0 (http://help.goo.ne.jp/door/crawler.html)",
        "ichiro/5.0 (http://help.goo.ne.jp/door/crawler.html)":
		return &Crawler{
			Pattern:      `ichiro`,
			Instances:    []string{ "DoCoMo/2.0 P900i(c100;TB;W24H11) (compatible; ichiro/mobile goo; +http://help.goo.ne.jp/help/article/1142/)", "DoCoMo/2.0 P900i(c100;TB;W24H11) (compatible; ichiro/mobile goo; +http://search.goo.ne.jp/option/use/sub4/sub4-1/)", "DoCoMo/2.0 P900i(c100;TB;W24H11) (compatible; ichiro/mobile goo;+http://search.goo.ne.jp/option/use/sub4/sub4-1/)", "DoCoMo/2.0 P900i(c100;TB;W24H11)(compatible; ichiro/mobile goo;+http://help.goo.ne.jp/door/crawler.html)", "DoCoMo/2.0 P901i(c100;TB;W24H11) (compatible; ichiro/mobile goo; +http://help.goo.ne.jp/door/crawler.html)", "KDDI-CA31 UP.Browser/6.2.0.7.3.129 (GUI) MMP/2.0 (compatible; ichiro/mobile goo; +http://help.goo.ne.jp/help/article/1142/)", "KDDI-CA31 UP.Browser/6.2.0.7.3.129 (GUI) MMP/2.0 (compatible; ichiro/mobile goo; +http://search.goo.ne.jp/option/use/sub4/sub4-1/)", "KDDI-CA31 UP.Browser/6.2.0.7.3.129 (GUI) MMP/2.0 (compatible; ichiro/mobile goo;+http://search.goo.ne.jp/option/use/sub4/sub4-1/)", "ichiro/2.0 (http://help.goo.ne.jp/door/crawler.html)", "ichiro/2.0 (ichiro@nttr.co.jp)", "ichiro/3.0 (http://help.goo.ne.jp/door/crawler.html)", "ichiro/3.0 (http://help.goo.ne.jp/help/article/1142)", "ichiro/3.0 (http://search.goo.ne.jp/option/use/sub4/sub4-1/)", "ichiro/4.0 (http://help.goo.ne.jp/door/crawler.html)", "ichiro/5.0 (http://help.goo.ne.jp/door/crawler.html)" },
			URL:          stringer("http://help.goo.ne.jp/help/article/1142"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "DuckDuckBot/1.0; (+http://duckduckgo.com/duckduckbot.html)",
        "DuckDuckBot/1.1; (+http://duckduckgo.com/duckduckbot.html)":
		return &Crawler{
			Pattern:      `DuckDuckBot`,
			Instances:    []string{ "DuckDuckBot/1.0; (+http://duckduckgo.com/duckduckbot.html)", "DuckDuckBot/1.1; (+http://duckduckgo.com/duckduckbot.html)" },
			URL:          stringer("http://duckduckgo.com/duckduckbot.html"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "drupact/0.7; http://www.arocom.de/drupact":
		return &Crawler{
			Pattern:      `drupact`,
			Instances:    []string{ "drupact/0.7; http://www.arocom.de/drupact" },
			URL:          stringer("http://www.arocom.de/drupact"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Mozilla/5.0 (compatible; coccoc/1.0; +http://help.coccoc.com/)",
        "Mozilla/5.0 (compatible; coccoc/1.0; +http://help.coccoc.com/searchengine)",
        "Mozilla/5.0 (compatible; coccocbot-image/1.0; +http://help.coccoc.com/searchengine)",
        "Mozilla/5.0 (compatible; coccocbot-web/1.0; +http://help.coccoc.com/searchengine)",
        "Mozilla/5.0 (compatible; image.coccoc/1.0; +http://help.coccoc.com/)",
        "Mozilla/5.0 (compatible; imagecoccoc/1.0; +http://help.coccoc.com/)",
        "Mozilla/5.0 (compatible; imagecoccoc/1.0; +http://help.coccoc.com/searchengine)",
        "coccoc",
        "coccoc/1.0 ()",
        "coccoc/1.0 (http://help.coccoc.com/)",
        "coccoc/1.0 (http://help.coccoc.vn/)":
		return &Crawler{
			Pattern:      `coccoc`,
			Instances:    []string{ "Mozilla/5.0 (compatible; coccoc/1.0; +http://help.coccoc.com/)", "Mozilla/5.0 (compatible; coccoc/1.0; +http://help.coccoc.com/searchengine)", "Mozilla/5.0 (compatible; coccocbot-image/1.0; +http://help.coccoc.com/searchengine)", "Mozilla/5.0 (compatible; coccocbot-web/1.0; +http://help.coccoc.com/searchengine)", "Mozilla/5.0 (compatible; image.coccoc/1.0; +http://help.coccoc.com/)", "Mozilla/5.0 (compatible; imagecoccoc/1.0; +http://help.coccoc.com/)", "Mozilla/5.0 (compatible; imagecoccoc/1.0; +http://help.coccoc.com/searchengine)", "coccoc", "coccoc/1.0 ()", "coccoc/1.0 (http://help.coccoc.com/)", "coccoc/1.0 (http://help.coccoc.vn/)" },
			URL:          stringer("http://help.coccoc.vn/"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "www.integromedb.org/Crawler":
		return &Crawler{
			Pattern:      `integromedb`,
			Instances:    []string{ "www.integromedb.org/Crawler" },
			URL:          stringer("http://www.integromedb.org/Crawler"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "it2media-domain-crawler/1.0 on crawler-prod.it2media.de",
        "it2media-domain-crawler/2.0":
		return &Crawler{
			Pattern:      `it2media-domain-crawler`,
			Instances:    []string{ "it2media-domain-crawler/1.0 on crawler-prod.it2media.de", "it2media-domain-crawler/2.0" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Mozilla/5.0 (compatible; SiteExplorer/1.0b; +http://siteexplorer.info/)",
        "Mozilla/5.0 (compatible; SiteExplorer/1.1b; +http://siteexplorer.info/Backlink-Checker-Spider/)":
		return &Crawler{
			Pattern:      `siteexplorer.info`,
			Instances:    []string{ "Mozilla/5.0 (compatible; SiteExplorer/1.0b; +http://siteexplorer.info/)", "Mozilla/5.0 (compatible; SiteExplorer/1.1b; +http://siteexplorer.info/Backlink-Checker-Spider/)" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Mozilla/5.0 (compatible; proximic; +http://www.proximic.com)",
        "Mozilla/5.0 (compatible; proximic; +http://www.proximic.com/info/spider.php)":
		return &Crawler{
			Pattern:      `proximic`,
			Instances:    []string{ "Mozilla/5.0 (compatible; proximic; +http://www.proximic.com)", "Mozilla/5.0 (compatible; proximic; +http://www.proximic.com/info/spider.php)" },
			URL:          stringer("http://www.proximic.com/info/spider.php"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1; SV1;  http://www.changedetection.com/bot.html )":
		return &Crawler{
			Pattern:      `changedetection`,
			Instances:    []string{ "Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1; SV1;  http://www.changedetection.com/bot.html )" },
			URL:          stringer("http://www.changedetection.com/bot.html"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "WeSEE:Search",
        "WeSEE:Search/0.1 (Alpha, http://www.wesee.com/en/support/bot/)":
		return &Crawler{
			Pattern:      `WeSEE:Search`,
			Instances:    []string{ "WeSEE:Search", "WeSEE:Search/0.1 (Alpha, http://www.wesee.com/en/support/bot/)" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Mozilla/5.0 (compatible; rogerBot/1.0; UrlCrawler; http://www.seomoz.org/dp/rogerbot)",
        "rogerbot/1.0 (http://moz.com/help/pro/what-is-rogerbot-, rogerbot-crawler+partager@moz.com)",
        "rogerbot/1.0 (http://moz.com/help/pro/what-is-rogerbot-, rogerbot-crawler+shiny@moz.com)",
        "rogerbot/1.0 (http://moz.com/help/pro/what-is-rogerbot-, rogerbot-wherecat@moz.com",
        "rogerbot/1.0 (http://moz.com/help/pro/what-is-rogerbot-, rogerbot-wherecat@moz.com)",
        "rogerbot/1.0 (http://www.moz.com/dp/rogerbot, rogerbot-crawler@moz.com)",
        "rogerbot/1.0 (http://www.seomoz.org/dp/rogerbot, rogerbot-crawler+shiny@seomoz.org)",
        "rogerbot/1.0 (http://www.seomoz.org/dp/rogerbot, rogerbot-crawler@seomoz.org)",
        "rogerbot/1.0 (http://www.seomoz.org/dp/rogerbot, rogerbot-wherecat@moz.com)",
        "rogerbot/1.1 (http://moz.com/help/guides/search-overview/crawl-diagnostics#more-help, rogerbot-crawler+pr2-crawler-05@moz.com)",
        "rogerbot/1.1 (http://moz.com/help/guides/search-overview/crawl-diagnostics#more-help, rogerbot-crawler+pr4-crawler-11@moz.com)",
        "rogerbot/1.1 (http://moz.com/help/guides/search-overview/crawl-diagnostics#more-help, rogerbot-crawler+pr4-crawler-15@moz.com)",
        "rogerbot/1.2 (http://moz.com/help/pro/what-is-rogerbot-, rogerbot-crawler+phaser-testing-crawler-01@moz.com)":
		return &Crawler{
			Pattern:      `rogerbot`,
			Instances:    []string{ "Mozilla/5.0 (compatible; rogerBot/1.0; UrlCrawler; http://www.seomoz.org/dp/rogerbot)", "rogerbot/1.0 (http://moz.com/help/pro/what-is-rogerbot-, rogerbot-crawler+partager@moz.com)", "rogerbot/1.0 (http://moz.com/help/pro/what-is-rogerbot-, rogerbot-crawler+shiny@moz.com)", "rogerbot/1.0 (http://moz.com/help/pro/what-is-rogerbot-, rogerbot-wherecat@moz.com", "rogerbot/1.0 (http://moz.com/help/pro/what-is-rogerbot-, rogerbot-wherecat@moz.com)", "rogerbot/1.0 (http://www.moz.com/dp/rogerbot, rogerbot-crawler@moz.com)", "rogerbot/1.0 (http://www.seomoz.org/dp/rogerbot, rogerbot-crawler+shiny@seomoz.org)", "rogerbot/1.0 (http://www.seomoz.org/dp/rogerbot, rogerbot-crawler@seomoz.org)", "rogerbot/1.0 (http://www.seomoz.org/dp/rogerbot, rogerbot-wherecat@moz.com)", "rogerbot/1.1 (http://moz.com/help/guides/search-overview/crawl-diagnostics#more-help, rogerbot-crawler+pr2-crawler-05@moz.com)", "rogerbot/1.1 (http://moz.com/help/guides/search-overview/crawl-diagnostics#more-help, rogerbot-crawler+pr4-crawler-11@moz.com)", "rogerbot/1.1 (http://moz.com/help/guides/search-overview/crawl-diagnostics#more-help, rogerbot-crawler+pr4-crawler-15@moz.com)", "rogerbot/1.2 (http://moz.com/help/pro/what-is-rogerbot-, rogerbot-crawler+phaser-testing-crawler-01@moz.com)" },
			URL:          stringer("http://moz.com/help/pro/what-is-rogerbot-"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.1 (KHTML, like Gecko) Chrome/21.0.1180.89 Safari/537.1; 360Spider",
        "Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.1 (KHTML, like Gecko) Chrome/21.0.1180.89 Safari/537.1; 360Spider(compatible; HaosouSpider; http://www.haosou.com/help/help_3_2.html)",
        "Mozilla/5.0 (Windows NT 6.2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/31.0.1650.63 Safari/537.36 QIHU 360SE; 360Spider",
        "Mozilla/5.0 (Windows; U; Windows NT 5.1; zh-CN; )  Firefox/1.5.0.11; 360Spider",
        "Mozilla/5.0 (Windows; U; Windows NT 5.1; zh-CN; rv:1.8.0.11)  Firefox/1.5.0.11; 360Spider",
        "Mozilla/5.0 (Windows; U; Windows NT 5.1; zh-CN; rv:1.8.0.11) Firefox/1.5.0.11 360Spider;",
        "Mozilla/5.0 (Windows; U; Windows NT 5.1; zh-CN; rv:1.8.0.11) Gecko/20070312 Firefox/1.5.0.11; 360Spider",
        "Mozilla/5.0 (compatible; MSIE 9.0; Windows NT 6.1; Trident/5.0); 360Spider",
        "Mozilla/5.0 (compatible; MSIE 9.0; Windows NT 6.1; Trident/5.0); 360Spider(compatible; HaosouSpider; http://www.haosou.com/help/help_3_2.html)":
		return &Crawler{
			Pattern:      `360Spider`,
			Instances:    []string{ "Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.1 (KHTML, like Gecko) Chrome/21.0.1180.89 Safari/537.1; 360Spider", "Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.1 (KHTML, like Gecko) Chrome/21.0.1180.89 Safari/537.1; 360Spider(compatible; HaosouSpider; http://www.haosou.com/help/help_3_2.html)", "Mozilla/5.0 (Windows NT 6.2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/31.0.1650.63 Safari/537.36 QIHU 360SE; 360Spider", "Mozilla/5.0 (Windows; U; Windows NT 5.1; zh-CN; )  Firefox/1.5.0.11; 360Spider", "Mozilla/5.0 (Windows; U; Windows NT 5.1; zh-CN; rv:1.8.0.11)  Firefox/1.5.0.11; 360Spider", "Mozilla/5.0 (Windows; U; Windows NT 5.1; zh-CN; rv:1.8.0.11) Firefox/1.5.0.11 360Spider;", "Mozilla/5.0 (Windows; U; Windows NT 5.1; zh-CN; rv:1.8.0.11) Gecko/20070312 Firefox/1.5.0.11; 360Spider", "Mozilla/5.0 (compatible; MSIE 9.0; Windows NT 6.1; Trident/5.0); 360Spider", "Mozilla/5.0 (compatible; MSIE 9.0; Windows NT 6.1; Trident/5.0); 360Spider(compatible; HaosouSpider; http://www.haosou.com/help/help_3_2.html)" },
			URL:          stringer("http://needs-be.blogspot.co.uk/2013/02/how-to-block-spider360.html"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "psbot-image (+http://www.picsearch.com/bot.html)",
        "psbot-page (+http://www.picsearch.com/bot.html)",
        "psbot/0.1 (+http://www.picsearch.com/bot.html)":
		return &Crawler{
			Pattern:      `psbot`,
			Instances:    []string{ "psbot-image (+http://www.picsearch.com/bot.html)", "psbot-page (+http://www.picsearch.com/bot.html)", "psbot/0.1 (+http://www.picsearch.com/bot.html)" },
			URL:          stringer("http://www.picsearch.com/bot.html"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "CC Metadata Scaper http://wiki.creativecommons.org/Metadata_Scraper":
		return &Crawler{
			Pattern:      `CC Metadata Scaper`,
			Instances:    []string{ "CC Metadata Scaper http://wiki.creativecommons.org/Metadata_Scraper" },
			URL:          stringer("http://wiki.creativecommons.org/Metadata_Scraper"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Mozilla/5.0 (compatible; GrapeshotCrawler/2.0; +http://www.grapeshot.co.uk/crawler.php)":
		return &Crawler{
			Pattern:      `GrapeshotCrawler`,
			Instances:    []string{ "Mozilla/5.0 (compatible; GrapeshotCrawler/2.0; +http://www.grapeshot.co.uk/crawler.php)" },
			URL:          stringer("http://www.grapeshot.co.uk/crawler.php"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Mozilla/5.0 (compatible; URLAppendBot/1.0; +http://www.profound.net/urlappendbot.html)":
		return &Crawler{
			Pattern:      `urlappendbot`,
			Instances:    []string{ "Mozilla/5.0 (compatible; URLAppendBot/1.0; +http://www.profound.net/urlappendbot.html)" },
			URL:          stringer("http://www.profound.net/urlappendbot.html"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Mozilla/5.0 (compatible; fr-crawler/1.1)":
		return &Crawler{
			Pattern:      `fr-crawler`,
			Instances:    []string{ "Mozilla/5.0 (compatible; fr-crawler/1.1)" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "binlar_2.6.3 binlar2.6.3@unspecified.mail",
        "binlar_2.6.3 binlar_2.6.3@unspecified.mail",
        "binlar_2.6.3 larbin2.6.3@unspecified.mail",
        "binlar_2.6.3 phanendra_kalapala@McAfee.com",
        "binlar_2.6.3 test@mgmt.mic":
		return &Crawler{
			Pattern:      `binlar`,
			Instances:    []string{ "binlar_2.6.3 binlar2.6.3@unspecified.mail", "binlar_2.6.3 binlar_2.6.3@unspecified.mail", "binlar_2.6.3 larbin2.6.3@unspecified.mail", "binlar_2.6.3 phanendra_kalapala@McAfee.com", "binlar_2.6.3 test@mgmt.mic" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "SimpleCrawler/0.1":
		return &Crawler{
			Pattern:      `SimpleCrawler`,
			Instances:    []string{ "SimpleCrawler/0.1" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Twitterbot/0.1",
        "Twitterbot/1.0":
		return &Crawler{
			Pattern:      `Twitterbot`,
			Instances:    []string{ "Twitterbot/0.1", "Twitterbot/1.0" },
			URL:          stringer("https://dev.twitter.com/cards/getting-started"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "cXensebot/1.1a":
		return &Crawler{
			Pattern:      `cXensebot`,
			Instances:    []string{ "cXensebot/1.1a" },
			URL:          stringer("http://www.cxense.com/bot.html"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Mozilla/5.0 (compatible; SMTBot/1.0; +http://www.similartech.com/smtbot)",
        "SMTBot (similartech.com/smtbot)":
		return &Crawler{
			Pattern:      `smtbot`,
			Instances:    []string{ "Mozilla/5.0 (compatible; SMTBot/1.0; +http://www.similartech.com/smtbot)", "SMTBot (similartech.com/smtbot)" },
			URL:          stringer("http://www.similartech.com/smtbot"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Mozilla/5.0 (compatible; bnf.fr_bot; +http://bibnum.bnf.fr/robot/bnf.html)",
        "Mozilla/5.0 (compatible; bnf.fr_bot; +http://www.bnf.fr/fr/outils/a.dl_web_capture_robot.html)":
		return &Crawler{
			Pattern:      `bnf.fr_bot`,
			Instances:    []string{ "Mozilla/5.0 (compatible; bnf.fr_bot; +http://bibnum.bnf.fr/robot/bnf.html)", "Mozilla/5.0 (compatible; bnf.fr_bot; +http://www.bnf.fr/fr/outils/a.dl_web_capture_robot.html)" },
			URL:          stringer("http://www.bnf.fr/fr/outils/a.dl_web_capture_robot.html"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "A6-Indexer":
		return &Crawler{
			Pattern:      `A6-Indexer`,
			Instances:    []string{ "A6-Indexer" },
			URL:          stringer("http://www.a6corp.com/a6-web-scraping-policy/"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "ADmantX Platform Semantic Analyzer - ADmantX Inc. - www.admantx.com - support@admantx.com":
		return &Crawler{
			Pattern:      `ADmantX`,
			Instances:    []string{ "ADmantX Platform Semantic Analyzer - ADmantX Inc. - www.admantx.com - support@admantx.com" },
			URL:          stringer("http://www.admantx.com"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Facebot/1.0":
		return &Crawler{
			Pattern:      `Facebot`,
			Instances:    []string{ "Facebot/1.0" },
			URL:          stringer("https://developers.facebook.com/docs/sharing/best-practices#crawl"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Mozilla/5.0 (compatible; OrangeBot/2.0; support.orangebot@orange.com":
		return &Crawler{
			Pattern:      `OrangeBot\/`,
			Instances:    []string{ "Mozilla/5.0 (compatible; OrangeBot/2.0; support.orangebot@orange.com" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Mozilla/5.0 (compatible; memorybot/1.21.14 +http://mignify.com/bot.html)":
		return &Crawler{
			Pattern:      `memorybot`,
			Instances:    []string{ "Mozilla/5.0 (compatible; memorybot/1.21.14 +http://mignify.com/bot.html)" },
			URL:          stringer("http://mignify.com/bot.htm"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Mozilla/5.0 (compatible; AdvBot/2.0; +http://advbot.net/bot.html)":
		return &Crawler{
			Pattern:      `AdvBot`,
			Instances:    []string{ "Mozilla/5.0 (compatible; AdvBot/2.0; +http://advbot.net/bot.html)" },
			URL:          stringer("http://advbot.net/bot.html"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Mozilla/5.0 (compatible; MegaIndex.ru/2.0; +https://www.megaindex.ru/?tab=linkAnalyze)":
		return &Crawler{
			Pattern:      `MegaIndex`,
			Instances:    []string{ "Mozilla/5.0 (compatible; MegaIndex.ru/2.0; +https://www.megaindex.ru/?tab=linkAnalyze)" },
			URL:          stringer("https://www.megaindex.ru/?tab=linkAnalyze"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "SemanticScholarBot/1.0 (+http://s2.allenai.org/bot.html)":
		return &Crawler{
			Pattern:      `SemanticScholarBot`,
			Instances:    []string{ "SemanticScholarBot/1.0 (+http://s2.allenai.org/bot.html)" },
			URL:          stringer("http://s2.allenai.org/bot.html"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "ltx71 - (http://ltx71.com/)":
		return &Crawler{
			Pattern:      `ltx71`,
			Instances:    []string{ "ltx71 - (http://ltx71.com/)" },
			URL:          stringer("http://ltx71.com/"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "nerdybot":
		return &Crawler{
			Pattern:      `nerdybot`,
			Instances:    []string{ "nerdybot" },
			URL:          stringer("http://nerdybot.com/"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Mozilla/5.0 (compatible; XoviBot/2.0; +http://www.xovibot.net/)":
		return &Crawler{
			Pattern:      `xovibot`,
			Instances:    []string{ "Mozilla/5.0 (compatible; XoviBot/2.0; +http://www.xovibot.net/)" },
			URL:          stringer("http://www.xovibot.net/"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "BUbiNG (+http://law.di.unimi.it/BUbiNG.html)":
		return &Crawler{
			Pattern:      `BUbiNG`,
			Instances:    []string{ "BUbiNG (+http://law.di.unimi.it/BUbiNG.html)" },
			URL:          stringer("http://law.di.unimi.it/BUbiNG.html"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Mozilla/5.0 (compatible; Qwantify/2.0n; +https://www.qwant.com/)/*",
        "Mozilla/5.0 (compatible; Qwantify/2.4w; +https://www.qwant.com/)/2.4w":
		return &Crawler{
			Pattern:      `Qwantify`,
			Instances:    []string{ "Mozilla/5.0 (compatible; Qwantify/2.0n; +https://www.qwant.com/)/*", "Mozilla/5.0 (compatible; Qwantify/2.4w; +https://www.qwant.com/)/2.4w" },
			URL:          stringer("https://www.qwant.com/"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Mozilla/5.0 (compatible; heritrix/3.1.1-SNAPSHOT-20120116.200628 +http://www.archive.org/details/archive.org_bot)",
        "Mozilla/5.0 (compatible; archive.org_bot/heritrix-1.15.4 +http://www.archive.org)",
        "Mozilla/5.0 (compatible; heritrix/3.3.0-SNAPSHOT-20140702-2247 +http://archive.org/details/archive.org_bot)",
        "Mozilla/5.0 (compatible; archive.org_bot +http://www.archive.org/details/archive.org_bot)":
		return &Crawler{
			Pattern:      `archive.org_bot`,
			Instances:    []string{ "Mozilla/5.0 (compatible; heritrix/3.1.1-SNAPSHOT-20120116.200628 +http://www.archive.org/details/archive.org_bot)", "Mozilla/5.0 (compatible; archive.org_bot/heritrix-1.15.4 +http://www.archive.org)", "Mozilla/5.0 (compatible; heritrix/3.3.0-SNAPSHOT-20140702-2247 +http://archive.org/details/archive.org_bot)", "Mozilla/5.0 (compatible; archive.org_bot +http://www.archive.org/details/archive.org_bot)" },
			URL:          stringer("http://www.archive.org/details/archive.org_bot"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{ "heritrix" },
		}

	case "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_1) AppleWebKit/600.2.5 (KHTML, like Gecko) Version/8.0.2 Safari/600.2.5 (Applebot/0.1)",
        "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_1) AppleWebKit/600.2.5 (KHTML, like Gecko) Version/8.0.2 Safari/600.2.5 (Applebot/0.1; +http://www.apple.com/go/applebot)",
        "Mozilla/5.0 (compatible; Applebot/0.3; +http://www.apple.com/go/applebot)",
        "Mozilla/5.0 (iPhone; CPU iPhone OS 6_0 like Mac OS X) AppleWebKit/536.26 (KHTML, like Gecko) Version/6.0 Mobile/10A5376e Safari/8536.25 (compatible; Applebot/0.3; +http://www.apple.com/go/applebot)",
        "Mozilla/5.0 (iPhone; CPU iPhone OS 8_1 like Mac OS X) AppleWebKit/600.1.4 (KHTML, like Gecko) Version/8.0 Mobile/12B410 Safari/600.1.4 (Applebot/0.1; +http://www.apple.com/go/applebot)":
		return &Crawler{
			Pattern:      `Applebot`,
			Instances:    []string{ "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_1) AppleWebKit/600.2.5 (KHTML, like Gecko) Version/8.0.2 Safari/600.2.5 (Applebot/0.1)", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_1) AppleWebKit/600.2.5 (KHTML, like Gecko) Version/8.0.2 Safari/600.2.5 (Applebot/0.1; +http://www.apple.com/go/applebot)", "Mozilla/5.0 (compatible; Applebot/0.3; +http://www.apple.com/go/applebot)", "Mozilla/5.0 (iPhone; CPU iPhone OS 6_0 like Mac OS X) AppleWebKit/536.26 (KHTML, like Gecko) Version/6.0 Mobile/10A5376e Safari/8536.25 (compatible; Applebot/0.3; +http://www.apple.com/go/applebot)", "Mozilla/5.0 (iPhone; CPU iPhone OS 8_1 like Mac OS X) AppleWebKit/600.1.4 (KHTML, like Gecko) Version/8.0 Mobile/12B410 Safari/600.1.4 (Applebot/0.1; +http://www.apple.com/go/applebot)" },
			URL:          stringer("http://www.apple.com/go/applebot"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Mozilla/5.0 (TweetmemeBot/4.0; +http://datasift.com/bot.html) Gecko/20100101 Firefox/31.0":
		return &Crawler{
			Pattern:      `TweetmemeBot`,
			Instances:    []string{ "Mozilla/5.0 (TweetmemeBot/4.0; +http://datasift.com/bot.html) Gecko/20100101 Firefox/31.0" },
			URL:          stringer("http://datasift.com/bot.html"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "crawler4j (http://code.google.com/p/crawler4j/)":
		return &Crawler{
			Pattern:      `crawler4j`,
			Instances:    []string{ "crawler4j (http://code.google.com/p/crawler4j/)" },
			URL:          stringer("https://github.com/yasserg/crawler4j"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Mozilla/5.0 (compatible; Findxbot/1.0; +http://www.findxbot.com)":
		return &Crawler{
			Pattern:      `findxbot`,
			Instances:    []string{ "Mozilla/5.0 (compatible; Findxbot/1.0; +http://www.findxbot.com)" },
			URL:          stringer("http://www.findxbot.com"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Mozilla/5.0 (compatible; SemrushBot-SA/0.97; +http://www.semrush.com/bot.html)",
        "Mozilla/5.0 (compatible; SemrushBot-SI/0.97; +http://www.semrush.com/bot.html)",
        "Mozilla/5.0 (compatible; SemrushBot/3~bl; +http://www.semrush.com/bot.html)",
        "Mozilla/5.0 (compatible; SemrushBot/0.98~bl; +http://www.semrush.com/bot.html)",
        "Mozilla/5.0 (compatible; SemrushBot-BA; +http://www.semrush.com/bot.html)",
        "SEMrushBot":
		return &Crawler{
			Pattern:      `S[eE][mM]rushBot`,
			Instances:    []string{ "Mozilla/5.0 (compatible; SemrushBot-SA/0.97; +http://www.semrush.com/bot.html)", "Mozilla/5.0 (compatible; SemrushBot-SI/0.97; +http://www.semrush.com/bot.html)", "Mozilla/5.0 (compatible; SemrushBot/3~bl; +http://www.semrush.com/bot.html)", "Mozilla/5.0 (compatible; SemrushBot/0.98~bl; +http://www.semrush.com/bot.html)", "Mozilla/5.0 (compatible; SemrushBot-BA; +http://www.semrush.com/bot.html)", "SEMrushBot" },
			URL:          stringer("http://www.semrush.com/bot.html"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Mozilla/5.0 (compatible; yoozBot-2.2; http://yooz.ir; info@yooz.ir)":
		return &Crawler{
			Pattern:      `yoozBot`,
			Instances:    []string{ "Mozilla/5.0 (compatible; yoozBot-2.2; http://yooz.ir; info@yooz.ir)" },
			URL:          stringer("http://yooz.ir"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Mozilla/5.0 (compatible; Lipperhey Link Explorer; http://www.lipperhey.com/)",
        "Mozilla/5.0 (compatible; Lipperhey SEO Service; http://www.lipperhey.com/)",
        "Mozilla/5.0 (compatible; Lipperhey Site Explorer; http://www.lipperhey.com/)",
        "Mozilla/5.0 (compatible; Lipperhey-Kaus-Australis/5.0; +https://www.lipperhey.com/en/about/)":
		return &Crawler{
			Pattern:      `lipperhey`,
			Instances:    []string{ "Mozilla/5.0 (compatible; Lipperhey Link Explorer; http://www.lipperhey.com/)", "Mozilla/5.0 (compatible; Lipperhey SEO Service; http://www.lipperhey.com/)", "Mozilla/5.0 (compatible; Lipperhey Site Explorer; http://www.lipperhey.com/)", "Mozilla/5.0 (compatible; Lipperhey-Kaus-Australis/5.0; +https://www.lipperhey.com/en/about/)" },
			URL:          stringer("http://www.lipperhey.com/"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Y!J-ASR/0.1 crawler (http://www.yahoo-help.jp/app/answers/detail/p/595/a_id/42716/)",
        "Y!J-BRJ/YATS crawler (http://help.yahoo.co.jp/help/jp/search/indexing/indexing-15.html)",
        "Y!J-PSC/1.0 crawler (http://help.yahoo.co.jp/help/jp/search/indexing/indexing-15.html)",
        "Y!J-BRW/1.0 crawler (http://help.yahoo.co.jp/help/jp/search/indexing/indexing-15.html)",
        "Mozilla/5.0 (iPhone; Y!J-BRY/YATSH crawler; http://help.yahoo.co.jp/help/jp/search/indexing/indexing-15.html)",
        "Mozilla/5.0 (compatible; Y!J SearchMonkey/1.0 (Y!J-AGENT; http://help.yahoo.co.jp/help/jp/search/indexing/indexing-15.html))":
		return &Crawler{
			Pattern:      `Y!J`,
			Instances:    []string{ "Y!J-ASR/0.1 crawler (http://www.yahoo-help.jp/app/answers/detail/p/595/a_id/42716/)", "Y!J-BRJ/YATS crawler (http://help.yahoo.co.jp/help/jp/search/indexing/indexing-15.html)", "Y!J-PSC/1.0 crawler (http://help.yahoo.co.jp/help/jp/search/indexing/indexing-15.html)", "Y!J-BRW/1.0 crawler (http://help.yahoo.co.jp/help/jp/search/indexing/indexing-15.html)", "Mozilla/5.0 (iPhone; Y!J-BRY/YATSH crawler; http://help.yahoo.co.jp/help/jp/search/indexing/indexing-15.html)", "Mozilla/5.0 (compatible; Y!J SearchMonkey/1.0 (Y!J-AGENT; http://help.yahoo.co.jp/help/jp/search/indexing/indexing-15.html))" },
			URL:          stringer("https://www.yahoo-help.jp/app/answers/detail/p/595/a_id/42716/~/%E3%82%A6%E3%82%A7%E3%83%96%E3%83%9A%E3%83%BC%E3%82%B8%E3%81%AB%E3%82%A2%E3%82%AF%E3%82%BB%E3%82%B9%E3%81%99%E3%82%8B%E3%82%B7%E3%82%B9%E3%83%86%E3%83%A0%E3%81%AE%E3%83%A6%E3%83%BC%E3%82%B6%E3%83%BC%E3%82%A8%E3%83%BC%E3%82%B8%E3%82%A7%E3%83%B3%E3%83%88%E3%81%AB%E3%81%A4%E3%81%84%E3%81%A6"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Domain Re-Animator Bot (http://domainreanimator.com) - support@domainreanimator.com":
		return &Crawler{
			Pattern:      `Domain Re-Animator Bot`,
			Instances:    []string{ "Domain Re-Animator Bot (http://domainreanimator.com) - support@domainreanimator.com" },
			URL:          stringer("http://domainreanimator.com"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "AddThis.com robot tech.support@clearspring.com":
		return &Crawler{
			Pattern:      `AddThis`,
			Instances:    []string{ "AddThis.com robot tech.support@clearspring.com" },
			URL:          stringer("https://www.addthis.com"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Screaming Frog SEO Spider/5.1":
		return &Crawler{
			Pattern:      `Screaming Frog SEO Spider`,
			Instances:    []string{ "Screaming Frog SEO Spider/5.1" },
			URL:          stringer("http://www.screamingfrog.co.uk/seo-spider"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "MetaURI API/2.0 +metauri.com":
		return &Crawler{
			Pattern:      `MetaURI`,
			Instances:    []string{ "MetaURI API/2.0 +metauri.com" },
			URL:          stringer("http://www.useragentstring.com/MetaURI_id_17683.php"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Scrapy/1.0.3 (+http://scrapy.org)":
		return &Crawler{
			Pattern:      `Scrapy`,
			Instances:    []string{ "Scrapy/1.0.3 (+http://scrapy.org)" },
			URL:          stringer("http://scrapy.org/"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "LivelapBot/0.2 (http://site.livelap.com/crawler)",
        "Livelapbot/0.1":
		return &Crawler{
			Pattern:      `Livelap[bB]ot`,
			Instances:    []string{ "LivelapBot/0.2 (http://site.livelap.com/crawler)", "Livelapbot/0.1" },
			URL:          stringer("http://site.livelap.com/crawler"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Mozilla/5.0 (compatible; OpenHoseBot/2.1; +http://www.openhose.org/bot.html)":
		return &Crawler{
			Pattern:      `OpenHoseBot`,
			Instances:    []string{ "Mozilla/5.0 (compatible; OpenHoseBot/2.1; +http://www.openhose.org/bot.html)" },
			URL:          stringer("http://www.openhose.org/bot.html"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "CapsuleChecker (http://www.capsulink.com/)":
		return &Crawler{
			Pattern:      `CapsuleChecker`,
			Instances:    []string{ "CapsuleChecker (http://www.capsulink.com/)" },
			URL:          stringer("http://www.capsulink.com/about"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Mozilla/5.0 (compatible) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/47.0.2526.73 Safari/537.36 collection@infegy.com":
		return &Crawler{
			Pattern:      `collection@infegy.com`,
			Instances:    []string{ "Mozilla/5.0 (compatible) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/47.0.2526.73 Safari/537.36 collection@infegy.com" },
			URL:          stringer("http://infegy.com/"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Mozilla/5.0 (compatible; IstellaBot/1.23.15 +http://www.tiscali.it/)":
		return &Crawler{
			Pattern:      `IstellaBot`,
			Instances:    []string{ "Mozilla/5.0 (compatible; IstellaBot/1.23.15 +http://www.tiscali.it/)" },
			URL:          stringer("http://www.tiscali.it/"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Mozilla/5.0 (compatible; DeuSu/0.1.0; +https://deusu.org)",
        "Mozilla/5.0 (compatible; DeuSu/5.0.2; +https://deusu.de/robot.html)":
		return &Crawler{
			Pattern:      `DeuSu\/`,
			Instances:    []string{ "Mozilla/5.0 (compatible; DeuSu/0.1.0; +https://deusu.org)", "Mozilla/5.0 (compatible; DeuSu/5.0.2; +https://deusu.de/robot.html)" },
			URL:          stringer("https://deusu.de/robot.html"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Mozilla/5.0 (compatible; Cliqzbot/2.0; +http://cliqz.com/company/cliqzbot)",
        "Cliqzbot/0.1 (+http://cliqz.com +cliqzbot@cliqz.com)",
        "Cliqzbot/0.1 (+http://cliqz.com/company/cliqzbot)",
        "Mozilla/5.0 (compatible; Cliqzbot/0.1 +http://cliqz.com/company/cliqzbot)",
        "Mozilla/5.0 (compatible; Cliqzbot/1.0 +http://cliqz.com/company/cliqzbot)":
		return &Crawler{
			Pattern:      `Cliqzbot\/`,
			Instances:    []string{ "Mozilla/5.0 (compatible; Cliqzbot/2.0; +http://cliqz.com/company/cliqzbot)", "Cliqzbot/0.1 (+http://cliqz.com +cliqzbot@cliqz.com)", "Cliqzbot/0.1 (+http://cliqz.com/company/cliqzbot)", "Mozilla/5.0 (compatible; Cliqzbot/0.1 +http://cliqz.com/company/cliqzbot)", "Mozilla/5.0 (compatible; Cliqzbot/1.0 +http://cliqz.com/company/cliqzbot)" },
			URL:          stringer("http://cliqz.com/company/cliqzbot"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "MojeekBot/0.2 (archi; http://www.mojeek.com/bot.html)",
        "Mozilla/5.0 (compatible; MojeekBot/0.2; http://www.mojeek.com/bot.html#relaunch)",
        "Mozilla/5.0 (compatible; MojeekBot/0.2; http://www.mojeek.com/bot.html)",
        "Mozilla/5.0 (compatible; MojeekBot/0.5; http://www.mojeek.com/bot.html)",
        "Mozilla/5.0 (compatible; MojeekBot/0.6; +https://www.mojeek.com/bot.html)",
        "Mozilla/5.0 (compatible; MojeekBot/0.6; http://www.mojeek.com/bot.html)":
		return &Crawler{
			Pattern:      `MojeekBot\/`,
			Instances:    []string{ "MojeekBot/0.2 (archi; http://www.mojeek.com/bot.html)", "Mozilla/5.0 (compatible; MojeekBot/0.2; http://www.mojeek.com/bot.html#relaunch)", "Mozilla/5.0 (compatible; MojeekBot/0.2; http://www.mojeek.com/bot.html)", "Mozilla/5.0 (compatible; MojeekBot/0.5; http://www.mojeek.com/bot.html)", "Mozilla/5.0 (compatible; MojeekBot/0.6; +https://www.mojeek.com/bot.html)", "Mozilla/5.0 (compatible; MojeekBot/0.6; http://www.mojeek.com/bot.html)" },
			URL:          stringer("https://www.mojeek.com/bot.html"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "netEstate NE Crawler (+http://www.sengine.info/)",
        "netEstate NE Crawler (+http://www.website-datenbank.de/)":
		return &Crawler{
			Pattern:      `netEstate NE Crawler`,
			Instances:    []string{ "netEstate NE Crawler (+http://www.sengine.info/)", "netEstate NE Crawler (+http://www.website-datenbank.de/)" },
			URL:          stringer("+http://www.website-datenbank.de/"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "SafeSearch microdata crawler (https://safesearch.avira.com, safesearch-abuse@avira.com)":
		return &Crawler{
			Pattern:      `SafeSearch microdata crawler`,
			Instances:    []string{ "SafeSearch microdata crawler (https://safesearch.avira.com, safesearch-abuse@avira.com)" },
			URL:          stringer("https://safesearch.avira.com"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Mozilla/5.0 (compatible; Gluten Free Crawler/1.0; +http://glutenfreepleasure.com/)":
		return &Crawler{
			Pattern:      `Gluten Free Crawler\/`,
			Instances:    []string{ "Mozilla/5.0 (compatible; Gluten Free Crawler/1.0; +http://glutenfreepleasure.com/)" },
			URL:          stringer("http://glutenfreepleasure.com/"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Mozilla/5.0 (compatible; RankSonicSiteAuditor/1.0; +https://ranksonic.com/ranksonic_sab.html)",
        "Mozilla/5.0 (compatible; Sonic/1.0; http://www.yama.info.waseda.ac.jp/~crawler/info.html)",
        "Mozzila/5.0 (compatible; Sonic/1.0; http://www.yama.info.waseda.ac.jp/~crawler/info.html)":
		return &Crawler{
			Pattern:      `Sonic`,
			Instances:    []string{ "Mozilla/5.0 (compatible; RankSonicSiteAuditor/1.0; +https://ranksonic.com/ranksonic_sab.html)", "Mozilla/5.0 (compatible; Sonic/1.0; http://www.yama.info.waseda.ac.jp/~crawler/info.html)", "Mozzila/5.0 (compatible; Sonic/1.0; http://www.yama.info.waseda.ac.jp/~crawler/info.html)" },
			URL:          stringer("http://www.yama.info.waseda.ac.jp/~crawler/info.html"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Mozilla/5.0 (compatible; Sysomos/1.0; +http://www.sysomos.com/; Sysomos)":
		return &Crawler{
			Pattern:      `Sysomos`,
			Instances:    []string{ "Mozilla/5.0 (compatible; Sysomos/1.0; +http://www.sysomos.com/; Sysomos)" },
			URL:          stringer("http://www.sysomos.com"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "www.deadlinkchecker.com Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/46.0.2490.86 Safari/537.36",
        "www.deadlinkchecker.com XMLHTTP/1.0",
        "www.deadlinkchecker.com XMLHTTP/1.0 Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/46.0.2490.86 Safari/537.36":
		return &Crawler{
			Pattern:      `deadlinkchecker`,
			Instances:    []string{ "www.deadlinkchecker.com Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/46.0.2490.86 Safari/537.36", "www.deadlinkchecker.com XMLHTTP/1.0", "www.deadlinkchecker.com XMLHTTP/1.0 Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/46.0.2490.86 Safari/537.36" },
			URL:          stringer("http://www.deadlinkchecker.com"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Slack-ImgProxy (+https://api.slack.com/robots)",
        "Slack-ImgProxy 0.59 (+https://api.slack.com/robots)",
        "Slack-ImgProxy 0.66 (+https://api.slack.com/robots)",
        "Slack-ImgProxy 1.106 (+https://api.slack.com/robots)",
        "Slack-ImgProxy 1.138 (+https://api.slack.com/robots)",
        "Slack-ImgProxy 149 (+https://api.slack.com/robots)":
		return &Crawler{
			Pattern:      `Slack-ImgProxy`,
			Instances:    []string{ "Slack-ImgProxy (+https://api.slack.com/robots)", "Slack-ImgProxy 0.59 (+https://api.slack.com/robots)", "Slack-ImgProxy 0.66 (+https://api.slack.com/robots)", "Slack-ImgProxy 1.106 (+https://api.slack.com/robots)", "Slack-ImgProxy 1.138 (+https://api.slack.com/robots)", "Slack-ImgProxy 149 (+https://api.slack.com/robots)" },
			URL:          stringer("https://api.slack.com/robots"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Embedly +support@embed.ly",
        "Mozilla/5.0 (compatible; Embedly/0.2; +http://support.embed.ly/)",
        "Mozilla/5.0 (compatible; Embedly/0.2; snap; +http://support.embed.ly/)":
		return &Crawler{
			Pattern:      `Embedly`,
			Instances:    []string{ "Embedly +support@embed.ly", "Mozilla/5.0 (compatible; Embedly/0.2; +http://support.embed.ly/)", "Mozilla/5.0 (compatible; Embedly/0.2; snap; +http://support.embed.ly/)" },
			URL:          stringer("http://support.embed.ly"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Mozilla/5.0 (compatible; RankActiveLinkBot; +https://rankactive.com/resources/rankactive-linkbot)":
		return &Crawler{
			Pattern:      `RankActiveLinkBot`,
			Instances:    []string{ "Mozilla/5.0 (compatible; RankActiveLinkBot; +https://rankactive.com/resources/rankactive-linkbot)" },
			URL:          stringer("https://rankactive.com/resources/rankactive-linkbot"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "iskanie (+http://www.iskanie.com)":
		return &Crawler{
			Pattern:      `iskanie`,
			Instances:    []string{ "iskanie (+http://www.iskanie.com)" },
			URL:          stringer("http://www.iskanie.com"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "SafeDNSBot (https://www.safedns.com/searchbot)":
		return &Crawler{
			Pattern:      `SafeDNSBot`,
			Instances:    []string{ "SafeDNSBot (https://www.safedns.com/searchbot)" },
			URL:          stringer("https://www.safedns.com/searchbot"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Mozilla/5.0 (Windows NT 6.1; WOW64) SkypeUriPreview Preview/0.5":
		return &Crawler{
			Pattern:      `SkypeUriPreview`,
			Instances:    []string{ "Mozilla/5.0 (Windows NT 6.1; WOW64) SkypeUriPreview Preview/0.5" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Mozilla/5.0 (compatible; Veoozbot/1.0; +http://www.veooz.com/veoozbot.html)":
		return &Crawler{
			Pattern:      `Veoozbot`,
			Instances:    []string{ "Mozilla/5.0 (compatible; Veoozbot/1.0; +http://www.veooz.com/veoozbot.html)" },
			URL:          stringer("http://www.veooz.com/veoozbot.html"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Slackbot-LinkExpanding (+https://api.slack.com/robots)",
        "Slackbot-LinkExpanding 1.0 (+https://api.slack.com/robots)":
		return &Crawler{
			Pattern:      `Slackbot`,
			Instances:    []string{ "Slackbot-LinkExpanding (+https://api.slack.com/robots)", "Slackbot-LinkExpanding 1.0 (+https://api.slack.com/robots)" },
			URL:          stringer("https://api.slack.com/robots"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Mozilla/5.0 (compatible; redditbot/1.0; +http://www.reddit.com/feedback)":
		return &Crawler{
			Pattern:      `redditbot`,
			Instances:    []string{ "Mozilla/5.0 (compatible; redditbot/1.0; +http://www.reddit.com/feedback)" },
			URL:          stringer("http://www.reddit.com/feedback"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "datagnionbot (+http://www.datagnion.com/bot.html)":
		return &Crawler{
			Pattern:      `datagnionbot`,
			Instances:    []string{ "datagnionbot (+http://www.datagnion.com/bot.html)" },
			URL:          stringer("http://www.datagnion.com/bot.html"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Google-Adwords-Instant (+http://www.google.com/adsbot.html)":
		return &Crawler{
			Pattern:      `Google-Adwords-Instant`,
			Instances:    []string{ "Google-Adwords-Instant (+http://www.google.com/adsbot.html)" },
			URL:          stringer("http://www.google.com/adsbot.html"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Mozilla/5.0 (compatible; adbeat_bot; +support@adbeat.com; support@adbeat.com)",
        "adbeat_bot":
		return &Crawler{
			Pattern:      `adbeat_bot`,
			Instances:    []string{ "Mozilla/5.0 (compatible; adbeat_bot; +support@adbeat.com; support@adbeat.com)", "adbeat_bot" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "WhatsApp",
        "WhatsApp/2.12.15/i",
        "WhatsApp/2.12.16/i",
        "WhatsApp/2.12.17/i",
        "WhatsApp/2.12.449 A",
        "WhatsApp/2.12.453 A",
        "WhatsApp/2.12.510 A",
        "WhatsApp/2.12.540 A",
        "WhatsApp/2.12.548 A",
        "WhatsApp/2.12.555 A",
        "WhatsApp/2.12.556 A",
        "WhatsApp/2.16.1/i",
        "WhatsApp/2.16.13 A",
        "WhatsApp/2.16.2/i",
        "WhatsApp/2.16.42 A",
        "WhatsApp/2.16.57 A":
		return &Crawler{
			Pattern:      `WhatsApp`,
			Instances:    []string{ "WhatsApp", "WhatsApp/2.12.15/i", "WhatsApp/2.12.16/i", "WhatsApp/2.12.17/i", "WhatsApp/2.12.449 A", "WhatsApp/2.12.453 A", "WhatsApp/2.12.510 A", "WhatsApp/2.12.540 A", "WhatsApp/2.12.548 A", "WhatsApp/2.12.555 A", "WhatsApp/2.12.556 A", "WhatsApp/2.16.1/i", "WhatsApp/2.16.13 A", "WhatsApp/2.16.2/i", "WhatsApp/2.16.42 A", "WhatsApp/2.16.57 A" },
			URL:          stringer("https://www.whatsapp.com/"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Mozilla/5.0 (compatible;contxbot/1.0)":
		return &Crawler{
			Pattern:      `contxbot`,
			Instances:    []string{ "Mozilla/5.0 (compatible;contxbot/1.0)" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Mozilla/5.0 (compatible; Pinterestbot/1.0; +http://www.pinterest.com/bot.html)",
        "Pinterest/0.2 (+http://www.pinterest.com/bot.html)":
		return &Crawler{
			Pattern:      `pinterest.com.bot`,
			Instances:    []string{ "Mozilla/5.0 (compatible; Pinterestbot/1.0; +http://www.pinterest.com/bot.html)", "Pinterest/0.2 (+http://www.pinterest.com/bot.html)" },
			URL:          stringer("http://www.pinterest.com/bot.html"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Mozilla/5.0 (compatible; electricmonk/3.2.0 +https://www.duedil.com/our-crawler/)":
		return &Crawler{
			Pattern:      `electricmonk`,
			Instances:    []string{ "Mozilla/5.0 (compatible; electricmonk/3.2.0 +https://www.duedil.com/our-crawler/)" },
			URL:          stringer("https://www.duedil.com/our-crawler/"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "GarlikCrawler/1.2 (http://garlik.com/, crawler@garlik.com)":
		return &Crawler{
			Pattern:      `GarlikCrawler`,
			Instances:    []string{ "GarlikCrawler/1.2 (http://garlik.com/, crawler@garlik.com)" },
			URL:          stringer("http://garlik.com/"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/534+ (KHTML, like Gecko) BingPreview/1.0b",
        "Mozilla/5.0 (Windows NT 6.3; WOW64; Trident/7.0; rv:11.0; BingPreview/1.0b) like Gecko",
        "Mozilla/5.0 (compatible; MSIE 10.0; Windows NT 6.2; Trident/6.0;  WOW64;  Trident/6.0;  BingPreview/1.0b)",
        "Mozilla/5.0 (compatible; MSIE 9.0; Windows NT 6.1; Trident/5.0;  WOW64;  Trident/5.0;  BingPreview/1.0b)",
        "Mozilla/5.0 (iPhone; CPU iPhone OS 7_0 like Mac OS X) AppleWebKit/537.51.1 (KHTML, like Gecko) Version/7.0 Mobile/11A465 Safari/9537.53 BingPreview/1.0b":
		return &Crawler{
			Pattern:      `BingPreview\/`,
			Instances:    []string{ "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/534+ (KHTML, like Gecko) BingPreview/1.0b", "Mozilla/5.0 (Windows NT 6.3; WOW64; Trident/7.0; rv:11.0; BingPreview/1.0b) like Gecko", "Mozilla/5.0 (compatible; MSIE 10.0; Windows NT 6.2; Trident/6.0;  WOW64;  Trident/6.0;  BingPreview/1.0b)", "Mozilla/5.0 (compatible; MSIE 9.0; Windows NT 6.1; Trident/5.0;  WOW64;  Trident/5.0;  BingPreview/1.0b)", "Mozilla/5.0 (iPhone; CPU iPhone OS 7_0 like Mac OS X) AppleWebKit/537.51.1 (KHTML, like Gecko) Version/7.0 Mobile/11A465 Safari/9537.53 BingPreview/1.0b" },
			URL:          stringer("https://www.bing.com/webmaster/help/which-crawlers-does-bing-use-8c184ec0"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Mozilla/5.0 (compatible; vebidoobot/1.0; +https://blog.vebidoo.de/vebidoobot/":
		return &Crawler{
			Pattern:      `vebidoobot`,
			Instances:    []string{ "Mozilla/5.0 (compatible; vebidoobot/1.0; +https://blog.vebidoo.de/vebidoobot/" },
			URL:          stringer("https://blog.vebidoo.de/vebidoobot/"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Mozilla/5.0 (compatible; FemtosearchBot/1.0; http://femtosearch.com)":
		return &Crawler{
			Pattern:      `FemtosearchBot`,
			Instances:    []string{ "Mozilla/5.0 (compatible; FemtosearchBot/1.0; http://femtosearch.com)" },
			URL:          stringer("http://femtosearch.com"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Mozilla/5.0 (compatible; Yahoo Link Preview; https://help.yahoo.com/kb/mail/yahoo-link-preview-SLN23615.html)":
		return &Crawler{
			Pattern:      `Yahoo Link Preview`,
			Instances:    []string{ "Mozilla/5.0 (compatible; Yahoo Link Preview; https://help.yahoo.com/kb/mail/yahoo-link-preview-SLN23615.html)" },
			URL:          stringer("https://help.yahoo.com/kb/mail/yahoo-link-preview-SLN23615.html"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Mozilla/5.0 (compatible; MetaJobBot; http://www.metajob.de/crawler)":
		return &Crawler{
			Pattern:      `MetaJobBot`,
			Instances:    []string{ "Mozilla/5.0 (compatible; MetaJobBot; http://www.metajob.de/crawler)" },
			URL:          stringer("http://www.metajob.de/the/crawler"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "DomainStatsBot/1.0 (http://domainstats.io/our-bot)":
		return &Crawler{
			Pattern:      `DomainStatsBot`,
			Instances:    []string{ "DomainStatsBot/1.0 (http://domainstats.io/our-bot)" },
			URL:          stringer("http://domainstats.io/our-bot"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "mindUpBot (datenbutler.de)":
		return &Crawler{
			Pattern:      `mindUpBot`,
			Instances:    []string{ "mindUpBot (datenbutler.de)" },
			URL:          stringer("http://www.datenbutler.de/"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Mozilla/5.0 (compatible; Daum/4.1; +http://cs.daum.net/faq/15/4118.html?faqId=28966)":
		return &Crawler{
			Pattern:      `Daum\/`,
			Instances:    []string{ "Mozilla/5.0 (compatible; Daum/4.1; +http://cs.daum.net/faq/15/4118.html?faqId=28966)" },
			URL:          stringer("http://cs.daum.net/faq/15/4118.html?faqId=28966"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Jugendschutzprogramm-Crawler; Info: http://www.jugendschutzprogramm.de":
		return &Crawler{
			Pattern:      `Jugendschutzprogramm-Crawler`,
			Instances:    []string{ "Jugendschutzprogramm-Crawler; Info: http://www.jugendschutzprogramm.de" },
			URL:          stringer("http://www.jugendschutzprogramm.de"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Xenu Link Sleuth/1.3.8":
		return &Crawler{
			Pattern:      `Xenu Link Sleuth`,
			Instances:    []string{ "Xenu Link Sleuth/1.3.8" },
			URL:          stringer("http://home.snafu.de/tilman/xenulink.html"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Pcore-HTTP/v0.40.3":
		return &Crawler{
			Pattern:      `Pcore-HTTP`,
			Instances:    []string{ "Pcore-HTTP/v0.40.3" },
			URL:          stringer("https://bitbucket.org/softvisio/pcore/overview"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/40.0.2214.111 Safari/537.36 moatbot",
        "Mozilla/5.0 (iPhone; CPU iPhone OS 8_0 like Mac OS X) AppleWebKit/600.1.3 (KHTML, like Gecko) Version/8.0 Mobile/12A4345d Safari/600.1.4 moatbot":
		return &Crawler{
			Pattern:      `moatbot`,
			Instances:    []string{ "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/40.0.2214.111 Safari/537.36 moatbot", "Mozilla/5.0 (iPhone; CPU iPhone OS 8_0 like Mac OS X) AppleWebKit/600.1.3 (KHTML, like Gecko) Version/8.0 Mobile/12A4345d Safari/600.1.4 moatbot" },
			URL:          stringer("https://moat.com"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/44.0.2403.125 Safari/537.36 (compatible; KosmioBot/1.0; +http://kosm.io/bot.html)":
		return &Crawler{
			Pattern:      `KosmioBot`,
			Instances:    []string{ "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/44.0.2403.125 Safari/537.36 (compatible; KosmioBot/1.0; +http://kosm.io/bot.html)" },
			URL:          stringer("http://kosm.io/bot.html"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Ubuntu Chromium/59.0.3071.109 Chrome/59.0.3071.109 Safari/537.36 PingdomPageSpeed/1.0 (pingbot/2.0; +http://www.pingdom.com/)",
        "Mozilla/5.0 (compatible; pingbot/2.0; +http://www.pingdom.com/)":
		return &Crawler{
			Pattern:      `pingdom`,
			Instances:    []string{ "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Ubuntu Chromium/59.0.3071.109 Chrome/59.0.3071.109 Safari/537.36 PingdomPageSpeed/1.0 (pingbot/2.0; +http://www.pingdom.com/)", "Mozilla/5.0 (compatible; pingbot/2.0; +http://www.pingdom.com/)" },
			URL:          stringer("http://www.pingdom.com"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Mozilla/5.0 (compatible; MSIE 9.0; Windows NT 6.1; Trident/5.0; AppInsights)":
		return &Crawler{
			Pattern:      `AppInsights`,
			Instances:    []string{ "Mozilla/5.0 (compatible; MSIE 9.0; Windows NT 6.1; Trident/5.0; AppInsights)" },
			URL:          stringer("https://docs.microsoft.com/en-us/azure/azure-monitor/app/app-insights-overview"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Mozilla/5.0 (Unknown; Linux x86_64) AppleWebKit/538.1 (KHTML, like Gecko) PhantomJS/2.1.1 Safari/538.1 bl.uk_lddc_renderbot/2.0.0 (+ http://www.bl.uk/aboutus/legaldeposit/websites/websites/faqswebmaster/index.html)":
		return &Crawler{
			Pattern:      `PhantomJS`,
			Instances:    []string{ "Mozilla/5.0 (Unknown; Linux x86_64) AppleWebKit/538.1 (KHTML, like Gecko) PhantomJS/2.1.1 Safari/538.1 bl.uk_lddc_renderbot/2.0.0 (+ http://www.bl.uk/aboutus/legaldeposit/websites/websites/faqswebmaster/index.html)" },
			URL:          stringer("http://phantomjs.org/"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Mozilla/5.0 (compatible; Gowikibot/1.0; +http://www.gowikibot.com)":
		return &Crawler{
			Pattern:      `Gowikibot`,
			Instances:    []string{ "Mozilla/5.0 (compatible; Gowikibot/1.0; +http://www.gowikibot.com)" },
			URL:          stringer("http://www.gowikibot.com"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "PiplBot (+http://www.pipl.com/bot/)",
        "Mozilla/5.0+(compatible;+PiplBot;+http://www.pipl.com/bot/)":
		return &Crawler{
			Pattern:      `PiplBot`,
			Instances:    []string{ "PiplBot (+http://www.pipl.com/bot/)", "Mozilla/5.0+(compatible;+PiplBot;+http://www.pipl.com/bot/)" },
			URL:          stringer("http://www.pipl.com/bot/"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Mozilla/5.0 (compatible; Discordbot/2.0; +https://discordapp.com)":
		return &Crawler{
			Pattern:      `Discordbot`,
			Instances:    []string{ "Mozilla/5.0 (compatible; Discordbot/2.0; +https://discordapp.com)" },
			URL:          stringer("https://discordapp.com"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "TelegramBot (like TwitterBot)":
		return &Crawler{
			Pattern:      `TelegramBot`,
			Instances:    []string{ "TelegramBot (like TwitterBot)" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Mozilla/5.0 (compatible; Jetslide; +http://jetsli.de/crawler)":
		return &Crawler{
			Pattern:      `Jetslide`,
			Instances:    []string{ "Mozilla/5.0 (compatible; Jetslide; +http://jetsli.de/crawler)" },
			URL:          stringer("http://jetsli.de/crawler"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Mozilla/5.0 (compatible; NewShareCounts.com/1.0; +http://newsharecounts.com/crawler)":
		return &Crawler{
			Pattern:      `newsharecounts`,
			Instances:    []string{ "Mozilla/5.0 (compatible; NewShareCounts.com/1.0; +http://newsharecounts.com/crawler)" },
			URL:          stringer("http://newsharecounts.com/crawler"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Mozilla/5.0 (Windows; U; Windows NT 5.1; en-US; rv:1.8.1.6) Gecko/20070725 Firefox/2.0.0.6 - James BOT - WebCrawler http://cognitiveseo.com/bot.html":
		return &Crawler{
			Pattern:      `James BOT`,
			Instances:    []string{ "Mozilla/5.0 (Windows; U; Windows NT 5.1; en-US; rv:1.8.1.6) Gecko/20070725 Firefox/2.0.0.6 - James BOT - WebCrawler http://cognitiveseo.com/bot.html" },
			URL:          stringer("http://cognitiveseo.com/bot.html"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Barkrowler/0.5.1 (experimenting / debugging - sorry for your logs ) http://www.exensa.com/crawl - admin@exensa.com -- based on BuBiNG",
        "Barkrowler/0.7 (+http://www.exensa.com/crawl)":
		return &Crawler{
			Pattern:      `Barkrowler`,
			Instances:    []string{ "Barkrowler/0.5.1 (experimenting / debugging - sorry for your logs ) http://www.exensa.com/crawl - admin@exensa.com -- based on BuBiNG", "Barkrowler/0.7 (+http://www.exensa.com/crawl)" },
			URL:          stringer("http://www.exensa.com/crawl"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Mozilla/5.0 (compatible; TinEye-bot/1.31; +http://www.tineye.com/crawler.html)",
        "TinEye/1.1 (http://tineye.com/crawler.html)":
		return &Crawler{
			Pattern:      `TinEye`,
			Instances:    []string{ "Mozilla/5.0 (compatible; TinEye-bot/1.31; +http://www.tineye.com/crawler.html)", "TinEye/1.1 (http://tineye.com/crawler.html)" },
			URL:          stringer("http://www.tineye.com/crawler.html"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "SocialRankIOBot; http://socialrank.io/about":
		return &Crawler{
			Pattern:      `SocialRankIOBot`,
			Instances:    []string{ "SocialRankIOBot; http://socialrank.io/about" },
			URL:          stringer("http://socialrank.io/about"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Mozilla/5.0 (Windows; U; Windows NT 6.0; en-GB; rv:1.0; trendictionbot0.5.0; trendiction search; http://www.trendiction.de/bot; please let us know of any problems; web at trendiction.com) Gecko/20071127 Firefox/3.0.0.11":
		return &Crawler{
			Pattern:      `trendictionbot`,
			Instances:    []string{ "Mozilla/5.0 (Windows; U; Windows NT 6.0; en-GB; rv:1.0; trendictionbot0.5.0; trendiction search; http://www.trendiction.de/bot; please let us know of any problems; web at trendiction.com) Gecko/20071127 Firefox/3.0.0.11" },
			URL:          stringer("http://www.trendiction.de/bot"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Ocarinabot":
		return &Crawler{
			Pattern:      `Ocarinabot`,
			Instances:    []string{ "Ocarinabot" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Mozilla/5.0 (compatible; epicbot; +http://www.epictions.com/epicbot)":
		return &Crawler{
			Pattern:      `epicbot`,
			Instances:    []string{ "Mozilla/5.0 (compatible; epicbot; +http://www.epictions.com/epicbot)" },
			URL:          stringer("http://www.epictions.com/epicbot"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Mozilla/5.0 (compatible; Primalbot; +https://www.primal.com;)":
		return &Crawler{
			Pattern:      `Primalbot`,
			Instances:    []string{ "Mozilla/5.0 (compatible; Primalbot; +https://www.primal.com;)" },
			URL:          stringer("https://www.primal.com"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Mozilla/5.0 (compatible; DuckDuckGo-Favicons-Bot/1.0; +http://duckduckgo.com)":
		return &Crawler{
			Pattern:      `DuckDuckGo-Favicons-Bot`,
			Instances:    []string{ "Mozilla/5.0 (compatible; DuckDuckGo-Favicons-Bot/1.0; +http://duckduckgo.com)" },
			URL:          stringer("http://duckduckgo.com"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:49.0) Gecko/20100101 Firefox/49.0 / GnowitNewsbot / Contact information at http://www.gnowit.com":
		return &Crawler{
			Pattern:      `GnowitNewsbot`,
			Instances:    []string{ "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:49.0) Gecko/20100101 Firefox/49.0 / GnowitNewsbot / Contact information at http://www.gnowit.com" },
			URL:          stringer("http://www.gnowit.com"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Mozilla/5.0 (Windows NT 6.3;compatible; Leikibot/1.0; +http://www.leiki.com)":
		return &Crawler{
			Pattern:      `Leikibot`,
			Instances:    []string{ "Mozilla/5.0 (Windows NT 6.3;compatible; Leikibot/1.0; +http://www.leiki.com)" },
			URL:          stringer("http://www.leiki.com"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "@LinkArchiver twitter bot":
		return &Crawler{
			Pattern:      `LinkArchiver`,
			Instances:    []string{ "@LinkArchiver twitter bot" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Mozilla/5.0 (compatible; YaK/1.0; http://linkfluence.com/; bot@linkfluence.com)":
		return &Crawler{
			Pattern:      `YaK\/`,
			Instances:    []string{ "Mozilla/5.0 (compatible; YaK/1.0; http://linkfluence.com/; bot@linkfluence.com)" },
			URL:          stringer("http://linkfluence.com"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Mozilla/5.0 (compatible; PaperLiBot/2.1; http://support.paper.li/entries/20023257-what-is-paper-li)",
        "Mozilla/5.0 (compatible; PaperLiBot/2.1; https://support.paper.li/entries/20023257-what-is-paper-li)":
		return &Crawler{
			Pattern:      `PaperLiBot`,
			Instances:    []string{ "Mozilla/5.0 (compatible; PaperLiBot/2.1; http://support.paper.li/entries/20023257-what-is-paper-li)", "Mozilla/5.0 (compatible; PaperLiBot/2.1; https://support.paper.li/entries/20023257-what-is-paper-li)" },
			URL:          stringer("http://support.paper.li/entries/20023257-what-is-paper-li"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Digg Deeper/v1 (http://digg.com/about)":
		return &Crawler{
			Pattern:      `Digg Deeper`,
			Instances:    []string{ "Digg Deeper/v1 (http://digg.com/about)" },
			URL:          stringer("http://digg.com/about"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "dcrawl/1.0":
		return &Crawler{
			Pattern:      `dcrawl`,
			Instances:    []string{ "dcrawl/1.0" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Mozilla/5.0 (compatible; Snacktory; +https://github.com/karussell/snacktory)":
		return &Crawler{
			Pattern:      `Snacktory`,
			Instances:    []string{ "Mozilla/5.0 (compatible; Snacktory; +https://github.com/karussell/snacktory)" },
			URL:          stringer("https://github.com/karussell/snacktory"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Mozilla/5.0 (compatible; AndersPinkBot/1.0; +http://anderspink.com/bot.html)":
		return &Crawler{
			Pattern:      `AndersPinkBot`,
			Instances:    []string{ "Mozilla/5.0 (compatible; AndersPinkBot/1.0; +http://anderspink.com/bot.html)" },
			URL:          stringer("http://anderspink.com/bot.html"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Fyrebot/1.0":
		return &Crawler{
			Pattern:      `Fyrebot`,
			Instances:    []string{ "Fyrebot/1.0" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Mozilla/5.0 (compatible; EveryoneSocialBot/1.0; support@everyonesocial.com http://everyonesocial.com/)":
		return &Crawler{
			Pattern:      `EveryoneSocialBot`,
			Instances:    []string{ "Mozilla/5.0 (compatible; EveryoneSocialBot/1.0; support@everyonesocial.com http://everyonesocial.com/)" },
			URL:          stringer("http://everyonesocial.com"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Mediatoolkitbot (complaints@mediatoolkit.com)":
		return &Crawler{
			Pattern:      `Mediatoolkitbot`,
			Instances:    []string{ "Mediatoolkitbot (complaints@mediatoolkit.com)" },
			URL:          stringer("http://mediatoolkit.com"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_8_2) AppleWebKit/537.13 (KHTML, like Gecko) Chrome/30.0.1599.66 Safari/537.13 Luminator-robots/2.0":
		return &Crawler{
			Pattern:      `Luminator-robots`,
			Instances:    []string{ "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_8_2) AppleWebKit/537.13 (KHTML, like Gecko) Chrome/30.0.1599.66 Safari/537.13 Luminator-robots/2.0" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Mozilla/5.0 (compatible; ExtLinksBot/1.5 +https://extlinks.com/Bot.html)":
		return &Crawler{
			Pattern:      `ExtLinksBot`,
			Instances:    []string{ "Mozilla/5.0 (compatible; ExtLinksBot/1.5 +https://extlinks.com/Bot.html)" },
			URL:          stringer("https://extlinks.com/Bot.html"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Mozilla/5.0 (Windows; U; Windows NT 5.1; en; rv:1.9.0.13) Gecko/2009073022 Firefox/3.5.2 (.NET CLR 3.5.30729) SurveyBot/2.3 (DomainTools)":
		return &Crawler{
			Pattern:      `SurveyBot`,
			Instances:    []string{ "Mozilla/5.0 (Windows; U; Windows NT 5.1; en; rv:1.9.0.13) Gecko/2009073022 Firefox/3.5.2 (.NET CLR 3.5.30729) SurveyBot/2.3 (DomainTools)" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "NING/1.0":
		return &Crawler{
			Pattern:      `NING\/`,
			Instances:    []string{ "NING/1.0" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "okhttp/2.5.0",
        "okhttp/2.7.5",
        "okhttp/3.2.0",
        "okhttp/3.5.0":
		return &Crawler{
			Pattern:      `okhttp`,
			Instances:    []string{ "okhttp/2.5.0", "okhttp/2.7.5", "okhttp/3.2.0", "okhttp/3.5.0" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Nuzzel":
		return &Crawler{
			Pattern:      `Nuzzel`,
			Instances:    []string{ "Nuzzel" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "omgili/0.5 +http://omgili.com":
		return &Crawler{
			Pattern:      `omgili`,
			Instances:    []string{ "omgili/0.5 +http://omgili.com" },
			URL:          stringer("http://omgili.com"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "PocketParser/2.0 (+https://getpocket.com/pocketparser_ua)":
		return &Crawler{
			Pattern:      `PocketParser`,
			Instances:    []string{ "PocketParser/2.0 (+https://getpocket.com/pocketparser_ua)" },
			URL:          stringer("https://getpocket.com/pocketparser_ua"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "YisouSpider":
		return &Crawler{
			Pattern:      `YisouSpider`,
			Instances:    []string{ "YisouSpider" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Mozilla/5.0 (compatible; um-LN/1.0; mailto: techinfo@ubermetrics-technologies.com)":
		return &Crawler{
			Pattern:      `um-LN`,
			Instances:    []string{ "Mozilla/5.0 (compatible; um-LN/1.0; mailto: techinfo@ubermetrics-technologies.com)" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Mozilla/5.0 (compatible; ToutiaoSpider/1.0; http://web.toutiao.com/media_cooperation/;)":
		return &Crawler{
			Pattern:      `ToutiaoSpider`,
			Instances:    []string{ "Mozilla/5.0 (compatible; ToutiaoSpider/1.0; http://web.toutiao.com/media_cooperation/;)" },
			URL:          stringer("http://web.toutiao.com/media_cooperation/"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Mozilla/5.0 (compatible; MuckRack/1.0; +http://muckrack.com)":
		return &Crawler{
			Pattern:      `MuckRack`,
			Instances:    []string{ "Mozilla/5.0 (compatible; MuckRack/1.0; +http://muckrack.com)" },
			URL:          stringer("http://muckrack.com"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Jamie's Spider (http://jamiembrown.com/)":
		return &Crawler{
			Pattern:      `Jamie's Spider`,
			Instances:    []string{ "Jamie's Spider (http://jamiembrown.com/)" },
			URL:          stringer("http://jamiembrown.com/"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "AHC/2.0":
		return &Crawler{
			Pattern:      `AHC\/`,
			Instances:    []string{ "AHC/2.0" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Mozilla/5.0 (compatible; NetcraftSurveyAgent/1.0; +info@netcraft.com)":
		return &Crawler{
			Pattern:      `NetcraftSurveyAgent`,
			Instances:    []string{ "Mozilla/5.0 (compatible; NetcraftSurveyAgent/1.0; +info@netcraft.com)" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Mozilla/5.0 (iPhone; CPU iPhone OS 8_3 like Mac OS X) AppleWebKit/600.1.4 (KHTML, like Gecko) Version/8.0 Mobile/12F70 Safari/600.1.4 (compatible; Laserlikebot/0.1)":
		return &Crawler{
			Pattern:      `Laserlikebot`,
			Instances:    []string{ "Mozilla/5.0 (iPhone; CPU iPhone OS 8_3 like Mac OS X) AppleWebKit/600.1.4 (KHTML, like Gecko) Version/8.0 Mobile/12F70 Safari/600.1.4 (compatible; Laserlikebot/0.1)" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Apache-HttpClient/4.2.3 (java 1.5)",
        "Apache-HttpClient/4.2.5 (java 1.5)",
        "Apache-HttpClient/4.3.1 (java 1.5)",
        "Apache-HttpClient/4.3.3 (java 1.5)",
        "Apache-HttpClient/4.3.5 (java 1.5)",
        "Apache-HttpClient/4.4.1 (Java/1.8.0_65)",
        "Apache-HttpClient/4.5.3 (Java/1.8.0_121)":
		return &Crawler{
			Pattern:      `Apache-HttpClient`,
			Instances:    []string{ "Apache-HttpClient/4.2.3 (java 1.5)", "Apache-HttpClient/4.2.5 (java 1.5)", "Apache-HttpClient/4.3.1 (java 1.5)", "Apache-HttpClient/4.3.3 (java 1.5)", "Apache-HttpClient/4.3.5 (java 1.5)", "Apache-HttpClient/4.4.1 (Java/1.8.0_65)", "Apache-HttpClient/4.5.3 (Java/1.8.0_121)" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "AppEngine-Google; (+http://code.google.com/appengine; appid: example)":
		return &Crawler{
			Pattern:      `AppEngine-Google`,
			Instances:    []string{ "AppEngine-Google; (+http://code.google.com/appengine; appid: example)" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Jetty/9.3.z-SNAPSHOT":
		return &Crawler{
			Pattern:      `Jetty`,
			Instances:    []string{ "Jetty/9.3.z-SNAPSHOT" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Upflow/1.0":
		return &Crawler{
			Pattern:      `Upflow`,
			Instances:    []string{ "Upflow/1.0" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Thinklab (thinklab.com)":
		return &Crawler{
			Pattern:      `Thinklab`,
			Instances:    []string{ "Thinklab (thinklab.com)" },
			URL:          stringer("thinklab.com"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Traackr.com":
		return &Crawler{
			Pattern:      `Traackr.com`,
			Instances:    []string{ "Traackr.com" },
			URL:          stringer("Traackr.com"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Ruby, Twurly v1.1 (http://twurly.org)":
		return &Crawler{
			Pattern:      `Twurly`,
			Instances:    []string{ "Ruby, Twurly v1.1 (http://twurly.org)" },
			URL:          stringer("http://twurly.org"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "http.rb/2.2.2 (Mastodon/1.5.1; +https://example-masto-instance.org/)":
		return &Crawler{
			Pattern:      `Mastodon`,
			Instances:    []string{ "http.rb/2.2.2 (Mastodon/1.5.1; +https://example-masto-instance.org/)" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "http_get":
		return &Crawler{
			Pattern:      `http_get`,
			Instances:    []string{ "http_get" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Mozilla/5.0 (compatible; DnyzBot/1.0)":
		return &Crawler{
			Pattern:      `DnyzBot`,
			Instances:    []string{ "Mozilla/5.0 (compatible; DnyzBot/1.0)" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Mozilla/5.0 (compatible; botify; http://botify.com)":
		return &Crawler{
			Pattern:      `botify`,
			Instances:    []string{ "Mozilla/5.0 (compatible; botify; http://botify.com)" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Mozilla/5.0 (compatible; 007ac9 Crawler; http://crawler.007ac9.net/)":
		return &Crawler{
			Pattern:      `007ac9 Crawler`,
			Instances:    []string{ "Mozilla/5.0 (compatible; 007ac9 Crawler; http://crawler.007ac9.net/)" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Mozilla/5.0 (compatible; BehloolBot/beta; +http://www.webeaver.com/bot)":
		return &Crawler{
			Pattern:      `BehloolBot`,
			Instances:    []string{ "Mozilla/5.0 (compatible; BehloolBot/beta; +http://www.webeaver.com/bot)" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.10; rv:41.0) Gecko/20100101 Firefox/55.0 BrandVerity/1.0 (http://www.brandverity.com/why-is-brandverity-visiting-me)":
		return &Crawler{
			Pattern:      `BrandVerity`,
			Instances:    []string{ "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.10; rv:41.0) Gecko/20100101 Firefox/55.0 BrandVerity/1.0 (http://www.brandverity.com/why-is-brandverity-visiting-me)" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "check_http/v2.2.1 (nagios-plugins 2.2.1)":
		return &Crawler{
			Pattern:      `check_http`,
			Instances:    []string{ "check_http/v2.2.1 (nagios-plugins 2.2.1)" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Mozilla/5.0 (Windows NT 6.1; compatible; BDCbot/1.0; +http://bigweb.bigdatacorp.com.br/faq.aspx) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/41.0.2272.118 Safari/537.36",
        "Mozilla/5.0 (Windows NT 10.0; Win64; x64; BDCbot/1.0; +http://bigweb.bigdatacorp.com.br/faq.aspx) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.100 Safari/537.36":
		return &Crawler{
			Pattern:      `BDCbot`,
			Instances:    []string{ "Mozilla/5.0 (Windows NT 6.1; compatible; BDCbot/1.0; +http://bigweb.bigdatacorp.com.br/faq.aspx) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/41.0.2272.118 Safari/537.36", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; BDCbot/1.0; +http://bigweb.bigdatacorp.com.br/faq.aspx) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.100 Safari/537.36" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Mozilla/5.0 (compatible; ZumBot/1.0; http://help.zum.com/inquiry)":
		return &Crawler{
			Pattern:      `ZumBot`,
			Instances:    []string{ "Mozilla/5.0 (compatible; ZumBot/1.0; http://help.zum.com/inquiry)" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "EZID (EZID link checker; https://ezid.cdlib.org/)":
		return &Crawler{
			Pattern:      `EZID`,
			Instances:    []string{ "EZID (EZID link checker; https://ezid.cdlib.org/)" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "ICC-Crawler/2.0 (Mozilla-compatible; ; http://ucri.nict.go.jp/en/icccrawler.html)":
		return &Crawler{
			Pattern:      `ICC-Crawler`,
			Instances:    []string{ "ICC-Crawler/2.0 (Mozilla-compatible; ; http://ucri.nict.go.jp/en/icccrawler.html)" },
			URL:          stringer("http://ucri.nict.go.jp/en/icccrawler.html"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "ArchiveTeam ArchiveBot/20170106.02 (wpull 2.0.2)":
		return &Crawler{
			Pattern:      `ArchiveBot`,
			Instances:    []string{ "ArchiveTeam ArchiveBot/20170106.02 (wpull 2.0.2)" },
			URL:          stringer("https://github.com/ArchiveTeam/ArchiveBot"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "LCC (+http://corpora.informatik.uni-leipzig.de/crawler_faq.html)":
		return &Crawler{
			Pattern:      `^LCC `,
			Instances:    []string{ "LCC (+http://corpora.informatik.uni-leipzig.de/crawler_faq.html)" },
			URL:          stringer("http://corpora.informatik.uni-leipzig.de/crawler_faq.html"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Mozilla/5.0 (compatible; oBot/2.3.1; +http://filterdb.iss.net/crawler/)":
		return &Crawler{
			Pattern:      `filterdb.iss.net\/crawler`,
			Instances:    []string{ "Mozilla/5.0 (compatible; oBot/2.3.1; +http://filterdb.iss.net/crawler/)" },
			URL:          stringer("http://filterdb.iss.net/crawler/"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "BLP_bbot/0.1":
		return &Crawler{
			Pattern:      `BLP_bbot`,
			Instances:    []string{ "BLP_bbot/0.1" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Mozilla/5.0 (compatible; BomboraBot/1.0; +http://www.bombora.com/bot)":
		return &Crawler{
			Pattern:      `BomboraBot`,
			Instances:    []string{ "Mozilla/5.0 (compatible; BomboraBot/1.0; +http://www.bombora.com/bot)" },
			URL:          stringer("http://www.bombora.com/bot"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Buck/2.2; (+https://app.hypefactors.com/media-monitoring/about.html)":
		return &Crawler{
			Pattern:      `Buck\/`,
			Instances:    []string{ "Buck/2.2; (+https://app.hypefactors.com/media-monitoring/about.html)" },
			URL:          stringer("https://app.hypefactors.com/media-monitoring/about.html"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Companybook-Crawler (+https://www.companybooknetworking.com/)":
		return &Crawler{
			Pattern:      `Companybook-Crawler`,
			Instances:    []string{ "Companybook-Crawler (+https://www.companybooknetworking.com/)" },
			URL:          stringer("https://www.companybooknetworking.com/"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Mozilla/5.0 (compatible; Genieo/1.0 http://www.genieo.com/webfilter.html)":
		return &Crawler{
			Pattern:      `Genieo`,
			Instances:    []string{ "Mozilla/5.0 (compatible; Genieo/1.0 http://www.genieo.com/webfilter.html)" },
			URL:          stringer("http://www.genieo.com/webfilter.html"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "magpie-crawler/1.1 (U; Linux amd64; en-GB; +http://www.brandwatch.net)":
		return &Crawler{
			Pattern:      `magpie-crawler`,
			Instances:    []string{ "magpie-crawler/1.1 (U; Linux amd64; en-GB; +http://www.brandwatch.net)" },
			URL:          stringer("http://www.brandwatch.net"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "MeltwaterNews www.meltwater.com":
		return &Crawler{
			Pattern:      `MeltwaterNews`,
			Instances:    []string{ "MeltwaterNews www.meltwater.com" },
			URL:          stringer("http://www.meltwater.com"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Mozilla/5.0 Moreover/5.1 (+http://www.moreover.com)":
		return &Crawler{
			Pattern:      `Moreover`,
			Instances:    []string{ "Mozilla/5.0 Moreover/5.1 (+http://www.moreover.com)" },
			URL:          stringer("http://www.moreover.com"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "newspaper/0.2.5",
        "newspaper/0.2.6",
        "newspaper/0.1.0.7":
		return &Crawler{
			Pattern:      `newspaper\/`,
			Instances:    []string{ "newspaper/0.2.5", "newspaper/0.2.6", "newspaper/0.1.0.7" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Mozilla/5.0 (compatible; ScoutJet; +http://www.scoutjet.com/)":
		return &Crawler{
			Pattern:      `ScoutJet`,
			Instances:    []string{ "Mozilla/5.0 (compatible; ScoutJet; +http://www.scoutjet.com/)" },
			URL:          stringer("http://www.scoutjet.com/"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "sentry/8.22.0 (https://sentry.io)":
		return &Crawler{
			Pattern:      `(^| )sentry\/`,
			Instances:    []string{ "sentry/8.22.0 (https://sentry.io)" },
			URL:          stringer("https://sentry.io"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Mozilla/5.0 (compatible; StorygizeBot; http://www.storygize.com)":
		return &Crawler{
			Pattern:      `StorygizeBot`,
			Instances:    []string{ "Mozilla/5.0 (compatible; StorygizeBot; http://www.storygize.com)" },
			URL:          stringer("http://www.storygize.com"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Mozilla/5.0+(compatible; UptimeRobot/2.0; http://www.uptimerobot.com/)":
		return &Crawler{
			Pattern:      `UptimeRobot`,
			Instances:    []string{ "Mozilla/5.0+(compatible; UptimeRobot/2.0; http://www.uptimerobot.com/)" },
			URL:          stringer("http://www.uptimerobot.com/"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "OutclicksBot/2 +https://www.outclicks.net/agent/VjzDygCuk4ubNmg40ZMbFqT0sIh7UfOKk8s8ZMiupUR",
        "OutclicksBot/2 +https://www.outclicks.net/agent/gIYbZ38dfAuhZkrFVl7sJBFOUhOVct6J1SvxgmBZgCe",
        "OutclicksBot/2 +https://www.outclicks.net/agent/PryJzTl8POCRHfvEUlRN5FKtZoWDQOBEvFJ2wh6KH5J",
        "OutclicksBot/2 +https://www.outclicks.net/agent/p2i4sNUh7eylJF1S6SGgRs5mP40ExlYvsr9GBxVQG6h":
		return &Crawler{
			Pattern:      `OutclicksBot`,
			Instances:    []string{ "OutclicksBot/2 +https://www.outclicks.net/agent/VjzDygCuk4ubNmg40ZMbFqT0sIh7UfOKk8s8ZMiupUR", "OutclicksBot/2 +https://www.outclicks.net/agent/gIYbZ38dfAuhZkrFVl7sJBFOUhOVct6J1SvxgmBZgCe", "OutclicksBot/2 +https://www.outclicks.net/agent/PryJzTl8POCRHfvEUlRN5FKtZoWDQOBEvFJ2wh6KH5J", "OutclicksBot/2 +https://www.outclicks.net/agent/p2i4sNUh7eylJF1S6SGgRs5mP40ExlYvsr9GBxVQG6h" },
			URL:          stringer("https://www.outclicks.net"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Mozilla/5.0 (compatible; seoscanners.net/1; +spider@seoscanners.net)":
		return &Crawler{
			Pattern:      `seoscanners`,
			Instances:    []string{ "Mozilla/5.0 (compatible; seoscanners.net/1; +spider@seoscanners.net)" },
			URL:          stringer("http://www.seoscanners.net/"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Hatena Antenna/0.3",
        "Hatena::Russia::Crawler/0.01":
		return &Crawler{
			Pattern:      `Hatena`,
			Instances:    []string{ "Hatena Antenna/0.3", "Hatena::Russia::Crawler/0.01" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Mozilla/5.0 (Linux; U; Android 2.3.4; generic) AppleWebKit/537.36 (KHTML, like Gecko; Google Web Preview) Version/4.0 Mobile Safari/537.36":
		return &Crawler{
			Pattern:      `Google Web Preview`,
			Instances:    []string{ "Mozilla/5.0 (Linux; U; Android 2.3.4; generic) AppleWebKit/537.36 (KHTML, like Gecko; Google Web Preview) Version/4.0 Mobile Safari/537.36" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "MauiBot (crawler.feedback+wc@gmail.com)":
		return &Crawler{
			Pattern:      `MauiBot`,
			Instances:    []string{ "MauiBot (crawler.feedback+wc@gmail.com)" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Mozilla/5.0 (compatible; AlphaBot/3.2; +http://alphaseobot.com/bot.html)":
		return &Crawler{
			Pattern:      `AlphaBot`,
			Instances:    []string{ "Mozilla/5.0 (compatible; AlphaBot/3.2; +http://alphaseobot.com/bot.html)" },
			URL:          stringer("http://alphaseobot.com/bot.html"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "SBL-BOT (http://sbl.net)":
		return &Crawler{
			Pattern:      `SBL-BOT`,
			Instances:    []string{ "SBL-BOT (http://sbl.net)" },
			URL:          stringer("http://sbl.net"),
			Description:  stringer("Bot of SoftByte BlackWidow"),
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "IAS crawler (ias_crawler; http://integralads.com/site-indexing-policy/)":
		return &Crawler{
			Pattern:      `IAS crawler`,
			Instances:    []string{ "IAS crawler (ias_crawler; http://integralads.com/site-indexing-policy/)" },
			URL:          stringer("http://integralads.com/site-indexing-policy/"),
			Description:  stringer("Bot of Integral Ad Science, Inc."),
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Mozilla/5.0 (compatible; adscanner/)":
		return &Crawler{
			Pattern:      `adscanner`,
			Instances:    []string{ "Mozilla/5.0 (compatible; adscanner/)" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Netvibes (crawler/bot; http://www.netvibes.com",
        "Netvibes (crawler; http://www.netvibes.com)":
		return &Crawler{
			Pattern:      `Netvibes`,
			Instances:    []string{ "Netvibes (crawler/bot; http://www.netvibes.com", "Netvibes (crawler; http://www.netvibes.com)" },
			URL:          stringer("http://www.netvibes.com"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Mozilla/5.0 (compatible;acapbot/0.1;treat like Googlebot)",
        "Mozilla/5.0 (compatible;acapbot/0.1.;treat like Googlebot)":
		return &Crawler{
			Pattern:      `acapbot`,
			Instances:    []string{ "Mozilla/5.0 (compatible;acapbot/0.1;treat like Googlebot)", "Mozilla/5.0 (compatible;acapbot/0.1.;treat like Googlebot)" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Baidu-YunGuanCe-Bot(ce.baidu.com)",
        "Baidu-YunGuanCe-SLABot(ce.baidu.com)",
        "Baidu-YunGuanCe-ScanBot(ce.baidu.com)",
        "Baidu-YunGuanCe-PerfBot(ce.baidu.com)",
        "Baidu-YunGuanCe-VSBot(ce.baidu.com)":
		return &Crawler{
			Pattern:      `Baidu-YunGuanCe`,
			Instances:    []string{ "Baidu-YunGuanCe-Bot(ce.baidu.com)", "Baidu-YunGuanCe-SLABot(ce.baidu.com)", "Baidu-YunGuanCe-ScanBot(ce.baidu.com)", "Baidu-YunGuanCe-PerfBot(ce.baidu.com)", "Baidu-YunGuanCe-VSBot(ce.baidu.com)" },
			URL:          stringer("https://ce.baidu.com/topic/topic20150908"),
			Description:  stringer("Baidu Cloud Watch"),
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "bitlybot/3.0 (+http://bit.ly/)",
        "bitlybot/2.0",
        "bitlybot":
		return &Crawler{
			Pattern:      `bitlybot`,
			Instances:    []string{ "bitlybot/3.0 (+http://bit.ly/)", "bitlybot/2.0", "bitlybot" },
			URL:          stringer("http://bit.ly/"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "blogmuraBot (+http://www.blogmura.com)":
		return &Crawler{
			Pattern:      `blogmuraBot`,
			Instances:    []string{ "blogmuraBot (+http://www.blogmura.com)" },
			URL:          stringer("http://www.blogmura.com"),
			Description:  stringer("A blog ranking site which links to blogs on just about every theme possible."),
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Bot.AraTurka.com/0.0.1":
		return &Crawler{
			Pattern:      `Bot.AraTurka.com`,
			Instances:    []string{ "Bot.AraTurka.com/0.0.1" },
			URL:          stringer("http://www.araturka.com"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "bot-pge.chlooe.com/1.0.0 (+http://www.chlooe.com/)":
		return &Crawler{
			Pattern:      `bot-pge.chlooe.com`,
			Instances:    []string{ "bot-pge.chlooe.com/1.0.0 (+http://www.chlooe.com/)" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Mozilla/5.0 (compatible; BoxcarBot/1.1; +awesome@boxcar.io)":
		return &Crawler{
			Pattern:      `BoxcarBot`,
			Instances:    []string{ "Mozilla/5.0 (compatible; BoxcarBot/1.1; +awesome@boxcar.io)" },
			URL:          stringer("https://boxcar.io/"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "BTWebClient/180B(9704)":
		return &Crawler{
			Pattern:      `BTWebClient`,
			Instances:    []string{ "BTWebClient/180B(9704)" },
			URL:          stringer("http://www.utorrent.com/"),
			Description:  stringer("µTorrent BitTorrent Client"),
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.0;.NET CLR 1.0.3705; ContextAd Bot 1.0)",
        "ContextAd Bot 1.0":
		return &Crawler{
			Pattern:      `ContextAd Bot`,
			Instances:    []string{ "Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.0;.NET CLR 1.0.3705; ContextAd Bot 1.0)", "ContextAd Bot 1.0" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Mozilla/5.0 (compatible; Digincore bot; https://www.digincore.com/crawler.html for rules and instructions.)":
		return &Crawler{
			Pattern:      `Digincore bot`,
			Instances:    []string{ "Mozilla/5.0 (compatible; Digincore bot; https://www.digincore.com/crawler.html for rules and instructions.)" },
			URL:          stringer("http://www.digincore.com/crawler.html"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Disqus/1.0":
		return &Crawler{
			Pattern:      `Disqus`,
			Instances:    []string{ "Disqus/1.0" },
			URL:          stringer("https://disqus.com/"),
			Description:  stringer("validate and quality check pages."),
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Feedly/1.0 (+http://www.feedly.com/fetcher.html; like FeedFetcher-Google)",
        "FeedlyBot/1.0 (http://feedly.com)":
		return &Crawler{
			Pattern:      `Feedly`,
			Instances:    []string{ "Feedly/1.0 (+http://www.feedly.com/fetcher.html; like FeedFetcher-Google)", "FeedlyBot/1.0 (http://feedly.com)" },
			URL:          stringer("https://www.feedly.com/fetcher.html"),
			Description:  stringer("Feedly Fetcher is how Feedly grabs RSS or Atom feeds when users choose to add them to their Feedly or any of the other applications built on top of the feedly cloud."),
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Fetch/2.0a (CMS Detection/Web/SEO analysis tool, see http://guess.scritch.org)":
		return &Crawler{
			Pattern:      `Fetch\/`,
			Instances:    []string{ "Fetch/2.0a (CMS Detection/Web/SEO analysis tool, see http://guess.scritch.org)" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Fever/1.38 (Feed Parser; http://feedafever.com; Allow like Gecko)":
		return &Crawler{
			Pattern:      `Fever`,
			Instances:    []string{ "Fever/1.38 (Feed Parser; http://feedafever.com; Allow like Gecko)" },
			URL:          stringer("http://feedafever.com"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Flamingo_SearchEngine (+http://www.flamingosearch.com/bot)":
		return &Crawler{
			Pattern:      `Flamingo_SearchEngine`,
			Instances:    []string{ "Flamingo_SearchEngine (+http://www.flamingosearch.com/bot)" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Mozilla/5.0 (compatible; FlipboardProxy/1.1; +http://flipboard.com/browserproxy)",
        "Mozilla/5.0 (compatible; FlipboardProxy/1.2; +http://flipboard.com/browserproxy)",
        "Mozilla/5.0 (Macintosh; U; Intel Mac OS X 10.6; en-US; rv:1.9.2) Gecko/20100115 Firefox/3.6 (FlipboardProxy/1.1; +http://flipboard.com/browserproxy)",
        "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.9; rv:28.0) Gecko/20100101 Firefox/28.0 (FlipboardProxy/1.1; +http://flipboard.com/browserproxy)":
		return &Crawler{
			Pattern:      `FlipboardProxy`,
			Instances:    []string{ "Mozilla/5.0 (compatible; FlipboardProxy/1.1; +http://flipboard.com/browserproxy)", "Mozilla/5.0 (compatible; FlipboardProxy/1.2; +http://flipboard.com/browserproxy)", "Mozilla/5.0 (Macintosh; U; Intel Mac OS X 10.6; en-US; rv:1.9.2) Gecko/20100115 Firefox/3.6 (FlipboardProxy/1.1; +http://flipboard.com/browserproxy)", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.9; rv:28.0) Gecko/20100101 Firefox/28.0 (FlipboardProxy/1.1; +http://flipboard.com/browserproxy)" },
			URL:          stringer("https://about.flipboard.com/browserproxy/"),
			Description:  stringer("a proxy service to fetch, validate, and prepare certain elements of websites for presentation through the Flipboard Application"),
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "g2reader-bot/1.0 (+http://www.g2reader.com/)":
		return &Crawler{
			Pattern:      `g2reader-bot`,
			Instances:    []string{ "g2reader-bot/1.0 (+http://www.g2reader.com/)" },
			URL:          stringer("http://www.g2reader.com/"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "G2 Web Services/1.0 (built with StormCrawler Archetype 1.8; https://www.g2webservices.com/; developers@g2llc.com)":
		return &Crawler{
			Pattern:      `G2 Web Services`,
			Instances:    []string{ "G2 Web Services/1.0 (built with StormCrawler Archetype 1.8; https://www.g2webservices.com/; developers@g2llc.com)" },
			URL:          stringer("https://www.g2webservices.com/"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Mozilla/5.0 (compatible; imrbot/1.10.8 +http://www.mignify.com)":
		return &Crawler{
			Pattern:      `imrbot`,
			Instances:    []string{ "Mozilla/5.0 (compatible; imrbot/1.10.8 +http://www.mignify.com)" },
			URL:          stringer("http://www.mignify.com"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "K7MLWCBot/1.0 (+http://www.k7computing.com)":
		return &Crawler{
			Pattern:      `K7MLWCBot`,
			Instances:    []string{ "K7MLWCBot/1.0 (+http://www.k7computing.com)" },
			URL:          stringer("http://www.k7computing.com"),
			Description:  stringer("Virus scanner"),
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Kemvibot/1.0 (http://kemvi.com, marco@kemvi.com)":
		return &Crawler{
			Pattern:      `Kemvibot`,
			Instances:    []string{ "Kemvibot/1.0 (http://kemvi.com, marco@kemvi.com)" },
			URL:          stringer("http://kemvi.com"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Landau-Media-Spider/1.0(http://bots.landaumedia.de/bot.html)":
		return &Crawler{
			Pattern:      `Landau-Media-Spider`,
			Instances:    []string{ "Landau-Media-Spider/1.0(http://bots.landaumedia.de/bot.html)" },
			URL:          stringer("http://bots.landaumedia.de/bot.html"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "linkapediabot (+http://www.linkapedia.com)":
		return &Crawler{
			Pattern:      `linkapediabot`,
			Instances:    []string{ "linkapediabot (+http://www.linkapedia.com)" },
			URL:          stringer("http://www.linkapedia.com"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Mozilla/5.0 (compatible; vkShare; +http://vk.com/dev/Share)":
		return &Crawler{
			Pattern:      `vkShare`,
			Instances:    []string{ "Mozilla/5.0 (compatible; vkShare; +http://vk.com/dev/Share)" },
			URL:          stringer("http://vk.com/dev/Share"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Mozilla/5.0 (compatible; MSIE 10.0; Windows NT 6.1; Trident/6.0) LinkCheck by Siteimprove.com",
        "Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 5.0) Match by Siteimprove.com",
        "Mozilla/5.0 (compatible; MSIE 10.0; Windows NT 6.1; Trident/6.0) SiteCheck-sitecrawl by Siteimprove.com",
        "Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 5.0) LinkCheck by Siteimprove.com":
		return &Crawler{
			Pattern:      `Siteimprove.com`,
			Instances:    []string{ "Mozilla/5.0 (compatible; MSIE 10.0; Windows NT 6.1; Trident/6.0) LinkCheck by Siteimprove.com", "Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 5.0) Match by Siteimprove.com", "Mozilla/5.0 (compatible; MSIE 10.0; Windows NT 6.1; Trident/6.0) SiteCheck-sitecrawl by Siteimprove.com", "Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 5.0) LinkCheck by Siteimprove.com" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Mozilla/5.0 (compatible; BLEXBot/1.0; +http://webmeup-crawler.com/)":
		return &Crawler{
			Pattern:      `BLEXBot\/`,
			Instances:    []string{ "Mozilla/5.0 (compatible; BLEXBot/1.0; +http://webmeup-crawler.com/)" },
			URL:          stringer("http://webmeup-crawler.com"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/62.0.3202.75 Safari/537.36 DareBoost":
		return &Crawler{
			Pattern:      `DareBoost`,
			Instances:    []string{ "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/62.0.3202.75 Safari/537.36 DareBoost" },
			URL:          stringer("https://www.dareboost.com/"),
			Description:  stringer("Bot to test, Analyze and Optimize website"),
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Mozilla/5.0 (compatible; ZuperlistBot/1.0)":
		return &Crawler{
			Pattern:      `ZuperlistBot\/`,
			Instances:    []string{ "Mozilla/5.0 (compatible; ZuperlistBot/1.0)" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Mozilla/5.0 (compatible; Miniflux/2.0.x-dev; +https://miniflux.net)",
        "Mozilla/5.0 (compatible; Miniflux/2.0.3; +https://miniflux.net)",
        "Mozilla/5.0 (compatible; Miniflux/2.0.7; +https://miniflux.net)",
        "Mozilla/5.0 (compatible; Miniflux/2.0.10; +https://miniflux.net)",
        "Mozilla/5.0 (compatibl$; Miniflux/2.0.x-dev; +https://miniflux.app)",
        "Mozilla/5.0 (compatible; Miniflux/2.0.11; +https://miniflux.app)",
        "Mozilla/5.0 (compatible; Miniflux/2.0.12; +https://miniflux.app)",
        "Mozilla/5.0 (compatible; Miniflux/ae1dc1a; +https://miniflux.app)",
        "Mozilla/5.0 (compatible; Miniflux/3b6e44c; +https://miniflux.app)":
		return &Crawler{
			Pattern:      `Miniflux\/`,
			Instances:    []string{ "Mozilla/5.0 (compatible; Miniflux/2.0.x-dev; +https://miniflux.net)", "Mozilla/5.0 (compatible; Miniflux/2.0.3; +https://miniflux.net)", "Mozilla/5.0 (compatible; Miniflux/2.0.7; +https://miniflux.net)", "Mozilla/5.0 (compatible; Miniflux/2.0.10; +https://miniflux.net)", "Mozilla/5.0 (compatibl$; Miniflux/2.0.x-dev; +https://miniflux.app)", "Mozilla/5.0 (compatible; Miniflux/2.0.11; +https://miniflux.app)", "Mozilla/5.0 (compatible; Miniflux/2.0.12; +https://miniflux.app)", "Mozilla/5.0 (compatible; Miniflux/ae1dc1a; +https://miniflux.app)", "Mozilla/5.0 (compatible; Miniflux/3b6e44c; +https://miniflux.app)" },
			URL:          stringer("https://miniflux.net"),
			Description:  stringer("Miniflux is a minimalist and opinionated feed reader."),
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Mozilla/5.0 (compatible; Feedspotbot/1.0; +http://www.feedspot.com/fs/bot)",
        "Mozilla/5.0 (compatible; Feedspot/1.0 (+https://www.feedspot.com/fs/fetcher; like FeedFetcher-Google)":
		return &Crawler{
			Pattern:      `Feedspot`,
			Instances:    []string{ "Mozilla/5.0 (compatible; Feedspotbot/1.0; +http://www.feedspot.com/fs/bot)", "Mozilla/5.0 (compatible; Feedspot/1.0 (+https://www.feedspot.com/fs/fetcher; like FeedFetcher-Google)" },
			URL:          stringer("http://www.feedspot.com/fs/bot"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Mozilla/5.0 (Windows; U; Windows NT 5.1; en-US; rv:1.9.1.2) Gecko/20090729 Firefox/3.5.2 (.NET CLR 3.5.30729; Diffbot/0.1; +http://www.diffbot.com)":
		return &Crawler{
			Pattern:      `Diffbot\/`,
			Instances:    []string{ "Mozilla/5.0 (Windows; U; Windows NT 5.1; en-US; rv:1.9.1.2) Gecko/20090729 Firefox/3.5.2 (.NET CLR 3.5.30729; Diffbot/0.1; +http://www.diffbot.com)" },
			URL:          stringer("http://www.diffbot.com"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Mozilla/5.0 (compatible; SEOkicks; +https://www.seokicks.de/robot.html)":
		return &Crawler{
			Pattern:      `SEOkicks`,
			Instances:    []string{ "Mozilla/5.0 (compatible; SEOkicks; +https://www.seokicks.de/robot.html)" },
			URL:          stringer("https://www.seokicks.de/robot.html"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Mozilla/5.0 (compatible; tracemyfile/1.0; +bot@tracemyfile.com)":
		return &Crawler{
			Pattern:      `tracemyfile`,
			Instances:    []string{ "Mozilla/5.0 (compatible; tracemyfile/1.0; +bot@tracemyfile.com)" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Mozilla/5.0 (compatible; Nimbostratus-Bot/v1.3.2; http://cloudsystemnetworks.com)":
		return &Crawler{
			Pattern:      `Nimbostratus-Bot`,
			Instances:    []string{ "Mozilla/5.0 (compatible; Nimbostratus-Bot/v1.3.2; http://cloudsystemnetworks.com)" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Mozilla/5.0 zgrab/0.x":
		return &Crawler{
			Pattern:      `zgrab`,
			Instances:    []string{ "Mozilla/5.0 zgrab/0.x" },
			URL:          stringer("https://zmap.io/"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Mozilla/5.0 (compatible; PR-CY.RU; + https://a.pr-cy.ru)":
		return &Crawler{
			Pattern:      `PR-CY.RU`,
			Instances:    []string{ "Mozilla/5.0 (compatible; PR-CY.RU; + https://a.pr-cy.ru)" },
			URL:          stringer("https://a.pr-cy.ru/"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "AdsTxtCrawler/1.0":
		return &Crawler{
			Pattern:      `AdsTxtCrawler`,
			Instances:    []string{ "AdsTxtCrawler/1.0" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Datafeedwatch/2.1.x":
		return &Crawler{
			Pattern:      `Datafeedwatch`,
			Instances:    []string{ "Datafeedwatch/2.1.x" },
			URL:          stringer("https://www.datafeedwatch.com/"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Zabbix":
		return &Crawler{
			Pattern:      `Zabbix`,
			Instances:    []string{ "Zabbix" },
			URL:          stringer("https://www.zabbix.com/documentation/3.4/manual/web_monitoring"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "TangibleeBot/1.0.0.0 (http://tangiblee.com/bot)":
		return &Crawler{
			Pattern:      `TangibleeBot`,
			Instances:    []string{ "TangibleeBot/1.0.0.0 (http://tangiblee.com/bot)" },
			URL:          stringer("http://tangiblee.com/bot"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "google-xrawler":
		return &Crawler{
			Pattern:      `google-xrawler`,
			Instances:    []string{ "google-xrawler" },
			URL:          stringer("https://webmasters.stackexchange.com/questions/105560/what-is-the-google-xrawler-user-agent-used-for"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "axios/0.18.0":
		return &Crawler{
			Pattern:      `axios`,
			Instances:    []string{ "axios/0.18.0" },
			URL:          stringer("https://github.com/axios/axios"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Amazon CloudFront":
		return &Crawler{
			Pattern:      `Amazon CloudFront`,
			Instances:    []string{ "Amazon CloudFront" },
			URL:          stringer("https://aws.amazon.com/cloudfront/"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Pulsepoint XT3 web scraper":
		return &Crawler{
			Pattern:      `Pulsepoint`,
			Instances:    []string{ "Pulsepoint XT3 web scraper" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Mozilla/5.0 (compatible; CloudFlare-AlwaysOnline/1.0; +http://www.cloudflare.com/always-online) AppleWebKit/534.34",
        "Mozilla/5.0 (compatible; CloudFlare-AlwaysOnline/1.0; +https://www.cloudflare.com/always-online) AppleWebKit/534.34":
		return &Crawler{
			Pattern:      `CloudFlare-AlwaysOnline`,
			Instances:    []string{ "Mozilla/5.0 (compatible; CloudFlare-AlwaysOnline/1.0; +http://www.cloudflare.com/always-online) AppleWebKit/534.34", "Mozilla/5.0 (compatible; CloudFlare-AlwaysOnline/1.0; +https://www.cloudflare.com/always-online) AppleWebKit/534.34" },
			URL:          stringer("https://www.cloudflare.com/always-online/"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Mozilla/5.0 (compatible; Google-Structured-Data-Testing-Tool +https://search.google.com/structured-data/testing-tool)",
        "Mozilla/5.0 (compatible; Google-Structured-Data-Testing-Tool +http://developers.google.com/structured-data/testing-tool/)":
		return &Crawler{
			Pattern:      `Google-Structured-Data-Testing-Tool`,
			Instances:    []string{ "Mozilla/5.0 (compatible; Google-Structured-Data-Testing-Tool +https://search.google.com/structured-data/testing-tool)", "Mozilla/5.0 (compatible; Google-Structured-Data-Testing-Tool +http://developers.google.com/structured-data/testing-tool/)" },
			URL:          stringer("https://search.google.com/structured-data/testing-tool"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "WordupInfoSearch/1.0":
		return &Crawler{
			Pattern:      `WordupInfoSearch`,
			Instances:    []string{ "WordupInfoSearch/1.0" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Mozilla/5.0 (compatible; WebDataStats/1.0 ; +https://webdatastats.com/policy.html)":
		return &Crawler{
			Pattern:      `WebDataStats`,
			Instances:    []string{ "Mozilla/5.0 (compatible; WebDataStats/1.0 ; +https://webdatastats.com/policy.html)" },
			URL:          stringer("https://webdatastats.com/"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Jersey/2.25.1 (HttpUrlConnection 1.8.0_141)":
		return &Crawler{
			Pattern:      `HttpUrlConnection`,
			Instances:    []string{ "Jersey/2.25.1 (HttpUrlConnection 1.8.0_141)" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Mozilla/5.0 (compatible; Seekport Crawler; http://seekport.com/)":
		return &Crawler{
			Pattern:      `Seekport Crawler`,
			Instances:    []string{ "Mozilla/5.0 (compatible; Seekport Crawler; http://seekport.com/)" },
			URL:          stringer("http://seekport.com/"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "ZoomBot (Linkbot 1.0 http://suite.seozoom.it/bot.html)":
		return &Crawler{
			Pattern:      `ZoomBot`,
			Instances:    []string{ "ZoomBot (Linkbot 1.0 http://suite.seozoom.it/bot.html)" },
			URL:          stringer("http://suite.seozoom.it/bot.html"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "VelenPublicWebCrawler (velen.io)":
		return &Crawler{
			Pattern:      `VelenPublicWebCrawler`,
			Instances:    []string{ "VelenPublicWebCrawler (velen.io)" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "MoodleBot/1.0":
		return &Crawler{
			Pattern:      `MoodleBot`,
			Instances:    []string{ "MoodleBot/1.0" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "jpg-newsbot/2.0; (+https://vipnytt.no/bots/)":
		return &Crawler{
			Pattern:      `jpg-newsbot`,
			Instances:    []string{ "jpg-newsbot/2.0; (+https://vipnytt.no/bots/)" },
			URL:          stringer("https://vipnytt.no/bots/"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Mozilla/5.0 (Java) outbrain":
		return &Crawler{
			Pattern:      `outbrain`,
			Instances:    []string{ "Mozilla/5.0 (Java) outbrain" },
			URL:          stringer("https://www.outbrain.com/help/advertisers/invalid-url/"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "W3C_Validator/1.3":
		return &Crawler{
			Pattern:      `W3C_Validator`,
			Instances:    []string{ "W3C_Validator/1.3" },
			URL:          stringer("https://validator.w3.org/services"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Validator.nu/LV":
		return &Crawler{
			Pattern:      `Validator\.nu`,
			Instances:    []string{ "Validator.nu/LV" },
			URL:          stringer("https://validator.w3.org/services"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "W3C-checklink/2.90 libwww-perl/5.64",
        "W3C-checklink/3.6.2.3 libwww-perl/5.64",
        "W3C-checklink/4.2 [4.20] libwww-perl/5.803",
        "W3C-checklink/4.2.1 [4.21] libwww-perl/5.803",
        "W3C-checklink/4.3 [4.42] libwww-perl/5.805",
        "W3C-checklink/4.3 [4.42] libwww-perl/5.808",
        "W3C-checklink/4.3 [4.42] libwww-perl/5.820",
        "W3C-checklink/4.5 [4.154] libwww-perl/5.823",
        "W3C-checklink/4.5 [4.160] libwww-perl/5.823":
		return &Crawler{
			Pattern:      `W3C-checklink`,
			Instances:    []string{ "W3C-checklink/2.90 libwww-perl/5.64", "W3C-checklink/3.6.2.3 libwww-perl/5.64", "W3C-checklink/4.2 [4.20] libwww-perl/5.803", "W3C-checklink/4.2.1 [4.21] libwww-perl/5.803", "W3C-checklink/4.3 [4.42] libwww-perl/5.805", "W3C-checklink/4.3 [4.42] libwww-perl/5.808", "W3C-checklink/4.3 [4.42] libwww-perl/5.820", "W3C-checklink/4.5 [4.154] libwww-perl/5.823", "W3C-checklink/4.5 [4.160] libwww-perl/5.823" },
			URL:          stringer("https://validator.w3.org/services"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{ "libwww-perl" },
		}

	case "W3C-mobileOK/DDC-1.0":
		return &Crawler{
			Pattern:      `W3C-mobileOK`,
			Instances:    []string{ "W3C-mobileOK/DDC-1.0" },
			URL:          stringer("https://validator.w3.org/services"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "W3C_I18n-Checker/1.0":
		return &Crawler{
			Pattern:      `W3C_I18n-Checker`,
			Instances:    []string{ "W3C_I18n-Checker/1.0" },
			URL:          stringer("https://validator.w3.org/services"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "FeedValidator/1.3":
		return &Crawler{
			Pattern:      `FeedValidator`,
			Instances:    []string{ "FeedValidator/1.3" },
			URL:          stringer("https://validator.w3.org/services"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Jigsaw/2.3.0 W3C_CSS_Validator_JFouffa/2.0":
		return &Crawler{
			Pattern:      `W3C_CSS_Validator`,
			Instances:    []string{ "Jigsaw/2.3.0 W3C_CSS_Validator_JFouffa/2.0" },
			URL:          stringer("https://validator.w3.org/services"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "W3C_Unicorn/1.0":
		return &Crawler{
			Pattern:      `W3C_Unicorn`,
			Instances:    []string{ "W3C_Unicorn/1.0" },
			URL:          stringer("https://validator.w3.org/services"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Mozilla/5.0 (Google-PhysicalWeb)":
		return &Crawler{
			Pattern:      `Google-PhysicalWeb`,
			Instances:    []string{ "Mozilla/5.0 (Google-PhysicalWeb)" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Blackboard Safeassign":
		return &Crawler{
			Pattern:      `Blackboard`,
			Instances:    []string{ "Blackboard Safeassign" },
			URL:          stringer("https://help.blackboard.com/Learn/Administrator/Hosting/Tools_Management/SafeAssign"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Mozilla/5.0 (compatible; ICBot/0.1; +https://ideasandcode.xyz":
		return &Crawler{
			Pattern:      `ICBot\/`,
			Instances:    []string{ "Mozilla/5.0 (compatible; ICBot/0.1; +https://ideasandcode.xyz" },
			URL:          stringer("https://ideasandcode.xyz"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Mozilla/5.0 (compatible; BazQux/2.4; +https://bazqux.com/fetcher; 1 subscribers)":
		return &Crawler{
			Pattern:      `BazQux`,
			Instances:    []string{ "Mozilla/5.0 (compatible; BazQux/2.4; +https://bazqux.com/fetcher; 1 subscribers)" },
			URL:          stringer("https://bazqux.com/fetcher"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Mozilla/5.0 (compatible; Twingly Recon; twingly.com)":
		return &Crawler{
			Pattern:      `Twingly`,
			Instances:    []string{ "Mozilla/5.0 (compatible; Twingly Recon; twingly.com)" },
			URL:          stringer("https://twingly.com"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Mozilla/5.0 (compatible; Rivva; http://rivva.de)":
		return &Crawler{
			Pattern:      `Rivva`,
			Instances:    []string{ "Mozilla/5.0 (compatible; Rivva; http://rivva.de)" },
			URL:          stringer("http://rivva.de"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Experibot-v2 http://goo.gl/ZAr8wX",
        "Experibot-v3 http://goo.gl/ZAr8wX":
		return &Crawler{
			Pattern:      `Experibot`,
			Instances:    []string{ "Experibot-v2 http://goo.gl/ZAr8wX", "Experibot-v3 http://goo.gl/ZAr8wX" },
			URL:          stringer("https://amirkr.wixsite.com/experibot"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Mozilla/5.0 (Windows NT 6.2; WOW64) AppleWebKit/537.22 (KHTML, like Gecko) Chrome/25.0.1364.5 Safari/537.22 +awesomecrawler":
		return &Crawler{
			Pattern:      `awesomecrawler`,
			Instances:    []string{ "Mozilla/5.0 (Windows NT 6.2; WOW64) AppleWebKit/537.22 (KHTML, like Gecko) Chrome/25.0.1364.5 Safari/537.22 +awesomecrawler" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Mozilla/5.0 (compatible; Dataprovider.com)":
		return &Crawler{
			Pattern:      `Dataprovider.com`,
			Instances:    []string{ "Mozilla/5.0 (compatible; Dataprovider.com)" },
			URL:          stringer("https://www.dataprovider.com/"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Mozilla/5.0 (compatible; GroupHigh/1.0; +http://www.grouphigh.com/":
		return &Crawler{
			Pattern:      `GroupHigh\/`,
			Instances:    []string{ "Mozilla/5.0 (compatible; GroupHigh/1.0; +http://www.grouphigh.com/" },
			URL:          stringer("http://www.grouphigh.com/"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Mozilla/5.0 (compatible; theoldreader.com)":
		return &Crawler{
			Pattern:      `theoldreader.com`,
			Instances:    []string{ "Mozilla/5.0 (compatible; theoldreader.com)" },
			URL:          stringer("https://www.theoldreader.com/"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Mozilla/5.0 (compatible; U; AnyEvent-HTTP/2.24; +http://software.schmorp.de/pkg/AnyEvent)":
		return &Crawler{
			Pattern:      `AnyEvent`,
			Instances:    []string{ "Mozilla/5.0 (compatible; U; AnyEvent-HTTP/2.24; +http://software.schmorp.de/pkg/AnyEvent)" },
			URL:          stringer("http://software.schmorp.de/pkg/AnyEvent.html"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Uptimebot.org - Free website monitoring":
		return &Crawler{
			Pattern:      `Uptimebot\.org`,
			Instances:    []string{ "Uptimebot.org - Free website monitoring" },
			URL:          stringer("http://uptimebot.org/"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Mozilla/5.0 (compatible; Nmap Scripting Engine; https://nmap.org/book/nse.html)":
		return &Crawler{
			Pattern:      `Nmap Scripting Engine`,
			Instances:    []string{ "Mozilla/5.0 (compatible; Nmap Scripting Engine; https://nmap.org/book/nse.html)" },
			URL:          stringer("https://nmap.org/book/nse.html"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "2ip.ru CMS Detector (https://2ip.ru/cms/)":
		return &Crawler{
			Pattern:      `2ip.ru`,
			Instances:    []string{ "2ip.ru CMS Detector (https://2ip.ru/cms/)" },
			URL:          stringer("https://2ip.ru/cms/"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Clickagy Intelligence Bot v2":
		return &Crawler{
			Pattern:      `Clickagy`,
			Instances:    []string{ "Clickagy Intelligence Bot v2" },
			URL:          stringer("https://www.clickagy.com"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Caliperbot/1.0 (+http://www.conductor.com/caliperbot)":
		return &Crawler{
			Pattern:      `Caliperbot`,
			Instances:    []string{ "Caliperbot/1.0 (+http://www.conductor.com/caliperbot)" },
			URL:          stringer("http://www.conductor.com/caliperbot"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "MBCrawler/1.0 (https://monitorbacklinks.com)":
		return &Crawler{
			Pattern:      `MBCrawler`,
			Instances:    []string{ "MBCrawler/1.0 (https://monitorbacklinks.com)" },
			URL:          stringer("https://monitorbacklinks.com"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Mozilla/5.0 (compatible; online-webceo-bot/1.0; +http://online.webceo.com)":
		return &Crawler{
			Pattern:      `online-webceo-bot`,
			Instances:    []string{ "Mozilla/5.0 (compatible; online-webceo-bot/1.0; +http://online.webceo.com)" },
			URL:          stringer("http://online.webceo.com"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "B2B Bot":
		return &Crawler{
			Pattern:      `B2B Bot`,
			Instances:    []string{ "B2B Bot" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Mozilla/5.0 (compatible; AddSearchBot/0.9; +http://www.addsearch.com/bot; info@addsearch.com)":
		return &Crawler{
			Pattern:      `AddSearchBot`,
			Instances:    []string{ "Mozilla/5.0 (compatible; AddSearchBot/0.9; +http://www.addsearch.com/bot; info@addsearch.com)" },
			URL:          stringer("http://www.addsearch.com/bot"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/49.0.2623.75 Safari/537.36 Google Favicon":
		return &Crawler{
			Pattern:      `Google Favicon`,
			Instances:    []string{ "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/49.0.2623.75 Safari/537.36 Google Favicon" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Mozilla/5.0 (Windows; U; Windows NT 5.1; en-US) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/27.0.1453.116 Safari/537.36 HubSpot Webcrawler - web-crawlers@hubspot.com",
        "Mozilla/5.0 (X11; Linux x86_64; HubSpot Single Page link check; web-crawlers+links@hubspot.com) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/51.0.2704.103 Safari/537.36",
        "Mozilla/5.0 (compatible; HubSpot Crawler; web-crawlers@hubspot.com)":
		return &Crawler{
			Pattern:      `HubSpot`,
			Instances:    []string{ "Mozilla/5.0 (Windows; U; Windows NT 5.1; en-US) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/27.0.1453.116 Safari/537.36 HubSpot Webcrawler - web-crawlers@hubspot.com", "Mozilla/5.0 (X11; Linux x86_64; HubSpot Single Page link check; web-crawlers+links@hubspot.com) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/51.0.2704.103 Safari/537.36", "Mozilla/5.0 (compatible; HubSpot Crawler; web-crawlers@hubspot.com)" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Mozilla/5.0 (Linux; Android 6.0.1; Nexus 5 Build/MRA58N) AppleWebKit/537.36(KHTML, like Gecko) Chrome/69.0.3464.0 Mobile Safari/537.36 Chrome-Lighthouse",
        "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36(KHTML, like Gecko) Chrome/69.0.3464.0 Safari/537.36 Chrome-Lighthouse",
        "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3694.0 Safari/537.36 Chrome-Lighthouse",
        "Mozilla/5.0 (Linux; Android 6.0.1; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3694.0 Mobile Safari/537.36 Chrome-Lighthouse":
		return &Crawler{
			Pattern:      `Chrome-Lighthouse`,
			Instances:    []string{ "Mozilla/5.0 (Linux; Android 6.0.1; Nexus 5 Build/MRA58N) AppleWebKit/537.36(KHTML, like Gecko) Chrome/69.0.3464.0 Mobile Safari/537.36 Chrome-Lighthouse", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36(KHTML, like Gecko) Chrome/69.0.3464.0 Safari/537.36 Chrome-Lighthouse", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3694.0 Safari/537.36 Chrome-Lighthouse", "Mozilla/5.0 (Linux; Android 6.0.1; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3694.0 Mobile Safari/537.36 Chrome-Lighthouse" },
			URL:          stringer("https://developers.google.com/speed/pagespeed/insights"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) HeadlessChrome/74.0.3729.169 Safari/537.36":
		return &Crawler{
			Pattern:      `HeadlessChrome`,
			Instances:    []string{ "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) HeadlessChrome/74.0.3729.169 Safari/537.36" },
			URL:          stringer("https://developers.google.com/web/updates/2017/04/headless-chrome"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "CheckMarkNetwork/1.0 (+http://www.checkmarknetwork.com/spider.html)":
		return &Crawler{
			Pattern:      `CheckMarkNetwork\/`,
			Instances:    []string{ "CheckMarkNetwork/1.0 (+http://www.checkmarknetwork.com/spider.html)" },
			URL:          stringer("https://www.checkmarknetwork.com/"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case "Mozilla/5.0 (compatible; Uptimebot/1.0; +http://www.uptime.com/uptimebot)":
		return &Crawler{
			Pattern:      `www\.uptime\.com`,
			Instances:    []string{ "Mozilla/5.0 (compatible; Uptimebot/1.0; +http://www.uptime.com/uptimebot)" },
			URL:          stringer("http://www.uptime.com/uptimebot"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	default:
		return nil
	}
}

