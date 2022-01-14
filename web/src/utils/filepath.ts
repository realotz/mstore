export function pathFmt(path: string) {
  const paths = path.split('/');
  for (let i = 0; i < paths.length; i++) {
    if (paths[i] == '' || paths[i] == null || typeof paths[i] == undefined) {
      paths.splice(i, 1);
      i = i - 1;
    }
  }
  return '/' + paths.join('/');
}

export function getPathInfo(path: string) {
  const paths = path.split('/');
  for (let i = 0; i < paths.length; i++) {
    if (paths[i] == '' || paths[i] == null || typeof paths[i] == undefined) {
      paths.splice(i, 1);
      i = i - 1;
    }
  }
  const id = paths[0];
  paths.splice(0, 1);
  return [id, '/' + paths.join('/')];
}
