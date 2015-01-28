
// Open "returnCapsule" from Go.
goatse._wrapgofunc = function(ret) {
  if (ret.error !== undefined) {
    throw new Error(ret.error);
  }
  return ret.value;
};

// Calls Go log.Println.
goatse.log = function(msg) {
  return this._wrapgofunc(this._log(msg));
};

// Like nodeJS require. Bad scope.
goatse.require = function(filename) {
  return this._wrapgofunc(this._require(filename));
};
