// OS/System functions
goatse.os = {};

// Sleep milliseconds
goatse.os.sleep = function(ms) {
    return goatse._wrapgofunc(goatse._sleep(ms));
};

// Get int 0-n
goatse.os.randInt = function(n) {
    return goatse._wrapgofunc(goatse._randInt(n));
};

// command args starting from goatse.
goatse.os.argv = function() {
  return goatse._wrapgofunc(goatse._argv());
};

// Locate a binary in path.
goatse.os.where = function(cmd) {
  return goatse._wrapgofunc(goatse._where(cmd));
};

// Get environment variable "v" or "".
goatse.os.getenv = function(v) {
  return goatse._wrapgofunc(goatse._getenv(v));
};

// (Many args) Run program & returns output.
// Use: `goatse.exec("cmd", "foo", "bar") -> "> cmd foo bar"`
goatse.os.exec = function(executable) {
  var args = Array.prototype.slice.call(arguments);
  var exed = goatse._exec.apply(this, args);
  return goatse._wrapgofunc(exed);
};

// Is file here
goatse.os.fileExists = function(v) {
  return goatse._wrapgofunc(goatse._exists(v));
};

// Rename/Move file
goatse.os.rename = function(f) {
    return goatse._wrapgofunc(goatse._rename(f));
};
goatse.os.move = goatse.os.rename;

// Make directories
goatse.os.mkdirs = function(f) {
    return goatse._wrapgofunc(goatse._mkdirs(f));
};

// Remove directories
goatse.os.rmtree = function(f) {
    return goatse._wrapgofunc(goatse._rmtree(f));
};
