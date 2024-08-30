export function getFallbackImageBasedOnKind(kind) {
  const fallbackComponent = {
    meshplay: 'static/img/meshplay-logo.png',
    kubernetes: 'static/img/kubernetes.svg',
  };
  return fallbackComponent[kind];
}
