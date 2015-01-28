// Namespace for IO file functions
goatse.io = {};

goatse.io.loadFromFile = function(filename) {
  return goatse._wrapgofunc(goatse._loadFromFile(filename));
};

goatse.io.saveToFile = function(filename, filecontents){
  return goatse._wrapgofunc(goatse._saveToFile(filename, filecontents));
};

goatse.io.fileLines = function(filename) {
    var contents = this.loadFromFile(filename);
    var lines = contents.split("\r\n");
    return lines;
};

goatse.io.md5SumFile = function(filename){
  return goatse._wrapgofunc(goatse._md5SumFile(filename));
};

goatse.io.fileNameSafe = function(s) {
    return goatse.strings.urlEscape(s);
};

goatse.io.fileSize = function(filename) {
    return goatse._wrapgofunc(goatse._fileSize(filename));
};

goatse.io.fileExists = function(filename) {
    return goatse.os.fileExists(filename);
};
