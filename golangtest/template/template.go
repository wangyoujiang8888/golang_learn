package template

import (
	"fmt"
	"html/template"
	"os"
	"strings"
)

type Person struct {
	UserName string
}

func init() {
	t := template.New("fieldname example")
	t.Parse("hello {{.UserName}}")
	p := Person{UserName: "Astaxie"}
	t.Execute(os.Stdout, p)
	manyOut()
	emptyTest()
	templateTest()

}

type Friend struct {
	Fname string
}

type Person1 struct {
	UserName string
	Emails   []string
	Friends  []*Friend
	Tels     []string
}

func EmailDealWith(args ...interface{}) string {
	ok := false
	var s string
	if len(args) == 1 {
		s, ok = args[0].(string)
	}
	if !ok {
		s = fmt.Sprint(args...)
	}
	// find the @ symbol
	substrs := strings.Split(s, "@")
	if len(substrs) != 2 {
		return s
	}
	// replace the @ by " at "
	return (substrs[0] + " at " + substrs[1])
}

func manyOut() {
	f1 := Friend{Fname: "minux.ma"}
	f2 := Friend{Fname: "xushiwei"}
	t := template.New("fieldname example")
	t = t.Funcs(template.FuncMap{"emailDeal": EmailDealWith})

	t, _ = t.Parse(`hello {{.UserName}}!
                {{range .Emails}}
                an email {{.|emailDeal}}
                {{end}}
                {{range .Tels}} an tel {{.}}{{end}}
                {{with .Friends}}
                {{range .}}
                    my friend name is {{.Fname}}
                {{end}}
                {{end}}
                `)
	p := Person1{UserName: "Astaxie",
		Emails:  []string{"astaxie@beego.me", "astaxie@gmail.com"},
		Tels:    []string{"187185655522", "1245646411212"},
		Friends: []*Friend{&f1, &f2}}
	t.Execute(os.Stdout, p)

}

func emptyTest() {
	tEmpty := template.New("template test")
	tEmpty = template.Must(tEmpty.Parse("空 pipeline if demo: {{if ``}} 不会输出. {{end}}\n"))
	tEmpty.Execute(os.Stdout, nil)

	tWithValue := template.New("template test")
	tWithValue = template.Must(tWithValue.Parse("不为空的 pipeline if demo: {{if `anything`}} 我有内容，我会输出. {{end}}\n"))
	tWithValue.Execute(os.Stdout, nil)

	tIfElse := template.New("template test")
	tIfElse = template.Must(tIfElse.Parse("if-else demo: {{if `anything`}} if部分 {{else}} else部分.{{end}}\n"))
	tIfElse.Execute(os.Stdout, nil)

}

func templateTest() {

	s1, _ := template.ParseFiles("header.tmpl", "content.tmpl", "footer.tmpl")
	s1.ExecuteTemplate(os.Stdout, "content", nil)
}
