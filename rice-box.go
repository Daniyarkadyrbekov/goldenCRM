package main

import (
	"time"

	"github.com/GeertJohan/go.rice/embedded"
)

func init() {

	// define files
	file2 := &embedded.EmbeddedFile{
		Filename:    "base_bottom.html",
		FileModTime: time.Unix(1588322637, 0),

		Content: string("{{ define \"base_bottom\" }}\n    <!-- Optional JavaScript -->\n    <!-- jQuery first, then Popper.js, then Bootstrap JS -->\n    <script src=\"https://code.jquery.com/jquery-3.4.1.slim.min.js\" integrity=\"sha384-J6qa4849blE2+poT4WnyKhv5vZF5SrPo0iEjwBvKU7imGFAV0wwj1yYfoRSJoZ+n\" crossorigin=\"anonymous\"></script>\n    <script src=\"https://cdn.jsdelivr.net/npm/popper.js@1.16.0/dist/umd/popper.min.js\" integrity=\"sha384-Q6E9RHvbIyZFJoft+2mJbHaEWldlvI9IOYy5n3zV9zzTtmI3UksdQRVvoxMfooAo\" crossorigin=\"anonymous\"></script>\n    <script src=\"https://stackpath.bootstrapcdn.com/bootstrap/4.4.1/js/bootstrap.min.js\" integrity=\"sha384-wfSDF2E50Y2D1uUdj0O3uMBJnjuUD4Ih7YwaYd1iqfktj0Uod8GCExl3Og8ifwB6\" crossorigin=\"anonymous\"></script>\n</body>\n</html>\n{{ end }}"),
	}
	file3 := &embedded.EmbeddedFile{
		Filename:    "base_header.html",
		FileModTime: time.Unix(1588322680, 0),

		Content: string("{{ define \"base_header\" }}\n<html lang=\"en\">\n<head>\n    <!-- Required meta tags -->\n    <meta charset=\"utf-8\">\n    <meta name=\"viewport\" content=\"width=device-width, initial-scale=1, shrink-to-fit=no\">\n\n    <!-- Bootstrap CSS -->\n    <link rel=\"stylesheet\" href=\"https://stackpath.bootstrapcdn.com/bootstrap/4.4.1/css/bootstrap.min.css\" integrity=\"sha384-Vkoo8x4CGsO3+Hhxv8T/Q5PaXtkKtu6ug5TOeNV6gBiFeWPGFN9MuhOf23Q9Ifjh\" crossorigin=\"anonymous\">\n</head>\n{{ end }}"),
	}
	file4 := &embedded.EmbeddedFile{
		Filename:    "base_nav_flat_form.html",
		FileModTime: time.Unix(1589025448, 0),

		Content: string("{{ define \"base_nav_flat_form\" }}\n    <form id=\"my-form\" action=\"/flat/new\" method=\"POST\">\n        <div class=\"form-row\">\n            <div class=\"form-group col-md-6\">\n                <label for=\"inputEmail4\">Email</label>\n                <input type=\"email\" class=\"form-control\" id=\"inputEmail4\">\n            </div>\n            <div class=\"form-group col-md-6\">\n                <label for=\"inputPassword4\">Password</label>\n                <input type=\"password\" class=\"form-control\" id=\"inputPassword4\">\n            </div>\n        </div>\n        <div class=\"form-group\">\n            <label for=\"inputAddress\">Address</label>\n            <input type=\"text\" class=\"form-control\" id=\"inputAddress\" placeholder=\"1234 Main St\">\n        </div>\n        <div class=\"form-group\">\n            <label for=\"inputAddress2\">Address 2</label>\n            <input type=\"text\" class=\"form-control\" id=\"inputAddress2\" placeholder=\"Apartment, studio, or floor\">\n        </div>\n        <div class=\"form-row\">\n            <div class=\"form-group col-md-6\">\n                <label for=\"inputCity\">City</label>\n                <input type=\"text\" class=\"form-control\" id=\"inputCity\">\n            </div>\n            <div class=\"form-group col-md-4\">\n                <label for=\"inputState\">State</label>\n                <select id=\"inputState\" class=\"form-control\">\n                    <option selected>Choose...</option>\n                    <option>...</option>\n                </select>\n            </div>\n            <div class=\"form-group col-md-2\">\n                <label for=\"inputZip\">Zip</label>\n                <input type=\"text\" class=\"form-control\" id=\"inputZip\">\n            </div>\n        </div>\n        <div class=\"form-group\">\n            <div class=\"form-check\">\n                <input class=\"form-check-input\" type=\"checkbox\" id=\"gridCheck\">\n                <label class=\"form-check-label\" for=\"gridCheck\">\n                    Check me out\n                </label>\n            </div>\n        </div>\n        <button type=\"submit\" class=\"btn btn-primary\">Sign in</button>\n    </form>\n{{end}}"),
	}
	file5 := &embedded.EmbeddedFile{
		Filename:    "base_navbar.html",
		FileModTime: time.Unix(1588828393, 0),

		Content: string("{{ define \"base_navbar\"}}\n    <div class=\"pos-f-t\">\n        <nav class=\"navbar navbar-expand-lg navbar-light bg-light\">\n            <a class=\"navbar-brand\" href=\"#\">Золотой Ключ</a>\n            <button class=\"navbar-toggler\" type=\"button\" data-toggle=\"collapse\" data-target=\"#navbarSupportedContent\" aria-controls=\"navbarSupportedContent\" aria-expanded=\"false\" aria-label=\"Toggle navigation\">\n                <span class=\"navbar-toggler-icon\"></span>\n            </button>\n\n            <div class=\"collapse navbar-collapse\" id=\"navbarSupportedContent\">\n                <form class=\"form-inline my-2 my-lg-0\">\n                    <button class=\"btn btn-sm btn-outline-secondary\" type=\"button\" data-toggle=\"collapse\" data-target=\"#navbarCollapseAdd\" aria-controls=\"navbarToggleExternalContent\" aria-expanded=\"false\" aria-label=\"Toggle navigation\">\n                        добавить\n                    </button>\n                </form>\n                <form class=\"form-inline my-2 my-lg-0\">\n                    <button class=\"btn btn-sm btn-outline-secondary\" type=\"button\" data-toggle=\"collapse\" data-target=\"#navbarCollapseSearch\" aria-controls=\"navbarToggleExternalContent\" aria-expanded=\"false\" aria-label=\"Toggle navigation\">\n                        найти\n                    </button>\n                </form>\n                <ul class=\"navbar-nav mr-auto\">\n                </ul>\n                <ul class=\"nav justify-content-end\">\n                    <li class=\"nav-item dropdown\">\n                        <a class=\"nav-link dropdown-toggle\" href=\"#\" id=\"navbarDropdown\" role=\"button\" data-toggle=\"dropdown\" aria-haspopup=\"true\" aria-expanded=\"false\">\n                            {{ .user.GetFullName }}\n                        </a>\n                        <div class=\"dropdown-menu\" aria-labelledby=\"navbarDropdown\">\n                            <a class=\"dropdown-item\" href=\"#\">Профиль</a>\n                            <div class=\"dropdown-divider\"></div>\n                            <a class=\"dropdown-item\" href=\"#\">Выйти</a>\n                        </div>\n                    </li>\n                </ul>\n            </div>\n        </nav>\n        <div class=\"collapse\" id=\"navbarCollapseSearch\">\n            <div class=\"p-4\">\n                <h5 class=\"h4\">Найти</h5>\n                {{template \"base_nav_flat_form\" .}}\n            </div>\n        </div>\n        <div class=\"collapse\" id=\"navbarCollapseAdd\">\n            <div class=\"p-4\">\n                <h5 class=\"h4\">Добавить</h5>\n                {{template \"base_nav_flat_form\" .}}\n            </div>\n        </div>\n    </div>\n\n{{ end }}"),
	}
	file6 := &embedded.EmbeddedFile{
		Filename:    "base_search_form.html",
		FileModTime: time.Unix(1587907986, 0),

		Content: string("{{ define \"base_search_form\" }}\n    <h1>Search form</h1>\n{{ end }}"),
	}
	file7 := &embedded.EmbeddedFile{
		Filename:    "index.html",
		FileModTime: time.Unix(1589026849, 0),

		Content: string("{{ template \"base_header\" . }}\n{{ template \"base_navbar\" . }}\n\n<table class=\"table\">\n    <thead>\n    <tr>\n        <th scope=\"col\">ID</th>\n        <th scope=\"col\">Улица</th>\n        <th scope=\"col\">дом</th>\n        <th scope=\"col\">строение</th>\n        <th scope=\"col\">номер Квартиры</th>\n        <th scope=\"col\">состояние</th>\n        <th scope=\"col\">этаж</th>\n        <th scope=\"col\">угловая</th>\n        <th scope=\"col\">тип</th>\n        <th scope=\"col\">описание</th>\n        <th scope=\"col\">владелец</th>\n    </tr>\n    </thead>\n\n{{/*    ID          int64*/}}\n{{/*    Street      string*/}}\n{{/*    Home        string*/}}\n{{/*    Structure   uint*/}}\n{{/*    FlatNumber  uint*/}}\n{{/*    State       flatState*/}}\n{{/*    Floor       uint*/}}\n{{/*    IsCorner    bool*/}}\n{{/*    FlatType    string*/}}\n{{/*    Description string*/}}\n{{/*    PictureURLs []string*/}}\n{{/*    Owner       string*/}}\n    <tbody>\n    {{range .flats}}\n        <tr>\n            <td>{{.ID}}</td>\n            <td>{{.Street}}</td>\n            <td>{{.Home}}</td>\n            <td>{{.Structure}}</td>\n            <td>{{.FlatNumber}}</td>\n            <td>{{.State}}</td>\n            <td>{{.Floor}}</td>\n            <td>{{.IsCorner}}</td>\n            <td>{{.FlatType}}</td>\n            <td>{{.Description}}</td>\n            <td>{{.Owner}}</td>\n        </tr>\n    {{end}}\n    </tbody>\n</table>\n{{ template \"base_bottom\" . }}"),
	}

	// define dirs
	dir1 := &embedded.EmbeddedDir{
		Filename:   "",
		DirModTime: time.Unix(1589026849, 0),
		ChildFiles: []*embedded.EmbeddedFile{
			file2, // "base_bottom.html"
			file3, // "base_header.html"
			file4, // "base_nav_flat_form.html"
			file5, // "base_navbar.html"
			file6, // "base_search_form.html"
			file7, // "index.html"

		},
	}

	// link ChildDirs
	dir1.ChildDirs = []*embedded.EmbeddedDir{}

	// register embeddedBox
	embedded.RegisterEmbeddedBox(`pages/templates`, &embedded.EmbeddedBox{
		Name: `pages/templates`,
		Time: time.Unix(1589026849, 0),
		Dirs: map[string]*embedded.EmbeddedDir{
			"": dir1,
		},
		Files: map[string]*embedded.EmbeddedFile{
			"base_bottom.html":        file2,
			"base_header.html":        file3,
			"base_nav_flat_form.html": file4,
			"base_navbar.html":        file5,
			"base_search_form.html":   file6,
			"index.html":              file7,
		},
	})
}
