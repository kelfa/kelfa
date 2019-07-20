package crawlerflagger

type Crawler struct {
	Pattern      string
	Instances    []string
	URL          *string
	Description  *string
	AdditionDate *string
	DependsOn    *string
}

func IdentifyCrawler(ua string) *Crawler {
	switch ua {
{{ range . }}
	case {{ range $i, $e := .Instances }} {{if $i}}, {{end}} "{{ . }}"{{ end }}:
		return Crawler{
			Pattern:      "{{ .Pattern }}",
			Instances:    []string{ {{ range $i, $e := .Instances }} {{if $i}}, {{end}} "{{ . }}"{{ end }} },
			URL:          {{ if .URL }}"{{ .URL }}"{{ else }}nil{{ end }},
			Description:  {{ if .Description }}"{{ .Description }}"{{ else }}nil{{ end }},
			AdditionDate: {{ if .AdditionDate }}"{{ .AdditionDate }}"{{ else }}nil{{ end }},
			DependsOn:    {{ if .DependsOn }}"{{ .DependsOn }}"{{ else }}nil{{ end }},
		}
{{ end }}
	default:
		return nil
	}
}
