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
		Filename:    "base_navbar.html",
		FileModTime: time.Unix(1588825432, 0),

		Content: string("{{ define \"base_navbar\"}}\n    <div class=\"pos-f-t\">\n        <nav class=\"navbar navbar-expand-lg navbar-light bg-light\">\n            <a class=\"navbar-brand\" href=\"#\">Золотой Ключ</a>\n            <button class=\"navbar-toggler\" type=\"button\" data-toggle=\"collapse\" data-target=\"#navbarSupportedContent\" aria-controls=\"navbarSupportedContent\" aria-expanded=\"false\" aria-label=\"Toggle navigation\">\n                <span class=\"navbar-toggler-icon\"></span>\n            </button>\n\n            <div class=\"collapse navbar-collapse\" id=\"navbarSupportedContent\">\n                <form class=\"form-inline my-2 my-lg-0\">\n                    <button class=\"btn btn-sm btn-outline-secondary\" type=\"button\" data-toggle=\"collapse\" data-target=\"#navbarCollapseAdd\" aria-controls=\"navbarToggleExternalContent\" aria-expanded=\"false\" aria-label=\"Toggle navigation\">\n                        добавить\n                    </button>\n                </form>\n                <form class=\"form-inline my-2 my-lg-0\">\n                    <button class=\"btn btn-sm btn-outline-secondary\" type=\"button\" data-toggle=\"collapse\" data-target=\"#navbarCollapseSearch\" aria-controls=\"navbarToggleExternalContent\" aria-expanded=\"false\" aria-label=\"Toggle navigation\">\n                        найти\n                    </button>\n                </form>\n{{/*                <form class=\"form-inline my-2 my-lg-0\">*/}}\n{{/*                    <button class=\"btn btn-sm btn-outline-secondary\" type=\"button\" data-toggle=\"modal\" data-target=\"#addModal\">*/}}\n{{/*                        добавить*/}}\n{{/*                    </button>*/}}\n{{/*                    {{ template \"modal_add\" . }}*/}}\n{{/*                </form>*/}}\n{{/*                <form class=\"form-inline my-2 my-lg-0\">*/}}\n{{/*                    <button class=\"btn btn-sm btn-outline-secondary\" type=\"button\" data-toggle=\"modal\" data-target=\"#searchModal\">*/}}\n{{/*                        найти*/}}\n{{/*                    </button>*/}}\n{{/*                    {{ template \"modal_search\" . }}*/}}\n{{/*                </form>*/}}\n                <ul class=\"navbar-nav mr-auto\">\n                </ul>\n                <ul class=\"nav justify-content-end\">\n                    <li class=\"nav-item dropdown\">\n                        <a class=\"nav-link dropdown-toggle\" href=\"#\" id=\"navbarDropdown\" role=\"button\" data-toggle=\"dropdown\" aria-haspopup=\"true\" aria-expanded=\"false\">\n                            {{ .user.GetFullName }}\n                        </a>\n                        <div class=\"dropdown-menu\" aria-labelledby=\"navbarDropdown\">\n                            <a class=\"dropdown-item\" href=\"#\">Профиль</a>\n                            <div class=\"dropdown-divider\"></div>\n                            <a class=\"dropdown-item\" href=\"#\">Выйти</a>\n                        </div>\n                    </li>\n                </ul>\n            </div>\n        </nav>\n        <div class=\"collapse\" id=\"navbarCollapseSearch\">\n            <div class=\"bg-dark p-4\">\n                <h5 class=\"text-white h4\">Найти</h5>\n                <span class=\"text-muted\">Toggleable via the navbar brand.</span>\n            </div>\n        </div>\n        <div class=\"collapse\" id=\"navbarCollapseAdd\">\n            <div class=\"bg-dark p-4\">\n                <h5 class=\"text-white h4\">Добавить</h5>\n                <span class=\"text-muted\">Toggleable via the navbar brand.</span>\n            </div>\n        </div>\n    </div>\n\n{{ end }}"),
	}
	file5 := &embedded.EmbeddedFile{
		Filename:    "base_search_form.html",
		FileModTime: time.Unix(1587907986, 0),

		Content: string("{{ define \"base_search_form\" }}\n    <h1>Search form</h1>\n{{ end }}"),
	}
	file6 := &embedded.EmbeddedFile{
		Filename:    "index.html",
		FileModTime: time.Unix(1587909560, 0),

		Content: string("{{ template \"base_header\" . }}\n{{ template \"base_navbar\" . }}\n\n<table class=\"table\">\n    <thead>\n    <tr>\n        <th scope=\"col\">ID</th>\n        <th scope=\"col\">Улица</th>\n        <th scope=\"col\">дом</th>\n        <th scope=\"col\">строение</th>\n        <th scope=\"col\">номер Квартиры</th>\n        <th scope=\"col\">состояние</th>\n        <th scope=\"col\">этаж</th>\n        <th scope=\"col\">угловая</th>\n        <th scope=\"col\">тип</th>\n        <th scope=\"col\">описание</th>\n    </tr>\n    </thead>\n    <tbody>\n    <tr>\n        <th scope=\"row\">1</th>\n        <td>Mark</td>\n        <td>Otto</td>\n        <td>@mdo</td>\n    </tr>\n    <tr>\n        <th scope=\"row\">2</th>\n        <td>Jacob</td>\n        <td>Thornton</td>\n        <td>@fat</td>\n    </tr>\n    <tr>\n        <th scope=\"row\">3</th>\n        <td>Larry</td>\n        <td>the Bird</td>\n        <td>@twitter</td>\n    </tr>\n    </tbody>\n</table>\n{{ template \"base_bottom\" . }}"),
	}
	file7 := &embedded.EmbeddedFile{
		Filename:    "modal_add.html",
		FileModTime: time.Unix(1587908217, 0),

		Content: string("{{ define \"modal_add\" }}\n    <!-- Модаль -->\n    <div class=\"modal fade\" id=\"addModal\" tabindex=\"-1\" role=\"dialog\" aria-labelledby=\"myModalLabel\">\n        <div class=\"modal-dialog modal-lg\" role=\"document\">\n            <div class=\"modal-content\">\n                <div class=\"modal-header\">\n                    <button type=\"button\" class=\"close\" data-dismiss=\"modal\" aria-label=\"Close\"><span aria-hidden=\"true\">&times;</span></button>\n                    <h4 class=\"modal-title\" id=\"myModalLabel\">Добавить квартиру</h4>\n                </div>\n                <div class=\"modal-body\">\n                    {{ template \"base_search_form\" }}\n                </div>\n                <div class=\"modal-footer\">\n                    <button type=\"button\" class=\"btn btn-default\" data-dismiss=\"modal\">Отменить</button>\n                    <button type=\"button\" class=\"btn btn-primary\">Добавить</button>\n                </div>\n            </div>\n        </div>\n    </div>\n{{end}}"),
	}
	file8 := &embedded.EmbeddedFile{
		Filename:    "modal_search.html",
		FileModTime: time.Unix(1587908266, 0),

		Content: string("{{ define \"modal_search\" }}\n    <!-- Модаль -->\n    <div class=\"modal fade\" id=\"searchModal\" tabindex=\"-1\" role=\"dialog\" aria-labelledby=\"myModalLabel\">\n        <div class=\"modal-dialog modal-lg\" role=\"document\">\n            <div class=\"modal-content\">\n                <div class=\"modal-header\">\n                    <button type=\"button\" class=\"close\" data-dismiss=\"modal\" aria-label=\"Close\"><span aria-hidden=\"true\">&times;</span></button>\n                    <h4 class=\"modal-title\" id=\"myModalLabel\">Найти квартиру</h4>\n                </div>\n                <div class=\"modal-body\">\n                    {{ template \"base_search_form\" }}\n                </div>\n                <div class=\"modal-footer\">\n                    <button type=\"button\" class=\"btn btn-default\" data-dismiss=\"modal\">Отменить</button>\n                    <button type=\"button\" class=\"btn btn-primary\">Найти</button>\n                </div>\n            </div>\n        </div>\n    </div>\n{{end}}"),
	}

	// define dirs
	dir1 := &embedded.EmbeddedDir{
		Filename:   "",
		DirModTime: time.Unix(1588825432, 0),
		ChildFiles: []*embedded.EmbeddedFile{
			file2, // "base_bottom.html"
			file3, // "base_header.html"
			file4, // "base_navbar.html"
			file5, // "base_search_form.html"
			file6, // "index.html"
			file7, // "modal_add.html"
			file8, // "modal_search.html"

		},
	}

	// link ChildDirs
	dir1.ChildDirs = []*embedded.EmbeddedDir{}

	// register embeddedBox
	embedded.RegisterEmbeddedBox(`pages/templates`, &embedded.EmbeddedBox{
		Name: `pages/templates`,
		Time: time.Unix(1588825432, 0),
		Dirs: map[string]*embedded.EmbeddedDir{
			"": dir1,
		},
		Files: map[string]*embedded.EmbeddedFile{
			"base_bottom.html":      file2,
			"base_header.html":      file3,
			"base_navbar.html":      file4,
			"base_search_form.html": file5,
			"index.html":            file6,
			"modal_add.html":        file7,
			"modal_search.html":     file8,
		},
	})
}
