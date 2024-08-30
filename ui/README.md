# <a name="contributing">Contributing Overview</a>

Please do! Thank you for your help in improving Meshplay! :balloon:

---

<details>

  <summary><h3>Find the complete set of contributor guides at https://docs.meshplay.khulnasoft.com/project/contributing</h3></summary>

# Contributor Guide for UI component

This guide is specific to the Meshplay UI component and involves steps/methods one need to follow while working on issues related to Meshplay UI.

## How to run Meshplay UI?

Meshplay UI can be built and run in different ways. You will choose one of the two ways to build and run Meshplay UI depending upon whether you are actively developing it (whether you are creating a new feature or fixing a bug in Meshplay UI) or whether you simply need to use it as a user. Let's refer to these two methods as a _Development Build_ and _User Build._

#### 1. User Build:

For general usage, one can run Meshplay UI using Meshplay's command client `meshplayctl`, by simply running the `meshplayctl system start` command.
If you don't have the `meshplayctl` tool installed already, you can follow the [meshplayctl installation docs](https://docs.meshplay.khulnasoft.com/installation/meshplayctl) to install `meshplayctl` using various `package management` tools supported.

#### 2. Development Build:

For purposes of actively developing Meshplay UI, you first need to ensure you have npm v7 installed (`npm -v`) and if not install it (`npm -g i npm@7`), then install the dependencies using `make ui-setup` and then you can use either of the following approaches to build Meshplay UI:

1. Follow the procedure mentioned in Step 1 (User build) above, and start Meshplay UI sever on the 9081 port, and login to Meshplay UI using either of the providers mentioned on the login page. Then, to run a development server of Meshplay UI, install the dependencies using the command mentioned above, then execute `make ui` to run the livereload-nodemon server on port 3000.

   > **NOTE:** Please run the steps in order to avoid issues, as Meshplay server should be running and logged-in before accessing the development server
   > on 3000 port.

   > **NOTE:** Its strongly recommended to use either [Node Version Manager](https://github.com/nvm-sh/nvm#node-version-manager---) in linux/mac os systems or [NVM for Windows](https://github.com/coreybutler/nvm-windows#nvm-for-windows) on Windows systems so single `nvm use` / `nvm install` simplifies installing and using correct node version locallly **(v18)**, see [NVM Intro](https://github.com/nvm-sh/nvm#intro) for details. Otherwise, you might experience issues during local `npm i` similar to [4674](https://github.com/meshplay/meshplay/issues/4674) due to how optional dependencies are resolved in npm v6.

1. **`make server`** - Alternatively, build all of Meshplay UI's components upfront before serving the UI. Do this in two steps:

- Execute `make ui-setup` to initialize your environment and then `make ui-build` to build and export all Meshplay UI components.
- Execute `make server` to serve the prebuilt components.
  This method doesn't provide a live reload server. You will have to build Meshplay UI after making changes to the code and rerun these steps again in order to see those subsequent code changes reflected in the UI. > **NOTE:** If you are using this method, make sure you don't have Meshplay already running on 9081 port, using `meshplayctl`.

## Tech stack used in Meshplay-UI

- Meshplay UI uses NextJs to do server side rendering of ReactJS components. The folder `ui/components` contains all the ReactJS components involved in
  building Meshplay UI.
- MaterialUI is being used extensively for the visual components of Meshplay UI.
- Billboard.js library is being used to display various charts, and comparison graphs in Meshplay UI.

## Component naming convention

For reference and easy code search, the components are named accordingly following the rule 'Meshplay<Part of UI it involves>', for example: components
involved in rendering the Results page of Meshplay UI are named as 'MeshplayResults.js', 'MeshplayResultDialog.js', 'MeshplaySMIResults.js'. Please follow this convention if you are creating a new component.

<p style="text-align: center"><em>If you'll like to go to the main Meshplay Contributor guide <a href="../CONTRIBUTING.md">click here</a></em></p>

## Testing

- Meshplay UI uses Playwright for end-to-end testing. The tests are written in JavaScript and are located in the `ui/tests` directory.

- Install the dependencies by running the following command:

```bash

npm install

npx playwright install --with-deps

```

- To run the tests, you can use the following command:

```bash

// for running the whole test suite with all browsers
npm run test:e2e

// for running only on chromium
npm run test:e2e:chromium

// for only running fast tests
npm run test:e2e:fast

```

</details>
