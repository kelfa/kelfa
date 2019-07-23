package crawlerflagger
// DO NOT EDIT: Auto generated

import "regexp"

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
    // #0
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
    // #1
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
    // #2
	case "Googlebot-Image/1.0":
		return &Crawler{
			Pattern:      `Googlebot-Image`,
			Instances:    []string{ "Googlebot-Image/1.0" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #3
	case "Googlebot-News":
		return &Crawler{
			Pattern:      `Googlebot-News`,
			Instances:    []string{ "Googlebot-News" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #4
	case "Googlebot-Video/1.0":
		return &Crawler{
			Pattern:      `Googlebot-Video`,
			Instances:    []string{ "Googlebot-Video/1.0" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #5
	case "AdsBot-Google (+http://www.google.com/adsbot.html)":
		return &Crawler{
			Pattern:      `AdsBot-Google([^-]|$)`,
			Instances:    []string{ "AdsBot-Google (+http://www.google.com/adsbot.html)" },
			URL:          stringer("https://support.google.com/webmasters/answer/1061943?hl=en"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #6
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
    // #7
	case "Feedfetcher-Google; (+http://www.google.com/feedfetcher.html; 1 subscribers; feed-id=728742641706423)":
		return &Crawler{
			Pattern:      `Feedfetcher-Google`,
			Instances:    []string{ "Feedfetcher-Google; (+http://www.google.com/feedfetcher.html; 1 subscribers; feed-id=728742641706423)" },
			URL:          stringer("https://support.google.com/webmasters/answer/178852"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #8
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
    // #10
	case "APIs-Google (+https://developers.google.com/webmasters/APIs-Google.html)":
		return &Crawler{
			Pattern:      `APIs-Google`,
			Instances:    []string{ "APIs-Google (+https://developers.google.com/webmasters/APIs-Google.html)" },
			URL:          stringer("https://support.google.com/webmasters/answer/1061943?hl=en"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #11
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
    // #12
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
    // #13
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
    // #14
	case "eCairn-Grabber/1.0 (+http://ecairn.com/grabber) curl/7.15":
		return &Crawler{
			Pattern:      `curl`,
			Instances:    []string{ "eCairn-Grabber/1.0 (+http://ecairn.com/grabber) curl/7.15" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #15
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
    // #16
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
    // #17
	case "python-requests/2.18.4":
		return &Crawler{
			Pattern:      `python-requests`,
			Instances:    []string{ "python-requests/2.18.4" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #18
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
    // #19
	case "httpunit/1.x":
		return &Crawler{
			Pattern:      `httpunit`,
			Instances:    []string{ "httpunit/1.x" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #20
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
    // #21
	case "Go-http-client/1.1":
		return &Crawler{
			Pattern:      `Go-http-client`,
			Instances:    []string{ "Go-http-client/1.1" },
			URL:          stringer("https://golang.org/pkg/net/http/"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #22
	case "phpcrawl":
		return &Crawler{
			Pattern:      `phpcrawl`,
			Instances:    []string{ "phpcrawl" },
			URL:          stringer("http://phpcrawl.cuab.de/"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #23
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
    // #25
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
    // #26
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
    // #27
	case "BIGLOTRON (Beta 2;GNU/Linux)":
		return &Crawler{
			Pattern:      `BIGLOTRON`,
			Instances:    []string{ "BIGLOTRON (Beta 2;GNU/Linux)" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #28
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
    // #29
	case "ConveraCrawler/0.9e (+http://ews.converasearch.com/crawl.htm)":
		return &Crawler{
			Pattern:      `convera`,
			Instances:    []string{ "ConveraCrawler/0.9e (+http://ews.converasearch.com/crawl.htm)" },
			URL:          stringer("http://ews.converasearch.com/crawl.htm"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #30
	case "Seekbot/1.0 (http://www.seekbot.net/bot.html) RobotsTxtFetcher/1.2":
		return &Crawler{
			Pattern:      `seekbot`,
			Instances:    []string{ "Seekbot/1.0 (http://www.seekbot.net/bot.html) RobotsTxtFetcher/1.2" },
			URL:          stringer("http://www.seekbot.net/bot.html"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #31
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
    // #32
	case "GigablastOpenSource/1.0":
		return &Crawler{
			Pattern:      `Gigablast`,
			Instances:    []string{ "GigablastOpenSource/1.0" },
			URL:          stringer("https://github.com/gigablast/open-source-search-engine"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #33
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
    // #34
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
    // #35
	case "GingerCrawler/1.0 (Language Assistant for Dyslexics; www.gingersoftware.com/crawler_agent.htm; support at ginger software dot com)":
		return &Crawler{
			Pattern:      `GingerCrawler`,
			Instances:    []string{ "GingerCrawler/1.0 (Language Assistant for Dyslexics; www.gingersoftware.com/crawler_agent.htm; support at ginger software dot com)" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #37
	case "Mozilla/4.5 (compatible; HTTrack 3.0x; Windows 98)":
		return &Crawler{
			Pattern:      `HTTrack`,
			Instances:    []string{ "Mozilla/4.5 (compatible; HTTrack 3.0x; Windows 98)" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #38
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
    // #42
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
    // #44
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
    // #46
	case "panscient.com":
		return &Crawler{
			Pattern:      `panscient`,
			Instances:    []string{ "panscient.com" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #47
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
    // #49
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
    // #51
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
    // #52
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
    // #53
	case "Yanga WorldSearch Bot v1.1/beta (http://www.yanga.co.uk/)":
		return &Crawler{
			Pattern:      `yanga`,
			Instances:    []string{ "Yanga WorldSearch Bot v1.1/beta (http://www.yanga.co.uk/)" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #54
	case "Buzzbot/1.0 (Buzzbot; http://www.buzzstream.com; buzzbot@buzzstream.com)":
		return &Crawler{
			Pattern:      `buzzbot`,
			Instances:    []string{ "Buzzbot/1.0 (Buzzbot; http://www.buzzstream.com; buzzbot@buzzstream.com)" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #55
	case "MLBot (www.metadatalabs.com/mlbot)":
		return &Crawler{
			Pattern:      `mlbot`,
			Instances:    []string{ "MLBot (www.metadatalabs.com/mlbot)" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #56
	case "Mozilla/5.0 (compatible; YandexBot/3.0; +http://yandex.com/bots)":
		return &Crawler{
			Pattern:      `YandexBot`,
			Instances:    []string{ "Mozilla/5.0 (compatible; YandexBot/3.0; +http://yandex.com/bots)" },
			URL:          stringer("http://yandex.com/bots"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #57
	case "Mozilla/5.0 (compatible; YandexImages/3.0; +http://yandex.com/bots)":
		return &Crawler{
			Pattern:      `YandexImages`,
			Instances:    []string{ "Mozilla/5.0 (compatible; YandexImages/3.0; +http://yandex.com/bots)" },
			URL:          stringer("http://yandex.com/bots"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #58
	case "Mozilla/5.0 (compatible; YandexAccessibilityBot/3.0; +http://yandex.com/bots":
		return &Crawler{
			Pattern:      `YandexAccessibilityBot`,
			Instances:    []string{ "Mozilla/5.0 (compatible; YandexAccessibilityBot/3.0; +http://yandex.com/bots" },
			URL:          stringer("http://yandex.com/bots"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #59
	case "Mozilla/5.0 (iPhone; CPU iPhone OS 8_1 like Mac OS X) AppleWebKit/600.1.4 (KHTML, like Gecko) Version/8.0 Mobile/12B411 Safari/600.1.4 (compatible; YandexMobileBot/3.0; +http://yandex.com/bots)":
		return &Crawler{
			Pattern:      `YandexMobileBot`,
			Instances:    []string{ "Mozilla/5.0 (iPhone; CPU iPhone OS 8_1 like Mac OS X) AppleWebKit/600.1.4 (KHTML, like Gecko) Version/8.0 Mobile/12B411 Safari/600.1.4 (compatible; YandexMobileBot/3.0; +http://yandex.com/bots)" },
			URL:          stringer("https://yandex.com/support/webmaster/robot-workings/check-yandex-robots.xml#robot-in-logs"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #61
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
    // #62
	case "CyberPatrol SiteCat Webbot (http://www.cyberpatrol.com/cyberpatrolcrawler.asp)":
		return &Crawler{
			Pattern:      `CyberPatrol`,
			Instances:    []string{ "CyberPatrol SiteCat Webbot (http://www.cyberpatrol.com/cyberpatrolcrawler.asp)" },
			URL:          stringer("http://www.cyberpatrol.com/cyberpatrolcrawler.asp"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #63
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
    // #64
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
    // #66
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
    // #68
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
    // #69
	case "TurnitinBot (https://turnitin.com/robot/crawlerinfo.html)":
		return &Crawler{
			Pattern:      `TurnitinBot`,
			Instances:    []string{ "TurnitinBot (https://turnitin.com/robot/crawlerinfo.html)" },
			URL:          stringer("http://www.turnitin.com"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #71
	case "Mozilla/5.0 (compatible;  Page2RSS/0.7; +http://page2rss.com/)":
		return &Crawler{
			Pattern:      `page2rss`,
			Instances:    []string{ "Mozilla/5.0 (compatible;  Page2RSS/0.7; +http://page2rss.com/)" },
			URL:          stringer("http://www.page2rss.com"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #72
	case "Mozilla/5.0 (compatible; Whoiswebsitebot/0.1; +http://www.whoiswebsite.net)":
		return &Crawler{
			Pattern:      `sitebot`,
			Instances:    []string{ "Mozilla/5.0 (compatible; Whoiswebsitebot/0.1; +http://www.whoiswebsite.net)" },
			URL:          stringer("http://www.sitebot.org"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #73
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
    // #75
	case "Mozilla/5.0 (compatible; Ezooms/1.0; ezooms.bot@gmail.com)":
		return &Crawler{
			Pattern:      `ezooms`,
			Instances:    []string{ "Mozilla/5.0 (compatible; Ezooms/1.0; ezooms.bot@gmail.com)" },
			URL:          stringer("http://www.phpbb.com/community/viewtopic.php?f=64&t=935605&start=450#p12948289"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #76
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
    // #77
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
    // #78
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
    // #79
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
    // #81
	case "Mozilla/5.0 (compatible; MSIE 7.0 +http://www.europarchive.org)":
		return &Crawler{
			Pattern:      `europarchive.org`,
			Instances:    []string{ "Mozilla/5.0 (compatible; MSIE 7.0 +http://www.europarchive.org)" },
			URL:          stringer(""),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #82
	case "Mozilla/5.0 (compatible; NerdByNature.Bot; http://www.nerdbynature.net/bot)":
		return &Crawler{
			Pattern:      `NerdByNature.Bot`,
			Instances:    []string{ "Mozilla/5.0 (compatible; NerdByNature.Bot; http://www.nerdbynature.net/bot)" },
			URL:          stringer("http://www.nerdbynature.net/bot"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #84
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
    // #85
	case "fuelbot":
		return &Crawler{
			Pattern:      `fuelbot`,
			Instances:    []string{ "fuelbot" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #86
	case "CrunchBot/1.0 (+http://www.leadcrunch.com/crunchbot)":
		return &Crawler{
			Pattern:      `CrunchBot`,
			Instances:    []string{ "CrunchBot/1.0 (+http://www.leadcrunch.com/crunchbot)" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #87
	case "Mozilla/5.0 (Windows NT 6.1; rv:38.0) Gecko/20100101 Firefox/38.0 (IndeedBot 1.1)":
		return &Crawler{
			Pattern:      `IndeedBot`,
			Instances:    []string{ "Mozilla/5.0 (Windows NT 6.1; rv:38.0) Gecko/20100101 Firefox/38.0 (IndeedBot 1.1)" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #88
	case "Mozilla/5.0 (compatible; Mappy/1.0; +http://mappydata.net/bot/)":
		return &Crawler{
			Pattern:      `mappydata`,
			Instances:    []string{ "Mozilla/5.0 (compatible; Mappy/1.0; +http://mappydata.net/bot/)" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #89
	case "woobot":
		return &Crawler{
			Pattern:      `woobot`,
			Instances:    []string{ "woobot" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #90
	case "ZoominfoBot (zoominfobot at zoominfo dot com)":
		return &Crawler{
			Pattern:      `ZoominfoBot`,
			Instances:    []string{ "ZoominfoBot (zoominfobot at zoominfo dot com)" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #91
	case "Mozilla/5.0 (compatible; PrivacyAwareBot/1.1; +http://www.privacyaware.org)":
		return &Crawler{
			Pattern:      `PrivacyAwareBot`,
			Instances:    []string{ "Mozilla/5.0 (compatible; PrivacyAwareBot/1.1; +http://www.privacyaware.org)" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #92
	case "Mozilla/5.0 (compatible; MSIE 9.0; Windows NT 6.1; Multiviewbot":
		return &Crawler{
			Pattern:      `Multiviewbot`,
			Instances:    []string{ "Mozilla/5.0 (compatible; MSIE 9.0; Windows NT 6.1; Multiviewbot" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #93
	case "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/45.0.2454.101 Safari/537.36 SWIMGBot":
		return &Crawler{
			Pattern:      `SWIMGBot`,
			Instances:    []string{ "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/45.0.2454.101 Safari/537.36 SWIMGBot" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #94
	case "Mozilla/5.0 (compatible; Grobbot/2.2; +https://grob.it)":
		return &Crawler{
			Pattern:      `Grobbot`,
			Instances:    []string{ "Mozilla/5.0 (compatible; Grobbot/2.2; +https://grob.it)" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #95
	case "Mozilla/5.0 (compatible; eright/1.0; +bot@eright.com)":
		return &Crawler{
			Pattern:      `eright`,
			Instances:    []string{ "Mozilla/5.0 (compatible; eright/1.0; +bot@eright.com)" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #96
	case "Mozilla/5.0 (compatible; Apercite; +http://www.apercite.fr/robot/index.html)":
		return &Crawler{
			Pattern:      `Apercite`,
			Instances:    []string{ "Mozilla/5.0 (compatible; Apercite; +http://www.apercite.fr/robot/index.html)" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #97
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
    // #98
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
    // #99
	case "CipaCrawler/3.0 (info@domaincrawler.com; http://www.domaincrawler.com/www.example.com)":
		return &Crawler{
			Pattern:      `domaincrawler`,
			Instances:    []string{ "CipaCrawler/3.0 (info@domaincrawler.com; http://www.domaincrawler.com/www.example.com)" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #101
	case "Summify (Summify/1.0.1; +http://summify.com)":
		return &Crawler{
			Pattern:      `summify`,
			Instances:    []string{ "Summify (Summify/1.0.1; +http://summify.com)" },
			URL:          stringer("http://summify.com"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #102
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
    // #104
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
    // #105
	case "ec2linkfinder":
		return &Crawler{
			Pattern:      `ec2linkfinder`,
			Instances:    []string{ "ec2linkfinder" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #107
	case "Mozilla/5.0 (compatible; aiHitBot/2.9; +https://www.aihitdata.com/about)":
		return &Crawler{
			Pattern:      `aiHitBot`,
			Instances:    []string{ "Mozilla/5.0 (compatible; aiHitBot/2.9; +https://www.aihitdata.com/about)" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #109
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
    // #110
	case "Mozilla/5.0 (compatible; Yeti/1.1; +http://naver.me/bot)":
		return &Crawler{
			Pattern:      `Yeti`,
			Instances:    []string{ "Mozilla/5.0 (compatible; Yeti/1.1; +http://naver.me/bot)" },
			URL:          stringer("http://naver.me/bot"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #111
	case "Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 5.1; RetrevoPageAnalyzer; +http://www.retrevo.com/content/about-us)":
		return &Crawler{
			Pattern:      `RetrevoPageAnalyzer`,
			Instances:    []string{ "Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 5.1; RetrevoPageAnalyzer; +http://www.retrevo.com/content/about-us)" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #113
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
    // #116
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
    // #118
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
    // #119
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
    // #121
	case "drupact/0.7; http://www.arocom.de/drupact":
		return &Crawler{
			Pattern:      `drupact`,
			Instances:    []string{ "drupact/0.7; http://www.arocom.de/drupact" },
			URL:          stringer("http://www.arocom.de/drupact"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #128
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
    // #129
	case "www.integromedb.org/Crawler":
		return &Crawler{
			Pattern:      `integromedb`,
			Instances:    []string{ "www.integromedb.org/Crawler" },
			URL:          stringer("http://www.integromedb.org/Crawler"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #132
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
    // #134
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
    // #136
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
    // #137
	case "Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1; SV1;  http://www.changedetection.com/bot.html )":
		return &Crawler{
			Pattern:      `changedetection`,
			Instances:    []string{ "Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1; SV1;  http://www.changedetection.com/bot.html )" },
			URL:          stringer("http://www.changedetection.com/bot.html"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #139
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
    // #142
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
    // #143
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
    // #144
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
    // #146
	case "CC Metadata Scaper http://wiki.creativecommons.org/Metadata_Scraper":
		return &Crawler{
			Pattern:      `CC Metadata Scaper`,
			Instances:    []string{ "CC Metadata Scaper http://wiki.creativecommons.org/Metadata_Scraper" },
			URL:          stringer("http://wiki.creativecommons.org/Metadata_Scraper"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #148
	case "Mozilla/5.0 (compatible; GrapeshotCrawler/2.0; +http://www.grapeshot.co.uk/crawler.php)":
		return &Crawler{
			Pattern:      `GrapeshotCrawler`,
			Instances:    []string{ "Mozilla/5.0 (compatible; GrapeshotCrawler/2.0; +http://www.grapeshot.co.uk/crawler.php)" },
			URL:          stringer("http://www.grapeshot.co.uk/crawler.php"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #149
	case "Mozilla/5.0 (compatible; URLAppendBot/1.0; +http://www.profound.net/urlappendbot.html)":
		return &Crawler{
			Pattern:      `urlappendbot`,
			Instances:    []string{ "Mozilla/5.0 (compatible; URLAppendBot/1.0; +http://www.profound.net/urlappendbot.html)" },
			URL:          stringer("http://www.profound.net/urlappendbot.html"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #151
	case "Mozilla/5.0 (compatible; fr-crawler/1.1)":
		return &Crawler{
			Pattern:      `fr-crawler`,
			Instances:    []string{ "Mozilla/5.0 (compatible; fr-crawler/1.1)" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #152
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
    // #153
	case "SimpleCrawler/0.1":
		return &Crawler{
			Pattern:      `SimpleCrawler`,
			Instances:    []string{ "SimpleCrawler/0.1" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #154
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
    // #155
	case "cXensebot/1.1a":
		return &Crawler{
			Pattern:      `cXensebot`,
			Instances:    []string{ "cXensebot/1.1a" },
			URL:          stringer("http://www.cxense.com/bot.html"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #156
	case "Mozilla/5.0 (compatible; SMTBot/1.0; +http://www.similartech.com/smtbot)",
        "SMTBot (similartech.com/smtbot)",
        "Mozilla/5.0 (iPhone; CPU iPhone OS 6_0 like Mac OS X) AppleWebKit/536.26 (KHTML, like Gecko)                 Version/6.0 Mobile/10A5376e Safari/8536.25 (compatible; SMTBot/1.0; +http://www.similartech.com/smtbot)",
        "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.75 Safari/537.36 (compatible; SMTBot/1.0; +http://www.similartech.com/smtbot)":
		return &Crawler{
			Pattern:      `smtbot`,
			Instances:    []string{ "Mozilla/5.0 (compatible; SMTBot/1.0; +http://www.similartech.com/smtbot)", "SMTBot (similartech.com/smtbot)", "Mozilla/5.0 (iPhone; CPU iPhone OS 6_0 like Mac OS X) AppleWebKit/536.26 (KHTML, like Gecko)                 Version/6.0 Mobile/10A5376e Safari/8536.25 (compatible; SMTBot/1.0; +http://www.similartech.com/smtbot)", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.75 Safari/537.36 (compatible; SMTBot/1.0; +http://www.similartech.com/smtbot)" },
			URL:          stringer("http://www.similartech.com/smtbot"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #157
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
    // #158
	case "A6-Indexer":
		return &Crawler{
			Pattern:      `A6-Indexer`,
			Instances:    []string{ "A6-Indexer" },
			URL:          stringer("http://www.a6corp.com/a6-web-scraping-policy/"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #159
	case "ADmantX Platform Semantic Analyzer - ADmantX Inc. - www.admantx.com - support@admantx.com":
		return &Crawler{
			Pattern:      `ADmantX`,
			Instances:    []string{ "ADmantX Platform Semantic Analyzer - ADmantX Inc. - www.admantx.com - support@admantx.com" },
			URL:          stringer("http://www.admantx.com"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #160
	case "Facebot/1.0":
		return &Crawler{
			Pattern:      `Facebot`,
			Instances:    []string{ "Facebot/1.0" },
			URL:          stringer("https://developers.facebook.com/docs/sharing/best-practices#crawl"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #161
	case "Mozilla/5.0 (compatible; OrangeBot/2.0; support.orangebot@orange.com":
		return &Crawler{
			Pattern:      `OrangeBot\/`,
			Instances:    []string{ "Mozilla/5.0 (compatible; OrangeBot/2.0; support.orangebot@orange.com" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #162
	case "Mozilla/5.0 (compatible; memorybot/1.21.14 +http://mignify.com/bot.html)":
		return &Crawler{
			Pattern:      `memorybot`,
			Instances:    []string{ "Mozilla/5.0 (compatible; memorybot/1.21.14 +http://mignify.com/bot.html)" },
			URL:          stringer("http://mignify.com/bot.htm"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #163
	case "Mozilla/5.0 (compatible; AdvBot/2.0; +http://advbot.net/bot.html)":
		return &Crawler{
			Pattern:      `AdvBot`,
			Instances:    []string{ "Mozilla/5.0 (compatible; AdvBot/2.0; +http://advbot.net/bot.html)" },
			URL:          stringer("http://advbot.net/bot.html"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #164
	case "Mozilla/5.0 (compatible; MegaIndex.ru/2.0; +https://www.megaindex.ru/?tab=linkAnalyze)":
		return &Crawler{
			Pattern:      `MegaIndex`,
			Instances:    []string{ "Mozilla/5.0 (compatible; MegaIndex.ru/2.0; +https://www.megaindex.ru/?tab=linkAnalyze)" },
			URL:          stringer("https://www.megaindex.ru/?tab=linkAnalyze"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #165
	case "SemanticScholarBot/1.0 (+http://s2.allenai.org/bot.html)",
        "Mozilla/5.0 (compatible) SemanticScholarBot (+https://www.semanticscholar.org/crawler)":
		return &Crawler{
			Pattern:      `SemanticScholarBot`,
			Instances:    []string{ "SemanticScholarBot/1.0 (+http://s2.allenai.org/bot.html)", "Mozilla/5.0 (compatible) SemanticScholarBot (+https://www.semanticscholar.org/crawler)" },
			URL:          stringer("https://www.semanticscholar.org/crawler"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #166
	case "ltx71 - (http://ltx71.com/)":
		return &Crawler{
			Pattern:      `ltx71`,
			Instances:    []string{ "ltx71 - (http://ltx71.com/)" },
			URL:          stringer("http://ltx71.com/"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #167
	case "nerdybot":
		return &Crawler{
			Pattern:      `nerdybot`,
			Instances:    []string{ "nerdybot" },
			URL:          stringer("http://nerdybot.com/"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #168
	case "Mozilla/5.0 (compatible; XoviBot/2.0; +http://www.xovibot.net/)":
		return &Crawler{
			Pattern:      `xovibot`,
			Instances:    []string{ "Mozilla/5.0 (compatible; XoviBot/2.0; +http://www.xovibot.net/)" },
			URL:          stringer("http://www.xovibot.net/"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #169
	case "BUbiNG (+http://law.di.unimi.it/BUbiNG.html)":
		return &Crawler{
			Pattern:      `BUbiNG`,
			Instances:    []string{ "BUbiNG (+http://law.di.unimi.it/BUbiNG.html)" },
			URL:          stringer("http://law.di.unimi.it/BUbiNG.html"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #170
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
    // #171
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
    // #172
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
    // #173
	case "Mozilla/5.0 (TweetmemeBot/4.0; +http://datasift.com/bot.html) Gecko/20100101 Firefox/31.0":
		return &Crawler{
			Pattern:      `TweetmemeBot`,
			Instances:    []string{ "Mozilla/5.0 (TweetmemeBot/4.0; +http://datasift.com/bot.html) Gecko/20100101 Firefox/31.0" },
			URL:          stringer("http://datasift.com/bot.html"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #174
	case "crawler4j (http://code.google.com/p/crawler4j/)":
		return &Crawler{
			Pattern:      `crawler4j`,
			Instances:    []string{ "crawler4j (http://code.google.com/p/crawler4j/)" },
			URL:          stringer("https://github.com/yasserg/crawler4j"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #175
	case "Mozilla/5.0 (compatible; Findxbot/1.0; +http://www.findxbot.com)":
		return &Crawler{
			Pattern:      `findxbot`,
			Instances:    []string{ "Mozilla/5.0 (compatible; Findxbot/1.0; +http://www.findxbot.com)" },
			URL:          stringer("http://www.findxbot.com"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #176
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
    // #177
	case "Mozilla/5.0 (compatible; yoozBot-2.2; http://yooz.ir; info@yooz.ir)":
		return &Crawler{
			Pattern:      `yoozBot`,
			Instances:    []string{ "Mozilla/5.0 (compatible; yoozBot-2.2; http://yooz.ir; info@yooz.ir)" },
			URL:          stringer("http://yooz.ir"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #178
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
    // #179
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
    // #180
	case "Domain Re-Animator Bot (http://domainreanimator.com) - support@domainreanimator.com":
		return &Crawler{
			Pattern:      `Domain Re-Animator Bot`,
			Instances:    []string{ "Domain Re-Animator Bot (http://domainreanimator.com) - support@domainreanimator.com" },
			URL:          stringer("http://domainreanimator.com"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #181
	case "AddThis.com robot tech.support@clearspring.com":
		return &Crawler{
			Pattern:      `AddThis`,
			Instances:    []string{ "AddThis.com robot tech.support@clearspring.com" },
			URL:          stringer("https://www.addthis.com"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #182
	case "Screaming Frog SEO Spider/5.1":
		return &Crawler{
			Pattern:      `Screaming Frog SEO Spider`,
			Instances:    []string{ "Screaming Frog SEO Spider/5.1" },
			URL:          stringer("http://www.screamingfrog.co.uk/seo-spider"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #183
	case "MetaURI API/2.0 +metauri.com":
		return &Crawler{
			Pattern:      `MetaURI`,
			Instances:    []string{ "MetaURI API/2.0 +metauri.com" },
			URL:          stringer("http://www.useragentstring.com/MetaURI_id_17683.php"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #184
	case "Scrapy/1.0.3 (+http://scrapy.org)":
		return &Crawler{
			Pattern:      `Scrapy`,
			Instances:    []string{ "Scrapy/1.0.3 (+http://scrapy.org)" },
			URL:          stringer("http://scrapy.org/"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #185
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
    // #186
	case "Mozilla/5.0 (compatible; OpenHoseBot/2.1; +http://www.openhose.org/bot.html)":
		return &Crawler{
			Pattern:      `OpenHoseBot`,
			Instances:    []string{ "Mozilla/5.0 (compatible; OpenHoseBot/2.1; +http://www.openhose.org/bot.html)" },
			URL:          stringer("http://www.openhose.org/bot.html"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #187
	case "CapsuleChecker (http://www.capsulink.com/)":
		return &Crawler{
			Pattern:      `CapsuleChecker`,
			Instances:    []string{ "CapsuleChecker (http://www.capsulink.com/)" },
			URL:          stringer("http://www.capsulink.com/about"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #188
	case "Mozilla/5.0 (compatible) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/47.0.2526.73 Safari/537.36 collection@infegy.com":
		return &Crawler{
			Pattern:      `collection@infegy.com`,
			Instances:    []string{ "Mozilla/5.0 (compatible) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/47.0.2526.73 Safari/537.36 collection@infegy.com" },
			URL:          stringer("http://infegy.com/"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #189
	case "Mozilla/5.0 (compatible; IstellaBot/1.23.15 +http://www.tiscali.it/)":
		return &Crawler{
			Pattern:      `IstellaBot`,
			Instances:    []string{ "Mozilla/5.0 (compatible; IstellaBot/1.23.15 +http://www.tiscali.it/)" },
			URL:          stringer("http://www.tiscali.it/"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #190
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
    // #192
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
    // #193
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
    // #194
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
    // #195
	case "SafeSearch microdata crawler (https://safesearch.avira.com, safesearch-abuse@avira.com)":
		return &Crawler{
			Pattern:      `SafeSearch microdata crawler`,
			Instances:    []string{ "SafeSearch microdata crawler (https://safesearch.avira.com, safesearch-abuse@avira.com)" },
			URL:          stringer("https://safesearch.avira.com"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #196
	case "Mozilla/5.0 (compatible; Gluten Free Crawler/1.0; +http://glutenfreepleasure.com/)":
		return &Crawler{
			Pattern:      `Gluten Free Crawler\/`,
			Instances:    []string{ "Mozilla/5.0 (compatible; Gluten Free Crawler/1.0; +http://glutenfreepleasure.com/)" },
			URL:          stringer("http://glutenfreepleasure.com/"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #197
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
    // #198
	case "Mozilla/5.0 (compatible; Sysomos/1.0; +http://www.sysomos.com/; Sysomos)":
		return &Crawler{
			Pattern:      `Sysomos`,
			Instances:    []string{ "Mozilla/5.0 (compatible; Sysomos/1.0; +http://www.sysomos.com/; Sysomos)" },
			URL:          stringer("http://www.sysomos.com"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #200
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
    // #201
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
    // #202
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
    // #203
	case "Mozilla/5.0 (compatible; RankActiveLinkBot; +https://rankactive.com/resources/rankactive-linkbot)":
		return &Crawler{
			Pattern:      `RankActiveLinkBot`,
			Instances:    []string{ "Mozilla/5.0 (compatible; RankActiveLinkBot; +https://rankactive.com/resources/rankactive-linkbot)" },
			URL:          stringer("https://rankactive.com/resources/rankactive-linkbot"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #204
	case "iskanie (+http://www.iskanie.com)":
		return &Crawler{
			Pattern:      `iskanie`,
			Instances:    []string{ "iskanie (+http://www.iskanie.com)" },
			URL:          stringer("http://www.iskanie.com"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #205
	case "SafeDNSBot (https://www.safedns.com/searchbot)":
		return &Crawler{
			Pattern:      `SafeDNSBot`,
			Instances:    []string{ "SafeDNSBot (https://www.safedns.com/searchbot)" },
			URL:          stringer("https://www.safedns.com/searchbot"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #206
	case "Mozilla/5.0 (Windows NT 6.1; WOW64) SkypeUriPreview Preview/0.5":
		return &Crawler{
			Pattern:      `SkypeUriPreview`,
			Instances:    []string{ "Mozilla/5.0 (Windows NT 6.1; WOW64) SkypeUriPreview Preview/0.5" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #207
	case "Mozilla/5.0 (compatible; Veoozbot/1.0; +http://www.veooz.com/veoozbot.html)":
		return &Crawler{
			Pattern:      `Veoozbot`,
			Instances:    []string{ "Mozilla/5.0 (compatible; Veoozbot/1.0; +http://www.veooz.com/veoozbot.html)" },
			URL:          stringer("http://www.veooz.com/veoozbot.html"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #208
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
    // #209
	case "Mozilla/5.0 (compatible; redditbot/1.0; +http://www.reddit.com/feedback)":
		return &Crawler{
			Pattern:      `redditbot`,
			Instances:    []string{ "Mozilla/5.0 (compatible; redditbot/1.0; +http://www.reddit.com/feedback)" },
			URL:          stringer("http://www.reddit.com/feedback"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #210
	case "datagnionbot (+http://www.datagnion.com/bot.html)":
		return &Crawler{
			Pattern:      `datagnionbot`,
			Instances:    []string{ "datagnionbot (+http://www.datagnion.com/bot.html)" },
			URL:          stringer("http://www.datagnion.com/bot.html"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #211
	case "Google-Adwords-Instant (+http://www.google.com/adsbot.html)":
		return &Crawler{
			Pattern:      `Google-Adwords-Instant`,
			Instances:    []string{ "Google-Adwords-Instant (+http://www.google.com/adsbot.html)" },
			URL:          stringer("http://www.google.com/adsbot.html"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #212
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
    // #213
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
    // #214
	case "Mozilla/5.0 (compatible;contxbot/1.0)":
		return &Crawler{
			Pattern:      `contxbot`,
			Instances:    []string{ "Mozilla/5.0 (compatible;contxbot/1.0)" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #215
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
    // #216
	case "Mozilla/5.0 (compatible; electricmonk/3.2.0 +https://www.duedil.com/our-crawler/)":
		return &Crawler{
			Pattern:      `electricmonk`,
			Instances:    []string{ "Mozilla/5.0 (compatible; electricmonk/3.2.0 +https://www.duedil.com/our-crawler/)" },
			URL:          stringer("https://www.duedil.com/our-crawler/"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #217
	case "GarlikCrawler/1.2 (http://garlik.com/, crawler@garlik.com)":
		return &Crawler{
			Pattern:      `GarlikCrawler`,
			Instances:    []string{ "GarlikCrawler/1.2 (http://garlik.com/, crawler@garlik.com)" },
			URL:          stringer("http://garlik.com/"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #218
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
    // #219
	case "Mozilla/5.0 (compatible; vebidoobot/1.0; +https://blog.vebidoo.de/vebidoobot/":
		return &Crawler{
			Pattern:      `vebidoobot`,
			Instances:    []string{ "Mozilla/5.0 (compatible; vebidoobot/1.0; +https://blog.vebidoo.de/vebidoobot/" },
			URL:          stringer("https://blog.vebidoo.de/vebidoobot/"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #220
	case "Mozilla/5.0 (compatible; FemtosearchBot/1.0; http://femtosearch.com)":
		return &Crawler{
			Pattern:      `FemtosearchBot`,
			Instances:    []string{ "Mozilla/5.0 (compatible; FemtosearchBot/1.0; http://femtosearch.com)" },
			URL:          stringer("http://femtosearch.com"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #221
	case "Mozilla/5.0 (compatible; Yahoo Link Preview; https://help.yahoo.com/kb/mail/yahoo-link-preview-SLN23615.html)":
		return &Crawler{
			Pattern:      `Yahoo Link Preview`,
			Instances:    []string{ "Mozilla/5.0 (compatible; Yahoo Link Preview; https://help.yahoo.com/kb/mail/yahoo-link-preview-SLN23615.html)" },
			URL:          stringer("https://help.yahoo.com/kb/mail/yahoo-link-preview-SLN23615.html"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #222
	case "Mozilla/5.0 (compatible; MetaJobBot; http://www.metajob.de/crawler)":
		return &Crawler{
			Pattern:      `MetaJobBot`,
			Instances:    []string{ "Mozilla/5.0 (compatible; MetaJobBot; http://www.metajob.de/crawler)" },
			URL:          stringer("http://www.metajob.de/the/crawler"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #223
	case "DomainStatsBot/1.0 (http://domainstats.io/our-bot)":
		return &Crawler{
			Pattern:      `DomainStatsBot`,
			Instances:    []string{ "DomainStatsBot/1.0 (http://domainstats.io/our-bot)" },
			URL:          stringer("http://domainstats.io/our-bot"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #224
	case "mindUpBot (datenbutler.de)":
		return &Crawler{
			Pattern:      `mindUpBot`,
			Instances:    []string{ "mindUpBot (datenbutler.de)" },
			URL:          stringer("http://www.datenbutler.de/"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #225
	case "Mozilla/5.0 (compatible; Daum/4.1; +http://cs.daum.net/faq/15/4118.html?faqId=28966)":
		return &Crawler{
			Pattern:      `Daum\/`,
			Instances:    []string{ "Mozilla/5.0 (compatible; Daum/4.1; +http://cs.daum.net/faq/15/4118.html?faqId=28966)" },
			URL:          stringer("http://cs.daum.net/faq/15/4118.html?faqId=28966"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #226
	case "Jugendschutzprogramm-Crawler; Info: http://www.jugendschutzprogramm.de":
		return &Crawler{
			Pattern:      `Jugendschutzprogramm-Crawler`,
			Instances:    []string{ "Jugendschutzprogramm-Crawler; Info: http://www.jugendschutzprogramm.de" },
			URL:          stringer("http://www.jugendschutzprogramm.de"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #227
	case "Xenu Link Sleuth/1.3.8":
		return &Crawler{
			Pattern:      `Xenu Link Sleuth`,
			Instances:    []string{ "Xenu Link Sleuth/1.3.8" },
			URL:          stringer("http://home.snafu.de/tilman/xenulink.html"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #228
	case "Pcore-HTTP/v0.40.3":
		return &Crawler{
			Pattern:      `Pcore-HTTP`,
			Instances:    []string{ "Pcore-HTTP/v0.40.3" },
			URL:          stringer("https://bitbucket.org/softvisio/pcore/overview"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #229
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
    // #230
	case "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/44.0.2403.125 Safari/537.36 (compatible; KosmioBot/1.0; +http://kosm.io/bot.html)":
		return &Crawler{
			Pattern:      `KosmioBot`,
			Instances:    []string{ "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/44.0.2403.125 Safari/537.36 (compatible; KosmioBot/1.0; +http://kosm.io/bot.html)" },
			URL:          stringer("http://kosm.io/bot.html"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #231
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
    // #232
	case "Mozilla/5.0 (compatible; MSIE 9.0; Windows NT 6.1; Trident/5.0; AppInsights)":
		return &Crawler{
			Pattern:      `AppInsights`,
			Instances:    []string{ "Mozilla/5.0 (compatible; MSIE 9.0; Windows NT 6.1; Trident/5.0; AppInsights)" },
			URL:          stringer("https://docs.microsoft.com/en-us/azure/azure-monitor/app/app-insights-overview"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #233
	case "Mozilla/5.0 (Unknown; Linux x86_64) AppleWebKit/538.1 (KHTML, like Gecko) PhantomJS/2.1.1 Safari/538.1 bl.uk_lddc_renderbot/2.0.0 (+ http://www.bl.uk/aboutus/legaldeposit/websites/websites/faqswebmaster/index.html)":
		return &Crawler{
			Pattern:      `PhantomJS`,
			Instances:    []string{ "Mozilla/5.0 (Unknown; Linux x86_64) AppleWebKit/538.1 (KHTML, like Gecko) PhantomJS/2.1.1 Safari/538.1 bl.uk_lddc_renderbot/2.0.0 (+ http://www.bl.uk/aboutus/legaldeposit/websites/websites/faqswebmaster/index.html)" },
			URL:          stringer("http://phantomjs.org/"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #234
	case "Mozilla/5.0 (compatible; Gowikibot/1.0; +http://www.gowikibot.com)":
		return &Crawler{
			Pattern:      `Gowikibot`,
			Instances:    []string{ "Mozilla/5.0 (compatible; Gowikibot/1.0; +http://www.gowikibot.com)" },
			URL:          stringer("http://www.gowikibot.com"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #235
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
    // #236
	case "Mozilla/5.0 (compatible; Discordbot/2.0; +https://discordapp.com)":
		return &Crawler{
			Pattern:      `Discordbot`,
			Instances:    []string{ "Mozilla/5.0 (compatible; Discordbot/2.0; +https://discordapp.com)" },
			URL:          stringer("https://discordapp.com"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #237
	case "TelegramBot (like TwitterBot)":
		return &Crawler{
			Pattern:      `TelegramBot`,
			Instances:    []string{ "TelegramBot (like TwitterBot)" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #238
	case "Mozilla/5.0 (compatible; Jetslide; +http://jetsli.de/crawler)":
		return &Crawler{
			Pattern:      `Jetslide`,
			Instances:    []string{ "Mozilla/5.0 (compatible; Jetslide; +http://jetsli.de/crawler)" },
			URL:          stringer("http://jetsli.de/crawler"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #239
	case "Mozilla/5.0 (compatible; NewShareCounts.com/1.0; +http://newsharecounts.com/crawler)":
		return &Crawler{
			Pattern:      `newsharecounts`,
			Instances:    []string{ "Mozilla/5.0 (compatible; NewShareCounts.com/1.0; +http://newsharecounts.com/crawler)" },
			URL:          stringer("http://newsharecounts.com/crawler"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #240
	case "Mozilla/5.0 (Windows; U; Windows NT 5.1; en-US; rv:1.8.1.6) Gecko/20070725 Firefox/2.0.0.6 - James BOT - WebCrawler http://cognitiveseo.com/bot.html":
		return &Crawler{
			Pattern:      `James BOT`,
			Instances:    []string{ "Mozilla/5.0 (Windows; U; Windows NT 5.1; en-US; rv:1.8.1.6) Gecko/20070725 Firefox/2.0.0.6 - James BOT - WebCrawler http://cognitiveseo.com/bot.html" },
			URL:          stringer("http://cognitiveseo.com/bot.html"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #241
	case "Barkrowler/0.5.1 (experimenting / debugging - sorry for your logs ) http://www.exensa.com/crawl - admin@exensa.com -- based on BuBiNG",
        "Barkrowler/0.7 (+http://www.exensa.com/crawl)",
        "BarkRowler/0.7 (+http://www.exensa.com/crawling)",
        "Barkrowler/0.9 (+http://www.exensa.com/crawl)":
		return &Crawler{
			Pattern:      `Bark[rR]owler`,
			Instances:    []string{ "Barkrowler/0.5.1 (experimenting / debugging - sorry for your logs ) http://www.exensa.com/crawl - admin@exensa.com -- based on BuBiNG", "Barkrowler/0.7 (+http://www.exensa.com/crawl)", "BarkRowler/0.7 (+http://www.exensa.com/crawling)", "Barkrowler/0.9 (+http://www.exensa.com/crawl)" },
			URL:          stringer("http://www.exensa.com/crawl"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #242
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
    // #243
	case "SocialRankIOBot; http://socialrank.io/about":
		return &Crawler{
			Pattern:      `SocialRankIOBot`,
			Instances:    []string{ "SocialRankIOBot; http://socialrank.io/about" },
			URL:          stringer("http://socialrank.io/about"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #244
	case "Mozilla/5.0 (Windows; U; Windows NT 6.0; en-GB; rv:1.0; trendictionbot0.5.0; trendiction search; http://www.trendiction.de/bot; please let us know of any problems; web at trendiction.com) Gecko/20071127 Firefox/3.0.0.11":
		return &Crawler{
			Pattern:      `trendictionbot`,
			Instances:    []string{ "Mozilla/5.0 (Windows; U; Windows NT 6.0; en-GB; rv:1.0; trendictionbot0.5.0; trendiction search; http://www.trendiction.de/bot; please let us know of any problems; web at trendiction.com) Gecko/20071127 Firefox/3.0.0.11" },
			URL:          stringer("http://www.trendiction.de/bot"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #245
	case "Ocarinabot":
		return &Crawler{
			Pattern:      `Ocarinabot`,
			Instances:    []string{ "Ocarinabot" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #246
	case "Mozilla/5.0 (compatible; epicbot; +http://www.epictions.com/epicbot)":
		return &Crawler{
			Pattern:      `epicbot`,
			Instances:    []string{ "Mozilla/5.0 (compatible; epicbot; +http://www.epictions.com/epicbot)" },
			URL:          stringer("http://www.epictions.com/epicbot"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #247
	case "Mozilla/5.0 (compatible; Primalbot; +https://www.primal.com;)":
		return &Crawler{
			Pattern:      `Primalbot`,
			Instances:    []string{ "Mozilla/5.0 (compatible; Primalbot; +https://www.primal.com;)" },
			URL:          stringer("https://www.primal.com"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #248
	case "Mozilla/5.0 (compatible; DuckDuckGo-Favicons-Bot/1.0; +http://duckduckgo.com)":
		return &Crawler{
			Pattern:      `DuckDuckGo-Favicons-Bot`,
			Instances:    []string{ "Mozilla/5.0 (compatible; DuckDuckGo-Favicons-Bot/1.0; +http://duckduckgo.com)" },
			URL:          stringer("http://duckduckgo.com"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #249
	case "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:49.0) Gecko/20100101 Firefox/49.0 / GnowitNewsbot / Contact information at http://www.gnowit.com":
		return &Crawler{
			Pattern:      `GnowitNewsbot`,
			Instances:    []string{ "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:49.0) Gecko/20100101 Firefox/49.0 / GnowitNewsbot / Contact information at http://www.gnowit.com" },
			URL:          stringer("http://www.gnowit.com"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #250
	case "Mozilla/5.0 (Windows NT 6.3;compatible; Leikibot/1.0; +http://www.leiki.com)":
		return &Crawler{
			Pattern:      `Leikibot`,
			Instances:    []string{ "Mozilla/5.0 (Windows NT 6.3;compatible; Leikibot/1.0; +http://www.leiki.com)" },
			URL:          stringer("http://www.leiki.com"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #251
	case "@LinkArchiver twitter bot":
		return &Crawler{
			Pattern:      `LinkArchiver`,
			Instances:    []string{ "@LinkArchiver twitter bot" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #252
	case "Mozilla/5.0 (compatible; YaK/1.0; http://linkfluence.com/; bot@linkfluence.com)":
		return &Crawler{
			Pattern:      `YaK\/`,
			Instances:    []string{ "Mozilla/5.0 (compatible; YaK/1.0; http://linkfluence.com/; bot@linkfluence.com)" },
			URL:          stringer("http://linkfluence.com"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #253
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
    // #254
	case "Digg Deeper/v1 (http://digg.com/about)":
		return &Crawler{
			Pattern:      `Digg Deeper`,
			Instances:    []string{ "Digg Deeper/v1 (http://digg.com/about)" },
			URL:          stringer("http://digg.com/about"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #255
	case "dcrawl/1.0":
		return &Crawler{
			Pattern:      `dcrawl`,
			Instances:    []string{ "dcrawl/1.0" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #256
	case "Mozilla/5.0 (compatible; Snacktory; +https://github.com/karussell/snacktory)":
		return &Crawler{
			Pattern:      `Snacktory`,
			Instances:    []string{ "Mozilla/5.0 (compatible; Snacktory; +https://github.com/karussell/snacktory)" },
			URL:          stringer("https://github.com/karussell/snacktory"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #257
	case "Mozilla/5.0 (compatible; AndersPinkBot/1.0; +http://anderspink.com/bot.html)":
		return &Crawler{
			Pattern:      `AndersPinkBot`,
			Instances:    []string{ "Mozilla/5.0 (compatible; AndersPinkBot/1.0; +http://anderspink.com/bot.html)" },
			URL:          stringer("http://anderspink.com/bot.html"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #258
	case "Fyrebot/1.0":
		return &Crawler{
			Pattern:      `Fyrebot`,
			Instances:    []string{ "Fyrebot/1.0" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #259
	case "Mozilla/5.0 (compatible; EveryoneSocialBot/1.0; support@everyonesocial.com http://everyonesocial.com/)":
		return &Crawler{
			Pattern:      `EveryoneSocialBot`,
			Instances:    []string{ "Mozilla/5.0 (compatible; EveryoneSocialBot/1.0; support@everyonesocial.com http://everyonesocial.com/)" },
			URL:          stringer("http://everyonesocial.com"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #260
	case "Mediatoolkitbot (complaints@mediatoolkit.com)":
		return &Crawler{
			Pattern:      `Mediatoolkitbot`,
			Instances:    []string{ "Mediatoolkitbot (complaints@mediatoolkit.com)" },
			URL:          stringer("http://mediatoolkit.com"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #261
	case "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_8_2) AppleWebKit/537.13 (KHTML, like Gecko) Chrome/30.0.1599.66 Safari/537.13 Luminator-robots/2.0":
		return &Crawler{
			Pattern:      `Luminator-robots`,
			Instances:    []string{ "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_8_2) AppleWebKit/537.13 (KHTML, like Gecko) Chrome/30.0.1599.66 Safari/537.13 Luminator-robots/2.0" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #262
	case "Mozilla/5.0 (compatible; ExtLinksBot/1.5 +https://extlinks.com/Bot.html)":
		return &Crawler{
			Pattern:      `ExtLinksBot`,
			Instances:    []string{ "Mozilla/5.0 (compatible; ExtLinksBot/1.5 +https://extlinks.com/Bot.html)" },
			URL:          stringer("https://extlinks.com/Bot.html"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #263
	case "Mozilla/5.0 (Windows; U; Windows NT 5.1; en; rv:1.9.0.13) Gecko/2009073022 Firefox/3.5.2 (.NET CLR 3.5.30729) SurveyBot/2.3 (DomainTools)":
		return &Crawler{
			Pattern:      `SurveyBot`,
			Instances:    []string{ "Mozilla/5.0 (Windows; U; Windows NT 5.1; en; rv:1.9.0.13) Gecko/2009073022 Firefox/3.5.2 (.NET CLR 3.5.30729) SurveyBot/2.3 (DomainTools)" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #264
	case "NING/1.0":
		return &Crawler{
			Pattern:      `NING\/`,
			Instances:    []string{ "NING/1.0" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #265
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
    // #266
	case "Nuzzel":
		return &Crawler{
			Pattern:      `Nuzzel`,
			Instances:    []string{ "Nuzzel" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #267
	case "omgili/0.5 +http://omgili.com":
		return &Crawler{
			Pattern:      `omgili`,
			Instances:    []string{ "omgili/0.5 +http://omgili.com" },
			URL:          stringer("http://omgili.com"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #268
	case "PocketParser/2.0 (+https://getpocket.com/pocketparser_ua)":
		return &Crawler{
			Pattern:      `PocketParser`,
			Instances:    []string{ "PocketParser/2.0 (+https://getpocket.com/pocketparser_ua)" },
			URL:          stringer("https://getpocket.com/pocketparser_ua"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #269
	case "YisouSpider":
		return &Crawler{
			Pattern:      `YisouSpider`,
			Instances:    []string{ "YisouSpider" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #270
	case "Mozilla/5.0 (compatible; um-LN/1.0; mailto: techinfo@ubermetrics-technologies.com)":
		return &Crawler{
			Pattern:      `um-LN`,
			Instances:    []string{ "Mozilla/5.0 (compatible; um-LN/1.0; mailto: techinfo@ubermetrics-technologies.com)" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #271
	case "Mozilla/5.0 (compatible; ToutiaoSpider/1.0; http://web.toutiao.com/media_cooperation/;)":
		return &Crawler{
			Pattern:      `ToutiaoSpider`,
			Instances:    []string{ "Mozilla/5.0 (compatible; ToutiaoSpider/1.0; http://web.toutiao.com/media_cooperation/;)" },
			URL:          stringer("http://web.toutiao.com/media_cooperation/"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #272
	case "Mozilla/5.0 (compatible; MuckRack/1.0; +http://muckrack.com)":
		return &Crawler{
			Pattern:      `MuckRack`,
			Instances:    []string{ "Mozilla/5.0 (compatible; MuckRack/1.0; +http://muckrack.com)" },
			URL:          stringer("http://muckrack.com"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #273
	case "Jamie's Spider (http://jamiembrown.com/)":
		return &Crawler{
			Pattern:      `Jamie's Spider`,
			Instances:    []string{ "Jamie's Spider (http://jamiembrown.com/)" },
			URL:          stringer("http://jamiembrown.com/"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #274
	case "AHC/2.0":
		return &Crawler{
			Pattern:      `AHC\/`,
			Instances:    []string{ "AHC/2.0" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #275
	case "Mozilla/5.0 (compatible; NetcraftSurveyAgent/1.0; +info@netcraft.com)":
		return &Crawler{
			Pattern:      `NetcraftSurveyAgent`,
			Instances:    []string{ "Mozilla/5.0 (compatible; NetcraftSurveyAgent/1.0; +info@netcraft.com)" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #276
	case "Mozilla/5.0 (iPhone; CPU iPhone OS 8_3 like Mac OS X) AppleWebKit/600.1.4 (KHTML, like Gecko) Version/8.0 Mobile/12F70 Safari/600.1.4 (compatible; Laserlikebot/0.1)":
		return &Crawler{
			Pattern:      `Laserlikebot`,
			Instances:    []string{ "Mozilla/5.0 (iPhone; CPU iPhone OS 8_3 like Mac OS X) AppleWebKit/600.1.4 (KHTML, like Gecko) Version/8.0 Mobile/12F70 Safari/600.1.4 (compatible; Laserlikebot/0.1)" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #277
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
    // #278
	case "AppEngine-Google; (+http://code.google.com/appengine; appid: example)":
		return &Crawler{
			Pattern:      `AppEngine-Google`,
			Instances:    []string{ "AppEngine-Google; (+http://code.google.com/appengine; appid: example)" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #279
	case "Jetty/9.3.z-SNAPSHOT":
		return &Crawler{
			Pattern:      `Jetty`,
			Instances:    []string{ "Jetty/9.3.z-SNAPSHOT" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #280
	case "Upflow/1.0":
		return &Crawler{
			Pattern:      `Upflow`,
			Instances:    []string{ "Upflow/1.0" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #281
	case "Thinklab (thinklab.com)":
		return &Crawler{
			Pattern:      `Thinklab`,
			Instances:    []string{ "Thinklab (thinklab.com)" },
			URL:          stringer("thinklab.com"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #282
	case "Traackr.com":
		return &Crawler{
			Pattern:      `Traackr.com`,
			Instances:    []string{ "Traackr.com" },
			URL:          stringer("Traackr.com"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #283
	case "Ruby, Twurly v1.1 (http://twurly.org)":
		return &Crawler{
			Pattern:      `Twurly`,
			Instances:    []string{ "Ruby, Twurly v1.1 (http://twurly.org)" },
			URL:          stringer("http://twurly.org"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #284
	case "http.rb/2.2.2 (Mastodon/1.5.1; +https://example-masto-instance.org/)":
		return &Crawler{
			Pattern:      `Mastodon`,
			Instances:    []string{ "http.rb/2.2.2 (Mastodon/1.5.1; +https://example-masto-instance.org/)" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #285
	case "http_get":
		return &Crawler{
			Pattern:      `http_get`,
			Instances:    []string{ "http_get" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #286
	case "Mozilla/5.0 (compatible; DnyzBot/1.0)":
		return &Crawler{
			Pattern:      `DnyzBot`,
			Instances:    []string{ "Mozilla/5.0 (compatible; DnyzBot/1.0)" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #287
	case "Mozilla/5.0 (compatible; botify; http://botify.com)":
		return &Crawler{
			Pattern:      `botify`,
			Instances:    []string{ "Mozilla/5.0 (compatible; botify; http://botify.com)" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #288
	case "Mozilla/5.0 (compatible; 007ac9 Crawler; http://crawler.007ac9.net/)":
		return &Crawler{
			Pattern:      `007ac9 Crawler`,
			Instances:    []string{ "Mozilla/5.0 (compatible; 007ac9 Crawler; http://crawler.007ac9.net/)" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #289
	case "Mozilla/5.0 (compatible; BehloolBot/beta; +http://www.webeaver.com/bot)":
		return &Crawler{
			Pattern:      `BehloolBot`,
			Instances:    []string{ "Mozilla/5.0 (compatible; BehloolBot/beta; +http://www.webeaver.com/bot)" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #290
	case "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.10; rv:41.0) Gecko/20100101 Firefox/55.0 BrandVerity/1.0 (http://www.brandverity.com/why-is-brandverity-visiting-me)":
		return &Crawler{
			Pattern:      `BrandVerity`,
			Instances:    []string{ "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.10; rv:41.0) Gecko/20100101 Firefox/55.0 BrandVerity/1.0 (http://www.brandverity.com/why-is-brandverity-visiting-me)" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #291
	case "check_http/v2.2.1 (nagios-plugins 2.2.1)":
		return &Crawler{
			Pattern:      `check_http`,
			Instances:    []string{ "check_http/v2.2.1 (nagios-plugins 2.2.1)" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #292
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
    // #293
	case "Mozilla/5.0 (compatible; ZumBot/1.0; http://help.zum.com/inquiry)":
		return &Crawler{
			Pattern:      `ZumBot`,
			Instances:    []string{ "Mozilla/5.0 (compatible; ZumBot/1.0; http://help.zum.com/inquiry)" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #294
	case "EZID (EZID link checker; https://ezid.cdlib.org/)":
		return &Crawler{
			Pattern:      `EZID`,
			Instances:    []string{ "EZID (EZID link checker; https://ezid.cdlib.org/)" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #295
	case "ICC-Crawler/2.0 (Mozilla-compatible; ; http://ucri.nict.go.jp/en/icccrawler.html)":
		return &Crawler{
			Pattern:      `ICC-Crawler`,
			Instances:    []string{ "ICC-Crawler/2.0 (Mozilla-compatible; ; http://ucri.nict.go.jp/en/icccrawler.html)" },
			URL:          stringer("http://ucri.nict.go.jp/en/icccrawler.html"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #296
	case "ArchiveTeam ArchiveBot/20170106.02 (wpull 2.0.2)":
		return &Crawler{
			Pattern:      `ArchiveBot`,
			Instances:    []string{ "ArchiveTeam ArchiveBot/20170106.02 (wpull 2.0.2)" },
			URL:          stringer("https://github.com/ArchiveTeam/ArchiveBot"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #297
	case "LCC (+http://corpora.informatik.uni-leipzig.de/crawler_faq.html)":
		return &Crawler{
			Pattern:      `^LCC `,
			Instances:    []string{ "LCC (+http://corpora.informatik.uni-leipzig.de/crawler_faq.html)" },
			URL:          stringer("http://corpora.informatik.uni-leipzig.de/crawler_faq.html"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #298
	case "Mozilla/5.0 (compatible; oBot/2.3.1; +http://filterdb.iss.net/crawler/)":
		return &Crawler{
			Pattern:      `filterdb.iss.net\/crawler`,
			Instances:    []string{ "Mozilla/5.0 (compatible; oBot/2.3.1; +http://filterdb.iss.net/crawler/)" },
			URL:          stringer("http://filterdb.iss.net/crawler/"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #299
	case "BLP_bbot/0.1":
		return &Crawler{
			Pattern:      `BLP_bbot`,
			Instances:    []string{ "BLP_bbot/0.1" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #300
	case "Mozilla/5.0 (compatible; BomboraBot/1.0; +http://www.bombora.com/bot)":
		return &Crawler{
			Pattern:      `BomboraBot`,
			Instances:    []string{ "Mozilla/5.0 (compatible; BomboraBot/1.0; +http://www.bombora.com/bot)" },
			URL:          stringer("http://www.bombora.com/bot"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #301
	case "Buck/2.2; (+https://app.hypefactors.com/media-monitoring/about.html)":
		return &Crawler{
			Pattern:      `Buck\/`,
			Instances:    []string{ "Buck/2.2; (+https://app.hypefactors.com/media-monitoring/about.html)" },
			URL:          stringer("https://app.hypefactors.com/media-monitoring/about.html"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #302
	case "Companybook-Crawler (+https://www.companybooknetworking.com/)":
		return &Crawler{
			Pattern:      `Companybook-Crawler`,
			Instances:    []string{ "Companybook-Crawler (+https://www.companybooknetworking.com/)" },
			URL:          stringer("https://www.companybooknetworking.com/"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #303
	case "Mozilla/5.0 (compatible; Genieo/1.0 http://www.genieo.com/webfilter.html)":
		return &Crawler{
			Pattern:      `Genieo`,
			Instances:    []string{ "Mozilla/5.0 (compatible; Genieo/1.0 http://www.genieo.com/webfilter.html)" },
			URL:          stringer("http://www.genieo.com/webfilter.html"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #304
	case "magpie-crawler/1.1 (U; Linux amd64; en-GB; +http://www.brandwatch.net)":
		return &Crawler{
			Pattern:      `magpie-crawler`,
			Instances:    []string{ "magpie-crawler/1.1 (U; Linux amd64; en-GB; +http://www.brandwatch.net)" },
			URL:          stringer("http://www.brandwatch.net"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #305
	case "MeltwaterNews www.meltwater.com":
		return &Crawler{
			Pattern:      `MeltwaterNews`,
			Instances:    []string{ "MeltwaterNews www.meltwater.com" },
			URL:          stringer("http://www.meltwater.com"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #306
	case "Mozilla/5.0 Moreover/5.1 (+http://www.moreover.com)":
		return &Crawler{
			Pattern:      `Moreover`,
			Instances:    []string{ "Mozilla/5.0 Moreover/5.1 (+http://www.moreover.com)" },
			URL:          stringer("http://www.moreover.com"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #307
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
    // #308
	case "Mozilla/5.0 (compatible; ScoutJet; +http://www.scoutjet.com/)":
		return &Crawler{
			Pattern:      `ScoutJet`,
			Instances:    []string{ "Mozilla/5.0 (compatible; ScoutJet; +http://www.scoutjet.com/)" },
			URL:          stringer("http://www.scoutjet.com/"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #309
	case "sentry/8.22.0 (https://sentry.io)":
		return &Crawler{
			Pattern:      `(^| )sentry\/`,
			Instances:    []string{ "sentry/8.22.0 (https://sentry.io)" },
			URL:          stringer("https://sentry.io"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #310
	case "Mozilla/5.0 (compatible; StorygizeBot; http://www.storygize.com)":
		return &Crawler{
			Pattern:      `StorygizeBot`,
			Instances:    []string{ "Mozilla/5.0 (compatible; StorygizeBot; http://www.storygize.com)" },
			URL:          stringer("http://www.storygize.com"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #311
	case "Mozilla/5.0+(compatible; UptimeRobot/2.0; http://www.uptimerobot.com/)":
		return &Crawler{
			Pattern:      `UptimeRobot`,
			Instances:    []string{ "Mozilla/5.0+(compatible; UptimeRobot/2.0; http://www.uptimerobot.com/)" },
			URL:          stringer("http://www.uptimerobot.com/"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #312
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
    // #313
	case "Mozilla/5.0 (compatible; seoscanners.net/1; +spider@seoscanners.net)":
		return &Crawler{
			Pattern:      `seoscanners`,
			Instances:    []string{ "Mozilla/5.0 (compatible; seoscanners.net/1; +spider@seoscanners.net)" },
			URL:          stringer("http://www.seoscanners.net/"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #314
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
    // #315
	case "Mozilla/5.0 (Linux; U; Android 2.3.4; generic) AppleWebKit/537.36 (KHTML, like Gecko; Google Web Preview) Version/4.0 Mobile Safari/537.36":
		return &Crawler{
			Pattern:      `Google Web Preview`,
			Instances:    []string{ "Mozilla/5.0 (Linux; U; Android 2.3.4; generic) AppleWebKit/537.36 (KHTML, like Gecko; Google Web Preview) Version/4.0 Mobile Safari/537.36" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #316
	case "MauiBot (crawler.feedback+wc@gmail.com)":
		return &Crawler{
			Pattern:      `MauiBot`,
			Instances:    []string{ "MauiBot (crawler.feedback+wc@gmail.com)" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #317
	case "Mozilla/5.0 (compatible; AlphaBot/3.2; +http://alphaseobot.com/bot.html)":
		return &Crawler{
			Pattern:      `AlphaBot`,
			Instances:    []string{ "Mozilla/5.0 (compatible; AlphaBot/3.2; +http://alphaseobot.com/bot.html)" },
			URL:          stringer("http://alphaseobot.com/bot.html"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #318
	case "SBL-BOT (http://sbl.net)":
		return &Crawler{
			Pattern:      `SBL-BOT`,
			Instances:    []string{ "SBL-BOT (http://sbl.net)" },
			URL:          stringer("http://sbl.net"),
			Description:  stringer("Bot of SoftByte BlackWidow"),
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #319
	case "IAS crawler (ias_crawler; http://integralads.com/site-indexing-policy/)":
		return &Crawler{
			Pattern:      `IAS crawler`,
			Instances:    []string{ "IAS crawler (ias_crawler; http://integralads.com/site-indexing-policy/)" },
			URL:          stringer("http://integralads.com/site-indexing-policy/"),
			Description:  stringer("Bot of Integral Ad Science, Inc."),
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #320
	case "Mozilla/5.0 (compatible; adscanner/)":
		return &Crawler{
			Pattern:      `adscanner`,
			Instances:    []string{ "Mozilla/5.0 (compatible; adscanner/)" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #321
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
    // #322
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
    // #323
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
    // #324
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
    // #325
	case "blogmuraBot (+http://www.blogmura.com)":
		return &Crawler{
			Pattern:      `blogmuraBot`,
			Instances:    []string{ "blogmuraBot (+http://www.blogmura.com)" },
			URL:          stringer("http://www.blogmura.com"),
			Description:  stringer("A blog ranking site which links to blogs on just about every theme possible."),
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #326
	case "Bot.AraTurka.com/0.0.1":
		return &Crawler{
			Pattern:      `Bot.AraTurka.com`,
			Instances:    []string{ "Bot.AraTurka.com/0.0.1" },
			URL:          stringer("http://www.araturka.com"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #327
	case "bot-pge.chlooe.com/1.0.0 (+http://www.chlooe.com/)":
		return &Crawler{
			Pattern:      `bot-pge.chlooe.com`,
			Instances:    []string{ "bot-pge.chlooe.com/1.0.0 (+http://www.chlooe.com/)" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #328
	case "Mozilla/5.0 (compatible; BoxcarBot/1.1; +awesome@boxcar.io)":
		return &Crawler{
			Pattern:      `BoxcarBot`,
			Instances:    []string{ "Mozilla/5.0 (compatible; BoxcarBot/1.1; +awesome@boxcar.io)" },
			URL:          stringer("https://boxcar.io/"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #329
	case "BTWebClient/180B(9704)":
		return &Crawler{
			Pattern:      `BTWebClient`,
			Instances:    []string{ "BTWebClient/180B(9704)" },
			URL:          stringer("http://www.utorrent.com/"),
			Description:  stringer("µTorrent BitTorrent Client"),
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #330
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
    // #331
	case "Mozilla/5.0 (compatible; Digincore bot; https://www.digincore.com/crawler.html for rules and instructions.)":
		return &Crawler{
			Pattern:      `Digincore bot`,
			Instances:    []string{ "Mozilla/5.0 (compatible; Digincore bot; https://www.digincore.com/crawler.html for rules and instructions.)" },
			URL:          stringer("http://www.digincore.com/crawler.html"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #332
	case "Disqus/1.0":
		return &Crawler{
			Pattern:      `Disqus`,
			Instances:    []string{ "Disqus/1.0" },
			URL:          stringer("https://disqus.com/"),
			Description:  stringer("validate and quality check pages."),
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #333
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
    // #334
	case "Fetch/2.0a (CMS Detection/Web/SEO analysis tool, see http://guess.scritch.org)":
		return &Crawler{
			Pattern:      `Fetch\/`,
			Instances:    []string{ "Fetch/2.0a (CMS Detection/Web/SEO analysis tool, see http://guess.scritch.org)" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #335
	case "Fever/1.38 (Feed Parser; http://feedafever.com; Allow like Gecko)":
		return &Crawler{
			Pattern:      `Fever`,
			Instances:    []string{ "Fever/1.38 (Feed Parser; http://feedafever.com; Allow like Gecko)" },
			URL:          stringer("http://feedafever.com"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #336
	case "Flamingo_SearchEngine (+http://www.flamingosearch.com/bot)":
		return &Crawler{
			Pattern:      `Flamingo_SearchEngine`,
			Instances:    []string{ "Flamingo_SearchEngine (+http://www.flamingosearch.com/bot)" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #337
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
    // #338
	case "g2reader-bot/1.0 (+http://www.g2reader.com/)":
		return &Crawler{
			Pattern:      `g2reader-bot`,
			Instances:    []string{ "g2reader-bot/1.0 (+http://www.g2reader.com/)" },
			URL:          stringer("http://www.g2reader.com/"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #339
	case "G2 Web Services/1.0 (built with StormCrawler Archetype 1.8; https://www.g2webservices.com/; developers@g2llc.com)":
		return &Crawler{
			Pattern:      `G2 Web Services`,
			Instances:    []string{ "G2 Web Services/1.0 (built with StormCrawler Archetype 1.8; https://www.g2webservices.com/; developers@g2llc.com)" },
			URL:          stringer("https://www.g2webservices.com/"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #340
	case "Mozilla/5.0 (compatible; imrbot/1.10.8 +http://www.mignify.com)":
		return &Crawler{
			Pattern:      `imrbot`,
			Instances:    []string{ "Mozilla/5.0 (compatible; imrbot/1.10.8 +http://www.mignify.com)" },
			URL:          stringer("http://www.mignify.com"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #341
	case "K7MLWCBot/1.0 (+http://www.k7computing.com)":
		return &Crawler{
			Pattern:      `K7MLWCBot`,
			Instances:    []string{ "K7MLWCBot/1.0 (+http://www.k7computing.com)" },
			URL:          stringer("http://www.k7computing.com"),
			Description:  stringer("Virus scanner"),
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #342
	case "Kemvibot/1.0 (http://kemvi.com, marco@kemvi.com)":
		return &Crawler{
			Pattern:      `Kemvibot`,
			Instances:    []string{ "Kemvibot/1.0 (http://kemvi.com, marco@kemvi.com)" },
			URL:          stringer("http://kemvi.com"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #343
	case "Landau-Media-Spider/1.0(http://bots.landaumedia.de/bot.html)":
		return &Crawler{
			Pattern:      `Landau-Media-Spider`,
			Instances:    []string{ "Landau-Media-Spider/1.0(http://bots.landaumedia.de/bot.html)" },
			URL:          stringer("http://bots.landaumedia.de/bot.html"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #344
	case "linkapediabot (+http://www.linkapedia.com)":
		return &Crawler{
			Pattern:      `linkapediabot`,
			Instances:    []string{ "linkapediabot (+http://www.linkapedia.com)" },
			URL:          stringer("http://www.linkapedia.com"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #345
	case "Mozilla/5.0 (compatible; vkShare; +http://vk.com/dev/Share)":
		return &Crawler{
			Pattern:      `vkShare`,
			Instances:    []string{ "Mozilla/5.0 (compatible; vkShare; +http://vk.com/dev/Share)" },
			URL:          stringer("http://vk.com/dev/Share"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #346
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
    // #347
	case "Mozilla/5.0 (compatible; BLEXBot/1.0; +http://webmeup-crawler.com/)":
		return &Crawler{
			Pattern:      `BLEXBot\/`,
			Instances:    []string{ "Mozilla/5.0 (compatible; BLEXBot/1.0; +http://webmeup-crawler.com/)" },
			URL:          stringer("http://webmeup-crawler.com"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #348
	case "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/62.0.3202.75 Safari/537.36 DareBoost":
		return &Crawler{
			Pattern:      `DareBoost`,
			Instances:    []string{ "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/62.0.3202.75 Safari/537.36 DareBoost" },
			URL:          stringer("https://www.dareboost.com/"),
			Description:  stringer("Bot to test, Analyze and Optimize website"),
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #349
	case "Mozilla/5.0 (compatible; ZuperlistBot/1.0)":
		return &Crawler{
			Pattern:      `ZuperlistBot\/`,
			Instances:    []string{ "Mozilla/5.0 (compatible; ZuperlistBot/1.0)" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #350
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
    // #351
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
    // #352
	case "Mozilla/5.0 (Windows; U; Windows NT 5.1; en-US; rv:1.9.1.2) Gecko/20090729 Firefox/3.5.2 (.NET CLR 3.5.30729; Diffbot/0.1; +http://www.diffbot.com)":
		return &Crawler{
			Pattern:      `Diffbot\/`,
			Instances:    []string{ "Mozilla/5.0 (Windows; U; Windows NT 5.1; en-US; rv:1.9.1.2) Gecko/20090729 Firefox/3.5.2 (.NET CLR 3.5.30729; Diffbot/0.1; +http://www.diffbot.com)" },
			URL:          stringer("http://www.diffbot.com"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #353
	case "Mozilla/5.0 (compatible; SEOkicks; +https://www.seokicks.de/robot.html)":
		return &Crawler{
			Pattern:      `SEOkicks`,
			Instances:    []string{ "Mozilla/5.0 (compatible; SEOkicks; +https://www.seokicks.de/robot.html)" },
			URL:          stringer("https://www.seokicks.de/robot.html"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #354
	case "Mozilla/5.0 (compatible; tracemyfile/1.0; +bot@tracemyfile.com)":
		return &Crawler{
			Pattern:      `tracemyfile`,
			Instances:    []string{ "Mozilla/5.0 (compatible; tracemyfile/1.0; +bot@tracemyfile.com)" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #355
	case "Mozilla/5.0 (compatible; Nimbostratus-Bot/v1.3.2; http://cloudsystemnetworks.com)":
		return &Crawler{
			Pattern:      `Nimbostratus-Bot`,
			Instances:    []string{ "Mozilla/5.0 (compatible; Nimbostratus-Bot/v1.3.2; http://cloudsystemnetworks.com)" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #356
	case "Mozilla/5.0 zgrab/0.x":
		return &Crawler{
			Pattern:      `zgrab`,
			Instances:    []string{ "Mozilla/5.0 zgrab/0.x" },
			URL:          stringer("https://zmap.io/"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #357
	case "Mozilla/5.0 (compatible; PR-CY.RU; + https://a.pr-cy.ru)":
		return &Crawler{
			Pattern:      `PR-CY.RU`,
			Instances:    []string{ "Mozilla/5.0 (compatible; PR-CY.RU; + https://a.pr-cy.ru)" },
			URL:          stringer("https://a.pr-cy.ru/"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #358
	case "AdsTxtCrawler/1.0":
		return &Crawler{
			Pattern:      `AdsTxtCrawler`,
			Instances:    []string{ "AdsTxtCrawler/1.0" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #359
	case "Datafeedwatch/2.1.x":
		return &Crawler{
			Pattern:      `Datafeedwatch`,
			Instances:    []string{ "Datafeedwatch/2.1.x" },
			URL:          stringer("https://www.datafeedwatch.com/"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #360
	case "Zabbix":
		return &Crawler{
			Pattern:      `Zabbix`,
			Instances:    []string{ "Zabbix" },
			URL:          stringer("https://www.zabbix.com/documentation/3.4/manual/web_monitoring"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #361
	case "TangibleeBot/1.0.0.0 (http://tangiblee.com/bot)":
		return &Crawler{
			Pattern:      `TangibleeBot`,
			Instances:    []string{ "TangibleeBot/1.0.0.0 (http://tangiblee.com/bot)" },
			URL:          stringer("http://tangiblee.com/bot"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #362
	case "google-xrawler":
		return &Crawler{
			Pattern:      `google-xrawler`,
			Instances:    []string{ "google-xrawler" },
			URL:          stringer("https://webmasters.stackexchange.com/questions/105560/what-is-the-google-xrawler-user-agent-used-for"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #363
	case "axios/0.18.0":
		return &Crawler{
			Pattern:      `axios`,
			Instances:    []string{ "axios/0.18.0" },
			URL:          stringer("https://github.com/axios/axios"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #364
	case "Amazon CloudFront":
		return &Crawler{
			Pattern:      `Amazon CloudFront`,
			Instances:    []string{ "Amazon CloudFront" },
			URL:          stringer("https://aws.amazon.com/cloudfront/"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #365
	case "Pulsepoint XT3 web scraper":
		return &Crawler{
			Pattern:      `Pulsepoint`,
			Instances:    []string{ "Pulsepoint XT3 web scraper" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #366
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
    // #367
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
    // #368
	case "WordupInfoSearch/1.0":
		return &Crawler{
			Pattern:      `WordupInfoSearch`,
			Instances:    []string{ "WordupInfoSearch/1.0" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #369
	case "Mozilla/5.0 (compatible; WebDataStats/1.0 ; +https://webdatastats.com/policy.html)":
		return &Crawler{
			Pattern:      `WebDataStats`,
			Instances:    []string{ "Mozilla/5.0 (compatible; WebDataStats/1.0 ; +https://webdatastats.com/policy.html)" },
			URL:          stringer("https://webdatastats.com/"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #370
	case "Jersey/2.25.1 (HttpUrlConnection 1.8.0_141)":
		return &Crawler{
			Pattern:      `HttpUrlConnection`,
			Instances:    []string{ "Jersey/2.25.1 (HttpUrlConnection 1.8.0_141)" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #371
	case "Mozilla/5.0 (compatible; Seekport Crawler; http://seekport.com/)":
		return &Crawler{
			Pattern:      `Seekport Crawler`,
			Instances:    []string{ "Mozilla/5.0 (compatible; Seekport Crawler; http://seekport.com/)" },
			URL:          stringer("http://seekport.com/"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #372
	case "ZoomBot (Linkbot 1.0 http://suite.seozoom.it/bot.html)":
		return &Crawler{
			Pattern:      `ZoomBot`,
			Instances:    []string{ "ZoomBot (Linkbot 1.0 http://suite.seozoom.it/bot.html)" },
			URL:          stringer("http://suite.seozoom.it/bot.html"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #373
	case "VelenPublicWebCrawler (velen.io)":
		return &Crawler{
			Pattern:      `VelenPublicWebCrawler`,
			Instances:    []string{ "VelenPublicWebCrawler (velen.io)" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #374
	case "MoodleBot/1.0":
		return &Crawler{
			Pattern:      `MoodleBot`,
			Instances:    []string{ "MoodleBot/1.0" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #375
	case "jpg-newsbot/2.0; (+https://vipnytt.no/bots/)":
		return &Crawler{
			Pattern:      `jpg-newsbot`,
			Instances:    []string{ "jpg-newsbot/2.0; (+https://vipnytt.no/bots/)" },
			URL:          stringer("https://vipnytt.no/bots/"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #376
	case "Mozilla/5.0 (Java) outbrain":
		return &Crawler{
			Pattern:      `outbrain`,
			Instances:    []string{ "Mozilla/5.0 (Java) outbrain" },
			URL:          stringer("https://www.outbrain.com/help/advertisers/invalid-url/"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #377
	case "W3C_Validator/1.3":
		return &Crawler{
			Pattern:      `W3C_Validator`,
			Instances:    []string{ "W3C_Validator/1.3" },
			URL:          stringer("https://validator.w3.org/services"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #378
	case "Validator.nu/LV":
		return &Crawler{
			Pattern:      `Validator\.nu`,
			Instances:    []string{ "Validator.nu/LV" },
			URL:          stringer("https://validator.w3.org/services"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #379
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
    // #380
	case "W3C-mobileOK/DDC-1.0":
		return &Crawler{
			Pattern:      `W3C-mobileOK`,
			Instances:    []string{ "W3C-mobileOK/DDC-1.0" },
			URL:          stringer("https://validator.w3.org/services"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #381
	case "W3C_I18n-Checker/1.0":
		return &Crawler{
			Pattern:      `W3C_I18n-Checker`,
			Instances:    []string{ "W3C_I18n-Checker/1.0" },
			URL:          stringer("https://validator.w3.org/services"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #382
	case "FeedValidator/1.3":
		return &Crawler{
			Pattern:      `FeedValidator`,
			Instances:    []string{ "FeedValidator/1.3" },
			URL:          stringer("https://validator.w3.org/services"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #383
	case "Jigsaw/2.3.0 W3C_CSS_Validator_JFouffa/2.0":
		return &Crawler{
			Pattern:      `W3C_CSS_Validator`,
			Instances:    []string{ "Jigsaw/2.3.0 W3C_CSS_Validator_JFouffa/2.0" },
			URL:          stringer("https://validator.w3.org/services"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #384
	case "W3C_Unicorn/1.0":
		return &Crawler{
			Pattern:      `W3C_Unicorn`,
			Instances:    []string{ "W3C_Unicorn/1.0" },
			URL:          stringer("https://validator.w3.org/services"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #385
	case "Mozilla/5.0 (Google-PhysicalWeb)":
		return &Crawler{
			Pattern:      `Google-PhysicalWeb`,
			Instances:    []string{ "Mozilla/5.0 (Google-PhysicalWeb)" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #386
	case "Blackboard Safeassign":
		return &Crawler{
			Pattern:      `Blackboard`,
			Instances:    []string{ "Blackboard Safeassign" },
			URL:          stringer("https://help.blackboard.com/Learn/Administrator/Hosting/Tools_Management/SafeAssign"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #387
	case "Mozilla/5.0 (compatible; ICBot/0.1; +https://ideasandcode.xyz":
		return &Crawler{
			Pattern:      `ICBot\/`,
			Instances:    []string{ "Mozilla/5.0 (compatible; ICBot/0.1; +https://ideasandcode.xyz" },
			URL:          stringer("https://ideasandcode.xyz"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #388
	case "Mozilla/5.0 (compatible; BazQux/2.4; +https://bazqux.com/fetcher; 1 subscribers)":
		return &Crawler{
			Pattern:      `BazQux`,
			Instances:    []string{ "Mozilla/5.0 (compatible; BazQux/2.4; +https://bazqux.com/fetcher; 1 subscribers)" },
			URL:          stringer("https://bazqux.com/fetcher"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #389
	case "Mozilla/5.0 (compatible; Twingly Recon; twingly.com)":
		return &Crawler{
			Pattern:      `Twingly`,
			Instances:    []string{ "Mozilla/5.0 (compatible; Twingly Recon; twingly.com)" },
			URL:          stringer("https://twingly.com"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #390
	case "Mozilla/5.0 (compatible; Rivva; http://rivva.de)":
		return &Crawler{
			Pattern:      `Rivva`,
			Instances:    []string{ "Mozilla/5.0 (compatible; Rivva; http://rivva.de)" },
			URL:          stringer("http://rivva.de"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #391
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
    // #392
	case "Mozilla/5.0 (Windows NT 6.2; WOW64) AppleWebKit/537.22 (KHTML, like Gecko) Chrome/25.0.1364.5 Safari/537.22 +awesomecrawler":
		return &Crawler{
			Pattern:      `awesomecrawler`,
			Instances:    []string{ "Mozilla/5.0 (Windows NT 6.2; WOW64) AppleWebKit/537.22 (KHTML, like Gecko) Chrome/25.0.1364.5 Safari/537.22 +awesomecrawler" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #393
	case "Mozilla/5.0 (compatible; Dataprovider.com)":
		return &Crawler{
			Pattern:      `Dataprovider.com`,
			Instances:    []string{ "Mozilla/5.0 (compatible; Dataprovider.com)" },
			URL:          stringer("https://www.dataprovider.com/"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #394
	case "Mozilla/5.0 (compatible; GroupHigh/1.0; +http://www.grouphigh.com/":
		return &Crawler{
			Pattern:      `GroupHigh\/`,
			Instances:    []string{ "Mozilla/5.0 (compatible; GroupHigh/1.0; +http://www.grouphigh.com/" },
			URL:          stringer("http://www.grouphigh.com/"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #395
	case "Mozilla/5.0 (compatible; theoldreader.com)":
		return &Crawler{
			Pattern:      `theoldreader.com`,
			Instances:    []string{ "Mozilla/5.0 (compatible; theoldreader.com)" },
			URL:          stringer("https://www.theoldreader.com/"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #396
	case "Mozilla/5.0 (compatible; U; AnyEvent-HTTP/2.24; +http://software.schmorp.de/pkg/AnyEvent)":
		return &Crawler{
			Pattern:      `AnyEvent`,
			Instances:    []string{ "Mozilla/5.0 (compatible; U; AnyEvent-HTTP/2.24; +http://software.schmorp.de/pkg/AnyEvent)" },
			URL:          stringer("http://software.schmorp.de/pkg/AnyEvent.html"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #397
	case "Uptimebot.org - Free website monitoring":
		return &Crawler{
			Pattern:      `Uptimebot\.org`,
			Instances:    []string{ "Uptimebot.org - Free website monitoring" },
			URL:          stringer("http://uptimebot.org/"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #398
	case "Mozilla/5.0 (compatible; Nmap Scripting Engine; https://nmap.org/book/nse.html)":
		return &Crawler{
			Pattern:      `Nmap Scripting Engine`,
			Instances:    []string{ "Mozilla/5.0 (compatible; Nmap Scripting Engine; https://nmap.org/book/nse.html)" },
			URL:          stringer("https://nmap.org/book/nse.html"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #399
	case "2ip.ru CMS Detector (https://2ip.ru/cms/)":
		return &Crawler{
			Pattern:      `2ip.ru`,
			Instances:    []string{ "2ip.ru CMS Detector (https://2ip.ru/cms/)" },
			URL:          stringer("https://2ip.ru/cms/"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #400
	case "Clickagy Intelligence Bot v2":
		return &Crawler{
			Pattern:      `Clickagy`,
			Instances:    []string{ "Clickagy Intelligence Bot v2" },
			URL:          stringer("https://www.clickagy.com"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #401
	case "Caliperbot/1.0 (+http://www.conductor.com/caliperbot)":
		return &Crawler{
			Pattern:      `Caliperbot`,
			Instances:    []string{ "Caliperbot/1.0 (+http://www.conductor.com/caliperbot)" },
			URL:          stringer("http://www.conductor.com/caliperbot"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #402
	case "MBCrawler/1.0 (https://monitorbacklinks.com)":
		return &Crawler{
			Pattern:      `MBCrawler`,
			Instances:    []string{ "MBCrawler/1.0 (https://monitorbacklinks.com)" },
			URL:          stringer("https://monitorbacklinks.com"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #403
	case "Mozilla/5.0 (compatible; online-webceo-bot/1.0; +http://online.webceo.com)":
		return &Crawler{
			Pattern:      `online-webceo-bot`,
			Instances:    []string{ "Mozilla/5.0 (compatible; online-webceo-bot/1.0; +http://online.webceo.com)" },
			URL:          stringer("http://online.webceo.com"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #404
	case "B2B Bot":
		return &Crawler{
			Pattern:      `B2B Bot`,
			Instances:    []string{ "B2B Bot" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #405
	case "Mozilla/5.0 (compatible; AddSearchBot/0.9; +http://www.addsearch.com/bot; info@addsearch.com)":
		return &Crawler{
			Pattern:      `AddSearchBot`,
			Instances:    []string{ "Mozilla/5.0 (compatible; AddSearchBot/0.9; +http://www.addsearch.com/bot; info@addsearch.com)" },
			URL:          stringer("http://www.addsearch.com/bot"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #406
	case "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/49.0.2623.75 Safari/537.36 Google Favicon":
		return &Crawler{
			Pattern:      `Google Favicon`,
			Instances:    []string{ "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/49.0.2623.75 Safari/537.36 Google Favicon" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #407
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
    // #408
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
    // #409
	case "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) HeadlessChrome/74.0.3729.169 Safari/537.36":
		return &Crawler{
			Pattern:      `HeadlessChrome`,
			Instances:    []string{ "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) HeadlessChrome/74.0.3729.169 Safari/537.36" },
			URL:          stringer("https://developers.google.com/web/updates/2017/04/headless-chrome"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #410
	case "CheckMarkNetwork/1.0 (+http://www.checkmarknetwork.com/spider.html)":
		return &Crawler{
			Pattern:      `CheckMarkNetwork\/`,
			Instances:    []string{ "CheckMarkNetwork/1.0 (+http://www.checkmarknetwork.com/spider.html)" },
			URL:          stringer("https://www.checkmarknetwork.com/"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #411
	case "Mozilla/5.0 (compatible; Uptimebot/1.0; +http://www.uptime.com/uptimebot)":
		return &Crawler{
			Pattern:      `www\.uptime\.com`,
			Instances:    []string{ "Mozilla/5.0 (compatible; Uptimebot/1.0; +http://www.uptime.com/uptimebot)" },
			URL:          stringer("http://www.uptime.com/uptimebot"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}
    // #412
	case "Mozilla/5.0 (compatible; MSIE 8.0; Windows NT 5.1) Streamline3Bot/1.0",
        "Mozilla/5.0 (Windows NT 6.1; Win64; x64; +https://www.ubtsupport.com/legal/Streamline3Bot.php) Streamline3Bot/1.0":
		return &Crawler{
			Pattern:      `Streamline3Bot\/`,
			Instances:    []string{ "Mozilla/5.0 (compatible; MSIE 8.0; Windows NT 5.1) Streamline3Bot/1.0", "Mozilla/5.0 (Windows NT 6.1; Win64; x64; +https://www.ubtsupport.com/legal/Streamline3Bot.php) Streamline3Bot/1.0" },
			URL:          stringer("https://www.ubtsupport.com/legal/Streamline3Bot.php"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	default:
		return nil
	}
}


var rgxp0 = regexp.MustCompile(`Googlebot\/`)
var rgxp1 = regexp.MustCompile(`Googlebot-Mobile`)
var rgxp2 = regexp.MustCompile(`Googlebot-Image`)
var rgxp3 = regexp.MustCompile(`Googlebot-News`)
var rgxp4 = regexp.MustCompile(`Googlebot-Video`)
var rgxp5 = regexp.MustCompile(`AdsBot-Google([^-]|$)`)
var rgxp6 = regexp.MustCompile(`AdsBot-Google-Mobile`)
var rgxp7 = regexp.MustCompile(`Feedfetcher-Google`)
var rgxp8 = regexp.MustCompile(`Mediapartners-Google`)
var rgxp9 = regexp.MustCompile(`Mediapartners \(Googlebot\)`)
var rgxp10 = regexp.MustCompile(`APIs-Google`)
var rgxp11 = regexp.MustCompile(`bingbot`)
var rgxp12 = regexp.MustCompile(`Slurp`)
var rgxp13 = regexp.MustCompile(`[wW]get`)
var rgxp14 = regexp.MustCompile(`curl`)
var rgxp15 = regexp.MustCompile(`LinkedInBot`)
var rgxp16 = regexp.MustCompile(`Python-urllib`)
var rgxp17 = regexp.MustCompile(`python-requests`)
var rgxp18 = regexp.MustCompile(`libwww-perl`)
var rgxp19 = regexp.MustCompile(`httpunit`)
var rgxp20 = regexp.MustCompile(`nutch`)
var rgxp21 = regexp.MustCompile(`Go-http-client`)
var rgxp22 = regexp.MustCompile(`phpcrawl`)
var rgxp23 = regexp.MustCompile(`msnbot`)
var rgxp24 = regexp.MustCompile(`jyxobot`)
var rgxp25 = regexp.MustCompile(`FAST-WebCrawler`)
var rgxp26 = regexp.MustCompile(`FAST Enterprise Crawler`)
var rgxp27 = regexp.MustCompile(`BIGLOTRON`)
var rgxp28 = regexp.MustCompile(`Teoma`)
var rgxp29 = regexp.MustCompile(`convera`)
var rgxp30 = regexp.MustCompile(`seekbot`)
var rgxp31 = regexp.MustCompile(`Gigabot`)
var rgxp32 = regexp.MustCompile(`Gigablast`)
var rgxp33 = regexp.MustCompile(`exabot`)
var rgxp34 = regexp.MustCompile(`ia_archiver`)
var rgxp35 = regexp.MustCompile(`GingerCrawler`)
var rgxp36 = regexp.MustCompile(`webmon `)
var rgxp37 = regexp.MustCompile(`HTTrack`)
var rgxp38 = regexp.MustCompile(`grub.org`)
var rgxp39 = regexp.MustCompile(`UsineNouvelleCrawler`)
var rgxp40 = regexp.MustCompile(`antibot`)
var rgxp41 = regexp.MustCompile(`netresearchserver`)
var rgxp42 = regexp.MustCompile(`speedy`)
var rgxp43 = regexp.MustCompile(`fluffy`)
var rgxp44 = regexp.MustCompile(`findlink`)
var rgxp45 = regexp.MustCompile(`msrbot`)
var rgxp46 = regexp.MustCompile(`panscient`)
var rgxp47 = regexp.MustCompile(`yacybot`)
var rgxp48 = regexp.MustCompile(`AISearchBot`)
var rgxp49 = regexp.MustCompile(`ips-agent`)
var rgxp50 = regexp.MustCompile(`tagoobot`)
var rgxp51 = regexp.MustCompile(`MJ12bot`)
var rgxp52 = regexp.MustCompile(`woriobot`)
var rgxp53 = regexp.MustCompile(`yanga`)
var rgxp54 = regexp.MustCompile(`buzzbot`)
var rgxp55 = regexp.MustCompile(`mlbot`)
var rgxp56 = regexp.MustCompile(`YandexBot`)
var rgxp57 = regexp.MustCompile(`YandexImages`)
var rgxp58 = regexp.MustCompile(`YandexAccessibilityBot`)
var rgxp59 = regexp.MustCompile(`YandexMobileBot`)
var rgxp60 = regexp.MustCompile(`purebot`)
var rgxp61 = regexp.MustCompile(`Linguee Bot`)
var rgxp62 = regexp.MustCompile(`CyberPatrol`)
var rgxp63 = regexp.MustCompile(`voilabot`)
var rgxp64 = regexp.MustCompile(`Baiduspider`)
var rgxp65 = regexp.MustCompile(`citeseerxbot`)
var rgxp66 = regexp.MustCompile(`spbot`)
var rgxp67 = regexp.MustCompile(`twengabot`)
var rgxp68 = regexp.MustCompile(`postrank`)
var rgxp69 = regexp.MustCompile(`TurnitinBot`)
var rgxp70 = regexp.MustCompile(`scribdbot`)
var rgxp71 = regexp.MustCompile(`page2rss`)
var rgxp72 = regexp.MustCompile(`sitebot`)
var rgxp73 = regexp.MustCompile(`linkdex`)
var rgxp74 = regexp.MustCompile(`Adidxbot`)
var rgxp75 = regexp.MustCompile(`ezooms`)
var rgxp76 = regexp.MustCompile(`dotbot`)
var rgxp77 = regexp.MustCompile(`Mail.RU_Bot`)
var rgxp78 = regexp.MustCompile(`discobot`)
var rgxp79 = regexp.MustCompile(`heritrix`)
var rgxp80 = regexp.MustCompile(`findthatfile`)
var rgxp81 = regexp.MustCompile(`europarchive.org`)
var rgxp82 = regexp.MustCompile(`NerdByNature.Bot`)
var rgxp83 = regexp.MustCompile(`sistrix crawler`)
var rgxp84 = regexp.MustCompile(`Ahrefs(Bot|SiteAudit)`)
var rgxp85 = regexp.MustCompile(`fuelbot`)
var rgxp86 = regexp.MustCompile(`CrunchBot`)
var rgxp87 = regexp.MustCompile(`IndeedBot`)
var rgxp88 = regexp.MustCompile(`mappydata`)
var rgxp89 = regexp.MustCompile(`woobot`)
var rgxp90 = regexp.MustCompile(`ZoominfoBot`)
var rgxp91 = regexp.MustCompile(`PrivacyAwareBot`)
var rgxp92 = regexp.MustCompile(`Multiviewbot`)
var rgxp93 = regexp.MustCompile(`SWIMGBot`)
var rgxp94 = regexp.MustCompile(`Grobbot`)
var rgxp95 = regexp.MustCompile(`eright`)
var rgxp96 = regexp.MustCompile(`Apercite`)
var rgxp97 = regexp.MustCompile(`semanticbot`)
var rgxp98 = regexp.MustCompile(`Aboundex`)
var rgxp99 = regexp.MustCompile(`domaincrawler`)
var rgxp100 = regexp.MustCompile(`wbsearchbot`)
var rgxp101 = regexp.MustCompile(`summify`)
var rgxp102 = regexp.MustCompile(`CCBot`)
var rgxp103 = regexp.MustCompile(`edisterbot`)
var rgxp104 = regexp.MustCompile(`seznambot`)
var rgxp105 = regexp.MustCompile(`ec2linkfinder`)
var rgxp106 = regexp.MustCompile(`gslfbot`)
var rgxp107 = regexp.MustCompile(`aiHitBot`)
var rgxp108 = regexp.MustCompile(`intelium_bot`)
var rgxp109 = regexp.MustCompile(`facebookexternalhit`)
var rgxp110 = regexp.MustCompile(`Yeti`)
var rgxp111 = regexp.MustCompile(`RetrevoPageAnalyzer`)
var rgxp112 = regexp.MustCompile(`lb-spider`)
var rgxp113 = regexp.MustCompile(`Sogou`)
var rgxp114 = regexp.MustCompile(`lssbot`)
var rgxp115 = regexp.MustCompile(`careerbot`)
var rgxp116 = regexp.MustCompile(`wotbox`)
var rgxp117 = regexp.MustCompile(`wocbot`)
var rgxp118 = regexp.MustCompile(`ichiro`)
var rgxp119 = regexp.MustCompile(`DuckDuckBot`)
var rgxp120 = regexp.MustCompile(`lssrocketcrawler`)
var rgxp121 = regexp.MustCompile(`drupact`)
var rgxp122 = regexp.MustCompile(`webcompanycrawler`)
var rgxp123 = regexp.MustCompile(`acoonbot`)
var rgxp124 = regexp.MustCompile(`openindexspider`)
var rgxp125 = regexp.MustCompile(`gnam gnam spider`)
var rgxp126 = regexp.MustCompile(`web-archive-net.com.bot`)
var rgxp127 = regexp.MustCompile(`backlinkcrawler`)
var rgxp128 = regexp.MustCompile(`coccoc`)
var rgxp129 = regexp.MustCompile(`integromedb`)
var rgxp130 = regexp.MustCompile(`content crawler spider`)
var rgxp131 = regexp.MustCompile(`toplistbot`)
var rgxp132 = regexp.MustCompile(`it2media-domain-crawler`)
var rgxp133 = regexp.MustCompile(`ip-web-crawler.com`)
var rgxp134 = regexp.MustCompile(`siteexplorer.info`)
var rgxp135 = regexp.MustCompile(`elisabot`)
var rgxp136 = regexp.MustCompile(`proximic`)
var rgxp137 = regexp.MustCompile(`changedetection`)
var rgxp138 = regexp.MustCompile(`arabot`)
var rgxp139 = regexp.MustCompile(`WeSEE:Search`)
var rgxp140 = regexp.MustCompile(`niki-bot`)
var rgxp141 = regexp.MustCompile(`CrystalSemanticsBot`)
var rgxp142 = regexp.MustCompile(`rogerbot`)
var rgxp143 = regexp.MustCompile(`360Spider`)
var rgxp144 = regexp.MustCompile(`psbot`)
var rgxp145 = regexp.MustCompile(`InterfaxScanBot`)
var rgxp146 = regexp.MustCompile(`CC Metadata Scaper`)
var rgxp147 = regexp.MustCompile(`g00g1e.net`)
var rgxp148 = regexp.MustCompile(`GrapeshotCrawler`)
var rgxp149 = regexp.MustCompile(`urlappendbot`)
var rgxp150 = regexp.MustCompile(`brainobot`)
var rgxp151 = regexp.MustCompile(`fr-crawler`)
var rgxp152 = regexp.MustCompile(`binlar`)
var rgxp153 = regexp.MustCompile(`SimpleCrawler`)
var rgxp154 = regexp.MustCompile(`Twitterbot`)
var rgxp155 = regexp.MustCompile(`cXensebot`)
var rgxp156 = regexp.MustCompile(`smtbot`)
var rgxp157 = regexp.MustCompile(`bnf.fr_bot`)
var rgxp158 = regexp.MustCompile(`A6-Indexer`)
var rgxp159 = regexp.MustCompile(`ADmantX`)
var rgxp160 = regexp.MustCompile(`Facebot`)
var rgxp161 = regexp.MustCompile(`OrangeBot\/`)
var rgxp162 = regexp.MustCompile(`memorybot`)
var rgxp163 = regexp.MustCompile(`AdvBot`)
var rgxp164 = regexp.MustCompile(`MegaIndex`)
var rgxp165 = regexp.MustCompile(`SemanticScholarBot`)
var rgxp166 = regexp.MustCompile(`ltx71`)
var rgxp167 = regexp.MustCompile(`nerdybot`)
var rgxp168 = regexp.MustCompile(`xovibot`)
var rgxp169 = regexp.MustCompile(`BUbiNG`)
var rgxp170 = regexp.MustCompile(`Qwantify`)
var rgxp171 = regexp.MustCompile(`archive.org_bot`)
var rgxp172 = regexp.MustCompile(`Applebot`)
var rgxp173 = regexp.MustCompile(`TweetmemeBot`)
var rgxp174 = regexp.MustCompile(`crawler4j`)
var rgxp175 = regexp.MustCompile(`findxbot`)
var rgxp176 = regexp.MustCompile(`S[eE][mM]rushBot`)
var rgxp177 = regexp.MustCompile(`yoozBot`)
var rgxp178 = regexp.MustCompile(`lipperhey`)
var rgxp179 = regexp.MustCompile(`Y!J`)
var rgxp180 = regexp.MustCompile(`Domain Re-Animator Bot`)
var rgxp181 = regexp.MustCompile(`AddThis`)
var rgxp182 = regexp.MustCompile(`Screaming Frog SEO Spider`)
var rgxp183 = regexp.MustCompile(`MetaURI`)
var rgxp184 = regexp.MustCompile(`Scrapy`)
var rgxp185 = regexp.MustCompile(`Livelap[bB]ot`)
var rgxp186 = regexp.MustCompile(`OpenHoseBot`)
var rgxp187 = regexp.MustCompile(`CapsuleChecker`)
var rgxp188 = regexp.MustCompile(`collection@infegy.com`)
var rgxp189 = regexp.MustCompile(`IstellaBot`)
var rgxp190 = regexp.MustCompile(`DeuSu\/`)
var rgxp191 = regexp.MustCompile(`betaBot`)
var rgxp192 = regexp.MustCompile(`Cliqzbot\/`)
var rgxp193 = regexp.MustCompile(`MojeekBot\/`)
var rgxp194 = regexp.MustCompile(`netEstate NE Crawler`)
var rgxp195 = regexp.MustCompile(`SafeSearch microdata crawler`)
var rgxp196 = regexp.MustCompile(`Gluten Free Crawler\/`)
var rgxp197 = regexp.MustCompile(`Sonic`)
var rgxp198 = regexp.MustCompile(`Sysomos`)
var rgxp199 = regexp.MustCompile(`Trove`)
var rgxp200 = regexp.MustCompile(`deadlinkchecker`)
var rgxp201 = regexp.MustCompile(`Slack-ImgProxy`)
var rgxp202 = regexp.MustCompile(`Embedly`)
var rgxp203 = regexp.MustCompile(`RankActiveLinkBot`)
var rgxp204 = regexp.MustCompile(`iskanie`)
var rgxp205 = regexp.MustCompile(`SafeDNSBot`)
var rgxp206 = regexp.MustCompile(`SkypeUriPreview`)
var rgxp207 = regexp.MustCompile(`Veoozbot`)
var rgxp208 = regexp.MustCompile(`Slackbot`)
var rgxp209 = regexp.MustCompile(`redditbot`)
var rgxp210 = regexp.MustCompile(`datagnionbot`)
var rgxp211 = regexp.MustCompile(`Google-Adwords-Instant`)
var rgxp212 = regexp.MustCompile(`adbeat_bot`)
var rgxp213 = regexp.MustCompile(`WhatsApp`)
var rgxp214 = regexp.MustCompile(`contxbot`)
var rgxp215 = regexp.MustCompile(`pinterest.com.bot`)
var rgxp216 = regexp.MustCompile(`electricmonk`)
var rgxp217 = regexp.MustCompile(`GarlikCrawler`)
var rgxp218 = regexp.MustCompile(`BingPreview\/`)
var rgxp219 = regexp.MustCompile(`vebidoobot`)
var rgxp220 = regexp.MustCompile(`FemtosearchBot`)
var rgxp221 = regexp.MustCompile(`Yahoo Link Preview`)
var rgxp222 = regexp.MustCompile(`MetaJobBot`)
var rgxp223 = regexp.MustCompile(`DomainStatsBot`)
var rgxp224 = regexp.MustCompile(`mindUpBot`)
var rgxp225 = regexp.MustCompile(`Daum\/`)
var rgxp226 = regexp.MustCompile(`Jugendschutzprogramm-Crawler`)
var rgxp227 = regexp.MustCompile(`Xenu Link Sleuth`)
var rgxp228 = regexp.MustCompile(`Pcore-HTTP`)
var rgxp229 = regexp.MustCompile(`moatbot`)
var rgxp230 = regexp.MustCompile(`KosmioBot`)
var rgxp231 = regexp.MustCompile(`pingdom`)
var rgxp232 = regexp.MustCompile(`AppInsights`)
var rgxp233 = regexp.MustCompile(`PhantomJS`)
var rgxp234 = regexp.MustCompile(`Gowikibot`)
var rgxp235 = regexp.MustCompile(`PiplBot`)
var rgxp236 = regexp.MustCompile(`Discordbot`)
var rgxp237 = regexp.MustCompile(`TelegramBot`)
var rgxp238 = regexp.MustCompile(`Jetslide`)
var rgxp239 = regexp.MustCompile(`newsharecounts`)
var rgxp240 = regexp.MustCompile(`James BOT`)
var rgxp241 = regexp.MustCompile(`Bark[rR]owler`)
var rgxp242 = regexp.MustCompile(`TinEye`)
var rgxp243 = regexp.MustCompile(`SocialRankIOBot`)
var rgxp244 = regexp.MustCompile(`trendictionbot`)
var rgxp245 = regexp.MustCompile(`Ocarinabot`)
var rgxp246 = regexp.MustCompile(`epicbot`)
var rgxp247 = regexp.MustCompile(`Primalbot`)
var rgxp248 = regexp.MustCompile(`DuckDuckGo-Favicons-Bot`)
var rgxp249 = regexp.MustCompile(`GnowitNewsbot`)
var rgxp250 = regexp.MustCompile(`Leikibot`)
var rgxp251 = regexp.MustCompile(`LinkArchiver`)
var rgxp252 = regexp.MustCompile(`YaK\/`)
var rgxp253 = regexp.MustCompile(`PaperLiBot`)
var rgxp254 = regexp.MustCompile(`Digg Deeper`)
var rgxp255 = regexp.MustCompile(`dcrawl`)
var rgxp256 = regexp.MustCompile(`Snacktory`)
var rgxp257 = regexp.MustCompile(`AndersPinkBot`)
var rgxp258 = regexp.MustCompile(`Fyrebot`)
var rgxp259 = regexp.MustCompile(`EveryoneSocialBot`)
var rgxp260 = regexp.MustCompile(`Mediatoolkitbot`)
var rgxp261 = regexp.MustCompile(`Luminator-robots`)
var rgxp262 = regexp.MustCompile(`ExtLinksBot`)
var rgxp263 = regexp.MustCompile(`SurveyBot`)
var rgxp264 = regexp.MustCompile(`NING\/`)
var rgxp265 = regexp.MustCompile(`okhttp`)
var rgxp266 = regexp.MustCompile(`Nuzzel`)
var rgxp267 = regexp.MustCompile(`omgili`)
var rgxp268 = regexp.MustCompile(`PocketParser`)
var rgxp269 = regexp.MustCompile(`YisouSpider`)
var rgxp270 = regexp.MustCompile(`um-LN`)
var rgxp271 = regexp.MustCompile(`ToutiaoSpider`)
var rgxp272 = regexp.MustCompile(`MuckRack`)
var rgxp273 = regexp.MustCompile(`Jamie's Spider`)
var rgxp274 = regexp.MustCompile(`AHC\/`)
var rgxp275 = regexp.MustCompile(`NetcraftSurveyAgent`)
var rgxp276 = regexp.MustCompile(`Laserlikebot`)
var rgxp277 = regexp.MustCompile(`Apache-HttpClient`)
var rgxp278 = regexp.MustCompile(`AppEngine-Google`)
var rgxp279 = regexp.MustCompile(`Jetty`)
var rgxp280 = regexp.MustCompile(`Upflow`)
var rgxp281 = regexp.MustCompile(`Thinklab`)
var rgxp282 = regexp.MustCompile(`Traackr.com`)
var rgxp283 = regexp.MustCompile(`Twurly`)
var rgxp284 = regexp.MustCompile(`Mastodon`)
var rgxp285 = regexp.MustCompile(`http_get`)
var rgxp286 = regexp.MustCompile(`DnyzBot`)
var rgxp287 = regexp.MustCompile(`botify`)
var rgxp288 = regexp.MustCompile(`007ac9 Crawler`)
var rgxp289 = regexp.MustCompile(`BehloolBot`)
var rgxp290 = regexp.MustCompile(`BrandVerity`)
var rgxp291 = regexp.MustCompile(`check_http`)
var rgxp292 = regexp.MustCompile(`BDCbot`)
var rgxp293 = regexp.MustCompile(`ZumBot`)
var rgxp294 = regexp.MustCompile(`EZID`)
var rgxp295 = regexp.MustCompile(`ICC-Crawler`)
var rgxp296 = regexp.MustCompile(`ArchiveBot`)
var rgxp297 = regexp.MustCompile(`^LCC `)
var rgxp298 = regexp.MustCompile(`filterdb.iss.net\/crawler`)
var rgxp299 = regexp.MustCompile(`BLP_bbot`)
var rgxp300 = regexp.MustCompile(`BomboraBot`)
var rgxp301 = regexp.MustCompile(`Buck\/`)
var rgxp302 = regexp.MustCompile(`Companybook-Crawler`)
var rgxp303 = regexp.MustCompile(`Genieo`)
var rgxp304 = regexp.MustCompile(`magpie-crawler`)
var rgxp305 = regexp.MustCompile(`MeltwaterNews`)
var rgxp306 = regexp.MustCompile(`Moreover`)
var rgxp307 = regexp.MustCompile(`newspaper\/`)
var rgxp308 = regexp.MustCompile(`ScoutJet`)
var rgxp309 = regexp.MustCompile(`(^| )sentry\/`)
var rgxp310 = regexp.MustCompile(`StorygizeBot`)
var rgxp311 = regexp.MustCompile(`UptimeRobot`)
var rgxp312 = regexp.MustCompile(`OutclicksBot`)
var rgxp313 = regexp.MustCompile(`seoscanners`)
var rgxp314 = regexp.MustCompile(`Hatena`)
var rgxp315 = regexp.MustCompile(`Google Web Preview`)
var rgxp316 = regexp.MustCompile(`MauiBot`)
var rgxp317 = regexp.MustCompile(`AlphaBot`)
var rgxp318 = regexp.MustCompile(`SBL-BOT`)
var rgxp319 = regexp.MustCompile(`IAS crawler`)
var rgxp320 = regexp.MustCompile(`adscanner`)
var rgxp321 = regexp.MustCompile(`Netvibes`)
var rgxp322 = regexp.MustCompile(`acapbot`)
var rgxp323 = regexp.MustCompile(`Baidu-YunGuanCe`)
var rgxp324 = regexp.MustCompile(`bitlybot`)
var rgxp325 = regexp.MustCompile(`blogmuraBot`)
var rgxp326 = regexp.MustCompile(`Bot.AraTurka.com`)
var rgxp327 = regexp.MustCompile(`bot-pge.chlooe.com`)
var rgxp328 = regexp.MustCompile(`BoxcarBot`)
var rgxp329 = regexp.MustCompile(`BTWebClient`)
var rgxp330 = regexp.MustCompile(`ContextAd Bot`)
var rgxp331 = regexp.MustCompile(`Digincore bot`)
var rgxp332 = regexp.MustCompile(`Disqus`)
var rgxp333 = regexp.MustCompile(`Feedly`)
var rgxp334 = regexp.MustCompile(`Fetch\/`)
var rgxp335 = regexp.MustCompile(`Fever`)
var rgxp336 = regexp.MustCompile(`Flamingo_SearchEngine`)
var rgxp337 = regexp.MustCompile(`FlipboardProxy`)
var rgxp338 = regexp.MustCompile(`g2reader-bot`)
var rgxp339 = regexp.MustCompile(`G2 Web Services`)
var rgxp340 = regexp.MustCompile(`imrbot`)
var rgxp341 = regexp.MustCompile(`K7MLWCBot`)
var rgxp342 = regexp.MustCompile(`Kemvibot`)
var rgxp343 = regexp.MustCompile(`Landau-Media-Spider`)
var rgxp344 = regexp.MustCompile(`linkapediabot`)
var rgxp345 = regexp.MustCompile(`vkShare`)
var rgxp346 = regexp.MustCompile(`Siteimprove.com`)
var rgxp347 = regexp.MustCompile(`BLEXBot\/`)
var rgxp348 = regexp.MustCompile(`DareBoost`)
var rgxp349 = regexp.MustCompile(`ZuperlistBot\/`)
var rgxp350 = regexp.MustCompile(`Miniflux\/`)
var rgxp351 = regexp.MustCompile(`Feedspot`)
var rgxp352 = regexp.MustCompile(`Diffbot\/`)
var rgxp353 = regexp.MustCompile(`SEOkicks`)
var rgxp354 = regexp.MustCompile(`tracemyfile`)
var rgxp355 = regexp.MustCompile(`Nimbostratus-Bot`)
var rgxp356 = regexp.MustCompile(`zgrab`)
var rgxp357 = regexp.MustCompile(`PR-CY.RU`)
var rgxp358 = regexp.MustCompile(`AdsTxtCrawler`)
var rgxp359 = regexp.MustCompile(`Datafeedwatch`)
var rgxp360 = regexp.MustCompile(`Zabbix`)
var rgxp361 = regexp.MustCompile(`TangibleeBot`)
var rgxp362 = regexp.MustCompile(`google-xrawler`)
var rgxp363 = regexp.MustCompile(`axios`)
var rgxp364 = regexp.MustCompile(`Amazon CloudFront`)
var rgxp365 = regexp.MustCompile(`Pulsepoint`)
var rgxp366 = regexp.MustCompile(`CloudFlare-AlwaysOnline`)
var rgxp367 = regexp.MustCompile(`Google-Structured-Data-Testing-Tool`)
var rgxp368 = regexp.MustCompile(`WordupInfoSearch`)
var rgxp369 = regexp.MustCompile(`WebDataStats`)
var rgxp370 = regexp.MustCompile(`HttpUrlConnection`)
var rgxp371 = regexp.MustCompile(`Seekport Crawler`)
var rgxp372 = regexp.MustCompile(`ZoomBot`)
var rgxp373 = regexp.MustCompile(`VelenPublicWebCrawler`)
var rgxp374 = regexp.MustCompile(`MoodleBot`)
var rgxp375 = regexp.MustCompile(`jpg-newsbot`)
var rgxp376 = regexp.MustCompile(`outbrain`)
var rgxp377 = regexp.MustCompile(`W3C_Validator`)
var rgxp378 = regexp.MustCompile(`Validator\.nu`)
var rgxp379 = regexp.MustCompile(`W3C-checklink`)
var rgxp380 = regexp.MustCompile(`W3C-mobileOK`)
var rgxp381 = regexp.MustCompile(`W3C_I18n-Checker`)
var rgxp382 = regexp.MustCompile(`FeedValidator`)
var rgxp383 = regexp.MustCompile(`W3C_CSS_Validator`)
var rgxp384 = regexp.MustCompile(`W3C_Unicorn`)
var rgxp385 = regexp.MustCompile(`Google-PhysicalWeb`)
var rgxp386 = regexp.MustCompile(`Blackboard`)
var rgxp387 = regexp.MustCompile(`ICBot\/`)
var rgxp388 = regexp.MustCompile(`BazQux`)
var rgxp389 = regexp.MustCompile(`Twingly`)
var rgxp390 = regexp.MustCompile(`Rivva`)
var rgxp391 = regexp.MustCompile(`Experibot`)
var rgxp392 = regexp.MustCompile(`awesomecrawler`)
var rgxp393 = regexp.MustCompile(`Dataprovider.com`)
var rgxp394 = regexp.MustCompile(`GroupHigh\/`)
var rgxp395 = regexp.MustCompile(`theoldreader.com`)
var rgxp396 = regexp.MustCompile(`AnyEvent`)
var rgxp397 = regexp.MustCompile(`Uptimebot\.org`)
var rgxp398 = regexp.MustCompile(`Nmap Scripting Engine`)
var rgxp399 = regexp.MustCompile(`2ip.ru`)
var rgxp400 = regexp.MustCompile(`Clickagy`)
var rgxp401 = regexp.MustCompile(`Caliperbot`)
var rgxp402 = regexp.MustCompile(`MBCrawler`)
var rgxp403 = regexp.MustCompile(`online-webceo-bot`)
var rgxp404 = regexp.MustCompile(`B2B Bot`)
var rgxp405 = regexp.MustCompile(`AddSearchBot`)
var rgxp406 = regexp.MustCompile(`Google Favicon`)
var rgxp407 = regexp.MustCompile(`HubSpot`)
var rgxp408 = regexp.MustCompile(`Chrome-Lighthouse`)
var rgxp409 = regexp.MustCompile(`HeadlessChrome`)
var rgxp410 = regexp.MustCompile(`CheckMarkNetwork\/`)
var rgxp411 = regexp.MustCompile(`www\.uptime\.com`)
var rgxp412 = regexp.MustCompile(`Streamline3Bot\/`)

// RegexpMatch allows you to identify if and which kind of crawler a User Agent belongs to
// The data are taken from the https://github.com/monperrus/crawler-user-agents/ project
func RegexpMatch(ua string) *Crawler {
	switch {

	case rgxp0.MatchString(ua):
		return &Crawler{
			Pattern:      `Googlebot\/`,
			Instances:    []string{ "Googlebot/2.1 (+http://www.google.com/bot.html)", "Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)", "Mozilla/5.0 (iPhone; CPU iPhone OS 6_0 like Mac OS X) AppleWebKit/536.26 (KHTML, like Gecko) Version/6.0 Mobile/10A5376e Safari/8536.25 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)", "Mozilla/5.0 (iPhone; CPU iPhone OS 8_3 like Mac OS X) AppleWebKit/537.36 (KHTML, like Gecko) Version/8.0 Mobile/12F70 Safari/600.1.4 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)", "Mozilla/5.0 (iPhone; CPU iPhone OS 8_3 like Mac OS X) AppleWebKit/600.1.4 (KHTML, like Gecko) Version/8.0 Mobile/12F70 Safari/600.1.4 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)", "Mozilla/5.0 (Linux; Android 6.0.1; Nexus 5X Build/MMB29P) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/27.0.1453 Mobile Safari/537.36 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)", "Mozilla/5.0 (Linux; Android 6.0.1; Nexus 5X Build/MMB29P) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/41.0.2272.96 Mobile Safari/537.36 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)", "Mozilla/5.0 AppleWebKit/537.36 (KHTML, like Gecko; compatible; Googlebot/2.1; +http://www.google.com/bot.html) Safari/537.36" },
			URL:          stringer("http://www.google.com/bot.html"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp1.MatchString(ua):
		return &Crawler{
			Pattern:      `Googlebot-Mobile`,
			Instances:    []string{ "DoCoMo/2.0 N905i(c100;TB;W24H16) (compatible; Googlebot-Mobile/2.1; +http://www.google.com/bot.html)", "Mozilla/5.0 (iPhone; CPU iPhone OS 6_0 like Mac OS X) AppleWebKit/536.26 (KHTML, like Gecko) Version/6.0 Mobile/10A5376e Safari/8536.25 (compatible; Googlebot-Mobile/2.1; +http://www.google.com/bot.html)", "Mozilla/5.0 (iPhone; U; CPU iPhone OS 4_1 like Mac OS X; en-us) AppleWebKit/532.9 (KHTML, like Gecko) Version/4.0.5 Mobile/8B117 Safari/6531.22.7 (compatible; Googlebot-Mobile/2.1; +http://www.google.com/bot.html)", "Nokia6820/2.0 (4.83) Profile/MIDP-1.0 Configuration/CLDC-1.0 (compatible; Googlebot-Mobile/2.1; +http://www.google.com/bot.html)", "SAMSUNG-SGH-E250/1.0 Profile/MIDP-2.0 Configuration/CLDC-1.1 UP.Browser/6.2.3.3.c.1.101 (GUI) MMP/2.0 (compatible; Googlebot-Mobile/2.1; +http://www.google.com/bot.html)" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp2.MatchString(ua):
		return &Crawler{
			Pattern:      `Googlebot-Image`,
			Instances:    []string{ "Googlebot-Image/1.0" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp3.MatchString(ua):
		return &Crawler{
			Pattern:      `Googlebot-News`,
			Instances:    []string{ "Googlebot-News" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp4.MatchString(ua):
		return &Crawler{
			Pattern:      `Googlebot-Video`,
			Instances:    []string{ "Googlebot-Video/1.0" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp5.MatchString(ua):
		return &Crawler{
			Pattern:      `AdsBot-Google([^-]|$)`,
			Instances:    []string{ "AdsBot-Google (+http://www.google.com/adsbot.html)" },
			URL:          stringer("https://support.google.com/webmasters/answer/1061943?hl=en"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp6.MatchString(ua):
		return &Crawler{
			Pattern:      `AdsBot-Google-Mobile`,
			Instances:    []string{ "AdsBot-Google-Mobile-Apps", "Mozilla/5.0 (Linux; Android 5.0; SM-G920A) AppleWebKit (KHTML, like Gecko) Chrome Mobile Safari (compatible; AdsBot-Google-Mobile; +http://www.google.com/mobile/adsbot.html)", "Mozilla/5.0 (iPhone; CPU iPhone OS 9_1 like Mac OS X) AppleWebKit/601.1.46 (KHTML, like Gecko) Version/9.0 Mobile/13B143 Safari/601.1 (compatible; AdsBot-Google-Mobile; +http://www.google.com/mobile/adsbot.html)" },
			URL:          stringer("https://support.google.com/adwords/answer/2404197"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp7.MatchString(ua):
		return &Crawler{
			Pattern:      `Feedfetcher-Google`,
			Instances:    []string{ "Feedfetcher-Google; (+http://www.google.com/feedfetcher.html; 1 subscribers; feed-id=728742641706423)" },
			URL:          stringer("https://support.google.com/webmasters/answer/178852"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp8.MatchString(ua):
		return &Crawler{
			Pattern:      `Mediapartners-Google`,
			Instances:    []string{ "Mediapartners-Google", "Mozilla/5.0 (compatible; MSIE or Firefox mutant; not on Windows server;) Daumoa/4.0 (Following Mediapartners-Google)", "Mozilla/5.0 (iPhone; U; CPU iPhone OS 10_0 like Mac OS X; en-us) AppleWebKit/602.1.38 (KHTML, like Gecko) Version/10.0 Mobile/14A5297c Safari/602.1 (compatible; Mediapartners-Google/2.1; +http://www.google.com/bot.html)", "Mozilla/5.0 (iPhone; U; CPU iPhone OS 4_1 like Mac OS X; en-us) AppleWebKit/532.9 (KHTML, like Gecko) Version/4.0.5 Mobile/8B117 Safari/6531.22.7 (compatible; Mediapartners-Google/2.1; +http://www.google.com/bot.html)" },
			URL:          stringer("https://support.google.com/webmasters/answer/1061943?hl=en"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp9.MatchString(ua):
		return &Crawler{
			Pattern:      `Mediapartners \(Googlebot\)`,
			Instances:    []string{  },
			URL:          stringer("https://support.google.com/webmasters/answer/1061943?hl=en"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp10.MatchString(ua):
		return &Crawler{
			Pattern:      `APIs-Google`,
			Instances:    []string{ "APIs-Google (+https://developers.google.com/webmasters/APIs-Google.html)" },
			URL:          stringer("https://support.google.com/webmasters/answer/1061943?hl=en"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp11.MatchString(ua):
		return &Crawler{
			Pattern:      `bingbot`,
			Instances:    []string{ "Mozilla/5.0 (Windows Phone 8.1; ARM; Trident/7.0; Touch; rv:11.0; IEMobile/11.0; NOKIA; Lumia 530) like Gecko (compatible; adidxbot/2.0; +http://www.bing.com/bingbot.htm)", "Mozilla/5.0 (compatible; adidxbot/2.0;  http://www.bing.com/bingbot.htm)", "Mozilla/5.0 (compatible; adidxbot/2.0; +http://www.bing.com/bingbot.htm)", "Mozilla/5.0 (compatible; bingbot/2.0;  http://www.bing.com/bingbot.htm)", "Mozilla/5.0 (compatible; bingbot/2.0; +http://www.bing.com/bingbot.htm", "Mozilla/5.0 (compatible; bingbot/2.0; +http://www.bing.com/bingbot.htm)", "Mozilla/5.0 (compatible; bingbot/2.0; +http://www.bing.com/bingbot.htm) SitemapProbe", "Mozilla/5.0 (iPhone; CPU iPhone OS 7_0 like Mac OS X) AppleWebKit/537.51.1 (KHTML, like Gecko) Version/7.0 Mobile/11A465 Safari/9537.53 (compatible; adidxbot/2.0;  http://www.bing.com/bingbot.htm)", "Mozilla/5.0 (iPhone; CPU iPhone OS 7_0 like Mac OS X) AppleWebKit/537.51.1 (KHTML, like Gecko) Version/7.0 Mobile/11A465 Safari/9537.53 (compatible; adidxbot/2.0; +http://www.bing.com/bingbot.htm)", "Mozilla/5.0 (iPhone; CPU iPhone OS 7_0 like Mac OS X) AppleWebKit/537.51.1 (KHTML, like Gecko) Version/7.0 Mobile/11A465 Safari/9537.53 (compatible; bingbot/2.0;  http://www.bing.com/bingbot.htm)", "Mozilla/5.0 (iPhone; CPU iPhone OS 7_0 like Mac OS X) AppleWebKit/537.51.1 (KHTML, like Gecko) Version/7.0 Mobile/11A465 Safari/9537.53 (compatible; bingbot/2.0; +http://www.bing.com/bingbot.htm)", "Mozilla/5.0 (seoanalyzer; compatible; bingbot/2.0; +http://www.bing.com/bingbot.htm)" },
			URL:          stringer("http://www.bing.com/bingbot.htm"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp12.MatchString(ua):
		return &Crawler{
			Pattern:      `Slurp`,
			Instances:    []string{ "Mozilla/5.0 (compatible; Yahoo! Slurp/3.0; http://help.yahoo.com/help/us/ysearch/slurp)", "Mozilla/5.0 (compatible; Yahoo! Slurp; http://help.yahoo.com/help/us/ysearch/slurp)", "Mozilla/5.0 (compatible; Yahoo! Slurp China; http://misc.yahoo.com.cn/help.html)" },
			URL:          stringer("http://help.yahoo.com/help/us/ysearch/slurp"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp13.MatchString(ua):
		return &Crawler{
			Pattern:      `[wW]get`,
			Instances:    []string{ "WGETbot/1.0 (+http://wget.alanreed.org)", "Wget/1.14 (linux-gnu)" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp14.MatchString(ua):
		return &Crawler{
			Pattern:      `curl`,
			Instances:    []string{ "eCairn-Grabber/1.0 (+http://ecairn.com/grabber) curl/7.15" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp15.MatchString(ua):
		return &Crawler{
			Pattern:      `LinkedInBot`,
			Instances:    []string{ "LinkedInBot/1.0 (compatible; Mozilla/5.0; Jakarta Commons-HttpClient/3.1 +http://www.linkedin.com)", "LinkedInBot/1.0 (compatible; Mozilla/5.0; Jakarta Commons-HttpClient/4.3 +http://www.linkedin.com)" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp16.MatchString(ua):
		return &Crawler{
			Pattern:      `Python-urllib`,
			Instances:    []string{ "Python-urllib/2.5", "Python-urllib/2.6", "Python-urllib/2.7", "Python-urllib/3.1", "Python-urllib/3.2", "Python-urllib/3.3", "Python-urllib/3.4", "Python-urllib/3.5", "Python-urllib/3.6" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp17.MatchString(ua):
		return &Crawler{
			Pattern:      `python-requests`,
			Instances:    []string{ "python-requests/2.18.4" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp18.MatchString(ua):
		return &Crawler{
			Pattern:      `libwww-perl`,
			Instances:    []string{ "2Bone_LinkChecker/1.0 libwww-perl/6.03", "2Bone_LinkChkr/1.0 libwww-perl/6.03", "amibot - http://www.amidalla.de - tech@amidalla.com libwww-perl/5.831" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp19.MatchString(ua):
		return &Crawler{
			Pattern:      `httpunit`,
			Instances:    []string{ "httpunit/1.x" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp20.MatchString(ua):
		return &Crawler{
			Pattern:      `nutch`,
			Instances:    []string{ "NutchCVS/0.7.1 (Nutch; http://lucene.apache.org/nutch/bot.html; nutch-agent@lucene.apache.org)", "istellabot-nutch/Nutch-1.10" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp21.MatchString(ua):
		return &Crawler{
			Pattern:      `Go-http-client`,
			Instances:    []string{ "Go-http-client/1.1" },
			URL:          stringer("https://golang.org/pkg/net/http/"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp22.MatchString(ua):
		return &Crawler{
			Pattern:      `phpcrawl`,
			Instances:    []string{ "phpcrawl" },
			URL:          stringer("http://phpcrawl.cuab.de/"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp23.MatchString(ua):
		return &Crawler{
			Pattern:      `msnbot`,
			Instances:    []string{ "adidxbot/1.1 (+http://search.msn.com/msnbot.htm)", "adidxbot/2.0 (+http://search.msn.com/msnbot.htm)", "librabot/1.0 (+http://search.msn.com/msnbot.htm)", "librabot/2.0 (+http://search.msn.com/msnbot.htm)", "msnbot-NewsBlogs/2.0b (+http://search.msn.com/msnbot.htm)", "msnbot-UDiscovery/2.0b (+http://search.msn.com/msnbot.htm)", "msnbot-media/1.0 (+http://search.msn.com/msnbot.htm)", "msnbot-media/1.1 (+http://search.msn.com/msnbot.htm)", "msnbot-media/2.0b (+http://search.msn.com/msnbot.htm)", "msnbot/1.0 (+http://search.msn.com/msnbot.htm)", "msnbot/1.1 (+http://search.msn.com/msnbot.htm)", "msnbot/2.0b (+http://search.msn.com/msnbot.htm)", "msnbot/2.0b (+http://search.msn.com/msnbot.htm).", "msnbot/2.0b (+http://search.msn.com/msnbot.htm)._" },
			URL:          stringer("http://search.msn.com/msnbot.htm"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp24.MatchString(ua):
		return &Crawler{
			Pattern:      `jyxobot`,
			Instances:    []string{  },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp25.MatchString(ua):
		return &Crawler{
			Pattern:      `FAST-WebCrawler`,
			Instances:    []string{ "FAST-WebCrawler/3.6/FirstPage (atw-crawler at fast dot no;http://fast.no/support/crawler.asp)", "FAST-WebCrawler/3.7 (atw-crawler at fast dot no; http://fast.no/support/crawler.asp)", "FAST-WebCrawler/3.7/FirstPage (atw-crawler at fast dot no;http://fast.no/support/crawler.asp)", "FAST-WebCrawler/3.8" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp26.MatchString(ua):
		return &Crawler{
			Pattern:      `FAST Enterprise Crawler`,
			Instances:    []string{ "FAST Enterprise Crawler 6 / Scirus scirus-crawler@fast.no; http://www.scirus.com/srsapp/contactus/", "FAST Enterprise Crawler 6 used by Schibsted (webcrawl@schibstedsok.no)" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp27.MatchString(ua):
		return &Crawler{
			Pattern:      `BIGLOTRON`,
			Instances:    []string{ "BIGLOTRON (Beta 2;GNU/Linux)" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp28.MatchString(ua):
		return &Crawler{
			Pattern:      `Teoma`,
			Instances:    []string{ "Mozilla/2.0 (compatible; Ask Jeeves/Teoma; +http://sp.ask.com/docs/about/tech_crawling.html)", "Mozilla/2.0 (compatible; Ask Jeeves/Teoma; +http://about.ask.com/en/docs/about/webmasters.shtml)" },
			URL:          stringer("http://about.ask.com/en/docs/about/webmasters.shtml"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp29.MatchString(ua):
		return &Crawler{
			Pattern:      `convera`,
			Instances:    []string{ "ConveraCrawler/0.9e (+http://ews.converasearch.com/crawl.htm)" },
			URL:          stringer("http://ews.converasearch.com/crawl.htm"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp30.MatchString(ua):
		return &Crawler{
			Pattern:      `seekbot`,
			Instances:    []string{ "Seekbot/1.0 (http://www.seekbot.net/bot.html) RobotsTxtFetcher/1.2" },
			URL:          stringer("http://www.seekbot.net/bot.html"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp31.MatchString(ua):
		return &Crawler{
			Pattern:      `Gigabot`,
			Instances:    []string{ "Gigabot/1.0", "Gigabot/2.0 (http://www.gigablast.com/spider.html)" },
			URL:          stringer("http://www.gigablast.com/spider.html"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp32.MatchString(ua):
		return &Crawler{
			Pattern:      `Gigablast`,
			Instances:    []string{ "GigablastOpenSource/1.0" },
			URL:          stringer("https://github.com/gigablast/open-source-search-engine"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp33.MatchString(ua):
		return &Crawler{
			Pattern:      `exabot`,
			Instances:    []string{ "Mozilla/5.0 (compatible; Alexabot/1.0; +http://www.alexa.com/help/certifyscan; certifyscan@alexa.com)", "Mozilla/5.0 (compatible; Exabot PyExalead/3.0; +http://www.exabot.com/go/robot)", "Mozilla/5.0 (compatible; Exabot-Images/3.0; +http://www.exabot.com/go/robot)", "Mozilla/5.0 (compatible; Exabot/3.0 (BiggerBetter); +http://www.exabot.com/go/robot)", "Mozilla/5.0 (compatible; Exabot/3.0; +http://www.exabot.com/go/robot)", "Mozilla/5.0 (compatible; Exabot/3.0;  http://www.exabot.com/go/robot)" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp34.MatchString(ua):
		return &Crawler{
			Pattern:      `ia_archiver`,
			Instances:    []string{ "ia_archiver (+http://www.alexa.com/site/help/webmasters; crawler@alexa.com)", "ia_archiver-web.archive.org" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp35.MatchString(ua):
		return &Crawler{
			Pattern:      `GingerCrawler`,
			Instances:    []string{ "GingerCrawler/1.0 (Language Assistant for Dyslexics; www.gingersoftware.com/crawler_agent.htm; support at ginger software dot com)" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp36.MatchString(ua):
		return &Crawler{
			Pattern:      `webmon `,
			Instances:    []string{  },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp37.MatchString(ua):
		return &Crawler{
			Pattern:      `HTTrack`,
			Instances:    []string{ "Mozilla/4.5 (compatible; HTTrack 3.0x; Windows 98)" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp38.MatchString(ua):
		return &Crawler{
			Pattern:      `grub.org`,
			Instances:    []string{ "Mozilla/4.0 (compatible; grub-client-0.3.0; Crawl your own stuff with http://grub.org)", "Mozilla/4.0 (compatible; grub-client-1.0.4; Crawl your own stuff with http://grub.org)", "Mozilla/4.0 (compatible; grub-client-1.0.5; Crawl your own stuff with http://grub.org)", "Mozilla/4.0 (compatible; grub-client-1.0.6; Crawl your own stuff with http://grub.org)", "Mozilla/4.0 (compatible; grub-client-1.0.7; Crawl your own stuff with http://grub.org)", "Mozilla/4.0 (compatible; grub-client-1.1.1; Crawl your own stuff with http://grub.org)", "Mozilla/4.0 (compatible; grub-client-1.2.1; Crawl your own stuff with http://grub.org)", "Mozilla/4.0 (compatible; grub-client-1.3.1; Crawl your own stuff with http://grub.org)", "Mozilla/4.0 (compatible; grub-client-1.3.7; Crawl your own stuff with http://grub.org)", "Mozilla/4.0 (compatible; grub-client-1.4.3; Crawl your own stuff with http://grub.org)", "Mozilla/4.0 (compatible; grub-client-1.5.3; Crawl your own stuff with http://grub.org)" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp39.MatchString(ua):
		return &Crawler{
			Pattern:      `UsineNouvelleCrawler`,
			Instances:    []string{  },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp40.MatchString(ua):
		return &Crawler{
			Pattern:      `antibot`,
			Instances:    []string{  },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp41.MatchString(ua):
		return &Crawler{
			Pattern:      `netresearchserver`,
			Instances:    []string{  },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp42.MatchString(ua):
		return &Crawler{
			Pattern:      `speedy`,
			Instances:    []string{ "Mozilla/5.0 (Windows; U; Windows NT 5.1; en-US) Speedy Spider (http://www.entireweb.com/about/search_tech/speedy_spider/)", "Mozilla/5.0 (Windows; U; Windows NT 5.1; en-US) Speedy Spider for SpeedyAds (http://www.entireweb.com/about/search_tech/speedy_spider/)", "Mozilla/5.0 (compatible; Speedy Spider; http://www.entireweb.com/about/search_tech/speedy_spider/)", "Speedy Spider (Entireweb; Beta/1.2; http://www.entireweb.com/about/search_tech/speedyspider/)", "Speedy Spider (http://www.entireweb.com/about/search_tech/speedy_spider/)" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp43.MatchString(ua):
		return &Crawler{
			Pattern:      `fluffy`,
			Instances:    []string{  },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp44.MatchString(ua):
		return &Crawler{
			Pattern:      `findlink`,
			Instances:    []string{ "findlinks/1.0 (+http://wortschatz.uni-leipzig.de/findlinks/)", "findlinks/1.1.3-beta8 (+http://wortschatz.uni-leipzig.de/findlinks/)", "findlinks/1.1.3-beta9 (+http://wortschatz.uni-leipzig.de/findlinks/)", "findlinks/1.1.5-beta7 (+http://wortschatz.uni-leipzig.de/findlinks/)", "findlinks/1.1.6-beta1 (+http://wortschatz.uni-leipzig.de/findlinks/)", "findlinks/1.1.6-beta1 (+http://wortschatz.uni-leipzig.de/findlinks/; YaCy 0.1; yacy.net)", "findlinks/1.1.6-beta2 (+http://wortschatz.uni-leipzig.de/findlinks/)", "findlinks/1.1.6-beta3 (+http://wortschatz.uni-leipzig.de/findlinks/)", "findlinks/1.1.6-beta4 (+http://wortschatz.uni-leipzig.de/findlinks/)", "findlinks/1.1.6-beta5 (+http://wortschatz.uni-leipzig.de/findlinks/)", "findlinks/1.1.6-beta6 (+http://wortschatz.uni-leipzig.de/findlinks/)", "findlinks/2.0 (+http://wortschatz.uni-leipzig.de/findlinks/)", "findlinks/2.0.1 (+http://wortschatz.uni-leipzig.de/findlinks/)", "findlinks/2.0.2 (+http://wortschatz.uni-leipzig.de/findlinks/)", "findlinks/2.0.4 (+http://wortschatz.uni-leipzig.de/findlinks/)", "findlinks/2.0.5 (+http://wortschatz.uni-leipzig.de/findlinks/)", "findlinks/2.0.9 (+http://wortschatz.uni-leipzig.de/findlinks/)", "findlinks/2.1 (+http://wortschatz.uni-leipzig.de/findlinks/)", "findlinks/2.1.3 (+http://wortschatz.uni-leipzig.de/findlinks/)", "findlinks/2.1.5 (+http://wortschatz.uni-leipzig.de/findlinks/)", "findlinks/2.2 (+http://wortschatz.uni-leipzig.de/findlinks/)", "findlinks/2.5 (+http://wortschatz.uni-leipzig.de/findlinks/)", "findlinks/2.6 (+http://wortschatz.uni-leipzig.de/findlinks/)" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp45.MatchString(ua):
		return &Crawler{
			Pattern:      `msrbot`,
			Instances:    []string{  },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp46.MatchString(ua):
		return &Crawler{
			Pattern:      `panscient`,
			Instances:    []string{ "panscient.com" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp47.MatchString(ua):
		return &Crawler{
			Pattern:      `yacybot`,
			Instances:    []string{ "yacybot (/global; amd64 FreeBSD 10.3-RELEASE; java 1.8.0_77; GMT/en) http://yacy.net/bot.html", "yacybot (/global; amd64 FreeBSD 10.3-RELEASE-p7; java 1.7.0_95; GMT/en) http://yacy.net/bot.html", "yacybot (-global; amd64 FreeBSD 9.2-RELEASE-p10; java 1.7.0_65; Europe/en) http://yacy.net/bot.html", "yacybot (/global; amd64 Linux 2.6.32-042stab093.4; java 1.7.0_65; Etc/en) http://yacy.net/bot.html", "yacybot (/global; amd64 Linux 2.6.32-042stab094.8; java 1.7.0_79; America/en) http://yacy.net/bot.html", "yacybot (/global; amd64 Linux 2.6.32-042stab108.8; java 1.7.0_91; America/en) http://yacy.net/bot.html", "yacybot (-global; amd64 Linux 2.6.32-042stab111.11; java 1.7.0_79; Europe/en) http://yacy.net/bot.html", "yacybot (-global; amd64 Linux 2.6.32-042stab116.1; java 1.7.0_79; Europe/en) http://yacy.net/bot.html", "yacybot (/global; amd64 Linux 2.6.32-573.3.1.el6.x86_64; java 1.7.0_85; Europe/en) http://yacy.net/bot.html", "yacybot (-global; amd64 Linux 3.10.0-229.4.2.el7.x86_64; java 1.7.0_79; Europe/en) http://yacy.net/bot.html", "yacybot (-global; amd64 Linux 3.10.0-229.4.2.el7.x86_64; java 1.8.0_45; Europe/en) http://yacy.net/bot.html", "yacybot (/global; amd64 Linux 3.10.0-229.7.2.el7.x86_64; java 1.8.0_45; Europe/en) http://yacy.net/bot.html", "yacybot (/global; amd64 Linux 3.10.0-327.22.2.el7.x86_64; java 1.7.0_101; Etc/en) http://yacy.net/bot.html", "yacybot (/global; amd64 Linux 3.11.10-21-desktop; java 1.7.0_51; America/en) http://yacy.net/bot.html", "yacybot (/global; amd64 Linux 3.12.1; java 1.7.0_65; Europe/en) http://yacy.net/bot.html", "yacybot (/global; amd64 Linux 3.13.0-042stab093.4; java 1.7.0_79; Europe/de) http://yacy.net/bot.html", "yacybot (/global; amd64 Linux 3.13.0-042stab093.4; java 1.7.0_79; Europe/en) http://yacy.net/bot.html", "yacybot (/global; amd64 Linux 3.13.0-45-generic; java 1.7.0_75; Europe/en) http://yacy.net/bot.html", "yacybot (-global; amd64 Linux 3.13.0-61-generic; java 1.7.0_79; Europe/en) http://yacy.net/bot.html", "yacybot (/global; amd64 Linux 3.13.0-74-generic; java 1.7.0_91; Europe/en) http://yacy.net/bot.html", "yacybot (/global; amd64 Linux 3.13.0-83-generic; java 1.7.0_95; Europe/de) http://yacy.net/bot.html", "yacybot (/global; amd64 Linux 3.13.0-83-generic; java 1.7.0_95; Europe/en) http://yacy.net/bot.html", "yacybot (/global; amd64 Linux 3.13.0-85-generic; java 1.7.0_101; Europe/en) http://yacy.net/bot.html", "yacybot (/global; amd64 Linux 3.13.0-85-generic; java 1.7.0_95; Europe/en) http://yacy.net/bot.html", "yacybot (/global; amd64 Linux 3.13.0-88-generic; java 1.7.0_101; Europe/en) http://yacy.net/bot.html", "yacybot (/global; amd64 Linux 3.14-0.bpo.1-amd64; java 1.7.0_55; Europe/de) http://yacy.net/bot.html", "yacybot (/global; amd64 Linux 3.14.32-xxxx-grs-ipv6-64; java 1.7.0_75; Europe/en) http://yacy.net/bot.html", "yacybot (-global; amd64 Linux 3.14.32-xxxx-grs-ipv6-64; java 1.8.0_111; Europe/de) http://yacy.net/bot.html", "yacybot (/global; amd64 Linux 3.16.0-4-amd64; java 1.7.0_111; Europe/de) http://yacy.net/bot.html", "yacybot (/global; amd64 Linux 3.16.0-4-amd64; java 1.7.0_75; America/en) http://yacy.net/bot.html", "yacybot (-global; amd64 Linux 3.16.0-4-amd64; java 1.7.0_75; Europe/en) http://yacy.net/bot.html", "yacybot (/global; amd64 Linux 3.16.0-4-amd64; java 1.7.0_75; Europe/en) http://yacy.net/bot.html", "yacybot (/global; amd64 Linux 3.16.0-4-amd64; java 1.7.0_79; Europe/de) http://yacy.net/bot.html", "yacybot (/global; amd64 Linux 3.16.0-4-amd64; java 1.7.0_79; Europe/en) http://yacy.net/bot.html", "yacybot (/global; amd64 Linux 3.16.0-4-amd64; java 1.7.0_91; Europe/de) http://yacy.net/bot.html", "yacybot (/global; amd64 Linux 3.16.0-4-amd64; java 1.7.0_95; Europe/en) http://yacy.net/bot.html", "yacybot (/global; amd64 Linux 3.16.0-4-amd64; java 1.8.0_111; Europe/en) http://yacy.net/bot.html", "yacybot (/global; amd64 Linux 3.16-0.bpo.2-amd64; java 1.7.0_65; Europe/en) http://yacy.net/bot.html", "yacybot (-global; amd64 Linux 3.19.0-15-generic; java 1.8.0_45-internal; Europe/de) http://yacy.net/bot.html", "yacybot (-global; amd64 Linux 3.2.0-4-amd64; java 1.7.0_65; Europe/en) http://yacy.net/bot.html", "yacybot (-global; amd64 Linux 3.2.0-4-amd64; java 1.7.0_67; Europe/en) http://yacy.net/bot.html", "yacybot (-global; amd64 Linux 4.4.0-57-generic; java 9-internal; Europe/en) http://yacy.net/bot.html", "yacybot (-global; amd64 Windows 8.1 6.3; java 1.7.0_55; Europe/de) http://yacy.net/bot.html", "yacybot (-global; amd64 Windows 8 6.2; java 1.7.0_55; Europe/de) http://yacy.net/bot.html" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp48.MatchString(ua):
		return &Crawler{
			Pattern:      `AISearchBot`,
			Instances:    []string{  },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp49.MatchString(ua):
		return &Crawler{
			Pattern:      `ips-agent`,
			Instances:    []string{ "BlackBerry9000/4.6.0.167 Profile/MIDP-2.0 Configuration/CLDC-1.1 VendorID/102 ips-agent", "Mozilla/5.0 (X11; U; Linux i686; en-US; rv:1.7.12; ips-agent) Gecko/20050922 Fedora/1.0.7-1.1.fc4 Firefox/1.0.7", "Mozilla/5.0 (X11; U; Linux i686; en-US; rv:1.9.1.3; ips-agent) Gecko/20090824 Fedora/1.0.7-1.1.fc4  Firefox/3.5.3", "Mozilla/5.0 (X11; U; Linux i686; en-US; rv:1.9.2.24; ips-agent) Gecko/20111107 Ubuntu/10.04 (lucid) Firefox/3.6.24", "Mozilla/5.0 (X11; Ubuntu; Linux i686; rv:14.0; ips-agent) Gecko/20100101 Firefox/14.0.1" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp50.MatchString(ua):
		return &Crawler{
			Pattern:      `tagoobot`,
			Instances:    []string{  },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp51.MatchString(ua):
		return &Crawler{
			Pattern:      `MJ12bot`,
			Instances:    []string{ "MJ12bot/v1.2.0 (http://majestic12.co.uk/bot.php?+)", "Mozilla/5.0 (compatible; MJ12bot/v1.2.1; http://www.majestic12.co.uk/bot.php?+)", "Mozilla/5.0 (compatible; MJ12bot/v1.2.3; http://www.majestic12.co.uk/bot.php?+)", "Mozilla/5.0 (compatible; MJ12bot/v1.2.4; http://www.majestic12.co.uk/bot.php?+)", "Mozilla/5.0 (compatible; MJ12bot/v1.2.5; http://www.majestic12.co.uk/bot.php?+)", "Mozilla/5.0 (compatible; MJ12bot/v1.3.0; http://www.majestic12.co.uk/bot.php?+)", "Mozilla/5.0 (compatible; MJ12bot/v1.3.1; http://www.majestic12.co.uk/bot.php?+)", "Mozilla/5.0 (compatible; MJ12bot/v1.3.2; http://www.majestic12.co.uk/bot.php?+)", "Mozilla/5.0 (compatible; MJ12bot/v1.3.3; http://www.majestic12.co.uk/bot.php?+)", "Mozilla/5.0 (compatible; MJ12bot/v1.4.0; http://www.majestic12.co.uk/bot.php?+)", "Mozilla/5.0 (compatible; MJ12bot/v1.4.1; http://www.majestic12.co.uk/bot.php?+)", "Mozilla/5.0 (compatible; MJ12bot/v1.4.2; http://www.majestic12.co.uk/bot.php?+)", "Mozilla/5.0 (compatible; MJ12bot/v1.4.3; http://www.majestic12.co.uk/bot.php?+)", "Mozilla/5.0 (compatible; MJ12bot/v1.4.4 (domain ownership verifier); http://www.majestic12.co.uk/bot.php?+)", "Mozilla/5.0 (compatible; MJ12bot/v1.4.4; http://www.majestic12.co.uk/bot.php?+)", "Mozilla/5.0 (compatible; MJ12bot/v1.4.5; http://www.majestic12.co.uk/bot.php?+)", "Mozilla/5.0 (compatible; MJ12bot/v1.4.6; http://mj12bot.com/)", "Mozilla/5.0 (compatible; MJ12bot/v1.4.7; http://mj12bot.com/)", "Mozilla/5.0 (compatible; MJ12bot/v1.4.7; http://www.majestic12.co.uk/bot.php?+)", "Mozilla/5.0 (compatible; MJ12bot/v1.4.8; http://mj12bot.com/)" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp52.MatchString(ua):
		return &Crawler{
			Pattern:      `woriobot`,
			Instances:    []string{ "Mozilla/5.0 (compatible; woriobot +http://worio.com)", "Mozilla/5.0 (compatible; woriobot support [at] zite [dot] com +http://zite.com)" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp53.MatchString(ua):
		return &Crawler{
			Pattern:      `yanga`,
			Instances:    []string{ "Yanga WorldSearch Bot v1.1/beta (http://www.yanga.co.uk/)" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp54.MatchString(ua):
		return &Crawler{
			Pattern:      `buzzbot`,
			Instances:    []string{ "Buzzbot/1.0 (Buzzbot; http://www.buzzstream.com; buzzbot@buzzstream.com)" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp55.MatchString(ua):
		return &Crawler{
			Pattern:      `mlbot`,
			Instances:    []string{ "MLBot (www.metadatalabs.com/mlbot)" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp56.MatchString(ua):
		return &Crawler{
			Pattern:      `YandexBot`,
			Instances:    []string{ "Mozilla/5.0 (compatible; YandexBot/3.0; +http://yandex.com/bots)" },
			URL:          stringer("http://yandex.com/bots"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp57.MatchString(ua):
		return &Crawler{
			Pattern:      `YandexImages`,
			Instances:    []string{ "Mozilla/5.0 (compatible; YandexImages/3.0; +http://yandex.com/bots)" },
			URL:          stringer("http://yandex.com/bots"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp58.MatchString(ua):
		return &Crawler{
			Pattern:      `YandexAccessibilityBot`,
			Instances:    []string{ "Mozilla/5.0 (compatible; YandexAccessibilityBot/3.0; +http://yandex.com/bots" },
			URL:          stringer("http://yandex.com/bots"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp59.MatchString(ua):
		return &Crawler{
			Pattern:      `YandexMobileBot`,
			Instances:    []string{ "Mozilla/5.0 (iPhone; CPU iPhone OS 8_1 like Mac OS X) AppleWebKit/600.1.4 (KHTML, like Gecko) Version/8.0 Mobile/12B411 Safari/600.1.4 (compatible; YandexMobileBot/3.0; +http://yandex.com/bots)" },
			URL:          stringer("https://yandex.com/support/webmaster/robot-workings/check-yandex-robots.xml#robot-in-logs"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp60.MatchString(ua):
		return &Crawler{
			Pattern:      `purebot`,
			Instances:    []string{  },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp61.MatchString(ua):
		return &Crawler{
			Pattern:      `Linguee Bot`,
			Instances:    []string{ "Linguee Bot (http://www.linguee.com/bot)", "Linguee Bot (http://www.linguee.com/bot; bot@linguee.com)" },
			URL:          stringer("http://www.linguee.com/bot"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp62.MatchString(ua):
		return &Crawler{
			Pattern:      `CyberPatrol`,
			Instances:    []string{ "CyberPatrol SiteCat Webbot (http://www.cyberpatrol.com/cyberpatrolcrawler.asp)" },
			URL:          stringer("http://www.cyberpatrol.com/cyberpatrolcrawler.asp"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp63.MatchString(ua):
		return &Crawler{
			Pattern:      `voilabot`,
			Instances:    []string{ "Mozilla/5.0 (Windows NT 5.1; U; Win64; fr; rv:1.8.1) VoilaBot BETA 1.2 (support.voilabot@orange-ftgroup.com)", "Mozilla/5.0 (Windows; U; Windows NT 5.1; fr; rv:1.8.1) VoilaBot BETA 1.2 (support.voilabot@orange-ftgroup.com)" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp64.MatchString(ua):
		return &Crawler{
			Pattern:      `Baiduspider`,
			Instances:    []string{ "Mozilla/5.0 (compatible; Baiduspider/2.0; +http://www.baidu.com/search/spider.html)", "Mozilla/5.0 (compatible; Baiduspider-render/2.0; +http://www.baidu.com/search/spider.html)" },
			URL:          stringer("http://www.baidu.jp/spider/"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp65.MatchString(ua):
		return &Crawler{
			Pattern:      `citeseerxbot`,
			Instances:    []string{  },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp66.MatchString(ua):
		return &Crawler{
			Pattern:      `spbot`,
			Instances:    []string{ "Mozilla/5.0 (compatible; spbot/1.0; +http://www.seoprofiler.com/bot/ )", "Mozilla/5.0 (compatible; spbot/1.1; +http://www.seoprofiler.com/bot/ )", "Mozilla/5.0 (compatible; spbot/1.2; +http://www.seoprofiler.com/bot/ )", "Mozilla/5.0 (compatible; spbot/2.0.1; +http://www.seoprofiler.com/bot/ )", "Mozilla/5.0 (compatible; spbot/2.0.2; +http://www.seoprofiler.com/bot/ )", "Mozilla/5.0 (compatible; spbot/2.0.3; +http://www.seoprofiler.com/bot/ )", "Mozilla/5.0 (compatible; spbot/2.0.4; +http://www.seoprofiler.com/bot )", "Mozilla/5.0 (compatible; spbot/2.0; +http://www.seoprofiler.com/bot/ )", "Mozilla/5.0 (compatible; spbot/2.1; +http://www.seoprofiler.com/bot )", "Mozilla/5.0 (compatible; spbot/3.0; +http://www.seoprofiler.com/bot )", "Mozilla/5.0 (compatible; spbot/3.1; +http://www.seoprofiler.com/bot )", "Mozilla/5.0 (compatible; spbot/4.0.1; +http://www.seoprofiler.com/bot )", "Mozilla/5.0 (compatible; spbot/4.0.2; +http://www.seoprofiler.com/bot )", "Mozilla/5.0 (compatible; spbot/4.0.3; +http://www.seoprofiler.com/bot )", "Mozilla/5.0 (compatible; spbot/4.0.4; +http://www.seoprofiler.com/bot )", "Mozilla/5.0 (compatible; spbot/4.0.5; +http://www.seoprofiler.com/bot )", "Mozilla/5.0 (compatible; spbot/4.0.6; +http://www.seoprofiler.com/bot )", "Mozilla/5.0 (compatible; spbot/4.0.7; +http://OpenLinkProfiler.org/bot )", "Mozilla/5.0 (compatible; spbot/4.0.7; +https://www.seoprofiler.com/bot )", "Mozilla/5.0 (compatible; spbot/4.0.8; +http://OpenLinkProfiler.org/bot )", "Mozilla/5.0 (compatible; spbot/4.0.9; +http://OpenLinkProfiler.org/bot )", "Mozilla/5.0 (compatible; spbot/4.0; +http://www.seoprofiler.com/bot )", "Mozilla/5.0 (compatible; spbot/4.0a; +http://www.seoprofiler.com/bot )", "Mozilla/5.0 (compatible; spbot/4.0b; +http://www.seoprofiler.com/bot )", "Mozilla/5.0 (compatible; spbot/4.1.0; +http://OpenLinkProfiler.org/bot )", "Mozilla/5.0 (compatible; spbot/4.2.0; +http://OpenLinkProfiler.org/bot )", "Mozilla/5.0 (compatible; spbot/4.3.0; +http://OpenLinkProfiler.org/bot )", "Mozilla/5.0 (compatible; spbot/4.4.0; +http://OpenLinkProfiler.org/bot )", "Mozilla/5.0 (compatible; spbot/4.4.1; +http://OpenLinkProfiler.org/bot )", "Mozilla/5.0 (compatible; spbot/4.4.2; +http://OpenLinkProfiler.org/bot )", "Mozilla/5.0 (compatible; spbot/5.0.1; +http://OpenLinkProfiler.org/bot )", "Mozilla/5.0 (compatible; spbot/5.0.2; +http://OpenLinkProfiler.org/bot )", "Mozilla/5.0 (compatible; spbot/5.0.3; +http://OpenLinkProfiler.org/bot )", "Mozilla/5.0 (compatible; spbot/5.0; +http://OpenLinkProfiler.org/bot )" },
			URL:          stringer("http://www.seoprofiler.com/bot"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp67.MatchString(ua):
		return &Crawler{
			Pattern:      `twengabot`,
			Instances:    []string{  },
			URL:          stringer("http://www.twenga.com/bot.html"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp68.MatchString(ua):
		return &Crawler{
			Pattern:      `postrank`,
			Instances:    []string{ "PostRank/2.0 (postrank.com)", "PostRank/2.0 (postrank.com; 1 subscribers)" },
			URL:          stringer("http://www.postrank.com"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp69.MatchString(ua):
		return &Crawler{
			Pattern:      `TurnitinBot`,
			Instances:    []string{ "TurnitinBot (https://turnitin.com/robot/crawlerinfo.html)" },
			URL:          stringer("http://www.turnitin.com"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp70.MatchString(ua):
		return &Crawler{
			Pattern:      `scribdbot`,
			Instances:    []string{  },
			URL:          stringer("http://www.scribd.com"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp71.MatchString(ua):
		return &Crawler{
			Pattern:      `page2rss`,
			Instances:    []string{ "Mozilla/5.0 (compatible;  Page2RSS/0.7; +http://page2rss.com/)" },
			URL:          stringer("http://www.page2rss.com"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp72.MatchString(ua):
		return &Crawler{
			Pattern:      `sitebot`,
			Instances:    []string{ "Mozilla/5.0 (compatible; Whoiswebsitebot/0.1; +http://www.whoiswebsite.net)" },
			URL:          stringer("http://www.sitebot.org"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp73.MatchString(ua):
		return &Crawler{
			Pattern:      `linkdex`,
			Instances:    []string{ "Mozilla/5.0 (compatible; linkdexbot/2.0; +http://www.linkdex.com/about/bots/)", "Mozilla/5.0 (compatible; linkdexbot/2.0; +http://www.linkdex.com/bots/)", "Mozilla/5.0 (compatible; linkdexbot/2.1; +http://www.linkdex.com/about/bots/)", "Mozilla/5.0 (compatible; linkdexbot/2.1; +http://www.linkdex.com/bots/)", "Mozilla/5.0 (compatible; linkdexbot/2.2; +http://www.linkdex.com/bots/)", "linkdex.com/v2.0", "linkdexbot/Nutch-1.0-dev (http://www.linkdex.com/; crawl at linkdex dot com)" },
			URL:          stringer("http://www.linkdex.com"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp74.MatchString(ua):
		return &Crawler{
			Pattern:      `Adidxbot`,
			Instances:    []string{  },
			URL:          stringer("http://onlinehelp.microsoft.com/en-us/bing/hh204496.aspx"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp75.MatchString(ua):
		return &Crawler{
			Pattern:      `ezooms`,
			Instances:    []string{ "Mozilla/5.0 (compatible; Ezooms/1.0; ezooms.bot@gmail.com)" },
			URL:          stringer("http://www.phpbb.com/community/viewtopic.php?f=64&t=935605&start=450#p12948289"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp76.MatchString(ua):
		return &Crawler{
			Pattern:      `dotbot`,
			Instances:    []string{ "Mozilla/5.0 (compatible; DotBot/1.1; http://www.opensiteexplorer.org/dotbot, help@moz.com)", "dotbot" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp77.MatchString(ua):
		return &Crawler{
			Pattern:      `Mail.RU_Bot`,
			Instances:    []string{ "Mozilla/5.0 (compatible; Linux x86_64; Mail.RU_Bot/2.0; +http://go.mail.ru/help/robots)", "Mozilla/5.0 (compatible; Linux x86_64; Mail.RU_Bot/2.0; +http://go.mail.ru/", "Mozilla/5.0 (compatible; Mail.RU_Bot/2.0; +http://go.mail.ru/", "Mozilla/5.0 (compatible; Linux x86_64; Mail.RU_Bot/Robots/2.0; +http://go.mail.ru/help/robots)" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp78.MatchString(ua):
		return &Crawler{
			Pattern:      `discobot`,
			Instances:    []string{ "Mozilla/5.0 (compatible; discobot/1.0; +http://discoveryengine.com/discobot.html)", "Mozilla/5.0 (compatible; discobot/2.0; +http://discoveryengine.com/discobot.html)", "mozilla/5.0 (compatible; discobot/1.1; +http://discoveryengine.com/discobot.html)" },
			URL:          stringer("http://discoveryengine.com/discobot.html"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp79.MatchString(ua):
		return &Crawler{
			Pattern:      `heritrix`,
			Instances:    []string{ "Mozilla/5.0 (compatible; heritrix/1.12.1 +http://www.webarchiv.cz)", "Mozilla/5.0 (compatible; heritrix/1.12.1b +http://netarkivet.dk/website/info.html)", "Mozilla/5.0 (compatible; heritrix/1.14.2 +http://rjpower.org)", "Mozilla/5.0 (compatible; heritrix/1.14.2 +http://www.webarchiv.cz)", "Mozilla/5.0 (compatible; heritrix/1.14.3 +http://archive.org)", "Mozilla/5.0 (compatible; heritrix/1.14.3 +http://www.accelobot.com)", "Mozilla/5.0 (compatible; heritrix/1.14.3 +http://www.webarchiv.cz)", "Mozilla/5.0 (compatible; heritrix/1.14.3.r6601 +http://www.buddybuzz.net/yptrino)", "Mozilla/5.0 (compatible; heritrix/1.14.4 +http://parsijoo.ir)", "Mozilla/5.0 (compatible; heritrix/1.14.4 +http://www.exif-search.com)", "Mozilla/5.0 (compatible; heritrix/2.0.2 +http://aihit.com)", "Mozilla/5.0 (compatible; heritrix/2.0.2 +http://seekda.com)", "Mozilla/5.0 (compatible; heritrix/3.0.0-SNAPSHOT-20091120.021634 +http://crawler.archive.org)", "Mozilla/5.0 (compatible; heritrix/3.1.0-RC1 +http://boston.lti.cs.cmu.edu/crawler_12/)", "Mozilla/5.0 (compatible; heritrix/3.1.1 +http://places.tomtom.com/crawlerinfo)", "Mozilla/5.0 (compatible; heritrix/3.1.1 +http://www.mixdata.com)", "Mozilla/5.0 (compatible; heritrix/3.1.1; UniLeipzigASV +http://corpora.informatik.uni-leipzig.de/crawler_faq.html)", "Mozilla/5.0 (compatible; heritrix/3.2.0 +http://www.crim.ca)", "Mozilla/5.0 (compatible; heritrix/3.2.0 +http://www.exif-search.com)", "Mozilla/5.0 (compatible; heritrix/3.2.0 +http://www.mixdata.com)", "Mozilla/5.0 (compatible; heritrix/3.3.0-SNAPSHOT-20160309-0050; UniLeipzigASV +http://corpora.informatik.uni-leipzig.de/crawler_faq.html)", "Mozilla/5.0 (compatible; sukibot_heritrix/3.1.1 +http://suki.ling.helsinki.fi/eng/webmasters.html)" },
			URL:          stringer("https://github.com/internetarchive/heritrix3/wiki"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp80.MatchString(ua):
		return &Crawler{
			Pattern:      `findthatfile`,
			Instances:    []string{  },
			URL:          stringer("http://www.findthatfile.com/"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp81.MatchString(ua):
		return &Crawler{
			Pattern:      `europarchive.org`,
			Instances:    []string{ "Mozilla/5.0 (compatible; MSIE 7.0 +http://www.europarchive.org)" },
			URL:          stringer(""),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp82.MatchString(ua):
		return &Crawler{
			Pattern:      `NerdByNature.Bot`,
			Instances:    []string{ "Mozilla/5.0 (compatible; NerdByNature.Bot; http://www.nerdbynature.net/bot)" },
			URL:          stringer("http://www.nerdbynature.net/bot"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp83.MatchString(ua):
		return &Crawler{
			Pattern:      `sistrix crawler`,
			Instances:    []string{  },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp84.MatchString(ua):
		return &Crawler{
			Pattern:      `Ahrefs(Bot|SiteAudit)`,
			Instances:    []string{ "Mozilla/5.0 (compatible; AhrefsBot/6.1; +http://ahrefs.com/robot/)", "Mozilla/5.0 (compatible; AhrefsSiteAudit/6.1; +http://ahrefs.com/robot/)", "Mozilla/5.0 (compatible; AhrefsBot/5.2; News; +http://ahrefs.com/robot/)", "Mozilla/5.0 (compatible; AhrefsBot/5.2; +http://ahrefs.com/robot/)", "Mozilla/5.0 (compatible; AhrefsSiteAudit/5.2; +http://ahrefs.com/robot/)", "Mozilla/5.0 (compatible; AhrefsBot/6.1; News; +http://ahrefs.com/robot/)" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp85.MatchString(ua):
		return &Crawler{
			Pattern:      `fuelbot`,
			Instances:    []string{ "fuelbot" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp86.MatchString(ua):
		return &Crawler{
			Pattern:      `CrunchBot`,
			Instances:    []string{ "CrunchBot/1.0 (+http://www.leadcrunch.com/crunchbot)" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp87.MatchString(ua):
		return &Crawler{
			Pattern:      `IndeedBot`,
			Instances:    []string{ "Mozilla/5.0 (Windows NT 6.1; rv:38.0) Gecko/20100101 Firefox/38.0 (IndeedBot 1.1)" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp88.MatchString(ua):
		return &Crawler{
			Pattern:      `mappydata`,
			Instances:    []string{ "Mozilla/5.0 (compatible; Mappy/1.0; +http://mappydata.net/bot/)" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp89.MatchString(ua):
		return &Crawler{
			Pattern:      `woobot`,
			Instances:    []string{ "woobot" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp90.MatchString(ua):
		return &Crawler{
			Pattern:      `ZoominfoBot`,
			Instances:    []string{ "ZoominfoBot (zoominfobot at zoominfo dot com)" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp91.MatchString(ua):
		return &Crawler{
			Pattern:      `PrivacyAwareBot`,
			Instances:    []string{ "Mozilla/5.0 (compatible; PrivacyAwareBot/1.1; +http://www.privacyaware.org)" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp92.MatchString(ua):
		return &Crawler{
			Pattern:      `Multiviewbot`,
			Instances:    []string{ "Mozilla/5.0 (compatible; MSIE 9.0; Windows NT 6.1; Multiviewbot" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp93.MatchString(ua):
		return &Crawler{
			Pattern:      `SWIMGBot`,
			Instances:    []string{ "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/45.0.2454.101 Safari/537.36 SWIMGBot" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp94.MatchString(ua):
		return &Crawler{
			Pattern:      `Grobbot`,
			Instances:    []string{ "Mozilla/5.0 (compatible; Grobbot/2.2; +https://grob.it)" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp95.MatchString(ua):
		return &Crawler{
			Pattern:      `eright`,
			Instances:    []string{ "Mozilla/5.0 (compatible; eright/1.0; +bot@eright.com)" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp96.MatchString(ua):
		return &Crawler{
			Pattern:      `Apercite`,
			Instances:    []string{ "Mozilla/5.0 (compatible; Apercite; +http://www.apercite.fr/robot/index.html)" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp97.MatchString(ua):
		return &Crawler{
			Pattern:      `semanticbot`,
			Instances:    []string{ "semanticbot", "semanticbot (info@semanticaudience.com)" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp98.MatchString(ua):
		return &Crawler{
			Pattern:      `Aboundex`,
			Instances:    []string{ "Aboundex/0.2 (http://www.aboundex.com/crawler/)", "Aboundex/0.3 (http://www.aboundex.com/crawler/)" },
			URL:          stringer("http://www.aboundex.com/crawler/"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp99.MatchString(ua):
		return &Crawler{
			Pattern:      `domaincrawler`,
			Instances:    []string{ "CipaCrawler/3.0 (info@domaincrawler.com; http://www.domaincrawler.com/www.example.com)" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp100.MatchString(ua):
		return &Crawler{
			Pattern:      `wbsearchbot`,
			Instances:    []string{  },
			URL:          stringer("http://www.warebay.com/bot.html"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp101.MatchString(ua):
		return &Crawler{
			Pattern:      `summify`,
			Instances:    []string{ "Summify (Summify/1.0.1; +http://summify.com)" },
			URL:          stringer("http://summify.com"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp102.MatchString(ua):
		return &Crawler{
			Pattern:      `CCBot`,
			Instances:    []string{ "CCBot/2.0 (http://commoncrawl.org/faq/)", "CCBot/2.0 (https://commoncrawl.org/faq/)" },
			URL:          stringer("http://www.commoncrawl.org/bot.html"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp103.MatchString(ua):
		return &Crawler{
			Pattern:      `edisterbot`,
			Instances:    []string{  },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp104.MatchString(ua):
		return &Crawler{
			Pattern:      `seznambot`,
			Instances:    []string{ "Mozilla/5.0 (compatible; SeznamBot/3.2-test1-1; +http://napoveda.seznam.cz/en/seznambot-intro/)", "Mozilla/5.0 (compatible; SeznamBot/3.2-test1; +http://napoveda.seznam.cz/en/seznambot-intro/)", "Mozilla/5.0 (compatible; SeznamBot/3.2-test2; +http://napoveda.seznam.cz/en/seznambot-intro/)", "Mozilla/5.0 (compatible; SeznamBot/3.2-test4; +http://napoveda.seznam.cz/en/seznambot-intro/)", "Mozilla/5.0 (compatible; SeznamBot/3.2; +http://napoveda.seznam.cz/en/seznambot-intro/)" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp105.MatchString(ua):
		return &Crawler{
			Pattern:      `ec2linkfinder`,
			Instances:    []string{ "ec2linkfinder" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp106.MatchString(ua):
		return &Crawler{
			Pattern:      `gslfbot`,
			Instances:    []string{  },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp107.MatchString(ua):
		return &Crawler{
			Pattern:      `aiHitBot`,
			Instances:    []string{ "Mozilla/5.0 (compatible; aiHitBot/2.9; +https://www.aihitdata.com/about)" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp108.MatchString(ua):
		return &Crawler{
			Pattern:      `intelium_bot`,
			Instances:    []string{  },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp109.MatchString(ua):
		return &Crawler{
			Pattern:      `facebookexternalhit`,
			Instances:    []string{ "facebookexternalhit/1.0 (+http://www.facebook.com/externalhit_uatext.php)", "facebookexternalhit/1.1", "facebookexternalhit/1.1 (+http://www.facebook.com/externalhit_uatext.php)" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp110.MatchString(ua):
		return &Crawler{
			Pattern:      `Yeti`,
			Instances:    []string{ "Mozilla/5.0 (compatible; Yeti/1.1; +http://naver.me/bot)" },
			URL:          stringer("http://naver.me/bot"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp111.MatchString(ua):
		return &Crawler{
			Pattern:      `RetrevoPageAnalyzer`,
			Instances:    []string{ "Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 5.1; RetrevoPageAnalyzer; +http://www.retrevo.com/content/about-us)" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp112.MatchString(ua):
		return &Crawler{
			Pattern:      `lb-spider`,
			Instances:    []string{  },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp113.MatchString(ua):
		return &Crawler{
			Pattern:      `Sogou`,
			Instances:    []string{ "Sogou News Spider/4.0(+http://www.sogou.com/docs/help/webmasters.htm#07)", "Sogou Pic Spider/3.0(+http://www.sogou.com/docs/help/webmasters.htm#07)", "Sogou web spider/4.0(+http://www.sogou.com/docs/help/webmasters.htm#07)" },
			URL:          stringer("http://www.sogou.com/docs/help/webmasters.htm#07"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp114.MatchString(ua):
		return &Crawler{
			Pattern:      `lssbot`,
			Instances:    []string{  },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp115.MatchString(ua):
		return &Crawler{
			Pattern:      `careerbot`,
			Instances:    []string{  },
			URL:          stringer("http://www.career-x.de/bot.html"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp116.MatchString(ua):
		return &Crawler{
			Pattern:      `wotbox`,
			Instances:    []string{ "Wotbox/2.0 (bot@wotbox.com; http://www.wotbox.com)", "Wotbox/2.01 (+http://www.wotbox.com/bot/)" },
			URL:          stringer("http://www.wotbox.com"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp117.MatchString(ua):
		return &Crawler{
			Pattern:      `wocbot`,
			Instances:    []string{  },
			URL:          stringer("http://www.wocodi.com/crawler"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp118.MatchString(ua):
		return &Crawler{
			Pattern:      `ichiro`,
			Instances:    []string{ "DoCoMo/2.0 P900i(c100;TB;W24H11) (compatible; ichiro/mobile goo; +http://help.goo.ne.jp/help/article/1142/)", "DoCoMo/2.0 P900i(c100;TB;W24H11) (compatible; ichiro/mobile goo; +http://search.goo.ne.jp/option/use/sub4/sub4-1/)", "DoCoMo/2.0 P900i(c100;TB;W24H11) (compatible; ichiro/mobile goo;+http://search.goo.ne.jp/option/use/sub4/sub4-1/)", "DoCoMo/2.0 P900i(c100;TB;W24H11)(compatible; ichiro/mobile goo;+http://help.goo.ne.jp/door/crawler.html)", "DoCoMo/2.0 P901i(c100;TB;W24H11) (compatible; ichiro/mobile goo; +http://help.goo.ne.jp/door/crawler.html)", "KDDI-CA31 UP.Browser/6.2.0.7.3.129 (GUI) MMP/2.0 (compatible; ichiro/mobile goo; +http://help.goo.ne.jp/help/article/1142/)", "KDDI-CA31 UP.Browser/6.2.0.7.3.129 (GUI) MMP/2.0 (compatible; ichiro/mobile goo; +http://search.goo.ne.jp/option/use/sub4/sub4-1/)", "KDDI-CA31 UP.Browser/6.2.0.7.3.129 (GUI) MMP/2.0 (compatible; ichiro/mobile goo;+http://search.goo.ne.jp/option/use/sub4/sub4-1/)", "ichiro/2.0 (http://help.goo.ne.jp/door/crawler.html)", "ichiro/2.0 (ichiro@nttr.co.jp)", "ichiro/3.0 (http://help.goo.ne.jp/door/crawler.html)", "ichiro/3.0 (http://help.goo.ne.jp/help/article/1142)", "ichiro/3.0 (http://search.goo.ne.jp/option/use/sub4/sub4-1/)", "ichiro/4.0 (http://help.goo.ne.jp/door/crawler.html)", "ichiro/5.0 (http://help.goo.ne.jp/door/crawler.html)" },
			URL:          stringer("http://help.goo.ne.jp/help/article/1142"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp119.MatchString(ua):
		return &Crawler{
			Pattern:      `DuckDuckBot`,
			Instances:    []string{ "DuckDuckBot/1.0; (+http://duckduckgo.com/duckduckbot.html)", "DuckDuckBot/1.1; (+http://duckduckgo.com/duckduckbot.html)" },
			URL:          stringer("http://duckduckgo.com/duckduckbot.html"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp120.MatchString(ua):
		return &Crawler{
			Pattern:      `lssrocketcrawler`,
			Instances:    []string{  },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp121.MatchString(ua):
		return &Crawler{
			Pattern:      `drupact`,
			Instances:    []string{ "drupact/0.7; http://www.arocom.de/drupact" },
			URL:          stringer("http://www.arocom.de/drupact"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp122.MatchString(ua):
		return &Crawler{
			Pattern:      `webcompanycrawler`,
			Instances:    []string{  },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp123.MatchString(ua):
		return &Crawler{
			Pattern:      `acoonbot`,
			Instances:    []string{  },
			URL:          stringer("http://www.acoon.de/robot.asp"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp124.MatchString(ua):
		return &Crawler{
			Pattern:      `openindexspider`,
			Instances:    []string{  },
			URL:          stringer("http://www.openindex.io/en/webmasters/spider.html"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp125.MatchString(ua):
		return &Crawler{
			Pattern:      `gnam gnam spider`,
			Instances:    []string{  },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp126.MatchString(ua):
		return &Crawler{
			Pattern:      `web-archive-net.com.bot`,
			Instances:    []string{  },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp127.MatchString(ua):
		return &Crawler{
			Pattern:      `backlinkcrawler`,
			Instances:    []string{  },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp128.MatchString(ua):
		return &Crawler{
			Pattern:      `coccoc`,
			Instances:    []string{ "Mozilla/5.0 (compatible; coccoc/1.0; +http://help.coccoc.com/)", "Mozilla/5.0 (compatible; coccoc/1.0; +http://help.coccoc.com/searchengine)", "Mozilla/5.0 (compatible; coccocbot-image/1.0; +http://help.coccoc.com/searchengine)", "Mozilla/5.0 (compatible; coccocbot-web/1.0; +http://help.coccoc.com/searchengine)", "Mozilla/5.0 (compatible; image.coccoc/1.0; +http://help.coccoc.com/)", "Mozilla/5.0 (compatible; imagecoccoc/1.0; +http://help.coccoc.com/)", "Mozilla/5.0 (compatible; imagecoccoc/1.0; +http://help.coccoc.com/searchengine)", "coccoc", "coccoc/1.0 ()", "coccoc/1.0 (http://help.coccoc.com/)", "coccoc/1.0 (http://help.coccoc.vn/)" },
			URL:          stringer("http://help.coccoc.vn/"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp129.MatchString(ua):
		return &Crawler{
			Pattern:      `integromedb`,
			Instances:    []string{ "www.integromedb.org/Crawler" },
			URL:          stringer("http://www.integromedb.org/Crawler"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp130.MatchString(ua):
		return &Crawler{
			Pattern:      `content crawler spider`,
			Instances:    []string{  },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp131.MatchString(ua):
		return &Crawler{
			Pattern:      `toplistbot`,
			Instances:    []string{  },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp132.MatchString(ua):
		return &Crawler{
			Pattern:      `it2media-domain-crawler`,
			Instances:    []string{ "it2media-domain-crawler/1.0 on crawler-prod.it2media.de", "it2media-domain-crawler/2.0" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp133.MatchString(ua):
		return &Crawler{
			Pattern:      `ip-web-crawler.com`,
			Instances:    []string{  },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp134.MatchString(ua):
		return &Crawler{
			Pattern:      `siteexplorer.info`,
			Instances:    []string{ "Mozilla/5.0 (compatible; SiteExplorer/1.0b; +http://siteexplorer.info/)", "Mozilla/5.0 (compatible; SiteExplorer/1.1b; +http://siteexplorer.info/Backlink-Checker-Spider/)" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp135.MatchString(ua):
		return &Crawler{
			Pattern:      `elisabot`,
			Instances:    []string{  },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp136.MatchString(ua):
		return &Crawler{
			Pattern:      `proximic`,
			Instances:    []string{ "Mozilla/5.0 (compatible; proximic; +http://www.proximic.com)", "Mozilla/5.0 (compatible; proximic; +http://www.proximic.com/info/spider.php)" },
			URL:          stringer("http://www.proximic.com/info/spider.php"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp137.MatchString(ua):
		return &Crawler{
			Pattern:      `changedetection`,
			Instances:    []string{ "Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1; SV1;  http://www.changedetection.com/bot.html )" },
			URL:          stringer("http://www.changedetection.com/bot.html"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp138.MatchString(ua):
		return &Crawler{
			Pattern:      `arabot`,
			Instances:    []string{  },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp139.MatchString(ua):
		return &Crawler{
			Pattern:      `WeSEE:Search`,
			Instances:    []string{ "WeSEE:Search", "WeSEE:Search/0.1 (Alpha, http://www.wesee.com/en/support/bot/)" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp140.MatchString(ua):
		return &Crawler{
			Pattern:      `niki-bot`,
			Instances:    []string{  },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp141.MatchString(ua):
		return &Crawler{
			Pattern:      `CrystalSemanticsBot`,
			Instances:    []string{  },
			URL:          stringer("http://www.crystalsemantics.com/user-agent/"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp142.MatchString(ua):
		return &Crawler{
			Pattern:      `rogerbot`,
			Instances:    []string{ "Mozilla/5.0 (compatible; rogerBot/1.0; UrlCrawler; http://www.seomoz.org/dp/rogerbot)", "rogerbot/1.0 (http://moz.com/help/pro/what-is-rogerbot-, rogerbot-crawler+partager@moz.com)", "rogerbot/1.0 (http://moz.com/help/pro/what-is-rogerbot-, rogerbot-crawler+shiny@moz.com)", "rogerbot/1.0 (http://moz.com/help/pro/what-is-rogerbot-, rogerbot-wherecat@moz.com", "rogerbot/1.0 (http://moz.com/help/pro/what-is-rogerbot-, rogerbot-wherecat@moz.com)", "rogerbot/1.0 (http://www.moz.com/dp/rogerbot, rogerbot-crawler@moz.com)", "rogerbot/1.0 (http://www.seomoz.org/dp/rogerbot, rogerbot-crawler+shiny@seomoz.org)", "rogerbot/1.0 (http://www.seomoz.org/dp/rogerbot, rogerbot-crawler@seomoz.org)", "rogerbot/1.0 (http://www.seomoz.org/dp/rogerbot, rogerbot-wherecat@moz.com)", "rogerbot/1.1 (http://moz.com/help/guides/search-overview/crawl-diagnostics#more-help, rogerbot-crawler+pr2-crawler-05@moz.com)", "rogerbot/1.1 (http://moz.com/help/guides/search-overview/crawl-diagnostics#more-help, rogerbot-crawler+pr4-crawler-11@moz.com)", "rogerbot/1.1 (http://moz.com/help/guides/search-overview/crawl-diagnostics#more-help, rogerbot-crawler+pr4-crawler-15@moz.com)", "rogerbot/1.2 (http://moz.com/help/pro/what-is-rogerbot-, rogerbot-crawler+phaser-testing-crawler-01@moz.com)" },
			URL:          stringer("http://moz.com/help/pro/what-is-rogerbot-"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp143.MatchString(ua):
		return &Crawler{
			Pattern:      `360Spider`,
			Instances:    []string{ "Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.1 (KHTML, like Gecko) Chrome/21.0.1180.89 Safari/537.1; 360Spider", "Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.1 (KHTML, like Gecko) Chrome/21.0.1180.89 Safari/537.1; 360Spider(compatible; HaosouSpider; http://www.haosou.com/help/help_3_2.html)", "Mozilla/5.0 (Windows NT 6.2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/31.0.1650.63 Safari/537.36 QIHU 360SE; 360Spider", "Mozilla/5.0 (Windows; U; Windows NT 5.1; zh-CN; )  Firefox/1.5.0.11; 360Spider", "Mozilla/5.0 (Windows; U; Windows NT 5.1; zh-CN; rv:1.8.0.11)  Firefox/1.5.0.11; 360Spider", "Mozilla/5.0 (Windows; U; Windows NT 5.1; zh-CN; rv:1.8.0.11) Firefox/1.5.0.11 360Spider;", "Mozilla/5.0 (Windows; U; Windows NT 5.1; zh-CN; rv:1.8.0.11) Gecko/20070312 Firefox/1.5.0.11; 360Spider", "Mozilla/5.0 (compatible; MSIE 9.0; Windows NT 6.1; Trident/5.0); 360Spider", "Mozilla/5.0 (compatible; MSIE 9.0; Windows NT 6.1; Trident/5.0); 360Spider(compatible; HaosouSpider; http://www.haosou.com/help/help_3_2.html)" },
			URL:          stringer("http://needs-be.blogspot.co.uk/2013/02/how-to-block-spider360.html"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp144.MatchString(ua):
		return &Crawler{
			Pattern:      `psbot`,
			Instances:    []string{ "psbot-image (+http://www.picsearch.com/bot.html)", "psbot-page (+http://www.picsearch.com/bot.html)", "psbot/0.1 (+http://www.picsearch.com/bot.html)" },
			URL:          stringer("http://www.picsearch.com/bot.html"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp145.MatchString(ua):
		return &Crawler{
			Pattern:      `InterfaxScanBot`,
			Instances:    []string{  },
			URL:          stringer("http://scan-interfax.ru"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp146.MatchString(ua):
		return &Crawler{
			Pattern:      `CC Metadata Scaper`,
			Instances:    []string{ "CC Metadata Scaper http://wiki.creativecommons.org/Metadata_Scraper" },
			URL:          stringer("http://wiki.creativecommons.org/Metadata_Scraper"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp147.MatchString(ua):
		return &Crawler{
			Pattern:      `g00g1e.net`,
			Instances:    []string{  },
			URL:          stringer("http://www.g00g1e.net/"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp148.MatchString(ua):
		return &Crawler{
			Pattern:      `GrapeshotCrawler`,
			Instances:    []string{ "Mozilla/5.0 (compatible; GrapeshotCrawler/2.0; +http://www.grapeshot.co.uk/crawler.php)" },
			URL:          stringer("http://www.grapeshot.co.uk/crawler.php"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp149.MatchString(ua):
		return &Crawler{
			Pattern:      `urlappendbot`,
			Instances:    []string{ "Mozilla/5.0 (compatible; URLAppendBot/1.0; +http://www.profound.net/urlappendbot.html)" },
			URL:          stringer("http://www.profound.net/urlappendbot.html"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp150.MatchString(ua):
		return &Crawler{
			Pattern:      `brainobot`,
			Instances:    []string{  },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp151.MatchString(ua):
		return &Crawler{
			Pattern:      `fr-crawler`,
			Instances:    []string{ "Mozilla/5.0 (compatible; fr-crawler/1.1)" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp152.MatchString(ua):
		return &Crawler{
			Pattern:      `binlar`,
			Instances:    []string{ "binlar_2.6.3 binlar2.6.3@unspecified.mail", "binlar_2.6.3 binlar_2.6.3@unspecified.mail", "binlar_2.6.3 larbin2.6.3@unspecified.mail", "binlar_2.6.3 phanendra_kalapala@McAfee.com", "binlar_2.6.3 test@mgmt.mic" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp153.MatchString(ua):
		return &Crawler{
			Pattern:      `SimpleCrawler`,
			Instances:    []string{ "SimpleCrawler/0.1" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp154.MatchString(ua):
		return &Crawler{
			Pattern:      `Twitterbot`,
			Instances:    []string{ "Twitterbot/0.1", "Twitterbot/1.0" },
			URL:          stringer("https://dev.twitter.com/cards/getting-started"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp155.MatchString(ua):
		return &Crawler{
			Pattern:      `cXensebot`,
			Instances:    []string{ "cXensebot/1.1a" },
			URL:          stringer("http://www.cxense.com/bot.html"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp156.MatchString(ua):
		return &Crawler{
			Pattern:      `smtbot`,
			Instances:    []string{ "Mozilla/5.0 (compatible; SMTBot/1.0; +http://www.similartech.com/smtbot)", "SMTBot (similartech.com/smtbot)", "Mozilla/5.0 (iPhone; CPU iPhone OS 6_0 like Mac OS X) AppleWebKit/536.26 (KHTML, like Gecko)                 Version/6.0 Mobile/10A5376e Safari/8536.25 (compatible; SMTBot/1.0; +http://www.similartech.com/smtbot)", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.75 Safari/537.36 (compatible; SMTBot/1.0; +http://www.similartech.com/smtbot)" },
			URL:          stringer("http://www.similartech.com/smtbot"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp157.MatchString(ua):
		return &Crawler{
			Pattern:      `bnf.fr_bot`,
			Instances:    []string{ "Mozilla/5.0 (compatible; bnf.fr_bot; +http://bibnum.bnf.fr/robot/bnf.html)", "Mozilla/5.0 (compatible; bnf.fr_bot; +http://www.bnf.fr/fr/outils/a.dl_web_capture_robot.html)" },
			URL:          stringer("http://www.bnf.fr/fr/outils/a.dl_web_capture_robot.html"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp158.MatchString(ua):
		return &Crawler{
			Pattern:      `A6-Indexer`,
			Instances:    []string{ "A6-Indexer" },
			URL:          stringer("http://www.a6corp.com/a6-web-scraping-policy/"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp159.MatchString(ua):
		return &Crawler{
			Pattern:      `ADmantX`,
			Instances:    []string{ "ADmantX Platform Semantic Analyzer - ADmantX Inc. - www.admantx.com - support@admantx.com" },
			URL:          stringer("http://www.admantx.com"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp160.MatchString(ua):
		return &Crawler{
			Pattern:      `Facebot`,
			Instances:    []string{ "Facebot/1.0" },
			URL:          stringer("https://developers.facebook.com/docs/sharing/best-practices#crawl"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp161.MatchString(ua):
		return &Crawler{
			Pattern:      `OrangeBot\/`,
			Instances:    []string{ "Mozilla/5.0 (compatible; OrangeBot/2.0; support.orangebot@orange.com" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp162.MatchString(ua):
		return &Crawler{
			Pattern:      `memorybot`,
			Instances:    []string{ "Mozilla/5.0 (compatible; memorybot/1.21.14 +http://mignify.com/bot.html)" },
			URL:          stringer("http://mignify.com/bot.htm"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp163.MatchString(ua):
		return &Crawler{
			Pattern:      `AdvBot`,
			Instances:    []string{ "Mozilla/5.0 (compatible; AdvBot/2.0; +http://advbot.net/bot.html)" },
			URL:          stringer("http://advbot.net/bot.html"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp164.MatchString(ua):
		return &Crawler{
			Pattern:      `MegaIndex`,
			Instances:    []string{ "Mozilla/5.0 (compatible; MegaIndex.ru/2.0; +https://www.megaindex.ru/?tab=linkAnalyze)" },
			URL:          stringer("https://www.megaindex.ru/?tab=linkAnalyze"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp165.MatchString(ua):
		return &Crawler{
			Pattern:      `SemanticScholarBot`,
			Instances:    []string{ "SemanticScholarBot/1.0 (+http://s2.allenai.org/bot.html)", "Mozilla/5.0 (compatible) SemanticScholarBot (+https://www.semanticscholar.org/crawler)" },
			URL:          stringer("https://www.semanticscholar.org/crawler"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp166.MatchString(ua):
		return &Crawler{
			Pattern:      `ltx71`,
			Instances:    []string{ "ltx71 - (http://ltx71.com/)" },
			URL:          stringer("http://ltx71.com/"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp167.MatchString(ua):
		return &Crawler{
			Pattern:      `nerdybot`,
			Instances:    []string{ "nerdybot" },
			URL:          stringer("http://nerdybot.com/"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp168.MatchString(ua):
		return &Crawler{
			Pattern:      `xovibot`,
			Instances:    []string{ "Mozilla/5.0 (compatible; XoviBot/2.0; +http://www.xovibot.net/)" },
			URL:          stringer("http://www.xovibot.net/"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp169.MatchString(ua):
		return &Crawler{
			Pattern:      `BUbiNG`,
			Instances:    []string{ "BUbiNG (+http://law.di.unimi.it/BUbiNG.html)" },
			URL:          stringer("http://law.di.unimi.it/BUbiNG.html"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp170.MatchString(ua):
		return &Crawler{
			Pattern:      `Qwantify`,
			Instances:    []string{ "Mozilla/5.0 (compatible; Qwantify/2.0n; +https://www.qwant.com/)/*", "Mozilla/5.0 (compatible; Qwantify/2.4w; +https://www.qwant.com/)/2.4w" },
			URL:          stringer("https://www.qwant.com/"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp171.MatchString(ua):
		return &Crawler{
			Pattern:      `archive.org_bot`,
			Instances:    []string{ "Mozilla/5.0 (compatible; heritrix/3.1.1-SNAPSHOT-20120116.200628 +http://www.archive.org/details/archive.org_bot)", "Mozilla/5.0 (compatible; archive.org_bot/heritrix-1.15.4 +http://www.archive.org)", "Mozilla/5.0 (compatible; heritrix/3.3.0-SNAPSHOT-20140702-2247 +http://archive.org/details/archive.org_bot)", "Mozilla/5.0 (compatible; archive.org_bot +http://www.archive.org/details/archive.org_bot)" },
			URL:          stringer("http://www.archive.org/details/archive.org_bot"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{ "heritrix" },
		}

	case rgxp172.MatchString(ua):
		return &Crawler{
			Pattern:      `Applebot`,
			Instances:    []string{ "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_1) AppleWebKit/600.2.5 (KHTML, like Gecko) Version/8.0.2 Safari/600.2.5 (Applebot/0.1)", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_1) AppleWebKit/600.2.5 (KHTML, like Gecko) Version/8.0.2 Safari/600.2.5 (Applebot/0.1; +http://www.apple.com/go/applebot)", "Mozilla/5.0 (compatible; Applebot/0.3; +http://www.apple.com/go/applebot)", "Mozilla/5.0 (iPhone; CPU iPhone OS 6_0 like Mac OS X) AppleWebKit/536.26 (KHTML, like Gecko) Version/6.0 Mobile/10A5376e Safari/8536.25 (compatible; Applebot/0.3; +http://www.apple.com/go/applebot)", "Mozilla/5.0 (iPhone; CPU iPhone OS 8_1 like Mac OS X) AppleWebKit/600.1.4 (KHTML, like Gecko) Version/8.0 Mobile/12B410 Safari/600.1.4 (Applebot/0.1; +http://www.apple.com/go/applebot)" },
			URL:          stringer("http://www.apple.com/go/applebot"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp173.MatchString(ua):
		return &Crawler{
			Pattern:      `TweetmemeBot`,
			Instances:    []string{ "Mozilla/5.0 (TweetmemeBot/4.0; +http://datasift.com/bot.html) Gecko/20100101 Firefox/31.0" },
			URL:          stringer("http://datasift.com/bot.html"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp174.MatchString(ua):
		return &Crawler{
			Pattern:      `crawler4j`,
			Instances:    []string{ "crawler4j (http://code.google.com/p/crawler4j/)" },
			URL:          stringer("https://github.com/yasserg/crawler4j"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp175.MatchString(ua):
		return &Crawler{
			Pattern:      `findxbot`,
			Instances:    []string{ "Mozilla/5.0 (compatible; Findxbot/1.0; +http://www.findxbot.com)" },
			URL:          stringer("http://www.findxbot.com"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp176.MatchString(ua):
		return &Crawler{
			Pattern:      `S[eE][mM]rushBot`,
			Instances:    []string{ "Mozilla/5.0 (compatible; SemrushBot-SA/0.97; +http://www.semrush.com/bot.html)", "Mozilla/5.0 (compatible; SemrushBot-SI/0.97; +http://www.semrush.com/bot.html)", "Mozilla/5.0 (compatible; SemrushBot/3~bl; +http://www.semrush.com/bot.html)", "Mozilla/5.0 (compatible; SemrushBot/0.98~bl; +http://www.semrush.com/bot.html)", "Mozilla/5.0 (compatible; SemrushBot-BA; +http://www.semrush.com/bot.html)", "SEMrushBot" },
			URL:          stringer("http://www.semrush.com/bot.html"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp177.MatchString(ua):
		return &Crawler{
			Pattern:      `yoozBot`,
			Instances:    []string{ "Mozilla/5.0 (compatible; yoozBot-2.2; http://yooz.ir; info@yooz.ir)" },
			URL:          stringer("http://yooz.ir"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp178.MatchString(ua):
		return &Crawler{
			Pattern:      `lipperhey`,
			Instances:    []string{ "Mozilla/5.0 (compatible; Lipperhey Link Explorer; http://www.lipperhey.com/)", "Mozilla/5.0 (compatible; Lipperhey SEO Service; http://www.lipperhey.com/)", "Mozilla/5.0 (compatible; Lipperhey Site Explorer; http://www.lipperhey.com/)", "Mozilla/5.0 (compatible; Lipperhey-Kaus-Australis/5.0; +https://www.lipperhey.com/en/about/)" },
			URL:          stringer("http://www.lipperhey.com/"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp179.MatchString(ua):
		return &Crawler{
			Pattern:      `Y!J`,
			Instances:    []string{ "Y!J-ASR/0.1 crawler (http://www.yahoo-help.jp/app/answers/detail/p/595/a_id/42716/)", "Y!J-BRJ/YATS crawler (http://help.yahoo.co.jp/help/jp/search/indexing/indexing-15.html)", "Y!J-PSC/1.0 crawler (http://help.yahoo.co.jp/help/jp/search/indexing/indexing-15.html)", "Y!J-BRW/1.0 crawler (http://help.yahoo.co.jp/help/jp/search/indexing/indexing-15.html)", "Mozilla/5.0 (iPhone; Y!J-BRY/YATSH crawler; http://help.yahoo.co.jp/help/jp/search/indexing/indexing-15.html)", "Mozilla/5.0 (compatible; Y!J SearchMonkey/1.0 (Y!J-AGENT; http://help.yahoo.co.jp/help/jp/search/indexing/indexing-15.html))" },
			URL:          stringer("https://www.yahoo-help.jp/app/answers/detail/p/595/a_id/42716/~/%E3%82%A6%E3%82%A7%E3%83%96%E3%83%9A%E3%83%BC%E3%82%B8%E3%81%AB%E3%82%A2%E3%82%AF%E3%82%BB%E3%82%B9%E3%81%99%E3%82%8B%E3%82%B7%E3%82%B9%E3%83%86%E3%83%A0%E3%81%AE%E3%83%A6%E3%83%BC%E3%82%B6%E3%83%BC%E3%82%A8%E3%83%BC%E3%82%B8%E3%82%A7%E3%83%B3%E3%83%88%E3%81%AB%E3%81%A4%E3%81%84%E3%81%A6"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp180.MatchString(ua):
		return &Crawler{
			Pattern:      `Domain Re-Animator Bot`,
			Instances:    []string{ "Domain Re-Animator Bot (http://domainreanimator.com) - support@domainreanimator.com" },
			URL:          stringer("http://domainreanimator.com"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp181.MatchString(ua):
		return &Crawler{
			Pattern:      `AddThis`,
			Instances:    []string{ "AddThis.com robot tech.support@clearspring.com" },
			URL:          stringer("https://www.addthis.com"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp182.MatchString(ua):
		return &Crawler{
			Pattern:      `Screaming Frog SEO Spider`,
			Instances:    []string{ "Screaming Frog SEO Spider/5.1" },
			URL:          stringer("http://www.screamingfrog.co.uk/seo-spider"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp183.MatchString(ua):
		return &Crawler{
			Pattern:      `MetaURI`,
			Instances:    []string{ "MetaURI API/2.0 +metauri.com" },
			URL:          stringer("http://www.useragentstring.com/MetaURI_id_17683.php"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp184.MatchString(ua):
		return &Crawler{
			Pattern:      `Scrapy`,
			Instances:    []string{ "Scrapy/1.0.3 (+http://scrapy.org)" },
			URL:          stringer("http://scrapy.org/"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp185.MatchString(ua):
		return &Crawler{
			Pattern:      `Livelap[bB]ot`,
			Instances:    []string{ "LivelapBot/0.2 (http://site.livelap.com/crawler)", "Livelapbot/0.1" },
			URL:          stringer("http://site.livelap.com/crawler"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp186.MatchString(ua):
		return &Crawler{
			Pattern:      `OpenHoseBot`,
			Instances:    []string{ "Mozilla/5.0 (compatible; OpenHoseBot/2.1; +http://www.openhose.org/bot.html)" },
			URL:          stringer("http://www.openhose.org/bot.html"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp187.MatchString(ua):
		return &Crawler{
			Pattern:      `CapsuleChecker`,
			Instances:    []string{ "CapsuleChecker (http://www.capsulink.com/)" },
			URL:          stringer("http://www.capsulink.com/about"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp188.MatchString(ua):
		return &Crawler{
			Pattern:      `collection@infegy.com`,
			Instances:    []string{ "Mozilla/5.0 (compatible) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/47.0.2526.73 Safari/537.36 collection@infegy.com" },
			URL:          stringer("http://infegy.com/"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp189.MatchString(ua):
		return &Crawler{
			Pattern:      `IstellaBot`,
			Instances:    []string{ "Mozilla/5.0 (compatible; IstellaBot/1.23.15 +http://www.tiscali.it/)" },
			URL:          stringer("http://www.tiscali.it/"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp190.MatchString(ua):
		return &Crawler{
			Pattern:      `DeuSu\/`,
			Instances:    []string{ "Mozilla/5.0 (compatible; DeuSu/0.1.0; +https://deusu.org)", "Mozilla/5.0 (compatible; DeuSu/5.0.2; +https://deusu.de/robot.html)" },
			URL:          stringer("https://deusu.de/robot.html"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp191.MatchString(ua):
		return &Crawler{
			Pattern:      `betaBot`,
			Instances:    []string{  },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp192.MatchString(ua):
		return &Crawler{
			Pattern:      `Cliqzbot\/`,
			Instances:    []string{ "Mozilla/5.0 (compatible; Cliqzbot/2.0; +http://cliqz.com/company/cliqzbot)", "Cliqzbot/0.1 (+http://cliqz.com +cliqzbot@cliqz.com)", "Cliqzbot/0.1 (+http://cliqz.com/company/cliqzbot)", "Mozilla/5.0 (compatible; Cliqzbot/0.1 +http://cliqz.com/company/cliqzbot)", "Mozilla/5.0 (compatible; Cliqzbot/1.0 +http://cliqz.com/company/cliqzbot)" },
			URL:          stringer("http://cliqz.com/company/cliqzbot"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp193.MatchString(ua):
		return &Crawler{
			Pattern:      `MojeekBot\/`,
			Instances:    []string{ "MojeekBot/0.2 (archi; http://www.mojeek.com/bot.html)", "Mozilla/5.0 (compatible; MojeekBot/0.2; http://www.mojeek.com/bot.html#relaunch)", "Mozilla/5.0 (compatible; MojeekBot/0.2; http://www.mojeek.com/bot.html)", "Mozilla/5.0 (compatible; MojeekBot/0.5; http://www.mojeek.com/bot.html)", "Mozilla/5.0 (compatible; MojeekBot/0.6; +https://www.mojeek.com/bot.html)", "Mozilla/5.0 (compatible; MojeekBot/0.6; http://www.mojeek.com/bot.html)" },
			URL:          stringer("https://www.mojeek.com/bot.html"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp194.MatchString(ua):
		return &Crawler{
			Pattern:      `netEstate NE Crawler`,
			Instances:    []string{ "netEstate NE Crawler (+http://www.sengine.info/)", "netEstate NE Crawler (+http://www.website-datenbank.de/)" },
			URL:          stringer("+http://www.website-datenbank.de/"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp195.MatchString(ua):
		return &Crawler{
			Pattern:      `SafeSearch microdata crawler`,
			Instances:    []string{ "SafeSearch microdata crawler (https://safesearch.avira.com, safesearch-abuse@avira.com)" },
			URL:          stringer("https://safesearch.avira.com"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp196.MatchString(ua):
		return &Crawler{
			Pattern:      `Gluten Free Crawler\/`,
			Instances:    []string{ "Mozilla/5.0 (compatible; Gluten Free Crawler/1.0; +http://glutenfreepleasure.com/)" },
			URL:          stringer("http://glutenfreepleasure.com/"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp197.MatchString(ua):
		return &Crawler{
			Pattern:      `Sonic`,
			Instances:    []string{ "Mozilla/5.0 (compatible; RankSonicSiteAuditor/1.0; +https://ranksonic.com/ranksonic_sab.html)", "Mozilla/5.0 (compatible; Sonic/1.0; http://www.yama.info.waseda.ac.jp/~crawler/info.html)", "Mozzila/5.0 (compatible; Sonic/1.0; http://www.yama.info.waseda.ac.jp/~crawler/info.html)" },
			URL:          stringer("http://www.yama.info.waseda.ac.jp/~crawler/info.html"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp198.MatchString(ua):
		return &Crawler{
			Pattern:      `Sysomos`,
			Instances:    []string{ "Mozilla/5.0 (compatible; Sysomos/1.0; +http://www.sysomos.com/; Sysomos)" },
			URL:          stringer("http://www.sysomos.com"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp199.MatchString(ua):
		return &Crawler{
			Pattern:      `Trove`,
			Instances:    []string{  },
			URL:          stringer("http://www.trove.com"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp200.MatchString(ua):
		return &Crawler{
			Pattern:      `deadlinkchecker`,
			Instances:    []string{ "www.deadlinkchecker.com Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/46.0.2490.86 Safari/537.36", "www.deadlinkchecker.com XMLHTTP/1.0", "www.deadlinkchecker.com XMLHTTP/1.0 Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/46.0.2490.86 Safari/537.36" },
			URL:          stringer("http://www.deadlinkchecker.com"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp201.MatchString(ua):
		return &Crawler{
			Pattern:      `Slack-ImgProxy`,
			Instances:    []string{ "Slack-ImgProxy (+https://api.slack.com/robots)", "Slack-ImgProxy 0.59 (+https://api.slack.com/robots)", "Slack-ImgProxy 0.66 (+https://api.slack.com/robots)", "Slack-ImgProxy 1.106 (+https://api.slack.com/robots)", "Slack-ImgProxy 1.138 (+https://api.slack.com/robots)", "Slack-ImgProxy 149 (+https://api.slack.com/robots)" },
			URL:          stringer("https://api.slack.com/robots"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp202.MatchString(ua):
		return &Crawler{
			Pattern:      `Embedly`,
			Instances:    []string{ "Embedly +support@embed.ly", "Mozilla/5.0 (compatible; Embedly/0.2; +http://support.embed.ly/)", "Mozilla/5.0 (compatible; Embedly/0.2; snap; +http://support.embed.ly/)" },
			URL:          stringer("http://support.embed.ly"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp203.MatchString(ua):
		return &Crawler{
			Pattern:      `RankActiveLinkBot`,
			Instances:    []string{ "Mozilla/5.0 (compatible; RankActiveLinkBot; +https://rankactive.com/resources/rankactive-linkbot)" },
			URL:          stringer("https://rankactive.com/resources/rankactive-linkbot"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp204.MatchString(ua):
		return &Crawler{
			Pattern:      `iskanie`,
			Instances:    []string{ "iskanie (+http://www.iskanie.com)" },
			URL:          stringer("http://www.iskanie.com"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp205.MatchString(ua):
		return &Crawler{
			Pattern:      `SafeDNSBot`,
			Instances:    []string{ "SafeDNSBot (https://www.safedns.com/searchbot)" },
			URL:          stringer("https://www.safedns.com/searchbot"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp206.MatchString(ua):
		return &Crawler{
			Pattern:      `SkypeUriPreview`,
			Instances:    []string{ "Mozilla/5.0 (Windows NT 6.1; WOW64) SkypeUriPreview Preview/0.5" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp207.MatchString(ua):
		return &Crawler{
			Pattern:      `Veoozbot`,
			Instances:    []string{ "Mozilla/5.0 (compatible; Veoozbot/1.0; +http://www.veooz.com/veoozbot.html)" },
			URL:          stringer("http://www.veooz.com/veoozbot.html"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp208.MatchString(ua):
		return &Crawler{
			Pattern:      `Slackbot`,
			Instances:    []string{ "Slackbot-LinkExpanding (+https://api.slack.com/robots)", "Slackbot-LinkExpanding 1.0 (+https://api.slack.com/robots)" },
			URL:          stringer("https://api.slack.com/robots"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp209.MatchString(ua):
		return &Crawler{
			Pattern:      `redditbot`,
			Instances:    []string{ "Mozilla/5.0 (compatible; redditbot/1.0; +http://www.reddit.com/feedback)" },
			URL:          stringer("http://www.reddit.com/feedback"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp210.MatchString(ua):
		return &Crawler{
			Pattern:      `datagnionbot`,
			Instances:    []string{ "datagnionbot (+http://www.datagnion.com/bot.html)" },
			URL:          stringer("http://www.datagnion.com/bot.html"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp211.MatchString(ua):
		return &Crawler{
			Pattern:      `Google-Adwords-Instant`,
			Instances:    []string{ "Google-Adwords-Instant (+http://www.google.com/adsbot.html)" },
			URL:          stringer("http://www.google.com/adsbot.html"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp212.MatchString(ua):
		return &Crawler{
			Pattern:      `adbeat_bot`,
			Instances:    []string{ "Mozilla/5.0 (compatible; adbeat_bot; +support@adbeat.com; support@adbeat.com)", "adbeat_bot" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp213.MatchString(ua):
		return &Crawler{
			Pattern:      `WhatsApp`,
			Instances:    []string{ "WhatsApp", "WhatsApp/2.12.15/i", "WhatsApp/2.12.16/i", "WhatsApp/2.12.17/i", "WhatsApp/2.12.449 A", "WhatsApp/2.12.453 A", "WhatsApp/2.12.510 A", "WhatsApp/2.12.540 A", "WhatsApp/2.12.548 A", "WhatsApp/2.12.555 A", "WhatsApp/2.12.556 A", "WhatsApp/2.16.1/i", "WhatsApp/2.16.13 A", "WhatsApp/2.16.2/i", "WhatsApp/2.16.42 A", "WhatsApp/2.16.57 A" },
			URL:          stringer("https://www.whatsapp.com/"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp214.MatchString(ua):
		return &Crawler{
			Pattern:      `contxbot`,
			Instances:    []string{ "Mozilla/5.0 (compatible;contxbot/1.0)" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp215.MatchString(ua):
		return &Crawler{
			Pattern:      `pinterest.com.bot`,
			Instances:    []string{ "Mozilla/5.0 (compatible; Pinterestbot/1.0; +http://www.pinterest.com/bot.html)", "Pinterest/0.2 (+http://www.pinterest.com/bot.html)" },
			URL:          stringer("http://www.pinterest.com/bot.html"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp216.MatchString(ua):
		return &Crawler{
			Pattern:      `electricmonk`,
			Instances:    []string{ "Mozilla/5.0 (compatible; electricmonk/3.2.0 +https://www.duedil.com/our-crawler/)" },
			URL:          stringer("https://www.duedil.com/our-crawler/"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp217.MatchString(ua):
		return &Crawler{
			Pattern:      `GarlikCrawler`,
			Instances:    []string{ "GarlikCrawler/1.2 (http://garlik.com/, crawler@garlik.com)" },
			URL:          stringer("http://garlik.com/"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp218.MatchString(ua):
		return &Crawler{
			Pattern:      `BingPreview\/`,
			Instances:    []string{ "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/534+ (KHTML, like Gecko) BingPreview/1.0b", "Mozilla/5.0 (Windows NT 6.3; WOW64; Trident/7.0; rv:11.0; BingPreview/1.0b) like Gecko", "Mozilla/5.0 (compatible; MSIE 10.0; Windows NT 6.2; Trident/6.0;  WOW64;  Trident/6.0;  BingPreview/1.0b)", "Mozilla/5.0 (compatible; MSIE 9.0; Windows NT 6.1; Trident/5.0;  WOW64;  Trident/5.0;  BingPreview/1.0b)", "Mozilla/5.0 (iPhone; CPU iPhone OS 7_0 like Mac OS X) AppleWebKit/537.51.1 (KHTML, like Gecko) Version/7.0 Mobile/11A465 Safari/9537.53 BingPreview/1.0b" },
			URL:          stringer("https://www.bing.com/webmaster/help/which-crawlers-does-bing-use-8c184ec0"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp219.MatchString(ua):
		return &Crawler{
			Pattern:      `vebidoobot`,
			Instances:    []string{ "Mozilla/5.0 (compatible; vebidoobot/1.0; +https://blog.vebidoo.de/vebidoobot/" },
			URL:          stringer("https://blog.vebidoo.de/vebidoobot/"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp220.MatchString(ua):
		return &Crawler{
			Pattern:      `FemtosearchBot`,
			Instances:    []string{ "Mozilla/5.0 (compatible; FemtosearchBot/1.0; http://femtosearch.com)" },
			URL:          stringer("http://femtosearch.com"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp221.MatchString(ua):
		return &Crawler{
			Pattern:      `Yahoo Link Preview`,
			Instances:    []string{ "Mozilla/5.0 (compatible; Yahoo Link Preview; https://help.yahoo.com/kb/mail/yahoo-link-preview-SLN23615.html)" },
			URL:          stringer("https://help.yahoo.com/kb/mail/yahoo-link-preview-SLN23615.html"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp222.MatchString(ua):
		return &Crawler{
			Pattern:      `MetaJobBot`,
			Instances:    []string{ "Mozilla/5.0 (compatible; MetaJobBot; http://www.metajob.de/crawler)" },
			URL:          stringer("http://www.metajob.de/the/crawler"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp223.MatchString(ua):
		return &Crawler{
			Pattern:      `DomainStatsBot`,
			Instances:    []string{ "DomainStatsBot/1.0 (http://domainstats.io/our-bot)" },
			URL:          stringer("http://domainstats.io/our-bot"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp224.MatchString(ua):
		return &Crawler{
			Pattern:      `mindUpBot`,
			Instances:    []string{ "mindUpBot (datenbutler.de)" },
			URL:          stringer("http://www.datenbutler.de/"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp225.MatchString(ua):
		return &Crawler{
			Pattern:      `Daum\/`,
			Instances:    []string{ "Mozilla/5.0 (compatible; Daum/4.1; +http://cs.daum.net/faq/15/4118.html?faqId=28966)" },
			URL:          stringer("http://cs.daum.net/faq/15/4118.html?faqId=28966"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp226.MatchString(ua):
		return &Crawler{
			Pattern:      `Jugendschutzprogramm-Crawler`,
			Instances:    []string{ "Jugendschutzprogramm-Crawler; Info: http://www.jugendschutzprogramm.de" },
			URL:          stringer("http://www.jugendschutzprogramm.de"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp227.MatchString(ua):
		return &Crawler{
			Pattern:      `Xenu Link Sleuth`,
			Instances:    []string{ "Xenu Link Sleuth/1.3.8" },
			URL:          stringer("http://home.snafu.de/tilman/xenulink.html"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp228.MatchString(ua):
		return &Crawler{
			Pattern:      `Pcore-HTTP`,
			Instances:    []string{ "Pcore-HTTP/v0.40.3" },
			URL:          stringer("https://bitbucket.org/softvisio/pcore/overview"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp229.MatchString(ua):
		return &Crawler{
			Pattern:      `moatbot`,
			Instances:    []string{ "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/40.0.2214.111 Safari/537.36 moatbot", "Mozilla/5.0 (iPhone; CPU iPhone OS 8_0 like Mac OS X) AppleWebKit/600.1.3 (KHTML, like Gecko) Version/8.0 Mobile/12A4345d Safari/600.1.4 moatbot" },
			URL:          stringer("https://moat.com"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp230.MatchString(ua):
		return &Crawler{
			Pattern:      `KosmioBot`,
			Instances:    []string{ "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/44.0.2403.125 Safari/537.36 (compatible; KosmioBot/1.0; +http://kosm.io/bot.html)" },
			URL:          stringer("http://kosm.io/bot.html"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp231.MatchString(ua):
		return &Crawler{
			Pattern:      `pingdom`,
			Instances:    []string{ "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Ubuntu Chromium/59.0.3071.109 Chrome/59.0.3071.109 Safari/537.36 PingdomPageSpeed/1.0 (pingbot/2.0; +http://www.pingdom.com/)", "Mozilla/5.0 (compatible; pingbot/2.0; +http://www.pingdom.com/)" },
			URL:          stringer("http://www.pingdom.com"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp232.MatchString(ua):
		return &Crawler{
			Pattern:      `AppInsights`,
			Instances:    []string{ "Mozilla/5.0 (compatible; MSIE 9.0; Windows NT 6.1; Trident/5.0; AppInsights)" },
			URL:          stringer("https://docs.microsoft.com/en-us/azure/azure-monitor/app/app-insights-overview"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp233.MatchString(ua):
		return &Crawler{
			Pattern:      `PhantomJS`,
			Instances:    []string{ "Mozilla/5.0 (Unknown; Linux x86_64) AppleWebKit/538.1 (KHTML, like Gecko) PhantomJS/2.1.1 Safari/538.1 bl.uk_lddc_renderbot/2.0.0 (+ http://www.bl.uk/aboutus/legaldeposit/websites/websites/faqswebmaster/index.html)" },
			URL:          stringer("http://phantomjs.org/"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp234.MatchString(ua):
		return &Crawler{
			Pattern:      `Gowikibot`,
			Instances:    []string{ "Mozilla/5.0 (compatible; Gowikibot/1.0; +http://www.gowikibot.com)" },
			URL:          stringer("http://www.gowikibot.com"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp235.MatchString(ua):
		return &Crawler{
			Pattern:      `PiplBot`,
			Instances:    []string{ "PiplBot (+http://www.pipl.com/bot/)", "Mozilla/5.0+(compatible;+PiplBot;+http://www.pipl.com/bot/)" },
			URL:          stringer("http://www.pipl.com/bot/"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp236.MatchString(ua):
		return &Crawler{
			Pattern:      `Discordbot`,
			Instances:    []string{ "Mozilla/5.0 (compatible; Discordbot/2.0; +https://discordapp.com)" },
			URL:          stringer("https://discordapp.com"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp237.MatchString(ua):
		return &Crawler{
			Pattern:      `TelegramBot`,
			Instances:    []string{ "TelegramBot (like TwitterBot)" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp238.MatchString(ua):
		return &Crawler{
			Pattern:      `Jetslide`,
			Instances:    []string{ "Mozilla/5.0 (compatible; Jetslide; +http://jetsli.de/crawler)" },
			URL:          stringer("http://jetsli.de/crawler"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp239.MatchString(ua):
		return &Crawler{
			Pattern:      `newsharecounts`,
			Instances:    []string{ "Mozilla/5.0 (compatible; NewShareCounts.com/1.0; +http://newsharecounts.com/crawler)" },
			URL:          stringer("http://newsharecounts.com/crawler"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp240.MatchString(ua):
		return &Crawler{
			Pattern:      `James BOT`,
			Instances:    []string{ "Mozilla/5.0 (Windows; U; Windows NT 5.1; en-US; rv:1.8.1.6) Gecko/20070725 Firefox/2.0.0.6 - James BOT - WebCrawler http://cognitiveseo.com/bot.html" },
			URL:          stringer("http://cognitiveseo.com/bot.html"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp241.MatchString(ua):
		return &Crawler{
			Pattern:      `Bark[rR]owler`,
			Instances:    []string{ "Barkrowler/0.5.1 (experimenting / debugging - sorry for your logs ) http://www.exensa.com/crawl - admin@exensa.com -- based on BuBiNG", "Barkrowler/0.7 (+http://www.exensa.com/crawl)", "BarkRowler/0.7 (+http://www.exensa.com/crawling)", "Barkrowler/0.9 (+http://www.exensa.com/crawl)" },
			URL:          stringer("http://www.exensa.com/crawl"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp242.MatchString(ua):
		return &Crawler{
			Pattern:      `TinEye`,
			Instances:    []string{ "Mozilla/5.0 (compatible; TinEye-bot/1.31; +http://www.tineye.com/crawler.html)", "TinEye/1.1 (http://tineye.com/crawler.html)" },
			URL:          stringer("http://www.tineye.com/crawler.html"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp243.MatchString(ua):
		return &Crawler{
			Pattern:      `SocialRankIOBot`,
			Instances:    []string{ "SocialRankIOBot; http://socialrank.io/about" },
			URL:          stringer("http://socialrank.io/about"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp244.MatchString(ua):
		return &Crawler{
			Pattern:      `trendictionbot`,
			Instances:    []string{ "Mozilla/5.0 (Windows; U; Windows NT 6.0; en-GB; rv:1.0; trendictionbot0.5.0; trendiction search; http://www.trendiction.de/bot; please let us know of any problems; web at trendiction.com) Gecko/20071127 Firefox/3.0.0.11" },
			URL:          stringer("http://www.trendiction.de/bot"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp245.MatchString(ua):
		return &Crawler{
			Pattern:      `Ocarinabot`,
			Instances:    []string{ "Ocarinabot" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp246.MatchString(ua):
		return &Crawler{
			Pattern:      `epicbot`,
			Instances:    []string{ "Mozilla/5.0 (compatible; epicbot; +http://www.epictions.com/epicbot)" },
			URL:          stringer("http://www.epictions.com/epicbot"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp247.MatchString(ua):
		return &Crawler{
			Pattern:      `Primalbot`,
			Instances:    []string{ "Mozilla/5.0 (compatible; Primalbot; +https://www.primal.com;)" },
			URL:          stringer("https://www.primal.com"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp248.MatchString(ua):
		return &Crawler{
			Pattern:      `DuckDuckGo-Favicons-Bot`,
			Instances:    []string{ "Mozilla/5.0 (compatible; DuckDuckGo-Favicons-Bot/1.0; +http://duckduckgo.com)" },
			URL:          stringer("http://duckduckgo.com"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp249.MatchString(ua):
		return &Crawler{
			Pattern:      `GnowitNewsbot`,
			Instances:    []string{ "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:49.0) Gecko/20100101 Firefox/49.0 / GnowitNewsbot / Contact information at http://www.gnowit.com" },
			URL:          stringer("http://www.gnowit.com"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp250.MatchString(ua):
		return &Crawler{
			Pattern:      `Leikibot`,
			Instances:    []string{ "Mozilla/5.0 (Windows NT 6.3;compatible; Leikibot/1.0; +http://www.leiki.com)" },
			URL:          stringer("http://www.leiki.com"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp251.MatchString(ua):
		return &Crawler{
			Pattern:      `LinkArchiver`,
			Instances:    []string{ "@LinkArchiver twitter bot" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp252.MatchString(ua):
		return &Crawler{
			Pattern:      `YaK\/`,
			Instances:    []string{ "Mozilla/5.0 (compatible; YaK/1.0; http://linkfluence.com/; bot@linkfluence.com)" },
			URL:          stringer("http://linkfluence.com"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp253.MatchString(ua):
		return &Crawler{
			Pattern:      `PaperLiBot`,
			Instances:    []string{ "Mozilla/5.0 (compatible; PaperLiBot/2.1; http://support.paper.li/entries/20023257-what-is-paper-li)", "Mozilla/5.0 (compatible; PaperLiBot/2.1; https://support.paper.li/entries/20023257-what-is-paper-li)" },
			URL:          stringer("http://support.paper.li/entries/20023257-what-is-paper-li"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp254.MatchString(ua):
		return &Crawler{
			Pattern:      `Digg Deeper`,
			Instances:    []string{ "Digg Deeper/v1 (http://digg.com/about)" },
			URL:          stringer("http://digg.com/about"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp255.MatchString(ua):
		return &Crawler{
			Pattern:      `dcrawl`,
			Instances:    []string{ "dcrawl/1.0" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp256.MatchString(ua):
		return &Crawler{
			Pattern:      `Snacktory`,
			Instances:    []string{ "Mozilla/5.0 (compatible; Snacktory; +https://github.com/karussell/snacktory)" },
			URL:          stringer("https://github.com/karussell/snacktory"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp257.MatchString(ua):
		return &Crawler{
			Pattern:      `AndersPinkBot`,
			Instances:    []string{ "Mozilla/5.0 (compatible; AndersPinkBot/1.0; +http://anderspink.com/bot.html)" },
			URL:          stringer("http://anderspink.com/bot.html"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp258.MatchString(ua):
		return &Crawler{
			Pattern:      `Fyrebot`,
			Instances:    []string{ "Fyrebot/1.0" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp259.MatchString(ua):
		return &Crawler{
			Pattern:      `EveryoneSocialBot`,
			Instances:    []string{ "Mozilla/5.0 (compatible; EveryoneSocialBot/1.0; support@everyonesocial.com http://everyonesocial.com/)" },
			URL:          stringer("http://everyonesocial.com"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp260.MatchString(ua):
		return &Crawler{
			Pattern:      `Mediatoolkitbot`,
			Instances:    []string{ "Mediatoolkitbot (complaints@mediatoolkit.com)" },
			URL:          stringer("http://mediatoolkit.com"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp261.MatchString(ua):
		return &Crawler{
			Pattern:      `Luminator-robots`,
			Instances:    []string{ "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_8_2) AppleWebKit/537.13 (KHTML, like Gecko) Chrome/30.0.1599.66 Safari/537.13 Luminator-robots/2.0" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp262.MatchString(ua):
		return &Crawler{
			Pattern:      `ExtLinksBot`,
			Instances:    []string{ "Mozilla/5.0 (compatible; ExtLinksBot/1.5 +https://extlinks.com/Bot.html)" },
			URL:          stringer("https://extlinks.com/Bot.html"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp263.MatchString(ua):
		return &Crawler{
			Pattern:      `SurveyBot`,
			Instances:    []string{ "Mozilla/5.0 (Windows; U; Windows NT 5.1; en; rv:1.9.0.13) Gecko/2009073022 Firefox/3.5.2 (.NET CLR 3.5.30729) SurveyBot/2.3 (DomainTools)" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp264.MatchString(ua):
		return &Crawler{
			Pattern:      `NING\/`,
			Instances:    []string{ "NING/1.0" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp265.MatchString(ua):
		return &Crawler{
			Pattern:      `okhttp`,
			Instances:    []string{ "okhttp/2.5.0", "okhttp/2.7.5", "okhttp/3.2.0", "okhttp/3.5.0" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp266.MatchString(ua):
		return &Crawler{
			Pattern:      `Nuzzel`,
			Instances:    []string{ "Nuzzel" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp267.MatchString(ua):
		return &Crawler{
			Pattern:      `omgili`,
			Instances:    []string{ "omgili/0.5 +http://omgili.com" },
			URL:          stringer("http://omgili.com"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp268.MatchString(ua):
		return &Crawler{
			Pattern:      `PocketParser`,
			Instances:    []string{ "PocketParser/2.0 (+https://getpocket.com/pocketparser_ua)" },
			URL:          stringer("https://getpocket.com/pocketparser_ua"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp269.MatchString(ua):
		return &Crawler{
			Pattern:      `YisouSpider`,
			Instances:    []string{ "YisouSpider" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp270.MatchString(ua):
		return &Crawler{
			Pattern:      `um-LN`,
			Instances:    []string{ "Mozilla/5.0 (compatible; um-LN/1.0; mailto: techinfo@ubermetrics-technologies.com)" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp271.MatchString(ua):
		return &Crawler{
			Pattern:      `ToutiaoSpider`,
			Instances:    []string{ "Mozilla/5.0 (compatible; ToutiaoSpider/1.0; http://web.toutiao.com/media_cooperation/;)" },
			URL:          stringer("http://web.toutiao.com/media_cooperation/"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp272.MatchString(ua):
		return &Crawler{
			Pattern:      `MuckRack`,
			Instances:    []string{ "Mozilla/5.0 (compatible; MuckRack/1.0; +http://muckrack.com)" },
			URL:          stringer("http://muckrack.com"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp273.MatchString(ua):
		return &Crawler{
			Pattern:      `Jamie's Spider`,
			Instances:    []string{ "Jamie's Spider (http://jamiembrown.com/)" },
			URL:          stringer("http://jamiembrown.com/"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp274.MatchString(ua):
		return &Crawler{
			Pattern:      `AHC\/`,
			Instances:    []string{ "AHC/2.0" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp275.MatchString(ua):
		return &Crawler{
			Pattern:      `NetcraftSurveyAgent`,
			Instances:    []string{ "Mozilla/5.0 (compatible; NetcraftSurveyAgent/1.0; +info@netcraft.com)" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp276.MatchString(ua):
		return &Crawler{
			Pattern:      `Laserlikebot`,
			Instances:    []string{ "Mozilla/5.0 (iPhone; CPU iPhone OS 8_3 like Mac OS X) AppleWebKit/600.1.4 (KHTML, like Gecko) Version/8.0 Mobile/12F70 Safari/600.1.4 (compatible; Laserlikebot/0.1)" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp277.MatchString(ua):
		return &Crawler{
			Pattern:      `Apache-HttpClient`,
			Instances:    []string{ "Apache-HttpClient/4.2.3 (java 1.5)", "Apache-HttpClient/4.2.5 (java 1.5)", "Apache-HttpClient/4.3.1 (java 1.5)", "Apache-HttpClient/4.3.3 (java 1.5)", "Apache-HttpClient/4.3.5 (java 1.5)", "Apache-HttpClient/4.4.1 (Java/1.8.0_65)", "Apache-HttpClient/4.5.3 (Java/1.8.0_121)" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp278.MatchString(ua):
		return &Crawler{
			Pattern:      `AppEngine-Google`,
			Instances:    []string{ "AppEngine-Google; (+http://code.google.com/appengine; appid: example)" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp279.MatchString(ua):
		return &Crawler{
			Pattern:      `Jetty`,
			Instances:    []string{ "Jetty/9.3.z-SNAPSHOT" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp280.MatchString(ua):
		return &Crawler{
			Pattern:      `Upflow`,
			Instances:    []string{ "Upflow/1.0" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp281.MatchString(ua):
		return &Crawler{
			Pattern:      `Thinklab`,
			Instances:    []string{ "Thinklab (thinklab.com)" },
			URL:          stringer("thinklab.com"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp282.MatchString(ua):
		return &Crawler{
			Pattern:      `Traackr.com`,
			Instances:    []string{ "Traackr.com" },
			URL:          stringer("Traackr.com"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp283.MatchString(ua):
		return &Crawler{
			Pattern:      `Twurly`,
			Instances:    []string{ "Ruby, Twurly v1.1 (http://twurly.org)" },
			URL:          stringer("http://twurly.org"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp284.MatchString(ua):
		return &Crawler{
			Pattern:      `Mastodon`,
			Instances:    []string{ "http.rb/2.2.2 (Mastodon/1.5.1; +https://example-masto-instance.org/)" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp285.MatchString(ua):
		return &Crawler{
			Pattern:      `http_get`,
			Instances:    []string{ "http_get" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp286.MatchString(ua):
		return &Crawler{
			Pattern:      `DnyzBot`,
			Instances:    []string{ "Mozilla/5.0 (compatible; DnyzBot/1.0)" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp287.MatchString(ua):
		return &Crawler{
			Pattern:      `botify`,
			Instances:    []string{ "Mozilla/5.0 (compatible; botify; http://botify.com)" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp288.MatchString(ua):
		return &Crawler{
			Pattern:      `007ac9 Crawler`,
			Instances:    []string{ "Mozilla/5.0 (compatible; 007ac9 Crawler; http://crawler.007ac9.net/)" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp289.MatchString(ua):
		return &Crawler{
			Pattern:      `BehloolBot`,
			Instances:    []string{ "Mozilla/5.0 (compatible; BehloolBot/beta; +http://www.webeaver.com/bot)" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp290.MatchString(ua):
		return &Crawler{
			Pattern:      `BrandVerity`,
			Instances:    []string{ "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.10; rv:41.0) Gecko/20100101 Firefox/55.0 BrandVerity/1.0 (http://www.brandverity.com/why-is-brandverity-visiting-me)" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp291.MatchString(ua):
		return &Crawler{
			Pattern:      `check_http`,
			Instances:    []string{ "check_http/v2.2.1 (nagios-plugins 2.2.1)" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp292.MatchString(ua):
		return &Crawler{
			Pattern:      `BDCbot`,
			Instances:    []string{ "Mozilla/5.0 (Windows NT 6.1; compatible; BDCbot/1.0; +http://bigweb.bigdatacorp.com.br/faq.aspx) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/41.0.2272.118 Safari/537.36", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; BDCbot/1.0; +http://bigweb.bigdatacorp.com.br/faq.aspx) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.100 Safari/537.36" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp293.MatchString(ua):
		return &Crawler{
			Pattern:      `ZumBot`,
			Instances:    []string{ "Mozilla/5.0 (compatible; ZumBot/1.0; http://help.zum.com/inquiry)" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp294.MatchString(ua):
		return &Crawler{
			Pattern:      `EZID`,
			Instances:    []string{ "EZID (EZID link checker; https://ezid.cdlib.org/)" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp295.MatchString(ua):
		return &Crawler{
			Pattern:      `ICC-Crawler`,
			Instances:    []string{ "ICC-Crawler/2.0 (Mozilla-compatible; ; http://ucri.nict.go.jp/en/icccrawler.html)" },
			URL:          stringer("http://ucri.nict.go.jp/en/icccrawler.html"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp296.MatchString(ua):
		return &Crawler{
			Pattern:      `ArchiveBot`,
			Instances:    []string{ "ArchiveTeam ArchiveBot/20170106.02 (wpull 2.0.2)" },
			URL:          stringer("https://github.com/ArchiveTeam/ArchiveBot"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp297.MatchString(ua):
		return &Crawler{
			Pattern:      `^LCC `,
			Instances:    []string{ "LCC (+http://corpora.informatik.uni-leipzig.de/crawler_faq.html)" },
			URL:          stringer("http://corpora.informatik.uni-leipzig.de/crawler_faq.html"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp298.MatchString(ua):
		return &Crawler{
			Pattern:      `filterdb.iss.net\/crawler`,
			Instances:    []string{ "Mozilla/5.0 (compatible; oBot/2.3.1; +http://filterdb.iss.net/crawler/)" },
			URL:          stringer("http://filterdb.iss.net/crawler/"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp299.MatchString(ua):
		return &Crawler{
			Pattern:      `BLP_bbot`,
			Instances:    []string{ "BLP_bbot/0.1" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp300.MatchString(ua):
		return &Crawler{
			Pattern:      `BomboraBot`,
			Instances:    []string{ "Mozilla/5.0 (compatible; BomboraBot/1.0; +http://www.bombora.com/bot)" },
			URL:          stringer("http://www.bombora.com/bot"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp301.MatchString(ua):
		return &Crawler{
			Pattern:      `Buck\/`,
			Instances:    []string{ "Buck/2.2; (+https://app.hypefactors.com/media-monitoring/about.html)" },
			URL:          stringer("https://app.hypefactors.com/media-monitoring/about.html"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp302.MatchString(ua):
		return &Crawler{
			Pattern:      `Companybook-Crawler`,
			Instances:    []string{ "Companybook-Crawler (+https://www.companybooknetworking.com/)" },
			URL:          stringer("https://www.companybooknetworking.com/"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp303.MatchString(ua):
		return &Crawler{
			Pattern:      `Genieo`,
			Instances:    []string{ "Mozilla/5.0 (compatible; Genieo/1.0 http://www.genieo.com/webfilter.html)" },
			URL:          stringer("http://www.genieo.com/webfilter.html"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp304.MatchString(ua):
		return &Crawler{
			Pattern:      `magpie-crawler`,
			Instances:    []string{ "magpie-crawler/1.1 (U; Linux amd64; en-GB; +http://www.brandwatch.net)" },
			URL:          stringer("http://www.brandwatch.net"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp305.MatchString(ua):
		return &Crawler{
			Pattern:      `MeltwaterNews`,
			Instances:    []string{ "MeltwaterNews www.meltwater.com" },
			URL:          stringer("http://www.meltwater.com"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp306.MatchString(ua):
		return &Crawler{
			Pattern:      `Moreover`,
			Instances:    []string{ "Mozilla/5.0 Moreover/5.1 (+http://www.moreover.com)" },
			URL:          stringer("http://www.moreover.com"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp307.MatchString(ua):
		return &Crawler{
			Pattern:      `newspaper\/`,
			Instances:    []string{ "newspaper/0.2.5", "newspaper/0.2.6", "newspaper/0.1.0.7" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp308.MatchString(ua):
		return &Crawler{
			Pattern:      `ScoutJet`,
			Instances:    []string{ "Mozilla/5.0 (compatible; ScoutJet; +http://www.scoutjet.com/)" },
			URL:          stringer("http://www.scoutjet.com/"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp309.MatchString(ua):
		return &Crawler{
			Pattern:      `(^| )sentry\/`,
			Instances:    []string{ "sentry/8.22.0 (https://sentry.io)" },
			URL:          stringer("https://sentry.io"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp310.MatchString(ua):
		return &Crawler{
			Pattern:      `StorygizeBot`,
			Instances:    []string{ "Mozilla/5.0 (compatible; StorygizeBot; http://www.storygize.com)" },
			URL:          stringer("http://www.storygize.com"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp311.MatchString(ua):
		return &Crawler{
			Pattern:      `UptimeRobot`,
			Instances:    []string{ "Mozilla/5.0+(compatible; UptimeRobot/2.0; http://www.uptimerobot.com/)" },
			URL:          stringer("http://www.uptimerobot.com/"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp312.MatchString(ua):
		return &Crawler{
			Pattern:      `OutclicksBot`,
			Instances:    []string{ "OutclicksBot/2 +https://www.outclicks.net/agent/VjzDygCuk4ubNmg40ZMbFqT0sIh7UfOKk8s8ZMiupUR", "OutclicksBot/2 +https://www.outclicks.net/agent/gIYbZ38dfAuhZkrFVl7sJBFOUhOVct6J1SvxgmBZgCe", "OutclicksBot/2 +https://www.outclicks.net/agent/PryJzTl8POCRHfvEUlRN5FKtZoWDQOBEvFJ2wh6KH5J", "OutclicksBot/2 +https://www.outclicks.net/agent/p2i4sNUh7eylJF1S6SGgRs5mP40ExlYvsr9GBxVQG6h" },
			URL:          stringer("https://www.outclicks.net"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp313.MatchString(ua):
		return &Crawler{
			Pattern:      `seoscanners`,
			Instances:    []string{ "Mozilla/5.0 (compatible; seoscanners.net/1; +spider@seoscanners.net)" },
			URL:          stringer("http://www.seoscanners.net/"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp314.MatchString(ua):
		return &Crawler{
			Pattern:      `Hatena`,
			Instances:    []string{ "Hatena Antenna/0.3", "Hatena::Russia::Crawler/0.01" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp315.MatchString(ua):
		return &Crawler{
			Pattern:      `Google Web Preview`,
			Instances:    []string{ "Mozilla/5.0 (Linux; U; Android 2.3.4; generic) AppleWebKit/537.36 (KHTML, like Gecko; Google Web Preview) Version/4.0 Mobile Safari/537.36" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp316.MatchString(ua):
		return &Crawler{
			Pattern:      `MauiBot`,
			Instances:    []string{ "MauiBot (crawler.feedback+wc@gmail.com)" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp317.MatchString(ua):
		return &Crawler{
			Pattern:      `AlphaBot`,
			Instances:    []string{ "Mozilla/5.0 (compatible; AlphaBot/3.2; +http://alphaseobot.com/bot.html)" },
			URL:          stringer("http://alphaseobot.com/bot.html"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp318.MatchString(ua):
		return &Crawler{
			Pattern:      `SBL-BOT`,
			Instances:    []string{ "SBL-BOT (http://sbl.net)" },
			URL:          stringer("http://sbl.net"),
			Description:  stringer("Bot of SoftByte BlackWidow"),
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp319.MatchString(ua):
		return &Crawler{
			Pattern:      `IAS crawler`,
			Instances:    []string{ "IAS crawler (ias_crawler; http://integralads.com/site-indexing-policy/)" },
			URL:          stringer("http://integralads.com/site-indexing-policy/"),
			Description:  stringer("Bot of Integral Ad Science, Inc."),
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp320.MatchString(ua):
		return &Crawler{
			Pattern:      `adscanner`,
			Instances:    []string{ "Mozilla/5.0 (compatible; adscanner/)" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp321.MatchString(ua):
		return &Crawler{
			Pattern:      `Netvibes`,
			Instances:    []string{ "Netvibes (crawler/bot; http://www.netvibes.com", "Netvibes (crawler; http://www.netvibes.com)" },
			URL:          stringer("http://www.netvibes.com"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp322.MatchString(ua):
		return &Crawler{
			Pattern:      `acapbot`,
			Instances:    []string{ "Mozilla/5.0 (compatible;acapbot/0.1;treat like Googlebot)", "Mozilla/5.0 (compatible;acapbot/0.1.;treat like Googlebot)" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp323.MatchString(ua):
		return &Crawler{
			Pattern:      `Baidu-YunGuanCe`,
			Instances:    []string{ "Baidu-YunGuanCe-Bot(ce.baidu.com)", "Baidu-YunGuanCe-SLABot(ce.baidu.com)", "Baidu-YunGuanCe-ScanBot(ce.baidu.com)", "Baidu-YunGuanCe-PerfBot(ce.baidu.com)", "Baidu-YunGuanCe-VSBot(ce.baidu.com)" },
			URL:          stringer("https://ce.baidu.com/topic/topic20150908"),
			Description:  stringer("Baidu Cloud Watch"),
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp324.MatchString(ua):
		return &Crawler{
			Pattern:      `bitlybot`,
			Instances:    []string{ "bitlybot/3.0 (+http://bit.ly/)", "bitlybot/2.0", "bitlybot" },
			URL:          stringer("http://bit.ly/"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp325.MatchString(ua):
		return &Crawler{
			Pattern:      `blogmuraBot`,
			Instances:    []string{ "blogmuraBot (+http://www.blogmura.com)" },
			URL:          stringer("http://www.blogmura.com"),
			Description:  stringer("A blog ranking site which links to blogs on just about every theme possible."),
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp326.MatchString(ua):
		return &Crawler{
			Pattern:      `Bot.AraTurka.com`,
			Instances:    []string{ "Bot.AraTurka.com/0.0.1" },
			URL:          stringer("http://www.araturka.com"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp327.MatchString(ua):
		return &Crawler{
			Pattern:      `bot-pge.chlooe.com`,
			Instances:    []string{ "bot-pge.chlooe.com/1.0.0 (+http://www.chlooe.com/)" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp328.MatchString(ua):
		return &Crawler{
			Pattern:      `BoxcarBot`,
			Instances:    []string{ "Mozilla/5.0 (compatible; BoxcarBot/1.1; +awesome@boxcar.io)" },
			URL:          stringer("https://boxcar.io/"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp329.MatchString(ua):
		return &Crawler{
			Pattern:      `BTWebClient`,
			Instances:    []string{ "BTWebClient/180B(9704)" },
			URL:          stringer("http://www.utorrent.com/"),
			Description:  stringer("µTorrent BitTorrent Client"),
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp330.MatchString(ua):
		return &Crawler{
			Pattern:      `ContextAd Bot`,
			Instances:    []string{ "Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.0;.NET CLR 1.0.3705; ContextAd Bot 1.0)", "ContextAd Bot 1.0" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp331.MatchString(ua):
		return &Crawler{
			Pattern:      `Digincore bot`,
			Instances:    []string{ "Mozilla/5.0 (compatible; Digincore bot; https://www.digincore.com/crawler.html for rules and instructions.)" },
			URL:          stringer("http://www.digincore.com/crawler.html"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp332.MatchString(ua):
		return &Crawler{
			Pattern:      `Disqus`,
			Instances:    []string{ "Disqus/1.0" },
			URL:          stringer("https://disqus.com/"),
			Description:  stringer("validate and quality check pages."),
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp333.MatchString(ua):
		return &Crawler{
			Pattern:      `Feedly`,
			Instances:    []string{ "Feedly/1.0 (+http://www.feedly.com/fetcher.html; like FeedFetcher-Google)", "FeedlyBot/1.0 (http://feedly.com)" },
			URL:          stringer("https://www.feedly.com/fetcher.html"),
			Description:  stringer("Feedly Fetcher is how Feedly grabs RSS or Atom feeds when users choose to add them to their Feedly or any of the other applications built on top of the feedly cloud."),
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp334.MatchString(ua):
		return &Crawler{
			Pattern:      `Fetch\/`,
			Instances:    []string{ "Fetch/2.0a (CMS Detection/Web/SEO analysis tool, see http://guess.scritch.org)" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp335.MatchString(ua):
		return &Crawler{
			Pattern:      `Fever`,
			Instances:    []string{ "Fever/1.38 (Feed Parser; http://feedafever.com; Allow like Gecko)" },
			URL:          stringer("http://feedafever.com"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp336.MatchString(ua):
		return &Crawler{
			Pattern:      `Flamingo_SearchEngine`,
			Instances:    []string{ "Flamingo_SearchEngine (+http://www.flamingosearch.com/bot)" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp337.MatchString(ua):
		return &Crawler{
			Pattern:      `FlipboardProxy`,
			Instances:    []string{ "Mozilla/5.0 (compatible; FlipboardProxy/1.1; +http://flipboard.com/browserproxy)", "Mozilla/5.0 (compatible; FlipboardProxy/1.2; +http://flipboard.com/browserproxy)", "Mozilla/5.0 (Macintosh; U; Intel Mac OS X 10.6; en-US; rv:1.9.2) Gecko/20100115 Firefox/3.6 (FlipboardProxy/1.1; +http://flipboard.com/browserproxy)", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.9; rv:28.0) Gecko/20100101 Firefox/28.0 (FlipboardProxy/1.1; +http://flipboard.com/browserproxy)" },
			URL:          stringer("https://about.flipboard.com/browserproxy/"),
			Description:  stringer("a proxy service to fetch, validate, and prepare certain elements of websites for presentation through the Flipboard Application"),
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp338.MatchString(ua):
		return &Crawler{
			Pattern:      `g2reader-bot`,
			Instances:    []string{ "g2reader-bot/1.0 (+http://www.g2reader.com/)" },
			URL:          stringer("http://www.g2reader.com/"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp339.MatchString(ua):
		return &Crawler{
			Pattern:      `G2 Web Services`,
			Instances:    []string{ "G2 Web Services/1.0 (built with StormCrawler Archetype 1.8; https://www.g2webservices.com/; developers@g2llc.com)" },
			URL:          stringer("https://www.g2webservices.com/"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp340.MatchString(ua):
		return &Crawler{
			Pattern:      `imrbot`,
			Instances:    []string{ "Mozilla/5.0 (compatible; imrbot/1.10.8 +http://www.mignify.com)" },
			URL:          stringer("http://www.mignify.com"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp341.MatchString(ua):
		return &Crawler{
			Pattern:      `K7MLWCBot`,
			Instances:    []string{ "K7MLWCBot/1.0 (+http://www.k7computing.com)" },
			URL:          stringer("http://www.k7computing.com"),
			Description:  stringer("Virus scanner"),
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp342.MatchString(ua):
		return &Crawler{
			Pattern:      `Kemvibot`,
			Instances:    []string{ "Kemvibot/1.0 (http://kemvi.com, marco@kemvi.com)" },
			URL:          stringer("http://kemvi.com"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp343.MatchString(ua):
		return &Crawler{
			Pattern:      `Landau-Media-Spider`,
			Instances:    []string{ "Landau-Media-Spider/1.0(http://bots.landaumedia.de/bot.html)" },
			URL:          stringer("http://bots.landaumedia.de/bot.html"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp344.MatchString(ua):
		return &Crawler{
			Pattern:      `linkapediabot`,
			Instances:    []string{ "linkapediabot (+http://www.linkapedia.com)" },
			URL:          stringer("http://www.linkapedia.com"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp345.MatchString(ua):
		return &Crawler{
			Pattern:      `vkShare`,
			Instances:    []string{ "Mozilla/5.0 (compatible; vkShare; +http://vk.com/dev/Share)" },
			URL:          stringer("http://vk.com/dev/Share"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp346.MatchString(ua):
		return &Crawler{
			Pattern:      `Siteimprove.com`,
			Instances:    []string{ "Mozilla/5.0 (compatible; MSIE 10.0; Windows NT 6.1; Trident/6.0) LinkCheck by Siteimprove.com", "Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 5.0) Match by Siteimprove.com", "Mozilla/5.0 (compatible; MSIE 10.0; Windows NT 6.1; Trident/6.0) SiteCheck-sitecrawl by Siteimprove.com", "Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 5.0) LinkCheck by Siteimprove.com" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp347.MatchString(ua):
		return &Crawler{
			Pattern:      `BLEXBot\/`,
			Instances:    []string{ "Mozilla/5.0 (compatible; BLEXBot/1.0; +http://webmeup-crawler.com/)" },
			URL:          stringer("http://webmeup-crawler.com"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp348.MatchString(ua):
		return &Crawler{
			Pattern:      `DareBoost`,
			Instances:    []string{ "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/62.0.3202.75 Safari/537.36 DareBoost" },
			URL:          stringer("https://www.dareboost.com/"),
			Description:  stringer("Bot to test, Analyze and Optimize website"),
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp349.MatchString(ua):
		return &Crawler{
			Pattern:      `ZuperlistBot\/`,
			Instances:    []string{ "Mozilla/5.0 (compatible; ZuperlistBot/1.0)" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp350.MatchString(ua):
		return &Crawler{
			Pattern:      `Miniflux\/`,
			Instances:    []string{ "Mozilla/5.0 (compatible; Miniflux/2.0.x-dev; +https://miniflux.net)", "Mozilla/5.0 (compatible; Miniflux/2.0.3; +https://miniflux.net)", "Mozilla/5.0 (compatible; Miniflux/2.0.7; +https://miniflux.net)", "Mozilla/5.0 (compatible; Miniflux/2.0.10; +https://miniflux.net)", "Mozilla/5.0 (compatibl$; Miniflux/2.0.x-dev; +https://miniflux.app)", "Mozilla/5.0 (compatible; Miniflux/2.0.11; +https://miniflux.app)", "Mozilla/5.0 (compatible; Miniflux/2.0.12; +https://miniflux.app)", "Mozilla/5.0 (compatible; Miniflux/ae1dc1a; +https://miniflux.app)", "Mozilla/5.0 (compatible; Miniflux/3b6e44c; +https://miniflux.app)" },
			URL:          stringer("https://miniflux.net"),
			Description:  stringer("Miniflux is a minimalist and opinionated feed reader."),
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp351.MatchString(ua):
		return &Crawler{
			Pattern:      `Feedspot`,
			Instances:    []string{ "Mozilla/5.0 (compatible; Feedspotbot/1.0; +http://www.feedspot.com/fs/bot)", "Mozilla/5.0 (compatible; Feedspot/1.0 (+https://www.feedspot.com/fs/fetcher; like FeedFetcher-Google)" },
			URL:          stringer("http://www.feedspot.com/fs/bot"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp352.MatchString(ua):
		return &Crawler{
			Pattern:      `Diffbot\/`,
			Instances:    []string{ "Mozilla/5.0 (Windows; U; Windows NT 5.1; en-US; rv:1.9.1.2) Gecko/20090729 Firefox/3.5.2 (.NET CLR 3.5.30729; Diffbot/0.1; +http://www.diffbot.com)" },
			URL:          stringer("http://www.diffbot.com"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp353.MatchString(ua):
		return &Crawler{
			Pattern:      `SEOkicks`,
			Instances:    []string{ "Mozilla/5.0 (compatible; SEOkicks; +https://www.seokicks.de/robot.html)" },
			URL:          stringer("https://www.seokicks.de/robot.html"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp354.MatchString(ua):
		return &Crawler{
			Pattern:      `tracemyfile`,
			Instances:    []string{ "Mozilla/5.0 (compatible; tracemyfile/1.0; +bot@tracemyfile.com)" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp355.MatchString(ua):
		return &Crawler{
			Pattern:      `Nimbostratus-Bot`,
			Instances:    []string{ "Mozilla/5.0 (compatible; Nimbostratus-Bot/v1.3.2; http://cloudsystemnetworks.com)" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp356.MatchString(ua):
		return &Crawler{
			Pattern:      `zgrab`,
			Instances:    []string{ "Mozilla/5.0 zgrab/0.x" },
			URL:          stringer("https://zmap.io/"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp357.MatchString(ua):
		return &Crawler{
			Pattern:      `PR-CY.RU`,
			Instances:    []string{ "Mozilla/5.0 (compatible; PR-CY.RU; + https://a.pr-cy.ru)" },
			URL:          stringer("https://a.pr-cy.ru/"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp358.MatchString(ua):
		return &Crawler{
			Pattern:      `AdsTxtCrawler`,
			Instances:    []string{ "AdsTxtCrawler/1.0" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp359.MatchString(ua):
		return &Crawler{
			Pattern:      `Datafeedwatch`,
			Instances:    []string{ "Datafeedwatch/2.1.x" },
			URL:          stringer("https://www.datafeedwatch.com/"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp360.MatchString(ua):
		return &Crawler{
			Pattern:      `Zabbix`,
			Instances:    []string{ "Zabbix" },
			URL:          stringer("https://www.zabbix.com/documentation/3.4/manual/web_monitoring"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp361.MatchString(ua):
		return &Crawler{
			Pattern:      `TangibleeBot`,
			Instances:    []string{ "TangibleeBot/1.0.0.0 (http://tangiblee.com/bot)" },
			URL:          stringer("http://tangiblee.com/bot"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp362.MatchString(ua):
		return &Crawler{
			Pattern:      `google-xrawler`,
			Instances:    []string{ "google-xrawler" },
			URL:          stringer("https://webmasters.stackexchange.com/questions/105560/what-is-the-google-xrawler-user-agent-used-for"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp363.MatchString(ua):
		return &Crawler{
			Pattern:      `axios`,
			Instances:    []string{ "axios/0.18.0" },
			URL:          stringer("https://github.com/axios/axios"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp364.MatchString(ua):
		return &Crawler{
			Pattern:      `Amazon CloudFront`,
			Instances:    []string{ "Amazon CloudFront" },
			URL:          stringer("https://aws.amazon.com/cloudfront/"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp365.MatchString(ua):
		return &Crawler{
			Pattern:      `Pulsepoint`,
			Instances:    []string{ "Pulsepoint XT3 web scraper" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp366.MatchString(ua):
		return &Crawler{
			Pattern:      `CloudFlare-AlwaysOnline`,
			Instances:    []string{ "Mozilla/5.0 (compatible; CloudFlare-AlwaysOnline/1.0; +http://www.cloudflare.com/always-online) AppleWebKit/534.34", "Mozilla/5.0 (compatible; CloudFlare-AlwaysOnline/1.0; +https://www.cloudflare.com/always-online) AppleWebKit/534.34" },
			URL:          stringer("https://www.cloudflare.com/always-online/"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp367.MatchString(ua):
		return &Crawler{
			Pattern:      `Google-Structured-Data-Testing-Tool`,
			Instances:    []string{ "Mozilla/5.0 (compatible; Google-Structured-Data-Testing-Tool +https://search.google.com/structured-data/testing-tool)", "Mozilla/5.0 (compatible; Google-Structured-Data-Testing-Tool +http://developers.google.com/structured-data/testing-tool/)" },
			URL:          stringer("https://search.google.com/structured-data/testing-tool"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp368.MatchString(ua):
		return &Crawler{
			Pattern:      `WordupInfoSearch`,
			Instances:    []string{ "WordupInfoSearch/1.0" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp369.MatchString(ua):
		return &Crawler{
			Pattern:      `WebDataStats`,
			Instances:    []string{ "Mozilla/5.0 (compatible; WebDataStats/1.0 ; +https://webdatastats.com/policy.html)" },
			URL:          stringer("https://webdatastats.com/"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp370.MatchString(ua):
		return &Crawler{
			Pattern:      `HttpUrlConnection`,
			Instances:    []string{ "Jersey/2.25.1 (HttpUrlConnection 1.8.0_141)" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp371.MatchString(ua):
		return &Crawler{
			Pattern:      `Seekport Crawler`,
			Instances:    []string{ "Mozilla/5.0 (compatible; Seekport Crawler; http://seekport.com/)" },
			URL:          stringer("http://seekport.com/"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp372.MatchString(ua):
		return &Crawler{
			Pattern:      `ZoomBot`,
			Instances:    []string{ "ZoomBot (Linkbot 1.0 http://suite.seozoom.it/bot.html)" },
			URL:          stringer("http://suite.seozoom.it/bot.html"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp373.MatchString(ua):
		return &Crawler{
			Pattern:      `VelenPublicWebCrawler`,
			Instances:    []string{ "VelenPublicWebCrawler (velen.io)" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp374.MatchString(ua):
		return &Crawler{
			Pattern:      `MoodleBot`,
			Instances:    []string{ "MoodleBot/1.0" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp375.MatchString(ua):
		return &Crawler{
			Pattern:      `jpg-newsbot`,
			Instances:    []string{ "jpg-newsbot/2.0; (+https://vipnytt.no/bots/)" },
			URL:          stringer("https://vipnytt.no/bots/"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp376.MatchString(ua):
		return &Crawler{
			Pattern:      `outbrain`,
			Instances:    []string{ "Mozilla/5.0 (Java) outbrain" },
			URL:          stringer("https://www.outbrain.com/help/advertisers/invalid-url/"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp377.MatchString(ua):
		return &Crawler{
			Pattern:      `W3C_Validator`,
			Instances:    []string{ "W3C_Validator/1.3" },
			URL:          stringer("https://validator.w3.org/services"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp378.MatchString(ua):
		return &Crawler{
			Pattern:      `Validator\.nu`,
			Instances:    []string{ "Validator.nu/LV" },
			URL:          stringer("https://validator.w3.org/services"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp379.MatchString(ua):
		return &Crawler{
			Pattern:      `W3C-checklink`,
			Instances:    []string{ "W3C-checklink/2.90 libwww-perl/5.64", "W3C-checklink/3.6.2.3 libwww-perl/5.64", "W3C-checklink/4.2 [4.20] libwww-perl/5.803", "W3C-checklink/4.2.1 [4.21] libwww-perl/5.803", "W3C-checklink/4.3 [4.42] libwww-perl/5.805", "W3C-checklink/4.3 [4.42] libwww-perl/5.808", "W3C-checklink/4.3 [4.42] libwww-perl/5.820", "W3C-checklink/4.5 [4.154] libwww-perl/5.823", "W3C-checklink/4.5 [4.160] libwww-perl/5.823" },
			URL:          stringer("https://validator.w3.org/services"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{ "libwww-perl" },
		}

	case rgxp380.MatchString(ua):
		return &Crawler{
			Pattern:      `W3C-mobileOK`,
			Instances:    []string{ "W3C-mobileOK/DDC-1.0" },
			URL:          stringer("https://validator.w3.org/services"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp381.MatchString(ua):
		return &Crawler{
			Pattern:      `W3C_I18n-Checker`,
			Instances:    []string{ "W3C_I18n-Checker/1.0" },
			URL:          stringer("https://validator.w3.org/services"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp382.MatchString(ua):
		return &Crawler{
			Pattern:      `FeedValidator`,
			Instances:    []string{ "FeedValidator/1.3" },
			URL:          stringer("https://validator.w3.org/services"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp383.MatchString(ua):
		return &Crawler{
			Pattern:      `W3C_CSS_Validator`,
			Instances:    []string{ "Jigsaw/2.3.0 W3C_CSS_Validator_JFouffa/2.0" },
			URL:          stringer("https://validator.w3.org/services"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp384.MatchString(ua):
		return &Crawler{
			Pattern:      `W3C_Unicorn`,
			Instances:    []string{ "W3C_Unicorn/1.0" },
			URL:          stringer("https://validator.w3.org/services"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp385.MatchString(ua):
		return &Crawler{
			Pattern:      `Google-PhysicalWeb`,
			Instances:    []string{ "Mozilla/5.0 (Google-PhysicalWeb)" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp386.MatchString(ua):
		return &Crawler{
			Pattern:      `Blackboard`,
			Instances:    []string{ "Blackboard Safeassign" },
			URL:          stringer("https://help.blackboard.com/Learn/Administrator/Hosting/Tools_Management/SafeAssign"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp387.MatchString(ua):
		return &Crawler{
			Pattern:      `ICBot\/`,
			Instances:    []string{ "Mozilla/5.0 (compatible; ICBot/0.1; +https://ideasandcode.xyz" },
			URL:          stringer("https://ideasandcode.xyz"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp388.MatchString(ua):
		return &Crawler{
			Pattern:      `BazQux`,
			Instances:    []string{ "Mozilla/5.0 (compatible; BazQux/2.4; +https://bazqux.com/fetcher; 1 subscribers)" },
			URL:          stringer("https://bazqux.com/fetcher"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp389.MatchString(ua):
		return &Crawler{
			Pattern:      `Twingly`,
			Instances:    []string{ "Mozilla/5.0 (compatible; Twingly Recon; twingly.com)" },
			URL:          stringer("https://twingly.com"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp390.MatchString(ua):
		return &Crawler{
			Pattern:      `Rivva`,
			Instances:    []string{ "Mozilla/5.0 (compatible; Rivva; http://rivva.de)" },
			URL:          stringer("http://rivva.de"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp391.MatchString(ua):
		return &Crawler{
			Pattern:      `Experibot`,
			Instances:    []string{ "Experibot-v2 http://goo.gl/ZAr8wX", "Experibot-v3 http://goo.gl/ZAr8wX" },
			URL:          stringer("https://amirkr.wixsite.com/experibot"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp392.MatchString(ua):
		return &Crawler{
			Pattern:      `awesomecrawler`,
			Instances:    []string{ "Mozilla/5.0 (Windows NT 6.2; WOW64) AppleWebKit/537.22 (KHTML, like Gecko) Chrome/25.0.1364.5 Safari/537.22 +awesomecrawler" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp393.MatchString(ua):
		return &Crawler{
			Pattern:      `Dataprovider.com`,
			Instances:    []string{ "Mozilla/5.0 (compatible; Dataprovider.com)" },
			URL:          stringer("https://www.dataprovider.com/"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp394.MatchString(ua):
		return &Crawler{
			Pattern:      `GroupHigh\/`,
			Instances:    []string{ "Mozilla/5.0 (compatible; GroupHigh/1.0; +http://www.grouphigh.com/" },
			URL:          stringer("http://www.grouphigh.com/"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp395.MatchString(ua):
		return &Crawler{
			Pattern:      `theoldreader.com`,
			Instances:    []string{ "Mozilla/5.0 (compatible; theoldreader.com)" },
			URL:          stringer("https://www.theoldreader.com/"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp396.MatchString(ua):
		return &Crawler{
			Pattern:      `AnyEvent`,
			Instances:    []string{ "Mozilla/5.0 (compatible; U; AnyEvent-HTTP/2.24; +http://software.schmorp.de/pkg/AnyEvent)" },
			URL:          stringer("http://software.schmorp.de/pkg/AnyEvent.html"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp397.MatchString(ua):
		return &Crawler{
			Pattern:      `Uptimebot\.org`,
			Instances:    []string{ "Uptimebot.org - Free website monitoring" },
			URL:          stringer("http://uptimebot.org/"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp398.MatchString(ua):
		return &Crawler{
			Pattern:      `Nmap Scripting Engine`,
			Instances:    []string{ "Mozilla/5.0 (compatible; Nmap Scripting Engine; https://nmap.org/book/nse.html)" },
			URL:          stringer("https://nmap.org/book/nse.html"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp399.MatchString(ua):
		return &Crawler{
			Pattern:      `2ip.ru`,
			Instances:    []string{ "2ip.ru CMS Detector (https://2ip.ru/cms/)" },
			URL:          stringer("https://2ip.ru/cms/"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp400.MatchString(ua):
		return &Crawler{
			Pattern:      `Clickagy`,
			Instances:    []string{ "Clickagy Intelligence Bot v2" },
			URL:          stringer("https://www.clickagy.com"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp401.MatchString(ua):
		return &Crawler{
			Pattern:      `Caliperbot`,
			Instances:    []string{ "Caliperbot/1.0 (+http://www.conductor.com/caliperbot)" },
			URL:          stringer("http://www.conductor.com/caliperbot"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp402.MatchString(ua):
		return &Crawler{
			Pattern:      `MBCrawler`,
			Instances:    []string{ "MBCrawler/1.0 (https://monitorbacklinks.com)" },
			URL:          stringer("https://monitorbacklinks.com"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp403.MatchString(ua):
		return &Crawler{
			Pattern:      `online-webceo-bot`,
			Instances:    []string{ "Mozilla/5.0 (compatible; online-webceo-bot/1.0; +http://online.webceo.com)" },
			URL:          stringer("http://online.webceo.com"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp404.MatchString(ua):
		return &Crawler{
			Pattern:      `B2B Bot`,
			Instances:    []string{ "B2B Bot" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp405.MatchString(ua):
		return &Crawler{
			Pattern:      `AddSearchBot`,
			Instances:    []string{ "Mozilla/5.0 (compatible; AddSearchBot/0.9; +http://www.addsearch.com/bot; info@addsearch.com)" },
			URL:          stringer("http://www.addsearch.com/bot"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp406.MatchString(ua):
		return &Crawler{
			Pattern:      `Google Favicon`,
			Instances:    []string{ "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/49.0.2623.75 Safari/537.36 Google Favicon" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp407.MatchString(ua):
		return &Crawler{
			Pattern:      `HubSpot`,
			Instances:    []string{ "Mozilla/5.0 (Windows; U; Windows NT 5.1; en-US) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/27.0.1453.116 Safari/537.36 HubSpot Webcrawler - web-crawlers@hubspot.com", "Mozilla/5.0 (X11; Linux x86_64; HubSpot Single Page link check; web-crawlers+links@hubspot.com) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/51.0.2704.103 Safari/537.36", "Mozilla/5.0 (compatible; HubSpot Crawler; web-crawlers@hubspot.com)" },
			URL:          nil,
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp408.MatchString(ua):
		return &Crawler{
			Pattern:      `Chrome-Lighthouse`,
			Instances:    []string{ "Mozilla/5.0 (Linux; Android 6.0.1; Nexus 5 Build/MRA58N) AppleWebKit/537.36(KHTML, like Gecko) Chrome/69.0.3464.0 Mobile Safari/537.36 Chrome-Lighthouse", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36(KHTML, like Gecko) Chrome/69.0.3464.0 Safari/537.36 Chrome-Lighthouse", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3694.0 Safari/537.36 Chrome-Lighthouse", "Mozilla/5.0 (Linux; Android 6.0.1; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3694.0 Mobile Safari/537.36 Chrome-Lighthouse" },
			URL:          stringer("https://developers.google.com/speed/pagespeed/insights"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp409.MatchString(ua):
		return &Crawler{
			Pattern:      `HeadlessChrome`,
			Instances:    []string{ "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) HeadlessChrome/74.0.3729.169 Safari/537.36" },
			URL:          stringer("https://developers.google.com/web/updates/2017/04/headless-chrome"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp410.MatchString(ua):
		return &Crawler{
			Pattern:      `CheckMarkNetwork\/`,
			Instances:    []string{ "CheckMarkNetwork/1.0 (+http://www.checkmarknetwork.com/spider.html)" },
			URL:          stringer("https://www.checkmarknetwork.com/"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp411.MatchString(ua):
		return &Crawler{
			Pattern:      `www\.uptime\.com`,
			Instances:    []string{ "Mozilla/5.0 (compatible; Uptimebot/1.0; +http://www.uptime.com/uptimebot)" },
			URL:          stringer("http://www.uptime.com/uptimebot"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	case rgxp412.MatchString(ua):
		return &Crawler{
			Pattern:      `Streamline3Bot\/`,
			Instances:    []string{ "Mozilla/5.0 (compatible; MSIE 8.0; Windows NT 5.1) Streamline3Bot/1.0", "Mozilla/5.0 (Windows NT 6.1; Win64; x64; +https://www.ubtsupport.com/legal/Streamline3Bot.php) Streamline3Bot/1.0" },
			URL:          stringer("https://www.ubtsupport.com/legal/Streamline3Bot.php"),
			Description:  nil,
			AdditionDate: nil,
			DependsOn:    []string{  },
		}

	default:
		return nil
	}
}

