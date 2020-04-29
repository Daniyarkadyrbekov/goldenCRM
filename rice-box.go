package main

import (
	"time"

	"github.com/GeertJohan/go.rice/embedded"
)

func init() {

	// define files
	file2 := &embedded.EmbeddedFile{
		Filename:    "base_bottom.html",
		FileModTime: time.Unix(1587577707, 0),

		Content: string("{{ define \"base_bottom\" }}\n<!-- jQuery (necessary for Bootstrap's JavaScript plugins) -->\n<script src=\"../sources/js/jquery.min.js\"></script>\n<!-- Include all compiled plugins (below), or include individual files as needed -->\n<script src=\"../sources/js/bootstrap.min.js\"></script>\n</body>\n</html>\n{{ end }}"),
	}
	file3 := &embedded.EmbeddedFile{
		Filename:    "base_header.html",
		FileModTime: time.Unix(1587577686, 0),

		Content: string("{{ define \"base_header\" }}\n<html lang=\"en\">\n<head>\n    <meta charset=\"utf-8\">\n    <meta http-equiv=\"X-UA-Compatible\" content=\"IE=edge\">\n    <meta name=\"viewport\" content=\"width=device-width, initial-scale=1\">\n    <!-- The above 3 meta tags *must* come first in the head; any other head content must come *after* these tags -->\n    <title>Bootstrap 101 Template</title>\n\n    <!-- Bootstrap -->\n    <link href=\"../sources/css/bootstrap.min.css\" rel=\"stylesheet\">\n\n    <!-- HTML5 shim and Respond.js for IE8 support of HTML5 elements and media queries -->\n    <!-- WARNING: Respond.js doesn't work if you view the page via file:// -->\n    <!--[if lt IE 9]>\n    <script src=\"../sources/js/html5shiv.min.js\"></script>\n    <script src=\"../sources/js/respond.min.js\"></script>\n    <![endif]-->\n</head>\n{{ end }}"),
	}
	file4 := &embedded.EmbeddedFile{
		Filename:    "base_navbar.html",
		FileModTime: time.Unix(1587909596, 0),

		Content: string("{{ define \"base_navbar\"}}\n<nav class=\"navbar navbar-default\">\n    <div class=\"container-fluid\">\n        <!-- Brand and toggle get grouped for better mobile display -->\n        <div class=\"navbar-header\">\n            <a class=\"navbar-brand\" href=\"#\">Золотой Ключ</a>\n        </div>\n\n        <!-- Collect the nav links, forms, and other content for toggling -->\n        <div class=\"collapse navbar-collapse\" id=\"bs-example-navbar-collapse-1\">\n            <form class=\"navbar-form navbar-left\">\n                <button type=\"button\" class=\"btn btn-default\" data-toggle=\"modal\" data-target=\"#addModal\">\n                    добавить\n                </button>\n                {{ template \"modal_add\" . }}\n            </form>\n            <form class=\"navbar-form navbar-left\">\n                <!-- Триггер кнопка модали-->\n                <button type=\"button\" class=\"btn btn-default\" data-toggle=\"modal\" data-target=\"#searchModal\">\n                   найти\n                </button>\n                {{ template \"modal_search\" . }}\n            </form>\n            <ul class=\"nav navbar-nav navbar-right\">\n                <li class=\"dropdown\">\n                    <a href=\"#\" class=\"dropdown-toggle\" data-toggle=\"dropdown\" role=\"button\" aria-haspopup=\"true\" aria-expanded=\"false\">{{ .user.GetFullName }} <span class=\"caret\"></span></a>\n                    <ul class=\"dropdown-menu\">\n                        <li><a href=\"#\">профиль</a></li>\n                        <li role=\"separator\" class=\"divider\"></li>\n                        <li><a href=\"#\">выйти</a></li>\n                    </ul>\n                </li>\n            </ul>\n        </div><!-- /.navbar-collapse -->\n    </div><!-- /.container-fluid -->\n</nav>\n\n{{ end }}"),
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
		DirModTime: time.Unix(1587909596, 0),
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
		Time: time.Unix(1587909596, 0),
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
