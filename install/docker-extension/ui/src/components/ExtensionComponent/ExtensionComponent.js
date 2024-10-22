import React, { useState, useEffect } from 'react'
import { Typography, Button, Tooltip, Grid } from '@mui/material'
import OpenInNewIcon from '@mui/icons-material/OpenInNew';
// import ConsulIcon from '../../img/SVGs/consulIcon'
// import IstioIcon from '../../img/SVGs/IstioIcon'
// import KumaIcon from '../../img/SVGs/kumaIcon'
// import LinkerdIcon from '../../img/SVGs/linkerdIcon'
import Tour from '../Walkthrough/Tour'
import PublishIcon from '../../assets/design'
import { Avatar } from '@mui/material'
// import NginxIcon from '../../img/SVGs/nginxIcon'
// import AppmeshIcon from '../../img/SVGs/appmeshIcon'
// import CiliumIcon from '../../img/SVGs/ciliumIcon'
// import TraefikIcon from '../../img/SVGs/traefikIcon'
import Meshplay from '../../img/SVGs/meshplay'
import MeshplayIcon from '../../img/meshplay-logo/CustomMeshplayLogo'
import { DockerMuiThemeProvider } from '@docker/docker-mui-theme'
import CssBaseline from '@mui/material/CssBaseline'
import { LoadComp } from '../LoadingComponent/LoadComp'
import {
  LoadingDiv,
  AccountDiv,
  ExtensionWrapper,
  LinkButton,
  ComponentWrapper,
  SectionWrapper,
  VersionText,
  LogoutButton,
  StyledButton,
  StyledLink,
  MeshModels,
  PublishCard,
} from './styledComponents'
import { MeshplayAnimation } from '../MeshplayAnimation/MeshplayAnimation'
import { randomApplicationNameGenerator } from '../../utils'
import CatalogChart from '../Catalog/Chart'
import {CatalogCard, SistentThemeProviderWithoutBaseLine} from '@khulnasoft/sistent';
import { MESHMAP, meshplayCloudUrl } from '../utils/constants';

const AuthenticatedMsg = 'Authenticated'
const UnauthenticatedMsg = 'Unauthenticated'
const proxyUrl = 'http://127.0.0.1:7877'
const httpDelete = 'DELETE'

// const adapters = {
//   APP_MESH: {
//     displayName: 'App Mesh',
//     icon: <AppmeshIcon width={40} height={40} />,
//     name: 'APP_MESH',
//   },
//   CILIUM_SERVICE_MESH: {
//     displayName: 'Cilium',
//     icon: <CiliumIcon width={40} height={40} />,
//     name: 'CILIUM_SERVICE_MESH',
//   },
//   CONSUL: {
//     displayName: 'Consul',
//     icon: <ConsulIcon width={40} height={40} />,
//     name: 'CONSUL',
//   },
//   ISTIO: {
//     displayName: 'Istio',
//     icon: <IstioIcon width={40} height={40} />,
//     name: 'ISTIO',
//   },
//   KUMA: {
//     displayName: 'Kuma',
//     icon: <KumaIcon width={40} height={40} />,
//     name: 'KUMA',
//   },
//   LINKERD: {
//     displayName: 'Linkerd',
//     icon: <LinkerdIcon width={40} height={40} />,
//     name: 'LINKERD',
//   },
//   NGINX_SERVICE_MESH: {
//     displayName: 'NGINX',
//     icon: <NginxIcon width={38} height={40} />,
//     name: 'NGINX_SERVICE_MESH',
//   },
//   TRAEFIK_MESH: {
//     displayName: 'Traefix Mesh',
//     icon: <TraefikIcon width={40} height={40} />,
//     name: 'TRAEFIK_MESH',
//   },
// }

const useThemeDetector = () => {
  const getCurrentTheme = () =>
    window.matchMedia('(prefers-color-scheme: dark)').matches
  const [isDarkTheme, setIsDarkTheme] = useState(getCurrentTheme())
  const mqListener = (e) => {
    setIsDarkTheme(e.matches)
  }

  useEffect(() => {
    const darkThemeMq = window.matchMedia('(prefers-color-scheme: dark)')
    darkThemeMq.addListener(mqListener)
    return () => darkThemeMq.removeListener(mqListener)
  }, [])
  return isDarkTheme
}

const ExtensionsComponent = () => {
  // const [switchesState, setSwitchesState] = useState(null)
  const [isHovered, setIsHovered] = useState(false)
  const isDarkTheme = useThemeDetector()
  const [isLoggedIn, setIsLoggedIn] = useState(false)
  const [user, setUser] = useState('')
  const [token, setToken] = useState()
  const [changing, isChanging] = useState(false)
  // const [emptystate, isEmptystate] = useState(true)
  // const [meshAdapters, setMeshAdapters] = useState(null)
  const [catalogDesigns, setCatalogDesigns] = useState(null)
  const [filter, setFilter] = useState(null)
  const [userDesigns, setUserDesigns] = useState(null)

  // useEffect(() => {
  //   if (meshAdapters && meshAdapters.length !== 0) {
  //     setSwitchesState(
  //       meshAdapters.map((adapter) => ({
  //         [adapter.name]: false,
  //       })),
  //     )
  //   }
  // }, [meshAdapters])
  const [meshplayVersion, setMeshplayVersion] = useState(null)

  const logout = () => {
    fetch(proxyUrl + '/token', { method: httpDelete })
      .then(console.log)
      .catch(console.error)
  }

  useEffect(() => {
    let ws = new WebSocket('ws://127.0.0.1:7877/ws')
    ws.onmessage = (msg) => {
      if (msg.data === AuthenticatedMsg) setIsLoggedIn(true)
      if (msg.data === UnauthenticatedMsg) {
        setIsLoggedIn(false)
      }
    }
    return () => ws.close()
  }, [])

  useEffect(() => {
    fetch(proxyUrl + '/token')
      .then((res) => res.text())
      .then((res) => {
        setToken(res)
        if (res !== 'null') {
          setIsLoggedIn(true)

          fetch(proxyUrl + '/api/user')
            .then((res) => res.text())
            .then((res) => {
              setUser(JSON.parse(res))
            })
            .catch(console.error)
          // fetch(proxyUrl + '/api/system/sync')
          //   .then((res) => res.json())
          //   .then((data) => {
          //     setMeshAdapters(data.meshAdapters)
          //     isEmptystate(false)
          //   })
          //   .catch(console.err)
          fetch(proxyUrl + '/api/system/version')
            .then((result) => result.text())
            .then((result) => setMeshplayVersion(JSON.parse(result)?.build))
            .catch(console.error)
          fetch(`${meshplayCloudUrl}/api/catalog/content/pattern`)
            .then((result) => result.text())
            .then((result) => {
              setCatalogDesigns(JSON.parse(result))
            })
            .catch(console.error)
          fetch(`${meshplayCloudUrl}/api/catalog/content/filter`)
            .then((result) => result.text())
            .then((result) => {
              setFilter(JSON.parse(result))
            })
            .catch(console.error)
        }
      })
      .catch(console.error)
  }, [isLoggedIn])

  useEffect(() => {
    if (user?.id) {
      fetch(`${meshplayCloudUrl}/api/content/patterns?user_id=${user?.id}`)
        .then((result) => result.text())
        .then((result) => {
          setUserDesigns(JSON.parse(result))
        })
        .catch(console.error)
    }
  }, [user])

  const onMouseOver = (e) => {
    let target = e.target.closest('div')
    target.style.transition = 'all .5s'
    target.style.transform = 'scale(1)'
  }
  const onMouseOut = (e) => {
    setIsHovered(!isHovered)
    let target = e.target.closest('div')
    target.style.transition = 'all .8s'
    target.style.transform = 'scale(1)'
  }
  const onClick = (e) => {
    let target = e.target.closest('div')
    target.style.transition = 'all .2s'
    target.style.transform = 'scale(0.8)'
    isChanging(true)
    setIsHovered(true)
  }
  // const submitConfig = (mesh, deprovision = false, meshAdapters) => {
  //   const targetMesh = meshAdapters.find((msh) => msh.name === mesh)
  //   const deployQuery = targetMesh.ops.find((op) => !op.category).key
  //   const data = {
  //     adapter: targetMesh.adapter_location,
  //     query: deployQuery,
  //     namespace: 'default',
  //     customBody: '',
  //     deleteOp: deprovision ? 'on' : '',
  //   }

  //   const params = Object.keys(data)
  //     .map(
  //       (key) => `${encodeURIComponent(key)}=${encodeURIComponent(data[key])}`,
  //     )
  //     .join('&')
  //   fetch(proxyUrl + '/api/system/adapter/operation', {
  //     credentials: 'same-origin',
  //     method: 'POST',
  //     credentials: 'include',
  //     headers: {
  //       'Content-Type': 'application/x-www-form-urlencoded;charset=UTF-8',
  //     },
  //     mode: 'no-cors',
  //     body: params,
  //   })
  //     .then(() => {
  //       window.ddClient.desktopUI.toast.success(
  //         `Request received. ${deprovision ? 'Deprovisioning' : 'Provisioning'
  //         } Service Mesh...`,
  //       )
  //     })
  //     .catch(() => {
  //       window.ddClient.desktopUI.toast.error(
  //         `Could not ${deprovision ? 'Deprovision' : 'Provision'
  //         } the Service Mesh due to some error.`,
  //       )
  //     })
  // }

  const handleImport = () => {
    const file = document.getElementById('upload-button').files[0]
    // Create a reader
    const type = String(file.name)
    const reader = new FileReader()
    reader.addEventListener('load', (event) => {
      let body = { save: true }
      let name = randomApplicationNameGenerator()
      body = JSON.stringify({
        ...body,
        application_data: { name, application_file: event.target.result },
      })
      if (!(type.includes('.yaml') || type.includes('.yml'))) {
        window.ddClient.desktopUI.toast.error(
          'Some error occured while uploading the compose file. ',
        )
        return
      }

      fetch(proxyUrl + '/api/application', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/x-www-form-urlencoded;charset=UTF-8',
        },
        body,
      })
        .then((res) => {
          window.ddClient.desktopUI.toast.success(
            'Compose file has been uploaded with name: ' + name,
          )
        })
        .catch(() =>
          window.ddClient.desktopUI.toast.error(
            'Some error occured while uploading the compose file.',
          ),
        )
    })
    reader.readAsText(file)
  }

  return (
    <DockerMuiThemeProvider>
      <CssBaseline />
      {changing && (
        <LoadingDiv sx={{ opacity: '1' }}>
          <LoadComp />
        </LoadingDiv>
      )}
      <ComponentWrapper sx={{ opacity: changing ? '0.3' : '1' }}>
        {isLoggedIn && <Tour />}
        <div
          style={{
            display: 'flex',
            justifyContent: 'space-evenly',
          }}
        >
          <div>
            <MeshplayIcon CustomColor={isDarkTheme ? 'white' : '#3C494F'} />

            <Typography sx={{ margin: 'auto', paddingTop: '1rem' }}>
              Design and operate your cloud native deployments with the
              extensible management plane, Meshplay.
            </Typography>
          </div>
        </div>

        <SectionWrapper>
          <ExtensionWrapper
            className="third-step"
            sx={{ backgroundColor: isDarkTheme ? '#393F49' : '#D7DADE' }}
          >
            <AccountDiv>
              <div style={{ marginBottom: '0.5rem' }}>
                <a
                  style={{ textDecoration: 'none' }}
                  href={
                    token &&
                    'http://localhost:9081/api/user/token?token=' +
                    token +
                    '&provider=Meshplay'
                  }
                >
                  {isLoggedIn ? (
                    <div
                      onMouseEnter={() => setIsHovered(!isHovered)}
                      onMouseLeave={onMouseOut}
                      onClick={onClick}
                      onMouseOver={onMouseOver}
                    >
                      {isHovered ? (
                        <MeshplayAnimation height={70} width={72} />
                      ) : (
                        <Meshplay height={70} width={72} />
                      )}
                    </div>
                  ) : (
                    <Meshplay height={70} width={72} />
                  )}
                </a>
                {isLoggedIn ? (
                  <LinkButton>
                    <StyledLink
                      style={{ textDecoration: 'none', color: "white" }}
                      href={
                        token &&
                        'http://localhost:9081/api/user/token?token=' +
                        token +
                        '&provider=Meshplay'
                      }
                    >
                      Launch Meshplay
                    </StyledLink>
                  </LinkButton>
                ) : ""}
              </div>
              {!isLoggedIn ? (
                <StyledButton
                  sx={{ marginTop: '0.3rem' }}
                  variant="contained"
                  disabled={isLoggedIn}
                  color="primary"
                  component="span"
                  onClick={() => {
                    window.ddClient.host.openExternal(
                      'https://meshplay.khulnasoft.com?source=aHR0cDovL2xvY2FsaG9zdDo3ODc3L3Rva2VuL3N0b3Jl&provider_version=v0.3.14',
                    )
                  }}
                >
                  Login
                </StyledButton>
              ) : ( <></> )}
            </AccountDiv>
          </ExtensionWrapper>
          {isLoggedIn && (
            <ExtensionWrapper
              className="second-step"
              sx={{ backgroundColor: isDarkTheme ? '#393F49' : '#D7DADE' }}
            >
              <AccountDiv>
                <Typography
                  sx={{ marginBottom: '2rem', whiteSpace: ' nowrap' }}
                >
                  Import Compose App
                </Typography>
                <div style={{ paddingBottom: '2rem' }}>
                  <label htmlFor="upload-button">
                    <StyledButton
                      variant="contained"
                      color="primary"
                      disabled={!isLoggedIn}
                      aria-label="Upload Button"
                      component="span"
                    >
                      <input
                        id="upload-button"
                        type="file"
                        accept=".yaml, .yml"
                        hidden
                        name="upload-button"
                        onChange={handleImport}
                      />
                      Browse...
                    </StyledButton>
                  </label>
                </div>
              </AccountDiv>
            </ExtensionWrapper>
          )}
          {!isLoggedIn ? (
            <div sx={{ display: 'none' }}></div>
          ) : (
            <div>
              <ExtensionWrapper
                className="third-step"
                sx={{ backgroundColor: isDarkTheme ? '#393F49' : '#D7DADE' }}
              >
                <AccountDiv>
                  <div
                    style={{
                      display: 'flex',
                      flexDirection: 'column',
                      alignItems: 'center',
                    }}
                  >
                    {user?.user_id && (
                      <Typography
                        sx={{
                          display: 'flex',
                          flexDirection: 'column',
                          alignItems: 'center',
                          marginBottom: '1.5rem',
                        }}
                      >
                        {user?.user_id}
                        <Avatar
                          src={user?.avatar_url}
                          sx={{
                            width: '5rem',
                            height: '5rem',
                            marginTop: '1.5rem',
                          }}
                        />
                      </Typography>
                    )}
                    <LogoutButton
                      variant="p"
                      component="p"
                      style={{
                        transform: 'none',
                      }}
                    >
                      <Button
                        onClick={logout}
                        color="secondary"
                        component="span"
                        variant="contained"
                      >
                        Logout
                      </Button>
                    </LogoutButton>
                  </div>
                </AccountDiv>
              </ExtensionWrapper>
            </div>
          )}
        </SectionWrapper>
        {isLoggedIn &&
          (<SectionWrapper>
            <CatalogChart filter={filter} pattern={catalogDesigns} isTheme={isDarkTheme} />
            <Grid sx={{ backgroundColor: isDarkTheme ? '#666A75' : '#D7DADE', borderRadius: "15px", height: "23rem", display: "flex", justifyContent: "center" }}>

              <div style={{ paddingTop: isLoggedIn ? '1.2rem' : null, margin: "10px 0" }}>
                <ExtensionWrapper
                  className="first-step"
                  sx={{
                    height: ['22rem', '17rem', '14rem'],
                  }}
                >
                  {catalogDesigns?.patterns.length > 0 ? (
                    <div>
                      <Typography variant="h5" sx={{ padding: '3rem 0 1rem 0', fontWeight: "bold" }}>
                        Designs
                      </Typography>
                      <MeshModels>
                        {
                          catalogDesigns?.patterns?.slice(0, 2).map((pattern, index) => {
                            let patternType =
                              pattern.catalog_data && pattern.catalog_data.type && pattern.catalog_data.type !== ""
                                ? pattern.catalog_data.type
                                : "deployment";
                            return (
                             <SistentThemeProviderWithoutBaseLine>                          
                              <CatalogCard
                                pattern={pattern}
                                key={`design-${index}`}
                                patternType={patternType}
                                catalog={true}
                                cardHeight='18rem'
                                cardWidth='15rem'
                              />                            
                             </SistentThemeProviderWithoutBaseLine>
                            )
                          })
                        }
                      </MeshModels>
                    </div>
                  ) : (
                    <div>
                      <Typography variant="h5" sx={{ padding: '3rem 0 1rem 0', fontWeight: "bold" }}>
                        Designs
                      </Typography>
                      <a href={user?.role_names?.includes(MESHMAP) ? "https://playground.meshplay.khulnasoft.com/extension/meshmap" : "https://play.meshplay.khulnasoft.com"} style={{ textDecoration: "none" }}>
                      <PublishCard>
                        <PublishIcon width={"60"} height={"60"} />
                        <h5>Publish your own design</h5>
                      </PublishCard>
                      </a>
                    </div>
                  )}
                </ExtensionWrapper>
              </div>
            </Grid>
          </SectionWrapper>)
        }

        <SectionWrapper>
          {isLoggedIn && (
            <div style={{ paddingTop: isLoggedIn ? '1.2rem' : null }}>
              <Tooltip title="Meshplay Server version">
                <VersionText variant="span" component="span" align="end">
                  {meshplayVersion}
                </VersionText>
              </Tooltip>
              <a
                href={`https://docs.meshplay.khulnasoft.com/project/releases/${meshplayVersion}`}
                target="_blank"
                rel="noreferrer"
                style={{ color: isDarkTheme ? 'white' : 'black' }}
              >
                <OpenInNewIcon style={{ width: '0.85rem', verticalAlign: 'middle' }} />
              </a>
            </div>
          )}
        </SectionWrapper>
      </ComponentWrapper>
    </DockerMuiThemeProvider>
  )
}

export default ExtensionsComponent
