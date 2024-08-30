const adaptersDescription = (adapterName) => {
  return `Deploy the Meshplay Adapter for ${adapterName} in order to enable deeper lifecycle management of ${adapterName}.`;
};

/*
 * adaptersList.name  -> name of the adapter to display on the card.
 * adaptersList.label -> used as a payload for adapter deployment (like an adapterId).
 */
export const adaptersList = {
  ISTIO: {
    name: 'Istio',
    label: 'meshplay-istio',
    imageSrc: '/static/img/istio.svg',
    description: adaptersDescription('Istio'),
    defaultPort: 10000,
    enabled: false,
    url: '',
  },
  LINKERD: {
    name: 'Linkerd',
    label: 'meshplay-linkerd',
    imageSrc: '/static/img/linkerd.svg',
    description: adaptersDescription('Linkerd'),
    defaultPort: 10001,
    enabled: false,
    url: '',
  },
  CONSUL: {
    name: 'Consul',
    label: 'meshplay-consul',
    imageSrc: '/static/img/consul.svg',
    description: adaptersDescription('Consul'),
    defaultPort: 10002,
    enabled: false,
    url: '',
  },
  NETWORK_SERVICE_MESH: {
    name: 'Network Service Mesh',
    label: 'meshplay-nsm',
    imageSrc: '/static/img/networkservicemesh.svg',
    description: adaptersDescription('Network Service Mesh'),
    defaultPort: 10004,
    enabled: false,
    url: '',
  },
  APP_MESH: {
    name: 'App Mesh',
    label: 'meshplay-app-mesh',
    imageSrc: '/static/img/app_mesh.svg',
    description: adaptersDescription('App Mesh'),
    defaultPort: 10005,
    enabled: false,
    url: '',
  },
  TRAEFIK_MESH: {
    name: 'Traefik Mesh',
    label: 'meshplay-traefik-mesh',
    imageSrc: '/static/img/traefik_mesh.svg',
    description: adaptersDescription('Traefik Mesh'),
    defaultPort: 10006,
    enabled: false,
    url: '',
  },
  KUMA: {
    name: 'Kuma',
    label: 'meshplay-kuma',
    imageSrc: '/static/img/kuma.svg',
    description: adaptersDescription('Kuma'),
    defaultPort: 10007,
    enabled: false,
    url: '',
  },
  // TODO: Need to add icon for this.
  // "meshplay-cpx": {
  //   name: "Meshplay Cpx",
  // label: "meshplay-cpx",
  //   imageSrc: "/static/img/",
  //   description: adaptersDescription("Meshplay CPX"),
  //   defaultPort: 10008,
  //   enabled: false,
  //   url: "",
  // },

  NGINX_SERVICE_MESH: {
    name: 'NGINX Service Mesh',
    label: 'meshplay-nginx-sm',
    imageSrc: '/static/img/nginx.svg',
    description: adaptersDescription('NGINX Service Mesh'),
    defaultPort: 10010,
    enabled: false,
    url: '',
  },
  CILIUM_SERVICE_MESH: {
    name: 'Cilium Service Mesh',
    label: 'meshplay-cilium',
    imageSrc: '/static/img/cilium_service_mesh.svg',
    description: adaptersDescription('Cilium Service Mesh'),
    defaultPort: 10012,
    enabled: false,
    url: '',
  },
  NIGHTHAWK: {
    name: 'Nighthawk',
    label: 'meshplay-nighthawk',
    imageSrc: '/static/img/nighthawk-logo.svg',
    description: adaptersDescription('Performance Characterization by Meshplay Nighthawk'),
    defaultPort: 10013,
    enabled: false,
    url: '',
  },
};

export const ADAPTER_STATUS = {
  ENABLED: 'ENABLED',
  DISABLED: 'DISABLED',
};
