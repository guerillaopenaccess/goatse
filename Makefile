gosrcfiles = goatse.go goatse_otto_utils.go goatsejs.go goatsejs_files.go goatsejs_http.go goatsejs_os.go goatsejs_require.go goatsejs_query.go goatsejs_doiutils.go goatsejs_strings.go
srcfiles = $(gosrcfiles) httpplus doiutils goatse_js_builtins Makefile readme.txt
buildfiles = goatse goatse.exe sourcetree bindata.go

goatse: bindata.go $(srcfiles)
	go build -o goatse -i $(gosrcfiles) bindata.go

bindata.go: sourcetree
	go-bindata goatse_js_builtins sourcetree/...

sourcetree: $(srcfiles)
	mkdir sourcetree
	cp -Rf *.go goatse_js_builtins httpplus sourcetree/

clean:
	rm -Rf $(buildfiles)
