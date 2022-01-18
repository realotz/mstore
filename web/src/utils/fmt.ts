// sizefmt
export function sizeFmt(limit: number) {
  let size = '';
  if (limit < 1 * 1024) {
    //小于0.1KB，则转化成B
    size = limit + 'B';
  } else if (limit < 1024 * 1024) {
    size = (limit / 1024).toFixed(2) + 'KB';
  } else if (limit < 1024 * 1024 * 1024) {
    size = (limit / (1024 * 1024)).toFixed(2) + 'MB';
  } else {
    size = (limit / (1024 * 1024 * 1024)).toFixed(2) + 'GB';
  }
  const sizeStr = size + '';
  const index = sizeStr.indexOf('.');
  const dou = sizeStr.substr(index + 1, 2);
  if (dou == '00') {
    return sizeStr.substring(0, index) + sizeStr.substr(index + 3, 2);
  }
  return size;
}
