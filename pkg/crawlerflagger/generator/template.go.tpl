package crawlerflagger

type Crawler struct {
	Pattern      string
	Instances    []string
	URL          *string
	Description  *string
	AdditionDate *string
	DependsOn    []string
}

func IdentifyCrawler(ua string) *Crawler {
	switch ua {
{{ range . }}{{ if .Instances }}
	case {{ range $i, $e := .Instances }}{{if $i}}{{printf ",\n        "}}{{end}}"{{ . }}"{{ end }}:
		return &Crawler{
			Pattern:      `{{ .Pattern }}`,
			Instances:    []string{ {{ range $i, $e := .Instances }}{{if $i}}, {{end}}"{{ . }}"{{ end }} },
			URL:          {{ if .URL }}stringer("{{ .URL }}"){{ else }}nil{{ end }},
			Description:  {{ if .Description }}stringer("{{ .Description }}"){{ else }}nil{{ end }},
			AdditionDate: {{ if .AdditionDate }}stringer("{{ .AdditionDate }}"){{ else }}nil{{ end }},
			DependsOn:    []string{ {{ range $i, $e := .DependsOn }}{{if $i}}, {{end}}"{{ . }}"{{ end }} },
		}
{{ end }}{{ end }}
	default:
		return nil
	}
}
