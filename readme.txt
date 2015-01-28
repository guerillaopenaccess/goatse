Guerilla Open Access Toolkit Standard Edition
=============================================
R.E.P. Aaron Swartz.

GOATSE is a toolkit for Guerilla Open Access activists.
It is made for USB drives. Load GOATSE on to a USB hard drive and use it
from a computer with access to PDFs. GOATSE fetches PDFs, and may either
save them locally or HTTP POST to a remote server.

Goatse is CC-0: Share with everyone.

Usage
-----
GOATSE has four modes.

* In the default mode `goatse` it downloads random article with random pauses.
* In the script mode `goatse -mode=script -f <script>` it execute a Javascript file.
* In the script mode `goatse -mode=doilist -f <file>` it reads file lines for dois and fetches them.
* In `goatse --mode source` mode it dumps all source code to a folder "sourcetree".

### Default mode
Goatse in default sleeps and fetches 20 articles. It repeats forever.
Can start and then leave running. Log will be given. Files are saved
and uploaded to libgen.

### DOI list mode
List is 1 doi per line. May add more after tab which goatse ignores.
Lines with "#" ignored. Lines empty ignored. Files are saved and uploaded
to libgen.

Like:

    # Comment ignored
    10.000/foo/bar.baz	A comment about doi ignored

### GOATSE-JS API
GOATSE-JS uses Otto VM which is complete JS runtime. Contains underscore.js as `_`.

In GOATSE-JS script there is object `goatse` which contains a number
of valued functions. All of these will throw any errors.

#### goatse object

* `goatse.log(string) -> null`: log to stderr
* `goatse.httpMethod(method, url, httpargs) -> httpresponse`: Http method with random UA
* `goatse.head(url) -> httpresponse`: Http HEAD with random UA
* `goatse.get(url) -> httpresponse`: Http GET with random UA
* `goatse.post(url, body) -> httpresponse`: Http POST with random UA
* `goatse.postFile(url, fileName) -> httpresponse`: Http POST with random UA
* `goatse.query(httpresponse.body) -> QueryObj`: JQuery-like
    - `QueryObj.find(css) -> QueryObj`: Find all matching CSS select
    - `QueryObj.children() -> QueryObj`: Get all children
    - `QueryObj.attr(attr_name) -> [string]`: Get all values for attr_name in QueryObj
    - `QueryObj.text() -> [string]`: Get text content of all values in QueryObj
    - `QueryObj.value -> [string]`: Property with current values in QueryObj
* `goatse.query_find(doc, query) -> [string]`: Get all found by CSS select on doc
* `goatse.query_find(doc, query) -> [string]`: Get all found by CSS select on doc. Does not wrap doc in `html`+`head`+`body`
* `goatse.query_children(doc) -> [string]`: Get all children of doc
* `goatse.query_attr(doc, attr_name) -> string`: Get attribute of doc
* `goatse.query_text(doc) -> string`: Get text of doc
* `goatse.addResolver(resolver object) -> undefined`: Add a PDF resolver.
* `goatse.getResolver(httpresponse, [require_access]) -> resolver object`: Get a PDF resolver that matches URL. Optional is has access.
* `goatse.grabPdfByDoi(doi string) -> filename`: Gets PDF with crossref/resolver. Saves filename doi+title. Return content.
* `goatse.grabPdfsFromDoiList(dois array) -> {doi:filename}`: Calls `grabPdfByDoi` for each DOI. Failed doi are empty strings.
* `goatse.grabone(doimeta) -> filename`: Gets PDF with `goatse.grabPdfByDoi` and post to libgen.
* `goatse.doilist_mode = function(doilist) -> undefined`: Mode for reading from list file.
* `goatse.default_mode() -> undefined`: Forever fetch PDFs with random sleep.

`resolver` is object with functions interface:

    <resolver>.targetsUrl(httpresponse.url)->bool
    <resolver>.getDoi(httpresponse.body)->string
    <resolver>.haveAccess(httpresponse.body)->bool
    <resolver>.getUrlPdf(httpresponse.body)->string
    <resolver>.getPdf(httpresponse.body)->string (binary)

`httpargs` is httpplus.HttpArgs struct and has struct:

    {
        Timeout_seconds int64
        Cookies bool
        PostBody string
        PostFileByName string
        PostBodyParams map[string]string
        UrlParams map[string][]string
        Headers map[string]string
        BasicAuthUser     string
        BasicAuthPassword string
    }

Not all httpargs arguments work.
Field may be missing. `User-Agent` header will write over normal header.

`httpresponse` object has fields:

    {
        body string
        status string
        statusCode int
        contentLength int
        url string
        header { string: string }
        trailer { string: string }
    }

#### goatse.doi
Built-in to `goatse.doi` object:

* `goatse.doi.getMeta(doi string) -> doimeta`: Returns doi meta object
* `goatse.doi.memberSample(number) -> [doimeta..]`: Returns list of doi meta object
* `goatse.doi.notOpenSample(number) -> [doimeta..]`: Returns list of doi meta object
* `goatse.doi.onLibgen(doi string) -> bool`: Returns list of doi meta object
* `goatse.doi.ElsevierSample(number) -> [doimeta..]`: Sample from Elsevier
* `goatse.doi.WileySample(number) -> [doimeta..]`: Sample from Wiley
* `goatse.doi.SpringerSample(number) -> [doimeta..]`: Sample from Springer
* `goatse.doi.InformaSample(number) -> [doimeta..]`: Sample from Informa
* `goatse.doi.JSTORSample(number) -> [doimeta..]`: Sample from JSTOR
* `goatse.doi.OvidSample(number) -> [doimeta..]`: Sample from Ovid
* `goatse.doi.OUPSample(number) -> [doimeta..]`: Sample from OUP
* `goatse.doi.SageSample(number) -> [doimeta..]`: Sample from Sage
* `goatse.doi.ACSSample(number) -> [doimeta..]`: Sample from ACS
* `goatse.doi.CUPSample(number) -> [doimeta..]`: Sample from CUP
* `goatse.doi.NatureSample(number) -> [doimeta..]`: Sample from Nature
* `goatse.doi.IEEESample(number) -> [doimeta..]`: Sample from IEEE
* `goatse.doi.BMJSample(number) -> [doimeta..]`: Sample from BMJ
* `goatse.doi.IOPSample(number) -> [doimeta..]`: Sample from IOP
* `goatse.doi.AMASample(number) -> [doimeta..]`: Sample from AMA
* `goatse.doi.APSSample(number) -> [doimeta..]`: Sample from APS
* `goatse.doi.AIPSample(number) -> [doimeta..]`: Sample from AIP
* `goatse.doi.ThiemeSample(number) -> [doimeta..]`: Sample from Thieme
* `goatse.doi.WalterdeGruyterSample(number) -> [doimeta..]`: Sample from WalterdeGruyter
* `goatse.doi.RSCSample(number) -> [doimeta..]`: Sample from RSC

`doimeta` object is direct from CrossRef and has structure:

    {
        Issued {
                DateParts [][]int
                UnixMS int64
                }
        Publisher string
        License []{
                ContentVersion string
                Start   {
                        DateParts [][]int
                        UnixMS int64
                        }
                DelayInDays int
                Url string
                }
        Source string
        Url string
        Doi string
        ISSN []string
        Score float32
        Page string
        Issue string
        Deposited {
                    DateParts [][]int
                    UnixMS int64
                    }
        ReferenceCount int
        Author []{
                    Given string
                    Family string
                }
        Title []string
        Type string
        Volume string
        ContainerTitle []string
        Subtitle []string
    }

#### goatse.libgen
Built-in to `goatse.libgen` object: 

* `goatse.libgen.uploadUrl()` - Get url for upload file
* `goatse.libgen.registerUrl()` - Get url for register metadata to md5
* `goatse.libgen.getUrlPdf()` - Get url for get/check pdf with doi
* `goatse.libgen.hasPdf(doi)` - Has libgen.in this doi
* `goatse.libgen.buildRegister(doi, filesize, md5)` - Used for upload
* `goatse.libgen.upload(doi, filename)` - Upload a file as doi

Can set `goatse.libgen.urlroot` to a libgen.in mirror.

#### goatse.os
Built-in to `goatse.os` object:

* `goatse.os.argv() -> [string]`: argv
* `goatse.os.where(cmd) -> path_to_cmd`: path to cmd (relative to cwd or absolute in PATH)
* `goatse.os.getenv(varn) -> varv`: Environment variable
* `goatse.os.exec(cmd) -> cmd_output`: Execute cmd in shell
* `goatse.os.fileExists(path) -> bool`: Check file exists
* `goatse.os.rename(src, dst) -> undefined`: Rename or move
* `goatse.os.move(src, dst) -> undefined`: Rename or move
* `goatse.os.mkdirs(path) -> undefined`: Make paths
* `goatse.os.rmtree(path) -> undefined`: Remove tree

#### goatse.io
Built-in to `goatse.io` object:

* `goatse.io.loadFromFile(filename) -> string (binary?)`: Load from file
* `goatse.io.saveToFile(filename, string (binary?)) -> undefined`: Save to file
* `goatse.io.fileLines(filename) -> [string]`: Get lines of file.
* `goatse.io.md5SumFile(filename) -> string`: Get md5 of file.
* `goatse.io.fileNameSafe(filename) -> filename`: Urlencode filename
* `goatse.io.fileSize(filename) -> int`: Get filesize
* `goatse.io.fileExists(filename) -> bool`: ->goatse.os.fileExists

#### goatse.strings
Built-in to `goatse.strings` object:

* `goatse.strings.urlEscape(string) -> string`: Url escape
* `goatse.strings.urlUnescape(string) -> string`: Url Unescape

Exposes following golang `strings` functions: `contains`, `containsAny`,
`containsRune`, `count`, `equalFold`, `hasPrefix`, `hasSuffix`, `indexRune`,
`lastIndex`, `lastIndexAny`, `title`, `trim`, `trimSpace`, `trimLeft`,
`trimRight`, `trimPrefix`, `trimSuffix`

*All golang string functions return string*. Bool is `"true"` or `"false"`.
Number is `"1"`...
